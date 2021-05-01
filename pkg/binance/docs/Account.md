# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MakerCommission** | Pointer to **int32** |  | [optional] 
**TakerCommission** | Pointer to **int32** |  | [optional] 
**BuyerCommission** | Pointer to **int32** |  | [optional] 
**SellerCommission** | Pointer to **int32** |  | [optional] 
**CanTrade** | Pointer to **bool** |  | [optional] 
**CanWithdraw** | Pointer to **bool** |  | [optional] 
**CanDeposit** | Pointer to **bool** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 
**AccountType** | Pointer to **string** |  | [optional] 
**Balances** | Pointer to [**[]AccountBalances**](AccountBalances.md) |  | [optional] 

## Methods

### NewAccount

`func NewAccount() *Account`

NewAccount instantiates a new Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithDefaults

`func NewAccountWithDefaults() *Account`

NewAccountWithDefaults instantiates a new Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMakerCommission

`func (o *Account) GetMakerCommission() int32`

GetMakerCommission returns the MakerCommission field if non-nil, zero value otherwise.

### GetMakerCommissionOk

`func (o *Account) GetMakerCommissionOk() (*int32, bool)`

GetMakerCommissionOk returns a tuple with the MakerCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakerCommission

`func (o *Account) SetMakerCommission(v int32)`

SetMakerCommission sets MakerCommission field to given value.

### HasMakerCommission

`func (o *Account) HasMakerCommission() bool`

HasMakerCommission returns a boolean if a field has been set.

### GetTakerCommission

`func (o *Account) GetTakerCommission() int32`

GetTakerCommission returns the TakerCommission field if non-nil, zero value otherwise.

### GetTakerCommissionOk

`func (o *Account) GetTakerCommissionOk() (*int32, bool)`

GetTakerCommissionOk returns a tuple with the TakerCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakerCommission

`func (o *Account) SetTakerCommission(v int32)`

SetTakerCommission sets TakerCommission field to given value.

### HasTakerCommission

`func (o *Account) HasTakerCommission() bool`

HasTakerCommission returns a boolean if a field has been set.

### GetBuyerCommission

`func (o *Account) GetBuyerCommission() int32`

GetBuyerCommission returns the BuyerCommission field if non-nil, zero value otherwise.

### GetBuyerCommissionOk

`func (o *Account) GetBuyerCommissionOk() (*int32, bool)`

GetBuyerCommissionOk returns a tuple with the BuyerCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerCommission

`func (o *Account) SetBuyerCommission(v int32)`

SetBuyerCommission sets BuyerCommission field to given value.

### HasBuyerCommission

`func (o *Account) HasBuyerCommission() bool`

HasBuyerCommission returns a boolean if a field has been set.

### GetSellerCommission

`func (o *Account) GetSellerCommission() int32`

GetSellerCommission returns the SellerCommission field if non-nil, zero value otherwise.

### GetSellerCommissionOk

`func (o *Account) GetSellerCommissionOk() (*int32, bool)`

GetSellerCommissionOk returns a tuple with the SellerCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerCommission

`func (o *Account) SetSellerCommission(v int32)`

SetSellerCommission sets SellerCommission field to given value.

### HasSellerCommission

`func (o *Account) HasSellerCommission() bool`

HasSellerCommission returns a boolean if a field has been set.

### GetCanTrade

`func (o *Account) GetCanTrade() bool`

GetCanTrade returns the CanTrade field if non-nil, zero value otherwise.

### GetCanTradeOk

`func (o *Account) GetCanTradeOk() (*bool, bool)`

GetCanTradeOk returns a tuple with the CanTrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanTrade

`func (o *Account) SetCanTrade(v bool)`

SetCanTrade sets CanTrade field to given value.

### HasCanTrade

`func (o *Account) HasCanTrade() bool`

HasCanTrade returns a boolean if a field has been set.

### GetCanWithdraw

`func (o *Account) GetCanWithdraw() bool`

GetCanWithdraw returns the CanWithdraw field if non-nil, zero value otherwise.

### GetCanWithdrawOk

`func (o *Account) GetCanWithdrawOk() (*bool, bool)`

GetCanWithdrawOk returns a tuple with the CanWithdraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanWithdraw

`func (o *Account) SetCanWithdraw(v bool)`

SetCanWithdraw sets CanWithdraw field to given value.

### HasCanWithdraw

`func (o *Account) HasCanWithdraw() bool`

HasCanWithdraw returns a boolean if a field has been set.

### GetCanDeposit

`func (o *Account) GetCanDeposit() bool`

GetCanDeposit returns the CanDeposit field if non-nil, zero value otherwise.

### GetCanDepositOk

`func (o *Account) GetCanDepositOk() (*bool, bool)`

GetCanDepositOk returns a tuple with the CanDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeposit

`func (o *Account) SetCanDeposit(v bool)`

SetCanDeposit sets CanDeposit field to given value.

### HasCanDeposit

`func (o *Account) HasCanDeposit() bool`

HasCanDeposit returns a boolean if a field has been set.

### GetUpdateTime

`func (o *Account) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *Account) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *Account) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *Account) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetAccountType

`func (o *Account) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *Account) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *Account) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *Account) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetBalances

`func (o *Account) GetBalances() []AccountBalances`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *Account) GetBalancesOk() (*[]AccountBalances, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *Account) SetBalances(v []AccountBalances)`

SetBalances sets Balances field to given value.

### HasBalances

`func (o *Account) HasBalances() bool`

HasBalances returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


