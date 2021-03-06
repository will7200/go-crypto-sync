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

// InlineResponse2001 struct for InlineResponse2001
type InlineResponse2001 struct {
	// Nomics Currency ID
	Id *string `json:"id,omitempty"`
	// Display ticker symbol, not unique
	OriginalSymbol *string `json:"original_symbol,omitempty"`
	// Currency Name
	Name *string `json:"name,omitempty"`
	// Currency description
	Description *string `json:"description,omitempty"`
	// Currency website URL
	WebsiteUrl *string `json:"website_url,omitempty"`
	// Currency logo URL
	LogoUrl *string `json:"logo_url,omitempty"`
	// Currency blog URL
	BlogUrl *string `json:"blog_url,omitempty"`
	// Currency Discord URL
	DiscordUrl *string `json:"discord_url,omitempty"`
	// Currency Facebook URL
	FacebookUrl *string `json:"facebook_url,omitempty"`
	// Currency GitHub URL
	GithubUrl *string `json:"github_url,omitempty"`
	// Currency Medium URL
	MediumUrl *string `json:"medium_url,omitempty"`
	// Currency Reddit URL
	RedditUrl *string `json:"reddit_url,omitempty"`
	// Currency Telegram URL
	TelegramUrl *string `json:"telegram_url,omitempty"`
	// Currency Twitter URL
	TwitterUrl *string `json:"twitter_url,omitempty"`
	// Currency white paper URL
	WhitepaperUrl *string `json:"whitepaper_url,omitempty"`
	// Currency YouTube URL
	YoutubeUrl *string `json:"youtube_url,omitempty"`
	// Currency LinkedIn URL
	LinkedinUrl *string `json:"linkedin_url,omitempty"`
	// Currency BitcoinTalk URL
	BitcointalkUrl *string `json:"bitcointalk_url,omitempty"`
	// Currency Block Explorer URL
	BlockExplorerUrl *string `json:"block_explorer_url,omitempty"`
	// Nomics Currency ID of the currency that replaced the given currency. This can happen on hard forks or blockchain swaps as an example.
	ReplacedBy *string `json:"replaced_by,omitempty"`
	// Coin ID for CryptoControl.io
	CryptocontrolCoinId *string `json:"cryptocontrol_coin_id,omitempty"`
	// Parent BlockChain, if any
	PlatformCurrencyId *string `json:"platform_currency_id,omitempty"`
	// Platform contract address, if any
	PlatformContractAddress *string `json:"platform_contract_address,omitempty"`
	// Whether or not Nomics uses this currency to price other currencies
	UsedForPricing *bool `json:"used_for_pricing,omitempty"`
}

// NewInlineResponse2001 instantiates a new InlineResponse2001 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2001() *InlineResponse2001 {
	this := InlineResponse2001{}
	return &this
}

// NewInlineResponse2001WithDefaults instantiates a new InlineResponse2001 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2001WithDefaults() *InlineResponse2001 {
	this := InlineResponse2001{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse2001) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse2001) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse2001) SetId(v string) {
	o.Id = &v
}

// GetOriginalSymbol returns the OriginalSymbol field value if set, zero value otherwise.
func (o *InlineResponse2001) GetOriginalSymbol() string {
	if o == nil || o.OriginalSymbol == nil {
		var ret string
		return ret
	}
	return *o.OriginalSymbol
}

// GetOriginalSymbolOk returns a tuple with the OriginalSymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetOriginalSymbolOk() (*string, bool) {
	if o == nil || o.OriginalSymbol == nil {
		return nil, false
	}
	return o.OriginalSymbol, true
}

// HasOriginalSymbol returns a boolean if a field has been set.
func (o *InlineResponse2001) HasOriginalSymbol() bool {
	if o != nil && o.OriginalSymbol != nil {
		return true
	}

	return false
}

// SetOriginalSymbol gets a reference to the given string and assigns it to the OriginalSymbol field.
func (o *InlineResponse2001) SetOriginalSymbol(v string) {
	o.OriginalSymbol = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse2001) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse2001) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse2001) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InlineResponse2001) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InlineResponse2001) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InlineResponse2001) SetDescription(v string) {
	o.Description = &v
}

// GetWebsiteUrl returns the WebsiteUrl field value if set, zero value otherwise.
func (o *InlineResponse2001) GetWebsiteUrl() string {
	if o == nil || o.WebsiteUrl == nil {
		var ret string
		return ret
	}
	return *o.WebsiteUrl
}

