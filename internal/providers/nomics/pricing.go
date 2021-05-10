package nomics

import (
	"context"
	"errors"
	"net/http"
	"sync"
	"time"

	"github.com/HereMobilityDevelopers/mediary"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"

	"github.com/will7200/go-crypto-sync/internal/common"
	"github.com/will7200/go-crypto-sync/internal/providers"
	"github.com/will7200/go-crypto-sync/pkg/nomics"
)

func init() {
	providers.Register("nomics", &Provider{})
}

type Data struct {
	ApiKey string `mapstructure:"apiKey"`
	Debug  bool   `mapstructure:"debug"`
}

func (d *Data) SetDefaults() {

}

func (d *Data) Validate() (err error) {
	if d.ApiKey == "" {
		err = errors.New("Invalid API Key")
	}
	return
}

// Nomics Provider
type Provider struct {
	data *Data
	once sync.Once

	//
	client *nomics.APIClient
	logger *zap.SugaredLogger
}

// ascertain that provider implements the account interface
var _ providers.Price = &Provider{}
var _ providers.Provider = &Provider{}

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
	} else {
		return nil, errors.New("invalid parameters")
	}
	return p, nil
}

func (p *Provider) Once() {
	p.once.Do(func() {
		config := nomics.NewConfiguration()
		config.APIKey = nomics.APIKey{
			Key: p.data.ApiKey,
		}
		client := &http.Client{
			Timeout: 15 * time.Second,
		}
		httpClient := mediary.Init().WithPreconfiguredClient(client)

		if p.data.Debug {
			httpClient = httpClient.AddInterceptors(common.DumpRequestResponseWrappedLogger(p.logger))
		}
		config.HTTPClient = httpClient.Build()
		p.client = nomics.NewAPIClient(config)
	})
}

func (p *Provider) Name() string {
	return "nomics"
}

func (p *Provider) GetExchange(currency1, currency2 string) (string, error) {
	p.Once()
	failureCount := 0
retry:
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	resp, httpResponse, err := p.client.CurrenciesApi.GetCurrenciesTicker(ctx).Ids(currency1).QuoteCurrency(currency2).Execute()
	if httpResponse != nil && httpResponse.StatusCode == 429 && failureCount < 3 {
		failureCount += 1
		time.Sleep(time.Second)
		goto retry
	}
	if err != nil {
		return "", err
	}
	if len(resp) == 0 {
		return "", errors.New("not a valid currency with nomics")
	}
	return *resp[0].Price, nil
}
