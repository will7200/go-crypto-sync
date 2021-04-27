# InlineResponse2003

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumCurrencies** | Pointer to **string** | The total number of currencies | [optional] 
**NumCurrenciesActive** | Pointer to **string** | The total number of active currencies | [optional] 
**NumCurrenciesInative** | Pointer to **string** | The total number of inactive currencies | [optional] 
**NumCurrenciesDead** | Pointer to **string** | The total number of dead currencies | [optional] 
**NumCurrenciesNew** | Pointer to **string** | The total number of new currencies | [optional] 
**MarketCap** | Pointer to **string** | Current global market cap | [optional] 
**TransparentMarketCap** | Pointer to **string** | Current global transparent market cap | [optional] 
**Var1d** | Pointer to [**GlobalTickerInterval**](GlobalTickerInterval.md) |  | [optional] 
**Var7d** | Pointer to [**GlobalTickerInterval**](GlobalTickerInterval.md) |  | [optional] 
**Var30d** | Pointer to [**GlobalTickerInterval**](GlobalTickerInterval.md) |  | [optional] 
**Var365d** | Pointer to [**GlobalTickerInterval**](GlobalTickerInterval.md) |  | [optional] 
**Ytd** | Pointer to [**GlobalTickerInterval**](GlobalTickerInterval.md) |  | [optional] 

## Methods

### NewInlineResponse2003

`func NewInlineResponse2003() *InlineResponse2003`

NewInlineResponse2003 instantiates a new InlineResponse2003 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2003WithDefaults

`func NewInlineResponse2003WithDefaults() *InlineResponse2003`

NewInlineResponse2003WithDefaults instantiates a new InlineResponse2003 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumCurrencies

`func (o *InlineResponse2003) GetNumCurrencies() string`

GetNumCurrencies returns the NumCurrencies field if non-nil, zero value otherwise.

### GetNumCurrenciesOk

`func (o *InlineResponse2003) GetNumCurrenciesOk() (*string, bool)`

GetNumCurrenciesOk returns a tuple with the NumCurrencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCurrencies

`func (o *InlineResponse2003) SetNumCurrencies(v string)`

SetNumCurrencies sets NumCurrencies field to given value.

### HasNumCurrencies

`func (o *InlineResponse2003) HasNumCurrencies() bool`

HasNumCurrencies returns a boolean if a field has been set.

### GetNumCurrenciesActive

`func (o *InlineResponse2003) GetNumCurrenciesActive() string`

GetNumCurrenciesActive returns the NumCurrenciesActive field if non-nil, zero value otherwise.

### GetNumCurrenciesActiveOk

`func (o *InlineResponse2003) GetNumCurrenciesActiveOk() (*string, bool)`

GetNumCurrenciesActiveOk returns a tuple with the NumCurrenciesActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCurrenciesActive

`func (o *InlineResponse2003) SetNumCurrenciesActive(v string)`

SetNumCurrenciesActive sets NumCurrenciesActive field to given value.

### HasNumCurrenciesActive

`func (o *InlineResponse2003) HasNumCurrenciesActive() bool`

HasNumCurrenciesActive returns a boolean if a field has been set.

### GetNumCurrenciesInative

`func (o *InlineResponse2003) GetNumCurrenciesInative() string`

GetNumCurrenciesInative returns the NumCurrenciesInative field if non-nil, zero value otherwise.

### GetNumCurrenciesInativeOk

`func (o *InlineResponse2003) GetNumCurrenciesInativeOk() (*string, bool)`

GetNumCurrenciesInativeOk returns a tuple with the NumCurrenciesInative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCurrenciesInative

`func (o *InlineResponse2003) SetNumCurrenciesInative(v string)`

SetNumCurrenciesInative sets NumCurrenciesInative field to given value.

### HasNumCurrenciesInative

`func (o *InlineResponse2003) HasNumCurrenciesInative() bool`

HasNumCurrenciesInative returns a boolean if a field has been set.

### GetNumCurrenciesDead

`func (o *InlineResponse2003) GetNumCurrenciesDead() string`

GetNumCurrenciesDead returns the NumCurrenciesDead field if non-nil, zero value otherwise.

### GetNumCurrenciesDeadOk

`func (o *InlineResponse2003) GetNumCurrenciesDeadOk() (*string, bool)`

GetNumCurrenciesDeadOk returns a tuple with the NumCurrenciesDead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCurrenciesDead

`func (o *InlineResponse2003) SetNumCurrenciesDead(v string)`

SetNumCurrenciesDead sets NumCurrenciesDead field to given value.

### HasNumCurrenciesDead

`func (o *InlineResponse2003) HasNumCurrenciesDead() bool`

HasNumCurrenciesDead returns a boolean if a field has been set.

### GetNumCurrenciesNew

`func (o *InlineResponse2003) GetNumCurrenciesNew() string`

GetNumCurrenciesNew returns the NumCurrenciesNew field if non-nil, zero value otherwise.

### GetNumCurrenciesNewOk

