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
	mappedResults := holds.HasCurrencyMap(func(l providers.IHolding) string {
		return strings.ToLower(l.CurrencySymbolName())
	}, func(r providers.IHolding) string {
		if r.CurrencySymbolName() == "Cash" {
			return strings.ToLower(r.(personalcapital.HoldingsType).Currency)
		}
		return strings.ToLower(strings.TrimSpace(strings.Split(r.CurrencySymbolName(), "-")[0]))
	}, personalcapital.PCHoldingsToIHoldings(pcHoldings.Holdings)...)

	for key, value := range mappedResults {
		log.Debug(key, value)
		// Handle by key
		{
			switch strings.ToLower(key) {
			// Handle special case if is us dollar
			case "usd", "us dollar", "cash":
				log.Info("updating PC cash amount")
				left := holds[value.LPos]
				quantity, err := decimal.NewFromString(left.TotalSharesString())
				if err != nil {
					log.Fatal(err)
				}
				_, err = apiClient.Holdings.UpdateCashAmount(context.Background(), quantity.StringFixed(2), account.UserAccountID)
				if err != nil {
					log.Fatal(err)
				}
				continue
			}
		}

		var (
			holdingRequest interface{}
			target         providers.IHolding
		)
		switch value.Result() {
		case providers.ExistsInBoth:
			left, right := holds[value.LPos], pcHoldings.Holdings[value.RPos]
			log.Infof("%s found in pc, updating holding", left.CurrencySymbolName())
			holdingRequest = HoldingTypeToHoldingRequest(right)
			target = left
		case providers.ExistsInOriginationOnly:
			left := holds[value.LPos]
			log.Infof("%s not found in pc, create new holding", left.CurrencySymbolName())
			if err != nil {
				log.Fatal(err)
			}
			ticker := left.CurrencyName()
			if len(ticker) == 0 {
				ticker = left.CurrencySymbolName()
			}
			if len(ticker) < 5 {
				ticker = ticker + " - " + "Cryptocurrency"
			}
			holdingRequest = &personalcapital.HoldingAddRequest{
				Ticker:        ticker,
				Description:   left.CurrencySymbolName(),
				Source:        "USER",
				UserAccountId: account.UserAccountID,
			}
			target = left
		case providers.ExistsInTargetOnly:
			right := pcHoldings.Holdings[value.RPos]
			currencySymbol := strings.TrimSpace(strings.Split(right.CurrencySymbolName(), "-")[0])
			log.Infof("%s not found from source, setting quantity to zero\n", currencySymbol)
			if err != nil {
				log.Fatal(err)
			}
			holdingRequest = HoldingTypeToHoldingRequest(right)
			target = right
			right.Quantity = 0
		default:
			log.Fatal("Wow never expected to get here")
			return
		}

		p, err := pricing.GetExchange(target.CurrencySymbolName(), "USD")
		if err != nil {
			log.Fatal(err)
		}
		price, err := decimal.NewFromString(p)
		if err != nil {
			log.Fatal(err)
		}
		quantity, err := decimal.NewFromString(target.TotalSharesString())
		if err != nil {
			log.Fatal(err)
		}
		v := price.Mul(quantity)
		log.Infof("Holding=%s, Quantity=%s, TotalValue=%s", target.CurrencyName(), quantity.StringFixed(18), v.String())

		switch value.Result() {
		case providers.ExistsInBoth, providers.ExistsInTargetOnly:
			d := holdingRequest.(personalcapital.HoldingsUpdateRequest)
			d.Value = v.StringFixed(2)
			if quantity.GreaterThan(decimal.NewFromFloat(1)) {
				d.Quantity = quantity.StringFixed(2)
			} else {
				d.Quantity = quantity.StringFixed(18)
			}
			if price.GreaterThan(decimal.NewFromFloat(1)) {
				d.Price = price.StringFixed(2)
			} else {
				d.Price = price.StringFixed(18)
			}
			_, err = apiClient.Holdings.UpdateHoldings(context.Background(), d)
			if err != nil {
				log.Fatal(err)
			}
		case providers.ExistsInOriginationOnly:
			d := holdingRequest.(*personalcapital.HoldingAddRequest)
			d.Value = v.StringFixed(2)
			if quantity.GreaterThan(decimal.NewFromFloat(1)) {
				d.Quantity = quantity.StringFixed(2)
			} else {
				d.Quantity = quantity.StringFixed(18)
			}
			if price.GreaterThan(decimal.NewFromFloat(1)) {
				d.Price = price.StringFixed(2)
			} else {
				d.Price = price.StringFixed(18)
			}
			_, err = apiClient.Holdings.AddHolding(context.Background(), d)
			if err != nil {
				log.Fatal(err)
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
