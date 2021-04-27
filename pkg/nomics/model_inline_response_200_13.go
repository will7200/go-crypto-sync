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
	"encoding/json"
)

// InlineResponse20013 struct for InlineResponse20013
type InlineResponse20013 struct {
	// Nomics Exchange ID
	Exchange *string `json:"exchange,omitempty"`
	// Nomics Currency ID of the asset being traded
	Base *string `json:"base,omitempty"`
	// Nomics Currency ID of the asset used to quote a price
	Quote *string `json:"quote,omitempty"`
	// Total volume in the base currency
	VolumeBase *string `json:"volume_base,omitempty"`
	// Total volume in USD
	VolumeUsd *string `json:"volume_usd,omitempty"`
	// Open price in the quote currency
	OpenQuote *string `json:"open_quote,omitempty"`
	// Timestamp of the open price in RFC 3339
	OpenTimestamp *string `json:"open_timestamp,omitempty"`
	// Close price in the quote currency
	CloseQuote *string `json:"close_quote,omitempty"`
	// Timestamp of the close price in RFC 3339
	CloseTimestamp *string `json:"close_timestamp,omitempty"`
	// Total number of trades
	NumTrades *string `json:"num_trades,omitempty"`
}

// NewInlineResponse20013 instantiates a new InlineResponse20013 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20013() *InlineResponse20013 {
	this := InlineResponse20013{}
	return &this
}

// NewInlineResponse20013WithDefaults instantiates a new InlineResponse20013 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20013WithDefaults() *InlineResponse20013 {
	this := InlineResponse20013{}
	return &this
}

// GetExchange returns the Exchange field value if set, zero value otherwise.
func (o *InlineResponse20013) GetExchange() string {
	if o == nil || o.Exchange == nil {
		var ret string
		return ret
	}
	return *o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20013) GetExchangeOk() (*string, bool) {
	if o == nil || o.Exchange == nil {
		return nil, false
	}
	return o.Exchange, true
}

// HasExchange returns a boolean if a field has been set.
func (o *InlineResponse20013) HasExchange() bool {
	if o != nil && o.Exchange != nil {
		return true
	}

	return false
}

// SetExchange gets a reference to the given string and assigns it to the Exchange field.
func (o *InlineResponse20013) SetExchange(v string) {
	o.Exchange = &v
}

// GetBase returns the Base field value if set, zero value otherwise.
func (o *InlineResponse20013) GetBase() string {
	if o == nil || o.Base == nil {
		var ret string
		return ret
	}
	return *o.Base
}

// GetBaseOk returns a tuple with the Base field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20013) GetBaseOk() (*string, bool) {
	if o == nil || o.Base == nil {
		return nil, false
	}
	return o.Base, true
}

// HasBase returns a boolean if a field has been set.
func (o *InlineResponse20013) HasBase() bool {
	if o != nil && o.Base != nil {
		return true
	}

	return false
}

// SetBase gets a reference to the given string and assigns it to the Base field.
func (o *InlineResponse20013) SetBase(v string) {
	o.Base = &v
}

// GetQuote returns the Quote field value if set, zero value otherwise.
func (o *InlineResponse20013) GetQuote() string {
	if o == nil || o.Quote == nil {
		var ret string
		return ret
	}
	return *o.Quote
}

// GetQuoteOk returns a tuple with the Quote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20013) GetQuoteOk() (*string, bool) {
	if o == nil || o.Quote == nil {
		return nil, false
	}
	return o.Quote, true
}

// HasQuote returns a boolean if a field has been set.
func (o *InlineResponse20013) HasQuote() bool {
	if o != nil && o.Quote != nil {
		return true
	}

	return false
}

// SetQuote gets a reference to the given string and assigns it to the Quote field.
func (o *InlineResponse20013) SetQuote(v string) {
	o.Quote = &v
}

