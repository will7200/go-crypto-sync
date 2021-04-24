package etherscan

import (
	"errors"
	"math/big"
	"net/http"
	"sync"
	"time"

	"github.com/HereMobilityDevelopers/mediary"
	"github.com/mitchellh/mapstructure"
	"github.com/preichenberger/go-coinbasepro/v2"
	"go.uber.org/zap"

	"github.com/will7200/go-crypto-sync/internal/common"
	"github.com/will7200/go-crypto-sync/internal/holdings"
)

func init() {
	holdings.Register("coinbasepro", &Provider{})
}

type Portfolio struct {
	Name       string `mapstructure:"name"`
	ApiKey     string `mapstructure:"key"`
	Secret     string `mapstructure:"secret"`
	Passphrase string `mapstructure:"passphrase"`
	BaseUrl    string `mapstructure:"base_url"`
}

type Data struct {
	Debug      bool        `mapstructure:"debug"`
	Portfolios []Portfolio `mapstructure:"portfolios"`
}

func (d *Data) SetDefaults() {
	for _, p := range d.Portfolios {
		if p.BaseUrl == "" {
			p.BaseUrl = "https://api.pro.coinbase.com"
		}
	}
}

func (d *Data) Validate() error {
	err := new(common.Errors)
	for _, p := range d.Portfolios {
		if p.ApiKey == "" {
			err.Add(errors.New("Invalid API Key for portfolio " + p.Name))
		}
		if p.Secret == "" {
			err.Add(errors.New("Invalid Secret for portfolio " + p.Name))
		}
		if p.Passphrase == "" {
			err.Add(errors.New("Invalid Passphrase for portfolio " + p.Name))
		}
	}
	return err.AsError()
}

// Coinbase Provider
type Provider struct {
	once sync.Once
	data *Data

	//
	clients []*coinbasepro.Client
	logger  *zap.Logger
}

func (p *Provider) SetLogger(logger *zap.Logger) {
	p.logger = logger
}

// ascertain that provider implements the account interface
var _ holdings.Account = &Provider{}

func (p *Provider) Name() string {
	return "coinbase pro"
}

func (p *Provider) GetHoldings() (holdings.Holdings, error) {
	p.once.Do(func() {
		httpClient := mediary.Init().WithPreconfiguredClient(&http.Client{
			Timeout: 15 * time.Second,
		})
		if p.data.Debug {
			httpClient = httpClient.AddInterceptors(common.DumpRequestResponse)
		}
		clients := make([]*coinbasepro.Client, len(p.data.Portfolios))
		for i, portfolio := range p.data.Portfolios {
			clients[i] = &coinbasepro.Client{
				BaseURL:    portfolio.BaseUrl,
				Key:        portfolio.ApiKey,
				Passphrase: portfolio.Passphrase,
				Secret:     portfolio.Secret,
				HTTPClient: httpClient.Build(),
				RetryCount: 3,
			}
		}

		p.clients = clients
	})
	h := make(holdings.Holdings, 0, 10)
	for _, client := range p.clients {
		th, err := p.getHoldingsForPortfolio(client)
		if err != nil {
			return nil, err
		}
		h.AddHoldings(th)
	}
	return h, nil
}

func (p *Provider) getHoldingsForPortfolio(client *coinbasepro.Client) (holdings.Holdings, error) {
	accounts, err := client.GetAccounts()
	if err != nil {
		return nil, err
	}
	h := make(holdings.Holdings, 0, len(accounts))
	for _, account := range accounts {
		a, _, err := big.ParseFloat(account.Balance, 10, 17, big.ToNearestEven)
		if err != nil {
			panic(err)
		}
		if a.Cmp(new(big.Float).SetInt64(0)) == 1 {
			h.AddHolding(holdings.Holding{
				SymbolName:  account.Currency,
				FullName:    "",
				TotalShares: account.Balance,
			})
		}
	}
	return h, nil
}

func (p *Provider) Open(params ...interface{}) (holdings.Account, error) {
	if len(params) == 1 {
		m, ok := params[0].(map[string]interface{})
		if !ok {
			return nil, errors.New("invalid parameters")
		}
		d := &Data{}
		err := mapstructure.Decode(m, &d)
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
