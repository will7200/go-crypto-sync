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

// ExchangesTickerInterval Interval-specific ticker attributes. The key will be the name of the interval (IE: `1d` or `365d`) and will occur for every requested interval.
type ExchangesTickerInterval struct {
	// Total volume for the given interval in USD
	Volume *string `json:"volume,omitempty"`
	// Amount of volume change for the given interval in USD
	VolumeChange *string `json:"volume_change,omitempty"`
	// Percent change of volume for the given interval
	VolumeChangePct *string `json:"volume_change_pct,omitempty"`
	// Total spot market volume for the given interval in USD
	SpotVolume *string `json:"spot_volume,omitempty"`
	// Amount of spot market volume change for the given interval in USD
	SpotVolumeChange *string `json:"spot_volume_change,omitempty"`
	// Percent change of spot market volume for the given interval
	SpotVolumeChangePct *string `json:"spot_volume_change_pct,omitempty"`
	// Total derivative market volume for the given interval in USD
	DerivativeVolume *string `json:"derivative_volume,omitempty"`
	// Amount of derivative market volume change for the given interval in USD
	DerivativeVolumeChange *string `json:"derivative_volume_change,omitempty"`
	// Percent change of derivative market volume for the given interval
	DerivativeVolumeChangePct *string `json:"derivative_volume_change_pct,omitempty"`
	// Total trades for the given interval in USD
	Trades *string `json:"trades,omitempty"`
	// Amount of trades change for the given interval in USD
	TradesChange *string `json:"trades_change,omitempty"`
	// Percent change of trades for the given interval
	TradesChangePct *string `json:"trades_change_pct,omitempty"`
}

// NewExchangesTickerInterval instantiates a new ExchangesTickerInterval object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangesTickerInterval() *ExchangesTickerInterval {
	this := ExchangesTickerInterval{}
	return &this
}

// NewExchangesTickerIntervalWithDefaults instantiates a new ExchangesTickerInterval object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangesTickerIntervalWithDefaults() *ExchangesTickerInterval {
	this := ExchangesTickerInterval{}
	return &this
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *ExchangesTickerInterval) GetVolume() string {
	if o == nil || o.Volume == nil {
		var ret string
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangesTickerInterval) GetVolumeOk() (*string, bool) {
	if o == nil || o.Volume == nil {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *ExchangesTickerInterval) HasVolume() bool {
	if o != nil && o.Volume != nil {
		return true
	}

	return false
}

// SetVolume gets a reference to the given string and assigns it to the Volume field.
func (o *ExchangesTickerInterval) SetVolume(v string) {
	o.Volume = &v
}

// GetVolumeChange returns the VolumeChange field value if set, zero value otherwise.
func (o *ExchangesTickerInterval) GetVolumeChange() string {
	if o == nil || o.VolumeChange == nil {
		var ret string
		return ret
	}
	return *o.VolumeChange
}

// GetVolumeChangeOk returns a tuple with the VolumeChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangesTickerInterval) GetVolumeChangeOk() (*string, bool) {
	if o == nil || o.VolumeChange == nil {
		return nil, false
	}
	return o.VolumeChange, true
}

// HasVolumeChange returns a boolean if a field has been set.
func (o *ExchangesTickerInterval) HasVolumeChange() bool {
	if o != nil && o.VolumeChange != nil {
		return true
	}

	return false
}

// SetVolumeChange gets a reference to the given string and assigns it to the VolumeChange field.
func (o *ExchangesTickerInterval) SetVolumeChange(v string) {
	o.VolumeChange = &v
}

// GetVolumeChangePct returns the VolumeChangePct field value if set, zero value otherwise.
func (o *ExchangesTickerInterval) GetVolumeChangePct() string {
	if o == nil || o.VolumeChangePct == nil {
		var ret string
		return ret
	}
	return *o.VolumeChangePct
}

// GetVolumeChangePctOk returns a tuple with the VolumeChangePct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangesTickerInterval) GetVolumeChangePctOk() (*string, bool) {
	if o == nil || o.VolumeChangePct == nil {
		return nil, false
	}
	return o.VolumeChangePct, true
}

