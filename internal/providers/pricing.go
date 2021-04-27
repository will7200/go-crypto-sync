package providers

import "fmt"

// Interface to providing pricing data
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
	return pp.(Price), nil
}
