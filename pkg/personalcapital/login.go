package personalcapital

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
	"strings"
)

var (
	TwoFactorAuthenticationRequired = errors.New("Two Factor Authentication is Required")
)

type Authentication service

type LoginParams struct {
	Username, Password string
}

type LoginResponse struct {
	CRSF string
}

func (a *Authentication) Login(ctx context.Context, params LoginParams) (*LoginResponse, error) {
	initialCsrf, err := a.getCsrfFromHomePage(ctx)

	csrf, authLevel, err := a.identifyUser(ctx, params.Username, initialCsrf)
	if err != nil {
		return nil, err
	}
	if authLevel == UserRemembered {
		err := a.AuthenticateWithPassword(ctx, AuthenticateWithPasswordParams{
			Username: params.Username,
			Password: params.Password,
			CSRF:     csrf,
		})
		if err != nil {
			return nil, err
		}
	} else {
		return &LoginResponse{CRSF: csrf}, TwoFactorAuthenticationRequired
	}
	return &LoginResponse{CRSF: csrf}, err
}

type AuthenticateWithPasswordParams struct {
	Username, Password string
	CSRF               string
}

const (
	AuthenticateWithPasswordEndpoint = "/api/credential/authenticatePassword"
	SessionAuthenticated             = "SESSION_AUTHENTICATED"
	UserRemembered                   = "USER_REMEMBERED"
)

type AuthenticateWithPasswordResponse struct {
	SpHeader SpHeader `json:"spHeader"`
	SpData   struct {
		Credentials    []string `json:"credentials"`
		AllCredentials []struct {
			Name   string `json:"name"`
			Status string `json:"status"`
		} `json:"allCredentials"`
	} `json:"spData"`
}

func (a *Authentication) AuthenticateWithPassword(ctx context.Context, params AuthenticateWithPasswordParams) error {
	baseURL := a.client.cfg.Host
	client := a.client.cfg.HTTPClient

	var urlBuffer bytes.Buffer
	urlBuffer.WriteString(baseURL)
	urlBuffer.WriteString(AuthenticateWithPasswordEndpoint)

	data := url.Values{
		"bindDevice":      {"true"},
		"deviceName":      {""},
		"redirectTo":      {""},
		"skipFirstUse":    {""},
		"skipLinkAccount": {"false"},
		"referrerId":      {""},
		"passwd":          {params.Password},
		"apiClient":       {"WEB"},
		"csrf":            {params.CSRF}}

	req, err := http.NewRequest("POST", urlBuffer.String(), strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req = req.WithContext(ctx)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New("non 200 response")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyJSON := AuthenticateWithPasswordResponse{}
	err = json.Unmarshal(body, &bodyJSON)
	if err != nil {
		return err
	}

	if bodyJSON.SpHeader.AuthLevel != SessionAuthenticated || len(bodyJSON.SpHeader.Errors) > 0 {
		return errors.New("Invalid Authentication level")
	}

	return err
}

const (
	IdentifyUserEndpoint = "/api/login/identifyUser"
)

func (a *Authentication) identifyUser(ctx context.Context, username, inCsrf string) (csrf string, authlevel string, err error) {
	baseURL := a.client.cfg.Host
	client := a.client.cfg.HTTPClient

	var urlBuffer bytes.Buffer
	urlBuffer.WriteString(baseURL)
	urlBuffer.WriteString(IdentifyUserEndpoint)

	data := url.Values{
		"username":        {username},
		"csrf":            {inCsrf},
		"apiClient":       {"WEB"},
		"bindDevice":      {"false"},
		"skipLinkAccount": {"false"},
		"redirectTo":      {""},
		"skipFirstUse":    {""},
		"referrerId":      {""}}
	req, err := http.NewRequest("POST", urlBuffer.String(), strings.NewReader(data.Encode()))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req = req.WithContext(ctx)
	resp, err := client.Do(req)

	if err != nil {
		return
	}
	if resp.StatusCode == http.StatusOK {
		if a.client.cfg.Debug {
			requestDump, err := httputil.DumpResponse(resp, true)
			if err != nil {
				return "", "", err
			}
			log.Println(string(requestDump))
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", "", err
		}
		bodyJSON := loginJSON{}
		err = json.Unmarshal(body, &bodyJSON)
		if err != nil {
			return "", "", err
		}

		return bodyJSON.SpHeader.Csrf, bodyJSON.SpHeader.AuthLevel, nil
	}
	return "", "", errors.New("unable to get csrf and authlevel")
}

func (a *Authentication) getCsrfFromHomePage(ctx context.Context) (csrf string, err error) {
	baseURL := a.client.cfg.Host
	client := a.client.cfg.HTTPClient
	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		return
	}
	req = req.WithContext(ctx)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Save a copy of this request for debugging.
	dumpResponse, err := httputil.DumpResponse(resp, true)
	if err != nil {
		return "", err
	}
	if a.client.cfg.Debug {
		log.Println(string(dumpResponse))
	}

	re := regexp.MustCompile("globals.csrf='([a-f0-9-]+)'")

	csrf = re.FindString(string(dumpResponse))
	if len(strings.Split(csrf, "'")) > 1 {
		csrf = strings.Split(csrf, "'")[1]
		return csrf, nil
	}

	return "", errors.New("Unable to get csrf token")
}

func (a *Authentication) TwoFactorAuthentication() {

}

type ChallengeType string

const (
	SMS                    ChallengeType = "sms"
	Email                  ChallengeType = "email"
	EmailChallengeEndpoint               = "/api/credential/challengeEmail"
	SmsChallengeEndpoint                 = "/api/credential/challengeSms"
)

type CredentialChallengeParams struct {
	Username, Password string
	CSRF               string
	ChallengeType      ChallengeType
}

func (a *Authentication) CredentialChallenge(ctx context.Context, params CredentialChallengeParams) (*http.Response, error) {
	client := a.client.cfg.HTTPClient
	var (
		cType    string
		endpoint string
	)

	switch params.ChallengeType {
	case SMS:
		cType = string(params.ChallengeType)
		endpoint = SmsChallengeEndpoint
	case Email:
		cType = string(params.ChallengeType)
		endpoint = EmailChallengeEndpoint
	default:
		return nil, errors.New("Unknown challenge type")
	}

	var urlBuffer bytes.Buffer
	urlBuffer.WriteString(a.client.cfg.Host)
	urlBuffer.WriteString(endpoint)

	data := url.Values{
		"challengeReason": {"DEVICE_AUTH"},
		"challengeMethod": {"OP"},
		"challengeType":   {cType},
		"apiClient":       {"WEB"},
		"bindDevice":      {"false"},
		"csrf":            {params.CSRF}}

	req, err := http.NewRequest("POST", urlBuffer.String(), strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req = req.WithContext(ctx)
	resp, err := client.Do(req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
