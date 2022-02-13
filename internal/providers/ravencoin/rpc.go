package ravencoin

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"unicode"
)

func isWindowsDriveURIPath(uri string) bool {
	if len(uri) < 4 {
		return false
	}
	return uri[0] == '/' && unicode.IsLetter(rune(uri[1])) && uri[2] == ':'
}

// split slices s into two substrings separated by the first occurrence of
// sep. If cutc is true then sep is excluded from the second substring.
// If sep does not occur in s then s and the empty string is returned.
func split(s string, sep byte, cutc bool) (string, string) {
	i := strings.IndexByte(s, sep)
	if i < 0 {
		return s, ""
	}
	if cutc {
		return s[:i], s[i+1:]
	}
	return s[:i], s[i:]
}

type rpcNodeArgs struct {
	url          string
	userPassword string
}

type GetWalletInfoResult struct {
	Result struct {
		Walletname            string  `json:"walletname"`
		Walletversion         int     `json:"walletversion"`
		Balance               float64 `json:"balance"`
		UnconfirmedBalance    float64 `json:"unconfirmed_balance"`
		ImmatureBalance       float64 `json:"immature_balance"`
		Txcount               int     `json:"txcount"`
		Keypoololdest         int     `json:"keypoololdest"`
		Keypoolsize           int     `json:"keypoolsize"`
		KeypoolsizeHdInternal int     `json:"keypoolsize_hd_internal"`
		Paytxfee              float64 `json:"paytxfee"`
		Hdseedid              string  `json:"hdseedid"`
		Hdmasterkeyid         string  `json:"hdmasterkeyid"`
	} `json:"result"`
	Error interface{} `json:"error"`
	ID    int         `json:"id"`
}

func readFromRPCNode(args rpcNodeArgs) (*GetWalletInfoResult, error) {
	data, err := json.Marshal(map[string]interface{}{
		"method": "getwalletinfo",
		"id":     1,
	})
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(args.url)
	if err != nil {
		return nil, err
	}
	userPassword, err := url.ParseRequestURI(args.userPassword)
	if err != nil {
		return nil, err
	}
	switch userPassword.Scheme {
	case "file":
		path := userPassword.Path
		if isWindowsDriveURIPath(path) {
			path = strings.ToUpper(string(path[1])) + path[2:]
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}
		userinfo := string(data)
		u, err = url.Parse(fmt.Sprintf("%s://%s@%s", u.Scheme, userinfo, u.Host))
	default:
		// assume its user password
		username, password := split(args.userPassword, ':', true)
		u.User = url.UserPassword(username, password)
	}
	//u.User = url.UserPassword()

	resp, err := http.Post(u.String(),
		"application/json", strings.NewReader(string(data)))
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 401:
		return nil, errors.New("Authentication is required")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result GetWalletInfoResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
