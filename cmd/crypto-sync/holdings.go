package main

import (
	"strings"

	"github.com/shopspring/decimal"

	"github.com/will7200/go-crypto-sync/internal/providers"
)

type HoldingsCmd struct {
	Holdings []string `arg name:"holding-accounts" help:"Holdings to analyze"`
}

func (s *HoldingsCmd) Run(ctx *Context) error {
	log := ctx.SugaredLogger.Named("holdings")
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
		holdingsProvider, err := providers.GetProvider(holding)
		if err != nil {
			log.Warnf("Skipping holding %s since provider doesn't exist", holding)
			continue
		}
		account, err := holdingsProvider.Open(providers.Config{Logger: ctx.Logger}, ctx.Config.Holdings[holding])
		if err != nil {
			return err
		}
		uHolding, err := account.(providers.Account).GetHoldings()
		if err != nil {
			return err
		}
		allHoldings = append(allHoldings, uHolding...)
	}

	log.Infof("setting pricing data provider to %s", ctx.Config.PriceDataSource)
	pdProvider, err := providers.GetProvider(ctx.Config.PriceDataSource)
	if err != nil {
		return err
	}
	pricingData = pdProvider.(providers.Price)
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
		name := holding.CurrencyName()
		if len(name) == 0 {
			name = holding.CurrencySymbolName()
		}
		log.Infof("Holding=%s, Quantity=%s, TotalValue=%s", name, holding.TotalShares, v.String())
	}
	return nil
}
