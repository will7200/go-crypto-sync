package pc

import (
	"context"
	"encoding/gob"
	"errors"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/shopspring/decimal"

	"github.com/will7200/go-crypto-sync/internal/holdings"
	"github.com/will7200/go-crypto-sync/pkg/personalcapital"
)

var (
	_ = spew.Dump
)

func Sync(email, password string, holds holdings.Holdings, pricing holdings.Price) {

	// Get saved cookies if available
	var cookies []*(http.Cookie)
	personalcapital.LoadSession(&cookies, "cookies.json")

	// Setup a cookie jar so the session can be saved
	cookieJar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: cookieJar,
	}

	cfg := personalcapital.NewConfiguration()

	u, err := url.Parse(cfg.Host)
	if err != nil {
		log.Println(err)
	}
	client.Jar.SetCookies(u, cookies)
	cfg.HTTPClient = client

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
	pcHoldings, _ := apiClient.Holdings.GetHoldings(context.Background(), &personalcapital.GetHoldingsParams{FilterUserCreated: true, FilterAccountName: "CryptoSync managed automatically"})
	accounts, err := apiClient.Accounts.GetAccounts(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	var account personalcapital.AccountType
	for index, account1 := range accounts.SpData.Accounts {
		if account1.Name == "CryptoSync managed automatically" && account1.AccountTypeNew == "CRYPTO_CURRENCY" {
			account = accounts.SpData.Accounts[index]
		}
	}

	m := holds.HasCurrencyMap(func(l holdings.IHolding) string {
		return l.CurrencyName()
	}, func(r holdings.IHolding) string {
		return strings.TrimSpace(strings.Split(r.CurrencyName(), "-")[0])
	}, personalcapital.PCHoldingsToIHoldings(pcHoldings.Holdings)...)

	for key, value := range m {
		log.Println(key, value)
		if value.Exists {
			left, right := holds[value.LPos], pcHoldings.Holdings[value.RPos]
			quantity, err := decimal.NewFromString(left.TotalSharesString())
			if err != nil {
				panic(err)
			}
			quantityFloat, _ := quantity.Float64()
			p, err := pricing.GetExchange(left.CurrencySymbolName(), "USD")
			if err != nil {
				panic(err)
			}
			pf, err := decimal.NewFromString(p)
			if err != nil {
				panic(err)
			}
			priceFloat, _ := pf.Float64()
			v := pf.Mul(quantity)
			valueFloat, _ := v.Float64()
			d := HoldingTypeToHoldingRequest(right)
			d.PriceSource = "COINBASE"
			d.Quantity = quantityFloat
			d.Value = valueFloat
			d.Price = priceFloat
			d.Description = left.SymbolName + " - " + left.TotalShares + " actual shares"
			_, err = apiClient.Holdings.UpdateHoldings(context.Background(), d)
			if err != nil {
				panic(err)
			}
		} else {
			quantity, err := decimal.NewFromString(holds[value.LPos].TotalSharesString())
			if err != nil {
				panic(err)
			}
			quantityFloat, _ := quantity.Float64()
			p, err := pricing.GetExchange(holds[value.LPos].CurrencySymbolName(), "USD")
			if err != nil {
				panic(err)
			}
			pf, err := decimal.NewFromString(p)
			if err != nil {
				panic(err)
			}
			priceFloat, _ := pf.Float64()
			v := pf.Mul(quantity)
			valueFloat, _ := v.Float64()
			name := holds[value.LPos].CurrencyName()
			if len(name) < 4 {
				name = name + " - " + "Cryptocurrency"
			}
			d := &personalcapital.HoldingAddRequest{
				Ticker:        name,
				Quantity:      quantityFloat,
				Description:   holds[value.LPos].CurrencySymbolName(),
				Source:        "USER",
				Price:         priceFloat,
				UserAccountId: account.UserAccountID,
				Value:         valueFloat,
			}
			_, err = apiClient.Holdings.AddHolding(context.Background(), d)
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
