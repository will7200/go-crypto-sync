# InquireDepositListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActualAmount** | Pointer to **map[string]interface{}** | actual deposit amount | [optional] 
**ActualAmountDisplay** | Pointer to **map[string]interface{}** | display of actual deposit amount | [optional] 
**AddExplorer** | Pointer to **map[string]interface{}** | Depositor | [optional] 
**Amount** | Pointer to **map[string]interface{}** | Amount | [optional] 
**AmountDisplay** | Pointer to **map[string]interface{}** | Amount displayed | [optional] 
**CoinAddress** | Pointer to **map[string]interface{}** | Deposit add | [optional] 
**CoinAddressDisplay** | Pointer to **map[string]interface{}** | Deposit add displayed | [optional] 
**CoinDepositId** | Pointer to **map[string]interface{}** | Deposit ID | [optional] 
**CoinType** | Pointer to **map[string]interface{}** | Coin type | [optional] 
**Confirmations** | Pointer to **map[string]interface{}** | confirmations | [optional] 
**CreateTime** | Pointer to **map[string]interface{}** | create time | [optional] 
**Explorer** | Pointer to **map[string]interface{}** | explorer | [optional] 
**Remark** | Pointer to **map[string]interface{}** | remark | [optional] 
**Status** | Pointer to **map[string]interface{}** | processing，confirming，cancel，finish | [optional] 
**StatusDisplay** | Pointer to **map[string]interface{}** | Status displayed | [optional] 
**TransferMethod** | Pointer to **map[string]interface{}** | transfer method | [optional] 
**TxId** | Pointer to **map[string]interface{}** | tx id | [optional] 
**TxIdDisplay** | Pointer to **map[string]interface{}** | tx id display | [optional] 

## Methods

### NewInquireDepositListResponse

`func NewInquireDepositListResponse() *InquireDepositListResponse`

NewInquireDepositListResponse instantiates a new InquireDepositListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInquireDepositListResponseWithDefaults

`func NewInquireDepositListResponseWithDefaults() *InquireDepositListResponse`

NewInquireDepositListResponseWithDefaults instantiates a new InquireDepositListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActualAmount

`func (o *InquireDepositListResponse) GetActualAmount() map[string]interface{}`

GetActualAmount returns the ActualAmount field if non-nil, zero value otherwise.

### GetActualAmountOk

`func (o *InquireDepositListResponse) GetActualAmountOk() (*map[string]interface{}, bool)`

GetActualAmountOk returns a tuple with the ActualAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualAmount

`func (o *InquireDepositListResponse) SetActualAmount(v map[string]interface{})`

SetActualAmount sets ActualAmount field to given value.

### HasActualAmount

`func (o *InquireDepositListResponse) HasActualAmount() bool`

HasActualAmount returns a boolean if a field has been set.

### GetActualAmountDisplay

`func (o *InquireDepositListResponse) GetActualAmountDisplay() map[string]interface{}`

GetActualAmountDisplay returns the ActualAmountDisplay field if non-nil, zero value otherwise.

### GetActualAmountDisplayOk

`func (o *InquireDepositListResponse) GetActualAmountDisplayOk() (*map[string]interface{}, bool)`

GetActualAmountDisplayOk returns a tuple with the ActualAmountDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualAmountDisplay

`func (o *InquireDepositListResponse) SetActualAmountDisplay(v map[string]interface{})`

SetActualAmountDisplay sets ActualAmountDisplay field to given value.

### HasActualAmountDisplay

`func (o *InquireDepositListResponse) HasActualAmountDisplay() bool`

HasActualAmountDisplay returns a boolean if a field has been set.

### GetAddExplorer

`func (o *InquireDepositListResponse) GetAddExplorer() map[string]interface{}`

GetAddExplorer returns the AddExplorer field if non-nil, zero value otherwise.

### GetAddExplorerOk

`func (o *InquireDepositListResponse) GetAddExplorerOk() (*map[string]interface{}, bool)`

GetAddExplorerOk returns a tuple with the AddExplorer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddExplorer

`func (o *InquireDepositListResponse) SetAddExplorer(v map[string]interface{})`

SetAddExplorer sets AddExplorer field to given value.

### HasAddExplorer

`func (o *InquireDepositListResponse) HasAddExplorer() bool`

HasAddExplorer returns a boolean if a field has been set.

### GetAmount

