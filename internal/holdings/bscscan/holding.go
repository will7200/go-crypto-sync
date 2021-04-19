package bscscan

import (
	"errors"
	"math"
	"math/big"
	"sync"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/nanmu42/etherscan-api"

	"github.com/will7200/go-crypto-sync/internal/common"
	"github.com/will7200/go-crypto-sync/internal/holdings"
)

func init() {
	holdings.Register("bscscan", &Provider{})
}

type account struct {
	Address         string
	ContractAddress string `mapstructure:"contractAddress"`
	Decimals        int
	FullName        string `mapstructure:"fullName"`
	SymbolName      string `mapstructure:"symbolName"`
}

type Data struct {
	ApiKey   string    `mapstructure:"apiKey"`
	BaseURL  string    `mapstructure:"baseUrl"`
	Accounts []account `mapstructure:"account"`
	Debug    bool      `mapstructure:"debug"`
}

func (d *Data) SetDefaults() {
	if d.BaseURL == "" {
		d.BaseURL = "https://api.bscscan.com/api?"
	}
}

func (d *Data) Validate() error {
	errs := new(common.Errors)
	if d.ApiKey == "" {
		errs.Add(errors.New("Invalid API Key"))
	}
	if d.BaseURL == "" {
		errs.Add(errors.New("Invalid URL"))
	}
	return errs.AsError()
}

// Coinbase Provider
type Provider struct {
	once sync.Once
	data *Data

	//
	client *etherscan.Client
}

// ascertain that provider implements the account interface
var _ holdings.Account = &Provider{}

func (p *Provider) Name() string {
	return "bscscan"
}

func (p *Provider) GetHoldings() (holdings.Holdings, error) {
	p.once.Do(func() {
		client := etherscan.NewCustomized(etherscan.Customization{
			Timeout:       15 * time.Second,
			Key:           p.data.ApiKey,
			BaseURL:       p.data.BaseURL,
			Verbose:       p.data.Debug,
			BeforeRequest: nil,
			AfterRequest:  nil,
		})
		p.client = client
		//client.BeforeRequest = func(module, action string, param map[string]interface{}) error {
		//	// ...
		//}
		//client.AfterRequest = func(module, action string, param map[string]interface{}, outcome interface{}, requestErr error) {
		//	// ...
		//}
	})
	client := p.client
	h := make([]holdings.Holding, 0, 25)
	for _, account := range p.data.Accounts {
		var (
			balance *etherscan.BigInt
			err     error
		)
		if account.ContractAddress != "" {
			balance, err = client.TokenBalance(account.ContractAddress, account.Address)
		} else {
			balance, err = client.AccountBalance(account.Address)
		}
		if err != nil {
			return nil, err
		}
		h = append(h, holdings.Holding{
			SymbolName:  account.SymbolName,
			FullName:    account.FullName,
			TotalShares: new(big.Float).Quo(new(big.Float).SetInt(balance.Int()), big.NewFloat(math.Pow10(account.Decimals))).String(),
		})
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
