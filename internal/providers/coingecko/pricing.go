package coingecko

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/HereMobilityDevelopers/mediary"
	"github.com/mitchellh/mapstructure"
	"github.com/shopspring/decimal"
	"github.com/will7200/go-crypto-sync/internal/common"
	"github.com/will7200/go-crypto-sync/internal/providers"
	"go.uber.org/zap"
)

func init() {
	for _, val := range []string{"coingecko"} {
		providers.Register(val, &Provider{name: val})
	}
}

type Data struct {
	BaseURL string `mapstructure:"baseUrl"`
	Debug   bool   `mapstructure:"debug"`
}

func (d *Data) SetDefaults(p *Provider) {
	if p.data.BaseURL == "" {
		p.data.BaseURL = "https://api.coingecko.com"
	}
}

func (d *Data) Validate() error {
	err := new(common.Errors)
	return err.AsError()
}

// Provider for Binance
type Provider struct {
	name string
	data *Data
	once sync.Once

	//
	client *http.Client
	logger *zap.SugaredLogger
}

func (p *Provider) GetExchange(currency1, currency2 string) (string, error) {
	response, err := p.client.Get(fmt.Sprintf("%s%s", p.data.BaseURL, "/api/v3/coins/"+strings.ToLower(currency1)))
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	var data map[string]interface{}
	bdata, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(bdata, &data)
	if err != nil {
		return "", err
	}
	pricingData := ((data["market_data"].(map[string]interface{}))["current_price"]).(map[string]interface{})
	currency2 = strings.ToLower(currency2)
	if _, ok := pricingData[currency2]; !ok {
		return "", errors.New("coingecko does not provide " + currency2 + " data")
	}
	return decimal.NewFromFloat(pricingData[currency2].(float64)).String(), nil
}

var _ providers.Price = &Provider{}
var _ providers.Provider = &Provider{}

func (p *Provider) Name() string {
	return p.name
}

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
		p.data = d
		d.SetDefaults(p)
		err = d.Validate()
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("invalid parameters")
	}
	p.Once()
	return p, nil
}

func (p *Provider) Once() {
	p.once.Do(func() {
		client := &http.Client{
			Timeout: 15 * time.Second,
		}
		httpClient := mediary.Init().WithPreconfiguredClient(client)

		if p.data.Debug {
			httpClient = httpClient.AddInterceptors(common.DumpRequestResponseWrappedLogger(p.logger))
		}

		p.client = httpClient.Build()
	})
}
