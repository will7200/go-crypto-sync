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

	"github.com/will7200/go-crypto-sync/internal/common"
	_ "github.com/will7200/go-crypto-sync/internal/holdings"
	_ "github.com/will7200/go-crypto-sync/internal/holdings/coinbase"
)

type Context struct {
	Debug  bool
	Tree   *toml.Tree
	Config common.Config
}

var cli struct {
	tree           *toml.Tree
	Debug          bool   `help:"Enable debug mode."`
	ConfigFileName string `help:"File to read conf from" name:"file-name" default:"config.toml"`

	Sync SyncCmd `cmd help:"Sync holdings to another account" default:"1"`
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
	conf := common.Config{
		Holdings:     map[string]map[string]interface{}{},
		Destinations: map[string]map[string]interface{}{},
	}
	err := cli.tree.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}
	if err := conf.Validate(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = ctx.Run(&Context{Debug: cli.Debug, Tree: cli.tree, Config: conf})
	ctx.FatalIfErrorf(err)
}
