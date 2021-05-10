package personalcapital

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/will7200/go-crypto-sync/internal/providers"
)

type Holdings service

const (
	GetHoldingsEndpoint = "/api/invest/getHoldings"
)

type (
	GetHoldingsParams struct {
		FilterUserCreated bool
		FilterAccountName string
	}

	GetHoldingsResponse struct {
		SpData
	}

	GetHoldingsResponseExpecting struct {
		SpHeader SpHeader `json:"spHeader"`
		SpData   SpData   `json:"spData"`
	}

	HoldingsType struct {
		Cusip                        string  `json:"cusip,omitempty"`
		AccountName                  string  `json:"accountName"`
		Description                  string  `json:"description"`
		Source                       string  `json:"source"`
		MarketType                   int     `json:"marketType"`
		OriginalTicker               string  `json:"originalTicker,omitempty"`
		OriginalCusip                string  `json:"originalCusip,omitempty"`
		HoldingType                  string  `json:"holdingType"`
		Price                        float64 `json:"price"`
		HoldingPercentage            float64 `json:"holdingPercentage"`
		Value                        float64 `json:"value"`
		OriginalDescription          string  `json:"originalDescription,omitempty"`
		Ticker                       string  `json:"ticker,omitempty"`
		Quantity                     float64 `json:"quantity"`
		ManualClassification         string  `json:"manualClassification"`
		IsMarketMover                bool    `json:"isMarketMover"`
		OneDayPercentChangeSortIndex int     `json:"oneDayPercentChangeSortIndex"`
		OneDayValueChange            float64 `json:"oneDayValueChange"`
		Change                       float64 `json:"change"`
		ChangeSortIndex              int     `json:"changeSortIndex"`
		OneDayValueChangeSortIndex   int     `json:"oneDayValueChangeSortIndex"`
		SourceAssetID                string  `json:"sourceAssetId"`
		External                     string  `json:"external"`
		UserAccountID                int     `json:"userAccountId"`
		PriceSource                  string  `json:"priceSource"`
		ValueSortIndex               int     `json:"valueSortIndex"`
		Exchange                     string  `json:"exchange,omitempty"`
		OneDayPercentChange          float64 `json:"oneDayPercentChange"`
		TradingRatio                 float64 `json:"tradingRatio,omitempty"`
		Type                         string  `json:"type,omitempty"`
		TaxCost                      float64 `json:"taxCost,omitempty"`
		FundFees                     float64 `json:"fundFees,omitempty"`
		FeesPerYear                  float64 `json:"feesPerYear,omitempty"`
		CostBasis                    float64 `json:"costBasis,omitempty"`
		Currency                     string  `json:"currency,omitempty"`
	}
	SpData struct {
		Classifications    []interface{}  `json:"classifications"`
		Holdings           []HoldingsType `json:"holdings"`
		HoldingsTotalValue float64        `json:"holdingsTotalValue"`
	}
)

func (h HoldingsType) CurrencySymbolName() string {
	return h.Description
}

func (h HoldingsType) CurrencyName() string {
	return h.Ticker
}

func (h HoldingsType) TotalSharesString() string {
	s := strconv.FormatFloat(h.Quantity, 'f', -1, 64)
	return s
}

