package providers

import (
	"fmt"
	"strings"
)

// Price Interface to providing pricing data
type Price interface {
	IProvider
	// GetExchange gets the value of currency1 in currency2
	GetExchange(currency1, currency2 string) (string, error)
}

func GetPricingProvider(name string) (Provider, error) {
	providerMu.RLock()
	provider, ok := providerMap[name]
	providerMu.RUnlock()

	if !ok {
		return nil, fmt.Errorf("provider: unknown provider %q (forgotten import?)", name)
	}

	_, ok = provider.(Price)
	if !ok {
		return nil, fmt.Errorf("provider %s: does not implement Account interface", name)
	}

	return provider, nil
}

func OpenPricingProvider(pricingProvider string, config Config, params ...interface{}) (Price, error) {
	p, err := GetPricingProvider(pricingProvider)
	if err != nil {
		return nil, err
	}
	pp, err := p.Open(config, params...)
	if err != nil {
		return nil, err
	}

	var newmap map[string]string
	if config.Config != nil {
		globalmap := config.Config.Symbols.Global
		localmap, ok := config.Config.Symbols.Overrides[pricingProvider]
		if ok && len(localmap) > 0 {
			newmap = make(map[string]string, len(localmap)+len(globalmap))
			for key, value := range globalmap {
				newmap[key] = value
			}
			for key, value := range localmap {
				newmap[key] = value
			}
		} else if len(globalmap) > 0 {
			newmap = make(map[string]string, len(localmap)+len(globalmap))
			for key, value := range globalmap {
				newmap[key] = value
			}
		}
		if newmap != nil {
			return &ConvertSymbolPricingProvider{
				Price:      pp.(Price),
				LookupName: newmap,
			}, nil
		}
	}
	return pp.(Price), nil
}

type FallBackPricingProvider struct {
	Pricing []Price
}

var _ Price = &FallBackPricingProvider{}

func (f *FallBackPricingProvider) Name() string {
	b := new(strings.Builder)
	for i, p := range f.Pricing {
		b.WriteString(p.Name())
		if i != len(f.Pricing)-1 {
			b.WriteString(",")
		}
	}
	return b.String()
}

func (f *FallBackPricingProvider) GetExchange(currency1, currency2 string) (res string, err error) {
	for _, p := range f.Pricing {
		res, err = p.GetExchange(currency1, currency2)
		if err == nil {
			return res, err
		}
	}
	return res, err
}

type ConvertSymbolPricingProvider struct {
	Price
	LookupName map[string]string
}

var _ Price = &ConvertSymbolPricingProvider{}

func (c *ConvertSymbolPricingProvider) Name() string {
	return c.Price.Name()
}

func (c *ConvertSymbolPricingProvider) GetExchange(currency1, currency2 string) (res string, err error) {
	override := currency1
	if val, ok := c.LookupName[override]; ok {
		override = val
	}
	res, err = c.Price.GetExchange(override, currency2)
	return
}
