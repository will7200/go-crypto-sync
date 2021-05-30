# AcquireLatestTransactionDataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Transaction No | [optional] 
**Date** | Pointer to **int64** | Transaction time | [optional] 
**DateMs** | Pointer to **int64** | Transaction time\\(ms\\) | [optional] 
**Amount** | Pointer to **string** | Transaction amount | [optional] 
**Price** | Pointer to **string** | Transaction price | [optional] 
**Type** | Pointer to **string** | buy;&lt;br&gt;sell; | [optional] 

## Methods

### NewAcquireLatestTransactionDataResponse

`func NewAcquireLatestTransactionDataResponse() *AcquireLatestTransactionDataResponse`

NewAcquireLatestTransactionDataResponse instantiates a new AcquireLatestTransactionDataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcquireLatestTransactionDataResponseWithDefaults

`func NewAcquireLatestTransactionDataResponseWithDefaults() *AcquireLatestTransactionDataResponse`

NewAcquireLatestTransactionDataResponseWithDefaults instantiates a new AcquireLatestTransactionDataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AcquireLatestTransactionDataResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AcquireLatestTransactionDataResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AcquireLatestTransactionDataResponse) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AcquireLatestTransactionDataResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDate

`func (o *AcquireLatestTransactionDataResponse) GetDate() int64`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *AcquireLatestTransactionDataResponse) GetDateOk() (*int64, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *AcquireLatestTransactionDataResponse) SetDate(v int64)`

SetDate sets Date field to given value.

### HasDate

`func (o *AcquireLatestTransactionDataResponse) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetDateMs

`func (o *AcquireLatestTransactionDataResponse) GetDateMs() int64`

GetDateMs returns the DateMs field if non-nil, zero value otherwise.

### GetDateMsOk

`func (o *AcquireLatestTransactionDataResponse) GetDateMsOk() (*int64, bool)`

GetDateMsOk returns a tuple with the DateMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateMs

`func (o *AcquireLatestTransactionDataResponse) SetDateMs(v int64)`

SetDateMs sets DateMs field to given value.

### HasDateMs

`func (o *AcquireLatestTransactionDataResponse) HasDateMs() bool`

HasDateMs returns a boolean if a field has been set.

### GetAmount

`func (o *AcquireLatestTransactionDataResponse) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AcquireLatestTransactionDataResponse) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AcquireLatestTransactionDataResponse) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *AcquireLatestTransactionDataResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetPrice

`func (o *AcquireLatestTransactionDataResponse) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *AcquireLatestTransactionDataResponse) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *AcquireLatestTransactionDataResponse) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *AcquireLatestTransactionDataResponse) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetType

`func (o *AcquireLatestTransactionDataResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AcquireLatestTransactionDataResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AcquireLatestTransactionDataResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AcquireLatestTransactionDataResponse) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


