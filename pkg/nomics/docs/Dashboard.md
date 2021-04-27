# Dashboard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** | Currency ID | [optional] 
**DayOpen** | Pointer to **string** | Price one day ago | [optional] 
**DayVolume** | Pointer to **string** | Volume over the past day | [optional] 
**DayOpenVolume** | Pointer to **string** | Volume over the previous day | [optional] 
**WeekOpen** | Pointer to **string** | Price one week ago | [optional] 
**WeekVolume** | Pointer to **string** | Volume over the past week | [optional] 
**WeekOpenVolume** | Pointer to **string** | Volume over the previous week | [optional] 
**MonthOpen** | Pointer to **string** | Price one month ago | [optional] 
**MonthVolume** | Pointer to **string** | Volume over the past 30 days | [optional] 
**MonthOpenVolume** | Pointer to **string** | Volume over the previous 30 days | [optional] 
**YearOpen** | Pointer to **string** | Price one year ago | [optional] 
**YearVolume** | Pointer to **string** | Volume over the past year | [optional] 
**YearOpenVolume** | Pointer to **string** | Volume over the previous year | [optional] 
**Close** | Pointer to **string** | Current price | [optional] 
**High** | Pointer to **string** | Highest price | [optional] 
**HighTimestamp** | Pointer to **string** | Timestamp of the highest price in RFC3339 | [optional] 
**HighExchange** | Pointer to **string** | Exchange on which the highest price occurred | [optional] 
**HighQuoteCurrency** | Pointer to **string** | Quote currency against which highest price was recorded | [optional] 
**AvailableSupply** | Pointer to **string** | Circulating supply | [optional] 
**MaxSupply** | Pointer to **string** | Maximum supply | [optional] 

## Methods

### NewDashboard

`func NewDashboard() *Dashboard`

NewDashboard instantiates a new Dashboard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardWithDefaults

`func NewDashboardWithDefaults() *Dashboard`

NewDashboardWithDefaults instantiates a new Dashboard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *Dashboard) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Dashboard) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Dashboard) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Dashboard) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDayOpen

`func (o *Dashboard) GetDayOpen() string`

GetDayOpen returns the DayOpen field if non-nil, zero value otherwise.

### GetDayOpenOk

`func (o *Dashboard) GetDayOpenOk() (*string, bool)`

GetDayOpenOk returns a tuple with the DayOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOpen

`func (o *Dashboard) SetDayOpen(v string)`

SetDayOpen sets DayOpen field to given value.

### HasDayOpen

`func (o *Dashboard) HasDayOpen() bool`

HasDayOpen returns a boolean if a field has been set.

### GetDayVolume

`func (o *Dashboard) GetDayVolume() string`

GetDayVolume returns the DayVolume field if non-nil, zero value otherwise.

### GetDayVolumeOk

`func (o *Dashboard) GetDayVolumeOk() (*string, bool)`

GetDayVolumeOk returns a tuple with the DayVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayVolume

`func (o *Dashboard) SetDayVolume(v string)`

SetDayVolume sets DayVolume field to given value.

### HasDayVolume

`func (o *Dashboard) HasDayVolume() bool`

HasDayVolume returns a boolean if a field has been set.

### GetDayOpenVolume

`func (o *Dashboard) GetDayOpenVolume() string`

GetDayOpenVolume returns the DayOpenVolume field if non-nil, zero value otherwise.

### GetDayOpenVolumeOk

`func (o *Dashboard) GetDayOpenVolumeOk() (*string, bool)`

GetDayOpenVolumeOk returns a tuple with the DayOpenVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOpenVolume

`func (o *Dashboard) SetDayOpenVolume(v string)`

SetDayOpenVolume sets DayOpenVolume field to given value.

### HasDayOpenVolume

`func (o *Dashboard) HasDayOpenVolume() bool`

HasDayOpenVolume returns a boolean if a field has been set.

### GetWeekOpen

`func (o *Dashboard) GetWeekOpen() string`

GetWeekOpen returns the WeekOpen field if non-nil, zero value otherwise.

### GetWeekOpenOk

`func (o *Dashboard) GetWeekOpenOk() (*string, bool)`

GetWeekOpenOk returns a tuple with the WeekOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekOpen

`func (o *Dashboard) SetWeekOpen(v string)`

SetWeekOpen sets WeekOpen field to given value.

### HasWeekOpen

`func (o *Dashboard) HasWeekOpen() bool`

HasWeekOpen returns a boolean if a field has been set.

### GetWeekVolume

`func (o *Dashboard) GetWeekVolume() string`

