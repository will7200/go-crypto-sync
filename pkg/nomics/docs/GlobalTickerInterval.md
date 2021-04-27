# GlobalTickerInterval

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MarketCapChange** | Pointer to **string** | Amount of market cap change for the given interval | [optional] 
**MarketCapChangePct** | Pointer to **string** | Percent change of market cap for the given interval | [optional] 
**TransparentMarketCapChange** | Pointer to **string** | Amount of transparent market cap change for the given interval | [optional] 
**TransparentMarketCapChangePct** | Pointer to **string** | Percent change of transparent market cap for the given interval | [optional] 
**Volume** | Pointer to **string** | Total volume (spot + derivative volume) for the given interval | [optional] 
**VolumeChange** | Pointer to **string** | Amount of volume change for the given interval | [optional] 
**VolumeChangePct** | Pointer to **string** | Percent change of volume for the given interval | [optional] 
**SpotVolume** | Pointer to **string** | Total spot market volume for the given interval | [optional] 
**SpotVolumeChange** | Pointer to **string** | Amount of spot volume change for the given interval | [optional] 
**SpotVolumeChangePct** | Pointer to **string** | Percent change of spot volume for the given interval | [optional] 
**DerivativeVolume** | Pointer to **string** | Total derivative market volume for the given interval | [optional] 
**DerivativeVolumeChange** | Pointer to **string** | Amount of derivative volume change for the given interval | [optional] 
**DerivativeVolumeChangePct** | Pointer to **string** | Percent change of derivative volume for the given interval | [optional] 
**TransparentVolume** | Pointer to **string** | Total transparent volume for the given interval | [optional] 
**TransparentVolumeChange** | Pointer to **string** | Amount of transparent volume change for the given interval | [optional] 
**TransparentVolumeChangePct** | Pointer to **string** | Percent change of transparent volume for the given interval | [optional] 
**TransparentSpotVolume** | Pointer to **string** | Total transparent spot market volume for the given interval | [optional] 
**TransparentSpotVolumeChange** | Pointer to **string** | Amount of transparent spot market volume change for the given interval | [optional] 
**TransparentSpotVolumeChangePct** | Pointer to **string** | Percent change of transparent spot market volume for the given interval | [optional] 
**TransparentDerivativeVolume** | Pointer to **string** | Total transparent derivative market volume for the given interval | [optional] 
**TransparentDerivativeVolumeChange** | Pointer to **string** | Amount of transparent derivative market volume change for the given interval | [optional] 
**TransparentDerivativeVolumeChangePct** | Pointer to **string** | Percent change of transparent derivative market volume for the given interval | [optional] 
**VolumeTransparency** | Pointer to **[]map[string]interface{}** | An array of &#x60;volume&#x60;, &#x60;volume_change&#x60; and &#x60;volume_change_pct&#x60; by exchange grade for the given interval | [optional] 

## Methods

### NewGlobalTickerInterval

`func NewGlobalTickerInterval() *GlobalTickerInterval`

NewGlobalTickerInterval instantiates a new GlobalTickerInterval object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalTickerIntervalWithDefaults

`func NewGlobalTickerIntervalWithDefaults() *GlobalTickerInterval`

NewGlobalTickerIntervalWithDefaults instantiates a new GlobalTickerInterval object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarketCapChange

`func (o *GlobalTickerInterval) GetMarketCapChange() string`

GetMarketCapChange returns the MarketCapChange field if non-nil, zero value otherwise.

### GetMarketCapChangeOk

`func (o *GlobalTickerInterval) GetMarketCapChangeOk() (*string, bool)`

GetMarketCapChangeOk returns a tuple with the MarketCapChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketCapChange

`func (o *GlobalTickerInterval) SetMarketCapChange(v string)`

SetMarketCapChange sets MarketCapChange field to given value.

### HasMarketCapChange

`func (o *GlobalTickerInterval) HasMarketCapChange() bool`

HasMarketCapChange returns a boolean if a field has been set.

### GetMarketCapChangePct

`func (o *GlobalTickerInterval) GetMarketCapChangePct() string`

GetMarketCapChangePct returns the MarketCapChangePct field if non-nil, zero value otherwise.

### GetMarketCapChangePctOk

`func (o *GlobalTickerInterval) GetMarketCapChangePctOk() (*string, bool)`

GetMarketCapChangePctOk returns a tuple with the MarketCapChangePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketCapChangePct

