# PlaceLimitOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessId** | **string** | access_id | 
**Market** | **string** | See &lt;API invocation description market&gt; | 
**Type** | **string** | sell: sell order | 

## Methods

### NewPlaceLimitOrderRequest

`func NewPlaceLimitOrderRequest(accessId string, market string, type_ string, ) *PlaceLimitOrderRequest`

NewPlaceLimitOrderRequest instantiates a new PlaceLimitOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaceLimitOrderRequestWithDefaults

`func NewPlaceLimitOrderRequestWithDefaults() *PlaceLimitOrderRequest`

NewPlaceLimitOrderRequestWithDefaults instantiates a new PlaceLimitOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessId

`func (o *PlaceLimitOrderRequest) GetAccessId() string`

GetAccessId returns the AccessId field if non-nil, zero value otherwise.

### GetAccessIdOk

`func (o *PlaceLimitOrderRequest) GetAccessIdOk() (*string, bool)`

GetAccessIdOk returns a tuple with the AccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessId

`func (o *PlaceLimitOrderRequest) SetAccessId(v string)`

SetAccessId sets AccessId field to given value.


### GetMarket

`func (o *PlaceLimitOrderRequest) GetMarket() string`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *PlaceLimitOrderRequest) GetMarketOk() (*string, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *PlaceLimitOrderRequest) SetMarket(v string)`

SetMarket sets Market field to given value.


### GetType

`func (o *PlaceLimitOrderRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PlaceLimitOrderRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PlaceLimitOrderRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


