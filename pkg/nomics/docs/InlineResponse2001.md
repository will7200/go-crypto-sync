# InlineResponse2001

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Nomics Currency ID | [optional] 
**OriginalSymbol** | Pointer to **string** | Display ticker symbol, not unique | [optional] 
**Name** | Pointer to **string** | Currency Name | [optional] 
**Description** | Pointer to **string** | Currency description | [optional] 
**WebsiteUrl** | Pointer to **string** | Currency website URL | [optional] 
**LogoUrl** | Pointer to **string** | Currency logo URL | [optional] 
**BlogUrl** | Pointer to **string** | Currency blog URL | [optional] 
**DiscordUrl** | Pointer to **string** | Currency Discord URL | [optional] 
**FacebookUrl** | Pointer to **string** | Currency Facebook URL | [optional] 
**GithubUrl** | Pointer to **string** | Currency GitHub URL | [optional] 
**MediumUrl** | Pointer to **string** | Currency Medium URL | [optional] 
**RedditUrl** | Pointer to **string** | Currency Reddit URL | [optional] 
**TelegramUrl** | Pointer to **string** | Currency Telegram URL | [optional] 
**TwitterUrl** | Pointer to **string** | Currency Twitter URL | [optional] 
**WhitepaperUrl** | Pointer to **string** | Currency white paper URL | [optional] 
**YoutubeUrl** | Pointer to **string** | Currency YouTube URL | [optional] 
**LinkedinUrl** | Pointer to **string** | Currency LinkedIn URL | [optional] 
**BitcointalkUrl** | Pointer to **string** | Currency BitcoinTalk URL | [optional] 
**BlockExplorerUrl** | Pointer to **string** | Currency Block Explorer URL | [optional] 
**ReplacedBy** | Pointer to **string** | Nomics Currency ID of the currency that replaced the given currency. This can happen on hard forks or blockchain swaps as an example. | [optional] 
**CryptocontrolCoinId** | Pointer to **string** | Coin ID for CryptoControl.io | [optional] 
**PlatformCurrencyId** | Pointer to **string** | Parent BlockChain, if any | [optional] 
**PlatformContractAddress** | Pointer to **string** | Platform contract address, if any | [optional] 
**UsedForPricing** | Pointer to **bool** | Whether or not Nomics uses this currency to price other currencies | [optional] 

## Methods

### NewInlineResponse2001

`func NewInlineResponse2001() *InlineResponse2001`

NewInlineResponse2001 instantiates a new InlineResponse2001 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2001WithDefaults

`func NewInlineResponse2001WithDefaults() *InlineResponse2001`

NewInlineResponse2001WithDefaults instantiates a new InlineResponse2001 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse2001) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse2001) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse2001) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse2001) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOriginalSymbol

`func (o *InlineResponse2001) GetOriginalSymbol() string`

GetOriginalSymbol returns the OriginalSymbol field if non-nil, zero value otherwise.

### GetOriginalSymbolOk

`func (o *InlineResponse2001) GetOriginalSymbolOk() (*string, bool)`

GetOriginalSymbolOk returns a tuple with the OriginalSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalSymbol

`func (o *InlineResponse2001) SetOriginalSymbol(v string)`

SetOriginalSymbol sets OriginalSymbol field to given value.

### HasOriginalSymbol

`func (o *InlineResponse2001) HasOriginalSymbol() bool`

HasOriginalSymbol returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse2001) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse2001) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse2001) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse2001) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *InlineResponse2001) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse2001) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse2001) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse2001) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetWebsiteUrl

`func (o *InlineResponse2001) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *InlineResponse2001) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *InlineResponse2001) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.

### HasWebsiteUrl

`func (o *InlineResponse2001) HasWebsiteUrl() bool`

HasWebsiteUrl returns a boolean if a field has been set.

### GetLogoUrl

`func (o *InlineResponse2001) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *InlineResponse2001) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *InlineResponse2001) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *InlineResponse2001) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetBlogUrl

`func (o *InlineResponse2001) GetBlogUrl() string`

GetBlogUrl returns the BlogUrl field if non-nil, zero value otherwise.

### GetBlogUrlOk

