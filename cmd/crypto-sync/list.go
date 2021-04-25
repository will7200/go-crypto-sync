package main

import (
	"fmt"

	"github.com/will7200/go-crypto-sync/internal/providers"
)

type ListCmd struct {
	Providers Providers `cmd`
}

func (s *ListCmd) Run(ctx *Context) error {
	return nil
}

type Providers struct {
	Holdings HoldingProviders `cmd`
	Pricing  PricingProviders `cmd`
}
type HoldingProviders struct{}
type PricingProviders struct{}

func (hp *HoldingProviders) Run(ctx *Context) error {
	for it := providers.GetProviderIterator(); it.HasNext(); {
		n := it.Next()
		switch p := n.(type) {
		case providers.Account:
			fmt.Println(p.Name())
		}
	}
	return nil
}

func (hp *PricingProviders) Run(ctx *Context) error {
	for it := providers.GetProviderIterator(); it.HasNext(); {
		n := it.Next()
		switch p := n.(type) {
		case providers.Price:
			fmt.Println(p.Name())
		}
	}
	return nil
}
