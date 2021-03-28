# Pagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndingBefore** | Pointer to **string** |  | [optional] 
**StartingAfter** | Pointer to **string** |  | [optional] 
**NextEndingBefore** | Pointer to **string** |  | [optional] 
**NextStartingAfter** | Pointer to **string** |  | [optional] 
**Limit** | Pointer to **float32** |  | [optional] [default to 25]
**Order** | Pointer to **string** |  | [optional] [default to "desc"]
**PreviousUri** | Pointer to **string** |  | [optional] 
**NextUri** | Pointer to **string** |  | [optional] 

## Methods

### NewPagination

`func NewPagination() *Pagination`

NewPagination instantiates a new Pagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationWithDefaults

`func NewPaginationWithDefaults() *Pagination`

NewPaginationWithDefaults instantiates a new Pagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndingBefore

`func (o *Pagination) GetEndingBefore() string`

GetEndingBefore returns the EndingBefore field if non-nil, zero value otherwise.

### GetEndingBeforeOk

`func (o *Pagination) GetEndingBeforeOk() (*string, bool)`

GetEndingBeforeOk returns a tuple with the EndingBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndingBefore

`func (o *Pagination) SetEndingBefore(v string)`

SetEndingBefore sets EndingBefore field to given value.

### HasEndingBefore

`func (o *Pagination) HasEndingBefore() bool`

HasEndingBefore returns a boolean if a field has been set.

### GetStartingAfter

`func (o *Pagination) GetStartingAfter() string`

GetStartingAfter returns the StartingAfter field if non-nil, zero value otherwise.

### GetStartingAfterOk

`func (o *Pagination) GetStartingAfterOk() (*string, bool)`

GetStartingAfterOk returns a tuple with the StartingAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingAfter

`func (o *Pagination) SetStartingAfter(v string)`

SetStartingAfter sets StartingAfter field to given value.

### HasStartingAfter

`func (o *Pagination) HasStartingAfter() bool`

HasStartingAfter returns a boolean if a field has been set.

### GetNextEndingBefore

`func (o *Pagination) GetNextEndingBefore() string`

GetNextEndingBefore returns the NextEndingBefore field if non-nil, zero value otherwise.

### GetNextEndingBeforeOk

`func (o *Pagination) GetNextEndingBeforeOk() (*string, bool)`

GetNextEndingBeforeOk returns a tuple with the NextEndingBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextEndingBefore

`func (o *Pagination) SetNextEndingBefore(v string)`

SetNextEndingBefore sets NextEndingBefore field to given value.

### HasNextEndingBefore

`func (o *Pagination) HasNextEndingBefore() bool`

HasNextEndingBefore returns a boolean if a field has been set.

### GetNextStartingAfter

`func (o *Pagination) GetNextStartingAfter() string`

GetNextStartingAfter returns the NextStartingAfter field if non-nil, zero value otherwise.

### GetNextStartingAfterOk

`func (o *Pagination) GetNextStartingAfterOk() (*string, bool)`

GetNextStartingAfterOk returns a tuple with the NextStartingAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextStartingAfter

`func (o *Pagination) SetNextStartingAfter(v string)`

SetNextStartingAfter sets NextStartingAfter field to given value.

### HasNextStartingAfter

`func (o *Pagination) HasNextStartingAfter() bool`

HasNextStartingAfter returns a boolean if a field has been set.

### GetLimit

`func (o *Pagination) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Pagination) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Pagination) SetLimit(v float32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *Pagination) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOrder

`func (o *Pagination) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *Pagination) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *Pagination) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *Pagination) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPreviousUri

`func (o *Pagination) GetPreviousUri() string`

GetPreviousUri returns the PreviousUri field if non-nil, zero value otherwise.

### GetPreviousUriOk

`func (o *Pagination) GetPreviousUriOk() (*string, bool)`

GetPreviousUriOk returns a tuple with the PreviousUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousUri

`func (o *Pagination) SetPreviousUri(v string)`

SetPreviousUri sets PreviousUri field to given value.

### HasPreviousUri

`func (o *Pagination) HasPreviousUri() bool`

HasPreviousUri returns a boolean if a field has been set.

### GetNextUri

`func (o *Pagination) GetNextUri() string`

GetNextUri returns the NextUri field if non-nil, zero value otherwise.

### GetNextUriOk

`func (o *Pagination) GetNextUriOk() (*string, bool)`

GetNextUriOk returns a tuple with the NextUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextUri

`func (o *Pagination) SetNextUri(v string)`

SetNextUri sets NextUri field to given value.

### HasNextUri

`func (o *Pagination) HasNextUri() bool`

HasNextUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


