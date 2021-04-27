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

// InlineResponse20010 struct for InlineResponse20010
type InlineResponse20010 struct {
	// Nomics Exchange ID
	Exchange *string `json:"exchange,omitempty"`
	// Nomics Market ID
	Market *string `json:"market,omitempty"`
	// Market Type
	Type *string `json:"type,omitempty"`
	// Market Subtype
	Subtype *string `json:"subtype,omitempty"`
	// Indicates if the market is used to price the base and quote currencies. See [Nomics Pricing Methodology](https://blog.nomics.com/newsletter/nomics-pricing-methodology-explained/) for more information.
	Aggregated *bool `json:"aggregated,omitempty"`
	// Indicates if the market has been excluded from pricing the base and quote currencies.
	PriceExclude *bool `json:"price_exclude,omitempty"`
	// Indicates if the market has been excluded from counting towards volume for base and quote currencies.
	VolumeExclude *bool `json:"volume_exclude,omitempty"`
	// Nomics Currency ID of the base of the market
	Base *string `json:"base,omitempty"`
	// Nomics Currency ID of the quote of the market
	Quote *string `json:"quote,omitempty"`
	// Nomics display symbol of the base of the market
	BaseSymbol *string `json:"base_symbol,omitempty"`
	// Nomics display symbol of the quote of the market
	QuoteSymbol *string `json:"quote_symbol,omitempty"`
	// Latest price of the market in USD
	Price *string `json:"price,omitempty"`
	// Latest price of the market in the quote currency
	PriceQuote *string `json:"price_quote,omitempty"`
	// Market volume in USD based on the first `interval` (or `1d` if none specified)
	VolumeUsd *string `json:"volume_usd,omitempty"`
	// Timestamp of when the market was last updated
	LastUpdated *string                        `json:"last_updated,omitempty"`
	Interval    *ExchangeMarketsTickerInterval `json:"interval,omitempty"`
}

// NewInlineResponse20010 instantiates a new InlineResponse20010 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20010() *InlineResponse20010 {
	this := InlineResponse20010{}
	return &this
}

// NewInlineResponse20010WithDefaults instantiates a new InlineResponse20010 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20010WithDefaults() *InlineResponse20010 {
	this := InlineResponse20010{}
	return &this
}

// GetExchange returns the Exchange field value if set, zero value otherwise.
func (o *InlineResponse20010) GetExchange() string {
	if o == nil || o.Exchange == nil {
		var ret string
		return ret
	}
	return *o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20010) GetExchangeOk() (*string, bool) {
	if o == nil || o.Exchange == nil {
		return nil, false
	}
	return o.Exchange, true
}

// HasExchange returns a boolean if a field has been set.
func (o *InlineResponse20010) HasExchange() bool {
	if o != nil && o.Exchange != nil {
		return true
	}

	return false
}

// SetExchange gets a reference to the given string and assigns it to the Exchange field.
func (o *InlineResponse20010) SetExchange(v string) {
	o.Exchange = &v
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *InlineResponse20010) GetMarket() string {
	if o == nil || o.Market == nil {
		var ret string
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20010) GetMarketOk() (*string, bool) {
	if o == nil || o.Market == nil {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *InlineResponse20010) HasMarket() bool {
	if o != nil && o.Market != nil {
		return true
	}

	return false
}

// SetMarket gets a reference to the given string and assigns it to the Market field.
func (o *InlineResponse20010) SetMarket(v string) {
	o.Market = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineResponse20010) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20010) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineResponse20010) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineResponse20010) SetType(v string) {
	o.Type = &v
}

// GetSubtype returns the Subtype field value if set, zero value otherwise.
func (o *InlineResponse20010) GetSubtype() string {
	if o == nil || o.Subtype == nil {
		var ret string
		return ret
	}
	return *o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20010) GetSubtypeOk() (*string, bool) {
	if o == nil || o.Subtype == nil {
		return nil, false
	}
	return o.Subtype, true
}

// HasSubtype returns a boolean if a field has been set.
func (o *InlineResponse20010) HasSubtype() bool {
	if o != nil && o.Subtype != nil {
		return true
	}

	return false
}

// SetSubtype gets a reference to the given string and assigns it to the Subtype field.
func (o *InlineResponse20010) SetSubtype(v string) {
	o.Subtype = &v
}

// GetAggregated returns the Aggregated field value if set, zero value otherwise.
func (o *InlineResponse20010) GetAggregated() bool {
	if o == nil || o.Aggregated == nil {
		var ret bool
		return ret
	}
	return *o.Aggregated
}

// GetAggregatedOk returns a tuple with the Aggregated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20010) GetAggregatedOk() (*bool, bool) {
	if o == nil || o.Aggregated == nil {
		return nil, false
	}
	return o.Aggregated, true
}

// HasAggregated returns a boolean if a field has been set.
func (o *InlineResponse20010) HasAggregated() bool {
	if o != nil && o.Aggregated != nil {
		return true
	}

	return false
}

// SetAggregated gets a reference to the given bool and assigns it to the Aggregated field.
func (o *InlineResponse20010) SetAggregated(v bool) {
	o.Aggregated = &v
}

