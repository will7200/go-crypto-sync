# CurrenciesTickerInterval

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PriceChange** | Pointer to **string** | Amount of price change for the given interval in USD | [optional] 
**PriceChangePct** | Pointer to **string** | Percent change of price for the given interval | [optional] 
**Volume** | Pointer to **string** | Total volume for the given interval in USD | [optional] 
**VolumeChange** | Pointer to **string** | Amount of volume change for the given interval in USD | [optional] 
**VolumeChangePct** | Pointer to **string** | Percent change of volume for the given interval | [optional] 
**MarketCapChange** | Pointer to **string** | Amount of market cap change for the given interval in USD | [optional] 
**MarketCapChangePct** | Pointer to **string** | Percent change of market cap for the given interval | [optional] 
**TransparentMarketCapChange** | Pointer to **string** | Amount of transparent market cap change for the given interval in USD | [optional] 
**TransparentMarketCapChangePct** | Pointer to **string** | Percent change of transparent market cap for the given interval | [optional] 
**VolumeTransparency** | Pointer to **[]map[string]interface{}** | An array of &#x60;volume&#x60;, &#x60;volume_change&#x60; and &#x60;volume_change_pct&#x60; by exchange grade | [optional] 
**VolumeTransparencyGrade** | Pointer to **string** | The quartile grade assigned to this currency | [optional] 

## Methods

### NewCurrenciesTickerInterval

`func NewCurrenciesTickerInterval() *CurrenciesTickerInterval`

NewCurrenciesTickerInterval instantiates a new CurrenciesTickerInterval object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrenciesTickerIntervalWithDefaults

`func NewCurrenciesTickerIntervalWithDefaults() *CurrenciesTickerInterval`

NewCurrenciesTickerIntervalWithDefaults instantiates a new CurrenciesTickerInterval object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriceChange

`func (o *CurrenciesTickerInterval) GetPriceChange() string`

GetPriceChange returns the PriceChange field if non-nil, zero value otherwise.

### GetPriceChangeOk

`func (o *CurrenciesTickerInterval) GetPriceChangeOk() (*string, bool)`

GetPriceChangeOk returns a tuple with the PriceChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceChange

`func (o *CurrenciesTickerInterval) SetPriceChange(v string)`

SetPriceChange sets PriceChange field to given value.

### HasPriceChange

`func (o *CurrenciesTickerInterval) HasPriceChange() bool`

HasPriceChange returns a boolean if a field has been set.

### GetPriceChangePct

`func (o *CurrenciesTickerInterval) GetPriceChangePct() string`

GetPriceChangePct returns the PriceChangePct field if non-nil, zero value otherwise.

### GetPriceChangePctOk

`func (o *CurrenciesTickerInterval) GetPriceChangePctOk() (*string, bool)`

GetPriceChangePctOk returns a tuple with the PriceChangePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceChangePct

`func (o *CurrenciesTickerInterval) SetPriceChangePct(v string)`

SetPriceChangePct sets PriceChangePct field to given value.

### HasPriceChangePct

`func (o *CurrenciesTickerInterval) HasPriceChangePct() bool`

HasPriceChangePct returns a boolean if a field has been set.

### GetVolume

