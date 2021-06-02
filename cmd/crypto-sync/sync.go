package main

import (
	"errors"
	"regexp"
	"strings"

	"github.com/pelletier/go-toml"
	"github.com/pelletier/go-toml/query"
	"github.com/will7200/go-crypto-sync/internal/common"

	"github.com/will7200/go-crypto-sync/internal/pc"
	"github.com/will7200/go-crypto-sync/internal/providers"
	_ "github.com/will7200/go-crypto-sync/internal/providers/bscscan"
	_ "github.com/will7200/go-crypto-sync/internal/providers/coinbase"
	_ "github.com/will7200/go-crypto-sync/internal/providers/coinbasepro"
	_ "github.com/will7200/go-crypto-sync/internal/providers/etherscan"
	"github.com/will7200/go-crypto-sync/pkg/personalcapital"
)

var (
	personCapitalValues, _ = query.Compile("$.destinations.personalcapital")
)

type SyncCmd struct {
	Destination string `help:"Sync to Destination"`

	Holdings       []string `arg name:"holding-accounts" help:"Holdings to fetch from"`
	TokenPatterns  []string `name:"token-pattern" help:"Only sync specific tokens"`
	FilterTokensBy string   `name:"filter-by" help:"filter property" default:"symbol"`

	AddHoldings        bool `name:"holdings-addition" help:"Allow adding new holdings in destination" default:"true"`
	UpdateHoldings     bool `name:"holdings-update" help:"Allows updating holdings in destination" default:"true"`
	SkipZeroQuantities bool `name:"skip-zero-quantities" help:"Skip updates/removals for zero quantity holdings use this if you are filtering tokens" default:"false"`
}

func (s *SyncCmd) holdingOperations() common.HoldingOperation {
	ao := common.HoldingOperation(0)
	if s.AddHoldings {
		ao |= common.AddHolding
	}
	if s.UpdateHoldings {
		ao |= common.UpdateHolding
	}
	if s.SkipZeroQuantities {
		ao |= common.SkipZeroQuantity
	}
	return ao
}

func (s *SyncCmd) Run(ctx *Context) error {
	log := ctx.SugaredLogger.Named("sync")
	switch s.Destination {
	case "personalcapital":
	case "pc":
		s.Destination = "personalcapital"
	default:
		return errors.New("unknown destination for " + s.Destination)
	}
	var (
		pricingData providers.Price
	)
	allHoldings := make(providers.Holdings, 0, 2)
	if len(s.Holdings) == 1 && s.Holdings[0] == "all" {
		s.Holdings = []string{}
		for key, _ := range ctx.Config.Holdings {
			s.Holdings = append(s.Holdings, key)
		}
	}
	for _, holding := range s.Holdings {
		log.Info("Fetching holdings from ", strings.Trim(holding, ""))
		holdingsProvider, err := providers.GetAccountProvider(holding)
		if err != nil {
			log.Warnf("Skipping holding %s since provider doesn't exist", holding)
			continue
		}
		provider, err := holdingsProvider.Open(providers.Config{Logger: ctx.Logger.Named("provider").Named(holding)}, ctx.Config.Holdings[holding])
		if err != nil {
			return err
		}
		account := provider.(providers.Account)
		uHolding, err := account.GetHoldings()
		if err != nil {
			return err
		}
		allHoldings = append(allHoldings, uHolding...)
	}

	log.Infof("setting pricing data provider to %s", ctx.Config.PriceDataSource)
	pricingData = getPricingProvider(ctx)

	allHoldings = allHoldings.MapReduce()
	if len(s.TokenPatterns) > 0 {
		log.Infof("Filtering holdings with property %s matching [%s]", s.FilterTokensBy, strings.Join(s.TokenPatterns, ","))
		pats := make([]*regexp.Regexp, len(s.TokenPatterns))
		for index, pattern := range s.TokenPatterns {
			rex, err := regexp.Compile(pattern)
			if err != nil {
				log.Fatal(err)
			}
			pats[index] = rex
		}
		allHoldings = allHoldings.FilterTokens(s.FilterTokensBy, pats...)
	}
	switch s.Destination {
	case "personalcapital":
		raw := personCapitalValues.Execute(ctx.Tree)
		email := raw.Values()[0].(*toml.Tree).Get("email").(string)
		password := raw.Values()[0].(*toml.Tree).Get("password").(string)
		accountName := raw.Values()[0].(*toml.Tree).Get("accountName").(string)
		cfg := personalcapital.NewConfiguration()
		cfg.Debug = ctx.Debug
		cfg.Logger = ctx.Logger
		pc.Sync(pc.SyncParams{
			Email:       email,
			Password:    password,
			Operations:  s.holdingOperations(),
			AccountName: accountName,
		}, cfg, allHoldings, pricingData)
	}
	return nil
}
