# InlineResponse20018

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **string** | RFC3339 timestamp of the order book snapshot | [optional] 
**Bids** | Pointer to [**[][]float32**](array.md) | Tuples of bids in descending order (first entry is best bid). First entry is price in the quote currency. Second entry is volume of the base currency. | [optional] 
**Asks** | Pointer to [**[][]float32**](array.md) | Tuples of asks in ascending order (first entry is best ask). First entry is price in the quote currency. Second entry is volume of the base currency. | [optional] 

## Methods

### NewInlineResponse20018

`func NewInlineResponse20018() *InlineResponse20018`

NewInlineResponse20018 instantiates a new InlineResponse20018 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20018WithDefaults

`func NewInlineResponse20018WithDefaults() *InlineResponse20018`

NewInlineResponse20018WithDefaults instantiates a new InlineResponse20018 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *InlineResponse20018) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *InlineResponse20018) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *InlineResponse20018) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *InlineResponse20018) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetBids

`func (o *InlineResponse20018) GetBids() [][]float32`

GetBids returns the Bids field if non-nil, zero value otherwise.

### GetBidsOk

`func (o *InlineResponse20018) GetBidsOk() (*[][]float32, bool)`

GetBidsOk returns a tuple with the Bids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBids

`func (o *InlineResponse20018) SetBids(v [][]float32)`

SetBids sets Bids field to given value.

### HasBids

`func (o *InlineResponse20018) HasBids() bool`

HasBids returns a boolean if a field has been set.

### GetAsks

`func (o *InlineResponse20018) GetAsks() [][]float32`

GetAsks returns the Asks field if non-nil, zero value otherwise.

### GetAsksOk

`func (o *InlineResponse20018) GetAsksOk() (*[][]float32, bool)`

GetAsksOk returns a tuple with the Asks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsks

`func (o *InlineResponse20018) SetAsks(v [][]float32)`

SetAsks sets Asks field to given value.

### HasAsks

`func (o *InlineResponse20018) HasAsks() bool`

HasAsks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


