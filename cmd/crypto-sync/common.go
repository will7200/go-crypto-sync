package main

import (
	"strings"

	"github.com/will7200/go-crypto-sync/internal/providers"
)

func getPricingProvider(ctx *Context) providers.Price {
	provider := ctx.Config.PriceDataSource
	if strings.Contains(provider, ",") {
		var priceProviders []providers.Price
		for _, val := range strings.Split(provider, ",") {
			pd, err := providers.OpenPricingProvider(val, providers.Config{Logger: ctx.Logger, Config: ctx.Config}, ctx.Config.Pricing[val])
			if err != nil {
				ctx.SugaredLogger.Fatal(err)
			}
			priceProviders = append(priceProviders, pd)
		}
		return &providers.FallBackPricingProvider{Pricing: priceProviders}
	} else {
		pricingData, err := providers.OpenPricingProvider(ctx.Config.PriceDataSource, providers.Config{Logger: ctx.Logger, Config: ctx.Config}, ctx.Config.Pricing[ctx.Config.PriceDataSource])
		if err != nil {
			ctx.SugaredLogger.Fatal(err)
		}
		return pricingData
	}
	ctx.SugaredLogger.Fatal("Could not process pricing provider value " + provider)
	return nil
}
