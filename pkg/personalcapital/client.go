package personalcapital

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

type session struct {
	AWSELB           string `json:"AWSELB"`
	JSESSIONID       string `json:"JSESSIONID"`
	PMData           string `json:"PMData"`
	REMEMBERMECOOKIE string `json:"REMEMBER_ME_COOKIE"`
	Cfduid           string `json:"__cfduid"`
}

type loginJSON struct {
	SpHeader SpHeader `json:"spHeader"`
	SpData   struct {
		UserStatus     string   `json:"userStatus"`
		Credentials    []string `json:"credentials"`
		AllCredentials []struct {
			Name   string `json:"name"`
			Status string `json:"status"`
		} `json:"allCredentials"`
	} `json:"spData"`
}

type Configuration struct {
	Host          string            `json:"host,omitempty"`
	DefaultHeader map[string]string `json:"defaultHeader,omitempty"`
	UserAgent     string            `json:"userAgent,omitempty"`
	Debug         bool              `json:"debug,omitempty"`
	HTTPClient    *http.Client
}

// NewConfiguration returns a new Configuration object
func NewConfiguration() *Configuration {
	cfg := &Configuration{
		Host:          "https://home.personalcapital.com",
		DefaultHeader: make(map[string]string),
		UserAgent:     "Crypto-Sync/0.1.0/go",
		Debug:         false,
	}
	return cfg
}

type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// CSRF use this once authenticated
	CSRF *string
	// APIs
	Authentication *Authentication
	Holdings       *Holdings
	Accounts       *Accounts
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = http.DefaultClient
	}

	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// API Services
	c.Authentication = (*Authentication)(&c.common)
	c.Holdings = (*Holdings)(&c.common)
	c.Accounts = (*Accounts)(&c.common)

	return c
}

// ApiRequest can be used to send API request after successfully
// logging in to Personal Capital
func ApiRequest(client *http.Client, csrf string, payload map[string][]string, endpoint string) []byte {
	//Example to getAccounts and print out networth

	payload = map[string][]string{
		"lastServerChangeId": {"-1"},
		"csrf":               {csrf},
		"apiClient":          {"WEB"}}

	baseURL := "https://home.personalcapital.com"
	apiURL := baseURL + "/api"
	endpoint = "/newaccount/getAccounts"

	var urlBuffer bytes.Buffer
	urlBuffer.WriteString(apiURL)
	urlBuffer.WriteString(endpoint)

	resp, err := client.PostForm(
		urlBuffer.String(),
		payload)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	//prettyPrintJSON(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// accountJSON := getAccountsJSON{}
	// json.Unmarshal(body, &accountJSON)

	// return accountJSON.SpData.Networth
	return body
}

// TwoFactorAuthentication is required if the user has setup 2FA and
// doesn't have a saved session. "sms" and "email" 2FA can be used.
func TwoFactorAuthentication(client *http.Client, csrf string, username string, password string, challengeType string) error {
	baseURL := "https://home.personalcapital.com"
	apiURL := baseURL + "/api"
	endpoint := "/credential/challengeSms"

	if challengeType == "sms" {
		challengeType = "challengeSMS"
		endpoint = "/credential/challengeSms"
	} else if challengeType == "email" {
		challengeType = "challengeEmail"
		endpoint = "/credential/challengeEmail"
	} else {
		// default to sms
		challengeType = "challengeSMS"
		endpoint = "/credential/challengeSms"
	}

	var urlBuffer bytes.Buffer
	urlBuffer.WriteString(apiURL)
	urlBuffer.WriteString(endpoint)

	payload := map[string][]string{
		"challengeReason": {"DEVICE_AUTH"},
		"challengeMethod": {"OP"},
		"challengeType":   {challengeType},
		"apiClient":       {"WEB"},
		"bindDevice":      {"false"},
		"csrf":            {csrf}}

	resp, err := client.PostForm(
		urlBuffer.String(),
		payload)
	if err != nil {
		log.Fatalln(err)
	}

	if resp.StatusCode == http.StatusOK {
		// Authenticate 2 factor code
		baseURL := "https://home.personalcapital.com"
		apiURL := baseURL + "/api"
		endpoint := "/credential/authenticateSms"

		if challengeType == "challengeSMS" {
			endpoint = "/credential/authenticateSms"
		} else if challengeType == "challengeEmail" {
			endpoint = "/credential/authenticateEmailByCode"
		} else {
			// default to sms
			endpoint = "/credential/authenticateSms"
		}

		var urlBuffer bytes.Buffer
		urlBuffer.WriteString(apiURL)
		urlBuffer.WriteString(endpoint)

		fmt.Print("Please enter 2FA code: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		code := scanner.Text()
		if scanner.Err() != nil {
			log.Fatal(scanner.Err())
		}
		fmt.Println(urlBuffer.String())
		resp, err := client.PostForm(
			urlBuffer.String(),
			url.Values{
				"challengeReason": {"DEVICE_AUTH"},
				"challengeMethod": {"OP"},
				"apiClient":       {"WEB"},
				"bindDevice":      {"false"},
				"code":            {code},
				"csrf":            {csrf}})
		if err != nil {
			log.Fatalln(err)
		}
		// Save a copy of this request for debugging.
		requestDump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			log.Println(err)
		}
		log.Println(string(requestDump))
		if resp.StatusCode == http.StatusOK {
			authenticatePassword(client, username, password, csrf)
		} else {
			err = errors.New("bad http status code")
		}
	}
	return err
}

// SaveSession saves the cookies to a file.
func SaveSession(inClient *http.Client, baseURL string, fileName string) {
	u, err := url.Parse(baseURL)
	if err != nil {
		log.Println(err)
	}

	cookies := inClient.Jar.Cookies(u)
	jc, err := json.Marshal(cookies)
	if err != nil {
		log.Println(err)
	}
	err = ioutil.WriteFile(fileName, jc, 0644)
	if err != nil {
		log.Println(err)
	}
}

// LoadSession loads the cookies from a previously saved session file.
func LoadSession(cookies *[]*(http.Cookie), fileName string) {
	cookiesFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(cookiesFile, &cookies)
	if err != nil {
		log.Println(err)
	}
}

func authenticatePassword(client *http.Client, username string, password string, csrf string) {
	baseURL := "https://home.personalcapital.com"
	apiURL := baseURL + "/api"
	endpoint := "/credential/authenticatePassword"

	var urlBuffer bytes.Buffer
	urlBuffer.WriteString(apiURL)
	urlBuffer.WriteString(endpoint)

	resp, err := client.PostForm(
		urlBuffer.String(),
		url.Values{
			"bindDevice":      {"true"},
			"deviceName":      {""},
			"redirectTo":      {""},
			"skipFirstUse":    {""},
			"skipLinkAccount": {"false"},
			"referrerId":      {""},
			"passwd":          {password},
			"apiClient":       {"WEB"},
			"csrf":            {csrf}})
	if err != nil {
		log.Fatalln(err)
	}

	// Save a copy of this request for debugging.
	requestDump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(requestDump))
}
