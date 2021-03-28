package coinbase

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/shopspring/decimal"

	"github.com/will7200/go-crypto-sync/internal/holdings"
	"github.com/will7200/go-crypto-sync/pkg/coinbase"
)

func init() {
	holdings.Register("coinbase", &Provider{})
}

// Coinbase Provider
type Provider struct {
	apiKey    string
	apiSecret string
	debug     bool
	once      sync.Once

	//
	client *coinbase.APIClient
}

// ascertain that provider implements the account interface
var _ holdings.Account = &Provider{}
var _ holdings.Price = &Provider{}

// GetHoldings
func (p *Provider) GetHoldings() (holdings.Holdings, error) {
	p.once.Do(func() {
		config := coinbase.NewConfiguration()
		config.Debug = p.debug
		config.APIKey = coinbase.APIKey{
			Key:    p.apiKey,
			Secret: p.apiSecret,
		}
		p.client = coinbase.NewAPIClient(config)
	})
	client := p.client
	ctx := context.Background()
	var nextURI string
	h := make([]holdings.Holding, 0, 25)
loop:
	accountListRequest := client.AccountsApi.ListAccounts(ctx)
	if nextURI != "" {
		accountListRequest = accountListRequest.StartingAfter(nextURI)
	}
	coinHoldings, _, err := accountListRequest.Execute()
	if err != nil {
		return nil, err
	}
	for _, coinHolding := range coinHoldings.Data {
		amount, err := decimal.NewFromString(coinHolding.Balance.Amount)
		if err != nil {
			return nil, err
		}
		if amount.Equal(decimal.NewFromInt(0)) {
			continue
		}
		h = append(h, holdings.Holding{
			SymbolName:  coinHolding.Currency.Code,
			FullName:    coinHolding.Currency.Name,
			TotalShares: coinHolding.Balance.Amount,
		})
	}
	if p, ok := coinHoldings.GetPaginationOk(); ok {
		if p.NextStartingAfter != nil {
			log.Println("fetching next set", *p.NextStartingAfter)
			nextURI = *p.NextStartingAfter
			goto loop
		}
	}
	return h, nil
}

func (p *Provider) GetExchange(currency1, currency2 string) (string, error) {
	p.once.Do(func() {
		config := coinbase.NewConfiguration()
		config.Debug = p.debug
		config.APIKey = coinbase.APIKey{
			Key:    p.apiKey,
			Secret: p.apiSecret,
		}
		p.client = coinbase.NewAPIClient(config)
	})
	client := p.client
	ctx := context.Background()
	price, _, err := client.ExchangeRatesApi.GetExchangeRateFor(ctx, currency1, currency2).Execute()
	if err != nil {
		log.Println(err)
		return "", err
	}
	return price.Data.Amount, nil
}

// Params:
// 1. APIKey (string)
// 2. APISecret (string)
// Optional Provide a map
func (p *Provider) Open(params ...interface{}) (holdings.Account, error) {
	if len(params) == 1 {
		m, ok := params[0].(map[string]interface{})
		if !ok {
			return nil, errors.New("invalid parameters")
		}
		p.apiKey, ok = m["apiKey"].(string)
		if !ok {
			return nil, errors.New("apiKey parameter missing")
		}
		p.apiSecret, ok = m["apiSecret"].(string)
		if !ok {
			return nil, errors.New("apiSecret parameter missing")
		}
		p.debug, _ = m["debug"].(bool)
	} else if len(params) == 2 {
		for i, v := range params {
			if _, ok := v.(string); !ok {
				return nil, fmt.Errorf("invalid type(%T) for parameter index %d", v, i)
			}
		}
		p.apiKey = params[0].(string)
		p.apiSecret = params[1].(string)
	} else {
		return nil, errors.New("invalid parameters")
	}
	return p, nil
}