GetWeekVolume returns the WeekVolume field if non-nil, zero value otherwise.

### GetWeekVolumeOk

`func (o *Dashboard) GetWeekVolumeOk() (*string, bool)`

GetWeekVolumeOk returns a tuple with the WeekVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekVolume

`func (o *Dashboard) SetWeekVolume(v string)`

SetWeekVolume sets WeekVolume field to given value.

### HasWeekVolume

`func (o *Dashboard) HasWeekVolume() bool`

HasWeekVolume returns a boolean if a field has been set.

### GetWeekOpenVolume

`func (o *Dashboard) GetWeekOpenVolume() string`

GetWeekOpenVolume returns the WeekOpenVolume field if non-nil, zero value otherwise.

### GetWeekOpenVolumeOk

`func (o *Dashboard) GetWeekOpenVolumeOk() (*string, bool)`

GetWeekOpenVolumeOk returns a tuple with the WeekOpenVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekOpenVolume

`func (o *Dashboard) SetWeekOpenVolume(v string)`

SetWeekOpenVolume sets WeekOpenVolume field to given value.

### HasWeekOpenVolume

`func (o *Dashboard) HasWeekOpenVolume() bool`

HasWeekOpenVolume returns a boolean if a field has been set.

### GetMonthOpen

`func (o *Dashboard) GetMonthOpen() string`

GetMonthOpen returns the MonthOpen field if non-nil, zero value otherwise.

### GetMonthOpenOk

`func (o *Dashboard) GetMonthOpenOk() (*string, bool)`

GetMonthOpenOk returns a tuple with the MonthOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthOpen

`func (o *Dashboard) SetMonthOpen(v string)`

SetMonthOpen sets MonthOpen field to given value.

### HasMonthOpen

`func (o *Dashboard) HasMonthOpen() bool`

HasMonthOpen returns a boolean if a field has been set.

### GetMonthVolume

`func (o *Dashboard) GetMonthVolume() string`

GetMonthVolume returns the MonthVolume field if non-nil, zero value otherwise.

### GetMonthVolumeOk

`func (o *Dashboard) GetMonthVolumeOk() (*string, bool)`

GetMonthVolumeOk returns a tuple with the MonthVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthVolume

`func (o *Dashboard) SetMonthVolume(v string)`

SetMonthVolume sets MonthVolume field to given value.

### HasMonthVolume

`func (o *Dashboard) HasMonthVolume() bool`

HasMonthVolume returns a boolean if a field has been set.

### GetMonthOpenVolume

`func (o *Dashboard) GetMonthOpenVolume() string`

GetMonthOpenVolume returns the MonthOpenVolume field if non-nil, zero value otherwise.

### GetMonthOpenVolumeOk

`func (o *Dashboard) GetMonthOpenVolumeOk() (*string, bool)`

GetMonthOpenVolumeOk returns a tuple with the MonthOpenVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthOpenVolume

`func (o *Dashboard) SetMonthOpenVolume(v string)`

SetMonthOpenVolume sets MonthOpenVolume field to given value.

### HasMonthOpenVolume

`func (o *Dashboard) HasMonthOpenVolume() bool`

HasMonthOpenVolume returns a boolean if a field has been set.

### GetYearOpen

`func (o *Dashboard) GetYearOpen() string`

GetYearOpen returns the YearOpen field if non-nil, zero value otherwise.

### GetYearOpenOk

`func (o *Dashboard) GetYearOpenOk() (*string, bool)`

GetYearOpenOk returns a tuple with the YearOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearOpen

`func (o *Dashboard) SetYearOpen(v string)`

SetYearOpen sets YearOpen field to given value.

### HasYearOpen

`func (o *Dashboard) HasYearOpen() bool`

HasYearOpen returns a boolean if a field has been set.

### GetYearVolume

`func (o *Dashboard) GetYearVolume() string`

GetYearVolume returns the YearVolume field if non-nil, zero value otherwise.

### GetYearVolumeOk

`func (o *Dashboard) GetYearVolumeOk() (*string, bool)`

GetYearVolumeOk returns a tuple with the YearVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearVolume

`func (o *Dashboard) SetYearVolume(v string)`

SetYearVolume sets YearVolume field to given value.

### HasYearVolume

`func (o *Dashboard) HasYearVolume() bool`

HasYearVolume returns a boolean if a field has been set.

### GetYearOpenVolume

`func (o *Dashboard) GetYearOpenVolume() string`

GetYearOpenVolume returns the YearOpenVolume field if non-nil, zero value otherwise.

### GetYearOpenVolumeOk

`func (o *Dashboard) GetYearOpenVolumeOk() (*string, bool)`

