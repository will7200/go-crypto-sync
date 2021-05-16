package bscscan

import (
	"errors"
	"math"
	"math/big"
	"net/http"
	"sync"
	"time"

	"github.com/HereMobilityDevelopers/mediary"
	"github.com/mitchellh/mapstructure"
	"github.com/nanmu42/etherscan-api"
	"go.uber.org/zap"

	"github.com/will7200/go-crypto-sync/internal/common"
	"github.com/will7200/go-crypto-sync/internal/providers"
	ietherscan "github.com/will7200/go-crypto-sync/internal/providers/etherscan"
)

func init() {
	providers.Register("bscscan", &Provider{})
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

// Provider for bscscan
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

func (p *Provider) Once() {
	httpClientBase := &http.Client{
		Timeout: 15 * time.Second,
	}
	httpClient := mediary.Init().WithPreconfiguredClient(httpClientBase)

	if p.data.Debug {
		httpClient = httpClient.AddInterceptors(common.DumpRequestResponseWrappedLogger(p.logger))
	}
	httpClient = httpClient.AddInterceptors(ietherscan.AddTurtleForRateLimiter(p.logger, 3))
	p.client = etherscan.NewCustomized(etherscan.Customization{
		Key:     p.data.ApiKey,
		BaseURL: p.data.BaseURL,
		Client:  httpClient.Build(),
		Verbose: p.data.Debug,
	})
}

func (p *Provider) Name() string {
	return "bscscan"
}

func (p *Provider) GetHoldings() (providers.Holdings, error) {
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
	p.Once()
	return p, nil
}