// HasVolumeChangePct returns a boolean if a field has been set.
func (o *ExchangesTickerInterval) HasVolumeChangePct() bool {
	if o != nil && o.VolumeChangePct != nil {
		return true
	}

	return false
}

// SetVolumeChangePct gets a reference to the given string and assigns it to the VolumeChangePct field.
func (o *ExchangesTickerInterval) SetVolumeChangePct(v string) {
	o.VolumeChangePct = &v
}

// GetSpotVolume returns the SpotVolume field value if set, zero value otherwise.
func (o *ExchangesTickerInterval) GetSpotVolume() string {
	if o == nil || o.SpotVolume == nil {
		var ret string
		return ret
	}
	return *o.SpotVolume
}

// GetSpotVolumeOk returns a tuple with the SpotVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangesTickerInterval) GetSpotVolumeOk() (*string, bool) {
	if o == nil || o.SpotVolume == nil {
		return nil, false
	}
	return o.SpotVolume, true
}

// HasSpotVolume returns a boolean if a field has been set.
func (o *ExchangesTickerInterval) HasSpotVolume() bool {
	if o != nil && o.SpotVolume != nil {
		return true
	}

	return false
}

// SetSpotVolume gets a reference to the given string and assigns it to the SpotVolume field.
func (o *ExchangesTickerInterval) SetSpotVolume(v string) {
	o.SpotVolume = &v
}

// GetSpotVolumeChange returns the SpotVolumeChange field value if set, zero value otherwise.
func (o *ExchangesTickerInterval) GetSpotVolumeChange() string {
	if o == nil || o.SpotVolumeChange == nil {
		var ret string
		return ret
	}
	return *o.SpotVolumeChange
}

// GetSpotVolumeChangeOk returns a tuple with the SpotVolumeChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangesTickerInterval) GetSpotVolumeChangeOk() (*string, bool) {
	if o == nil || o.SpotVolumeChange == nil {
		return nil, false
	}
	return o.SpotVolumeChange, true
}

// HasSpotVolumeChange returns a boolean if a field has been set.
func (o *ExchangesTickerInterval) HasSpotVolumeChange() bool {
	if o != nil && o.SpotVolumeChange != nil {
		return true
	}

	return false
}

// SetSpotVolumeChange gets a reference to the given string and assigns it to the SpotVolumeChange field.
func (o *ExchangesTickerInterval) SetSpotVolumeChange(v string) {
	o.SpotVolumeChange = &v
}

// GetSpotVolumeChangePct returns the SpotVolumeChangePct field value if set, zero value otherwise.
func (o *ExchangesTickerInterval) GetSpotVolumeChangePct() string {
	if o == nil || o.SpotVolumeChangePct == nil {
		var ret string
		return ret
	}
	return *o.SpotVolumeChangePct
}

// GetSpotVolumeChangePctOk returns a tuple with the SpotVolumeChangePct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangesTickerInterval) GetSpotVolumeChangePctOk() (*string, bool) {
	if o == nil || o.SpotVolumeChangePct == nil {
		return nil, false
	}
	return o.SpotVolumeChangePct, true
}

// HasSpotVolumeChangePct returns a boolean if a field has been set.
func (o *ExchangesTickerInterval) HasSpotVolumeChangePct() bool {
	if o != nil && o.SpotVolumeChangePct != nil {
		return true
	}

	return false
}

// SetSpotVolumeChangePct gets a reference to the given string and assigns it to the SpotVolumeChangePct field.
func (o *ExchangesTickerInterval) SetSpotVolumeChangePct(v string) {
	o.SpotVolumeChangePct = &v
}

// GetDerivativeVolume returns the DerivativeVolume field value if set, zero value otherwise.
func (o *ExchangesTickerInterval) GetDerivativeVolume() string {
	if o == nil || o.DerivativeVolume == nil {
		var ret string
		return ret
	}
	return *o.DerivativeVolume
}