// GetVolumeBase returns the VolumeBase field value if set, zero value otherwise.
func (o *InlineResponse20013) GetVolumeBase() string {
	if o == nil || o.VolumeBase == nil {
		var ret string
		return ret
	}
	return *o.VolumeBase
}

// GetVolumeBaseOk returns a tuple with the VolumeBase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20013) GetVolumeBaseOk() (*string, bool) {
	if o == nil || o.VolumeBase == nil {
		return nil, false
	}
	return o.VolumeBase, true
}

// HasVolumeBase returns a boolean if a field has been set.
func (o *InlineResponse20013) HasVolumeBase() bool {
	if o != nil && o.VolumeBase != nil {
		return true
	}

	return false
}

// SetVolumeBase gets a reference to the given string and assigns it to the VolumeBase field.
func (o *InlineResponse20013) SetVolumeBase(v string) {
	o.VolumeBase = &v
}

// GetVolumeUsd returns the VolumeUsd field value if set, zero value otherwise.
func (o *InlineResponse20013) GetVolumeUsd() string {
	if o == nil || o.VolumeUsd == nil {
		var ret string
		return ret
	}
	return *o.VolumeUsd
}

// GetVolumeUsdOk returns a tuple with the VolumeUsd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20013) GetVolumeUsdOk() (*string, bool) {
	if o == nil || o.VolumeUsd == nil {
		return nil, false
	}
	return o.VolumeUsd, true
}

// HasVolumeUsd returns a boolean if a field has been set.
func (o *InlineResponse20013) HasVolumeUsd() bool {
	if o != nil && o.VolumeUsd != nil {
		return true
	}

	return false
}

// SetVolumeUsd gets a reference to the given string and assigns it to the VolumeUsd field.
func (o *InlineResponse20013) SetVolumeUsd(v string) {
	o.VolumeUsd = &v
}

// GetOpenQuote returns the OpenQuote field value if set, zero value otherwise.
func (o *InlineResponse20013) GetOpenQuote() string {
	if o == nil || o.OpenQuote == nil {
		var ret string
		return ret
	}
	return *o.OpenQuote
}

// GetOpenQuoteOk returns a tuple with the OpenQuote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20013) GetOpenQuoteOk() (*string, bool) {
	if o == nil || o.OpenQuote == nil {
		return nil, false
	}
	return o.OpenQuote, true
}

// HasOpenQuote returns a boolean if a field has been set.
func (o *InlineResponse20013) HasOpenQuote() bool {
	if o != nil && o.OpenQuote != nil {
		return true
	}

	return false
}

// SetOpenQuote gets a reference to the given string and assigns it to the OpenQuote field.
func (o *InlineResponse20013) SetOpenQuote(v string) {
	o.OpenQuote = &v
}

// GetOpenTimestamp returns the OpenTimestamp field value if set, zero value otherwise.
func (o *InlineResponse20013) GetOpenTimestamp() string {
	if o == nil || o.OpenTimestamp == nil {
		var ret string
		return ret
	}
	return *o.OpenTimestamp
}

// GetOpenTimestampOk returns a tuple with the OpenTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20013) GetOpenTimestampOk() (*string, bool) {
	if o == nil || o.OpenTimestamp == nil {
		return nil, false
	}
	return o.OpenTimestamp, true
}

// HasOpenTimestamp returns a boolean if a field has been set.
func (o *InlineResponse20013) HasOpenTimestamp() bool {
	if o != nil && o.OpenTimestamp != nil {
		return true
	}

	return false
}

// SetOpenTimestamp gets a reference to the given string and assigns it to the OpenTimestamp field.
func (o *InlineResponse20013) SetOpenTimestamp(v string) {
	o.OpenTimestamp = &v
}

// GetCloseQuote returns the CloseQuote field value if set, zero value otherwise.
func (o *InlineResponse20013) GetCloseQuote() string {
	if o == nil || o.CloseQuote == nil {
		var ret string
		return ret
	}
	return *o.CloseQuote
}