GetYearOpenVolumeOk returns a tuple with the YearOpenVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearOpenVolume

`func (o *Dashboard) SetYearOpenVolume(v string)`

SetYearOpenVolume sets YearOpenVolume field to given value.

### HasYearOpenVolume

`func (o *Dashboard) HasYearOpenVolume() bool`

HasYearOpenVolume returns a boolean if a field has been set.

### GetClose

`func (o *Dashboard) GetClose() string`

GetClose returns the Close field if non-nil, zero value otherwise.

### GetCloseOk

`func (o *Dashboard) GetCloseOk() (*string, bool)`

GetCloseOk returns a tuple with the Close field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClose

`func (o *Dashboard) SetClose(v string)`

SetClose sets Close field to given value.

### HasClose

`func (o *Dashboard) HasClose() bool`

HasClose returns a boolean if a field has been set.

### GetHigh

`func (o *Dashboard) GetHigh() string`

GetHigh returns the High field if non-nil, zero value otherwise.

### GetHighOk

`func (o *Dashboard) GetHighOk() (*string, bool)`

GetHighOk returns a tuple with the High field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigh

`func (o *Dashboard) SetHigh(v string)`

SetHigh sets High field to given value.

### HasHigh

`func (o *Dashboard) HasHigh() bool`

HasHigh returns a boolean if a field has been set.

### GetHighTimestamp

`func (o *Dashboard) GetHighTimestamp() string`

GetHighTimestamp returns the HighTimestamp field if non-nil, zero value otherwise.

### GetHighTimestampOk

`func (o *Dashboard) GetHighTimestampOk() (*string, bool)`

GetHighTimestampOk returns a tuple with the HighTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighTimestamp

`func (o *Dashboard) SetHighTimestamp(v string)`

SetHighTimestamp sets HighTimestamp field to given value.

### HasHighTimestamp

`func (o *Dashboard) HasHighTimestamp() bool`

HasHighTimestamp returns a boolean if a field has been set.

### GetHighExchange

`func (o *Dashboard) GetHighExchange() string`

GetHighExchange returns the HighExchange field if non-nil, zero value otherwise.

### GetHighExchangeOk

`func (o *Dashboard) GetHighExchangeOk() (*string, bool)`

GetHighExchangeOk returns a tuple with the HighExchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighExchange

`func (o *Dashboard) SetHighExchange(v string)`

SetHighExchange sets HighExchange field to given value.

### HasHighExchange

`func (o *Dashboard) HasHighExchange() bool`

HasHighExchange returns a boolean if a field has been set.

### GetHighQuoteCurrency

`func (o *Dashboard) GetHighQuoteCurrency() string`

GetHighQuoteCurrency returns the HighQuoteCurrency field if non-nil, zero value otherwise.

### GetHighQuoteCurrencyOk

`func (o *Dashboard) GetHighQuoteCurrencyOk() (*string, bool)`

GetHighQuoteCurrencyOk returns a tuple with the HighQuoteCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighQuoteCurrency

`func (o *Dashboard) SetHighQuoteCurrency(v string)`

SetHighQuoteCurrency sets HighQuoteCurrency field to given value.

### HasHighQuoteCurrency

`func (o *Dashboard) HasHighQuoteCurrency() bool`

HasHighQuoteCurrency returns a boolean if a field has been set.

### GetAvailableSupply

`func (o *Dashboard) GetAvailableSupply() string`

GetAvailableSupply returns the AvailableSupply field if non-nil, zero value otherwise.

### GetAvailableSupplyOk

`func (o *Dashboard) GetAvailableSupplyOk() (*string, bool)`

GetAvailableSupplyOk returns a tuple with the AvailableSupply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableSupply

`func (o *Dashboard) SetAvailableSupply(v string)`

SetAvailableSupply sets AvailableSupply field to given value.

### HasAvailableSupply

`func (o *Dashboard) HasAvailableSupply() bool`

HasAvailableSupply returns a boolean if a field has been set.

### GetMaxSupply

`func (o *Dashboard) GetMaxSupply() string`

GetMaxSupply returns the MaxSupply field if non-nil, zero value otherwise.

### GetMaxSupplyOk

`func (o *Dashboard) GetMaxSupplyOk() (*string, bool)`

GetMaxSupplyOk returns a tuple with the MaxSupply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSupply

`func (o *Dashboard) SetMaxSupply(v string)`

SetMaxSupply sets MaxSupply field to given value.

### HasMaxSupply

`func (o *Dashboard) HasMaxSupply() bool`

HasMaxSupply returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


