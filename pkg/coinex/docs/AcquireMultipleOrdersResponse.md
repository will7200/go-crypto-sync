# AcquireMultipleOrdersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **string** | order count | [optional] 
**AvgPrice** | Pointer to **string** | average price | [optional] 
**CreateTime** | Pointer to **int64** | time when placing order | [optional] 
**DealAmount** | Pointer to **string** | count | [optional] 
**DealFee** | Pointer to **string** | transaction fee | [optional] 
**DealMoney** | Pointer to **string** | executed value | [optional] 
**FinishedTime** | Pointer to **int64** | complete time | [optional] 
**Id** | Pointer to **int64** | Order No. | [optional] 
**MakerFeeRate** | Pointer to **string** | maker fee | [optional] 
**Market** | Pointer to **string** | See &lt;API invocation description market&gt; | [optional] 
**OrderType** | Pointer to **string** | limit:limit order;&lt;br&gt;market:market order; | [optional] 
**Price** | Pointer to **string** | order price | [optional] 
**Status** | Pointer to **string** | not\\_deal: unexecuted;&lt;br&gt;part\\_deal: partly executed;&lt;br&gt;done: executed;&lt;br&gt;cancel: cancelled;&lt;br&gt;To guarantee server performance, all cancelled unexecuted orders will be deleted. | [optional] 
**TakerFeeRate** | Pointer to **string** | taker fee | [optional] 
**Type** | Pointer to **string** | sell: sell order;&lt;br&gt;buy: buy order; | [optional] 

## Methods

### NewAcquireMultipleOrdersResponse

`func NewAcquireMultipleOrdersResponse() *AcquireMultipleOrdersResponse`

NewAcquireMultipleOrdersResponse instantiates a new AcquireMultipleOrdersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcquireMultipleOrdersResponseWithDefaults

`func NewAcquireMultipleOrdersResponseWithDefaults() *AcquireMultipleOrdersResponse`

NewAcquireMultipleOrdersResponseWithDefaults instantiates a new AcquireMultipleOrdersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *AcquireMultipleOrdersResponse) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AcquireMultipleOrdersResponse) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AcquireMultipleOrdersResponse) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *AcquireMultipleOrdersResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAvgPrice

`func (o *AcquireMultipleOrdersResponse) GetAvgPrice() string`

GetAvgPrice returns the AvgPrice field if non-nil, zero value otherwise.

### GetAvgPriceOk

`func (o *AcquireMultipleOrdersResponse) GetAvgPriceOk() (*string, bool)`

GetAvgPriceOk returns a tuple with the AvgPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPrice

`func (o *AcquireMultipleOrdersResponse) SetAvgPrice(v string)`

SetAvgPrice sets AvgPrice field to given value.

### HasAvgPrice

`func (o *AcquireMultipleOrdersResponse) HasAvgPrice() bool`

HasAvgPrice returns a boolean if a field has been set.

### GetCreateTime

