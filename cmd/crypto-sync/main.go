package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/alecthomas/kong"
	"github.com/pelletier/go-toml"
	"github.com/pelletier/go-toml/query"
	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"

	"github.com/will7200/go-crypto-sync/internal/common"
	_ "github.com/will7200/go-crypto-sync/internal/providers"
	_ "github.com/will7200/go-crypto-sync/internal/providers/coinbase"
)

var (
	zapConfig *zap.Config = &zap.Config{
		Level:             zap.NewAtomicLevelAt(zapcore.InfoLevel),
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "console",
		EncoderConfig:     zap.NewDevelopmentEncoderConfig(),
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stderr"},
		InitialFields:     nil,
	}
)

type Context struct {
	Debug         bool
	Tree          *toml.Tree
	Config        common.Config
	Logger        *zap.Logger
	SugaredLogger *zap.SugaredLogger
}

var cli struct {
	tree       *toml.Tree
	Debug      debugFlag  `help:"Enable debug mode."`
	TimeFormat timeFormat `help:"Specify output time format" name:"time-format" default:""`
	LogFormat  logFormat  `help:"Specify output log format" name:"log-format" default:"console"`
	ConfigFile string     `help:"File to read conf from" name:"config-file" default:"config.toml"`

	Sync     SyncCmd     `cmd help:"Sync holdings to another account" default:"1"`
	List     ListCmd     `cmd help:"List specific stuff like available holdings, providers"`
	Holdings HoldingsCmd `cmd help:"List holdings from various accounts"`
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
	ctx := kong.Parse(&cli, kong.Configuration(TOML, "config.toml"), kong.Bind())
	if cli.ConfigFile != "config.toml" {
		if _, err := os.Stat(cli.ConfigFile); err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
		ctx = kong.Parse(&cli, kong.Configuration(TOML, cli.ConfigFile))
	}
	conf := common.Config{
		Holdings:     map[string]map[string]interface{}{},
		Destinations: map[string]map[string]interface{}{},
	}
	err := cli.tree.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}
	if err := conf.Validate(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	logger, _ := zapConfig.Build()
	logger = logger.Named("crypto-sync")
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	err = ctx.Run(&Context{Debug: bool(cli.Debug), Tree: cli.tree, Config: conf, Logger: logger, SugaredLogger: sugar})
	ctx.FatalIfErrorf(err)
}