// GetWebsiteUrlOk returns a tuple with the WebsiteUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetWebsiteUrlOk() (*string, bool) {
	if o == nil || o.WebsiteUrl == nil {
		return nil, false
	}
	return o.WebsiteUrl, true
}

// HasWebsiteUrl returns a boolean if a field has been set.
func (o *InlineResponse2001) HasWebsiteUrl() bool {
	if o != nil && o.WebsiteUrl != nil {
		return true
	}

	return false
}

// SetWebsiteUrl gets a reference to the given string and assigns it to the WebsiteUrl field.
func (o *InlineResponse2001) SetWebsiteUrl(v string) {
	o.WebsiteUrl = &v
}

// GetLogoUrl returns the LogoUrl field value if set, zero value otherwise.
func (o *InlineResponse2001) GetLogoUrl() string {
	if o == nil || o.LogoUrl == nil {
		var ret string
		return ret
	}
	return *o.LogoUrl
}

// GetLogoUrlOk returns a tuple with the LogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetLogoUrlOk() (*string, bool) {
	if o == nil || o.LogoUrl == nil {
		return nil, false
	}
	return o.LogoUrl, true
}

// HasLogoUrl returns a boolean if a field has been set.
func (o *InlineResponse2001) HasLogoUrl() bool {
	if o != nil && o.LogoUrl != nil {
		return true
	}

	return false
}

// SetLogoUrl gets a reference to the given string and assigns it to the LogoUrl field.
func (o *InlineResponse2001) SetLogoUrl(v string) {
	o.LogoUrl = &v
}

// GetBlogUrl returns the BlogUrl field value if set, zero value otherwise.
func (o *InlineResponse2001) GetBlogUrl() string {
	if o == nil || o.BlogUrl == nil {
		var ret string
		return ret
	}
	return *o.BlogUrl
}

// GetBlogUrlOk returns a tuple with the BlogUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetBlogUrlOk() (*string, bool) {
	if o == nil || o.BlogUrl == nil {
		return nil, false
	}
	return o.BlogUrl, true
}

// HasBlogUrl returns a boolean if a field has been set.
func (o *InlineResponse2001) HasBlogUrl() bool {
	if o != nil && o.BlogUrl != nil {
		return true
	}

	return false
}

// SetBlogUrl gets a reference to the given string and assigns it to the BlogUrl field.
func (o *InlineResponse2001) SetBlogUrl(v string) {
	o.BlogUrl = &v
}

// GetDiscordUrl returns the DiscordUrl field value if set, zero value otherwise.
func (o *InlineResponse2001) GetDiscordUrl() string {
	if o == nil || o.DiscordUrl == nil {
		var ret string
		return ret
	}
	return *o.DiscordUrl
}

// GetDiscordUrlOk returns a tuple with the DiscordUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetDiscordUrlOk() (*string, bool) {
	if o == nil || o.DiscordUrl == nil {
		return nil, false
	}
	return o.DiscordUrl, true
}

// HasDiscordUrl returns a boolean if a field has been set.
func (o *InlineResponse2001) HasDiscordUrl() bool {
	if o != nil && o.DiscordUrl != nil {
		return true
	}

	return false
}

// SetDiscordUrl gets a reference to the given string and assigns it to the DiscordUrl field.
func (o *InlineResponse2001) SetDiscordUrl(v string) {
	o.DiscordUrl = &v
}

// GetFacebookUrl returns the FacebookUrl field value if set, zero value otherwise.
func (o *InlineResponse2001) GetFacebookUrl() string {
	if o == nil || o.FacebookUrl == nil {
		var ret string
		return ret
	}
	return *o.FacebookUrl
}

// GetFacebookUrlOk returns a tuple with the FacebookUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetFacebookUrlOk() (*string, bool) {
	if o == nil || o.FacebookUrl == nil {
		return nil, false
	}
	return o.FacebookUrl, true
}

// HasFacebookUrl returns a boolean if a field has been set.
func (o *InlineResponse2001) HasFacebookUrl() bool {
	if o != nil && o.FacebookUrl != nil {
		return true
	}

	return false
}

// SetFacebookUrl gets a reference to the given string and assigns it to the FacebookUrl field.
func (o *InlineResponse2001) SetFacebookUrl(v string) {
	o.FacebookUrl = &v
}

