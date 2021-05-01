# InlineResponse2001Symbols

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**BaseAsset** | Pointer to **string** |  | [optional] 
**BaseAssetPrecision** | Pointer to **int32** |  | [optional] 
**QuoteAsset** | Pointer to **string** |  | [optional] 
**QuoteAssetPrecision** | Pointer to **int32** |  | [optional] 
**BaseCommissionPrecision** | Pointer to **int32** |  | [optional] 
**QuoteCommissionPrecision** | Pointer to **int32** |  | [optional] 
**OrderTypes** | Pointer to **[]string** |  | [optional] 
**IcebergAllowed** | Pointer to **bool** |  | [optional] 
**OcoAllowed** | Pointer to **bool** |  | [optional] 
**QuoteOrderQtyMarketAllowed** | Pointer to **bool** |  | [optional] 
**IsSpotTradingAllowed** | Pointer to **bool** |  | [optional] 
**IsMarginTradingAllowed** | Pointer to **bool** |  | [optional] 
**Filters** | Pointer to [**[]InlineResponse2001Filters**](InlineResponse2001Filters.md) |  | [optional] 
**Permissions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewInlineResponse2001Symbols

`func NewInlineResponse2001Symbols() *InlineResponse2001Symbols`

NewInlineResponse2001Symbols instantiates a new InlineResponse2001Symbols object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2001SymbolsWithDefaults

`func NewInlineResponse2001SymbolsWithDefaults() *InlineResponse2001Symbols`

NewInlineResponse2001SymbolsWithDefaults instantiates a new InlineResponse2001Symbols object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *InlineResponse2001Symbols) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *InlineResponse2001Symbols) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *InlineResponse2001Symbols) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *InlineResponse2001Symbols) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse2001Symbols) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse2001Symbols) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse2001Symbols) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse2001Symbols) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetBaseAsset

`func (o *InlineResponse2001Symbols) GetBaseAsset() string`

GetBaseAsset returns the BaseAsset field if non-nil, zero value otherwise.

### GetBaseAssetOk

`func (o *InlineResponse2001Symbols) GetBaseAssetOk() (*string, bool)`

GetBaseAssetOk returns a tuple with the BaseAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAsset

`func (o *InlineResponse2001Symbols) SetBaseAsset(v string)`

SetBaseAsset sets BaseAsset field to given value.

### HasBaseAsset

`func (o *InlineResponse2001Symbols) HasBaseAsset() bool`

HasBaseAsset returns a boolean if a field has been set.

### GetBaseAssetPrecision

`func (o *InlineResponse2001Symbols) GetBaseAssetPrecision() int32`

GetBaseAssetPrecision returns the BaseAssetPrecision field if non-nil, zero value otherwise.

### GetBaseAssetPrecisionOk

`func (o *InlineResponse2001Symbols) GetBaseAssetPrecisionOk() (*int32, bool)`

GetBaseAssetPrecisionOk returns a tuple with the BaseAssetPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAssetPrecision

`func (o *InlineResponse2001Symbols) SetBaseAssetPrecision(v int32)`

SetBaseAssetPrecision sets BaseAssetPrecision field to given value.

### HasBaseAssetPrecision

`func (o *InlineResponse2001Symbols) HasBaseAssetPrecision() bool`

HasBaseAssetPrecision returns a boolean if a field has been set.

### GetQuoteAsset

`func (o *InlineResponse2001Symbols) GetQuoteAsset() string`

GetQuoteAsset returns the QuoteAsset field if non-nil, zero value otherwise.

### GetQuoteAssetOk

`func (o *InlineResponse2001Symbols) GetQuoteAssetOk() (*string, bool)`

GetQuoteAssetOk returns a tuple with the QuoteAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteAsset

`func (o *InlineResponse2001Symbols) SetQuoteAsset(v string)`

SetQuoteAsset sets QuoteAsset field to given value.

### HasQuoteAsset

`func (o *InlineResponse2001Symbols) HasQuoteAsset() bool`

HasQuoteAsset returns a boolean if a field has been set.

### GetQuoteAssetPrecision

