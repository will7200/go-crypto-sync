package personalcapital

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"go.uber.org/zap"
)

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
	Logger        *zap.Logger
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

	if cfg.Logger == nil {
		cfg.Logger = zap.NewNop()
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

// SaveSession saves the cookies to a file.
func SaveSession(inClient *http.Client, baseURL string, fileName string) (err error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return
	}

	cookies := inClient.Jar.Cookies(u)
	jc, err := json.Marshal(cookies)
	if err != nil {
		return
	}
	err = ioutil.WriteFile(fileName, jc, 0644)
	return
}

// LoadSession loads the cookies from a previously saved session file.
func LoadSession(cookies *[]*http.Cookie, fileName string) (err error) {
	cookiesFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	err = json.Unmarshal(cookiesFile, &cookies)
	return
}
