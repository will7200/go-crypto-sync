package main

import (
	"errors"
	"strings"

	"github.com/pelletier/go-toml"
	"github.com/pelletier/go-toml/query"

	"github.com/will7200/go-crypto-sync/internal/pc"
	"github.com/will7200/go-crypto-sync/internal/providers"
	_ "github.com/will7200/go-crypto-sync/internal/providers/bscscan"
	_ "github.com/will7200/go-crypto-sync/internal/providers/coinbase"
	_ "github.com/will7200/go-crypto-sync/internal/providers/coinbasepro"
	_ "github.com/will7200/go-crypto-sync/internal/providers/etherscan"
	"github.com/will7200/go-crypto-sync/pkg/personalcapital"
)

var (
	personCapitalValues, _ = query.Compile("$.destinations.personalcapital")
)

type SyncCmd struct {
	Destination string `help:"Sync to Destination"`

	Holdings []string `arg name:"holding-accounts" help:"Holdings to fetch from"`
}

func (s *SyncCmd) Run(ctx *Context) error {
	log := ctx.SugaredLogger.Named("sync")
	switch s.Destination {
	case "personalcapital":
	case "pc":
		s.Destination = "personalcapital"
	default:
		return errors.New("unknown destination for " + s.Destination)
	}
	var (
		pricingData providers.Price
	)
	allHoldings := make(providers.Holdings, 0, 2)
	if len(s.Holdings) == 1 && s.Holdings[0] == "all" {
		s.Holdings = []string{}
		for key, _ := range ctx.Config.Holdings {
			s.Holdings = append(s.Holdings, key)
		}
	}
	for _, holding := range s.Holdings {
		log.Info("Fetching holdings from ", strings.Trim(holding, ""))
		holdingsProvider, err := providers.GetAccountProvider(holding)
		if err != nil {
			log.Warnf("Skipping holding %s since provider doesn't exist", holding)
			continue
		}
		provider, err := holdingsProvider.Open(providers.Config{Logger: ctx.Logger.Named("provider").Named(holding)}, ctx.Config.Holdings[holding])
		if err != nil {
			return err
		}
		account := provider.(providers.Account)
		uHolding, err := account.GetHoldings()
		if err != nil {
			return err
		}
		allHoldings = append(allHoldings, uHolding...)
	}

	log.Infof("setting pricing data provider to %s", ctx.Config.PriceDataSource)
	pricingData, err := providers.OpenPricingProvider(ctx.Config.PriceDataSource, providers.Config{Logger: ctx.Logger, Config: ctx.Config}, ctx.Config.Holdings[ctx.Config.PriceDataSource])
	if err != nil {
		return err
	}
	switch s.Destination {
	case "personalcapital":
		raw := personCapitalValues.Execute(ctx.Tree)
		email := raw.Values()[0].(*toml.Tree).Get("email").(string)
		password := raw.Values()[0].(*toml.Tree).Get("password").(string)
		cfg := personalcapital.NewConfiguration()
		cfg.Debug = ctx.Debug
		cfg.Logger = ctx.Logger
		pc.Sync(email, password, cfg, allHoldings.MapReduce(), pricingData)
	}
	return nil
}
