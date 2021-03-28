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
)

type OnHoldingNotFoundType string

const (
	zeroQuantity  = "zeroQuantity"
	deleteHolding = "deleteHolding"
)

// Config holds details when syncing
type Config struct {
	// Debug
	// Prints Debug Information
	Debug bool `toml:"debug"`
	// Destination will the aggregated information will go
	// Supported: person capital
	Destination string `toml:"destination"`
	// PriceDataSource will fetch the currency pricing data from this
	// Supported: coinbase
	PriceDataSource string `toml:"priceDataSource"`
	// OnHoldingNotFound will determine the actions performed when a holding
	// exists in the destination but not in the source
	// Available values:
	// zeroQuantity - This will not remove the holding, and instead just set the quantity to zero
	// deleteHolding - This will remove the holding
	OnHoldingNotFound OnHoldingNotFoundType `toml:"onHoldingNotFound"`
	// DestinationCurrencyAs will fetch the converted pricing data of the concurrency in the specified format
	// Data Matrix:
	// Coinbase: USD, many others look at their api
	DestinationCurrencyAs string `toml:"destinationCurrencyAs"`
	// List of crypto currency holdings with their configurations
	// Supported: coinbase
	Holdings map[string]map[string]interface{} `toml:"holdings"`
	// List of Destinations holding their configuration
	// Supported: personalcapital
	Destinations map[string]map[string]interface{} `toml:"destinations"`
}

func getConfig() Config {
	config, err := toml.LoadFile("config.toml")
	if err != nil {
		panic(err)
	}
	conf := Config{
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
