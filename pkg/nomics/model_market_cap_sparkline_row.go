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

// MarketCapSparklineRow struct for MarketCapSparklineRow
type MarketCapSparklineRow struct {
	// Currency ID
	Currency *string `json:"currency,omitempty"`
	// Time values matching the close values
	Timestamps *[]string `json:"timestamps,omitempty"`
	// Closing market cap in USD corresponding to timestamp value
	Closes *[]string `json:"closes,omitempty"`
}

// NewMarketCapSparklineRow instantiates a new MarketCapSparklineRow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketCapSparklineRow() *MarketCapSparklineRow {
	this := MarketCapSparklineRow{}
	return &this
}

// NewMarketCapSparklineRowWithDefaults instantiates a new MarketCapSparklineRow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketCapSparklineRowWithDefaults() *MarketCapSparklineRow {
	this := MarketCapSparklineRow{}
	return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *MarketCapSparklineRow) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketCapSparklineRow) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *MarketCapSparklineRow) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *MarketCapSparklineRow) SetCurrency(v string) {
	o.Currency = &v
}

// GetTimestamps returns the Timestamps field value if set, zero value otherwise.
func (o *MarketCapSparklineRow) GetTimestamps() []string {
	if o == nil || o.Timestamps == nil {
		var ret []string
		return ret
	}
	return *o.Timestamps
}

// GetTimestampsOk returns a tuple with the Timestamps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketCapSparklineRow) GetTimestampsOk() (*[]string, bool) {
	if o == nil || o.Timestamps == nil {
		return nil, false
	}
	return o.Timestamps, true
}

// HasTimestamps returns a boolean if a field has been set.
func (o *MarketCapSparklineRow) HasTimestamps() bool {
	if o != nil && o.Timestamps != nil {
		return true
	}

	return false
}

// SetTimestamps gets a reference to the given []string and assigns it to the Timestamps field.
func (o *MarketCapSparklineRow) SetTimestamps(v []string) {
	o.Timestamps = &v
}

// GetCloses returns the Closes field value if set, zero value otherwise.
func (o *MarketCapSparklineRow) GetCloses() []string {
	if o == nil || o.Closes == nil {
		var ret []string
		return ret
	}
	return *o.Closes
}

// GetClosesOk returns a tuple with the Closes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketCapSparklineRow) GetClosesOk() (*[]string, bool) {
	if o == nil || o.Closes == nil {
		return nil, false
	}
	return o.Closes, true
}

// HasCloses returns a boolean if a field has been set.
func (o *MarketCapSparklineRow) HasCloses() bool {
	if o != nil && o.Closes != nil {
		return true
	}

	return false
}

// SetCloses gets a reference to the given []string and assigns it to the Closes field.
func (o *MarketCapSparklineRow) SetCloses(v []string) {
	o.Closes = &v
}

func (o MarketCapSparklineRow) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	if o.Timestamps != nil {
		toSerialize["timestamps"] = o.Timestamps
	}
	if o.Closes != nil {
		toSerialize["closes"] = o.Closes
	}
	return json.Marshal(toSerialize)
}

type NullableMarketCapSparklineRow struct {
	value *MarketCapSparklineRow
	isSet bool
}

func (v NullableMarketCapSparklineRow) Get() *MarketCapSparklineRow {
	return v.value
}

func (v *NullableMarketCapSparklineRow) Set(val *MarketCapSparklineRow) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketCapSparklineRow) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketCapSparklineRow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketCapSparklineRow(val *MarketCapSparklineRow) *NullableMarketCapSparklineRow {
	return &NullableMarketCapSparklineRow{value: val, isSet: true}
}

func (v NullableMarketCapSparklineRow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketCapSparklineRow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
