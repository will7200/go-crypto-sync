# InquireAccountInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int64** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**map[string]Object**](Object.md) |  | [optional] 

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

### GetCode

`func (o *InquireAccountInfoResponse) GetCode() int64`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *InquireAccountInfoResponse) GetCodeOk() (*int64, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *InquireAccountInfoResponse) SetCode(v int64)`

SetCode sets Code field to given value.

### HasCode

`func (o *InquireAccountInfoResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *InquireAccountInfoResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InquireAccountInfoResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InquireAccountInfoResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *InquireAccountInfoResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetData

`func (o *InquireAccountInfoResponse) GetData() map[string]Object`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InquireAccountInfoResponse) GetDataOk() (*map[string]Object, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InquireAccountInfoResponse) SetData(v map[string]Object)`

SetData sets Data field to given value.

### HasData

`func (o *InquireAccountInfoResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


