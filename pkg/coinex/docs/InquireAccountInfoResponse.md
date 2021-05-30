# InquireAccountInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Frozen** | Pointer to **map[string]interface{}** | frozen amount | [optional] 
**Available** | Pointer to **map[string]interface{}** | available amount | [optional] 

## Methods

### NewInquireAccountInfoResponse

`func NewInquireAccountInfoResponse() *InquireAccountInfoResponse`

NewInquireAccountInfoResponse instantiates a new InquireAccountInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInquireAccountInfoResponseWithDefaults

`func NewInquireAccountInfoResponseWithDefaults() *InquireAccountInfoResponse`

NewInquireAccountInfoResponseWithDefaults instantiates a new InquireAccountInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrozen

`func (o *InquireAccountInfoResponse) GetFrozen() map[string]interface{}`

GetFrozen returns the Frozen field if non-nil, zero value otherwise.

### GetFrozenOk

`func (o *InquireAccountInfoResponse) GetFrozenOk() (*map[string]interface{}, bool)`

GetFrozenOk returns a tuple with the Frozen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozen

`func (o *InquireAccountInfoResponse) SetFrozen(v map[string]interface{})`

SetFrozen sets Frozen field to given value.

### HasFrozen

`func (o *InquireAccountInfoResponse) HasFrozen() bool`

HasFrozen returns a boolean if a field has been set.

### GetAvailable

`func (o *InquireAccountInfoResponse) GetAvailable() map[string]interface{}`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *InquireAccountInfoResponse) GetAvailableOk() (*map[string]interface{}, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *InquireAccountInfoResponse) SetAvailable(v map[string]interface{})`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *InquireAccountInfoResponse) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


