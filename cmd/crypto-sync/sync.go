package main

import (
	"errors"
	"log"
	"strings"

	"github.com/pelletier/go-toml"
	"github.com/pelletier/go-toml/query"

	"github.com/will7200/go-crypto-sync/internal/holdings"
	_ "github.com/will7200/go-crypto-sync/internal/holdings/bscscan"
	_ "github.com/will7200/go-crypto-sync/internal/holdings/coinbase"
	_ "github.com/will7200/go-crypto-sync/internal/holdings/coinbasepro"
	_ "github.com/will7200/go-crypto-sync/internal/holdings/etherscan"
	"github.com/will7200/go-crypto-sync/internal/pc"
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
	switch s.Destination {
	case "personalcapital":
	case "pc":
		s.Destination = "personalcapital"
	default:
		return errors.New("unknown destination for " + s.Destination)
	}
	var (
		pricingData holdings.Price
	)
	allHoldings := make(holdings.Holdings, 0, 2)
	if len(s.Holdings) == 1 && s.Holdings[0] == "all" {
		s.Holdings = []string{}
		for key, _ := range ctx.Config.Holdings {
			s.Holdings = append(s.Holdings, key)
		}
	}
	for _, holding := range s.Holdings {
		log.Println("Fetching holdings from ", strings.Trim(holding, ""))
		holdingsProvider, err := holdings.GetProvider(holding)
		if err != nil {
			return err
		}
		account, err := holdingsProvider.Open(ctx.Config.Holdings[holding])
		if err != nil {
			return err
		}
		uHolding, err := account.GetHoldings()
		if err != nil {
			return err
		}
		allHoldings = append(allHoldings, uHolding...)
	}

	log.Println("setting pricing data provider to coinbase")
	pdProvider, err := holdings.GetProvider("coinbase")
	if err != nil {
		return err
	}
	pricingData = pdProvider.(holdings.Price)
	switch s.Destination {
	case "personalcapital":
		raw := personCapitalValues.Execute(ctx.Tree)
		email := raw.Values()[0].(*toml.Tree).Get("email").(string)
		password := raw.Values()[0].(*toml.Tree).Get("password").(string)
		cfg := personalcapital.NewConfiguration()
		cfg.Debug = ctx.Debug
		pc.Sync(email, password, cfg, allHoldings.MapReduce(), pricingData)
	}
	return nil
}
