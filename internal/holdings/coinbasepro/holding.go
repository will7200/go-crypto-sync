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

	"github.com/will7200/go-crypto-sync/internal/common"
	"github.com/will7200/go-crypto-sync/internal/holdings"
)

func init() {
	holdings.Register("coinbasepro", &Provider{})
}

type account struct {
	Address         string
	ContractAddress string `mapstructure:"contractAddress"`
	Decimals        int
	FullName        string `mapstructure:"fullName"`
	SymbolName      string `mapstructure:"symbolName"`
}

type Data struct {
	ApiKey     string `mapstructure:"key"`
	Secret     string `mapstructure:"secret"`
	Passphrase string `mapstructure:"passphrase"`
	BaseUrl    string `mapstructure:"base_url"`
	Debug      bool   `mapstructure:"debug"`
}

func (d *Data) SetDefaults() {
	if d.BaseUrl == "" {
		d.BaseUrl = "https://api.pro.coinbase.com"
	}
}

func (d *Data) Validate() (err error) {
	if d.ApiKey == "" {
		err = errors.New("Invalid API Key")
	}
	return
}

// Coinbase Provider
type Provider struct {
	once sync.Once
	data *Data

	//
	client *coinbasepro.Client
}

// ascertain that provider implements the account interface
var _ holdings.Account = &Provider{}

func (p *Provider) GetHoldings() (holdings.Holdings, error) {
	p.once.Do(func() {
		httpClient := mediary.Init().WithPreconfiguredClient(&http.Client{
			Timeout: 15 * time.Second,
		})
		if p.data.Debug {
			httpClient = httpClient.AddInterceptors(common.DumpRequestResponse)
		}
		client := &coinbasepro.Client{
			BaseURL:    p.data.BaseUrl,
			Key:        p.data.ApiKey,
			Passphrase: p.data.Passphrase,
			Secret:     p.data.Secret,
			HTTPClient: httpClient.Build(),
			RetryCount: 3,
		}

		p.client = client
	})
	client := p.client

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
