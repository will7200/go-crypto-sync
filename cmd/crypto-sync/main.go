package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/alecthomas/kong"
	"github.com/pelletier/go-toml"
	"github.com/pelletier/go-toml/query"

	_ "github.com/will7200/go-crypto-sync/internal/holdings"
	_ "github.com/will7200/go-crypto-sync/internal/holdings/coinbase"
)

type Context struct {
	Debug  bool
	Tree   *toml.Tree
	Config Config
}

var cli struct {
	tree           *toml.Tree
	Debug          bool   `help:"Enable debug mode."`
	ConfigFileName string `help:"File to read conf from" name:"file-name" default:"config.toml"`

	Sync SyncCmd `cmd help:"Sync holdings to another account" default:"1"`
}

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

// Validate configuration will add an error for each field not complying to its type
func (conf Config) Validate() (bool, []string) {
	errors := make([]string, 0)
	if conf.OnHoldingNotFound != zeroQuantity && conf.OnHoldingNotFound != deleteHolding {
		errors = append(errors, fmt.Sprintf("invalid configuration for field: OnHoldingNotFound\n"+
			"Expected Values: %s, %s", zeroQuantity, deleteHolding))
	}
	return len(errors) == 0, errors
}

// TOML returns a Resolver that retrieves values from a TOML source.
//
// Hyphens in flag names are replaced with underscores.
func TOML(r io.Reader) (kong.Resolver, error) {
	config, err := toml.LoadReader(r)
	if err != nil {
		return nil, err
	}
	cli.tree = config
	var f kong.ResolverFunc = func(context *kong.Context, parent *kong.Path, flag *kong.Flag) (interface{}, error) {
		name := strings.Replace(flag.Name, "-", "_", -1)
		raw, err := query.CompileAndExecute(fmt.Sprintf("$.%s", name), config)
		if err != nil || len(raw.Values()) == 0 {
			return nil, nil
		}
		values := raw.Values()
		if flag.IsBool() {
			return values[0], nil
		}
		if flag.IsSlice() {
			return raw.Values(), nil
		}
		return raw.Values()[0], nil
	}

	return f, nil
}

func main() {
	ctx := kong.Parse(&cli, kong.Configuration(TOML, "config.toml"))
	if cli.ConfigFileName != "config.toml" {
		if _, err := os.Stat(cli.ConfigFileName); err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
		ctx = kong.Parse(&cli, kong.Configuration(TOML, cli.ConfigFileName))
	}
	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)
	conf := Config{
		Holdings:     map[string]map[string]interface{}{},
		Destinations: map[string]map[string]interface{}{},
	}
	err := cli.tree.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}
	if ok, errors := conf.Validate(); !ok {
		for _, e := range errors {
			log.Println(e)
		}
		os.Exit(1)
	}
	err = ctx.Run(&Context{Debug: cli.Debug, Tree: cli.tree, Config: conf})
	ctx.FatalIfErrorf(err)
}
