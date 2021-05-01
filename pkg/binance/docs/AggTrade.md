# AggTrade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**A** | Pointer to **int32** | trade id | [optional] 
**P** | Pointer to **string** | price | [optional] 
**Q** | Pointer to **string** | amount of base asset | [optional] 
**F** | Pointer to **int32** | First tradeId | [optional] 
**L** | Pointer to **int32** | Last tradeId | [optional] 
**T** | Pointer to **bool** | Timestamp | [optional] 
**M** | Pointer to **bool** | Was the buyer the maker? | [optional] 
**M** | Pointer to **bool** | Was the trade the best price match? | [optional] 

## Methods

### NewAggTrade

`func NewAggTrade() *AggTrade`

NewAggTrade instantiates a new AggTrade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggTradeWithDefaults

`func NewAggTradeWithDefaults() *AggTrade`

NewAggTradeWithDefaults instantiates a new AggTrade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetA

`func (o *AggTrade) GetA() int32`

GetA returns the A field if non-nil, zero value otherwise.

### GetAOk

`func (o *AggTrade) GetAOk() (*int32, bool)`

GetAOk returns a tuple with the A field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetA

`func (o *AggTrade) SetA(v int32)`

SetA sets A field to given value.

### HasA

`func (o *AggTrade) HasA() bool`

HasA returns a boolean if a field has been set.

### GetP

`func (o *AggTrade) GetP() string`

GetP returns the P field if non-nil, zero value otherwise.

### GetPOk

`func (o *AggTrade) GetPOk() (*string, bool)`

GetPOk returns a tuple with the P field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP

`func (o *AggTrade) SetP(v string)`

SetP sets P field to given value.

### HasP

`func (o *AggTrade) HasP() bool`

HasP returns a boolean if a field has been set.

### GetQ

`func (o *AggTrade) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *AggTrade) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *AggTrade) SetQ(v string)`

SetQ sets Q field to given value.

### HasQ

`func (o *AggTrade) HasQ() bool`

HasQ returns a boolean if a field has been set.

### GetF

`func (o *AggTrade) GetF() int32`

GetF returns the F field if non-nil, zero value otherwise.

### GetFOk

`func (o *AggTrade) GetFOk() (*int32, bool)`

GetFOk returns a tuple with the F field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetF

`func (o *AggTrade) SetF(v int32)`

SetF sets F field to given value.

### HasF

`func (o *AggTrade) HasF() bool`

HasF returns a boolean if a field has been set.

### GetL

`func (o *AggTrade) GetL() int32`

GetL returns the L field if non-nil, zero value otherwise.

### GetLOk

`func (o *AggTrade) GetLOk() (*int32, bool)`

GetLOk returns a tuple with the L field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL

`func (o *AggTrade) SetL(v int32)`

SetL sets L field to given value.

### HasL

`func (o *AggTrade) HasL() bool`

HasL returns a boolean if a field has been set.

### GetT

`func (o *AggTrade) GetT() bool`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *AggTrade) GetTOk() (*bool, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *AggTrade) SetT(v bool)`

SetT sets T field to given value.

### HasT

`func (o *AggTrade) HasT() bool`

HasT returns a boolean if a field has been set.

### GetM

`func (o *AggTrade) GetM() bool`

GetM returns the M field if non-nil, zero value otherwise.

### GetMOk

`func (o *AggTrade) GetMOk() (*bool, bool)`

GetMOk returns a tuple with the M field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetM

`func (o *AggTrade) SetM(v bool)`

SetM sets M field to given value.

### HasM

`func (o *AggTrade) HasM() bool`

HasM returns a boolean if a field has been set.

### GetM

`func (o *AggTrade) GetM() bool`

GetM returns the M field if non-nil, zero value otherwise.

### GetMOk

`func (o *AggTrade) GetMOk() (*bool, bool)`

GetMOk returns a tuple with the M field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetM

`func (o *AggTrade) SetM(v bool)`

SetM sets M field to given value.

### HasM

`func (o *AggTrade) HasM() bool`

HasM returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


