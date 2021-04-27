# InlineResponse20015

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **string** | Start time of the candle in RFC3339 | [optional] 
**Open** | Pointer to **string** | First trade price in USD | [optional] 
**High** | Pointer to **string** | Highest price in USD | [optional] 
**Low** | Pointer to **string** | Lowest price in USD | [optional] 
**Close** | Pointer to **string** | Last trade price in USD | [optional] 
**Volume** | Pointer to **string** | Volume in USD | [optional] 
**TransparentOpen** | Pointer to **string** | First transparent trade price in USD | [optional] 
**TransparentHigh** | Pointer to **string** | Highest transparent price in USD | [optional] 
**TransparentLow** | Pointer to **string** | Lowest transparent price in USD | [optional] 
**TransparentClose** | Pointer to **string** | Last transparent trade price in USD | [optional] 
**TransparentVolume** | Pointer to **string** | Transparent volume in USD | [optional] 
**VolumeTransparency** | Pointer to [**CandlesVolumeTransparency**](_candles_volume_transparency.md) |  | [optional] 

## Methods

### NewInlineResponse20015

`func NewInlineResponse20015() *InlineResponse20015`

NewInlineResponse20015 instantiates a new InlineResponse20015 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20015WithDefaults

`func NewInlineResponse20015WithDefaults() *InlineResponse20015`

NewInlineResponse20015WithDefaults instantiates a new InlineResponse20015 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *InlineResponse20015) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *InlineResponse20015) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *InlineResponse20015) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *InlineResponse20015) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetOpen

`func (o *InlineResponse20015) GetOpen() string`

GetOpen returns the Open field if non-nil, zero value otherwise.

### GetOpenOk

`func (o *InlineResponse20015) GetOpenOk() (*string, bool)`

GetOpenOk returns a tuple with the Open field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpen

`func (o *InlineResponse20015) SetOpen(v string)`

SetOpen sets Open field to given value.

### HasOpen

`func (o *InlineResponse20015) HasOpen() bool`

HasOpen returns a boolean if a field has been set.

### GetHigh

`func (o *InlineResponse20015) GetHigh() string`

GetHigh returns the High field if non-nil, zero value otherwise.

### GetHighOk

`func (o *InlineResponse20015) GetHighOk() (*string, bool)`

GetHighOk returns a tuple with the High field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigh

`func (o *InlineResponse20015) SetHigh(v string)`

SetHigh sets High field to given value.

### HasHigh

`func (o *InlineResponse20015) HasHigh() bool`

HasHigh returns a boolean if a field has been set.

### GetLow

`func (o *InlineResponse20015) GetLow() string`

GetLow returns the Low field if non-nil, zero value otherwise.

### GetLowOk

`func (o *InlineResponse20015) GetLowOk() (*string, bool)`

GetLowOk returns a tuple with the Low field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLow

`func (o *InlineResponse20015) SetLow(v string)`

SetLow sets Low field to given value.

### HasLow

`func (o *InlineResponse20015) HasLow() bool`

HasLow returns a boolean if a field has been set.

### GetClose

`func (o *InlineResponse20015) GetClose() string`

GetClose returns the Close field if non-nil, zero value otherwise.

### GetCloseOk

`func (o *InlineResponse20015) GetCloseOk() (*string, bool)`

GetCloseOk returns a tuple with the Close field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClose

`func (o *InlineResponse20015) SetClose(v string)`

SetClose sets Close field to given value.

### HasClose

`func (o *InlineResponse20015) HasClose() bool`

HasClose returns a boolean if a field has been set.

### GetVolume

