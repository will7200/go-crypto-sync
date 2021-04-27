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

// InlineResponse200 struct for InlineResponse200
type InlineResponse200 struct {
	// Nomics ID of the currency
	Currency *string `json:"currency,omitempty"`
	// Unique Nomics ID of the currency
	Id *string `json:"id,omitempty"`
	// Current status
	Status *string `json:"status,omitempty"`
	// Current price
	Price *string `json:"price,omitempty"`
	// Date of the price
	PriceDate *string `json:"price_date,omitempty"`
	// Timestamp of the price
	PriceTimestamp *string `json:"price_timestamp,omitempty"`
	// Display symbol of the currency (may be duplicated)
	Symbol *string `json:"symbol,omitempty"`
	// Current circulating supply
	CirculatingSupply *string `json:"circulating_supply,omitempty"`
	// Current maximum supply
	MaxSupply *string `json:"max_supply,omitempty"`
	// Name of the currency
	Name *string `json:"name,omitempty"`
	// URL to logo of the currency
	LogoUrl *string `json:"logo_url,omitempty"`
	// Current market cap
	MarketCap *string `json:"market_cap,omitempty"`
	// Current transparent market cap
	TransparentMarketCap *string `json:"transparent_market_cap,omitempty"`
	// The number of exchanges on which this asset is traded
	NumExchanges *string `json:"num_exchanges,omitempty"`
	// Number of currency pairs (markets) this asset is a part of, where both assets are known
	NumPairs *string `json:"num_pairs,omitempty"`
	// Number of currency pairs (markets) this asset is a part of, but where the other asset is unknown
	NumPairsUnmapped *string `json:"num_pairs_unmapped,omitempty"`
	// RFC3999 timestamp of the first `1d` candle available via the Nomics API
	FirstCandle *string `json:"first_candle,omitempty"`
	// RFC3999 timestamp of the first trade available via the Nomics API
	FirstTrade *string `json:"first_trade,omitempty"`
	// RFC3999 timestamp of the first order book snapshot available via the Nomics API
	FirstOrderBook *string `json:"first_order_book,omitempty"`
	// RFC3999 timestamp representing the currency was first priced by Nomics
	FirstPricedAt *string `json:"first_priced_at,omitempty"`
	// Rank of the currency by market cap
	Rank *string `json:"rank,omitempty"`
	// Relative rank change based on the first specified `interval` other than `1h`.  This field is only present on `active` currencies.
	RankDelta *string `json:"rank_delta,omitempty"`
	// All time high price
	High *string `json:"high,omitempty"`
	// RFC3999 timestamp of the all time high
	HighTimestamp *string                   `json:"high_timestamp,omitempty"`
	Interval      *CurrenciesTickerInterval `json:"interval,omitempty"`
}

// NewInlineResponse200 instantiates a new InlineResponse200 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200() *InlineResponse200 {
	this := InlineResponse200{}
	return &this
}

// NewInlineResponse200WithDefaults instantiates a new InlineResponse200 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200WithDefaults() *InlineResponse200 {
	this := InlineResponse200{}
	return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *InlineResponse200) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *InlineResponse200) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *InlineResponse200) SetCurrency(v string) {
	o.Currency = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse200) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InlineResponse200) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineResponse200) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InlineResponse200) SetStatus(v string) {
	o.Status = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *InlineResponse200) GetPrice() string {
	if o == nil || o.Price == nil {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200) GetPriceOk() (*string, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *InlineResponse200) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *InlineResponse200) SetPrice(v string) {
	o.Price = &v
}

// GetPriceDate returns the PriceDate field value if set, zero value otherwise.
func (o *InlineResponse200) GetPriceDate() string {
	if o == nil || o.PriceDate == nil {
		var ret string
		return ret
	}
	return *o.PriceDate
}

// GetPriceDateOk returns a tuple with the PriceDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200) GetPriceDateOk() (*string, bool) {
	if o == nil || o.PriceDate == nil {
		return nil, false
	}
	return o.PriceDate, true
}

// HasPriceDate returns a boolean if a field has been set.
func (o *InlineResponse200) HasPriceDate() bool {
	if o != nil && o.PriceDate != nil {
		return true
	}

	return false
}

// SetPriceDate gets a reference to the given string and assigns it to the PriceDate field.
func (o *InlineResponse200) SetPriceDate(v string) {
	o.PriceDate = &v
}

// GetPriceTimestamp returns the PriceTimestamp field value if set, zero value otherwise.
func (o *InlineResponse200) GetPriceTimestamp() string {
	if o == nil || o.PriceTimestamp == nil {
		var ret string
		return ret
	}
	return *o.PriceTimestamp
}

