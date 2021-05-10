package coinbase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/HereMobilityDevelopers/mediary"
	"github.com/mitchellh/mapstructure"
	"github.com/shopspring/decimal"
	"github.com/will7200/go-crypto-sync/internal/common"
	"go.uber.org/zap"

	"github.com/will7200/go-crypto-sync/internal/providers"
	"github.com/will7200/go-crypto-sync/pkg/coinbase"
)

func init() {
	providers.Register("coinbase", &Provider{})
}

type Data struct {
	ApiKey    string `mapstructure:"apiKey"`
	ApiSecret string `mapstructure:"apiSecret"`
	Debug     bool   `mapstructure:"debug"`
}

func (d *Data) SetDefaults() {

}

func (d *Data) Validate() (err error) {
	errs := new(common.Errors)
	if d.ApiKey == "" {
		errs.Add(errors.New("Invalid API Key"))
	}
	if d.ApiSecret == "" {
		errs.Add(errors.New("Invalid API Secret"))
	}
	return errs.AsError()
}

// Provider for Coinbase
type Provider struct {
	once sync.Once
	data *Data

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

func (p *Provider) Once() {
	config := coinbase.NewConfiguration()
	config.Debug = p.data.Debug
	config.APIKey = coinbase.APIKey{
		Key:    p.data.ApiKey,
		Secret: p.data.ApiSecret,
	}
	client := &http.Client{
		Timeout: 15 * time.Second,
	}
	httpClient := mediary.Init().WithPreconfiguredClient(client)

	if p.data.Debug {
		httpClient = httpClient.AddInterceptors(common.DumpRequestResponseWrappedLogger(p.logger))
	}
	config.HTTPClient = httpClient.Build()
	p.client = coinbase.NewAPIClient(config)
}

// GetHoldings returns all holds for coinbase
func (p *Provider) GetHoldings() (providers.Holdings, error) {
	p.Once()
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
	p.Once()
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
		d := &Data{}
		err := mapstructure.Decode(m, d)
		if err != nil {
			return nil, err
		}
		d.SetDefaults()
		err = d.Validate()
		if err != nil {
			return nil, err
		}
		p.data = d
	} else if len(params) == 2 {
		for i, v := range params {
			if _, ok := v.(string); !ok {
				return nil, fmt.Errorf("invalid type(%T) for parameter index %d", v, i)
			}
		}
		p.data = &Data{}
		p.data.ApiKey = params[0].(string)
		p.data.ApiSecret = params[1].(string)
	} else {
		return nil, errors.New("invalid parameters")
	}
	p.Once()
	return p, nil
}