`func (o *InquireDepositListResponse) GetAmount() map[string]interface{}`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *InquireDepositListResponse) GetAmountOk() (*map[string]interface{}, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *InquireDepositListResponse) SetAmount(v map[string]interface{})`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *InquireDepositListResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAmountDisplay

`func (o *InquireDepositListResponse) GetAmountDisplay() map[string]interface{}`

GetAmountDisplay returns the AmountDisplay field if non-nil, zero value otherwise.

### GetAmountDisplayOk

`func (o *InquireDepositListResponse) GetAmountDisplayOk() (*map[string]interface{}, bool)`

GetAmountDisplayOk returns a tuple with the AmountDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountDisplay

`func (o *InquireDepositListResponse) SetAmountDisplay(v map[string]interface{})`

SetAmountDisplay sets AmountDisplay field to given value.

### HasAmountDisplay

`func (o *InquireDepositListResponse) HasAmountDisplay() bool`

HasAmountDisplay returns a boolean if a field has been set.

### GetCoinAddress

`func (o *InquireDepositListResponse) GetCoinAddress() map[string]interface{}`

GetCoinAddress returns the CoinAddress field if non-nil, zero value otherwise.

### GetCoinAddressOk

`func (o *InquireDepositListResponse) GetCoinAddressOk() (*map[string]interface{}, bool)`

GetCoinAddressOk returns a tuple with the CoinAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoinAddress

`func (o *InquireDepositListResponse) SetCoinAddress(v map[string]interface{})`

SetCoinAddress sets CoinAddress field to given value.

### HasCoinAddress

`func (o *InquireDepositListResponse) HasCoinAddress() bool`

HasCoinAddress returns a boolean if a field has been set.

### GetCoinAddressDisplay

`func (o *InquireDepositListResponse) GetCoinAddressDisplay() map[string]interface{}`

GetCoinAddressDisplay returns the CoinAddressDisplay field if non-nil, zero value otherwise.

### GetCoinAddressDisplayOk

`func (o *InquireDepositListResponse) GetCoinAddressDisplayOk() (*map[string]interface{}, bool)`

GetCoinAddressDisplayOk returns a tuple with the CoinAddressDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoinAddressDisplay

`func (o *InquireDepositListResponse) SetCoinAddressDisplay(v map[string]interface{})`

SetCoinAddressDisplay sets CoinAddressDisplay field to given value.

### HasCoinAddressDisplay

`func (o *InquireDepositListResponse) HasCoinAddressDisplay() bool`

HasCoinAddressDisplay returns a boolean if a field has been set.

### GetCoinDepositId

`func (o *InquireDepositListResponse) GetCoinDepositId() map[string]interface{}`

GetCoinDepositId returns the CoinDepositId field if non-nil, zero value otherwise.

### GetCoinDepositIdOk

`func (o *InquireDepositListResponse) GetCoinDepositIdOk() (*map[string]interface{}, bool)`

GetCoinDepositIdOk returns a tuple with the CoinDepositId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoinDepositId

`func (o *InquireDepositListResponse) SetCoinDepositId(v map[string]interface{})`

SetCoinDepositId sets CoinDepositId field to given value.

### HasCoinDepositId

`func (o *InquireDepositListResponse) HasCoinDepositId() bool`

HasCoinDepositId returns a boolean if a field has been set.

### GetCoinType

`func (o *InquireDepositListResponse) GetCoinType() map[string]interface{}`

GetCoinType returns the CoinType field if non-nil, zero value otherwise.

### GetCoinTypeOk

`func (o *InquireDepositListResponse) GetCoinTypeOk() (*map[string]interface{}, bool)`

GetCoinTypeOk returns a tuple with the CoinType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoinType

`func (o *InquireDepositListResponse) SetCoinType(v map[string]interface{})`

SetCoinType sets CoinType field to given value.

### HasCoinType

`func (o *InquireDepositListResponse) HasCoinType() bool`

HasCoinType returns a boolean if a field has been set.

### GetConfirmations

`func (o *InquireDepositListResponse) GetConfirmations() map[string]interface{}`

GetConfirmations returns the Confirmations field if non-nil, zero value otherwise.

### GetConfirmationsOk

`func (o *InquireDepositListResponse) GetConfirmationsOk() (*map[string]interface{}, bool)`

GetConfirmationsOk returns a tuple with the Confirmations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmations

`func (o *InquireDepositListResponse) SetConfirmations(v map[string]interface{})`

SetConfirmations sets Confirmations field to given value.

### HasConfirmations

`func (o *InquireDepositListResponse) HasConfirmations() bool`

HasConfirmations returns a boolean if a field has been set.

### GetCreateTime

`func (o *InquireDepositListResponse) GetCreateTime() map[string]interface{}`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *InquireDepositListResponse) GetCreateTimeOk() (*map[string]interface{}, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *InquireDepositListResponse) SetCreateTime(v map[string]interface{})`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *InquireDepositListResponse) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetExplorer

