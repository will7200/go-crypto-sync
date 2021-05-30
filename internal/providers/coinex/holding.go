package coinex

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/HereMobilityDevelopers/mediary"
	"github.com/mitchellh/mapstructure"
	"github.com/shopspring/decimal"
	"github.com/will7200/go-crypto-sync/internal/common"
	"github.com/will7200/go-crypto-sync/internal/providers"
	"github.com/will7200/go-crypto-sync/pkg/coinex"
	"go.uber.org/zap"
)

func init() {
	for _, val := range []string{"coinex"} {
		providers.Register(val, &Provider{name: val})
	}
}

type Data struct {
	AccessId  string `mapstructure:"accessId"`
	SecretKey string `mapstructure:"secretKey"`
	BaseURL   string `mapstructure:"baseUrl"`
	Debug     bool   `mapstructure:"debug"`
}

func (d *Data) SetDefaults(p *Provider) {
}

func (d *Data) Validate() error {
	err := new(common.Errors)
	if d.AccessId == "" {
		err.Add(errors.New("Invalid Access Id"))
	}
	if d.SecretKey == "" {
		err.Add(errors.New("Invalid Secret Key"))
	}
	return err.AsError()
}

// Provider for Coinex
type Provider struct {
	name string
	data *Data
	once sync.Once

	//
	client *coinex.APIClient
	logger *zap.SugaredLogger
}

var _ providers.Account = &Provider{}
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
		config := coinex.NewConfiguration()
		if p.data.BaseURL != "" {
			config.Host = p.data.BaseURL
		}
		p.logger.Debug(config.Host)
		client := &http.Client{
			Timeout: 15 * time.Second,
		}
		httpClient := mediary.Init().WithPreconfiguredClient(client)

		if true {
			httpClient = httpClient.AddInterceptors(common.DumpRequestResponseWrappedLogger(p.logger))
		}

		httpClient = httpClient.AddInterceptors(AddAuthorizationHeader(p.data.AccessId, p.data.SecretKey, p.logger))
		config.HTTPClient = httpClient.Build()
		p.client = coinex.NewAPIClient(config)
	})
}

func AddAuthorizationHeader(accessId, secretKey string, logger *zap.SugaredLogger) func(req *http.Request, handler mediary.Handler) (*http.Response, error) {
	return func(req *http.Request, handler mediary.Handler) (*http.Response, error) {
		query := req.URL.Query()
		if req.Method == "GET" || req.Method == "DELETE" {
			if query.Get("access_id") != "" {
				// this means a signature was required
				urlParameters := url.Values{}
				for k, v := range req.URL.Query() {
					value := fmt.Sprintf("%v", v)
					urlParameters.Set(k, value)
				}
				urlParameters.Set("access_id", accessId)
				urlParameters.Set("tonce", fmt.Sprintf("%d", time.Now().UnixNano()/1e6))

				queryParamsString := urlParameters.Encode()
				toEncodeParamsString := queryParamsString + "&secret_key=" + secretKey
				s, err := generateSignature(toEncodeParamsString)
				if err != nil {
					return nil, err
				}
				req.Header.Set("authorization", s)
				req.URL.RawQuery = queryParamsString
				logger.Debug(queryParamsString)
			}
		}
		r, err := handler(req)
		return r, err
	}
}

// generateSignature will create the sha256 HMAC using the secret key on the passed message. Value is encoded to hex value
func generateSignature(message string) (string, error) {
	w := md5.New()
	_, err := io.WriteString(w, message)
	if err != nil {
		return "", err
	}
	return strings.ToUpper(hex.EncodeToString(w.Sum(nil))), nil
}

func (p *Provider) GetHoldings() (providers.Holdings, error) {
	account, _, err := p.client.AccountApi.InquireAccountInfo(context.Background()).AccessId(p.data.AccessId).Authorization("").Tonce(0).Execute()
	if err != nil {
		return nil, err
	}
	h := make([]providers.Holding, 0, 25)
	data, ok := account.GetDataOk()
	if !ok {
		return nil, errors.New("data is nil whatttt")
	}
	for coin, subData := range *data {
		p1, err := decimal.NewFromString(subData["available"].(string))
		if err != nil {
			return nil, err
		}
		p2, err := decimal.NewFromString(subData["frozen"].(string))
		if err != nil {
			return nil, err
		}
		h = append(h, providers.Holding{
			SymbolName:  coin,
			FullName:    "",
			TotalShares: p1.Add(p2).String(),
		})
	}
	return h, nil
}