// GetDerivativeVolumeOk returns a tuple with the DerivativeVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangesTickerInterval) GetDerivativeVolumeOk() (*string, bool) {
	if o == nil || o.DerivativeVolume == nil {
		return nil, false
	}
	return o.DerivativeVolume, true
}

// HasDerivativeVolume returns a boolean if a field has been set.
func (o *ExchangesTickerInterval) HasDerivativeVolume() bool {
	if o != nil && o.DerivativeVolume != nil {
		return true
	}

	return false
}

// SetDerivativeVolume gets a reference to the given string and assigns it to the DerivativeVolume field.
func (o *ExchangesTickerInterval) SetDerivativeVolume(v string) {
	o.DerivativeVolume = &v
}

// GetDerivativeVolumeChange returns the DerivativeVolumeChange field value if set, zero value otherwise.
func (o *ExchangesTickerInterval) GetDerivativeVolumeChange() string {
	if o == nil || o.DerivativeVolumeChange == nil {
		var ret string
		return ret
	}
	return *o.DerivativeVolumeChange
}

// GetDerivativeVolumeChangeOk returns a tuple with the DerivativeVolumeChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangesTickerInterval) GetDerivativeVolumeChangeOk() (*string, bool) {
	if o == nil || o.DerivativeVolumeChange == nil {
		return nil, false
	}
	return o.DerivativeVolumeChange, true
}

// HasDerivativeVolumeChange returns a boolean if a field has been set.
func (o *ExchangesTickerInterval) HasDerivativeVolumeChange() bool {
	if o != nil && o.DerivativeVolumeChange != nil {
		return true
	}

	return false
}

// SetDerivativeVolumeChange gets a reference to the given string and assigns it to the DerivativeVolumeChange field.
func (o *ExchangesTickerInterval) SetDerivativeVolumeChange(v string) {
	o.DerivativeVolumeChange = &v
}

// GetDerivativeVolumeChangePct returns the DerivativeVolumeChangePct field value if set, zero value otherwise.
func (o *ExchangesTickerInterval) GetDerivativeVolumeChangePct() string {
	if o == nil || o.DerivativeVolumeChangePct == nil {
		var ret string
		return ret
	}
	return *o.DerivativeVolumeChangePct
}

// GetDerivativeVolumeChangePctOk returns a tuple with the DerivativeVolumeChangePct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangesTickerInterval) GetDerivativeVolumeChangePctOk() (*string, bool) {
	if o == nil || o.DerivativeVolumeChangePct == nil {
		return nil, false
	}
	return o.DerivativeVolumeChangePct, true
}

// HasDerivativeVolumeChangePct returns a boolean if a field has been set.
func (o *ExchangesTickerInterval) HasDerivativeVolumeChangePct() bool {
	if o != nil && o.DerivativeVolumeChangePct != nil {
		return true
	}

	return false
}

// SetDerivativeVolumeChangePct gets a reference to the given string and assigns it to the DerivativeVolumeChangePct field.
func (o *ExchangesTickerInterval) SetDerivativeVolumeChangePct(v string) {
	o.DerivativeVolumeChangePct = &v
}

// GetTrades returns the Trades field value if set, zero value otherwise.
func (o *ExchangesTickerInterval) GetTrades() string {
	if o == nil || o.Trades == nil {
		var ret string
		return ret
	}
	return *o.Trades
}

// GetTradesOk returns a tuple with the Trades field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangesTickerInterval) GetTradesOk() (*string, bool) {
	if o == nil || o.Trades == nil {
		return nil, false
	}
	return o.Trades, true
}

// HasTrades returns a boolean if a field has been set.
func (o *ExchangesTickerInterval) HasTrades() bool {
	if o != nil && o.Trades != nil {
		return true
	}

	return false
}

// SetTrades gets a reference to the given string and assigns it to the Trades field.
func (o *ExchangesTickerInterval) SetTrades(v string) {
	o.Trades = &v
}

