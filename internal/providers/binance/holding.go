package binance

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"github.com/HereMobilityDevelopers/mediary"
	"github.com/mitchellh/mapstructure"
	"github.com/shopspring/decimal"
	"github.com/will7200/go-crypto-sync/internal/common"
	"github.com/will7200/go-crypto-sync/internal/providers"
	"github.com/will7200/go-crypto-sync/pkg/binance"
	"go.uber.org/zap"
)

func init() {
	for _, val := range []string{"binance", "binance-us"} {
		providers.Register(val, &Provider{name: val})
	}
}

type Data struct {
	ApiKey  string `mapstructure:"apiKey"`
	Secret  string `mapstructure:"secret"`
	BaseURL string `mapstructure:"baseUrl"`
	Debug   bool   `mapstructure:"debug"`
}

func (d *Data) SetDefaults(p *Provider) {
	switch p.name {
	case "binance":
		d.BaseURL = "api.binance.com"
	case "binance-us":
		d.BaseURL = "api.binance.us"
	}
}

func (d *Data) Validate() error {
	err := new(common.Errors)
	if d.ApiKey == "" {
		err.Add(errors.New("Invalid API Key"))
	}
	if d.Secret == "" {
		err.Add(errors.New("Invalid Secret Key"))
	}
	return err.AsError()
}

// Provider for Binance
type Provider struct {
	name string
	data *Data
	once sync.Once

	//
	client *binance.APIClient
	logger *zap.SugaredLogger
}

func (p *Provider) GetExchange(currency1, currency2 string) (string, error) {
	data, _, err := p.client.MarketDataApi.ApiV3AvgPriceGet(context.Background()).Symbol(currency1).Execute()
	if err != nil {
		return "", err
	}
	if currency2 != "USD" {
		return "", errors.New("Binance only supports usd as a destination currency")
	}
	if price, ok := data.GetPriceOk(); ok {
		return *price, nil
	}
	return "", errors.New("Pricing was not returned")
}

var _ providers.Account = &Provider{}
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
		config := binance.NewConfiguration()
		config.Host = p.data.BaseURL
		p.logger.Debug(config.Host)
		client := &http.Client{
			Timeout: 15 * time.Second,
		}
		httpClient := mediary.Init().WithPreconfiguredClient(client)

		if p.data.Debug {
			httpClient = httpClient.AddInterceptors(common.DumpRequestResponseWrappedLogger(p.logger))
		}

		httpClient = httpClient.AddInterceptors(AddApiCreds(p.data.ApiKey, p.data.Secret, p.logger))
		config.HTTPClient = httpClient.Build()
		p.client = binance.NewAPIClient(config)
	})
}

var excludeAPISignature = map[string]struct{}{"/api/v3/avgPrice": {}}

func AddApiCreds(apiKey, secret string, logger *zap.SugaredLogger) func(req *http.Request, handler mediary.Handler) (*http.Response, error) {
	return func(req *http.Request, handler mediary.Handler) (*http.Response, error) {
		req.Header.Add("X-MBX-APIKEY", apiKey)
		if _, found := excludeAPISignature[req.URL.Path]; found {
			r, err := handler(req)
			return r, err
		}
		logger.Debug(req.URL.Path)
		query := req.URL.Query()
		if query.Get("recvWindow") == "" {
			query.Set("recvWindow", "10000")
		}
		if query.Get("timestamp") == "" {
			query.Set("timestamp", fmt.Sprintf("%d", time.Now().UnixNano()/int64(time.Millisecond)-1000))
		}
		if query.Get("signature") == "" {
			body := query.Encode()
			if req.Body != nil {
				// Read the content
				rawBody, err := ioutil.ReadAll(req.Body)
				if err != nil {
					return nil, err
				}
				// Restore the io.ReadCloser to it's original state
				req.Body = ioutil.NopCloser(bytes.NewBuffer(rawBody))
				body = body + string(rawBody)
			}
			sig, err := generateSignature(body, secret)
			if err != nil {
				return nil, err
			}
			query.Set("signature", sig)
		}
		req.URL.RawQuery = query.Encode()
		r, err := handler(req)
		return r, err
	}
}

// generateSignature will create the sha256 HMAC using the secret key on the passed message. Value is encoded to hex value
func generateSignature(message, secret string) (string, error) {
	signature := hmac.New(sha256.New, []byte(secret))
	_, err := signature.Write([]byte(message))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(signature.Sum(nil)), nil
}

func (p *Provider) GetHoldings() (providers.Holdings, error) {
	account, _, err := p.client.TradeApi.ApiV3AccountGet(context.Background()).RecvWindow(10000).Execute()
	if err != nil {
		return nil, err
	}
	h := make([]providers.Holding, 0, 25)
	for _, balance := range account.GetBalances() {
		p1, err := decimal.NewFromString(balance.GetFree())
		if err != nil {
			return nil, err
		}
		p2, err := decimal.NewFromString(balance.GetLocked())
		if err != nil {
			return nil, err
		}
		if p1.Add(p2).Cmp(decimal.NewFromFloat(0)) == 0 {
			continue
		}
		h = append(h, providers.Holding{
			SymbolName:  balance.GetAsset(),
			FullName:    "",
			TotalShares: p1.Add(p2).String(),
		})
	}
	return h, nil
}
