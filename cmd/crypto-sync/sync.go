package main

import (
	"errors"
	"log"

	"github.com/pelletier/go-toml"
	"github.com/pelletier/go-toml/query"

	"github.com/will7200/go-crypto-sync/internal/holdings"
	_ "github.com/will7200/go-crypto-sync/internal/holdings/coinbase"
	_ "github.com/will7200/go-crypto-sync/internal/holdings/etherscan"
	"github.com/will7200/go-crypto-sync/internal/pc"
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
	for _, holding := range s.Holdings {
		log.Println("Fetching holdings from ", holding)
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
		pc.Sync(email, password, allHoldings.MapReduce(), pricingData)
	}
	return nil
}