// GetPriceExclude returns the PriceExclude field value if set, zero value otherwise.
func (o *InlineResponse20010) GetPriceExclude() bool {
	if o == nil || o.PriceExclude == nil {
		var ret bool
		return ret
	}
	return *o.PriceExclude
}

// GetPriceExcludeOk returns a tuple with the PriceExclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20010) GetPriceExcludeOk() (*bool, bool) {
	if o == nil || o.PriceExclude == nil {
		return nil, false
	}
	return o.PriceExclude, true
}

// HasPriceExclude returns a boolean if a field has been set.
func (o *InlineResponse20010) HasPriceExclude() bool {
	if o != nil && o.PriceExclude != nil {
		return true
	}

	return false
}

// SetPriceExclude gets a reference to the given bool and assigns it to the PriceExclude field.
func (o *InlineResponse20010) SetPriceExclude(v bool) {
	o.PriceExclude = &v
}

// GetVolumeExclude returns the VolumeExclude field value if set, zero value otherwise.
func (o *InlineResponse20010) GetVolumeExclude() bool {
	if o == nil || o.VolumeExclude == nil {
		var ret bool
		return ret
	}
	return *o.VolumeExclude
}

// GetVolumeExcludeOk returns a tuple with the VolumeExclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20010) GetVolumeExcludeOk() (*bool, bool) {
	if o == nil || o.VolumeExclude == nil {
		return nil, false
	}
	return o.VolumeExclude, true
}

// HasVolumeExclude returns a boolean if a field has been set.
func (o *InlineResponse20010) HasVolumeExclude() bool {
	if o != nil && o.VolumeExclude != nil {
		return true
	}

	return false
}

// SetVolumeExclude gets a reference to the given bool and assigns it to the VolumeExclude field.
func (o *InlineResponse20010) SetVolumeExclude(v bool) {
	o.VolumeExclude = &v
}

// GetBase returns the Base field value if set, zero value otherwise.
func (o *InlineResponse20010) GetBase() string {
	if o == nil || o.Base == nil {
		var ret string
		return ret
	}
	return *o.Base
}

// GetBaseOk returns a tuple with the Base field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20010) GetBaseOk() (*string, bool) {
	if o == nil || o.Base == nil {
		return nil, false
	}
	return o.Base, true
}

// HasBase returns a boolean if a field has been set.
func (o *InlineResponse20010) HasBase() bool {
	if o != nil && o.Base != nil {
		return true
	}

	return false
}

// SetBase gets a reference to the given string and assigns it to the Base field.
func (o *InlineResponse20010) SetBase(v string) {
	o.Base = &v
}

// GetQuote returns the Quote field value if set, zero value otherwise.
func (o *InlineResponse20010) GetQuote() string {
	if o == nil || o.Quote == nil {
		var ret string
		return ret
	}
	return *o.Quote
}

// GetQuoteOk returns a tuple with the Quote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20010) GetQuoteOk() (*string, bool) {
	if o == nil || o.Quote == nil {
		return nil, false
	}
	return o.Quote, true
}

// HasQuote returns a boolean if a field has been set.
func (o *InlineResponse20010) HasQuote() bool {
	if o != nil && o.Quote != nil {
		return true
	}

	return false
}

// SetQuote gets a reference to the given string and assigns it to the Quote field.
func (o *InlineResponse20010) SetQuote(v string) {
	o.Quote = &v
}

// GetBaseSymbol returns the BaseSymbol field value if set, zero value otherwise.
func (o *InlineResponse20010) GetBaseSymbol() string {
	if o == nil || o.BaseSymbol == nil {
		var ret string
		return ret
	}
	return *o.BaseSymbol
}

// GetBaseSymbolOk returns a tuple with the BaseSymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20010) GetBaseSymbolOk() (*string, bool) {
	if o == nil || o.BaseSymbol == nil {
		return nil, false
	}
	return o.BaseSymbol, true
}

// HasBaseSymbol returns a boolean if a field has been set.
func (o *InlineResponse20010) HasBaseSymbol() bool {
	if o != nil && o.BaseSymbol != nil {
		return true
	}

	return false
}

// SetBaseSymbol gets a reference to the given string and assigns it to the BaseSymbol field.
func (o *InlineResponse20010) SetBaseSymbol(v string) {
	o.BaseSymbol = &v
}

// GetQuoteSymbol returns the QuoteSymbol field value if set, zero value otherwise.
func (o *InlineResponse20010) GetQuoteSymbol() string {
	if o == nil || o.QuoteSymbol == nil {
		var ret string
		return ret
	}
	return *o.QuoteSymbol
}

// GetQuoteSymbolOk returns a tuple with the QuoteSymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20010) GetQuoteSymbolOk() (*string, bool) {
	if o == nil || o.QuoteSymbol == nil {
		return nil, false
	}
	return o.QuoteSymbol, true
}

// HasQuoteSymbol returns a boolean if a field has been set.
func (o *InlineResponse20010) HasQuoteSymbol() bool {
	if o != nil && o.QuoteSymbol != nil {
		return true
	}

	return false
}