`func (o *InlineResponse2001Symbols) GetQuoteAssetPrecision() int32`

GetQuoteAssetPrecision returns the QuoteAssetPrecision field if non-nil, zero value otherwise.

### GetQuoteAssetPrecisionOk

`func (o *InlineResponse2001Symbols) GetQuoteAssetPrecisionOk() (*int32, bool)`

GetQuoteAssetPrecisionOk returns a tuple with the QuoteAssetPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteAssetPrecision

`func (o *InlineResponse2001Symbols) SetQuoteAssetPrecision(v int32)`

SetQuoteAssetPrecision sets QuoteAssetPrecision field to given value.

### HasQuoteAssetPrecision

`func (o *InlineResponse2001Symbols) HasQuoteAssetPrecision() bool`

HasQuoteAssetPrecision returns a boolean if a field has been set.

### GetBaseCommissionPrecision

`func (o *InlineResponse2001Symbols) GetBaseCommissionPrecision() int32`

GetBaseCommissionPrecision returns the BaseCommissionPrecision field if non-nil, zero value otherwise.

### GetBaseCommissionPrecisionOk

`func (o *InlineResponse2001Symbols) GetBaseCommissionPrecisionOk() (*int32, bool)`

GetBaseCommissionPrecisionOk returns a tuple with the BaseCommissionPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCommissionPrecision

`func (o *InlineResponse2001Symbols) SetBaseCommissionPrecision(v int32)`

SetBaseCommissionPrecision sets BaseCommissionPrecision field to given value.

### HasBaseCommissionPrecision

`func (o *InlineResponse2001Symbols) HasBaseCommissionPrecision() bool`

HasBaseCommissionPrecision returns a boolean if a field has been set.

### GetQuoteCommissionPrecision

`func (o *InlineResponse2001Symbols) GetQuoteCommissionPrecision() int32`

GetQuoteCommissionPrecision returns the QuoteCommissionPrecision field if non-nil, zero value otherwise.

### GetQuoteCommissionPrecisionOk

`func (o *InlineResponse2001Symbols) GetQuoteCommissionPrecisionOk() (*int32, bool)`

GetQuoteCommissionPrecisionOk returns a tuple with the QuoteCommissionPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteCommissionPrecision

`func (o *InlineResponse2001Symbols) SetQuoteCommissionPrecision(v int32)`

SetQuoteCommissionPrecision sets QuoteCommissionPrecision field to given value.

### HasQuoteCommissionPrecision

`func (o *InlineResponse2001Symbols) HasQuoteCommissionPrecision() bool`

HasQuoteCommissionPrecision returns a boolean if a field has been set.

### GetOrderTypes

`func (o *InlineResponse2001Symbols) GetOrderTypes() []string`

GetOrderTypes returns the OrderTypes field if non-nil, zero value otherwise.

### GetOrderTypesOk

`func (o *InlineResponse2001Symbols) GetOrderTypesOk() (*[]string, bool)`

GetOrderTypesOk returns a tuple with the OrderTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderTypes

`func (o *InlineResponse2001Symbols) SetOrderTypes(v []string)`

SetOrderTypes sets OrderTypes field to given value.

### HasOrderTypes

`func (o *InlineResponse2001Symbols) HasOrderTypes() bool`

HasOrderTypes returns a boolean if a field has been set.

### GetIcebergAllowed

`func (o *InlineResponse2001Symbols) GetIcebergAllowed() bool`

GetIcebergAllowed returns the IcebergAllowed field if non-nil, zero value otherwise.

### GetIcebergAllowedOk

`func (o *InlineResponse2001Symbols) GetIcebergAllowedOk() (*bool, bool)`

GetIcebergAllowedOk returns a tuple with the IcebergAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcebergAllowed

`func (o *InlineResponse2001Symbols) SetIcebergAllowed(v bool)`

SetIcebergAllowed sets IcebergAllowed field to given value.

### HasIcebergAllowed

`func (o *InlineResponse2001Symbols) HasIcebergAllowed() bool`

HasIcebergAllowed returns a boolean if a field has been set.

### GetOcoAllowed