`func (o *InquireDepositListResponse) GetExplorer() map[string]interface{}`

GetExplorer returns the Explorer field if non-nil, zero value otherwise.

### GetExplorerOk

`func (o *InquireDepositListResponse) GetExplorerOk() (*map[string]interface{}, bool)`

GetExplorerOk returns a tuple with the Explorer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplorer

`func (o *InquireDepositListResponse) SetExplorer(v map[string]interface{})`

SetExplorer sets Explorer field to given value.

### HasExplorer

`func (o *InquireDepositListResponse) HasExplorer() bool`

HasExplorer returns a boolean if a field has been set.

### GetRemark

`func (o *InquireDepositListResponse) GetRemark() map[string]interface{}`

GetRemark returns the Remark field if non-nil, zero value otherwise.

### GetRemarkOk

`func (o *InquireDepositListResponse) GetRemarkOk() (*map[string]interface{}, bool)`

GetRemarkOk returns a tuple with the Remark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemark

`func (o *InquireDepositListResponse) SetRemark(v map[string]interface{})`

SetRemark sets Remark field to given value.

### HasRemark

`func (o *InquireDepositListResponse) HasRemark() bool`

HasRemark returns a boolean if a field has been set.

### GetStatus

`func (o *InquireDepositListResponse) GetStatus() map[string]interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InquireDepositListResponse) GetStatusOk() (*map[string]interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InquireDepositListResponse) SetStatus(v map[string]interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InquireDepositListResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDisplay

`func (o *InquireDepositListResponse) GetStatusDisplay() map[string]interface{}`

GetStatusDisplay returns the StatusDisplay field if non-nil, zero value otherwise.

### GetStatusDisplayOk

`func (o *InquireDepositListResponse) GetStatusDisplayOk() (*map[string]interface{}, bool)`

GetStatusDisplayOk returns a tuple with the StatusDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDisplay

`func (o *InquireDepositListResponse) SetStatusDisplay(v map[string]interface{})`

SetStatusDisplay sets StatusDisplay field to given value.

### HasStatusDisplay

`func (o *InquireDepositListResponse) HasStatusDisplay() bool`

HasStatusDisplay returns a boolean if a field has been set.

### GetTransferMethod

`func (o *InquireDepositListResponse) GetTransferMethod() map[string]interface{}`

GetTransferMethod returns the TransferMethod field if non-nil, zero value otherwise.

### GetTransferMethodOk

`func (o *InquireDepositListResponse) GetTransferMethodOk() (*map[string]interface{}, bool)`

GetTransferMethodOk returns a tuple with the TransferMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferMethod

`func (o *InquireDepositListResponse) SetTransferMethod(v map[string]interface{})`

SetTransferMethod sets TransferMethod field to given value.

### HasTransferMethod

`func (o *InquireDepositListResponse) HasTransferMethod() bool`

HasTransferMethod returns a boolean if a field has been set.

### GetTxId

`func (o *InquireDepositListResponse) GetTxId() map[string]interface{}`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *InquireDepositListResponse) GetTxIdOk() (*map[string]interface{}, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *InquireDepositListResponse) SetTxId(v map[string]interface{})`

SetTxId sets TxId field to given value.

### HasTxId

`func (o *InquireDepositListResponse) HasTxId() bool`

HasTxId returns a boolean if a field has been set.

### GetTxIdDisplay

`func (o *InquireDepositListResponse) GetTxIdDisplay() map[string]interface{}`

GetTxIdDisplay returns the TxIdDisplay field if non-nil, zero value otherwise.

### GetTxIdDisplayOk

`func (o *InquireDepositListResponse) GetTxIdDisplayOk() (*map[string]interface{}, bool)`

GetTxIdDisplayOk returns a tuple with the TxIdDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxIdDisplay

`func (o *InquireDepositListResponse) SetTxIdDisplay(v map[string]interface{})`

SetTxIdDisplay sets TxIdDisplay field to given value.

### HasTxIdDisplay

`func (o *InquireDepositListResponse) HasTxIdDisplay() bool`

HasTxIdDisplay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