// GetPriceTimestampOk returns a tuple with the PriceTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200) GetPriceTimestampOk() (*string, bool) {
	if o == nil || o.PriceTimestamp == nil {
		return nil, false
	}
	return o.PriceTimestamp, true
}

// HasPriceTimestamp returns a boolean if a field has been set.
func (o *InlineResponse200) HasPriceTimestamp() bool {
	if o != nil && o.PriceTimestamp != nil {
		return true
	}

	return false
}

// SetPriceTimestamp gets a reference to the given string and assigns it to the PriceTimestamp field.
func (o *InlineResponse200) SetPriceTimestamp(v string) {
	o.PriceTimestamp = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *InlineResponse200) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *InlineResponse200) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *InlineResponse200) SetSymbol(v string) {
	o.Symbol = &v
}

// GetCirculatingSupply returns the CirculatingSupply field value if set, zero value otherwise.
func (o *InlineResponse200) GetCirculatingSupply() string {
	if o == nil || o.CirculatingSupply == nil {
		var ret string
		return ret
	}
	return *o.CirculatingSupply
}

// GetCirculatingSupplyOk returns a tuple with the CirculatingSupply field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200) GetCirculatingSupplyOk() (*string, bool) {
	if o == nil || o.CirculatingSupply == nil {
		return nil, false
	}
	return o.CirculatingSupply, true
}

// HasCirculatingSupply returns a boolean if a field has been set.
func (o *InlineResponse200) HasCirculatingSupply() bool {
	if o != nil && o.CirculatingSupply != nil {
		return true
	}

	return false
}

// SetCirculatingSupply gets a reference to the given string and assigns it to the CirculatingSupply field.
func (o *InlineResponse200) SetCirculatingSupply(v string) {
	o.CirculatingSupply = &v
}

// GetMaxSupply returns the MaxSupply field value if set, zero value otherwise.
func (o *InlineResponse200) GetMaxSupply() string {
	if o == nil || o.MaxSupply == nil {
		var ret string
		return ret
	}
	return *o.MaxSupply
}

// GetMaxSupplyOk returns a tuple with the MaxSupply field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200) GetMaxSupplyOk() (*string, bool) {
	if o == nil || o.MaxSupply == nil {
		return nil, false
	}
	return o.MaxSupply, true
}

// HasMaxSupply returns a boolean if a field has been set.
func (o *InlineResponse200) HasMaxSupply() bool {
	if o != nil && o.MaxSupply != nil {
		return true
	}

	return false
}

// SetMaxSupply gets a reference to the given string and assigns it to the MaxSupply field.
func (o *InlineResponse200) SetMaxSupply(v string) {
	o.MaxSupply = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200) SetName(v string) {
	o.Name = &v
}

// GetLogoUrl returns the LogoUrl field value if set, zero value otherwise.
func (o *InlineResponse200) GetLogoUrl() string {
	if o == nil || o.LogoUrl == nil {
		var ret string
		return ret
	}
	return *o.LogoUrl
}

// GetLogoUrlOk returns a tuple with the LogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200) GetLogoUrlOk() (*string, bool) {
	if o == nil || o.LogoUrl == nil {
		return nil, false
	}
	return o.LogoUrl, true
}

// HasLogoUrl returns a boolean if a field has been set.
func (o *InlineResponse200) HasLogoUrl() bool {
	if o != nil && o.LogoUrl != nil {
		return true
	}

	return false
}

// SetLogoUrl gets a reference to the given string and assigns it to the LogoUrl field.
func (o *InlineResponse200) SetLogoUrl(v string) {
	o.LogoUrl = &v
}

// GetMarketCap returns the MarketCap field value if set, zero value otherwise.
func (o *InlineResponse200) GetMarketCap() string {
	if o == nil || o.MarketCap == nil {
		var ret string
		return ret
	}
	return *o.MarketCap
}

// GetMarketCapOk returns a tuple with the MarketCap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200) GetMarketCapOk() (*string, bool) {
	if o == nil || o.MarketCap == nil {
		return nil, false
	}
	return o.MarketCap, true
}

// HasMarketCap returns a boolean if a field has been set.
func (o *InlineResponse200) HasMarketCap() bool {
	if o != nil && o.MarketCap != nil {
		return true
	}

	return false
}

// SetMarketCap gets a reference to the given string and assigns it to the MarketCap field.
func (o *InlineResponse200) SetMarketCap(v string) {
	o.MarketCap = &v
}

