# AcquireSingleMarketInformationResponse

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

### NewAcquireSingleMarketInformationResponse

`func NewAcquireSingleMarketInformationResponse() *AcquireSingleMarketInformationResponse`

NewAcquireSingleMarketInformationResponse instantiates a new AcquireSingleMarketInformationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcquireSingleMarketInformationResponseWithDefaults

`func NewAcquireSingleMarketInformationResponseWithDefaults() *AcquireSingleMarketInformationResponse`

NewAcquireSingleMarketInformationResponseWithDefaults instantiates a new AcquireSingleMarketInformationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AcquireSingleMarketInformationResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AcquireSingleMarketInformationResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AcquireSingleMarketInformationResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AcquireSingleMarketInformationResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTakerFeeRate

`func (o *AcquireSingleMarketInformationResponse) GetTakerFeeRate() string`

GetTakerFeeRate returns the TakerFeeRate field if non-nil, zero value otherwise.

### GetTakerFeeRateOk

`func (o *AcquireSingleMarketInformationResponse) GetTakerFeeRateOk() (*string, bool)`

GetTakerFeeRateOk returns a tuple with the TakerFeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakerFeeRate

`func (o *AcquireSingleMarketInformationResponse) SetTakerFeeRate(v string)`

SetTakerFeeRate sets TakerFeeRate field to given value.

### HasTakerFeeRate

`func (o *AcquireSingleMarketInformationResponse) HasTakerFeeRate() bool`

HasTakerFeeRate returns a boolean if a field has been set.

### GetMakerFeeRate

`func (o *AcquireSingleMarketInformationResponse) GetMakerFeeRate() string`

GetMakerFeeRate returns the MakerFeeRate field if non-nil, zero value otherwise.

### GetMakerFeeRateOk

`func (o *AcquireSingleMarketInformationResponse) GetMakerFeeRateOk() (*string, bool)`

GetMakerFeeRateOk returns a tuple with the MakerFeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakerFeeRate

`func (o *AcquireSingleMarketInformationResponse) SetMakerFeeRate(v string)`

SetMakerFeeRate sets MakerFeeRate field to given value.

### HasMakerFeeRate

`func (o *AcquireSingleMarketInformationResponse) HasMakerFeeRate() bool`

HasMakerFeeRate returns a boolean if a field has been set.

### GetMinAmount

`func (o *AcquireSingleMarketInformationResponse) GetMinAmount() string`

GetMinAmount returns the MinAmount field if non-nil, zero value otherwise.

### GetMinAmountOk

`func (o *AcquireSingleMarketInformationResponse) GetMinAmountOk() (*string, bool)`

GetMinAmountOk returns a tuple with the MinAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAmount

`func (o *AcquireSingleMarketInformationResponse) SetMinAmount(v string)`

SetMinAmount sets MinAmount field to given value.

### HasMinAmount

`func (o *AcquireSingleMarketInformationResponse) HasMinAmount() bool`

HasMinAmount returns a boolean if a field has been set.

### GetTradingName

`func (o *AcquireSingleMarketInformationResponse) GetTradingName() string`

GetTradingName returns the TradingName field if non-nil, zero value otherwise.

### GetTradingNameOk

`func (o *AcquireSingleMarketInformationResponse) GetTradingNameOk() (*string, bool)`

GetTradingNameOk returns a tuple with the TradingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingName

`func (o *AcquireSingleMarketInformationResponse) SetTradingName(v string)`

SetTradingName sets TradingName field to given value.

### HasTradingName

`func (o *AcquireSingleMarketInformationResponse) HasTradingName() bool`

HasTradingName returns a boolean if a field has been set.

### GetTradingDecimal

`func (o *AcquireSingleMarketInformationResponse) GetTradingDecimal() int64`

GetTradingDecimal returns the TradingDecimal field if non-nil, zero value otherwise.

### GetTradingDecimalOk

`func (o *AcquireSingleMarketInformationResponse) GetTradingDecimalOk() (*int64, bool)`

GetTradingDecimalOk returns a tuple with the TradingDecimal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingDecimal

`func (o *AcquireSingleMarketInformationResponse) SetTradingDecimal(v int64)`

SetTradingDecimal sets TradingDecimal field to given value.

### HasTradingDecimal

`func (o *AcquireSingleMarketInformationResponse) HasTradingDecimal() bool`

HasTradingDecimal returns a boolean if a field has been set.

### GetPricingName

`func (o *AcquireSingleMarketInformationResponse) GetPricingName() string`

GetPricingName returns the PricingName field if non-nil, zero value otherwise.

### GetPricingNameOk

`func (o *AcquireSingleMarketInformationResponse) GetPricingNameOk() (*string, bool)`

GetPricingNameOk returns a tuple with the PricingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingName

`func (o *AcquireSingleMarketInformationResponse) SetPricingName(v string)`

SetPricingName sets PricingName field to given value.

### HasPricingName

`func (o *AcquireSingleMarketInformationResponse) HasPricingName() bool`

HasPricingName returns a boolean if a field has been set.

### GetPricingDecimal

`func (o *AcquireSingleMarketInformationResponse) GetPricingDecimal() int64`

GetPricingDecimal returns the PricingDecimal field if non-nil, zero value otherwise.

### GetPricingDecimalOk

`func (o *AcquireSingleMarketInformationResponse) GetPricingDecimalOk() (*int64, bool)`

GetPricingDecimalOk returns a tuple with the PricingDecimal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingDecimal

`func (o *AcquireSingleMarketInformationResponse) SetPricingDecimal(v int64)`

SetPricingDecimal sets PricingDecimal field to given value.

### HasPricingDecimal

`func (o *AcquireSingleMarketInformationResponse) HasPricingDecimal() bool`

HasPricingDecimal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