`func (o *InlineResponse2001) GetBlogUrlOk() (*string, bool)`

GetBlogUrlOk returns a tuple with the BlogUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlogUrl

`func (o *InlineResponse2001) SetBlogUrl(v string)`

SetBlogUrl sets BlogUrl field to given value.

### HasBlogUrl

`func (o *InlineResponse2001) HasBlogUrl() bool`

HasBlogUrl returns a boolean if a field has been set.

### GetDiscordUrl

`func (o *InlineResponse2001) GetDiscordUrl() string`

GetDiscordUrl returns the DiscordUrl field if non-nil, zero value otherwise.

### GetDiscordUrlOk

`func (o *InlineResponse2001) GetDiscordUrlOk() (*string, bool)`

GetDiscordUrlOk returns a tuple with the DiscordUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordUrl

`func (o *InlineResponse2001) SetDiscordUrl(v string)`

SetDiscordUrl sets DiscordUrl field to given value.

### HasDiscordUrl

`func (o *InlineResponse2001) HasDiscordUrl() bool`

HasDiscordUrl returns a boolean if a field has been set.

### GetFacebookUrl

`func (o *InlineResponse2001) GetFacebookUrl() string`

GetFacebookUrl returns the FacebookUrl field if non-nil, zero value otherwise.

### GetFacebookUrlOk

`func (o *InlineResponse2001) GetFacebookUrlOk() (*string, bool)`

GetFacebookUrlOk returns a tuple with the FacebookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookUrl

`func (o *InlineResponse2001) SetFacebookUrl(v string)`

SetFacebookUrl sets FacebookUrl field to given value.

### HasFacebookUrl

`func (o *InlineResponse2001) HasFacebookUrl() bool`

HasFacebookUrl returns a boolean if a field has been set.

### GetGithubUrl

`func (o *InlineResponse2001) GetGithubUrl() string`

GetGithubUrl returns the GithubUrl field if non-nil, zero value otherwise.

### GetGithubUrlOk

`func (o *InlineResponse2001) GetGithubUrlOk() (*string, bool)`

GetGithubUrlOk returns a tuple with the GithubUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubUrl

`func (o *InlineResponse2001) SetGithubUrl(v string)`

SetGithubUrl sets GithubUrl field to given value.

### HasGithubUrl

`func (o *InlineResponse2001) HasGithubUrl() bool`

HasGithubUrl returns a boolean if a field has been set.

### GetMediumUrl

`func (o *InlineResponse2001) GetMediumUrl() string`

GetMediumUrl returns the MediumUrl field if non-nil, zero value otherwise.

### GetMediumUrlOk

`func (o *InlineResponse2001) GetMediumUrlOk() (*string, bool)`

GetMediumUrlOk returns a tuple with the MediumUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediumUrl

`func (o *InlineResponse2001) SetMediumUrl(v string)`

SetMediumUrl sets MediumUrl field to given value.

### HasMediumUrl

`func (o *InlineResponse2001) HasMediumUrl() bool`

HasMediumUrl returns a boolean if a field has been set.

### GetRedditUrl

`func (o *InlineResponse2001) GetRedditUrl() string`

GetRedditUrl returns the RedditUrl field if non-nil, zero value otherwise.

### GetRedditUrlOk

`func (o *InlineResponse2001) GetRedditUrlOk() (*string, bool)`

GetRedditUrlOk returns a tuple with the RedditUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedditUrl

`func (o *InlineResponse2001) SetRedditUrl(v string)`

SetRedditUrl sets RedditUrl field to given value.

### HasRedditUrl

`func (o *InlineResponse2001) HasRedditUrl() bool`

HasRedditUrl returns a boolean if a field has been set.

### GetTelegramUrl

`func (o *InlineResponse2001) GetTelegramUrl() string`

GetTelegramUrl returns the TelegramUrl field if non-nil, zero value otherwise.

### GetTelegramUrlOk

`func (o *InlineResponse2001) GetTelegramUrlOk() (*string, bool)`

GetTelegramUrlOk returns a tuple with the TelegramUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelegramUrl

`func (o *InlineResponse2001) SetTelegramUrl(v string)`