// GetTransparentMarketCap returns the TransparentMarketCap field value if set, zero value otherwise.
func (o *InlineResponse200) GetTransparentMarketCap() string {
	if o == nil || o.TransparentMarketCap == nil {
		var ret string
		return ret
	}
	return *o.TransparentMarketCap
}

// GetTransparentMarketCapOk returns a tuple with the TransparentMarketCap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200) GetTransparentMarketCapOk() (*string, bool) {
	if o == nil || o.TransparentMarketCap == nil {
		return nil, false
	}
	return o.TransparentMarketCap, true
}

// HasTransparentMarketCap returns a boolean if a field has been set.
func (o *InlineResponse200) HasTransparentMarketCap() bool {
	if o != nil && o.TransparentMarketCap != nil {
		return true
	}

	return false
}

// SetTransparentMarketCap gets a reference to the given string and assigns it to the TransparentMarketCap field.
func (o *InlineResponse200) SetTransparentMarketCap(v string) {
	o.TransparentMarketCap = &v
}

// GetNumExchanges returns the NumExchanges field value if set, zero value otherwise.
func (o *InlineResponse200) GetNumExchanges() string {
	if o == nil || o.NumExchanges == nil {
		var ret string
		return ret
	}
	return *o.NumExchanges
}

// GetNumExchangesOk returns a tuple with the NumExchanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200) GetNumExchangesOk() (*string, bool) {
	if o == nil || o.NumExchanges == nil {
		return nil, false
	}
	return o.NumExchanges, true
}

// HasNumExchanges returns a boolean if a field has been set.
func (o *InlineResponse200) HasNumExchanges() bool {
	if o != nil && o.NumExchanges != nil {
		return true
	}

	return false
}

// SetNumExchanges gets a reference to the given string and assigns it to the NumExchanges field.
func (o *InlineResponse200) SetNumExchanges(v string) {
	o.NumExchanges = &v
}

// GetNumPairs returns the NumPairs field value if set, zero value otherwise.
func (o *InlineResponse200) GetNumPairs() string {
	if o == nil || o.NumPairs == nil {
		var ret string
		return ret
	}
	return *o.NumPairs
}

// GetNumPairsOk returns a tuple with the NumPairs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200) GetNumPairsOk() (*string, bool) {
	if o == nil || o.NumPairs == nil {
		return nil, false
	}
	return o.NumPairs, true
}

// HasNumPairs returns a boolean if a field has been set.
func (o *InlineResponse200) HasNumPairs() bool {
	if o != nil && o.NumPairs != nil {
		return true
	}

	return false
}

// SetNumPairs gets a reference to the given string and assigns it to the NumPairs field.
func (o *InlineResponse200) SetNumPairs(v string) {
	o.NumPairs = &v
}

// GetNumPairsUnmapped returns the NumPairsUnmapped field value if set, zero value otherwise.
func (o *InlineResponse200) GetNumPairsUnmapped() string {
	if o == nil || o.NumPairsUnmapped == nil {
		var ret string
		return ret
	}
	return *o.NumPairsUnmapped
}

// GetNumPairsUnmappedOk returns a tuple with the NumPairsUnmapped field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200) GetNumPairsUnmappedOk() (*string, bool) {
	if o == nil || o.NumPairsUnmapped == nil {
		return nil, false
	}
	return o.NumPairsUnmapped, true
}

// HasNumPairsUnmapped returns a boolean if a field has been set.
func (o *InlineResponse200) HasNumPairsUnmapped() bool {
	if o != nil && o.NumPairsUnmapped != nil {
		return true
	}

	return false
}

// SetNumPairsUnmapped gets a reference to the given string and assigns it to the NumPairsUnmapped field.
func (o *InlineResponse200) SetNumPairsUnmapped(v string) {
	o.NumPairsUnmapped = &v
}

// GetFirstCandle returns the FirstCandle field value if set, zero value otherwise.
func (o *InlineResponse200) GetFirstCandle() string {
	if o == nil || o.FirstCandle == nil {
		var ret string
		return ret
	}
	return *o.FirstCandle
}

// GetFirstCandleOk returns a tuple with the FirstCandle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200) GetFirstCandleOk() (*string, bool) {
	if o == nil || o.FirstCandle == nil {
		return nil, false
	}
	return o.FirstCandle, true
}

// HasFirstCandle returns a boolean if a field has been set.
func (o *InlineResponse200) HasFirstCandle() bool {
	if o != nil && o.FirstCandle != nil {
		return true
	}

	return false
}

// SetFirstCandle gets a reference to the given string and assigns it to the FirstCandle field.
func (o *InlineResponse200) SetFirstCandle(v string) {
	o.FirstCandle = &v
}