// GetCloseQuoteOk returns a tuple with the CloseQuote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20013) GetCloseQuoteOk() (*string, bool) {
	if o == nil || o.CloseQuote == nil {
		return nil, false
	}
	return o.CloseQuote, true
}

// HasCloseQuote returns a boolean if a field has been set.
func (o *InlineResponse20013) HasCloseQuote() bool {
	if o != nil && o.CloseQuote != nil {
		return true
	}

	return false
}

// SetCloseQuote gets a reference to the given string and assigns it to the CloseQuote field.
func (o *InlineResponse20013) SetCloseQuote(v string) {
	o.CloseQuote = &v
}

// GetCloseTimestamp returns the CloseTimestamp field value if set, zero value otherwise.
func (o *InlineResponse20013) GetCloseTimestamp() string {
	if o == nil || o.CloseTimestamp == nil {
		var ret string
		return ret
	}
	return *o.CloseTimestamp
}

// GetCloseTimestampOk returns a tuple with the CloseTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20013) GetCloseTimestampOk() (*string, bool) {
	if o == nil || o.CloseTimestamp == nil {
		return nil, false
	}
	return o.CloseTimestamp, true
}

// HasCloseTimestamp returns a boolean if a field has been set.
func (o *InlineResponse20013) HasCloseTimestamp() bool {
	if o != nil && o.CloseTimestamp != nil {
		return true
	}

	return false
}

// SetCloseTimestamp gets a reference to the given string and assigns it to the CloseTimestamp field.
func (o *InlineResponse20013) SetCloseTimestamp(v string) {
	o.CloseTimestamp = &v
}

// GetNumTrades returns the NumTrades field value if set, zero value otherwise.
func (o *InlineResponse20013) GetNumTrades() string {
	if o == nil || o.NumTrades == nil {
		var ret string
		return ret
	}
	return *o.NumTrades
}

// GetNumTradesOk returns a tuple with the NumTrades field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20013) GetNumTradesOk() (*string, bool) {
	if o == nil || o.NumTrades == nil {
		return nil, false
	}
	return o.NumTrades, true
}

// HasNumTrades returns a boolean if a field has been set.
func (o *InlineResponse20013) HasNumTrades() bool {
	if o != nil && o.NumTrades != nil {
		return true
	}

	return false
}

// SetNumTrades gets a reference to the given string and assigns it to the NumTrades field.
func (o *InlineResponse20013) SetNumTrades(v string) {
	o.NumTrades = &v
}

func (o InlineResponse20013) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Exchange != nil {
		toSerialize["exchange"] = o.Exchange
	}
	if o.Base != nil {
		toSerialize["base"] = o.Base
	}
	if o.Quote != nil {
		toSerialize["quote"] = o.Quote
	}
	if o.VolumeBase != nil {
		toSerialize["volume_base"] = o.VolumeBase
	}
	if o.VolumeUsd != nil {
		toSerialize["volume_usd"] = o.VolumeUsd
	}
	if o.OpenQuote != nil {
		toSerialize["open_quote"] = o.OpenQuote
	}
	if o.OpenTimestamp != nil {
		toSerialize["open_timestamp"] = o.OpenTimestamp
	}
	if o.CloseQuote != nil {
		toSerialize["close_quote"] = o.CloseQuote
	}
	if o.CloseTimestamp != nil {
		toSerialize["close_timestamp"] = o.CloseTimestamp
	}
	if o.NumTrades != nil {
		toSerialize["num_trades"] = o.NumTrades
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20013 struct {
	value *InlineResponse20013
	isSet bool
}

func (v NullableInlineResponse20013) Get() *InlineResponse20013 {
	return v.value
}

func (v *NullableInlineResponse20013) Set(val *InlineResponse20013) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20013) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20013) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20013(val *InlineResponse20013) *NullableInlineResponse20013 {
	return &NullableInlineResponse20013{value: val, isSet: true}
}

func (v NullableInlineResponse20013) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20013) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
