package yieldwatch

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/HereMobilityDevelopers/mediary"
	"github.com/mitchellh/mapstructure"
	"github.com/shopspring/decimal"
	"github.com/will7200/go-crypto-sync/internal/common"
	"github.com/will7200/go-crypto-sync/internal/providers"
	"github.com/will7200/go-crypto-sync/pkg/yieldwatch"
	"go.uber.org/zap"
)

func init() {
	for _, val := range []string{"yieldwatch"} {
		providers.Register(val, &Provider{name: val})
	}
}

type Data struct {
	BaseURL         string   `mapstructure:"baseUrl"`
	Debug           bool     `mapstructure:"debug"`
	Platforms       []string `mapstructure:"platforms"`
	WalletAddresses []string `mapstructure:"walletAddresses"`
}

func (d *Data) SetDefaults(p *Provider) {
	if d.BaseURL == "" {
		d.BaseURL = "www.yieldwatch.net"
	}
}

func (d *Data) Validate() error {
	err := new(common.Errors)
	return err.AsError()
}

// Provider for Coinex
type Provider struct {
	name string
	data *Data
	once sync.Once

	//
	client *yieldwatch.APIClient
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
		config := yieldwatch.NewConfiguration()
		if p.data.BaseURL != "" {
			config.Host = p.data.BaseURL
		}
		config.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:89.0) Gecko/20100101 Firefox/89.0"
		config.DefaultHeader = map[string]string{
			"User-Agent":                "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:89.0) Gecko/20100101 Firefox/89.0",
			"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
			"Accept-Language":           "en-US,en;q=0.5",
			"Upgrade-Insecure-Requests": "1",
			"Cache-Control":             "max-age=0",
		}
		p.logger.Debug(config.Host)
		client := &http.Client{
			Timeout: 15 * time.Second,
		}
		httpClient := mediary.Init().WithPreconfiguredClient(client)

		if true {
			httpClient = httpClient.AddInterceptors(common.DumpRequestResponseWrappedLogger(p.logger))
		}

		config.HTTPClient = httpClient.Build()
		p.client = yieldwatch.NewAPIClient(config)
	})
}

func (p *Provider) GetHoldings() (providers.Holdings, error) {
	h := make(providers.Holdings, 0, 25)
	platforms := strings.Join(p.data.Platforms, ",")
	for _, wallet := range p.data.WalletAddresses {
		data, _, err := p.client.DefiApi.GetDefiPortfolioAll(context.Background(), wallet).Platforms(platforms).Execute()
		if err != nil {
			return nil, err
		}
		for _, stacks := range *data.Result.PancakeSwap.Vaults.Vaults {
			p1 := decimal.NewFromFloat32(stacks.GetCurrentTokens())
			h.AddHolding(providers.Holding{
				SymbolName:  stacks.GetDepositToken(),
				FullName:    "",
				TotalShares: p1.String(),
			})
		}
		for _, value := range *data.Result.PancakeSwap.LPStaking.Vaults {
			// LP Info
			lpinfo := value.GetLPInfo()
			p1 := decimal.NewFromFloat32(lpinfo.GetDepositToken0())
			h.AddHolding(providers.Holding{
				SymbolName:  lpinfo.GetSymbolToken0(),
				FullName:    "",
				TotalShares: p1.String(),
			})
			p2 := decimal.NewFromFloat32(lpinfo.GetDepositToken1())
			h.AddHolding(providers.Holding{
				SymbolName:  lpinfo.GetSymbolToken1(),
				FullName:    "",
				TotalShares: p2.String(),
			})
			p3 := decimal.NewFromFloat32(value.GetPendingRewards())
			h.AddHolding(providers.Holding{
				SymbolName:  value.GetRewardToken(),
				FullName:    "",
				TotalShares: p3.String(),
			})
		}
	}
	return h, nil
}
