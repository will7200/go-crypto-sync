package etherscan

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"math/big"
	"net/http"
	"sync"
	"time"

	"github.com/HereMobilityDevelopers/mediary"
	"github.com/mitchellh/mapstructure"
	"github.com/nanmu42/etherscan-api"
	"github.com/will7200/go-crypto-sync/internal/common"
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

// Provider for etherscan
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

func AddTurtleForRateLimiter(logger *zap.SugaredLogger, retryAmount int) func(req *http.Request, handler mediary.Handler) (*http.Response, error) {
	maxLimitMessage := "Max rate limit reached"
	return func(req *http.Request, handler mediary.Handler) (*http.Response, error) {
		r, err := handler(req)
		for ra := retryAmount; ra > 0 && err == nil; ra-- {
			if r.StatusCode == http.StatusOK {
				bodyBytes, _ := ioutil.ReadAll(r.Body)
				r.Body.Close() //  must close
				r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
				if bytes.Contains(bodyBytes, []byte(maxLimitMessage)) {
					logger.Warnf("Rate limited for %s. Sleeping for one second", req.Host)
					time.Sleep(1 * time.Second)
					r, err = handler(req)
				} else {
					break
				}
			} else {
				break
			}
		}
		return r, err
	}
}

func (p *Provider) Once() {
	httpClientBase := &http.Client{
		Timeout: 15 * time.Second,
	}
	httpClient := mediary.Init().WithPreconfiguredClient(httpClientBase)

	if p.data.Debug {
		httpClient = httpClient.AddInterceptors(common.DumpRequestResponseWrappedLogger(p.logger))
	}
	httpClient = httpClient.AddInterceptors(AddTurtleForRateLimiter(p.logger, 3))
	p.client = etherscan.NewCustomized(etherscan.Customization{
		Key:     p.data.ApiKey,
		BaseURL: fmt.Sprintf(`https://%s.etherscan.io/api?`, p.data.Network.SubDomain()),
		Client:  httpClient.Build(),
		//Verbose: p.data.Debug,
	})
}

func (p *Provider) Name() string {
	return "etherscan"
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