// GetGithubUrl returns the GithubUrl field value if set, zero value otherwise.
func (o *InlineResponse2001) GetGithubUrl() string {
	if o == nil || o.GithubUrl == nil {
		var ret string
		return ret
	}
	return *o.GithubUrl
}

// GetGithubUrlOk returns a tuple with the GithubUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetGithubUrlOk() (*string, bool) {
	if o == nil || o.GithubUrl == nil {
		return nil, false
	}
	return o.GithubUrl, true
}

// HasGithubUrl returns a boolean if a field has been set.
func (o *InlineResponse2001) HasGithubUrl() bool {
	if o != nil && o.GithubUrl != nil {
		return true
	}

	return false
}

// SetGithubUrl gets a reference to the given string and assigns it to the GithubUrl field.
func (o *InlineResponse2001) SetGithubUrl(v string) {
	o.GithubUrl = &v
}

// GetMediumUrl returns the MediumUrl field value if set, zero value otherwise.
func (o *InlineResponse2001) GetMediumUrl() string {
	if o == nil || o.MediumUrl == nil {
		var ret string
		return ret
	}
	return *o.MediumUrl
}

// GetMediumUrlOk returns a tuple with the MediumUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetMediumUrlOk() (*string, bool) {
	if o == nil || o.MediumUrl == nil {
		return nil, false
	}
	return o.MediumUrl, true
}

// HasMediumUrl returns a boolean if a field has been set.
func (o *InlineResponse2001) HasMediumUrl() bool {
	if o != nil && o.MediumUrl != nil {
		return true
	}

	return false
}

// SetMediumUrl gets a reference to the given string and assigns it to the MediumUrl field.
func (o *InlineResponse2001) SetMediumUrl(v string) {
	o.MediumUrl = &v
}

// GetRedditUrl returns the RedditUrl field value if set, zero value otherwise.
func (o *InlineResponse2001) GetRedditUrl() string {
	if o == nil || o.RedditUrl == nil {
		var ret string
		return ret
	}
	return *o.RedditUrl
}

// GetRedditUrlOk returns a tuple with the RedditUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetRedditUrlOk() (*string, bool) {
	if o == nil || o.RedditUrl == nil {
		return nil, false
	}
	return o.RedditUrl, true
}

// HasRedditUrl returns a boolean if a field has been set.
func (o *InlineResponse2001) HasRedditUrl() bool {
	if o != nil && o.RedditUrl != nil {
		return true
	}

	return false
}

// SetRedditUrl gets a reference to the given string and assigns it to the RedditUrl field.
func (o *InlineResponse2001) SetRedditUrl(v string) {
	o.RedditUrl = &v
}

// GetTelegramUrl returns the TelegramUrl field value if set, zero value otherwise.
func (o *InlineResponse2001) GetTelegramUrl() string {
	if o == nil || o.TelegramUrl == nil {
		var ret string
		return ret
	}
	return *o.TelegramUrl
}

// GetTelegramUrlOk returns a tuple with the TelegramUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetTelegramUrlOk() (*string, bool) {
	if o == nil || o.TelegramUrl == nil {
		return nil, false
	}
	return o.TelegramUrl, true
}

// HasTelegramUrl returns a boolean if a field has been set.
func (o *InlineResponse2001) HasTelegramUrl() bool {
	if o != nil && o.TelegramUrl != nil {
		return true
	}

	return false
}

// SetTelegramUrl gets a reference to the given string and assigns it to the TelegramUrl field.
func (o *InlineResponse2001) SetTelegramUrl(v string) {
	o.TelegramUrl = &v
}

// GetTwitterUrl returns the TwitterUrl field value if set, zero value otherwise.
func (o *InlineResponse2001) GetTwitterUrl() string {
	if o == nil || o.TwitterUrl == nil {
		var ret string
		return ret
	}
	return *o.TwitterUrl
}

// GetTwitterUrlOk returns a tuple with the TwitterUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetTwitterUrlOk() (*string, bool) {
	if o == nil || o.TwitterUrl == nil {
		return nil, false
	}
	return o.TwitterUrl, true
}

// HasTwitterUrl returns a boolean if a field has been set.
func (o *InlineResponse2001) HasTwitterUrl() bool {
	if o != nil && o.TwitterUrl != nil {
		return true
	}

	return false
}

