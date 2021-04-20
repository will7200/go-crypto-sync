package common

import (
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/HereMobilityDevelopers/mediary"
)

func DumpRequestResponse(req *http.Request, handler mediary.Handler) (*http.Response, error) {
	if bytes, err := httputil.DumpRequestOut(req, true); err == nil {
		fmt.Printf("%s\n", bytes)
	}
	r, err := handler(req)
	if err == nil {
		if bytes, err := httputil.DumpResponse(r, true); err == nil {
			fmt.Printf("%s\n", bytes)
		}
	}
	return r, err
}
