# InquireMarginAccountInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **int64** | market account ID, See &lt;  Inquire Margin Account Market Info&gt; | [optional] 
**Leverage** | Pointer to **int64** | maximum leverage | [optional] 
**MarketType** | Pointer to **string** | &lt;API invocation description market&gt; | [optional] 
**SellAssetType** | Pointer to **string** | sell coin name, equal \&quot;sell_type\&quot; | [optional] 
**BuyAssetType** | Pointer to **string** | buy coin name, equal \&quot;buy_type\&quot; | [optional] 
**Balance** | Pointer to **map[string]interface{}** | balance | [optional] 
**Frozen** | Pointer to **map[string]interface{}** | amount frozen(not available) | [optional] 
**Loan** | Pointer to **map[string]interface{}** | amount loaned info | [optional] 
**Interest** | Pointer to **map[string]interface{}** | interest info | [optional] 
**CanTransfer** | Pointer to **map[string]interface{}** | available transfer amount | [optional] 
**WarnRate** | Pointer to **string** | warn rate | [optional] 
**LiquidationPrice** | Pointer to **string** | burst price | [optional] 

## Methods

### NewInquireMarginAccountInfoResponse

`func NewInquireMarginAccountInfoResponse() *InquireMarginAccountInfoResponse`

NewInquireMarginAccountInfoResponse instantiates a new InquireMarginAccountInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInquireMarginAccountInfoResponseWithDefaults

`func NewInquireMarginAccountInfoResponseWithDefaults() *InquireMarginAccountInfoResponse`

NewInquireMarginAccountInfoResponseWithDefaults instantiates a new InquireMarginAccountInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *InquireMarginAccountInfoResponse) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *InquireMarginAccountInfoResponse) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *InquireMarginAccountInfoResponse) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *InquireMarginAccountInfoResponse) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetLeverage

`func (o *InquireMarginAccountInfoResponse) GetLeverage() int64`

GetLeverage returns the Leverage field if non-nil, zero value otherwise.

### GetLeverageOk

`func (o *InquireMarginAccountInfoResponse) GetLeverageOk() (*int64, bool)`

GetLeverageOk returns a tuple with the Leverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeverage

`func (o *InquireMarginAccountInfoResponse) SetLeverage(v int64)`

SetLeverage sets Leverage field to given value.

### HasLeverage

`func (o *InquireMarginAccountInfoResponse) HasLeverage() bool`

HasLeverage returns a boolean if a field has been set.

### GetMarketType

`func (o *InquireMarginAccountInfoResponse) GetMarketType() string`

GetMarketType returns the MarketType field if non-nil, zero value otherwise.

### GetMarketTypeOk

`func (o *InquireMarginAccountInfoResponse) GetMarketTypeOk() (*string, bool)`

GetMarketTypeOk returns a tuple with the MarketType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketType

`func (o *InquireMarginAccountInfoResponse) SetMarketType(v string)`

SetMarketType sets MarketType field to given value.

### HasMarketType

`func (o *InquireMarginAccountInfoResponse) HasMarketType() bool`

HasMarketType returns a boolean if a field has been set.

### GetSellAssetType

`func (o *InquireMarginAccountInfoResponse) GetSellAssetType() string`

GetSellAssetType returns the SellAssetType field if non-nil, zero value otherwise.

### GetSellAssetTypeOk

`func (o *InquireMarginAccountInfoResponse) GetSellAssetTypeOk() (*string, bool)`

GetSellAssetTypeOk returns a tuple with the SellAssetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellAssetType

`func (o *InquireMarginAccountInfoResponse) SetSellAssetType(v string)`

SetSellAssetType sets SellAssetType field to given value.

### HasSellAssetType

`func (o *InquireMarginAccountInfoResponse) HasSellAssetType() bool`

HasSellAssetType returns a boolean if a field has been set.

### GetBuyAssetType

`func (o *InquireMarginAccountInfoResponse) GetBuyAssetType() string`

GetBuyAssetType returns the BuyAssetType field if non-nil, zero value otherwise.

### GetBuyAssetTypeOk

`func (o *InquireMarginAccountInfoResponse) GetBuyAssetTypeOk() (*string, bool)`

GetBuyAssetTypeOk returns a tuple with the BuyAssetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyAssetType

`func (o *InquireMarginAccountInfoResponse) SetBuyAssetType(v string)`

SetBuyAssetType sets BuyAssetType field to given value.

### HasBuyAssetType

`func (o *InquireMarginAccountInfoResponse) HasBuyAssetType() bool`