// SetTwitterUrl gets a reference to the given string and assigns it to the TwitterUrl field.
func (o *InlineResponse2001) SetTwitterUrl(v string) {
	o.TwitterUrl = &v
}

// GetWhitepaperUrl returns the WhitepaperUrl field value if set, zero value otherwise.
func (o *InlineResponse2001) GetWhitepaperUrl() string {
	if o == nil || o.WhitepaperUrl == nil {
		var ret string
		return ret
	}
	return *o.WhitepaperUrl
}

// GetWhitepaperUrlOk returns a tuple with the WhitepaperUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetWhitepaperUrlOk() (*string, bool) {
	if o == nil || o.WhitepaperUrl == nil {
		return nil, false
	}
	return o.WhitepaperUrl, true
}

// HasWhitepaperUrl returns a boolean if a field has been set.
func (o *InlineResponse2001) HasWhitepaperUrl() bool {
	if o != nil && o.WhitepaperUrl != nil {
		return true
	}

	return false
}

// SetWhitepaperUrl gets a reference to the given string and assigns it to the WhitepaperUrl field.
func (o *InlineResponse2001) SetWhitepaperUrl(v string) {
	o.WhitepaperUrl = &v
}

// GetYoutubeUrl returns the YoutubeUrl field value if set, zero value otherwise.
func (o *InlineResponse2001) GetYoutubeUrl() string {
	if o == nil || o.YoutubeUrl == nil {
		var ret string
		return ret
	}
	return *o.YoutubeUrl
}

// GetYoutubeUrlOk returns a tuple with the YoutubeUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetYoutubeUrlOk() (*string, bool) {
	if o == nil || o.YoutubeUrl == nil {
		return nil, false
	}
	return o.YoutubeUrl, true
}

// HasYoutubeUrl returns a boolean if a field has been set.
func (o *InlineResponse2001) HasYoutubeUrl() bool {
	if o != nil && o.YoutubeUrl != nil {
		return true
	}

	return false
}

// SetYoutubeUrl gets a reference to the given string and assigns it to the YoutubeUrl field.
func (o *InlineResponse2001) SetYoutubeUrl(v string) {
	o.YoutubeUrl = &v
}

// GetLinkedinUrl returns the LinkedinUrl field value if set, zero value otherwise.
func (o *InlineResponse2001) GetLinkedinUrl() string {
	if o == nil || o.LinkedinUrl == nil {
		var ret string
		return ret
	}
	return *o.LinkedinUrl
}

// GetLinkedinUrlOk returns a tuple with the LinkedinUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetLinkedinUrlOk() (*string, bool) {
	if o == nil || o.LinkedinUrl == nil {
		return nil, false
	}
	return o.LinkedinUrl, true
}

// HasLinkedinUrl returns a boolean if a field has been set.
func (o *InlineResponse2001) HasLinkedinUrl() bool {
	if o != nil && o.LinkedinUrl != nil {
		return true
	}

	return false
}

// SetLinkedinUrl gets a reference to the given string and assigns it to the LinkedinUrl field.
func (o *InlineResponse2001) SetLinkedinUrl(v string) {
	o.LinkedinUrl = &v
}

// GetBitcointalkUrl returns the BitcointalkUrl field value if set, zero value otherwise.
func (o *InlineResponse2001) GetBitcointalkUrl() string {
	if o == nil || o.BitcointalkUrl == nil {
		var ret string
		return ret
	}
	return *o.BitcointalkUrl
}

// GetBitcointalkUrlOk returns a tuple with the BitcointalkUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetBitcointalkUrlOk() (*string, bool) {
	if o == nil || o.BitcointalkUrl == nil {
		return nil, false
	}
	return o.BitcointalkUrl, true
}

// HasBitcointalkUrl returns a boolean if a field has been set.
func (o *InlineResponse2001) HasBitcointalkUrl() bool {
	if o != nil && o.BitcointalkUrl != nil {
		return true
	}

	return false
}

// SetBitcointalkUrl gets a reference to the given string and assigns it to the BitcointalkUrl field.
func (o *InlineResponse2001) SetBitcointalkUrl(v string) {
	o.BitcointalkUrl = &v
}

// GetBlockExplorerUrl returns the BlockExplorerUrl field value if set, zero value otherwise.
func (o *InlineResponse2001) GetBlockExplorerUrl() string {
	if o == nil || o.BlockExplorerUrl == nil {
		var ret string
		return ret
	}
	return *o.BlockExplorerUrl
}

