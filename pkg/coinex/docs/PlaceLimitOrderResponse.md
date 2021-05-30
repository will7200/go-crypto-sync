# PlaceLimitOrderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **string** | order count | [optional] 
**AvgPrice** | Pointer to **string** | average price | [optional] 
**CreateTime** | Pointer to **int64** | time when placing order | [optional] 
**DealAmount** | Pointer to **string** | count | [optional] 
**DealFee** | Pointer to **string** | transaction fee | [optional] 
**DealMoney** | Pointer to **string** | amount | [optional] 
**FinishedTime** | Pointer to **int64** | complete time | [optional] 
**Id** | Pointer to **int64** | Order No. | [optional] 
**MakerFeeRate** | Pointer to **string** | maker fee | [optional] 
**Market** | Pointer to **string** | See &lt;API invocation description market&gt; | [optional] 
**OrderType** | Pointer to **string** | limit:limit order;&lt;br&gt;market:market order; | [optional] 
**Price** | Pointer to **string** | order price | [optional] 
**Status** | Pointer to **string** | not\\_deal: unexecuted;&lt;br&gt;part\\_deal: partly executed;&lt;br&gt;done: executed; | [optional] 
**TakerFeeRate** | Pointer to **string** | taker fee | [optional] 
**Type** | Pointer to **string** | sell: sell order;&lt;br&gt;buy: buy order; | [optional] 
**ClientId** | Pointer to **string** | client_id: what you give | [optional] 

## Methods

### NewPlaceLimitOrderResponse

`func NewPlaceLimitOrderResponse() *PlaceLimitOrderResponse`

NewPlaceLimitOrderResponse instantiates a new PlaceLimitOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaceLimitOrderResponseWithDefaults

`func NewPlaceLimitOrderResponseWithDefaults() *PlaceLimitOrderResponse`

NewPlaceLimitOrderResponseWithDefaults instantiates a new PlaceLimitOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *PlaceLimitOrderResponse) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PlaceLimitOrderResponse) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PlaceLimitOrderResponse) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PlaceLimitOrderResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAvgPrice

`func (o *PlaceLimitOrderResponse) GetAvgPrice() string`

GetAvgPrice returns the AvgPrice field if non-nil, zero value otherwise.

### GetAvgPriceOk

`func (o *PlaceLimitOrderResponse) GetAvgPriceOk() (*string, bool)`

GetAvgPriceOk returns a tuple with the AvgPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPrice

`func (o *PlaceLimitOrderResponse) SetAvgPrice(v string)`

SetAvgPrice sets AvgPrice field to given value.

### HasAvgPrice

`func (o *PlaceLimitOrderResponse) HasAvgPrice() bool`

HasAvgPrice returns a boolean if a field has been set.

### GetCreateTime

