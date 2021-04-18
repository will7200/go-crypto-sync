package etherscan

import (
	"fmt"
	"math"
	"math/big"
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/mitchellh/mapstructure"
	"github.com/nanmu42/etherscan-api"
	"github.com/pelletier/go-toml"
	"github.com/stretchr/testify/assert"

	"github.com/will7200/go-crypto-sync/internal/common"
)

func getConfig() common.Config {
	config, err := toml.LoadFile("config.toml")
	if err != nil {
		panic(err)
	}
	conf := common.Config{
		Holdings:     map[string]map[string]interface{}{},
		Destinations: map[string]map[string]interface{}{},
	}
	err = config.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}
	return conf
}

func TestEtherScanAPI(t *testing.T) {
	config := getConfig()
	val := os.Getenv("API_KEY")
	if val == "" {
		fmt.Println("Skipping etherscan api test")
		return
	}
	client := etherscan.New(etherscan.Mainnet, os.Getenv("API_KEY"))
	client.Verbose = true
	accounts := config.Holdings["ether"]["account"].([]map[string]interface{})
	for _, account := range accounts {
		var (
			balance *etherscan.BigInt
			err     error
		)
		if ca, ok := account["contractAddress"]; ok {
			balance, err = client.TokenBalance(ca.(string), account["address"].(string))
		} else {
			balance, err = client.AccountBalance(account["address"].(string))
		}
		assert.NoError(t, err)
		assert.NoError(t, err)
		spew.Dump(new(big.Float).Quo(new(big.Float).SetInt(balance.Int()), big.NewFloat(math.Pow10(int(account["decimals"].(int64))))))
	}
}

func TestDecode(t *testing.T) {
	config := getConfig()
	data := &Data{}
	err := mapstructure.Decode(config.Holdings["ether"], &data)
	assert.NoError(t, err)
	data.SetDefaults()
	err = data.Validate()
	assert.NoError(t, err)
	spew.Dump(data)
}
