package flatfile

import (
	"errors"
	"path/filepath"
	"sync"

	"github.com/mitchellh/mapstructure"
	"github.com/will7200/go-crypto-sync/internal/common"
	"github.com/will7200/go-crypto-sync/internal/providers"
	"go.uber.org/zap"
)

func init() {
	for _, val := range []string{"flatfile"} {
		providers.Register(val, &Provider{name: val})
	}
}

type Data struct {
	Debug    bool   `mapstructure:"debug"`
	Filename string `mapstructure:"filename"`
}

func (d *Data) SetDefaults(p *Provider) {

}

func (d *Data) Validate() error {
	err := new(common.Errors)
	if d.Filename == "" {
		err.Add(errors.New("FileName Invalid"))
	}
	return err.AsError()
}

// Provider for Flat files
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

var (
	mapColumnNameToValue = map[string]string{
		"symbol_name":  "SymbolName",
		"full_name":    "FullName",
		"total_shares": "TotalShares",
	}
)

func (p *Provider) GetHoldings() (providers.Holdings, error) {
	h := make(providers.Holdings, 0, 25)
	var (
		err     error
		records [][]string
	)
	ext := filepath.Ext(p.data.Filename)
	switch ext {
	case ".csv":
		records, err = readCsvFile(p.data.Filename)
	default:
		p.logger.Fatal("Unable handling of extension", ext)
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	mapping := make(map[int]string)
	for index, column := range records[0] {
		mapping[index] = mapColumnNameToValue[column]
	}
	p.logger.Debug(mapping)
	for _, record := range records[1:] {
		holding := providers.Holding{
			SymbolName:  "",
			FullName:    "",
			TotalShares: "",
		}
		for index, value := range record {
			switch mapping[index] {
			case "SymbolName":
				holding.SymbolName = value
			case "FullName":
				holding.FullName = value
			case "TotalShares":
				holding.TotalShares = value
			default:
				p.logger.Warnf("unknown column value %s", mapping[index])
			}
		}
		h.AddHolding(holding)
	}
	return h, nil
}