`func (o *PlaceLimitOrderResponse) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *PlaceLimitOrderResponse) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *PlaceLimitOrderResponse) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *PlaceLimitOrderResponse) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDealAmount

`func (o *PlaceLimitOrderResponse) GetDealAmount() string`

GetDealAmount returns the DealAmount field if non-nil, zero value otherwise.

### GetDealAmountOk

`func (o *PlaceLimitOrderResponse) GetDealAmountOk() (*string, bool)`

GetDealAmountOk returns a tuple with the DealAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealAmount

`func (o *PlaceLimitOrderResponse) SetDealAmount(v string)`

SetDealAmount sets DealAmount field to given value.

### HasDealAmount

`func (o *PlaceLimitOrderResponse) HasDealAmount() bool`

HasDealAmount returns a boolean if a field has been set.

### GetDealFee

`func (o *PlaceLimitOrderResponse) GetDealFee() string`

GetDealFee returns the DealFee field if non-nil, zero value otherwise.

### GetDealFeeOk

`func (o *PlaceLimitOrderResponse) GetDealFeeOk() (*string, bool)`

GetDealFeeOk returns a tuple with the DealFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealFee

`func (o *PlaceLimitOrderResponse) SetDealFee(v string)`

SetDealFee sets DealFee field to given value.

### HasDealFee

`func (o *PlaceLimitOrderResponse) HasDealFee() bool`

HasDealFee returns a boolean if a field has been set.

### GetDealMoney

`func (o *PlaceLimitOrderResponse) GetDealMoney() string`

GetDealMoney returns the DealMoney field if non-nil, zero value otherwise.

### GetDealMoneyOk

`func (o *PlaceLimitOrderResponse) GetDealMoneyOk() (*string, bool)`

GetDealMoneyOk returns a tuple with the DealMoney field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealMoney

`func (o *PlaceLimitOrderResponse) SetDealMoney(v string)`

SetDealMoney sets DealMoney field to given value.

### HasDealMoney

`func (o *PlaceLimitOrderResponse) HasDealMoney() bool`

HasDealMoney returns a boolean if a field has been set.

### GetFinishedTime

`func (o *PlaceLimitOrderResponse) GetFinishedTime() int64`

GetFinishedTime returns the FinishedTime field if non-nil, zero value otherwise.

### GetFinishedTimeOk

`func (o *PlaceLimitOrderResponse) GetFinishedTimeOk() (*int64, bool)`

GetFinishedTimeOk returns a tuple with the FinishedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedTime

`func (o *PlaceLimitOrderResponse) SetFinishedTime(v int64)`

SetFinishedTime sets FinishedTime field to given value.

### HasFinishedTime

`func (o *PlaceLimitOrderResponse) HasFinishedTime() bool`

HasFinishedTime returns a boolean if a field has been set.

### GetId

`func (o *PlaceLimitOrderResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlaceLimitOrderResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlaceLimitOrderResponse) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PlaceLimitOrderResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMakerFeeRate

`func (o *PlaceLimitOrderResponse) GetMakerFeeRate() string`

GetMakerFeeRate returns the MakerFeeRate field if non-nil, zero value otherwise.

### GetMakerFeeRateOk

`func (o *PlaceLimitOrderResponse) GetMakerFeeRateOk() (*string, bool)`

GetMakerFeeRateOk returns a tuple with the MakerFeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakerFeeRate

`func (o *PlaceLimitOrderResponse) SetMakerFeeRate(v string)`

SetMakerFeeRate sets MakerFeeRate field to given value.

### HasMakerFeeRate

`func (o *PlaceLimitOrderResponse) HasMakerFeeRate() bool`

HasMakerFeeRate returns a boolean if a field has been set.

### GetMarket

`func (o *PlaceLimitOrderResponse) GetMarket() string`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *PlaceLimitOrderResponse) GetMarketOk() (*string, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *PlaceLimitOrderResponse) SetMarket(v string)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *PlaceLimitOrderResponse) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetOrderType

`func (o *PlaceLimitOrderResponse) GetOrderType() string`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *PlaceLimitOrderResponse) GetOrderTypeOk() (*string, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *PlaceLimitOrderResponse) SetOrderType(v string)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *PlaceLimitOrderResponse) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetPrice

`func (o *PlaceLimitOrderResponse) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PlaceLimitOrderResponse) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *PlaceLimitOrderResponse) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *PlaceLimitOrderResponse) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetStatus

`func (o *PlaceLimitOrderResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PlaceLimitOrderResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PlaceLimitOrderResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PlaceLimitOrderResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTakerFeeRate

`func (o *PlaceLimitOrderResponse) GetTakerFeeRate() string`

GetTakerFeeRate returns the TakerFeeRate field if non-nil, zero value otherwise.

### GetTakerFeeRateOk

`func (o *PlaceLimitOrderResponse) GetTakerFeeRateOk() (*string, bool)`

GetTakerFeeRateOk returns a tuple with the TakerFeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakerFeeRate

`func (o *PlaceLimitOrderResponse) SetTakerFeeRate(v string)`

SetTakerFeeRate sets TakerFeeRate field to given value.

### HasTakerFeeRate

`func (o *PlaceLimitOrderResponse) HasTakerFeeRate() bool`

HasTakerFeeRate returns a boolean if a field has been set.

### GetType

`func (o *PlaceLimitOrderResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PlaceLimitOrderResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PlaceLimitOrderResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PlaceLimitOrderResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetClientId

`func (o *PlaceLimitOrderResponse) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *PlaceLimitOrderResponse) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *PlaceLimitOrderResponse) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *PlaceLimitOrderResponse) HasClientId() bool`

HasClientId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


