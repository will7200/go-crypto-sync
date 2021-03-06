package common

import (
	"fmt"
	"strings"

	"github.com/pelletier/go-toml"
)

type OnHoldingNotFoundType string

const (
	ZeroQuantity  = "ZeroQuantity"
	DeleteHolding = "DeleteHolding"
)

func (o OnHoldingNotFoundType) IsValid() bool {
	return o == ZeroQuantity || o == DeleteHolding
}

type SymbolConfig struct {
	Global    map[string]string            `toml:"global"`
	Overrides map[string]map[string]string `toml:"override"`
}

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
	// ZeroQuantity - This will not remove the holding, and instead just set the quantity to zero
	// DeleteHolding - This will remove the holding
	OnHoldingNotFound OnHoldingNotFoundType `toml:"onHoldingNotFound"`
	// DestinationCurrencyAs will fetch the converted pricing data of the concurrency in the specified format
	// Data Matrix:
	// Coinbase: USD, many others look at their api
	DestinationCurrencyAs string `toml:"destinationCurrencyAs"`
	// List of crypto currency holdings with their configurations
	// Supported: coinbase
	Holdings map[string]map[string]interface{} `toml:"holdings"`
	// Pricing holds a list of configuration for pricing data
	Pricing map[string]map[string]interface{} `toml:"pricing"`
	// Symbols holds a map of overrides for symbols
	Symbols SymbolConfig `toml:"symbols"`
	// List of Destinations holding their configuration
	// Supported: personalcapital
	Destinations map[string]map[string]interface{} `toml:"destinations"`
}

// Validate configuration will add an error for each field not complying to its type
func (conf Config) Validate() error {
	errs := new(Errors)
	conf.OnHoldingNotFound = OnHoldingNotFoundType(strings.Title(string(conf.OnHoldingNotFound)))
	if !conf.OnHoldingNotFound.IsValid() {
		errs.Add(fmt.Errorf("invalid configuration for field: OnHoldingNotFound\n"+
			"Expected Values: %s, %s\n"+
			"Found Value: %s", ZeroQuantity, DeleteHolding, conf.OnHoldingNotFound))
	}
	return errs.AsError()
}

func GetConfigFromTomlFile(vars ...string) Config {
	filename := "config.toml"
	if len(vars) == 1 {
		filename = vars[0]
	}
	config, err := toml.LoadFile(filename)
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
