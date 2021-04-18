package etherscan

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"

	"github.com/will7200/go-crypto-sync/internal/common"
)

func TestCoinbaseProAPI(t *testing.T) {
	config := common.GetConfigFromTomlFile()
	data := &Data{}
	err := mapstructure.Decode(config.Holdings["coinbasepro"], &data)
	assert.NoError(t, err)
	p := Provider{}
	a, _ := p.Open(config.Holdings["coinbasepro"])
	fmt.Println(a.GetHoldings())
}

func TestDecode(t *testing.T) {
	config := common.GetConfigFromTomlFile()
	data := &Data{}
	err := mapstructure.Decode(config.Holdings["coinbasepro"], &data)
	assert.NoError(t, err)
	data.SetDefaults()
	err = data.Validate()
	assert.NoError(t, err)
	spew.Dump(data)
}
