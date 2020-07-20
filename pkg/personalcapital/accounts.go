package personalcapital

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/google/go-querystring/query"
)

type Accounts service

const (
	GetAccountsEndpoint = `/api/newaccount/getAccounts2`
)

type (
	AccountType struct {
		AvailableCash string `json:"availableCash"`
		IsOnUs        bool   `json:"isOnUs"`
		UserProductID int    `json:"userProductId"`
		IsHome        bool   `json:"isHome"`
		NextAction    struct {
			IconType             string        `json:"iconType"`
			Action               string        `json:"action"`
			Prompts              []interface{} `json:"prompts"`
			AggregationErrorType string        `json:"aggregationErrorType"`
		} `json:"nextAction"`
		IsIAVEligible bool `json:"isIAVEligible"`
		LoginFields   []struct {
			IsUsername bool `json:"isUsername,omitempty"`
			IsRequired bool `json:"isRequired"`
			Parts      []struct {
				Size      int    `json:"size"`
				Name      string `json:"name"`
				ID        string `json:"id"`
				Type      string `json:"type"`
				MaxLength int    `json:"maxLength"`
			} `json:"parts"`
			Label      string `json:"label"`
			IsPassword bool   `json:"isPassword,omitempty"`
		} `json:"loginFields"`
		EnrollmentConciergeRequested int     `json:"enrollmentConciergeRequested"`
		IsCrypto                     bool    `json:"isCrypto"`
		IsPartner                    bool    `json:"isPartner"`
		PriorBalance                 float64 `json:"priorBalance"`
		IsCustomManual               bool    `json:"isCustomManual"`
		OriginalName                 string  `json:"originalName"`
		IsIAVAccountNumberValid      bool    `json:"isIAVAccountNumberValid"`
		IsExcludeFromHousehold       bool    `json:"isExcludeFromHousehold"`
		IsAsset                      bool    `json:"isAsset"`
		Aggregating                  bool    `json:"aggregating"`
		Balance                      float64 `json:"balance"`
		IsStatementDownloadEligible  bool    `json:"isStatementDownloadEligible"`
		Is401KEligible               bool    `json:"is401KEligible"`
		IsAccountUsedInFunding       bool    `json:"isAccountUsedInFunding"`
		IsTransferPending            bool    `json:"isTransferPending"`
		IsOnUs401K                   bool    `json:"isOnUs401K"`
		AdvisoryFeePercentage        float64 `json:"advisoryFeePercentage"`
		LastRefreshed                int64   `json:"lastRefreshed"`
		ProductID                    int     `json:"productId"`
		UserSiteID                   int     `json:"userSiteId"`
		Is365DayTransactionEligible  bool    `json:"is365DayTransactionEligible"`
		IsManual                     bool    `json:"isManual"`
		LogoPath                     string  `json:"logoPath"`
		CurrentBalance               float64 `json:"currentBalance"`
		AccountType                  string  `json:"accountType"`
		PaymentFromStatus            bool    `json:"paymentFromStatus"`
		IsRefetchTransactionEligible bool    `json:"isRefetchTransactionEligible"`
		AccountID                    string  `json:"accountId"`
		IsManualPortfolio            bool    `json:"isManualPortfolio"`
		ExcludeFromProposal          bool    `json:"excludeFromProposal"`
		UserAccountID                int     `json:"userAccountId"`
		Name                         string  `json:"name"`
		FirmName                     string  `json:"firmName"`
		AccountTypeGroup             string  `json:"accountTypeGroup"`
		PaymentToStatus              bool    `json:"paymentToStatus"`
		AdditionalAttributes         struct {
			Trust struct {
				KindOfTrust          string `json:"kindOfTrust"`
				KindOfTrustSecondary string `json:"kindOfTrustSecondary"`
			} `json:"trust"`
		} `json:"additionalAttributes"`
		IsOnUsBank                bool    `json:"isOnUsBank"`
		DisbursementType          string  `json:"disbursementType"`
		IsPaymentToCapable        bool    `json:"isPaymentToCapable"`
		ClosedComment             string  `json:"closedComment"`
		IsTaxDeferredOrNonTaxable bool    `json:"isTaxDeferredOrNonTaxable"`
		Currency                  string  `json:"currency"`
		FundFees                  float64 `json:"fundFees"`
		ProductType               string  `json:"productType"`
		IsAccountNumberValidated  bool    `json:"isAccountNumberValidated"`
		AccountTypeNew            string  `json:"accountTypeNew"`
		IsRoutingNumberValidated  bool    `json:"isRoutingNumberValidated"`
		IsLiability               bool    `json:"isLiability"`
		IsTransferEligible        bool    `json:"isTransferEligible"`
		InfoSource                string  `json:"infoSource"`
		DefaultAdvisoryFee        float64 `json:"defaultAdvisoryFee"`
		IsAdvised                 bool    `json:"isAdvised"`
		FeesPerYear               string  `json:"feesPerYear"`
		IsEsog                    bool    `json:"isEsog"`
		CreatedDate               int64   `json:"createdDate"`
		ClosedDate                string  `json:"closedDate"`
		TotalFee                  float64 `json:"totalFee"`
		IsPaymentFromCapable      bool    `json:"isPaymentFromCapable"`
		SiteID                    int     `json:"siteId"`
		OriginalFirmName          string  `json:"originalFirmName"`
	}
	GetAccountsRequest struct {
		LastServerChangeId string `url:"lastServerChangeId"`
		Csrf               string `url:"csrf"`
		ApiClient          string `url:"apiClient"`
	}
	GetAccountsResponse struct {
		SpHeader SpHeader `json:"spHeader"`
		SpData   struct {
			CreditCardAccountsTotal       float64       `json:"creditCardAccountsTotal"`
			Assets                        float64       `json:"assets"`
			OtherLiabilitiesAccountsTotal float64       `json:"otherLiabilitiesAccountsTotal"`
			CashAccountsTotal             float64       `json:"cashAccountsTotal"`
			Liabilities                   float64       `json:"liabilities"`
			Networth                      float64       `json:"networth"`
			InvestmentAccountsTotal       float64       `json:"investmentAccountsTotal"`
			MortgageAccountsTotal         float64       `json:"mortgageAccountsTotal"`
			LoanAccountsTotal             float64       `json:"loanAccountsTotal"`
			Accounts                      []AccountType `json:"accounts"`
			OtherAssetAccountsTotal       float64       `json:"otherAssetAccountsTotal"`
		} `json:"spData"`
	}
)

// GetAccounts will obtains all accounts for a user's account.
// request parameter may be nil
func (a *Accounts) GetAccounts(ctx context.Context, request *GetAccountsRequest) (*GetAccountsResponse, error) {
	baseURL := a.client.cfg.Host
	client := a.client.cfg.HTTPClient
	csrf := *a.client.CSRF

	var urlBuffer bytes.Buffer
	urlBuffer.WriteString(baseURL)
	urlBuffer.WriteString(GetAccountsEndpoint)
	if request == nil {
		request = &GetAccountsRequest{LastServerChangeId: "-1"}
	}
	if request.Csrf == "" {
		request.Csrf = csrf
	}

	if request.ApiClient == "" {
		request.ApiClient = "WEB"
	}

	data, _ := query.Values(request)
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
	update := GetAccountsResponse{}
	err = json.Unmarshal(body, &update)
	if err != nil {
		return nil, err
	}
	return &update, err
}