`func (o *GlobalTickerInterval) SetMarketCapChangePct(v string)`

SetMarketCapChangePct sets MarketCapChangePct field to given value.

### HasMarketCapChangePct

`func (o *GlobalTickerInterval) HasMarketCapChangePct() bool`

HasMarketCapChangePct returns a boolean if a field has been set.

### GetTransparentMarketCapChange

`func (o *GlobalTickerInterval) GetTransparentMarketCapChange() string`

GetTransparentMarketCapChange returns the TransparentMarketCapChange field if non-nil, zero value otherwise.

### GetTransparentMarketCapChangeOk

`func (o *GlobalTickerInterval) GetTransparentMarketCapChangeOk() (*string, bool)`

GetTransparentMarketCapChangeOk returns a tuple with the TransparentMarketCapChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransparentMarketCapChange

`func (o *GlobalTickerInterval) SetTransparentMarketCapChange(v string)`

SetTransparentMarketCapChange sets TransparentMarketCapChange field to given value.

### HasTransparentMarketCapChange

`func (o *GlobalTickerInterval) HasTransparentMarketCapChange() bool`

HasTransparentMarketCapChange returns a boolean if a field has been set.

### GetTransparentMarketCapChangePct

`func (o *GlobalTickerInterval) GetTransparentMarketCapChangePct() string`

GetTransparentMarketCapChangePct returns the TransparentMarketCapChangePct field if non-nil, zero value otherwise.

### GetTransparentMarketCapChangePctOk

`func (o *GlobalTickerInterval) GetTransparentMarketCapChangePctOk() (*string, bool)`

GetTransparentMarketCapChangePctOk returns a tuple with the TransparentMarketCapChangePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransparentMarketCapChangePct

`func (o *GlobalTickerInterval) SetTransparentMarketCapChangePct(v string)`

SetTransparentMarketCapChangePct sets TransparentMarketCapChangePct field to given value.

### HasTransparentMarketCapChangePct

`func (o *GlobalTickerInterval) HasTransparentMarketCapChangePct() bool`

HasTransparentMarketCapChangePct returns a boolean if a field has been set.

### GetVolume

