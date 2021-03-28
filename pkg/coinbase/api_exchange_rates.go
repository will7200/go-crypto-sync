/*
 * CoinBase API
 *
 * Coinbase provides a simple and powerful REST API to integrate bitcoin, bitcoin cash, litecoin and ethereum payments into your business or application.  This API reference provides information on available endpoints and how to interact with it. To read more about the API, visit our API documentation.
 *
 * API version: 2019-11-15
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package coinbase

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// ExchangeRatesApiService ExchangeRatesApi service
type ExchangeRatesApiService service

type ApiGetExchangeRateForRequest struct {
	ctx           _context.Context
	ApiService    *ExchangeRatesApiService
	currencyPair1 string
	currencyPair2 string
}

func (r ApiGetExchangeRateForRequest) Execute() (InlineResponse2001, *_nethttp.Response, error) {
	return r.ApiService.GetExchangeRateForExecute(r)
}

/*
 * GetExchangeRateFor Exchanges Rates for currency pair
 * Get the total price to buy one bitcoin or ether.

Note that exchange rates fluctuates so the price is only correct for seconds at the time. This buy price includes standard Coinbase fee (1%) but excludes any other fees including bank fees. If you need more accurate price estimate for a specific payment method or amount, see buy bitcoin endpoint and quote: true option.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param currencyPair1 currency 1
 * @param currencyPair2 currency 2
 * @return ApiGetExchangeRateForRequest
*/
func (a *ExchangeRatesApiService) GetExchangeRateFor(ctx _context.Context, currencyPair1 string, currencyPair2 string) ApiGetExchangeRateForRequest {
	return ApiGetExchangeRateForRequest{
		ApiService:    a,
		ctx:           ctx,
		currencyPair1: currencyPair1,
		currencyPair2: currencyPair2,
	}
}

/*
 * Execute executes the request
 * @return InlineResponse2001
 */
func (a *ExchangeRatesApiService) GetExchangeRateForExecute(r ApiGetExchangeRateForRequest) (InlineResponse2001, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse2001
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExchangeRatesApiService.GetExchangeRateFor")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/prices/{currency_pair_1}-{currency_pair_2}/sell"
	localVarPath = strings.Replace(localVarPath, "{"+"currency_pair_1"+"}", _neturl.PathEscape(parameterToString(r.currencyPair1, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"currency_pair_2"+"}", _neturl.PathEscape(parameterToString(r.currencyPair2, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