// GetFirstTrade returns the FirstTrade field value if set, zero value otherwise.
func (o *InlineResponse200) GetFirstTrade() string {
	if o == nil || o.FirstTrade == nil {
		var ret string
		return ret
	}
	return *o.FirstTrade
}

// GetFirstTradeOk returns a tuple with the FirstTrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200) GetFirstTradeOk() (*string, bool) {
	if o == nil || o.FirstTrade == nil {
		return nil, false
	}
	return o.FirstTrade, true
}

// HasFirstTrade returns a boolean if a field has been set.
func (o *InlineResponse200) HasFirstTrade() bool {
	if o != nil && o.FirstTrade != nil {
		return true
	}

	return false
}

// SetFirstTrade gets a reference to the given string and assigns it to the FirstTrade field.
func (o *InlineResponse200) SetFirstTrade(v string) {
	o.FirstTrade = &v
}

// GetFirstOrderBook returns the FirstOrderBook field value if set, zero value otherwise.
func (o *InlineResponse200) GetFirstOrderBook() string {
	if o == nil || o.FirstOrderBook == nil {
		var ret string
		return ret
	}
	return *o.FirstOrderBook
}

// GetFirstOrderBookOk returns a tuple with the FirstOrderBook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200) GetFirstOrderBookOk() (*string, bool) {
	if o == nil || o.FirstOrderBook == nil {
		return nil, false
	}
	return o.FirstOrderBook, true
}

// HasFirstOrderBook returns a boolean if a field has been set.
func (o *InlineResponse200) HasFirstOrderBook() bool {
	if o != nil && o.FirstOrderBook != nil {
		return true
	}

	return false
}

// SetFirstOrderBook gets a reference to the given string and assigns it to the FirstOrderBook field.
func (o *InlineResponse200) SetFirstOrderBook(v string) {
	o.FirstOrderBook = &v
}

// GetFirstPricedAt returns the FirstPricedAt field value if set, zero value otherwise.
func (o *InlineResponse200) GetFirstPricedAt() string {
	if o == nil || o.FirstPricedAt == nil {
		var ret string
		return ret
	}
	return *o.FirstPricedAt
}

// GetFirstPricedAtOk returns a tuple with the FirstPricedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200) GetFirstPricedAtOk() (*string, bool) {
	if o == nil || o.FirstPricedAt == nil {
		return nil, false
	}
	return o.FirstPricedAt, true
}

// HasFirstPricedAt returns a boolean if a field has been set.
func (o *InlineResponse200) HasFirstPricedAt() bool {
	if o != nil && o.FirstPricedAt != nil {
		return true
	}

	return false
}

// SetFirstPricedAt gets a reference to the given string and assigns it to the FirstPricedAt field.
func (o *InlineResponse200) SetFirstPricedAt(v string) {
	o.FirstPricedAt = &v
}

// GetRank returns the Rank field value if set, zero value otherwise.
func (o *InlineResponse200) GetRank() string {
	if o == nil || o.Rank == nil {
		var ret string
		return ret
	}
	return *o.Rank
}

// GetRankOk returns a tuple with the Rank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200) GetRankOk() (*string, bool) {
	if o == nil || o.Rank == nil {
		return nil, false
	}
	return o.Rank, true
}

// HasRank returns a boolean if a field has been set.
func (o *InlineResponse200) HasRank() bool {
	if o != nil && o.Rank != nil {
		return true
	}

	return false
}

// SetRank gets a reference to the given string and assigns it to the Rank field.
func (o *InlineResponse200) SetRank(v string) {
	o.Rank = &v
}

// GetRankDelta returns the RankDelta field value if set, zero value otherwise.
func (o *InlineResponse200) GetRankDelta() string {
	if o == nil || o.RankDelta == nil {
		var ret string
		return ret
	}
	return *o.RankDelta
}

// GetRankDeltaOk returns a tuple with the RankDelta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200) GetRankDeltaOk() (*string, bool) {
	if o == nil || o.RankDelta == nil {
		return nil, false
	}
	return o.RankDelta, true
}

// HasRankDelta returns a boolean if a field has been set.
func (o *InlineResponse200) HasRankDelta() bool {
	if o != nil && o.RankDelta != nil {
		return true
	}

	return false
}

// SetRankDelta gets a reference to the given string and assigns it to the RankDelta field.
func (o *InlineResponse200) SetRankDelta(v string) {
	o.RankDelta = &v
}

// GetHigh returns the High field value if set, zero value otherwise.
func (o *InlineResponse200) GetHigh() string {
	if o == nil || o.High == nil {
		var ret string
		return ret
	}
	return *o.High
}

