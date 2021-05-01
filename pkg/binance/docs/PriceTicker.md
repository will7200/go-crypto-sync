# PriceTicker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 

## Methods

### NewPriceTicker

`func NewPriceTicker() *PriceTicker`

NewPriceTicker instantiates a new PriceTicker object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceTickerWithDefaults

`func NewPriceTickerWithDefaults() *PriceTicker`

NewPriceTickerWithDefaults instantiates a new PriceTicker object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *PriceTicker) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *PriceTicker) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *PriceTicker) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *PriceTicker) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetPrice

`func (o *PriceTicker) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PriceTicker) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *PriceTicker) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *PriceTicker) HasPrice() bool`

HasPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