// SetQuoteSymbol gets a reference to the given string and assigns it to the QuoteSymbol field.
func (o *InlineResponse20010) SetQuoteSymbol(v string) {
	o.QuoteSymbol = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *InlineResponse20010) GetPrice() string {
	if o == nil || o.Price == nil {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20010) GetPriceOk() (*string, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *InlineResponse20010) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *InlineResponse20010) SetPrice(v string) {
	o.Price = &v
}

// GetPriceQuote returns the PriceQuote field value if set, zero value otherwise.
func (o *InlineResponse20010) GetPriceQuote() string {
	if o == nil || o.PriceQuote == nil {
		var ret string
		return ret
	}
	return *o.PriceQuote
}

// GetPriceQuoteOk returns a tuple with the PriceQuote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20010) GetPriceQuoteOk() (*string, bool) {
	if o == nil || o.PriceQuote == nil {
		return nil, false
	}
	return o.PriceQuote, true
}

// HasPriceQuote returns a boolean if a field has been set.
func (o *InlineResponse20010) HasPriceQuote() bool {
	if o != nil && o.PriceQuote != nil {
		return true
	}

	return false
}

// SetPriceQuote gets a reference to the given string and assigns it to the PriceQuote field.
func (o *InlineResponse20010) SetPriceQuote(v string) {
	o.PriceQuote = &v
}

// GetVolumeUsd returns the VolumeUsd field value if set, zero value otherwise.
func (o *InlineResponse20010) GetVolumeUsd() string {
	if o == nil || o.VolumeUsd == nil {
		var ret string
		return ret
	}
	return *o.VolumeUsd
}

// GetVolumeUsdOk returns a tuple with the VolumeUsd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20010) GetVolumeUsdOk() (*string, bool) {
	if o == nil || o.VolumeUsd == nil {
		return nil, false
	}
	return o.VolumeUsd, true
}

// HasVolumeUsd returns a boolean if a field has been set.
func (o *InlineResponse20010) HasVolumeUsd() bool {
	if o != nil && o.VolumeUsd != nil {
		return true
	}

	return false
}

// SetVolumeUsd gets a reference to the given string and assigns it to the VolumeUsd field.
func (o *InlineResponse20010) SetVolumeUsd(v string) {
	o.VolumeUsd = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *InlineResponse20010) GetLastUpdated() string {
	if o == nil || o.LastUpdated == nil {
		var ret string
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20010) GetLastUpdatedOk() (*string, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *InlineResponse20010) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given string and assigns it to the LastUpdated field.
func (o *InlineResponse20010) SetLastUpdated(v string) {
	o.LastUpdated = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *InlineResponse20010) GetInterval() ExchangeMarketsTickerInterval {
	if o == nil || o.Interval == nil {
		var ret ExchangeMarketsTickerInterval
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20010) GetIntervalOk() (*ExchangeMarketsTickerInterval, bool) {
	if o == nil || o.Interval == nil {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *InlineResponse20010) HasInterval() bool {
	if o != nil && o.Interval != nil {
		return true
	}

	return false
}

// SetInterval gets a reference to the given ExchangeMarketsTickerInterval and assigns it to the Interval field.
func (o *InlineResponse20010) SetInterval(v ExchangeMarketsTickerInterval) {
	o.Interval = &v
}

func (o InlineResponse20010) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Exchange != nil {
		toSerialize["exchange"] = o.Exchange
	}
	if o.Market != nil {
		toSerialize["market"] = o.Market
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Subtype != nil {
		toSerialize["subtype"] = o.Subtype
	}
	if o.Aggregated != nil {
		toSerialize["aggregated"] = o.Aggregated
	}
	if o.PriceExclude != nil {
		toSerialize["price_exclude"] = o.PriceExclude
	}
	if o.VolumeExclude != nil {
		toSerialize["volume_exclude"] = o.VolumeExclude
	}
	if o.Base != nil {
		toSerialize["base"] = o.Base
	}
	if o.Quote != nil {
		toSerialize["quote"] = o.Quote
	}
	if o.BaseSymbol != nil {
		toSerialize["base_symbol"] = o.BaseSymbol
	}
	if o.QuoteSymbol != nil {
		toSerialize["quote_symbol"] = o.QuoteSymbol
	}
	if o.Price != nil {
		toSerialize["price"] = o.Price
	}
	if o.PriceQuote != nil {
		toSerialize["price_quote"] = o.PriceQuote
	}
	if o.VolumeUsd != nil {
		toSerialize["volume_usd"] = o.VolumeUsd
	}
	if o.LastUpdated != nil {
		toSerialize["last_updated"] = o.LastUpdated
	}
	if o.Interval != nil {
		toSerialize["interval"] = o.Interval
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20010 struct {
	value *InlineResponse20010
	isSet bool
}

func (v NullableInlineResponse20010) Get() *InlineResponse20010 {
	return v.value
}

func (v *NullableInlineResponse20010) Set(val *InlineResponse20010) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20010) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20010) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20010(val *InlineResponse20010) *NullableInlineResponse20010 {
	return &NullableInlineResponse20010{value: val, isSet: true}
}

func (v NullableInlineResponse20010) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20010) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