func (h *Holdings) GetHoldings(ctx context.Context, params *GetHoldingsParams) (*GetHoldingsResponse, error) {
	baseURL := h.client.cfg.Host
	client := h.client.cfg.HTTPClient
	csrf := *h.client.CSRF

	var urlBuffer bytes.Buffer
	urlBuffer.WriteString(baseURL)
	urlBuffer.WriteString(GetHoldingsEndpoint)

	data := url.Values{
		"apiClient": {"WEB"},
		"csrf":      {csrf}}

	req, err := http.NewRequest("POST", urlBuffer.String(), strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req = req.WithContext(ctx)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("non 200 response")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	ghr := GetHoldingsResponseExpecting{}
	err = json.Unmarshal(body, &ghr)
	if err != nil {
		return nil, err
	}
	lHoldings := new(GetHoldingsResponse)
	lHoldings.SpData = ghr.SpData
	if params.FilterUserCreated {
		lHoldings.Holdings = FilterHoldsFromUser(lHoldings.Holdings)
	}
	if params.FilterAccountName != "" {
		lHoldings.Holdings = FilterHoldsByAccountName(lHoldings.Holdings, params.FilterAccountName)
	}
	return lHoldings, nil
}

func PCHoldingsToIHoldings(h []HoldingsType) []providers.IHolding {
	a := make([]providers.IHolding, len(h))
	for i := 0; i < len(h); i++ {
		a[i] = h[i]
	}
	return a
}

func DumpResponse(resp *http.Response) {
	b, _ := httputil.DumpResponse(resp, true)
	fmt.Println(string(b))
}

func FilterHoldsFromUser(holdings []HoldingsType) []HoldingsType {
	h := make([]HoldingsType, 0, len(holdings))
	for i, v := range holdings {
		if v.Source == "USER" {
			h = append(h, holdings[i])
		}
	}
	return h
}

func FilterHoldsByAccountName(holdings []HoldingsType, accountName string) []HoldingsType {
	h := make([]HoldingsType, 0, len(holdings))
	for i, v := range holdings {
		if v.AccountName == accountName {
			h = append(h, holdings[i])
		}
	}
	return h
}

const (
	UpdateHoldingsEndpoint = `/api/account/updateHolding`
)

type (
	// userAccountId=80828031&priceSource=USER&valueSortIndex=1&costBasis=0&value=44.35&oneDayPercentChange=0&lastServerChangeId=21&csrf=7e23adfc-8058-408a-bb39-d7ee428a0dbb&apiClient=WEB
	HoldingsUpdateRequest struct {
		Ticker                       string  `url:"ticker"`
		Quantity                     string  `url:"quantity"`
		ManualClassification         string  `url:"manualClassification"`
		IsMarketMover                bool    `url:"isMarketMover"`
		AccountName                  string  `url:"accountName"`
		OneDayPercentChangeSortIndex int     `url:"oneDayPercentChangeSortIndex"`
		OneDayValueChange            float64 `url:"oneDayValueChange"`
		Change                       float64 `url:"change"`
		Description                  string  `url:"description"`
		Source                       string  `url:"source"`
		ChangeSortIndex              int     `url:"changeSortIndex"`
		OneDayValueChangeSortIndex   int     `url:"oneDayValueChangeSortIndex"`
		MarketType                   int     `url:"marketType"`
		SourceAssetId                string  `url:"sourceAssetId"`
		External                     string  `url:"external"`
		HoldingType                  string  `url:"holdingType"`
		Price                        string  `url:"price"`
		HoldingPercentage            float64 `url:"holdingPercentage"`
		UserAccountId                int     `url:"userAccountId"`
		PriceSource                  string  `url:"priceSource"`
		ValueSortIndex               int     `url:"valueSortIndex"`
		CostBasis                    string  `url:"costBasis"`
		Value                        string  `url:"value"`
		OneDayPercentChange          float64 `url:"oneDayPercentChange"`
		LastServerChangeId           string  `url:"lastServerChangeId"`
		Csrf                         string  `url:"csrf"`
		ApiClient                    string  `url:"apiClient"`
	}
	UpdateHoldingsResponse struct {
		SpHeader SpHeader `json:"spHeader"`
		SpData   struct {
			Holding struct {
				Ticker        string  `json:"ticker"`
				Quantity      float64 `json:"quantity"`
				AccountName   string  `json:"accountName"`
				Description   string  `json:"description"`
				Source        string  `json:"source"`
				SourceAssetID string  `json:"sourceAssetId"`
				External      string  `json:"external"`
				HoldingType   string  `json:"holdingType"`
				Price         float64 `json:"price"`
				PriceSource   string  `json:"priceSource"`
				UserAccountID int     `json:"userAccountId"`
				CostBasis     float64 `json:"costBasis"`
				Value         float64 `json:"value"`
			} `json:"holding"`
		} `json:"spData"`
	}
)

func (h *Holdings) UpdateHoldings(ctx context.Context, holding HoldingsUpdateRequest) (*UpdateHoldingsResponse, error) {
	baseURL := h.client.cfg.Host
	client := h.client.cfg.HTTPClient
	csrf := *h.client.CSRF

	var urlBuffer bytes.Buffer
	urlBuffer.WriteString(baseURL)
	urlBuffer.WriteString(UpdateHoldingsEndpoint)
	if holding.Csrf == "" {
		holding.Csrf = csrf
	}

	if holding.ApiClient == "" {
		holding.ApiClient = "WEB"
	}

	data, _ := query.Values(holding)
	req, err := http.NewRequest("POST", urlBuffer.String(), strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req = req.WithContext(ctx)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("non 200 response")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	update := UpdateHoldingsResponse{}
	err = json.Unmarshal(body, &update)
	if len(update.SpHeader.Errors) > 0 {
		return nil, errors.New(update.SpHeader.Errors[0].Message)
	}
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return &update, err
}

// UpdateCashAmount this is a special Add/Update Holding because this one is built into personal capital for all manual investments accounts.
func (h *Holdings) UpdateCashAmount(ctx context.Context, cashAmount string, userAccountId int) (*UpdateHoldingsResponse, error) {
	return h.UpdateHoldings(ctx, HoldingsUpdateRequest{
		Quantity:      cashAmount,
		Description:   "Cash",
		SourceAssetId: "USER_DESCR_Cash",
		Price:         "1.00",
		UserAccountId: userAccountId,
		PriceSource:   "USER",
	})
}

const (
	AddHoldingEndpoint = `/api/account/addHolding`
)

type (
	HoldingAddRequest struct {
		// userAccountId=80828031&ticker=Stellar+Lumens&cusip=&description=XML&quantity=965.08&price=0.06&value=57.9048&source=USER&costBasis=&lastServerChangeId=152&csrf=3c19636c-175f-447c-9cc7-f56a5be07f0d&apiClient=WEB
		Ticker             string `url:"ticker"`
		Quantity           string `url:"quantity"`
		Description        string `url:"description"`
		Source             string `url:"source"`
		Price              string `url:"price"`
		UserAccountId      int    `url:"userAccountId"`
		CostBasis          string `url:"costBasis"`
		Value              string `url:"value"`
		LastServerChangeId string `url:"lastServerChangeId"`
		Csrf               string `url:"csrf"`
		ApiClient          string `url:"apiClient"`
	}

	AddHoldingResponse struct {
		SpHeader SpHeader `json:"spHeader"`
		SpData   struct {
			Holding struct {
				SourceAssetID string  `json:"sourceAssetId"`
				Ticker        string  `json:"ticker"`
				External      string  `json:"external"`
				Quantity      float64 `json:"quantity"`
				AccountName   string  `json:"accountName"`
				HoldingType   string  `json:"holdingType"`
				Price         float64 `json:"price"`
				PriceSource   string  `json:"priceSource"`
				UserAccountID int     `json:"userAccountId"`
				Description   string  `json:"description"`
				Source        string  `json:"source"`
				Value         float64 `json:"value"`
			} `json:"holding"`
		} `json:"spData"`
	}
)

func (h *Holdings) AddHolding(ctx context.Context, holding *HoldingAddRequest) (*AddHoldingResponse, error) {
	baseURL := h.client.cfg.Host
	client := h.client.cfg.HTTPClient
	csrf := *h.client.CSRF

	var urlBuffer bytes.Buffer
	urlBuffer.WriteString(baseURL)
	urlBuffer.WriteString(AddHoldingEndpoint)
	if holding.Csrf == "" {
		holding.Csrf = csrf
	}

	if holding.ApiClient == "" {
		holding.ApiClient = "WEB"
	}

	data, _ := query.Values(holding)
	req, err := http.NewRequest("POST", urlBuffer.String(), strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req = req.WithContext(ctx)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("non 200 response")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	update := AddHoldingResponse{}
	err = json.Unmarshal(body, &update)
	if len(update.SpHeader.Errors) > 0 {
		return nil, errors.New(update.SpHeader.Errors[0].Message)
	}
	if err != nil {
		return nil, err
	}
	return &update, err
}
