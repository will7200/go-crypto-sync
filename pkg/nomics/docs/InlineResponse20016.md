# InlineResponse20016

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **string** | RFC3339 timestamp | [optional] 
**Volume** | Pointer to **string** | Total volume (spot volume + derivative volume) in the target currency (&#x60;USD&#x60; by default) | [optional] 
**SpotVolume** | Pointer to **string** | Total spot market volume in the target currency (&#x60;USD&#x60; by default) | [optional] 
**DerivativeVolume** | Pointer to **string** | Total derivative market volume in the target currency (&#x60;USD&#x60; by default) | [optional] 
**TransparentVolume** | Pointer to **string** | Total transparent volume in the target currency (&#x60;USD&#x60; by default) | [optional] 
**TransparentSpotVolume** | Pointer to **string** | Total transparent spot market volume in the target currency (&#x60;USD&#x60; by default) | [optional] 
**TransparentDerivativeVolume** | Pointer to **string** | Total transparent derivative market volume in the target currency (&#x60;USD&#x60; by default) | [optional] 

## Methods

### NewInlineResponse20016

`func NewInlineResponse20016() *InlineResponse20016`

NewInlineResponse20016 instantiates a new InlineResponse20016 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20016WithDefaults

`func NewInlineResponse20016WithDefaults() *InlineResponse20016`

NewInlineResponse20016WithDefaults instantiates a new InlineResponse20016 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *InlineResponse20016) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *InlineResponse20016) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *InlineResponse20016) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *InlineResponse20016) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetVolume

`func (o *InlineResponse20016) GetVolume() string`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *InlineResponse20016) GetVolumeOk() (*string, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *InlineResponse20016) SetVolume(v string)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *InlineResponse20016) HasVolume() bool`

HasVolume returns a boolean if a field has been set.

### GetSpotVolume

`func (o *InlineResponse20016) GetSpotVolume() string`

GetSpotVolume returns the SpotVolume field if non-nil, zero value otherwise.

### GetSpotVolumeOk

`func (o *InlineResponse20016) GetSpotVolumeOk() (*string, bool)`

GetSpotVolumeOk returns a tuple with the SpotVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotVolume

`func (o *InlineResponse20016) SetSpotVolume(v string)`

SetSpotVolume sets SpotVolume field to given value.

### HasSpotVolume

`func (o *InlineResponse20016) HasSpotVolume() bool`

HasSpotVolume returns a boolean if a field has been set.

### GetDerivativeVolume

`func (o *InlineResponse20016) GetDerivativeVolume() string`

GetDerivativeVolume returns the DerivativeVolume field if non-nil, zero value otherwise.

### GetDerivativeVolumeOk

`func (o *InlineResponse20016) GetDerivativeVolumeOk() (*string, bool)`

GetDerivativeVolumeOk returns a tuple with the DerivativeVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivativeVolume

`func (o *InlineResponse20016) SetDerivativeVolume(v string)`

SetDerivativeVolume sets DerivativeVolume field to given value.

### HasDerivativeVolume

`func (o *InlineResponse20016) HasDerivativeVolume() bool`

HasDerivativeVolume returns a boolean if a field has been set.

### GetTransparentVolume

`func (o *InlineResponse20016) GetTransparentVolume() string`

GetTransparentVolume returns the TransparentVolume field if non-nil, zero value otherwise.

### GetTransparentVolumeOk

`func (o *InlineResponse20016) GetTransparentVolumeOk() (*string, bool)`

GetTransparentVolumeOk returns a tuple with the TransparentVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransparentVolume

`func (o *InlineResponse20016) SetTransparentVolume(v string)`

SetTransparentVolume sets TransparentVolume field to given value.

### HasTransparentVolume

`func (o *InlineResponse20016) HasTransparentVolume() bool`

HasTransparentVolume returns a boolean if a field has been set.

### GetTransparentSpotVolume

`func (o *InlineResponse20016) GetTransparentSpotVolume() string`

GetTransparentSpotVolume returns the TransparentSpotVolume field if non-nil, zero value otherwise.

### GetTransparentSpotVolumeOk

`func (o *InlineResponse20016) GetTransparentSpotVolumeOk() (*string, bool)`

GetTransparentSpotVolumeOk returns a tuple with the TransparentSpotVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransparentSpotVolume

`func (o *InlineResponse20016) SetTransparentSpotVolume(v string)`

SetTransparentSpotVolume sets TransparentSpotVolume field to given value.

### HasTransparentSpotVolume

`func (o *InlineResponse20016) HasTransparentSpotVolume() bool`

HasTransparentSpotVolume returns a boolean if a field has been set.

### GetTransparentDerivativeVolume

`func (o *InlineResponse20016) GetTransparentDerivativeVolume() string`

GetTransparentDerivativeVolume returns the TransparentDerivativeVolume field if non-nil, zero value otherwise.

### GetTransparentDerivativeVolumeOk

`func (o *InlineResponse20016) GetTransparentDerivativeVolumeOk() (*string, bool)`

GetTransparentDerivativeVolumeOk returns a tuple with the TransparentDerivativeVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransparentDerivativeVolume

`func (o *InlineResponse20016) SetTransparentDerivativeVolume(v string)`

SetTransparentDerivativeVolume sets TransparentDerivativeVolume field to given value.

### HasTransparentDerivativeVolume

`func (o *InlineResponse20016) HasTransparentDerivativeVolume() bool`

HasTransparentDerivativeVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


