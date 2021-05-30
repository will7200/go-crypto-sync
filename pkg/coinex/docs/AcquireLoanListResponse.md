# AcquireLoanListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoanId** | Pointer to **int64** | loan record ID | [optional] 
**CreateTime** | Pointer to **int64** | create timestamp | [optional] 
**Market** | Pointer to **string** | See&lt;API invocation description market&gt; | [optional] 
**CoinType** | Pointer to **string** | coin token name | [optional] 
**LoanAmount** | Pointer to **string** | borrow amount | [optional] 
**UnflatAmount** | Pointer to **string** | amount and interest need to repay | [optional] 
**ExpireTime** | Pointer to **int64** | expire timestamp | [optional] 
**Renew** | Pointer to **int64** | 0: close; &lt;br&gt;1: open; | [optional] 
**DayRate** | Pointer to **string** | daily interest rate | [optional] 
**Status** | Pointer to **string** | pass: in loan;&lt;br&gt; burst: bankrupt;&lt;br&gt; arrears: in debt;&lt;br&gt; finish: paid off; | [optional] 

## Methods

### NewAcquireLoanListResponse

`func NewAcquireLoanListResponse() *AcquireLoanListResponse`

NewAcquireLoanListResponse instantiates a new AcquireLoanListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcquireLoanListResponseWithDefaults

`func NewAcquireLoanListResponseWithDefaults() *AcquireLoanListResponse`

NewAcquireLoanListResponseWithDefaults instantiates a new AcquireLoanListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoanId

`func (o *AcquireLoanListResponse) GetLoanId() int64`

GetLoanId returns the LoanId field if non-nil, zero value otherwise.

### GetLoanIdOk

`func (o *AcquireLoanListResponse) GetLoanIdOk() (*int64, bool)`

GetLoanIdOk returns a tuple with the LoanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoanId

`func (o *AcquireLoanListResponse) SetLoanId(v int64)`

SetLoanId sets LoanId field to given value.

### HasLoanId

`func (o *AcquireLoanListResponse) HasLoanId() bool`

HasLoanId returns a boolean if a field has been set.

### GetCreateTime

`func (o *AcquireLoanListResponse) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *AcquireLoanListResponse) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *AcquireLoanListResponse) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *AcquireLoanListResponse) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetMarket

`func (o *AcquireLoanListResponse) GetMarket() string`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *AcquireLoanListResponse) GetMarketOk() (*string, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *AcquireLoanListResponse) SetMarket(v string)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *AcquireLoanListResponse) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetCoinType

`func (o *AcquireLoanListResponse) GetCoinType() string`

GetCoinType returns the CoinType field if non-nil, zero value otherwise.

### GetCoinTypeOk

`func (o *AcquireLoanListResponse) GetCoinTypeOk() (*string, bool)`

GetCoinTypeOk returns a tuple with the CoinType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoinType

`func (o *AcquireLoanListResponse) SetCoinType(v string)`

SetCoinType sets CoinType field to given value.

### HasCoinType

`func (o *AcquireLoanListResponse) HasCoinType() bool`

HasCoinType returns a boolean if a field has been set.

### GetLoanAmount

`func (o *AcquireLoanListResponse) GetLoanAmount() string`

GetLoanAmount returns the LoanAmount field if non-nil, zero value otherwise.

### GetLoanAmountOk

`func (o *AcquireLoanListResponse) GetLoanAmountOk() (*string, bool)`

GetLoanAmountOk returns a tuple with the LoanAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoanAmount

`func (o *AcquireLoanListResponse) SetLoanAmount(v string)`

SetLoanAmount sets LoanAmount field to given value.

### HasLoanAmount

`func (o *AcquireLoanListResponse) HasLoanAmount() bool`

HasLoanAmount returns a boolean if a field has been set.

### GetUnflatAmount

`func (o *AcquireLoanListResponse) GetUnflatAmount() string`

GetUnflatAmount returns the UnflatAmount field if non-nil, zero value otherwise.

### GetUnflatAmountOk

`func (o *AcquireLoanListResponse) GetUnflatAmountOk() (*string, bool)`

GetUnflatAmountOk returns a tuple with the UnflatAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnflatAmount

`func (o *AcquireLoanListResponse) SetUnflatAmount(v string)`

SetUnflatAmount sets UnflatAmount field to given value.

### HasUnflatAmount

`func (o *AcquireLoanListResponse) HasUnflatAmount() bool`

HasUnflatAmount returns a boolean if a field has been set.

### GetExpireTime

`func (o *AcquireLoanListResponse) GetExpireTime() int64`

GetExpireTime returns the ExpireTime field if non-nil, zero value otherwise.

### GetExpireTimeOk

`func (o *AcquireLoanListResponse) GetExpireTimeOk() (*int64, bool)`

GetExpireTimeOk returns a tuple with the ExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireTime

`func (o *AcquireLoanListResponse) SetExpireTime(v int64)`

SetExpireTime sets ExpireTime field to given value.

### HasExpireTime

`func (o *AcquireLoanListResponse) HasExpireTime() bool`

HasExpireTime returns a boolean if a field has been set.

### GetRenew

`func (o *AcquireLoanListResponse) GetRenew() int64`

GetRenew returns the Renew field if non-nil, zero value otherwise.

### GetRenewOk

`func (o *AcquireLoanListResponse) GetRenewOk() (*int64, bool)`

GetRenewOk returns a tuple with the Renew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenew

`func (o *AcquireLoanListResponse) SetRenew(v int64)`

SetRenew sets Renew field to given value.

### HasRenew

`func (o *AcquireLoanListResponse) HasRenew() bool`

HasRenew returns a boolean if a field has been set.

### GetDayRate

`func (o *AcquireLoanListResponse) GetDayRate() string`

GetDayRate returns the DayRate field if non-nil, zero value otherwise.

### GetDayRateOk

`func (o *AcquireLoanListResponse) GetDayRateOk() (*string, bool)`

GetDayRateOk returns a tuple with the DayRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayRate

`func (o *AcquireLoanListResponse) SetDayRate(v string)`

SetDayRate sets DayRate field to given value.

### HasDayRate

`func (o *AcquireLoanListResponse) HasDayRate() bool`

HasDayRate returns a boolean if a field has been set.

### GetStatus

`func (o *AcquireLoanListResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AcquireLoanListResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AcquireLoanListResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AcquireLoanListResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


