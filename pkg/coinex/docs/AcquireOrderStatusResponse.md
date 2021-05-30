# AcquireOrderStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **string** | order count | [optional] 
**AvgPrice** | Pointer to **string** | average price | [optional] 
**CreateTime** | Pointer to **int64** | time when placing order | [optional] 
**DealAmount** | Pointer to **string** | count | [optional] 
**DealFee** | Pointer to **string** | transaction fee | [optional] 
**DealMoney** | Pointer to **string** | executed value | [optional] 
**Id** | Pointer to **int64** | Order No. | [optional] 
**MakerFeeRate** | Pointer to **string** | maker fee | [optional] 
**Market** | Pointer to **string** | See &lt;API invocation description market&gt; | [optional] 
**OrderType** | Pointer to **string** | limit:limit order;&lt;br&gt;market:market order; | [optional] 
**Price** | Pointer to **string** | order price | [optional] 
**Status** | Pointer to **string** | not\\_deal: unexecuted;&lt;br&gt;part\\_deal: partly executed;&lt;br&gt;done: executed;&lt;br&gt;cancel: cancelled;&lt;br&gt;To guarantee server performance, all cancelled unexecuted orders will be deleted. | [optional] 
**TakerFeeRate** | Pointer to **string** | taker fee | [optional] 
**Type** | Pointer to **string** | sell: sell order;&lt;br&gt;buy: buy order; | [optional] 

## Methods

### NewAcquireOrderStatusResponse

`func NewAcquireOrderStatusResponse() *AcquireOrderStatusResponse`

NewAcquireOrderStatusResponse instantiates a new AcquireOrderStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcquireOrderStatusResponseWithDefaults

`func NewAcquireOrderStatusResponseWithDefaults() *AcquireOrderStatusResponse`

NewAcquireOrderStatusResponseWithDefaults instantiates a new AcquireOrderStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *AcquireOrderStatusResponse) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AcquireOrderStatusResponse) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AcquireOrderStatusResponse) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *AcquireOrderStatusResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAvgPrice

`func (o *AcquireOrderStatusResponse) GetAvgPrice() string`

GetAvgPrice returns the AvgPrice field if non-nil, zero value otherwise.

### GetAvgPriceOk

`func (o *AcquireOrderStatusResponse) GetAvgPriceOk() (*string, bool)`

GetAvgPriceOk returns a tuple with the AvgPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPrice

`func (o *AcquireOrderStatusResponse) SetAvgPrice(v string)`

SetAvgPrice sets AvgPrice field to given value.

### HasAvgPrice

`func (o *AcquireOrderStatusResponse) HasAvgPrice() bool`

HasAvgPrice returns a boolean if a field has been set.

### GetCreateTime

`func (o *AcquireOrderStatusResponse) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *AcquireOrderStatusResponse) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *AcquireOrderStatusResponse) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *AcquireOrderStatusResponse) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDealAmount

`func (o *AcquireOrderStatusResponse) GetDealAmount() string`

GetDealAmount returns the DealAmount field if non-nil, zero value otherwise.

### GetDealAmountOk

`func (o *AcquireOrderStatusResponse) GetDealAmountOk() (*string, bool)`

GetDealAmountOk returns a tuple with the DealAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealAmount

`func (o *AcquireOrderStatusResponse) SetDealAmount(v string)`

SetDealAmount sets DealAmount field to given value.

### HasDealAmount

`func (o *AcquireOrderStatusResponse) HasDealAmount() bool`

HasDealAmount returns a boolean if a field has been set.

### GetDealFee

`func (o *AcquireOrderStatusResponse) GetDealFee() string`

GetDealFee returns the DealFee field if non-nil, zero value otherwise.

### GetDealFeeOk

`func (o *AcquireOrderStatusResponse) GetDealFeeOk() (*string, bool)`

GetDealFeeOk returns a tuple with the DealFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealFee

`func (o *AcquireOrderStatusResponse) SetDealFee(v string)`

SetDealFee sets DealFee field to given value.

### HasDealFee

`func (o *AcquireOrderStatusResponse) HasDealFee() bool`

HasDealFee returns a boolean if a field has been set.

### GetDealMoney

`func (o *AcquireOrderStatusResponse) GetDealMoney() string`

GetDealMoney returns the DealMoney field if non-nil, zero value otherwise.

### GetDealMoneyOk

`func (o *AcquireOrderStatusResponse) GetDealMoneyOk() (*string, bool)`

GetDealMoneyOk returns a tuple with the DealMoney field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealMoney

`func (o *AcquireOrderStatusResponse) SetDealMoney(v string)`

SetDealMoney sets DealMoney field to given value.

### HasDealMoney

`func (o *AcquireOrderStatusResponse) HasDealMoney() bool`

HasDealMoney returns a boolean if a field has been set.

### GetId

`func (o *AcquireOrderStatusResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AcquireOrderStatusResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AcquireOrderStatusResponse) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AcquireOrderStatusResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMakerFeeRate

`func (o *AcquireOrderStatusResponse) GetMakerFeeRate() string`

GetMakerFeeRate returns the MakerFeeRate field if non-nil, zero value otherwise.

### GetMakerFeeRateOk

`func (o *AcquireOrderStatusResponse) GetMakerFeeRateOk() (*string, bool)`

GetMakerFeeRateOk returns a tuple with the MakerFeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakerFeeRate

`func (o *AcquireOrderStatusResponse) SetMakerFeeRate(v string)`

SetMakerFeeRate sets MakerFeeRate field to given value.

### HasMakerFeeRate

`func (o *AcquireOrderStatusResponse) HasMakerFeeRate() bool`

HasMakerFeeRate returns a boolean if a field has been set.

### GetMarket

`func (o *AcquireOrderStatusResponse) GetMarket() string`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *AcquireOrderStatusResponse) GetMarketOk() (*string, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *AcquireOrderStatusResponse) SetMarket(v string)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *AcquireOrderStatusResponse) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetOrderType

`func (o *AcquireOrderStatusResponse) GetOrderType() string`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *AcquireOrderStatusResponse) GetOrderTypeOk() (*string, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *AcquireOrderStatusResponse) SetOrderType(v string)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *AcquireOrderStatusResponse) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetPrice

`func (o *AcquireOrderStatusResponse) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *AcquireOrderStatusResponse) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *AcquireOrderStatusResponse) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *AcquireOrderStatusResponse) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetStatus

`func (o *AcquireOrderStatusResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AcquireOrderStatusResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AcquireOrderStatusResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AcquireOrderStatusResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTakerFeeRate

`func (o *AcquireOrderStatusResponse) GetTakerFeeRate() string`

GetTakerFeeRate returns the TakerFeeRate field if non-nil, zero value otherwise.

### GetTakerFeeRateOk

`func (o *AcquireOrderStatusResponse) GetTakerFeeRateOk() (*string, bool)`

GetTakerFeeRateOk returns a tuple with the TakerFeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakerFeeRate

`func (o *AcquireOrderStatusResponse) SetTakerFeeRate(v string)`

SetTakerFeeRate sets TakerFeeRate field to given value.

### HasTakerFeeRate

`func (o *AcquireOrderStatusResponse) HasTakerFeeRate() bool`

HasTakerFeeRate returns a boolean if a field has been set.

### GetType

`func (o *AcquireOrderStatusResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AcquireOrderStatusResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AcquireOrderStatusResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AcquireOrderStatusResponse) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