SetTelegramUrl sets TelegramUrl field to given value.

### HasTelegramUrl

`func (o *InlineResponse2001) HasTelegramUrl() bool`

HasTelegramUrl returns a boolean if a field has been set.

### GetTwitterUrl

`func (o *InlineResponse2001) GetTwitterUrl() string`

GetTwitterUrl returns the TwitterUrl field if non-nil, zero value otherwise.

### GetTwitterUrlOk

`func (o *InlineResponse2001) GetTwitterUrlOk() (*string, bool)`

GetTwitterUrlOk returns a tuple with the TwitterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterUrl

`func (o *InlineResponse2001) SetTwitterUrl(v string)`

SetTwitterUrl sets TwitterUrl field to given value.

### HasTwitterUrl

`func (o *InlineResponse2001) HasTwitterUrl() bool`

HasTwitterUrl returns a boolean if a field has been set.

### GetWhitepaperUrl

`func (o *InlineResponse2001) GetWhitepaperUrl() string`

GetWhitepaperUrl returns the WhitepaperUrl field if non-nil, zero value otherwise.

### GetWhitepaperUrlOk

`func (o *InlineResponse2001) GetWhitepaperUrlOk() (*string, bool)`

GetWhitepaperUrlOk returns a tuple with the WhitepaperUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitepaperUrl

`func (o *InlineResponse2001) SetWhitepaperUrl(v string)`

SetWhitepaperUrl sets WhitepaperUrl field to given value.

### HasWhitepaperUrl

`func (o *InlineResponse2001) HasWhitepaperUrl() bool`

HasWhitepaperUrl returns a boolean if a field has been set.

### GetYoutubeUrl

`func (o *InlineResponse2001) GetYoutubeUrl() string`

GetYoutubeUrl returns the YoutubeUrl field if non-nil, zero value otherwise.

### GetYoutubeUrlOk

`func (o *InlineResponse2001) GetYoutubeUrlOk() (*string, bool)`

GetYoutubeUrlOk returns a tuple with the YoutubeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYoutubeUrl

`func (o *InlineResponse2001) SetYoutubeUrl(v string)`

SetYoutubeUrl sets YoutubeUrl field to given value.

### HasYoutubeUrl

`func (o *InlineResponse2001) HasYoutubeUrl() bool`

HasYoutubeUrl returns a boolean if a field has been set.

### GetLinkedinUrl

`func (o *InlineResponse2001) GetLinkedinUrl() string`

GetLinkedinUrl returns the LinkedinUrl field if non-nil, zero value otherwise.

### GetLinkedinUrlOk

`func (o *InlineResponse2001) GetLinkedinUrlOk() (*string, bool)`

GetLinkedinUrlOk returns a tuple with the LinkedinUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedinUrl

`func (o *InlineResponse2001) SetLinkedinUrl(v string)`

SetLinkedinUrl sets LinkedinUrl field to given value.

### HasLinkedinUrl

`func (o *InlineResponse2001) HasLinkedinUrl() bool`

HasLinkedinUrl returns a boolean if a field has been set.

### GetBitcointalkUrl

`func (o *InlineResponse2001) GetBitcointalkUrl() string`

GetBitcointalkUrl returns the BitcointalkUrl field if non-nil, zero value otherwise.

### GetBitcointalkUrlOk

`func (o *InlineResponse2001) GetBitcointalkUrlOk() (*string, bool)`

GetBitcointalkUrlOk returns a tuple with the BitcointalkUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitcointalkUrl

`func (o *InlineResponse2001) SetBitcointalkUrl(v string)`

SetBitcointalkUrl sets BitcointalkUrl field to given value.

### HasBitcointalkUrl

`func (o *InlineResponse2001) HasBitcointalkUrl() bool`

HasBitcointalkUrl returns a boolean if a field has been set.

### GetBlockExplorerUrl

`func (o *InlineResponse2001) GetBlockExplorerUrl() string`

GetBlockExplorerUrl returns the BlockExplorerUrl field if non-nil, zero value otherwise.

### GetBlockExplorerUrlOk

`func (o *InlineResponse2001) GetBlockExplorerUrlOk() (*string, bool)`