// GetHighOk returns a tuple with the High field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200) GetHighOk() (*string, bool) {
	if o == nil || o.High == nil {
		return nil, false
	}
	return o.High, true
}

// HasHigh returns a boolean if a field has been set.
func (o *InlineResponse200) HasHigh() bool {
	if o != nil && o.High != nil {
		return true
	}

	return false
}

// SetHigh gets a reference to the given string and assigns it to the High field.
func (o *InlineResponse200) SetHigh(v string) {
	o.High = &v
}

// GetHighTimestamp returns the HighTimestamp field value if set, zero value otherwise.
func (o *InlineResponse200) GetHighTimestamp() string {
	if o == nil || o.HighTimestamp == nil {
		var ret string
		return ret
	}
	return *o.HighTimestamp
}

// GetHighTimestampOk returns a tuple with the HighTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200) GetHighTimestampOk() (*string, bool) {
	if o == nil || o.HighTimestamp == nil {
		return nil, false
	}
	return o.HighTimestamp, true
}

// HasHighTimestamp returns a boolean if a field has been set.
func (o *InlineResponse200) HasHighTimestamp() bool {
	if o != nil && o.HighTimestamp != nil {
		return true
	}

	return false
}

// SetHighTimestamp gets a reference to the given string and assigns it to the HighTimestamp field.
func (o *InlineResponse200) SetHighTimestamp(v string) {
	o.HighTimestamp = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *InlineResponse200) GetInterval() CurrenciesTickerInterval {
	if o == nil || o.Interval == nil {
		var ret CurrenciesTickerInterval
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200) GetIntervalOk() (*CurrenciesTickerInterval, bool) {
	if o == nil || o.Interval == nil {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *InlineResponse200) HasInterval() bool {
	if o != nil && o.Interval != nil {
		return true
	}

	return false
}

// SetInterval gets a reference to the given CurrenciesTickerInterval and assigns it to the Interval field.
func (o *InlineResponse200) SetInterval(v CurrenciesTickerInterval) {
	o.Interval = &v
}

func (o InlineResponse200) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Price != nil {
		toSerialize["price"] = o.Price
	}
	if o.PriceDate != nil {
		toSerialize["price_date"] = o.PriceDate
	}
	if o.PriceTimestamp != nil {
		toSerialize["price_timestamp"] = o.PriceTimestamp
	}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	if o.CirculatingSupply != nil {
		toSerialize["circulating_supply"] = o.CirculatingSupply
	}
	if o.MaxSupply != nil {
		toSerialize["max_supply"] = o.MaxSupply
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.LogoUrl != nil {
		toSerialize["logo_url"] = o.LogoUrl
	}
	if o.MarketCap != nil {
		toSerialize["market_cap"] = o.MarketCap
	}
	if o.TransparentMarketCap != nil {
		toSerialize["transparent_market_cap"] = o.TransparentMarketCap
	}
	if o.NumExchanges != nil {
		toSerialize["num_exchanges"] = o.NumExchanges
	}
	if o.NumPairs != nil {
		toSerialize["num_pairs"] = o.NumPairs
	}
	if o.NumPairsUnmapped != nil {
		toSerialize["num_pairs_unmapped"] = o.NumPairsUnmapped
	}
	if o.FirstCandle != nil {
		toSerialize["first_candle"] = o.FirstCandle
	}
	if o.FirstTrade != nil {
		toSerialize["first_trade"] = o.FirstTrade
	}
	if o.FirstOrderBook != nil {
		toSerialize["first_order_book"] = o.FirstOrderBook
	}
	if o.FirstPricedAt != nil {
		toSerialize["first_priced_at"] = o.FirstPricedAt
	}
	if o.Rank != nil {
		toSerialize["rank"] = o.Rank
	}
	if o.RankDelta != nil {
		toSerialize["rank_delta"] = o.RankDelta
	}
	if o.High != nil {
		toSerialize["high"] = o.High
	}
	if o.HighTimestamp != nil {
		toSerialize["high_timestamp"] = o.HighTimestamp
	}
	if o.Interval != nil {
		toSerialize["interval"] = o.Interval
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200 struct {
	value *InlineResponse200
	isSet bool
}

func (v NullableInlineResponse200) Get() *InlineResponse200 {
	return v.value
}

func (v *NullableInlineResponse200) Set(val *InlineResponse200) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200(val *InlineResponse200) *NullableInlineResponse200 {
	return &NullableInlineResponse200{value: val, isSet: true}
}

func (v NullableInlineResponse200) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
