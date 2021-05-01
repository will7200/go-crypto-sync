# InlineResponse20012

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BorrowEnabled** | Pointer to **bool** |  | [optional] 
**MarginLevel** | Pointer to **string** |  | [optional] 
**TotalAssetOfBtc** | Pointer to **string** |  | [optional] 
**TotalLiabilityOfBtc** | Pointer to **string** |  | [optional] 
**TotalNetAssetOfBtc** | Pointer to **string** |  | [optional] 
**TradeEnabled** | Pointer to **bool** |  | [optional] 
**TransferEnabled** | Pointer to **bool** |  | [optional] 
**UserAssets** | Pointer to [**[]InlineResponse20012UserAssets**](InlineResponse20012UserAssets.md) |  | [optional] 

## Methods

### NewInlineResponse20012

`func NewInlineResponse20012() *InlineResponse20012`

NewInlineResponse20012 instantiates a new InlineResponse20012 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20012WithDefaults

`func NewInlineResponse20012WithDefaults() *InlineResponse20012`

NewInlineResponse20012WithDefaults instantiates a new InlineResponse20012 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBorrowEnabled

`func (o *InlineResponse20012) GetBorrowEnabled() bool`

GetBorrowEnabled returns the BorrowEnabled field if non-nil, zero value otherwise.

### GetBorrowEnabledOk

`func (o *InlineResponse20012) GetBorrowEnabledOk() (*bool, bool)`

GetBorrowEnabledOk returns a tuple with the BorrowEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorrowEnabled

`func (o *InlineResponse20012) SetBorrowEnabled(v bool)`

SetBorrowEnabled sets BorrowEnabled field to given value.

### HasBorrowEnabled

`func (o *InlineResponse20012) HasBorrowEnabled() bool`

HasBorrowEnabled returns a boolean if a field has been set.

### GetMarginLevel

`func (o *InlineResponse20012) GetMarginLevel() string`

GetMarginLevel returns the MarginLevel field if non-nil, zero value otherwise.

### GetMarginLevelOk

`func (o *InlineResponse20012) GetMarginLevelOk() (*string, bool)`

GetMarginLevelOk returns a tuple with the MarginLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginLevel

`func (o *InlineResponse20012) SetMarginLevel(v string)`

SetMarginLevel sets MarginLevel field to given value.

### HasMarginLevel

`func (o *InlineResponse20012) HasMarginLevel() bool`

HasMarginLevel returns a boolean if a field has been set.

### GetTotalAssetOfBtc

`func (o *InlineResponse20012) GetTotalAssetOfBtc() string`

GetTotalAssetOfBtc returns the TotalAssetOfBtc field if non-nil, zero value otherwise.

### GetTotalAssetOfBtcOk

`func (o *InlineResponse20012) GetTotalAssetOfBtcOk() (*string, bool)`

GetTotalAssetOfBtcOk returns a tuple with the TotalAssetOfBtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAssetOfBtc

`func (o *InlineResponse20012) SetTotalAssetOfBtc(v string)`

SetTotalAssetOfBtc sets TotalAssetOfBtc field to given value.

### HasTotalAssetOfBtc

`func (o *InlineResponse20012) HasTotalAssetOfBtc() bool`

HasTotalAssetOfBtc returns a boolean if a field has been set.

### GetTotalLiabilityOfBtc

`func (o *InlineResponse20012) GetTotalLiabilityOfBtc() string`

GetTotalLiabilityOfBtc returns the TotalLiabilityOfBtc field if non-nil, zero value otherwise.

### GetTotalLiabilityOfBtcOk

`func (o *InlineResponse20012) GetTotalLiabilityOfBtcOk() (*string, bool)`

GetTotalLiabilityOfBtcOk returns a tuple with the TotalLiabilityOfBtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLiabilityOfBtc

`func (o *InlineResponse20012) SetTotalLiabilityOfBtc(v string)`

SetTotalLiabilityOfBtc sets TotalLiabilityOfBtc field to given value.

### HasTotalLiabilityOfBtc

`func (o *InlineResponse20012) HasTotalLiabilityOfBtc() bool`

HasTotalLiabilityOfBtc returns a boolean if a field has been set.

### GetTotalNetAssetOfBtc

`func (o *InlineResponse20012) GetTotalNetAssetOfBtc() string`

GetTotalNetAssetOfBtc returns the TotalNetAssetOfBtc field if non-nil, zero value otherwise.

### GetTotalNetAssetOfBtcOk

`func (o *InlineResponse20012) GetTotalNetAssetOfBtcOk() (*string, bool)`

GetTotalNetAssetOfBtcOk returns a tuple with the TotalNetAssetOfBtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNetAssetOfBtc

`func (o *InlineResponse20012) SetTotalNetAssetOfBtc(v string)`

SetTotalNetAssetOfBtc sets TotalNetAssetOfBtc field to given value.

### HasTotalNetAssetOfBtc

`func (o *InlineResponse20012) HasTotalNetAssetOfBtc() bool`

HasTotalNetAssetOfBtc returns a boolean if a field has been set.

### GetTradeEnabled

`func (o *InlineResponse20012) GetTradeEnabled() bool`

GetTradeEnabled returns the TradeEnabled field if non-nil, zero value otherwise.

### GetTradeEnabledOk

`func (o *InlineResponse20012) GetTradeEnabledOk() (*bool, bool)`

GetTradeEnabledOk returns a tuple with the TradeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeEnabled

`func (o *InlineResponse20012) SetTradeEnabled(v bool)`

SetTradeEnabled sets TradeEnabled field to given value.

### HasTradeEnabled

`func (o *InlineResponse20012) HasTradeEnabled() bool`

HasTradeEnabled returns a boolean if a field has been set.

### GetTransferEnabled

`func (o *InlineResponse20012) GetTransferEnabled() bool`

GetTransferEnabled returns the TransferEnabled field if non-nil, zero value otherwise.

### GetTransferEnabledOk

`func (o *InlineResponse20012) GetTransferEnabledOk() (*bool, bool)`

GetTransferEnabledOk returns a tuple with the TransferEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferEnabled

`func (o *InlineResponse20012) SetTransferEnabled(v bool)`

SetTransferEnabled sets TransferEnabled field to given value.

### HasTransferEnabled

`func (o *InlineResponse20012) HasTransferEnabled() bool`

HasTransferEnabled returns a boolean if a field has been set.

### GetUserAssets

`func (o *InlineResponse20012) GetUserAssets() []InlineResponse20012UserAssets`

GetUserAssets returns the UserAssets field if non-nil, zero value otherwise.

### GetUserAssetsOk

`func (o *InlineResponse20012) GetUserAssetsOk() (*[]InlineResponse20012UserAssets, bool)`

GetUserAssetsOk returns a tuple with the UserAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAssets

`func (o *InlineResponse20012) SetUserAssets(v []InlineResponse20012UserAssets)`

SetUserAssets sets UserAssets field to given value.

### HasUserAssets

`func (o *InlineResponse20012) HasUserAssets() bool`

HasUserAssets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