`func (o *AcquireMultipleOrdersResponse) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *AcquireMultipleOrdersResponse) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *AcquireMultipleOrdersResponse) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *AcquireMultipleOrdersResponse) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDealAmount

`func (o *AcquireMultipleOrdersResponse) GetDealAmount() string`

GetDealAmount returns the DealAmount field if non-nil, zero value otherwise.

### GetDealAmountOk

`func (o *AcquireMultipleOrdersResponse) GetDealAmountOk() (*string, bool)`

GetDealAmountOk returns a tuple with the DealAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealAmount

`func (o *AcquireMultipleOrdersResponse) SetDealAmount(v string)`

SetDealAmount sets DealAmount field to given value.

### HasDealAmount

`func (o *AcquireMultipleOrdersResponse) HasDealAmount() bool`

HasDealAmount returns a boolean if a field has been set.

### GetDealFee

`func (o *AcquireMultipleOrdersResponse) GetDealFee() string`

GetDealFee returns the DealFee field if non-nil, zero value otherwise.

### GetDealFeeOk

`func (o *AcquireMultipleOrdersResponse) GetDealFeeOk() (*string, bool)`

GetDealFeeOk returns a tuple with the DealFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealFee

`func (o *AcquireMultipleOrdersResponse) SetDealFee(v string)`

SetDealFee sets DealFee field to given value.

### HasDealFee

`func (o *AcquireMultipleOrdersResponse) HasDealFee() bool`

HasDealFee returns a boolean if a field has been set.

### GetDealMoney

`func (o *AcquireMultipleOrdersResponse) GetDealMoney() string`

GetDealMoney returns the DealMoney field if non-nil, zero value otherwise.

### GetDealMoneyOk

`func (o *AcquireMultipleOrdersResponse) GetDealMoneyOk() (*string, bool)`

GetDealMoneyOk returns a tuple with the DealMoney field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealMoney

`func (o *AcquireMultipleOrdersResponse) SetDealMoney(v string)`

SetDealMoney sets DealMoney field to given value.

### HasDealMoney

`func (o *AcquireMultipleOrdersResponse) HasDealMoney() bool`

HasDealMoney returns a boolean if a field has been set.

### GetFinishedTime

`func (o *AcquireMultipleOrdersResponse) GetFinishedTime() int64`

GetFinishedTime returns the FinishedTime field if non-nil, zero value otherwise.

### GetFinishedTimeOk

`func (o *AcquireMultipleOrdersResponse) GetFinishedTimeOk() (*int64, bool)`

GetFinishedTimeOk returns a tuple with the FinishedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedTime

`func (o *AcquireMultipleOrdersResponse) SetFinishedTime(v int64)`

SetFinishedTime sets FinishedTime field to given value.

### HasFinishedTime

`func (o *AcquireMultipleOrdersResponse) HasFinishedTime() bool`

HasFinishedTime returns a boolean if a field has been set.

### GetId

`func (o *AcquireMultipleOrdersResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AcquireMultipleOrdersResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AcquireMultipleOrdersResponse) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AcquireMultipleOrdersResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMakerFeeRate

`func (o *AcquireMultipleOrdersResponse) GetMakerFeeRate() string`

GetMakerFeeRate returns the MakerFeeRate field if non-nil, zero value otherwise.

### GetMakerFeeRateOk

`func (o *AcquireMultipleOrdersResponse) GetMakerFeeRateOk() (*string, bool)`

GetMakerFeeRateOk returns a tuple with the MakerFeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakerFeeRate

`func (o *AcquireMultipleOrdersResponse) SetMakerFeeRate(v string)`

SetMakerFeeRate sets MakerFeeRate field to given value.

### HasMakerFeeRate

`func (o *AcquireMultipleOrdersResponse) HasMakerFeeRate() bool`

HasMakerFeeRate returns a boolean if a field has been set.

### GetMarket

`func (o *AcquireMultipleOrdersResponse) GetMarket() string`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *AcquireMultipleOrdersResponse) GetMarketOk() (*string, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *AcquireMultipleOrdersResponse) SetMarket(v string)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *AcquireMultipleOrdersResponse) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetOrderType

`func (o *AcquireMultipleOrdersResponse) GetOrderType() string`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *AcquireMultipleOrdersResponse) GetOrderTypeOk() (*string, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *AcquireMultipleOrdersResponse) SetOrderType(v string)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *AcquireMultipleOrdersResponse) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetPrice

`func (o *AcquireMultipleOrdersResponse) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *AcquireMultipleOrdersResponse) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *AcquireMultipleOrdersResponse) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *AcquireMultipleOrdersResponse) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetStatus

`func (o *AcquireMultipleOrdersResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AcquireMultipleOrdersResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AcquireMultipleOrdersResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AcquireMultipleOrdersResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTakerFeeRate

`func (o *AcquireMultipleOrdersResponse) GetTakerFeeRate() string`

GetTakerFeeRate returns the TakerFeeRate field if non-nil, zero value otherwise.

### GetTakerFeeRateOk

`func (o *AcquireMultipleOrdersResponse) GetTakerFeeRateOk() (*string, bool)`

GetTakerFeeRateOk returns a tuple with the TakerFeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakerFeeRate

`func (o *AcquireMultipleOrdersResponse) SetTakerFeeRate(v string)`

SetTakerFeeRate sets TakerFeeRate field to given value.

### HasTakerFeeRate

`func (o *AcquireMultipleOrdersResponse) HasTakerFeeRate() bool`

HasTakerFeeRate returns a boolean if a field has been set.

### GetType

`func (o *AcquireMultipleOrdersResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AcquireMultipleOrdersResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AcquireMultipleOrdersResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AcquireMultipleOrdersResponse) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