`func (o *CurrenciesTickerInterval) GetVolume() string`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *CurrenciesTickerInterval) GetVolumeOk() (*string, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *CurrenciesTickerInterval) SetVolume(v string)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *CurrenciesTickerInterval) HasVolume() bool`

HasVolume returns a boolean if a field has been set.

### GetVolumeChange

`func (o *CurrenciesTickerInterval) GetVolumeChange() string`

GetVolumeChange returns the VolumeChange field if non-nil, zero value otherwise.

### GetVolumeChangeOk

`func (o *CurrenciesTickerInterval) GetVolumeChangeOk() (*string, bool)`

GetVolumeChangeOk returns a tuple with the VolumeChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeChange

`func (o *CurrenciesTickerInterval) SetVolumeChange(v string)`

SetVolumeChange sets VolumeChange field to given value.

### HasVolumeChange

`func (o *CurrenciesTickerInterval) HasVolumeChange() bool`

HasVolumeChange returns a boolean if a field has been set.

### GetVolumeChangePct

`func (o *CurrenciesTickerInterval) GetVolumeChangePct() string`

GetVolumeChangePct returns the VolumeChangePct field if non-nil, zero value otherwise.

### GetVolumeChangePctOk

`func (o *CurrenciesTickerInterval) GetVolumeChangePctOk() (*string, bool)`

GetVolumeChangePctOk returns a tuple with the VolumeChangePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeChangePct

`func (o *CurrenciesTickerInterval) SetVolumeChangePct(v string)`

SetVolumeChangePct sets VolumeChangePct field to given value.

### HasVolumeChangePct

`func (o *CurrenciesTickerInterval) HasVolumeChangePct() bool`

HasVolumeChangePct returns a boolean if a field has been set.

### GetMarketCapChange

`func (o *CurrenciesTickerInterval) GetMarketCapChange() string`

GetMarketCapChange returns the MarketCapChange field if non-nil, zero value otherwise.

### GetMarketCapChangeOk

`func (o *CurrenciesTickerInterval) GetMarketCapChangeOk() (*string, bool)`

GetMarketCapChangeOk returns a tuple with the MarketCapChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketCapChange

`func (o *CurrenciesTickerInterval) SetMarketCapChange(v string)`

SetMarketCapChange sets MarketCapChange field to given value.

### HasMarketCapChange

`func (o *CurrenciesTickerInterval) HasMarketCapChange() bool`

HasMarketCapChange returns a boolean if a field has been set.

### GetMarketCapChangePct

`func (o *CurrenciesTickerInterval) GetMarketCapChangePct() string`

GetMarketCapChangePct returns the MarketCapChangePct field if non-nil, zero value otherwise.

### GetMarketCapChangePctOk

`func (o *CurrenciesTickerInterval) GetMarketCapChangePctOk() (*string, bool)`

GetMarketCapChangePctOk returns a tuple with the MarketCapChangePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketCapChangePct

`func (o *CurrenciesTickerInterval) SetMarketCapChangePct(v string)`

SetMarketCapChangePct sets MarketCapChangePct field to given value.

### HasMarketCapChangePct

`func (o *CurrenciesTickerInterval) HasMarketCapChangePct() bool`

HasMarketCapChangePct returns a boolean if a field has been set.

### GetTransparentMarketCapChange

`func (o *CurrenciesTickerInterval) GetTransparentMarketCapChange() string`

GetTransparentMarketCapChange returns the TransparentMarketCapChange field if non-nil, zero value otherwise.

### GetTransparentMarketCapChangeOk

`func (o *CurrenciesTickerInterval) GetTransparentMarketCapChangeOk() (*string, bool)`

GetTransparentMarketCapChangeOk returns a tuple with the TransparentMarketCapChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransparentMarketCapChange

`func (o *CurrenciesTickerInterval) SetTransparentMarketCapChange(v string)`

SetTransparentMarketCapChange sets TransparentMarketCapChange field to given value.

### HasTransparentMarketCapChange

`func (o *CurrenciesTickerInterval) HasTransparentMarketCapChange() bool`

HasTransparentMarketCapChange returns a boolean if a field has been set.

### GetTransparentMarketCapChangePct

`func (o *CurrenciesTickerInterval) GetTransparentMarketCapChangePct() string`

GetTransparentMarketCapChangePct returns the TransparentMarketCapChangePct field if non-nil, zero value otherwise.

### GetTransparentMarketCapChangePctOk

`func (o *CurrenciesTickerInterval) GetTransparentMarketCapChangePctOk() (*string, bool)`

GetTransparentMarketCapChangePctOk returns a tuple with the TransparentMarketCapChangePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransparentMarketCapChangePct

`func (o *CurrenciesTickerInterval) SetTransparentMarketCapChangePct(v string)`

SetTransparentMarketCapChangePct sets TransparentMarketCapChangePct field to given value.

### HasTransparentMarketCapChangePct

`func (o *CurrenciesTickerInterval) HasTransparentMarketCapChangePct() bool`

HasTransparentMarketCapChangePct returns a boolean if a field has been set.

### GetVolumeTransparency

`func (o *CurrenciesTickerInterval) GetVolumeTransparency() []map[string]interface{}`

GetVolumeTransparency returns the VolumeTransparency field if non-nil, zero value otherwise.

### GetVolumeTransparencyOk

`func (o *CurrenciesTickerInterval) GetVolumeTransparencyOk() (*[]map[string]interface{}, bool)`

GetVolumeTransparencyOk returns a tuple with the VolumeTransparency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeTransparency

`func (o *CurrenciesTickerInterval) SetVolumeTransparency(v []map[string]interface{})`

SetVolumeTransparency sets VolumeTransparency field to given value.

### HasVolumeTransparency

`func (o *CurrenciesTickerInterval) HasVolumeTransparency() bool`

HasVolumeTransparency returns a boolean if a field has been set.

### GetVolumeTransparencyGrade

`func (o *CurrenciesTickerInterval) GetVolumeTransparencyGrade() string`

GetVolumeTransparencyGrade returns the VolumeTransparencyGrade field if non-nil, zero value otherwise.

### GetVolumeTransparencyGradeOk

`func (o *CurrenciesTickerInterval) GetVolumeTransparencyGradeOk() (*string, bool)`

GetVolumeTransparencyGradeOk returns a tuple with the VolumeTransparencyGrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeTransparencyGrade

`func (o *CurrenciesTickerInterval) SetVolumeTransparencyGrade(v string)`

SetVolumeTransparencyGrade sets VolumeTransparencyGrade field to given value.

### HasVolumeTransparencyGrade

`func (o *CurrenciesTickerInterval) HasVolumeTransparencyGrade() bool`

HasVolumeTransparencyGrade returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


