# InquireWithdrawalListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoinWithdrawId** | Pointer to **map[string]interface{}** | coin withdrawal id | [optional] 
**CreateTime** | Pointer to **map[string]interface{}** | create time | [optional] 
**Amount** | Pointer to **map[string]interface{}** | withdrawal amount | [optional] 
**ActualAmount** | Pointer to **map[string]interface{}** | actual withdrawal amount | [optional] 
**TxId** | Pointer to **map[string]interface{}** | tx id | [optional] 
**CoinAddress** | Pointer to **map[string]interface{}** | withdrawal address | [optional] 
**TxFee** | Pointer to **map[string]interface{}** | tx fee | [optional] 
**Confirmations** | Pointer to **map[string]interface{}** | confirmations | [optional] 
**CoinType** | Pointer to **map[string]interface{}** | coin type | [optional] 
**Status** | Pointer to **map[string]interface{}** | audit;&lt;br&gt;pass;&lt;br&gt;processing;&lt;br&gt;confirming;&lt;br&gt;not\\_pass;&lt;br&gt;cancel;&lt;br&gt;finish;&lt;br&gt;fail; | [optional] 

## Methods

### NewInquireWithdrawalListResponse

`func NewInquireWithdrawalListResponse() *InquireWithdrawalListResponse`

NewInquireWithdrawalListResponse instantiates a new InquireWithdrawalListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInquireWithdrawalListResponseWithDefaults

`func NewInquireWithdrawalListResponseWithDefaults() *InquireWithdrawalListResponse`

NewInquireWithdrawalListResponseWithDefaults instantiates a new InquireWithdrawalListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoinWithdrawId

`func (o *InquireWithdrawalListResponse) GetCoinWithdrawId() map[string]interface{}`

GetCoinWithdrawId returns the CoinWithdrawId field if non-nil, zero value otherwise.

### GetCoinWithdrawIdOk

`func (o *InquireWithdrawalListResponse) GetCoinWithdrawIdOk() (*map[string]interface{}, bool)`

GetCoinWithdrawIdOk returns a tuple with the CoinWithdrawId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoinWithdrawId

`func (o *InquireWithdrawalListResponse) SetCoinWithdrawId(v map[string]interface{})`

SetCoinWithdrawId sets CoinWithdrawId field to given value.

### HasCoinWithdrawId

`func (o *InquireWithdrawalListResponse) HasCoinWithdrawId() bool`

HasCoinWithdrawId returns a boolean if a field has been set.

### GetCreateTime

`func (o *InquireWithdrawalListResponse) GetCreateTime() map[string]interface{}`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *InquireWithdrawalListResponse) GetCreateTimeOk() (*map[string]interface{}, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *InquireWithdrawalListResponse) SetCreateTime(v map[string]interface{})`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *InquireWithdrawalListResponse) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetAmount

`func (o *InquireWithdrawalListResponse) GetAmount() map[string]interface{}`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *InquireWithdrawalListResponse) GetAmountOk() (*map[string]interface{}, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *InquireWithdrawalListResponse) SetAmount(v map[string]interface{})`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *InquireWithdrawalListResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetActualAmount

`func (o *InquireWithdrawalListResponse) GetActualAmount() map[string]interface{}`

GetActualAmount returns the ActualAmount field if non-nil, zero value otherwise.

### GetActualAmountOk

`func (o *InquireWithdrawalListResponse) GetActualAmountOk() (*map[string]interface{}, bool)`

GetActualAmountOk returns a tuple with the ActualAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualAmount

`func (o *InquireWithdrawalListResponse) SetActualAmount(v map[string]interface{})`

SetActualAmount sets ActualAmount field to given value.

### HasActualAmount

`func (o *InquireWithdrawalListResponse) HasActualAmount() bool`

HasActualAmount returns a boolean if a field has been set.

### GetTxId

`func (o *InquireWithdrawalListResponse) GetTxId() map[string]interface{}`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *InquireWithdrawalListResponse) GetTxIdOk() (*map[string]interface{}, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *InquireWithdrawalListResponse) SetTxId(v map[string]interface{})`

SetTxId sets TxId field to given value.

### HasTxId

`func (o *InquireWithdrawalListResponse) HasTxId() bool`

HasTxId returns a boolean if a field has been set.

### GetCoinAddress

`func (o *InquireWithdrawalListResponse) GetCoinAddress() map[string]interface{}`

GetCoinAddress returns the CoinAddress field if non-nil, zero value otherwise.

### GetCoinAddressOk

`func (o *InquireWithdrawalListResponse) GetCoinAddressOk() (*map[string]interface{}, bool)`

GetCoinAddressOk returns a tuple with the CoinAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoinAddress

`func (o *InquireWithdrawalListResponse) SetCoinAddress(v map[string]interface{})`

SetCoinAddress sets CoinAddress field to given value.

### HasCoinAddress

`func (o *InquireWithdrawalListResponse) HasCoinAddress() bool`

HasCoinAddress returns a boolean if a field has been set.

### GetTxFee

`func (o *InquireWithdrawalListResponse) GetTxFee() map[string]interface{}`

GetTxFee returns the TxFee field if non-nil, zero value otherwise.

### GetTxFeeOk

`func (o *InquireWithdrawalListResponse) GetTxFeeOk() (*map[string]interface{}, bool)`

GetTxFeeOk returns a tuple with the TxFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxFee

`func (o *InquireWithdrawalListResponse) SetTxFee(v map[string]interface{})`

SetTxFee sets TxFee field to given value.

### HasTxFee

`func (o *InquireWithdrawalListResponse) HasTxFee() bool`

HasTxFee returns a boolean if a field has been set.

### GetConfirmations

`func (o *InquireWithdrawalListResponse) GetConfirmations() map[string]interface{}`

GetConfirmations returns the Confirmations field if non-nil, zero value otherwise.

### GetConfirmationsOk

`func (o *InquireWithdrawalListResponse) GetConfirmationsOk() (*map[string]interface{}, bool)`

GetConfirmationsOk returns a tuple with the Confirmations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmations

`func (o *InquireWithdrawalListResponse) SetConfirmations(v map[string]interface{})`

SetConfirmations sets Confirmations field to given value.

### HasConfirmations

`func (o *InquireWithdrawalListResponse) HasConfirmations() bool`

HasConfirmations returns a boolean if a field has been set.

### GetCoinType

`func (o *InquireWithdrawalListResponse) GetCoinType() map[string]interface{}`

GetCoinType returns the CoinType field if non-nil, zero value otherwise.

### GetCoinTypeOk

`func (o *InquireWithdrawalListResponse) GetCoinTypeOk() (*map[string]interface{}, bool)`

GetCoinTypeOk returns a tuple with the CoinType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoinType

`func (o *InquireWithdrawalListResponse) SetCoinType(v map[string]interface{})`

SetCoinType sets CoinType field to given value.

### HasCoinType

`func (o *InquireWithdrawalListResponse) HasCoinType() bool`

HasCoinType returns a boolean if a field has been set.

### GetStatus

`func (o *InquireWithdrawalListResponse) GetStatus() map[string]interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InquireWithdrawalListResponse) GetStatusOk() (*map[string]interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InquireWithdrawalListResponse) SetStatus(v map[string]interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InquireWithdrawalListResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