HasBuyAssetType returns a boolean if a field has been set.

### GetBalance

`func (o *InquireMarginAccountInfoResponse) GetBalance() map[string]interface{}`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *InquireMarginAccountInfoResponse) GetBalanceOk() (*map[string]interface{}, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *InquireMarginAccountInfoResponse) SetBalance(v map[string]interface{})`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *InquireMarginAccountInfoResponse) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetFrozen

`func (o *InquireMarginAccountInfoResponse) GetFrozen() map[string]interface{}`

GetFrozen returns the Frozen field if non-nil, zero value otherwise.

### GetFrozenOk

`func (o *InquireMarginAccountInfoResponse) GetFrozenOk() (*map[string]interface{}, bool)`

GetFrozenOk returns a tuple with the Frozen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozen

`func (o *InquireMarginAccountInfoResponse) SetFrozen(v map[string]interface{})`

SetFrozen sets Frozen field to given value.

### HasFrozen

`func (o *InquireMarginAccountInfoResponse) HasFrozen() bool`

HasFrozen returns a boolean if a field has been set.

### GetLoan

`func (o *InquireMarginAccountInfoResponse) GetLoan() map[string]interface{}`

GetLoan returns the Loan field if non-nil, zero value otherwise.

### GetLoanOk

`func (o *InquireMarginAccountInfoResponse) GetLoanOk() (*map[string]interface{}, bool)`

GetLoanOk returns a tuple with the Loan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoan

`func (o *InquireMarginAccountInfoResponse) SetLoan(v map[string]interface{})`

SetLoan sets Loan field to given value.

### HasLoan

`func (o *InquireMarginAccountInfoResponse) HasLoan() bool`

HasLoan returns a boolean if a field has been set.

### GetInterest

`func (o *InquireMarginAccountInfoResponse) GetInterest() map[string]interface{}`

GetInterest returns the Interest field if non-nil, zero value otherwise.

### GetInterestOk

`func (o *InquireMarginAccountInfoResponse) GetInterestOk() (*map[string]interface{}, bool)`

GetInterestOk returns a tuple with the Interest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterest

`func (o *InquireMarginAccountInfoResponse) SetInterest(v map[string]interface{})`

SetInterest sets Interest field to given value.

### HasInterest

`func (o *InquireMarginAccountInfoResponse) HasInterest() bool`

HasInterest returns a boolean if a field has been set.

### GetCanTransfer

`func (o *InquireMarginAccountInfoResponse) GetCanTransfer() map[string]interface{}`

GetCanTransfer returns the CanTransfer field if non-nil, zero value otherwise.

### GetCanTransferOk

`func (o *InquireMarginAccountInfoResponse) GetCanTransferOk() (*map[string]interface{}, bool)`

GetCanTransferOk returns a tuple with the CanTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanTransfer

`func (o *InquireMarginAccountInfoResponse) SetCanTransfer(v map[string]interface{})`

SetCanTransfer sets CanTransfer field to given value.

### HasCanTransfer

`func (o *InquireMarginAccountInfoResponse) HasCanTransfer() bool`

HasCanTransfer returns a boolean if a field has been set.

### GetWarnRate

`func (o *InquireMarginAccountInfoResponse) GetWarnRate() string`

GetWarnRate returns the WarnRate field if non-nil, zero value otherwise.

### GetWarnRateOk

`func (o *InquireMarginAccountInfoResponse) GetWarnRateOk() (*string, bool)`

GetWarnRateOk returns a tuple with the WarnRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnRate

`func (o *InquireMarginAccountInfoResponse) SetWarnRate(v string)`

SetWarnRate sets WarnRate field to given value.

### HasWarnRate

`func (o *InquireMarginAccountInfoResponse) HasWarnRate() bool`

HasWarnRate returns a boolean if a field has been set.

### GetLiquidationPrice

`func (o *InquireMarginAccountInfoResponse) GetLiquidationPrice() string`

GetLiquidationPrice returns the LiquidationPrice field if non-nil, zero value otherwise.

### GetLiquidationPriceOk

`func (o *InquireMarginAccountInfoResponse) GetLiquidationPriceOk() (*string, bool)`

GetLiquidationPriceOk returns a tuple with the LiquidationPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidationPrice

`func (o *InquireMarginAccountInfoResponse) SetLiquidationPrice(v string)`

SetLiquidationPrice sets LiquidationPrice field to given value.

### HasLiquidationPrice

`func (o *InquireMarginAccountInfoResponse) HasLiquidationPrice() bool`

HasLiquidationPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


