# PlaceMultipleLimitOrdersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessId** | **string** | access_id | 
**BatchOrders** | Pointer to **string** |  | [optional] 

## Methods

### NewPlaceMultipleLimitOrdersRequest

`func NewPlaceMultipleLimitOrdersRequest(accessId string, ) *PlaceMultipleLimitOrdersRequest`

NewPlaceMultipleLimitOrdersRequest instantiates a new PlaceMultipleLimitOrdersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaceMultipleLimitOrdersRequestWithDefaults

`func NewPlaceMultipleLimitOrdersRequestWithDefaults() *PlaceMultipleLimitOrdersRequest`

NewPlaceMultipleLimitOrdersRequestWithDefaults instantiates a new PlaceMultipleLimitOrdersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessId

`func (o *PlaceMultipleLimitOrdersRequest) GetAccessId() string`

GetAccessId returns the AccessId field if non-nil, zero value otherwise.

### GetAccessIdOk

`func (o *PlaceMultipleLimitOrdersRequest) GetAccessIdOk() (*string, bool)`

GetAccessIdOk returns a tuple with the AccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessId

`func (o *PlaceMultipleLimitOrdersRequest) SetAccessId(v string)`

SetAccessId sets AccessId field to given value.


### GetBatchOrders

`func (o *PlaceMultipleLimitOrdersRequest) GetBatchOrders() string`

GetBatchOrders returns the BatchOrders field if non-nil, zero value otherwise.

### GetBatchOrdersOk

`func (o *PlaceMultipleLimitOrdersRequest) GetBatchOrdersOk() (*string, bool)`

GetBatchOrdersOk returns a tuple with the BatchOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchOrders

`func (o *PlaceMultipleLimitOrdersRequest) SetBatchOrders(v string)`

SetBatchOrders sets BatchOrders field to given value.

### HasBatchOrders

`func (o *PlaceMultipleLimitOrdersRequest) HasBatchOrders() bool`

HasBatchOrders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


