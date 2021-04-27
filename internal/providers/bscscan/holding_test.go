package bscscan

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"

	"github.com/will7200/go-crypto-sync/internal/common"
	"github.com/will7200/go-crypto-sync/internal/providers"
)

func TestCoinbaseProAPI(t *testing.T) {
	config := common.GetConfigFromTomlFile()
	data := &Data{}
	err := mapstructure.Decode(config.Holdings["bscscan"], &data)
	assert.NoError(t, err)
	p := Provider{}
	aa, err := p.Open(config.Holdings["bscscan"])
	a := aa.(providers.Account)
	assert.NoError(t, err)
	fmt.Println(a.GetHoldings())
}

func TestDecode(t *testing.T) {
	config := common.GetConfigFromTomlFile()
	data := &Data{}
	err := mapstructure.Decode(config.Holdings["ether"], &data)
	assert.NoError(t, err)
	data.SetDefaults()
	err = data.Validate()
	assert.NoError(t, err)
	spew.Dump(data)
}
