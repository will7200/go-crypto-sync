# AcquireAssetConfigResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanDeposit** | Pointer to **bool** | coin type deposit status | [optional] 
**CanWithdraw** | Pointer to **bool** | coin type withdraw status | [optional] 
**DepositLeastAmount** | Pointer to **string** | deposit least amount | [optional] 
**WithdrawLeastAmount** | Pointer to **string** | withdraw least amount | [optional] 
**WithdrawTxFee** | Pointer to **string** | withdraw tx fee | [optional] 

## Methods

### NewAcquireAssetConfigResponse

`func NewAcquireAssetConfigResponse() *AcquireAssetConfigResponse`

NewAcquireAssetConfigResponse instantiates a new AcquireAssetConfigResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcquireAssetConfigResponseWithDefaults

`func NewAcquireAssetConfigResponseWithDefaults() *AcquireAssetConfigResponse`

NewAcquireAssetConfigResponseWithDefaults instantiates a new AcquireAssetConfigResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanDeposit

`func (o *AcquireAssetConfigResponse) GetCanDeposit() bool`

GetCanDeposit returns the CanDeposit field if non-nil, zero value otherwise.

### GetCanDepositOk

`func (o *AcquireAssetConfigResponse) GetCanDepositOk() (*bool, bool)`

GetCanDepositOk returns a tuple with the CanDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeposit

`func (o *AcquireAssetConfigResponse) SetCanDeposit(v bool)`

SetCanDeposit sets CanDeposit field to given value.

### HasCanDeposit

`func (o *AcquireAssetConfigResponse) HasCanDeposit() bool`

HasCanDeposit returns a boolean if a field has been set.

### GetCanWithdraw

`func (o *AcquireAssetConfigResponse) GetCanWithdraw() bool`

GetCanWithdraw returns the CanWithdraw field if non-nil, zero value otherwise.

### GetCanWithdrawOk

`func (o *AcquireAssetConfigResponse) GetCanWithdrawOk() (*bool, bool)`

GetCanWithdrawOk returns a tuple with the CanWithdraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanWithdraw

`func (o *AcquireAssetConfigResponse) SetCanWithdraw(v bool)`

SetCanWithdraw sets CanWithdraw field to given value.

### HasCanWithdraw

`func (o *AcquireAssetConfigResponse) HasCanWithdraw() bool`

HasCanWithdraw returns a boolean if a field has been set.

### GetDepositLeastAmount

`func (o *AcquireAssetConfigResponse) GetDepositLeastAmount() string`

GetDepositLeastAmount returns the DepositLeastAmount field if non-nil, zero value otherwise.

### GetDepositLeastAmountOk

`func (o *AcquireAssetConfigResponse) GetDepositLeastAmountOk() (*string, bool)`

GetDepositLeastAmountOk returns a tuple with the DepositLeastAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositLeastAmount

`func (o *AcquireAssetConfigResponse) SetDepositLeastAmount(v string)`

SetDepositLeastAmount sets DepositLeastAmount field to given value.

### HasDepositLeastAmount

`func (o *AcquireAssetConfigResponse) HasDepositLeastAmount() bool`

HasDepositLeastAmount returns a boolean if a field has been set.

### GetWithdrawLeastAmount

`func (o *AcquireAssetConfigResponse) GetWithdrawLeastAmount() string`

GetWithdrawLeastAmount returns the WithdrawLeastAmount field if non-nil, zero value otherwise.

### GetWithdrawLeastAmountOk

`func (o *AcquireAssetConfigResponse) GetWithdrawLeastAmountOk() (*string, bool)`

GetWithdrawLeastAmountOk returns a tuple with the WithdrawLeastAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawLeastAmount

`func (o *AcquireAssetConfigResponse) SetWithdrawLeastAmount(v string)`

SetWithdrawLeastAmount sets WithdrawLeastAmount field to given value.

### HasWithdrawLeastAmount

`func (o *AcquireAssetConfigResponse) HasWithdrawLeastAmount() bool`

HasWithdrawLeastAmount returns a boolean if a field has been set.

### GetWithdrawTxFee

`func (o *AcquireAssetConfigResponse) GetWithdrawTxFee() string`

GetWithdrawTxFee returns the WithdrawTxFee field if non-nil, zero value otherwise.

### GetWithdrawTxFeeOk

`func (o *AcquireAssetConfigResponse) GetWithdrawTxFeeOk() (*string, bool)`

GetWithdrawTxFeeOk returns a tuple with the WithdrawTxFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawTxFee

`func (o *AcquireAssetConfigResponse) SetWithdrawTxFee(v string)`

SetWithdrawTxFee sets WithdrawTxFee field to given value.

### HasWithdrawTxFee

`func (o *AcquireAssetConfigResponse) HasWithdrawTxFee() bool`

HasWithdrawTxFee returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


