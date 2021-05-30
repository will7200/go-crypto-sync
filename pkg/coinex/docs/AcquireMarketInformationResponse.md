# AcquireMarketInformationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | market name | [optional] 
**TakerFeeRate** | Pointer to **string** | taker fee | [optional] 
**MakerFeeRate** | Pointer to **string** | maker fee | [optional] 
**MinAmount** | Pointer to **string** | Min. amount of transaction | [optional] 
**TradingName** | Pointer to **string** | tradiing coin type | [optional] 
**TradingDecimal** | Pointer to **int64** | decimal of tradiing coin | [optional] 
**PricingName** | Pointer to **string** | pricing coin type | [optional] 
**PricingDecimal** | Pointer to **int64** | decimal of pricing coin | [optional] 

## Methods

### NewAcquireMarketInformationResponse

`func NewAcquireMarketInformationResponse() *AcquireMarketInformationResponse`

NewAcquireMarketInformationResponse instantiates a new AcquireMarketInformationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcquireMarketInformationResponseWithDefaults

`func NewAcquireMarketInformationResponseWithDefaults() *AcquireMarketInformationResponse`

NewAcquireMarketInformationResponseWithDefaults instantiates a new AcquireMarketInformationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AcquireMarketInformationResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AcquireMarketInformationResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AcquireMarketInformationResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AcquireMarketInformationResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTakerFeeRate

`func (o *AcquireMarketInformationResponse) GetTakerFeeRate() string`

GetTakerFeeRate returns the TakerFeeRate field if non-nil, zero value otherwise.

### GetTakerFeeRateOk

`func (o *AcquireMarketInformationResponse) GetTakerFeeRateOk() (*string, bool)`

GetTakerFeeRateOk returns a tuple with the TakerFeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakerFeeRate

`func (o *AcquireMarketInformationResponse) SetTakerFeeRate(v string)`

SetTakerFeeRate sets TakerFeeRate field to given value.

### HasTakerFeeRate

`func (o *AcquireMarketInformationResponse) HasTakerFeeRate() bool`

HasTakerFeeRate returns a boolean if a field has been set.

### GetMakerFeeRate

`func (o *AcquireMarketInformationResponse) GetMakerFeeRate() string`

GetMakerFeeRate returns the MakerFeeRate field if non-nil, zero value otherwise.

### GetMakerFeeRateOk

`func (o *AcquireMarketInformationResponse) GetMakerFeeRateOk() (*string, bool)`

GetMakerFeeRateOk returns a tuple with the MakerFeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakerFeeRate

`func (o *AcquireMarketInformationResponse) SetMakerFeeRate(v string)`

SetMakerFeeRate sets MakerFeeRate field to given value.

### HasMakerFeeRate

`func (o *AcquireMarketInformationResponse) HasMakerFeeRate() bool`

HasMakerFeeRate returns a boolean if a field has been set.

### GetMinAmount

`func (o *AcquireMarketInformationResponse) GetMinAmount() string`

GetMinAmount returns the MinAmount field if non-nil, zero value otherwise.

### GetMinAmountOk

`func (o *AcquireMarketInformationResponse) GetMinAmountOk() (*string, bool)`

GetMinAmountOk returns a tuple with the MinAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAmount

`func (o *AcquireMarketInformationResponse) SetMinAmount(v string)`

SetMinAmount sets MinAmount field to given value.

### HasMinAmount

`func (o *AcquireMarketInformationResponse) HasMinAmount() bool`

HasMinAmount returns a boolean if a field has been set.

### GetTradingName

`func (o *AcquireMarketInformationResponse) GetTradingName() string`

GetTradingName returns the TradingName field if non-nil, zero value otherwise.

### GetTradingNameOk

`func (o *AcquireMarketInformationResponse) GetTradingNameOk() (*string, bool)`

GetTradingNameOk returns a tuple with the TradingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingName

`func (o *AcquireMarketInformationResponse) SetTradingName(v string)`

SetTradingName sets TradingName field to given value.

### HasTradingName

`func (o *AcquireMarketInformationResponse) HasTradingName() bool`

HasTradingName returns a boolean if a field has been set.

### GetTradingDecimal

`func (o *AcquireMarketInformationResponse) GetTradingDecimal() int64`

GetTradingDecimal returns the TradingDecimal field if non-nil, zero value otherwise.

### GetTradingDecimalOk

`func (o *AcquireMarketInformationResponse) GetTradingDecimalOk() (*int64, bool)`

GetTradingDecimalOk returns a tuple with the TradingDecimal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingDecimal

`func (o *AcquireMarketInformationResponse) SetTradingDecimal(v int64)`

SetTradingDecimal sets TradingDecimal field to given value.

### HasTradingDecimal

`func (o *AcquireMarketInformationResponse) HasTradingDecimal() bool`

HasTradingDecimal returns a boolean if a field has been set.

### GetPricingName

`func (o *AcquireMarketInformationResponse) GetPricingName() string`

GetPricingName returns the PricingName field if non-nil, zero value otherwise.

### GetPricingNameOk

`func (o *AcquireMarketInformationResponse) GetPricingNameOk() (*string, bool)`

GetPricingNameOk returns a tuple with the PricingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingName

`func (o *AcquireMarketInformationResponse) SetPricingName(v string)`

SetPricingName sets PricingName field to given value.

### HasPricingName

`func (o *AcquireMarketInformationResponse) HasPricingName() bool`

HasPricingName returns a boolean if a field has been set.

### GetPricingDecimal

`func (o *AcquireMarketInformationResponse) GetPricingDecimal() int64`

GetPricingDecimal returns the PricingDecimal field if non-nil, zero value otherwise.

### GetPricingDecimalOk

`func (o *AcquireMarketInformationResponse) GetPricingDecimalOk() (*int64, bool)`

GetPricingDecimalOk returns a tuple with the PricingDecimal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingDecimal

`func (o *AcquireMarketInformationResponse) SetPricingDecimal(v int64)`

SetPricingDecimal sets PricingDecimal field to given value.

### HasPricingDecimal

`func (o *AcquireMarketInformationResponse) HasPricingDecimal() bool`

HasPricingDecimal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