`func (o *InlineResponse20015) GetVolume() string`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *InlineResponse20015) GetVolumeOk() (*string, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *InlineResponse20015) SetVolume(v string)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *InlineResponse20015) HasVolume() bool`

HasVolume returns a boolean if a field has been set.

### GetTransparentOpen

`func (o *InlineResponse20015) GetTransparentOpen() string`

GetTransparentOpen returns the TransparentOpen field if non-nil, zero value otherwise.

### GetTransparentOpenOk

`func (o *InlineResponse20015) GetTransparentOpenOk() (*string, bool)`

GetTransparentOpenOk returns a tuple with the TransparentOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransparentOpen

`func (o *InlineResponse20015) SetTransparentOpen(v string)`

SetTransparentOpen sets TransparentOpen field to given value.

### HasTransparentOpen

`func (o *InlineResponse20015) HasTransparentOpen() bool`

HasTransparentOpen returns a boolean if a field has been set.

### GetTransparentHigh

`func (o *InlineResponse20015) GetTransparentHigh() string`

GetTransparentHigh returns the TransparentHigh field if non-nil, zero value otherwise.

### GetTransparentHighOk

`func (o *InlineResponse20015) GetTransparentHighOk() (*string, bool)`

GetTransparentHighOk returns a tuple with the TransparentHigh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransparentHigh

`func (o *InlineResponse20015) SetTransparentHigh(v string)`

SetTransparentHigh sets TransparentHigh field to given value.

### HasTransparentHigh

`func (o *InlineResponse20015) HasTransparentHigh() bool`

HasTransparentHigh returns a boolean if a field has been set.

### GetTransparentLow

`func (o *InlineResponse20015) GetTransparentLow() string`

GetTransparentLow returns the TransparentLow field if non-nil, zero value otherwise.

### GetTransparentLowOk

`func (o *InlineResponse20015) GetTransparentLowOk() (*string, bool)`

GetTransparentLowOk returns a tuple with the TransparentLow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransparentLow

`func (o *InlineResponse20015) SetTransparentLow(v string)`

SetTransparentLow sets TransparentLow field to given value.

### HasTransparentLow

`func (o *InlineResponse20015) HasTransparentLow() bool`

HasTransparentLow returns a boolean if a field has been set.

### GetTransparentClose

`func (o *InlineResponse20015) GetTransparentClose() string`

GetTransparentClose returns the TransparentClose field if non-nil, zero value otherwise.

### GetTransparentCloseOk

`func (o *InlineResponse20015) GetTransparentCloseOk() (*string, bool)`

GetTransparentCloseOk returns a tuple with the TransparentClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransparentClose

`func (o *InlineResponse20015) SetTransparentClose(v string)`

SetTransparentClose sets TransparentClose field to given value.

### HasTransparentClose

`func (o *InlineResponse20015) HasTransparentClose() bool`

HasTransparentClose returns a boolean if a field has been set.

### GetTransparentVolume

`func (o *InlineResponse20015) GetTransparentVolume() string`

GetTransparentVolume returns the TransparentVolume field if non-nil, zero value otherwise.

### GetTransparentVolumeOk

`func (o *InlineResponse20015) GetTransparentVolumeOk() (*string, bool)`

GetTransparentVolumeOk returns a tuple with the TransparentVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransparentVolume

`func (o *InlineResponse20015) SetTransparentVolume(v string)`

SetTransparentVolume sets TransparentVolume field to given value.

### HasTransparentVolume

`func (o *InlineResponse20015) HasTransparentVolume() bool`

HasTransparentVolume returns a boolean if a field has been set.

### GetVolumeTransparency

`func (o *InlineResponse20015) GetVolumeTransparency() CandlesVolumeTransparency`

GetVolumeTransparency returns the VolumeTransparency field if non-nil, zero value otherwise.

### GetVolumeTransparencyOk

`func (o *InlineResponse20015) GetVolumeTransparencyOk() (*CandlesVolumeTransparency, bool)`

GetVolumeTransparencyOk returns a tuple with the VolumeTransparency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeTransparency

`func (o *InlineResponse20015) SetVolumeTransparency(v CandlesVolumeTransparency)`

SetVolumeTransparency sets VolumeTransparency field to given value.

### HasVolumeTransparency

`func (o *InlineResponse20015) HasVolumeTransparency() bool`

HasVolumeTransparency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