// GetBlockExplorerUrlOk returns a tuple with the BlockExplorerUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetBlockExplorerUrlOk() (*string, bool) {
	if o == nil || o.BlockExplorerUrl == nil {
		return nil, false
	}
	return o.BlockExplorerUrl, true
}

// HasBlockExplorerUrl returns a boolean if a field has been set.
func (o *InlineResponse2001) HasBlockExplorerUrl() bool {
	if o != nil && o.BlockExplorerUrl != nil {
		return true
	}

	return false
}

// SetBlockExplorerUrl gets a reference to the given string and assigns it to the BlockExplorerUrl field.
func (o *InlineResponse2001) SetBlockExplorerUrl(v string) {
	o.BlockExplorerUrl = &v
}

// GetReplacedBy returns the ReplacedBy field value if set, zero value otherwise.
func (o *InlineResponse2001) GetReplacedBy() string {
	if o == nil || o.ReplacedBy == nil {
		var ret string
		return ret
	}
	return *o.ReplacedBy
}

// GetReplacedByOk returns a tuple with the ReplacedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetReplacedByOk() (*string, bool) {
	if o == nil || o.ReplacedBy == nil {
		return nil, false
	}
	return o.ReplacedBy, true
}

// HasReplacedBy returns a boolean if a field has been set.
func (o *InlineResponse2001) HasReplacedBy() bool {
	if o != nil && o.ReplacedBy != nil {
		return true
	}

	return false
}

// SetReplacedBy gets a reference to the given string and assigns it to the ReplacedBy field.
func (o *InlineResponse2001) SetReplacedBy(v string) {
	o.ReplacedBy = &v
}

// GetCryptocontrolCoinId returns the CryptocontrolCoinId field value if set, zero value otherwise.
func (o *InlineResponse2001) GetCryptocontrolCoinId() string {
	if o == nil || o.CryptocontrolCoinId == nil {
		var ret string
		return ret
	}
	return *o.CryptocontrolCoinId
}

// GetCryptocontrolCoinIdOk returns a tuple with the CryptocontrolCoinId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetCryptocontrolCoinIdOk() (*string, bool) {
	if o == nil || o.CryptocontrolCoinId == nil {
		return nil, false
	}
	return o.CryptocontrolCoinId, true
}

// HasCryptocontrolCoinId returns a boolean if a field has been set.
func (o *InlineResponse2001) HasCryptocontrolCoinId() bool {
	if o != nil && o.CryptocontrolCoinId != nil {
		return true
	}

	return false
}

// SetCryptocontrolCoinId gets a reference to the given string and assigns it to the CryptocontrolCoinId field.
func (o *InlineResponse2001) SetCryptocontrolCoinId(v string) {
	o.CryptocontrolCoinId = &v
}

// GetPlatformCurrencyId returns the PlatformCurrencyId field value if set, zero value otherwise.
func (o *InlineResponse2001) GetPlatformCurrencyId() string {
	if o == nil || o.PlatformCurrencyId == nil {
		var ret string
		return ret
	}
	return *o.PlatformCurrencyId
}

// GetPlatformCurrencyIdOk returns a tuple with the PlatformCurrencyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetPlatformCurrencyIdOk() (*string, bool) {
	if o == nil || o.PlatformCurrencyId == nil {
		return nil, false
	}
	return o.PlatformCurrencyId, true
}

// HasPlatformCurrencyId returns a boolean if a field has been set.
func (o *InlineResponse2001) HasPlatformCurrencyId() bool {
	if o != nil && o.PlatformCurrencyId != nil {
		return true
	}

	return false
}

// SetPlatformCurrencyId gets a reference to the given string and assigns it to the PlatformCurrencyId field.
func (o *InlineResponse2001) SetPlatformCurrencyId(v string) {
	o.PlatformCurrencyId = &v
}

// GetPlatformContractAddress returns the PlatformContractAddress field value if set, zero value otherwise.
func (o *InlineResponse2001) GetPlatformContractAddress() string {
	if o == nil || o.PlatformContractAddress == nil {
		var ret string
		return ret
	}
	return *o.PlatformContractAddress
}

// GetPlatformContractAddressOk returns a tuple with the PlatformContractAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetPlatformContractAddressOk() (*string, bool) {
	if o == nil || o.PlatformContractAddress == nil {
		return nil, false
	}
	return o.PlatformContractAddress, true
}