`func (o *InlineResponse2003) GetNumCurrenciesNewOk() (*string, bool)`

GetNumCurrenciesNewOk returns a tuple with the NumCurrenciesNew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCurrenciesNew

`func (o *InlineResponse2003) SetNumCurrenciesNew(v string)`

SetNumCurrenciesNew sets NumCurrenciesNew field to given value.

### HasNumCurrenciesNew

`func (o *InlineResponse2003) HasNumCurrenciesNew() bool`

HasNumCurrenciesNew returns a boolean if a field has been set.

### GetMarketCap

`func (o *InlineResponse2003) GetMarketCap() string`

GetMarketCap returns the MarketCap field if non-nil, zero value otherwise.

### GetMarketCapOk

`func (o *InlineResponse2003) GetMarketCapOk() (*string, bool)`

GetMarketCapOk returns a tuple with the MarketCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketCap

`func (o *InlineResponse2003) SetMarketCap(v string)`

SetMarketCap sets MarketCap field to given value.

### HasMarketCap

`func (o *InlineResponse2003) HasMarketCap() bool`

HasMarketCap returns a boolean if a field has been set.

### GetTransparentMarketCap

`func (o *InlineResponse2003) GetTransparentMarketCap() string`

GetTransparentMarketCap returns the TransparentMarketCap field if non-nil, zero value otherwise.

### GetTransparentMarketCapOk

`func (o *InlineResponse2003) GetTransparentMarketCapOk() (*string, bool)`

GetTransparentMarketCapOk returns a tuple with the TransparentMarketCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransparentMarketCap

`func (o *InlineResponse2003) SetTransparentMarketCap(v string)`

SetTransparentMarketCap sets TransparentMarketCap field to given value.

### HasTransparentMarketCap

`func (o *InlineResponse2003) HasTransparentMarketCap() bool`

HasTransparentMarketCap returns a boolean if a field has been set.

### GetVar1d

`func (o *InlineResponse2003) GetVar1d() GlobalTickerInterval`

GetVar1d returns the Var1d field if non-nil, zero value otherwise.

### GetVar1dOk

`func (o *InlineResponse2003) GetVar1dOk() (*GlobalTickerInterval, bool)`

GetVar1dOk returns a tuple with the Var1d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar1d

`func (o *InlineResponse2003) SetVar1d(v GlobalTickerInterval)`

SetVar1d sets Var1d field to given value.

### HasVar1d

`func (o *InlineResponse2003) HasVar1d() bool`

HasVar1d returns a boolean if a field has been set.

### GetVar7d

`func (o *InlineResponse2003) GetVar7d() GlobalTickerInterval`

GetVar7d returns the Var7d field if non-nil, zero value otherwise.

### GetVar7dOk

`func (o *InlineResponse2003) GetVar7dOk() (*GlobalTickerInterval, bool)`

GetVar7dOk returns a tuple with the Var7d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7d

`func (o *InlineResponse2003) SetVar7d(v GlobalTickerInterval)`

SetVar7d sets Var7d field to given value.

### HasVar7d

`func (o *InlineResponse2003) HasVar7d() bool`

HasVar7d returns a boolean if a field has been set.

### GetVar30d

`func (o *InlineResponse2003) GetVar30d() GlobalTickerInterval`

GetVar30d returns the Var30d field if non-nil, zero value otherwise.

### GetVar30dOk

`func (o *InlineResponse2003) GetVar30dOk() (*GlobalTickerInterval, bool)`

GetVar30dOk returns a tuple with the Var30d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar30d

`func (o *InlineResponse2003) SetVar30d(v GlobalTickerInterval)`

SetVar30d sets Var30d field to given value.

### HasVar30d

`func (o *InlineResponse2003) HasVar30d() bool`

HasVar30d returns a boolean if a field has been set.

### GetVar365d

`func (o *InlineResponse2003) GetVar365d() GlobalTickerInterval`

GetVar365d returns the Var365d field if non-nil, zero value otherwise.

### GetVar365dOk

`func (o *InlineResponse2003) GetVar365dOk() (*GlobalTickerInterval, bool)`

GetVar365dOk returns a tuple with the Var365d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar365d

`func (o *InlineResponse2003) SetVar365d(v GlobalTickerInterval)`

SetVar365d sets Var365d field to given value.

### HasVar365d

`func (o *InlineResponse2003) HasVar365d() bool`

HasVar365d returns a boolean if a field has been set.

### GetYtd

`func (o *InlineResponse2003) GetYtd() GlobalTickerInterval`

GetYtd returns the Ytd field if non-nil, zero value otherwise.

### GetYtdOk

`func (o *InlineResponse2003) GetYtdOk() (*GlobalTickerInterval, bool)`

GetYtdOk returns a tuple with the Ytd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYtd

`func (o *InlineResponse2003) SetYtd(v GlobalTickerInterval)`

SetYtd sets Ytd field to given value.

### HasYtd

`func (o *InlineResponse2003) HasYtd() bool`

HasYtd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