// GetTradesChange returns the TradesChange field value if set, zero value otherwise.
func (o *ExchangesTickerInterval) GetTradesChange() string {
	if o == nil || o.TradesChange == nil {
		var ret string
		return ret
	}
	return *o.TradesChange
}

// GetTradesChangeOk returns a tuple with the TradesChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangesTickerInterval) GetTradesChangeOk() (*string, bool) {
	if o == nil || o.TradesChange == nil {
		return nil, false
	}
	return o.TradesChange, true
}

// HasTradesChange returns a boolean if a field has been set.
func (o *ExchangesTickerInterval) HasTradesChange() bool {
	if o != nil && o.TradesChange != nil {
		return true
	}

	return false
}

// SetTradesChange gets a reference to the given string and assigns it to the TradesChange field.
func (o *ExchangesTickerInterval) SetTradesChange(v string) {
	o.TradesChange = &v
}

// GetTradesChangePct returns the TradesChangePct field value if set, zero value otherwise.
func (o *ExchangesTickerInterval) GetTradesChangePct() string {
	if o == nil || o.TradesChangePct == nil {
		var ret string
		return ret
	}
	return *o.TradesChangePct
}

// GetTradesChangePctOk returns a tuple with the TradesChangePct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangesTickerInterval) GetTradesChangePctOk() (*string, bool) {
	if o == nil || o.TradesChangePct == nil {
		return nil, false
	}
	return o.TradesChangePct, true
}

// HasTradesChangePct returns a boolean if a field has been set.
func (o *ExchangesTickerInterval) HasTradesChangePct() bool {
	if o != nil && o.TradesChangePct != nil {
		return true
	}

	return false
}

// SetTradesChangePct gets a reference to the given string and assigns it to the TradesChangePct field.
func (o *ExchangesTickerInterval) SetTradesChangePct(v string) {
	o.TradesChangePct = &v
}

func (o ExchangesTickerInterval) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Volume != nil {
		toSerialize["volume"] = o.Volume
	}
	if o.VolumeChange != nil {
		toSerialize["volume_change"] = o.VolumeChange
	}
	if o.VolumeChangePct != nil {
		toSerialize["volume_change_pct"] = o.VolumeChangePct
	}
	if o.SpotVolume != nil {
		toSerialize["spot_volume"] = o.SpotVolume
	}
	if o.SpotVolumeChange != nil {
		toSerialize["spot_volume_change"] = o.SpotVolumeChange
	}
	if o.SpotVolumeChangePct != nil {
		toSerialize["spot_volume_change_pct"] = o.SpotVolumeChangePct
	}
	if o.DerivativeVolume != nil {
		toSerialize["derivative_volume"] = o.DerivativeVolume
	}
	if o.DerivativeVolumeChange != nil {
		toSerialize["derivative_volume_change"] = o.DerivativeVolumeChange
	}
	if o.DerivativeVolumeChangePct != nil {
		toSerialize["derivative_volume_change_pct"] = o.DerivativeVolumeChangePct
	}
	if o.Trades != nil {
		toSerialize["trades"] = o.Trades
	}
	if o.TradesChange != nil {
		toSerialize["trades_change"] = o.TradesChange
	}
	if o.TradesChangePct != nil {
		toSerialize["trades_change_pct"] = o.TradesChangePct
	}
	return json.Marshal(toSerialize)
}

type NullableExchangesTickerInterval struct {
	value *ExchangesTickerInterval
	isSet bool
}

func (v NullableExchangesTickerInterval) Get() *ExchangesTickerInterval {
	return v.value
}

func (v *NullableExchangesTickerInterval) Set(val *ExchangesTickerInterval) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangesTickerInterval) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangesTickerInterval) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangesTickerInterval(val *ExchangesTickerInterval) *NullableExchangesTickerInterval {
	return &NullableExchangesTickerInterval{value: val, isSet: true}
}

func (v NullableExchangesTickerInterval) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangesTickerInterval) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}