// HasPlatformContractAddress returns a boolean if a field has been set.
func (o *InlineResponse2001) HasPlatformContractAddress() bool {
	if o != nil && o.PlatformContractAddress != nil {
		return true
	}

	return false
}

// SetPlatformContractAddress gets a reference to the given string and assigns it to the PlatformContractAddress field.
func (o *InlineResponse2001) SetPlatformContractAddress(v string) {
	o.PlatformContractAddress = &v
}

// GetUsedForPricing returns the UsedForPricing field value if set, zero value otherwise.
func (o *InlineResponse2001) GetUsedForPricing() bool {
	if o == nil || o.UsedForPricing == nil {
		var ret bool
		return ret
	}
	return *o.UsedForPricing
}

// GetUsedForPricingOk returns a tuple with the UsedForPricing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetUsedForPricingOk() (*bool, bool) {
	if o == nil || o.UsedForPricing == nil {
		return nil, false
	}
	return o.UsedForPricing, true
}

// HasUsedForPricing returns a boolean if a field has been set.
func (o *InlineResponse2001) HasUsedForPricing() bool {
	if o != nil && o.UsedForPricing != nil {
		return true
	}

	return false
}

// SetUsedForPricing gets a reference to the given bool and assigns it to the UsedForPricing field.
func (o *InlineResponse2001) SetUsedForPricing(v bool) {
	o.UsedForPricing = &v
}

func (o InlineResponse2001) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.OriginalSymbol != nil {
		toSerialize["original_symbol"] = o.OriginalSymbol
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.WebsiteUrl != nil {
		toSerialize["website_url"] = o.WebsiteUrl
	}
	if o.LogoUrl != nil {
		toSerialize["logo_url"] = o.LogoUrl
	}
	if o.BlogUrl != nil {
		toSerialize["blog_url"] = o.BlogUrl
	}
	if o.DiscordUrl != nil {
		toSerialize["discord_url"] = o.DiscordUrl
	}
	if o.FacebookUrl != nil {
		toSerialize["facebook_url"] = o.FacebookUrl
	}
	if o.GithubUrl != nil {
		toSerialize["github_url"] = o.GithubUrl
	}
	if o.MediumUrl != nil {
		toSerialize["medium_url"] = o.MediumUrl
	}
	if o.RedditUrl != nil {
		toSerialize["reddit_url"] = o.RedditUrl
	}
	if o.TelegramUrl != nil {
		toSerialize["telegram_url"] = o.TelegramUrl
	}
	if o.TwitterUrl != nil {
		toSerialize["twitter_url"] = o.TwitterUrl
	}
	if o.WhitepaperUrl != nil {
		toSerialize["whitepaper_url"] = o.WhitepaperUrl
	}
	if o.YoutubeUrl != nil {
		toSerialize["youtube_url"] = o.YoutubeUrl
	}
	if o.LinkedinUrl != nil {
		toSerialize["linkedin_url"] = o.LinkedinUrl
	}
	if o.BitcointalkUrl != nil {
		toSerialize["bitcointalk_url"] = o.BitcointalkUrl
	}
	if o.BlockExplorerUrl != nil {
		toSerialize["block_explorer_url"] = o.BlockExplorerUrl
	}
	if o.ReplacedBy != nil {
		toSerialize["replaced_by"] = o.ReplacedBy
	}
	if o.CryptocontrolCoinId != nil {
		toSerialize["cryptocontrol_coin_id"] = o.CryptocontrolCoinId
	}
	if o.PlatformCurrencyId != nil {
		toSerialize["platform_currency_id"] = o.PlatformCurrencyId
	}
	if o.PlatformContractAddress != nil {
		toSerialize["platform_contract_address"] = o.PlatformContractAddress
	}
	if o.UsedForPricing != nil {
		toSerialize["used_for_pricing"] = o.UsedForPricing
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2001 struct {
	value *InlineResponse2001
	isSet bool
}

func (v NullableInlineResponse2001) Get() *InlineResponse2001 {
	return v.value
}

func (v *NullableInlineResponse2001) Set(val *InlineResponse2001) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2001) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2001) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2001(val *InlineResponse2001) *NullableInlineResponse2001 {
	return &NullableInlineResponse2001{value: val, isSet: true}
}

func (v NullableInlineResponse2001) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2001) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
