# AcquireAllMarketDataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **string** | server time when returning | [optional] 
**Buy** | Pointer to **string** | buy 1 | [optional] 
**BuyAmount** | Pointer to **string** | buy 1 amount | [optional] 
**High** | Pointer to **string** | 24H highest price | [optional] 
**Last** | Pointer to **string** | latest price | [optional] 
**Low** | Pointer to **string** | 24H lowest price | [optional] 
**Sell** | Pointer to **string** | sell 1 | [optional] 
**SellAmount** | Pointer to **string** | sell 1 amount | [optional] 
**Vol** | Pointer to **string** | 24H volumn | [optional] 

## Methods

### NewAcquireAllMarketDataResponse

`func NewAcquireAllMarketDataResponse() *AcquireAllMarketDataResponse`

NewAcquireAllMarketDataResponse instantiates a new AcquireAllMarketDataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcquireAllMarketDataResponseWithDefaults

`func NewAcquireAllMarketDataResponseWithDefaults() *AcquireAllMarketDataResponse`

NewAcquireAllMarketDataResponseWithDefaults instantiates a new AcquireAllMarketDataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *AcquireAllMarketDataResponse) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *AcquireAllMarketDataResponse) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *AcquireAllMarketDataResponse) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *AcquireAllMarketDataResponse) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetBuy

`func (o *AcquireAllMarketDataResponse) GetBuy() string`

GetBuy returns the Buy field if non-nil, zero value otherwise.

### GetBuyOk

`func (o *AcquireAllMarketDataResponse) GetBuyOk() (*string, bool)`

GetBuyOk returns a tuple with the Buy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuy

`func (o *AcquireAllMarketDataResponse) SetBuy(v string)`

SetBuy sets Buy field to given value.

### HasBuy

`func (o *AcquireAllMarketDataResponse) HasBuy() bool`

HasBuy returns a boolean if a field has been set.

### GetBuyAmount

`func (o *AcquireAllMarketDataResponse) GetBuyAmount() string`

GetBuyAmount returns the BuyAmount field if non-nil, zero value otherwise.

### GetBuyAmountOk

`func (o *AcquireAllMarketDataResponse) GetBuyAmountOk() (*string, bool)`

GetBuyAmountOk returns a tuple with the BuyAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyAmount

`func (o *AcquireAllMarketDataResponse) SetBuyAmount(v string)`

SetBuyAmount sets BuyAmount field to given value.

### HasBuyAmount

`func (o *AcquireAllMarketDataResponse) HasBuyAmount() bool`

HasBuyAmount returns a boolean if a field has been set.

### GetHigh

`func (o *AcquireAllMarketDataResponse) GetHigh() string`

GetHigh returns the High field if non-nil, zero value otherwise.

### GetHighOk

`func (o *AcquireAllMarketDataResponse) GetHighOk() (*string, bool)`

GetHighOk returns a tuple with the High field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigh

`func (o *AcquireAllMarketDataResponse) SetHigh(v string)`

SetHigh sets High field to given value.

### HasHigh

`func (o *AcquireAllMarketDataResponse) HasHigh() bool`

HasHigh returns a boolean if a field has been set.

### GetLast

`func (o *AcquireAllMarketDataResponse) GetLast() string`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *AcquireAllMarketDataResponse) GetLastOk() (*string, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *AcquireAllMarketDataResponse) SetLast(v string)`

SetLast sets Last field to given value.

### HasLast

`func (o *AcquireAllMarketDataResponse) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetLow

`func (o *AcquireAllMarketDataResponse) GetLow() string`

GetLow returns the Low field if non-nil, zero value otherwise.

### GetLowOk

`func (o *AcquireAllMarketDataResponse) GetLowOk() (*string, bool)`

GetLowOk returns a tuple with the Low field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLow

`func (o *AcquireAllMarketDataResponse) SetLow(v string)`

SetLow sets Low field to given value.

### HasLow

`func (o *AcquireAllMarketDataResponse) HasLow() bool`

HasLow returns a boolean if a field has been set.

### GetSell

`func (o *AcquireAllMarketDataResponse) GetSell() string`

GetSell returns the Sell field if non-nil, zero value otherwise.

### GetSellOk

`func (o *AcquireAllMarketDataResponse) GetSellOk() (*string, bool)`

GetSellOk returns a tuple with the Sell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSell

`func (o *AcquireAllMarketDataResponse) SetSell(v string)`

SetSell sets Sell field to given value.

### HasSell

`func (o *AcquireAllMarketDataResponse) HasSell() bool`

HasSell returns a boolean if a field has been set.

### GetSellAmount

`func (o *AcquireAllMarketDataResponse) GetSellAmount() string`

GetSellAmount returns the SellAmount field if non-nil, zero value otherwise.

### GetSellAmountOk

`func (o *AcquireAllMarketDataResponse) GetSellAmountOk() (*string, bool)`

GetSellAmountOk returns a tuple with the SellAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellAmount

`func (o *AcquireAllMarketDataResponse) SetSellAmount(v string)`

SetSellAmount sets SellAmount field to given value.

### HasSellAmount

`func (o *AcquireAllMarketDataResponse) HasSellAmount() bool`

HasSellAmount returns a boolean if a field has been set.

### GetVol

`func (o *AcquireAllMarketDataResponse) GetVol() string`

GetVol returns the Vol field if non-nil, zero value otherwise.

### GetVolOk

`func (o *AcquireAllMarketDataResponse) GetVolOk() (*string, bool)`

GetVolOk returns a tuple with the Vol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVol

`func (o *AcquireAllMarketDataResponse) SetVol(v string)`

SetVol sets Vol field to given value.

### HasVol

`func (o *AcquireAllMarketDataResponse) HasVol() bool`

HasVol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