`func (o *GlobalTickerInterval) GetVolume() string`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *GlobalTickerInterval) GetVolumeOk() (*string, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *GlobalTickerInterval) SetVolume(v string)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *GlobalTickerInterval) HasVolume() bool`

HasVolume returns a boolean if a field has been set.

### GetVolumeChange

`func (o *GlobalTickerInterval) GetVolumeChange() string`

GetVolumeChange returns the VolumeChange field if non-nil, zero value otherwise.

### GetVolumeChangeOk

`func (o *GlobalTickerInterval) GetVolumeChangeOk() (*string, bool)`

GetVolumeChangeOk returns a tuple with the VolumeChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeChange

`func (o *GlobalTickerInterval) SetVolumeChange(v string)`

SetVolumeChange sets VolumeChange field to given value.

### HasVolumeChange

`func (o *GlobalTickerInterval) HasVolumeChange() bool`

HasVolumeChange returns a boolean if a field has been set.

### GetVolumeChangePct

`func (o *GlobalTickerInterval) GetVolumeChangePct() string`

GetVolumeChangePct returns the VolumeChangePct field if non-nil, zero value otherwise.

### GetVolumeChangePctOk

`func (o *GlobalTickerInterval) GetVolumeChangePctOk() (*string, bool)`

GetVolumeChangePctOk returns a tuple with the VolumeChangePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeChangePct

`func (o *GlobalTickerInterval) SetVolumeChangePct(v string)`

SetVolumeChangePct sets VolumeChangePct field to given value.

### HasVolumeChangePct

`func (o *GlobalTickerInterval) HasVolumeChangePct() bool`

HasVolumeChangePct returns a boolean if a field has been set.

### GetSpotVolume

`func (o *GlobalTickerInterval) GetSpotVolume() string`

GetSpotVolume returns the SpotVolume field if non-nil, zero value otherwise.

### GetSpotVolumeOk

`func (o *GlobalTickerInterval) GetSpotVolumeOk() (*string, bool)`

GetSpotVolumeOk returns a tuple with the SpotVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotVolume

`func (o *GlobalTickerInterval) SetSpotVolume(v string)`

SetSpotVolume sets SpotVolume field to given value.

### HasSpotVolume

`func (o *GlobalTickerInterval) HasSpotVolume() bool`

HasSpotVolume returns a boolean if a field has been set.

### GetSpotVolumeChange

`func (o *GlobalTickerInterval) GetSpotVolumeChange() string`

GetSpotVolumeChange returns the SpotVolumeChange field if non-nil, zero value otherwise.

### GetSpotVolumeChangeOk

`func (o *GlobalTickerInterval) GetSpotVolumeChangeOk() (*string, bool)`

GetSpotVolumeChangeOk returns a tuple with the SpotVolumeChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotVolumeChange

`func (o *GlobalTickerInterval) SetSpotVolumeChange(v string)`

SetSpotVolumeChange sets SpotVolumeChange field to given value.

### HasSpotVolumeChange

`func (o *GlobalTickerInterval) HasSpotVolumeChange() bool`

HasSpotVolumeChange returns a boolean if a field has been set.

### GetSpotVolumeChangePct

`func (o *GlobalTickerInterval) GetSpotVolumeChangePct() string`

GetSpotVolumeChangePct returns the SpotVolumeChangePct field if non-nil, zero value otherwise.

### GetSpotVolumeChangePctOk

`func (o *GlobalTickerInterval) GetSpotVolumeChangePctOk() (*string, bool)`

GetSpotVolumeChangePctOk returns a tuple with the SpotVolumeChangePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotVolumeChangePct

`func (o *GlobalTickerInterval) SetSpotVolumeChangePct(v string)`

SetSpotVolumeChangePct sets SpotVolumeChangePct field to given value.

### HasSpotVolumeChangePct

`func (o *GlobalTickerInterval) HasSpotVolumeChangePct() bool`

HasSpotVolumeChangePct returns a boolean if a field has been set.

### GetDerivativeVolume

`func (o *GlobalTickerInterval) GetDerivativeVolume() string`

GetDerivativeVolume returns the DerivativeVolume field if non-nil, zero value otherwise.

### GetDerivativeVolumeOk

`func (o *GlobalTickerInterval) GetDerivativeVolumeOk() (*string, bool)`

GetDerivativeVolumeOk returns a tuple with the DerivativeVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivativeVolume

`func (o *GlobalTickerInterval) SetDerivativeVolume(v string)`

SetDerivativeVolume sets DerivativeVolume field to given value.

### HasDerivativeVolume

`func (o *GlobalTickerInterval) HasDerivativeVolume() bool`

HasDerivativeVolume returns a boolean if a field has been set.

### GetDerivativeVolumeChange

`func (o *GlobalTickerInterval) GetDerivativeVolumeChange() string`

GetDerivativeVolumeChange returns the DerivativeVolumeChange field if non-nil, zero value otherwise.

### GetDerivativeVolumeChangeOk

`func (o *GlobalTickerInterval) GetDerivativeVolumeChangeOk() (*string, bool)`

GetDerivativeVolumeChangeOk returns a tuple with the DerivativeVolumeChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivativeVolumeChange

`func (o *GlobalTickerInterval) SetDerivativeVolumeChange(v string)`

SetDerivativeVolumeChange sets DerivativeVolumeChange field to given value.

### HasDerivativeVolumeChange

`func (o *GlobalTickerInterval) HasDerivativeVolumeChange() bool`

HasDerivativeVolumeChange returns a boolean if a field has been set.

### GetDerivativeVolumeChangePct

`func (o *GlobalTickerInterval) GetDerivativeVolumeChangePct() string`

GetDerivativeVolumeChangePct returns the DerivativeVolumeChangePct field if non-nil, zero value otherwise.

### GetDerivativeVolumeChangePctOk

`func (o *GlobalTickerInterval) GetDerivativeVolumeChangePctOk() (*string, bool)`

GetDerivativeVolumeChangePctOk returns a tuple with the DerivativeVolumeChangePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivativeVolumeChangePct

`func (o *GlobalTickerInterval) SetDerivativeVolumeChangePct(v string)`

SetDerivativeVolumeChangePct sets DerivativeVolumeChangePct field to given value.

### HasDerivativeVolumeChangePct

`func (o *GlobalTickerInterval) HasDerivativeVolumeChangePct() bool`

HasDerivativeVolumeChangePct returns a boolean if a field has been set.

### GetTransparentVolume

`func (o *GlobalTickerInterval) GetTransparentVolume() string`

GetTransparentVolume returns the TransparentVolume field if non-nil, zero value otherwise.

### GetTransparentVolumeOk

`func (o *GlobalTickerInterval) GetTransparentVolumeOk() (*string, bool)`

GetTransparentVolumeOk returns a tuple with the TransparentVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransparentVolume

`func (o *GlobalTickerInterval) SetTransparentVolume(v string)`

SetTransparentVolume sets TransparentVolume field to given value.

### HasTransparentVolume

`func (o *GlobalTickerInterval) HasTransparentVolume() bool`

HasTransparentVolume returns a boolean if a field has been set.

### GetTransparentVolumeChange

`func (o *GlobalTickerInterval) GetTransparentVolumeChange() string`

GetTransparentVolumeChange returns the TransparentVolumeChange field if non-nil, zero value otherwise.

### GetTransparentVolumeChangeOk

`func (o *GlobalTickerInterval) GetTransparentVolumeChangeOk() (*string, bool)`

GetTransparentVolumeChangeOk returns a tuple with the TransparentVolumeChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransparentVolumeChange

`func (o *GlobalTickerInterval) SetTransparentVolumeChange(v string)`

SetTransparentVolumeChange sets TransparentVolumeChange field to given value.

### HasTransparentVolumeChange

`func (o *GlobalTickerInterval) HasTransparentVolumeChange() bool`

HasTransparentVolumeChange returns a boolean if a field has been set.

### GetTransparentVolumeChangePct

`func (o *GlobalTickerInterval) GetTransparentVolumeChangePct() string`

GetTransparentVolumeChangePct returns the TransparentVolumeChangePct field if non-nil, zero value otherwise.

### GetTransparentVolumeChangePctOk

`func (o *GlobalTickerInterval) GetTransparentVolumeChangePctOk() (*string, bool)`

GetTransparentVolumeChangePctOk returns a tuple with the TransparentVolumeChangePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransparentVolumeChangePct

`func (o *GlobalTickerInterval) SetTransparentVolumeChangePct(v string)`

SetTransparentVolumeChangePct sets TransparentVolumeChangePct field to given value.

### HasTransparentVolumeChangePct

`func (o *GlobalTickerInterval) HasTransparentVolumeChangePct() bool`

HasTransparentVolumeChangePct returns a boolean if a field has been set.

### GetTransparentSpotVolume

`func (o *GlobalTickerInterval) GetTransparentSpotVolume() string`

GetTransparentSpotVolume returns the TransparentSpotVolume field if non-nil, zero value otherwise.

### GetTransparentSpotVolumeOk

`func (o *GlobalTickerInterval) GetTransparentSpotVolumeOk() (*string, bool)`

GetTransparentSpotVolumeOk returns a tuple with the TransparentSpotVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransparentSpotVolume

`func (o *GlobalTickerInterval) SetTransparentSpotVolume(v string)`

SetTransparentSpotVolume sets TransparentSpotVolume field to given value.

### HasTransparentSpotVolume

`func (o *GlobalTickerInterval) HasTransparentSpotVolume() bool`

HasTransparentSpotVolume returns a boolean if a field has been set.

### GetTransparentSpotVolumeChange

`func (o *GlobalTickerInterval) GetTransparentSpotVolumeChange() string`

GetTransparentSpotVolumeChange returns the TransparentSpotVolumeChange field if non-nil, zero value otherwise.

### GetTransparentSpotVolumeChangeOk

`func (o *GlobalTickerInterval) GetTransparentSpotVolumeChangeOk() (*string, bool)`

GetTransparentSpotVolumeChangeOk returns a tuple with the TransparentSpotVolumeChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransparentSpotVolumeChange

`func (o *GlobalTickerInterval) SetTransparentSpotVolumeChange(v string)`

SetTransparentSpotVolumeChange sets TransparentSpotVolumeChange field to given value.

### HasTransparentSpotVolumeChange

`func (o *GlobalTickerInterval) HasTransparentSpotVolumeChange() bool`

HasTransparentSpotVolumeChange returns a boolean if a field has been set.

### GetTransparentSpotVolumeChangePct

`func (o *GlobalTickerInterval) GetTransparentSpotVolumeChangePct() string`

GetTransparentSpotVolumeChangePct returns the TransparentSpotVolumeChangePct field if non-nil, zero value otherwise.

### GetTransparentSpotVolumeChangePctOk

`func (o *GlobalTickerInterval) GetTransparentSpotVolumeChangePctOk() (*string, bool)`

GetTransparentSpotVolumeChangePctOk returns a tuple with the TransparentSpotVolumeChangePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransparentSpotVolumeChangePct

`func (o *GlobalTickerInterval) SetTransparentSpotVolumeChangePct(v string)`

SetTransparentSpotVolumeChangePct sets TransparentSpotVolumeChangePct field to given value.

### HasTransparentSpotVolumeChangePct

`func (o *GlobalTickerInterval) HasTransparentSpotVolumeChangePct() bool`

HasTransparentSpotVolumeChangePct returns a boolean if a field has been set.

### GetTransparentDerivativeVolume

`func (o *GlobalTickerInterval) GetTransparentDerivativeVolume() string`

GetTransparentDerivativeVolume returns the TransparentDerivativeVolume field if non-nil, zero value otherwise.

### GetTransparentDerivativeVolumeOk

`func (o *GlobalTickerInterval) GetTransparentDerivativeVolumeOk() (*string, bool)`

GetTransparentDerivativeVolumeOk returns a tuple with the TransparentDerivativeVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransparentDerivativeVolume

`func (o *GlobalTickerInterval) SetTransparentDerivativeVolume(v string)`

SetTransparentDerivativeVolume sets TransparentDerivativeVolume field to given value.

### HasTransparentDerivativeVolume

`func (o *GlobalTickerInterval) HasTransparentDerivativeVolume() bool`

HasTransparentDerivativeVolume returns a boolean if a field has been set.

### GetTransparentDerivativeVolumeChange

`func (o *GlobalTickerInterval) GetTransparentDerivativeVolumeChange() string`

GetTransparentDerivativeVolumeChange returns the TransparentDerivativeVolumeChange field if non-nil, zero value otherwise.

### GetTransparentDerivativeVolumeChangeOk

`func (o *GlobalTickerInterval) GetTransparentDerivativeVolumeChangeOk() (*string, bool)`

GetTransparentDerivativeVolumeChangeOk returns a tuple with the TransparentDerivativeVolumeChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransparentDerivativeVolumeChange

`func (o *GlobalTickerInterval) SetTransparentDerivativeVolumeChange(v string)`

SetTransparentDerivativeVolumeChange sets TransparentDerivativeVolumeChange field to given value.

### HasTransparentDerivativeVolumeChange

`func (o *GlobalTickerInterval) HasTransparentDerivativeVolumeChange() bool`

HasTransparentDerivativeVolumeChange returns a boolean if a field has been set.

### GetTransparentDerivativeVolumeChangePct

`func (o *GlobalTickerInterval) GetTransparentDerivativeVolumeChangePct() string`

GetTransparentDerivativeVolumeChangePct returns the TransparentDerivativeVolumeChangePct field if non-nil, zero value otherwise.

### GetTransparentDerivativeVolumeChangePctOk

`func (o *GlobalTickerInterval) GetTransparentDerivativeVolumeChangePctOk() (*string, bool)`

GetTransparentDerivativeVolumeChangePctOk returns a tuple with the TransparentDerivativeVolumeChangePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransparentDerivativeVolumeChangePct

`func (o *GlobalTickerInterval) SetTransparentDerivativeVolumeChangePct(v string)`

SetTransparentDerivativeVolumeChangePct sets TransparentDerivativeVolumeChangePct field to given value.

### HasTransparentDerivativeVolumeChangePct

`func (o *GlobalTickerInterval) HasTransparentDerivativeVolumeChangePct() bool`

HasTransparentDerivativeVolumeChangePct returns a boolean if a field has been set.

### GetVolumeTransparency

`func (o *GlobalTickerInterval) GetVolumeTransparency() []map[string]interface{}`

GetVolumeTransparency returns the VolumeTransparency field if non-nil, zero value otherwise.

### GetVolumeTransparencyOk

`func (o *GlobalTickerInterval) GetVolumeTransparencyOk() (*[]map[string]interface{}, bool)`

GetVolumeTransparencyOk returns a tuple with the VolumeTransparency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeTransparency

`func (o *GlobalTickerInterval) SetVolumeTransparency(v []map[string]interface{})`

SetVolumeTransparency sets VolumeTransparency field to given value.

### HasVolumeTransparency

`func (o *GlobalTickerInterval) HasVolumeTransparency() bool`

HasVolumeTransparency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


