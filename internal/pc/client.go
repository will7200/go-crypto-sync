package pc

import (
	"context"
	"encoding/gob"
	"errors"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/HereMobilityDevelopers/mediary"
	"github.com/davecgh/go-spew/spew"
	"github.com/shopspring/decimal"

	"github.com/will7200/go-crypto-sync/internal/common"
	"github.com/will7200/go-crypto-sync/internal/providers"
	"github.com/will7200/go-crypto-sync/pkg/personalcapital"
)

var (
	_ = spew.Dump
)

// Sync source holdings to personal capital account
// Currently cookies will be saved in the working directory of where this command is run
func Sync(email, password string, cfg *personalcapital.Configuration, holds providers.Holdings, pricing providers.Price) {
	log := cfg.Logger.Sugar().Named("personal-capital")
	// Get saved cookies if available
	var cookies []*http.Cookie
	personalcapital.LoadSession(&cookies, "cookies.json")

	// Setup a cookie jar so the session can be saved
	cookieJar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar:     cookieJar,
		Timeout: 30 * time.Second,
	}
	httpClient := mediary.Init().WithPreconfiguredClient(client)

	if cfg.Debug {
		httpClient = httpClient.AddInterceptors(common.DumpRequestResponseWrappedLogger(cfg.Logger.Sugar().Named("personal-capital-client")))
	}
	u, err := url.Parse(cfg.Host)
	if err != nil {
		log.Fatal(err)
	}
	cfg.HTTPClient = httpClient.Build()
	client.Jar.SetCookies(u, cookies)

	apiClient := personalcapital.NewAPIClient(cfg)
	resp, err := apiClient.Authentication.Login(context.Background(), personalcapital.LoginParams{
		Username: email,
		Password: password,
	})
	// Handle two factor authentication
	if err != nil && resp != nil {
		if errors.Is(err, personalcapital.TwoFactorAuthenticationRequired) {
			err = personalcapital.TwoFactorAuthentication(client, resp.CRSF, email, password, "email")
			// csrf, err = personalcapital.Login(client, email, password, cookies)
		}
		if err != nil {
			log.Fatal(err)
		}
	} else if err != nil {
		log.Fatal(err)
	}
	apiClient.CSRF = &resp.CRSF
	pcHoldings, err := apiClient.Holdings.GetHoldings(context.Background(), &personalcapital.GetHoldingsParams{FilterUserCreated: true, FilterAccountName: "CryptoSync managed automatically"})
	if err != nil {
		log.Fatal(err)
	}
	accounts, err := apiClient.Accounts.GetAccounts(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	var account personalcapital.AccountType
	for index, account1 := range accounts.SpData.Accounts {
		if account1.Name == "CryptoSync managed automatically" && account1.AccountTypeNew == "CRYPTO_CURRENCY" {
			account = accounts.SpData.Accounts[index]
		}
	}
	m := holds.HasCurrencyMap(func(l providers.IHolding) string {
		return strings.ToLower(l.CurrencySymbolName())
	}, func(r providers.IHolding) string {
		if r.CurrencySymbolName() == "Cash" {
			return strings.ToLower(r.(personalcapital.HoldingsType).Currency)
		}
		return strings.ToLower(strings.TrimSpace(strings.Split(r.CurrencySymbolName(), "-")[0]))
	}, personalcapital.PCHoldingsToIHoldings(pcHoldings.Holdings)...)

	for key, value := range m {
		log.Debug(key, value)
		// Handle special case if is us dollar
		if key == "usd" {
			log.Info("updating PC cash amount")
			left := holds[value.LPos]
			quantity, err := decimal.NewFromString(left.TotalSharesString())
			if err != nil {
				panic(err)
			}
			_, err = apiClient.Holdings.UpdateCashAmount(context.Background(), quantity.StringFixed(2), account.UserAccountID)
			if err != nil {
				panic(err)
			}
			continue
		}
		if value.FoundBoth() {
			log.Infof("%s found in pc, updating holding", holds[value.LPos].CurrencySymbolName())
			left, right := holds[value.LPos], pcHoldings.Holdings[value.RPos]
			quantity, err := decimal.NewFromString(left.TotalSharesString())
			if err != nil {
				panic(err)
			}
			p, err := pricing.GetExchange(left.CurrencySymbolName(), "USD")
			if err != nil {
				panic(err)
			}
			pf, err := decimal.NewFromString(p)
			if err != nil {
				panic(err)
			}
			v := pf.Mul(quantity)
			d := HoldingTypeToHoldingRequest(right)
			d.PriceSource = "COINBASE"
			if quantity.GreaterThan(decimal.NewFromFloat(1)) {
				d.Quantity = quantity.StringFixed(2)
			} else {
				d.Quantity = quantity.StringFixed(18)
			}
			d.Value = v.StringFixed(2)
			if pf.GreaterThan(decimal.NewFromFloat(1)) {
				d.Price = pf.StringFixed(2)
			} else {
				d.Price = pf.StringFixed(18)
			}
			d.Description = left.SymbolName + " - " + left.TotalShares + " actual shares"
			log.Infof("Holding=%s, Quantity=%s, TotalValue=%s", left.CurrencySymbolName(), left.TotalShares, v.StringFixed(2))
			_, err = apiClient.Holdings.UpdateHoldings(context.Background(), d)
			if err != nil {
				log.Panic(err)
			}
		} else if value.LeftOnly() {
			log.Infof("%s not found in pc, create new holding\n", holds[value.LPos].CurrencySymbolName())
			quantity, err := decimal.NewFromString(holds[value.LPos].TotalSharesString())
			if err != nil {
				panic(err)
			}
			p, err := pricing.GetExchange(holds[value.LPos].CurrencySymbolName(), "USD")
			if err != nil {
				panic(err)
			}
			pf, err := decimal.NewFromString(p)
			if err != nil {
				panic(err)
			}
			v := pf.Mul(quantity)
			name := holds[value.LPos].CurrencyName()
			if len(name) < 5 {
				name = name + " - " + "Cryptocurrency"
			}
			d := &personalcapital.HoldingAddRequest{
				Ticker:        name,
				Quantity:      quantity.StringFixed(2),
				Description:   holds[value.LPos].CurrencySymbolName(),
				Source:        "USER",
				Price:         pf.StringFixed(14),
				UserAccountId: account.UserAccountID,
				Value:         v.StringFixed(2),
			}
			_, err = apiClient.Holdings.AddHolding(context.Background(), d)
			if err != nil {
				panic(err)
			}
		} else if value.RightOnly() {
			right := pcHoldings.Holdings[value.RPos]
			currencySymbol := strings.TrimSpace(strings.Split(right.CurrencySymbolName(), "-")[0])
			log.Infof("%s not found from source, setting quantity to zero\n", currencySymbol)
			quantity, err := decimal.NewFromString("0.00")
			if err != nil {
				panic(err)
			}
			p, err := pricing.GetExchange(currencySymbol, "USD")
			if err != nil {
				panic(err)
			}
			pf, err := decimal.NewFromString(p)
			if err != nil {
				panic(err)
			}
			v := pf.Mul(quantity)
			d := HoldingTypeToHoldingRequest(right)
			d.PriceSource = "COINBASE"
			d.Quantity = quantity.StringFixed(2)
			d.Value = v.StringFixed(2)
			d.Value = v.StringFixed(2)
			if pf.GreaterThan(decimal.NewFromFloat(1)) {
				d.Price = pf.StringFixed(2)
			} else {
				d.Price = pf.StringFixed(18)
			}
			_, err = apiClient.Holdings.UpdateHoldings(context.Background(), d)
			if err != nil {
				panic(err)
			}
		}
	}
	personalcapital.SaveSession(client, cfg.Host, "cookies.json")
}

func ToGlob(p interface{}, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	err = gob.NewEncoder(f).Encode(p)
	if err != nil {
		panic(err)
	}
}
