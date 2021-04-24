package main

import (
	"strings"

	"github.com/shopspring/decimal"

	"github.com/will7200/go-crypto-sync/internal/holdings"
)

type HoldingsCmd struct {
	Holdings []string `arg name:"holding-accounts" help:"Holdings to anaylze"`
}

func (s *HoldingsCmd) Run(ctx *Context) error {
	log := ctx.SugaredLogger.Named("holdings")
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
		log.Info("Fetching holdings from ", strings.Trim(holding, ""))
		holdingsProvider, err := holdings.GetProvider(holding)
		if err != nil {
			log.Warnf("Skipping holding %s since provider doesn't exist", holding)
			continue
		}
		account, err := holdingsProvider.Open(ctx.Config.Holdings[holding])
		if err != nil {
			return err
		}
		account.SetLogger(ctx.Logger)
		uHolding, err := account.GetHoldings()
		if err != nil {
			return err
		}
		allHoldings = append(allHoldings, uHolding...)
	}

	log.Info("setting pricing data provider to coinbase")
	pdProvider, err := holdings.GetProvider("coinbase")
	if err != nil {
		return err
	}
	pricingData = pdProvider.(holdings.Price)
	for _, holding := range allHoldings.MapReduce() {
		p, err := pricingData.GetExchange(holding.CurrencySymbolName(), "USD")
		if err != nil {
			panic(err)
		}
		pf, err := decimal.NewFromString(p)
		if err != nil {
			panic(err)
		}
		quantity, err := decimal.NewFromString(holding.TotalSharesString())
		if err != nil {
			panic(err)
		}
		v := pf.Mul(quantity)
		log.Infof("Holding=%s, Quantity=%s, TotalValue=%s", holding.CurrencyName(), holding.TotalShares, v.String())
	}
	return nil
}
