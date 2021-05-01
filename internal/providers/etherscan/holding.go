package etherscan

import (
	"errors"
	"math"
	"math/big"
	"sync"

	"github.com/mitchellh/mapstructure"
	"github.com/nanmu42/etherscan-api"
	"go.uber.org/zap"

	"github.com/will7200/go-crypto-sync/internal/providers"
)

func init() {
	providers.Register("etherscan", &Provider{})
}

type account struct {
	Address         string
	ContractAddress string `mapstructure:"contractAddress"`
	Decimals        int
	FullName        string `mapstructure:"fullName"`
	SymbolName      string `mapstructure:"symbolName"`
}

type Data struct {
	ApiKey   string            `mapstructure:"apiKey"`
	Network  etherscan.Network `mapstructure:"network"`
	Accounts []account         `mapstructure:"account"`
	Debug    bool              `mapstructure:"debug"`
}

func (d *Data) SetDefaults() {
	if d.Network == "" {
		d.Network = etherscan.Mainnet
	}
}

func (d *Data) Validate() (err error) {
	if d.ApiKey == "" {
		err = errors.New("Invalid API Key")
	}
	if d.Network == "" {
		err = errors.New("Invalid Network")
	}
	return
}

// Coinbase Provider
type Provider struct {
	once sync.Once
	data *Data

	//
	client *etherscan.Client
	logger *zap.SugaredLogger
}

// ascertain that provider implements the account interface
var _ providers.Account = &Provider{}
var _ providers.Provider = &Provider{}

func (p *Provider) Name() string {
	return "etherscan"
}

func (p *Provider) GetHoldings() (providers.Holdings, error) {
	p.once.Do(func() {
		client := etherscan.New(p.data.Network, p.data.ApiKey)
		client.Verbose = p.data.Debug
		p.client = client
		//client.BeforeRequest = func(module, action string, param map[string]interface{}) error {
		//	// ...
		//}
		//client.AfterRequest = func(module, action string, param map[string]interface{}, outcome interface{}, requestErr error) {
		//	// ...
		//}
	})
	client := p.client
	h := make([]providers.Holding, 0, 25)
	for _, account := range p.data.Accounts {
		var (
			balance *etherscan.BigInt
			err     error
		)
		if account.ContractAddress != "" {
			p.logger.Debugf("fetching token balance at %s", account.ContractAddress)
			balance, err = client.TokenBalance(account.ContractAddress, account.Address)
		} else {
			balance, err = client.AccountBalance(account.Address)
		}
		if err != nil {
			return nil, err
		}
		h = append(h, providers.Holding{
			SymbolName:  account.SymbolName,
			FullName:    account.FullName,
			TotalShares: new(big.Float).Quo(new(big.Float).SetInt(balance.Int()), big.NewFloat(math.Pow10(account.Decimals))).String(),
		})
	}

	return h, nil
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
