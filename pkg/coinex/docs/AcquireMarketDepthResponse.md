# AcquireMarketDepthResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Last** | Pointer to **string** | Last price | [optional] 
**Time** | Pointer to **float64** | Updated time of Depth | [optional] 
**Asks** | Pointer to **[]map[string]interface{}** | Seller depth | [optional] 
**Bids** | Pointer to **[]map[string]interface{}** | Buyer depth | [optional] 

## Methods

### NewAcquireMarketDepthResponse

`func NewAcquireMarketDepthResponse() *AcquireMarketDepthResponse`

NewAcquireMarketDepthResponse instantiates a new AcquireMarketDepthResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcquireMarketDepthResponseWithDefaults

`func NewAcquireMarketDepthResponseWithDefaults() *AcquireMarketDepthResponse`

NewAcquireMarketDepthResponseWithDefaults instantiates a new AcquireMarketDepthResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLast

`func (o *AcquireMarketDepthResponse) GetLast() string`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *AcquireMarketDepthResponse) GetLastOk() (*string, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *AcquireMarketDepthResponse) SetLast(v string)`

SetLast sets Last field to given value.

### HasLast

`func (o *AcquireMarketDepthResponse) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetTime

`func (o *AcquireMarketDepthResponse) GetTime() float64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *AcquireMarketDepthResponse) GetTimeOk() (*float64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *AcquireMarketDepthResponse) SetTime(v float64)`

SetTime sets Time field to given value.

### HasTime

`func (o *AcquireMarketDepthResponse) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetAsks

`func (o *AcquireMarketDepthResponse) GetAsks() []map[string]interface{}`

GetAsks returns the Asks field if non-nil, zero value otherwise.

### GetAsksOk

`func (o *AcquireMarketDepthResponse) GetAsksOk() (*[]map[string]interface{}, bool)`

GetAsksOk returns a tuple with the Asks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsks

`func (o *AcquireMarketDepthResponse) SetAsks(v []map[string]interface{})`

SetAsks sets Asks field to given value.

### HasAsks

`func (o *AcquireMarketDepthResponse) HasAsks() bool`

HasAsks returns a boolean if a field has been set.

### GetBids

`func (o *AcquireMarketDepthResponse) GetBids() []map[string]interface{}`

GetBids returns the Bids field if non-nil, zero value otherwise.

### GetBidsOk

`func (o *AcquireMarketDepthResponse) GetBidsOk() (*[]map[string]interface{}, bool)`

GetBidsOk returns a tuple with the Bids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBids

`func (o *AcquireMarketDepthResponse) SetBids(v []map[string]interface{})`

SetBids sets Bids field to given value.

### HasBids

`func (o *AcquireMarketDepthResponse) HasBids() bool`

HasBids returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


