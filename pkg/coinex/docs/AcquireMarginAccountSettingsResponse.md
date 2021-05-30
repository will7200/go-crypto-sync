# AcquireMarginAccountSettingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **string** | maximum loan amount | [optional] 
**DayRate** | Pointer to **string** | daily interest rate | [optional] 
**Leverage** | Pointer to **int64** | maximum leverage | [optional] 

## Methods

### NewAcquireMarginAccountSettingsResponse

`func NewAcquireMarginAccountSettingsResponse() *AcquireMarginAccountSettingsResponse`

NewAcquireMarginAccountSettingsResponse instantiates a new AcquireMarginAccountSettingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcquireMarginAccountSettingsResponseWithDefaults

`func NewAcquireMarginAccountSettingsResponseWithDefaults() *AcquireMarginAccountSettingsResponse`

NewAcquireMarginAccountSettingsResponseWithDefaults instantiates a new AcquireMarginAccountSettingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *AcquireMarginAccountSettingsResponse) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AcquireMarginAccountSettingsResponse) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AcquireMarginAccountSettingsResponse) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *AcquireMarginAccountSettingsResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetDayRate

`func (o *AcquireMarginAccountSettingsResponse) GetDayRate() string`

GetDayRate returns the DayRate field if non-nil, zero value otherwise.

### GetDayRateOk

`func (o *AcquireMarginAccountSettingsResponse) GetDayRateOk() (*string, bool)`

GetDayRateOk returns a tuple with the DayRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayRate

`func (o *AcquireMarginAccountSettingsResponse) SetDayRate(v string)`

SetDayRate sets DayRate field to given value.

### HasDayRate

`func (o *AcquireMarginAccountSettingsResponse) HasDayRate() bool`

HasDayRate returns a boolean if a field has been set.

### GetLeverage

`func (o *AcquireMarginAccountSettingsResponse) GetLeverage() int64`

GetLeverage returns the Leverage field if non-nil, zero value otherwise.

### GetLeverageOk

`func (o *AcquireMarginAccountSettingsResponse) GetLeverageOk() (*int64, bool)`

GetLeverageOk returns a tuple with the Leverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeverage

`func (o *AcquireMarginAccountSettingsResponse) SetLeverage(v int64)`

SetLeverage sets Leverage field to given value.

### HasLeverage

`func (o *AcquireMarginAccountSettingsResponse) HasLeverage() bool`

HasLeverage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


