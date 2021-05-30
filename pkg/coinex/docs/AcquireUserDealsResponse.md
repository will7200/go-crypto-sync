# AcquireUserDealsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **string** | executed amount | [optional] 
**CreateTime** | Pointer to **int64** | executed time | [optional] 
**DealMoney** | Pointer to **string** | executed value | [optional] 
**Fee** | Pointer to **string** | transaction fee | [optional] 
**FeeAsset** | Pointer to **string** | transaction fee asset | [optional] 
**Id** | Pointer to **int64** | executed id | [optional] 
**OrderId** | Pointer to **int64** | order no. | [optional] 
**Price** | Pointer to **string** | order price | [optional] 
**Role** | Pointer to **string** | order role | [optional] 
**Type** | Pointer to **string** | sell:sell;&lt;br&gt;buy:buy | [optional] 

## Methods

### NewAcquireUserDealsResponse

`func NewAcquireUserDealsResponse() *AcquireUserDealsResponse`

NewAcquireUserDealsResponse instantiates a new AcquireUserDealsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcquireUserDealsResponseWithDefaults

`func NewAcquireUserDealsResponseWithDefaults() *AcquireUserDealsResponse`

NewAcquireUserDealsResponseWithDefaults instantiates a new AcquireUserDealsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *AcquireUserDealsResponse) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AcquireUserDealsResponse) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AcquireUserDealsResponse) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *AcquireUserDealsResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCreateTime

`func (o *AcquireUserDealsResponse) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *AcquireUserDealsResponse) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *AcquireUserDealsResponse) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *AcquireUserDealsResponse) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDealMoney

`func (o *AcquireUserDealsResponse) GetDealMoney() string`

GetDealMoney returns the DealMoney field if non-nil, zero value otherwise.

### GetDealMoneyOk

`func (o *AcquireUserDealsResponse) GetDealMoneyOk() (*string, bool)`

GetDealMoneyOk returns a tuple with the DealMoney field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealMoney

`func (o *AcquireUserDealsResponse) SetDealMoney(v string)`

SetDealMoney sets DealMoney field to given value.

### HasDealMoney

`func (o *AcquireUserDealsResponse) HasDealMoney() bool`

HasDealMoney returns a boolean if a field has been set.

### GetFee

`func (o *AcquireUserDealsResponse) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *AcquireUserDealsResponse) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *AcquireUserDealsResponse) SetFee(v string)`

SetFee sets Fee field to given value.

### HasFee

`func (o *AcquireUserDealsResponse) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetFeeAsset

`func (o *AcquireUserDealsResponse) GetFeeAsset() string`

GetFeeAsset returns the FeeAsset field if non-nil, zero value otherwise.

### GetFeeAssetOk

`func (o *AcquireUserDealsResponse) GetFeeAssetOk() (*string, bool)`

GetFeeAssetOk returns a tuple with the FeeAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAsset

`func (o *AcquireUserDealsResponse) SetFeeAsset(v string)`

SetFeeAsset sets FeeAsset field to given value.

### HasFeeAsset

`func (o *AcquireUserDealsResponse) HasFeeAsset() bool`

HasFeeAsset returns a boolean if a field has been set.

### GetId

`func (o *AcquireUserDealsResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AcquireUserDealsResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AcquireUserDealsResponse) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AcquireUserDealsResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrderId

`func (o *AcquireUserDealsResponse) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *AcquireUserDealsResponse) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *AcquireUserDealsResponse) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *AcquireUserDealsResponse) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetPrice

`func (o *AcquireUserDealsResponse) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *AcquireUserDealsResponse) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *AcquireUserDealsResponse) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *AcquireUserDealsResponse) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetRole

`func (o *AcquireUserDealsResponse) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AcquireUserDealsResponse) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AcquireUserDealsResponse) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *AcquireUserDealsResponse) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetType

`func (o *AcquireUserDealsResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AcquireUserDealsResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AcquireUserDealsResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AcquireUserDealsResponse) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


