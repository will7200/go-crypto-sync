package coinbase

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/shopspring/decimal"
	"go.uber.org/zap"

	"github.com/will7200/go-crypto-sync/internal/providers"
	"github.com/will7200/go-crypto-sync/pkg/coinbase"
)

func init() {
	providers.Register("coinbase", &Provider{})
}

// Coinbase Provider
type Provider struct {
	apiKey    string
	apiSecret string
	debug     bool
	once      sync.Once

	//
	client *coinbase.APIClient
	logger *zap.SugaredLogger
}

// ascertain that provider implements the account interface
var _ providers.Account = &Provider{}
var _ providers.Price = &Provider{}
var _ providers.Provider = &Provider{}

func (p *Provider) Name() string {
	return "coinbase"
}

// GetHoldings returns all holds for coinbase
func (p *Provider) GetHoldings() (providers.Holdings, error) {
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
	h := make([]providers.Holding, 0, 25)
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
		h = append(h, providers.Holding{
			SymbolName:  coinHolding.Currency.Code,
			FullName:    coinHolding.Currency.Name,
			TotalShares: coinHolding.Balance.Amount,
		})
	}
	if page, ok := coinHoldings.GetPaginationOk(); ok {
		if page.NextStartingAfter != nil {
			p.logger.Debug("fetching next set", *page.NextStartingAfter)
			nextURI = *page.NextStartingAfter
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
	price, _, err := client.ExchangeRatesApi.GetExchangeRateFor(ctx).Currency(currency1).Execute()
	if err != nil {
		p.logger.Error(err)
		return "", err
	}
	return price.Data.Rates[currency2], nil
}

// Open the provider
// Params:
// 1. APIKey (string)
// 2. APISecret (string)
// Optional Provide a map
func (p *Provider) Open(config providers.Config, params ...interface{}) (providers.IProvider, error) {
	if config.Logger != nil {
		p.logger = config.Logger.Sugar()
	} else {
		l, _ := zap.NewDevelopmentConfig().Build()
		p.logger = l.Sugar()
	}
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