`func (o *InlineResponse2001Symbols) GetOcoAllowed() bool`

GetOcoAllowed returns the OcoAllowed field if non-nil, zero value otherwise.

### GetOcoAllowedOk

`func (o *InlineResponse2001Symbols) GetOcoAllowedOk() (*bool, bool)`

GetOcoAllowedOk returns a tuple with the OcoAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcoAllowed

`func (o *InlineResponse2001Symbols) SetOcoAllowed(v bool)`

SetOcoAllowed sets OcoAllowed field to given value.

### HasOcoAllowed

`func (o *InlineResponse2001Symbols) HasOcoAllowed() bool`

HasOcoAllowed returns a boolean if a field has been set.

### GetQuoteOrderQtyMarketAllowed

`func (o *InlineResponse2001Symbols) GetQuoteOrderQtyMarketAllowed() bool`

GetQuoteOrderQtyMarketAllowed returns the QuoteOrderQtyMarketAllowed field if non-nil, zero value otherwise.

### GetQuoteOrderQtyMarketAllowedOk

`func (o *InlineResponse2001Symbols) GetQuoteOrderQtyMarketAllowedOk() (*bool, bool)`

GetQuoteOrderQtyMarketAllowedOk returns a tuple with the QuoteOrderQtyMarketAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteOrderQtyMarketAllowed

`func (o *InlineResponse2001Symbols) SetQuoteOrderQtyMarketAllowed(v bool)`

SetQuoteOrderQtyMarketAllowed sets QuoteOrderQtyMarketAllowed field to given value.

### HasQuoteOrderQtyMarketAllowed

`func (o *InlineResponse2001Symbols) HasQuoteOrderQtyMarketAllowed() bool`

HasQuoteOrderQtyMarketAllowed returns a boolean if a field has been set.

### GetIsSpotTradingAllowed

`func (o *InlineResponse2001Symbols) GetIsSpotTradingAllowed() bool`

GetIsSpotTradingAllowed returns the IsSpotTradingAllowed field if non-nil, zero value otherwise.

### GetIsSpotTradingAllowedOk

`func (o *InlineResponse2001Symbols) GetIsSpotTradingAllowedOk() (*bool, bool)`

GetIsSpotTradingAllowedOk returns a tuple with the IsSpotTradingAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpotTradingAllowed

`func (o *InlineResponse2001Symbols) SetIsSpotTradingAllowed(v bool)`

SetIsSpotTradingAllowed sets IsSpotTradingAllowed field to given value.

### HasIsSpotTradingAllowed

`func (o *InlineResponse2001Symbols) HasIsSpotTradingAllowed() bool`

HasIsSpotTradingAllowed returns a boolean if a field has been set.

### GetIsMarginTradingAllowed

`func (o *InlineResponse2001Symbols) GetIsMarginTradingAllowed() bool`

GetIsMarginTradingAllowed returns the IsMarginTradingAllowed field if non-nil, zero value otherwise.

### GetIsMarginTradingAllowedOk

`func (o *InlineResponse2001Symbols) GetIsMarginTradingAllowedOk() (*bool, bool)`

GetIsMarginTradingAllowedOk returns a tuple with the IsMarginTradingAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMarginTradingAllowed

`func (o *InlineResponse2001Symbols) SetIsMarginTradingAllowed(v bool)`

SetIsMarginTradingAllowed sets IsMarginTradingAllowed field to given value.

### HasIsMarginTradingAllowed

`func (o *InlineResponse2001Symbols) HasIsMarginTradingAllowed() bool`

HasIsMarginTradingAllowed returns a boolean if a field has been set.

### GetFilters

`func (o *InlineResponse2001Symbols) GetFilters() []InlineResponse2001Filters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *InlineResponse2001Symbols) GetFiltersOk() (*[]InlineResponse2001Filters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *InlineResponse2001Symbols) SetFilters(v []InlineResponse2001Filters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *InlineResponse2001Symbols) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetPermissions

`func (o *InlineResponse2001Symbols) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *InlineResponse2001Symbols) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *InlineResponse2001Symbols) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *InlineResponse2001Symbols) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


