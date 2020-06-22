# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | 
**Name** | Pointer to **string** |  | 
**Primary** | Pointer to **bool** |  | 
**Type** | Pointer to **string** |  | 
**Currency** | Pointer to [**Currency**](Currency.md) |  | 
**Balance** | Pointer to [**AccountBalance**](Account_balance.md) |  | 
**CreatedAt** | Pointer to [**time.Time**](time.Time.md) |  | 
**UpdatedAt** | Pointer to [**time.Time**](time.Time.md) |  | 
**Resource** | Pointer to **string** |  | 
**ResourcePath** | Pointer to **string** |  | 
**Ready** | Pointer to **bool** |  | 
**AllowDeposits** | Pointer to **bool** |  | [optional] 
**AllowWithdrawals** | Pointer to **bool** |  | [optional] 

## Methods

### NewAccount

`func NewAccount(id string, name string, primary bool, type_ string, currency Currency, balance AccountBalance, createdAt time.Time, updatedAt time.Time, resource string, resourcePath string, ready bool, ) *Account`

NewAccount instantiates a new Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithDefaults

`func NewAccountWithDefaults() *Account`

NewAccountWithDefaults instantiates a new Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Account) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Account) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Account) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Account) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Account) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Account) SetName(v string)`

SetName sets Name field to given value.


### GetPrimary

`func (o *Account) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *Account) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *Account) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.


### GetType

`func (o *Account) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Account) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Account) SetType(v string)`

SetType sets Type field to given value.


### GetCurrency

`func (o *Account) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Account) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Account) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.


### GetBalance

`func (o *Account) GetBalance() AccountBalance`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *Account) GetBalanceOk() (*AccountBalance, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *Account) SetBalance(v AccountBalance)`

SetBalance sets Balance field to given value.


### GetCreatedAt

`func (o *Account) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Account) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Account) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Account) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Account) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Account) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetResource

`func (o *Account) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *Account) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *Account) SetResource(v string)`

SetResource sets Resource field to given value.


### GetResourcePath

`func (o *Account) GetResourcePath() string`

GetResourcePath returns the ResourcePath field if non-nil, zero value otherwise.

### GetResourcePathOk

`func (o *Account) GetResourcePathOk() (*string, bool)`

GetResourcePathOk returns a tuple with the ResourcePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePath

`func (o *Account) SetResourcePath(v string)`

SetResourcePath sets ResourcePath field to given value.


### GetReady

`func (o *Account) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *Account) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *Account) SetReady(v bool)`

SetReady sets Ready field to given value.


### GetAllowDeposits

`func (o *Account) GetAllowDeposits() bool`

GetAllowDeposits returns the AllowDeposits field if non-nil, zero value otherwise.

### GetAllowDepositsOk

`func (o *Account) GetAllowDepositsOk() (*bool, bool)`

GetAllowDepositsOk returns a tuple with the AllowDeposits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDeposits

`func (o *Account) SetAllowDeposits(v bool)`

SetAllowDeposits sets AllowDeposits field to given value.

### HasAllowDeposits

`func (o *Account) HasAllowDeposits() bool`

HasAllowDeposits returns a boolean if a field has been set.

### GetAllowWithdrawals

`func (o *Account) GetAllowWithdrawals() bool`

GetAllowWithdrawals returns the AllowWithdrawals field if non-nil, zero value otherwise.

### GetAllowWithdrawalsOk

`func (o *Account) GetAllowWithdrawalsOk() (*bool, bool)`

GetAllowWithdrawalsOk returns a tuple with the AllowWithdrawals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowWithdrawals

`func (o *Account) SetAllowWithdrawals(v bool)`

SetAllowWithdrawals sets AllowWithdrawals field to given value.

### HasAllowWithdrawals

`func (o *Account) HasAllowWithdrawals() bool`

HasAllowWithdrawals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


