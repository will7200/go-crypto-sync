package ravencoin

import (
	"errors"
	"math/big"
	"net/url"
	"sync"

	"github.com/mitchellh/mapstructure"
	"github.com/will7200/go-crypto-sync/internal/common"
	"github.com/will7200/go-crypto-sync/internal/providers"
	"go.uber.org/zap"
)

func init() {
	for _, val := range []string{"ravencoin"} {
		providers.Register(val, &Provider{name: val})
	}
}

type Data struct {
	Debug        bool   `mapstructure:"debug"`
	BaseURL      string `mapstructure:"baseUrl"`
	UserPassword string `mapstructure:"userPassword"`
}

func (d *Data) SetDefaults(p *Provider) {

}

func (d *Data) Validate() error {
	errs := new(common.Errors)
	if d.BaseURL == "" {
		errs.Add(errors.New("baseUrl Invalid"))
	}
	if _, err := url.Parse(d.BaseURL); err != nil {
		errs.Add(err)
	}
	return errs.AsError()
}

// Provider for RavenD
type Provider struct {
	name string
	data *Data
	once sync.Once

	//
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
	})
}

func (p *Provider) GetHoldings() (providers.Holdings, error) {
	h := make(providers.Holdings, 0, 1)
	data, err := readFromRPCNode(rpcNodeArgs{url: p.data.BaseURL, userPassword: p.data.UserPassword})
	if err != nil {
		return nil, err
	}
	h.AddHolding(providers.Holding{
		SymbolName:  "RVN",
		FullName:    "Ravencoin",
		TotalShares: big.NewFloat(data.Result.Balance).String(),
	})
	return h, nil
}