GetBlockExplorerUrlOk returns a tuple with the BlockExplorerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockExplorerUrl

`func (o *InlineResponse2001) SetBlockExplorerUrl(v string)`

SetBlockExplorerUrl sets BlockExplorerUrl field to given value.

### HasBlockExplorerUrl

`func (o *InlineResponse2001) HasBlockExplorerUrl() bool`

HasBlockExplorerUrl returns a boolean if a field has been set.

### GetReplacedBy

`func (o *InlineResponse2001) GetReplacedBy() string`

GetReplacedBy returns the ReplacedBy field if non-nil, zero value otherwise.

### GetReplacedByOk

`func (o *InlineResponse2001) GetReplacedByOk() (*string, bool)`

GetReplacedByOk returns a tuple with the ReplacedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacedBy

`func (o *InlineResponse2001) SetReplacedBy(v string)`

SetReplacedBy sets ReplacedBy field to given value.

### HasReplacedBy

`func (o *InlineResponse2001) HasReplacedBy() bool`

HasReplacedBy returns a boolean if a field has been set.

### GetCryptocontrolCoinId

`func (o *InlineResponse2001) GetCryptocontrolCoinId() string`

GetCryptocontrolCoinId returns the CryptocontrolCoinId field if non-nil, zero value otherwise.

### GetCryptocontrolCoinIdOk

`func (o *InlineResponse2001) GetCryptocontrolCoinIdOk() (*string, bool)`

GetCryptocontrolCoinIdOk returns a tuple with the CryptocontrolCoinId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptocontrolCoinId

`func (o *InlineResponse2001) SetCryptocontrolCoinId(v string)`

SetCryptocontrolCoinId sets CryptocontrolCoinId field to given value.

### HasCryptocontrolCoinId

`func (o *InlineResponse2001) HasCryptocontrolCoinId() bool`

HasCryptocontrolCoinId returns a boolean if a field has been set.

### GetPlatformCurrencyId

`func (o *InlineResponse2001) GetPlatformCurrencyId() string`

GetPlatformCurrencyId returns the PlatformCurrencyId field if non-nil, zero value otherwise.

### GetPlatformCurrencyIdOk

`func (o *InlineResponse2001) GetPlatformCurrencyIdOk() (*string, bool)`

GetPlatformCurrencyIdOk returns a tuple with the PlatformCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformCurrencyId

`func (o *InlineResponse2001) SetPlatformCurrencyId(v string)`

SetPlatformCurrencyId sets PlatformCurrencyId field to given value.

### HasPlatformCurrencyId

`func (o *InlineResponse2001) HasPlatformCurrencyId() bool`

HasPlatformCurrencyId returns a boolean if a field has been set.

### GetPlatformContractAddress

`func (o *InlineResponse2001) GetPlatformContractAddress() string`

GetPlatformContractAddress returns the PlatformContractAddress field if non-nil, zero value otherwise.

### GetPlatformContractAddressOk

`func (o *InlineResponse2001) GetPlatformContractAddressOk() (*string, bool)`

GetPlatformContractAddressOk returns a tuple with the PlatformContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformContractAddress

`func (o *InlineResponse2001) SetPlatformContractAddress(v string)`

SetPlatformContractAddress sets PlatformContractAddress field to given value.

### HasPlatformContractAddress

`func (o *InlineResponse2001) HasPlatformContractAddress() bool`

HasPlatformContractAddress returns a boolean if a field has been set.

### GetUsedForPricing

`func (o *InlineResponse2001) GetUsedForPricing() bool`

GetUsedForPricing returns the UsedForPricing field if non-nil, zero value otherwise.

### GetUsedForPricingOk

`func (o *InlineResponse2001) GetUsedForPricingOk() (*bool, bool)`

GetUsedForPricingOk returns a tuple with the UsedForPricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedForPricing

`func (o *InlineResponse2001) SetUsedForPricing(v bool)`

SetUsedForPricing sets UsedForPricing field to given value.

### HasUsedForPricing

`func (o *InlineResponse2001) HasUsedForPricing() bool`

HasUsedForPricing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


