/*
 * Nomics Cryptocurrency & Bitcoin API
 *
 * # Introduction  Welcome to the Nomics Cryptocurrency & Bitcoin API. To sign up for an API key please [go here](https://p.nomics.com/cryptocurrency-bitcoin-api/).  [nomics.com](https://nomics.com) is built entirely with the Nomics API. Everything we've done on [nomics.com](https://nomics.com) you can do with our API. There are no internal API endpoints.  If you need support, reach out to use at our [forums](https://forums.nomics.com/).  # General  ## API Server URL  The Nomics API runs at `https://api.nomics.com/v1`. All requests should be prefixed by the server URL.  ## JSON and CSV Support  By default, all endpoints serve data as JSON. However, by passing `format=csv` in the URL, some endpoints will return CSV data. This can be used in Google Sheets via the `IMPORTDATA` function.  CSV responses will not contain a header row, this is so that data can be easily concatenated from multiple requests. The fields will be rendered in the same order as the JSON fields. See the endpoint's documentation for an example.  Not all endpoints support CSV. Endpoints that support CSV will have the `format` parameter in the parameters section.  ## Errors  The Nomics API uses standard HTTP status codes to indicate success or failure. 200 represents success, 4xx represents a user error (such as a problem with your key), and 5xx represents a problem with our API.  ## Versioning  We follow Semantic Versioning. That means our API is versioned as Major.Minor.Patch. For example, Version 1.2.3 has major version 1, minor version 2, and patch version 3.  Major version changes indicate that we have altered the API significantly and it is no longer compatible with a previous version. Major versions are also used as the API URL prefix.  When we update the major version, we will not remove the previous version without notice to API customers and a deprecation period to allow everyone to smoothly update to the new version.  Minor version changes indicate that we have added new functionality without breaking any existing functionality. An API client is compatible with future minor versions. Note that a minor version update may add a new field to an existing API endpoint's response. Your API client must ignore fields it does not understand in order to be compatible with future minor versions.  Patch version changes indicate we fixed a bug or security vulnerability. Patch versions don't add new functionality.  ## Cross Origin Resource Sharing (CORS)  This API supports Cross Origin Resource Sharing, which allows you to make API requests directly from your user's browser.  To use CORS, you must provide Nomics with the domains on which your application will run so that we can whitelist them for CORS access.  Requests from `localhost`, `127.0.0.1`, and `0.0.0.0` will always succeed to aid in development.  ## Demo Application  A demo application using the Nomics API, CORS, and React is available on Glitch.com. This can help you get started using the Nomics API. Keep in mind it uses the demo key, which is rotated frequently. You should get your own API key before deploying an app to production. Check it out:  <div class=\"glitch-embed-wrap\" style=\"height: 420px; width: 100%;\">   <iframe src=\"https://glitch.com/embed/#!/embed/nomics-api-demo?path=README.md\" alt=\"nomics-api-demo on glitch\" style=\"height: 100%; width: 100%; border: 0;\"></iframe> </div>  ## Demo Spreadsheet  Here is a demo of using the Nomics API with Google Sheets.  <iframe width=\"100%\" height=\"400px\" src=\"https://docs.google.com/spreadsheets/d/e/2PACX-1vShn2iWjvqQ0ueBa9l9g1UBYVM92OZSgZ4nmp0rWuykvHPrvyMyMeSN4r0Orj0ACEIIKdCz6cc5abCw/pubhtml?widget=true&amp;headers=false\"></iframe>  ### Formulas  * A2: `=IMPORTDATA(\"https://api.nomics.com/v1/prices?key=your-key-here&format=csv\")` * Column F: `=LOOKUP(D2,A:A,B:B)` finds D2 (BTC) in column A and pulls the price from column B * Column G: `=E2*F2` * Column H: `=G2/I$2` * Column I: `=SUM(G:G)`  # SDKs and Libraries  ## By Nomics - [Nomics JavaScript Client](https://github.com/nomics-crypto/nomics-javascript)  ## Community Submissions - [Nomics.com Swift SDK](https://forums.nomics.com/t/swift-sdk-supporting-ios-macos-tvos-and-watchos/) by Nick DiZazzo - [Nomics Node.js Library](https://forums.nomics.com/t/i-made-a-library-for-node-js/) by mikunimaru - [Nomics Python Wrapper](https://forums.nomics.com/t/python-package-for-nomics-api/119) by Taylor Facen - [Python Wrapper for Nomics](https://github.com/AviFelman/py-nomics) by Avi Felman  We love watching developers explore new use-cases with our API. Whether you're tinkering on a small side project or building an open-source resource, please share what you're up to in our [forums](https://forums.nomics.com/).  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package nomics

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

// ExchangesApiService ExchangesApi service
type ExchangesApiService service

type ApiGetExchangeVolumeHistoryRequest struct {
	ctx                 _context.Context
	ApiService          *ExchangesApiService
	exchange            *string
	start               *string
	end                 *string
	convert             *string
	format              *string
	includeTransparency *bool
}

// exchange  Exchange ID
func (r ApiGetExchangeVolumeHistoryRequest) Exchange(exchange string) ApiGetExchangeVolumeHistoryRequest {
	r.exchange = &exchange
	return r
}

// start  Start time of the interval in RFC3339 (URI escaped)
func (r ApiGetExchangeVolumeHistoryRequest) Start(start string) ApiGetExchangeVolumeHistoryRequest {
	r.start = &start
	return r
}

// end  End time of the interval in RFC3339 (URI escaped). If not provided, the current time is used.
func (r ApiGetExchangeVolumeHistoryRequest) End(end string) ApiGetExchangeVolumeHistoryRequest {
	r.end = &end
	return r
}

// convert  Currency to quote ticker price, market cap, and volume values. May be a Fiat Currency or Cryptocurrency. Default is `USD`.
func (r ApiGetExchangeVolumeHistoryRequest) Convert(convert string) ApiGetExchangeVolumeHistoryRequest {
	r.convert = &convert
	return r
}

// format  Format of the response. Defaults to JSON when blank.
func (r ApiGetExchangeVolumeHistoryRequest) Format(format string) ApiGetExchangeVolumeHistoryRequest {
	r.format = &format
	return r
}

// includeTransparency  Whether to include [Transparent Volume](https://blog.nomics.com/essays/transparent-volume/) and transparent market cap information in the response. Default is `false`. This option is only available to customers of our paid API plans.
func (r ApiGetExchangeVolumeHistoryRequest) IncludeTransparency(includeTransparency bool) ApiGetExchangeVolumeHistoryRequest {
	r.includeTransparency = &includeTransparency
	return r
}

func (r ApiGetExchangeVolumeHistoryRequest) Execute() ([]InlineResponse2007, *_nethttp.Response, error) {
	return r.ApiService.GetExchangeVolumeHistoryExecute(r)
}

/*
 * GetExchangeVolumeHistory Exchanges Volume History
 * Exchange Volume History is the total volume for all
cryptoassets of an exchange in USD at intervals between the
requested time period.  For each entry, the `volume` field
represents the sum of the `spot_volume` and
`derivative_volume` fields.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetExchangeVolumeHistoryRequest
*/
func (a *ExchangesApiService) GetExchangeVolumeHistory(ctx _context.Context) ApiGetExchangeVolumeHistoryRequest {
	return ApiGetExchangeVolumeHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return []InlineResponse2007
 */
func (a *ExchangesApiService) GetExchangeVolumeHistoryExecute(r ApiGetExchangeVolumeHistoryRequest) ([]InlineResponse2007, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []InlineResponse2007
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExchangesApiService.GetExchangeVolumeHistory")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/exchanges/volume/history"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.exchange == nil {
		return localVarReturnValue, nil, reportError("exchange is required and must be specified")
	}
	if r.start == nil {
		return localVarReturnValue, nil, reportError("start is required and must be specified")
	}

	localVarQueryParams.Add("exchange", parameterToString(*r.exchange, ""))
	localVarQueryParams.Add("start", parameterToString(*r.start, ""))
	if r.end != nil {
		localVarQueryParams.Add("end", parameterToString(*r.end, ""))
	}
	if r.convert != nil {
		localVarQueryParams.Add("convert", parameterToString(*r.convert, ""))
	}
	if r.format != nil {
		localVarQueryParams.Add("format", parameterToString(*r.format, ""))
	}
	if r.includeTransparency != nil {
		localVarQueryParams.Add("include-transparency", parameterToString(*r.includeTransparency, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "test/csv"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Key"]; ok {
				var key string
				key = apiKey.Key
				localVarQueryParams.Add("key", key)
			}
		}
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

type ApiGetExchangesRequest struct {
	ctx           _context.Context
	ApiService    *ExchangesApiService
	ids           *string
	attributes    *string
	centralized   *string
	decentralized *string
	format        *string
}

// ids  Comma separated list of Nomics Exchange IDs to filter result rows
func (r ApiGetExchangesRequest) Ids(ids string) ApiGetExchangesRequest {
	r.ids = &ids
	return r
}

// attributes  Comma separated list of exchange attributes to filter result columns
func (r ApiGetExchangesRequest) Attributes(attributes string) ApiGetExchangesRequest {
	r.attributes = &attributes
	return r
}

// centralized  When `true`, selects centralized exchanges.  When `false`, removes centralized exchanges.  If omitted, all exchange types are shown.  May be used in conjunction with `decentralized` to form an \"and\" filter.
func (r ApiGetExchangesRequest) Centralized(centralized string) ApiGetExchangesRequest {
	r.centralized = &centralized
	return r
}

// decentralized  When `true`, selects decentralized exchanges.  When `false`, removes decentralized exchanges.  If omitted, all exchange types are shown.  May be used in conjunction with `centralized` to form an \"and\" filter.
func (r ApiGetExchangesRequest) Decentralized(decentralized string) ApiGetExchangesRequest {
	r.decentralized = &decentralized
	return r
}

// format  Format of the response. Defaults to JSON when blank.
func (r ApiGetExchangesRequest) Format(format string) ApiGetExchangesRequest {
	r.format = &format
	return r
}

func (r ApiGetExchangesRequest) Execute() ([]InlineResponse2008, *_nethttp.Response, error) {
	return r.ApiService.GetExchangesExecute(r)
}

/*
 * GetExchanges Exchanges Metadata
 * The exchanges endpoint returns all the exchanges and their metadata that Nomics supports.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetExchangesRequest
 */
func (a *ExchangesApiService) GetExchanges(ctx _context.Context) ApiGetExchangesRequest {
	return ApiGetExchangesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return []InlineResponse2008
 */
func (a *ExchangesApiService) GetExchangesExecute(r ApiGetExchangesRequest) ([]InlineResponse2008, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []InlineResponse2008
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExchangesApiService.GetExchanges")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/exchanges"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.ids != nil {
		localVarQueryParams.Add("ids", parameterToString(*r.ids, ""))
	}
	if r.attributes != nil {
		localVarQueryParams.Add("attributes", parameterToString(*r.attributes, ""))
	}
	if r.centralized != nil {
		localVarQueryParams.Add("centralized", parameterToString(*r.centralized, ""))
	}
	if r.decentralized != nil {
		localVarQueryParams.Add("decentralized", parameterToString(*r.decentralized, ""))
	}
	if r.format != nil {
		localVarQueryParams.Add("format", parameterToString(*r.format, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/csv"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Key"]; ok {
				var key string
				key = apiKey.Key
				localVarQueryParams.Add("key", key)
			}
		}
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

type ApiGetExchangesTickerRequest struct {
	ctx           _context.Context
	ApiService    *ExchangesApiService
	ids           *string
	interval      *string
	convert       *string
	status        *string
	type_         *string
	centralized   *string
	decentralized *string
	perPage       *int32
	page          *int32
}

// ids  Comma separated list of Nomics Exchange IDs to filter result rows
func (r ApiGetExchangesTickerRequest) Ids(ids string) ApiGetExchangesTickerRequest {
	r.ids = &ids
	return r
}

// interval  Comma separated time interval of the ticker(s). Default is `1d`.
func (r ApiGetExchangesTickerRequest) Interval(interval string) ApiGetExchangesTickerRequest {
	r.interval = &interval
	return r
}

// convert  Currency to quote ticker price, market cap, and volume values. May be a Fiat Currency or Cryptocurrency. Default is `USD`.
func (r ApiGetExchangesTickerRequest) Convert(convert string) ApiGetExchangesTickerRequest {
	r.convert = &convert
	return r
}

// status  Status by which to filter exchanges.  If not provided, all exchanges are shown.
func (r ApiGetExchangesTickerRequest) Status(status string) ApiGetExchangesTickerRequest {
	r.status = &status
	return r
}

// type_  Type by which to filter exchanges.  If not provided, all exchanges are shown.
func (r ApiGetExchangesTickerRequest) Type_(type_ string) ApiGetExchangesTickerRequest {
	r.type_ = &type_
	return r
}

// centralized  When `true`, selects centralized exchanges.  When `false`, removes centralized exchanges.  If omitted, all exchange types are shown.  May be used in conjunction with `decentralized` to form an \"and\" filter.
func (r ApiGetExchangesTickerRequest) Centralized(centralized string) ApiGetExchangesTickerRequest {
	r.centralized = &centralized
	return r
}

// decentralized  When `true`, selects decentralized exchanges.  When `false`, removes decentralized exchanges.  If omitted, all exchange types are shown.  May be used in conjunction with `centralized` to form an \"and\" filter.
func (r ApiGetExchangesTickerRequest) Decentralized(decentralized string) ApiGetExchangesTickerRequest {
	r.decentralized = &decentralized
	return r
}

// perPage  The maximum number of items to return per paginated response. Paginated responses include an additional response header, `X-Pagination-Total-Items`, which represents the total number of items available after all the request filters have been applied.  Must be between `1` and `100` (inclusive).
func (r ApiGetExchangesTickerRequest) PerPage(perPage int32) ApiGetExchangesTickerRequest {
	r.perPage = &perPage
	return r
}

// page  Which page of items to get.  Only applicable when `per-page` is also supplied.
func (r ApiGetExchangesTickerRequest) Page(page int32) ApiGetExchangesTickerRequest {
	r.page = &page
	return r
}

func (r ApiGetExchangesTickerRequest) Execute() ([]InlineResponse2006, *_nethttp.Response, error) {
	return r.ApiService.GetExchangesTickerExecute(r)
}

/*
 * GetExchangesTicker Exchanges Ticker
 * The Exchanges Ticker provides high level information about the exchanges integrated with Nomics.
It provides a limited amount of metadata, the type of integration, the time range of available data, pairs,
and interval information about the volume and, where applicable, the number of trades.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetExchangesTickerRequest
*/
func (a *ExchangesApiService) GetExchangesTicker(ctx _context.Context) ApiGetExchangesTickerRequest {
	return ApiGetExchangesTickerRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return []InlineResponse2006
 */
func (a *ExchangesApiService) GetExchangesTickerExecute(r ApiGetExchangesTickerRequest) ([]InlineResponse2006, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []InlineResponse2006
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExchangesApiService.GetExchangesTicker")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/exchanges/ticker"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.ids != nil {
		localVarQueryParams.Add("ids", parameterToString(*r.ids, ""))
	}
	if r.interval != nil {
		localVarQueryParams.Add("interval", parameterToString(*r.interval, ""))
	}
	if r.convert != nil {
		localVarQueryParams.Add("convert", parameterToString(*r.convert, ""))
	}
	if r.status != nil {
		localVarQueryParams.Add("status", parameterToString(*r.status, ""))
	}
	if r.type_ != nil {
		localVarQueryParams.Add("type", parameterToString(*r.type_, ""))
	}
	if r.centralized != nil {
		localVarQueryParams.Add("centralized", parameterToString(*r.centralized, ""))
	}
	if r.decentralized != nil {
		localVarQueryParams.Add("decentralized", parameterToString(*r.decentralized, ""))
	}
	if r.perPage != nil {
		localVarQueryParams.Add("per-page", parameterToString(*r.perPage, ""))
	}
	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Key"]; ok {
				var key string
				key = apiKey.Key
				localVarQueryParams.Add("key", key)
			}
		}
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
