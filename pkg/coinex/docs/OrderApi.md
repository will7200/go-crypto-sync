# \OrderApi

All URIs are relative to *https://api.coinex.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcquireExecutedOrderDetail**](OrderApi.md#AcquireExecutedOrderDetail) | **Get** /order/deals | Acquire Executed Order Detail
[**AcquireExecutedOrderList**](OrderApi.md#AcquireExecutedOrderList) | **Get** /order/finished | Acquire Executed Order List
[**AcquireFinishedOrderDetail**](OrderApi.md#AcquireFinishedOrderDetail) | **Get** /order/finished/{id} | Acquire Finished Order Detail
[**AcquireMiningDifficulty**](OrderApi.md#AcquireMiningDifficulty) | **Get** /order/mining/difficulty | Acquire Mining Difficulty
[**AcquireMultipleOrders**](OrderApi.md#AcquireMultipleOrders) | **Get** /order/status/batch | Acquire Multiple Orders
[**AcquireOrderStatus**](OrderApi.md#AcquireOrderStatus) | **Get** /order/status | Acquire Order Status 
[**AcquireUnexecutedOrderList**](OrderApi.md#AcquireUnexecutedOrderList) | **Get** /order/pending | Acquire Unexecuted Order List
[**AcquireUserDeals**](OrderApi.md#AcquireUserDeals) | **Get** /order/user/deals | Acquire User Deals
[**CancelAllOrder**](OrderApi.md#CancelAllOrder) | **Delete** /order/pending | Cancel All Order
[**CancelMultipleOrders**](OrderApi.md#CancelMultipleOrders) | **Delete** /order/pending/batch | Cancel Multiple Orders
[**PlaceIocOrder**](OrderApi.md#PlaceIocOrder) | **Post** /order/ioc | Place IOC Order
[**PlaceLimitOrder**](OrderApi.md#PlaceLimitOrder) | **Post** /order/limit | Place Limit Order
[**PlaceMarketOrder**](OrderApi.md#PlaceMarketOrder) | **Post** /order/market | Place Market Order
[**PlaceMultipleLimitOrders**](OrderApi.md#PlaceMultipleLimitOrders) | **Post** /order/limit/batch | Place Multiple Limit Orders



## AcquireExecutedOrderDetail

> InlineResponse2008 AcquireExecutedOrderDetail(ctx).Authorization(authorization).AccessId(accessId).Id(id).Page(page).Execute()

Acquire Executed Order Detail



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    authorization := "authorization_example" // string | Signature is required for Account API and trading API related interfaces. The signature data is placed in the authorization header of the HTTP header and authorization is the signature result string. No signature is required for market API related interfaces. Use 32-bit MD5 Algorithm Signature Use MD5 algorithm to encrypt the signature string, convert encrypted result to uppercase, get signature data and put signature data in HTTP Header - authorization. 
    accessId := "F0160AD40BAE4519A3836BFB057BD682" // string | access_id
    id := int64(123456) // int64 | Order no.
    page := int64(1) // int64 | page

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrderApi.AcquireExecutedOrderDetail(context.Background()).Authorization(authorization).AccessId(accessId).Id(id).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.AcquireExecutedOrderDetail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcquireExecutedOrderDetail`: InlineResponse2008
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.AcquireExecutedOrderDetail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAcquireExecutedOrderDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Signature is required for Account API and trading API related interfaces. The signature data is placed in the authorization header of the HTTP header and authorization is the signature result string. No signature is required for market API related interfaces. Use 32-bit MD5 Algorithm Signature Use MD5 algorithm to encrypt the signature string, convert encrypted result to uppercase, get signature data and put signature data in HTTP Header - authorization.  | 
 **accessId** | **string** | access_id | 
 **id** | **int64** | Order no. | 
 **page** | **int64** | page | 

### Return type

[**InlineResponse2008**](InlineResponse2008.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AcquireExecutedOrderList

> InlineResponse20017 AcquireExecutedOrderList(ctx).Authorization(authorization).AccessId(accessId).Market(market).Page(page).Execute()

Acquire Executed Order List



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    authorization := "authorization_example" // string | Signature is required for Account API and trading API related interfaces. The signature data is placed in the authorization header of the HTTP header and authorization is the signature result string. No signature is required for market API related interfaces. Use 32-bit MD5 Algorithm Signature Use MD5 algorithm to encrypt the signature string, convert encrypted result to uppercase, get signature data and put signature data in HTTP Header - authorization. 
    accessId := "EFF0C4057AA240F6AAE52634E07EC0DC" // string | access_id
    market := "BTCBCH" // string | See <API invocation description market>
    page := int64(1) // int64 | page

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrderApi.AcquireExecutedOrderList(context.Background()).Authorization(authorization).AccessId(accessId).Market(market).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.AcquireExecutedOrderList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcquireExecutedOrderList`: InlineResponse20017
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.AcquireExecutedOrderList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAcquireExecutedOrderListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Signature is required for Account API and trading API related interfaces. The signature data is placed in the authorization header of the HTTP header and authorization is the signature result string. No signature is required for market API related interfaces. Use 32-bit MD5 Algorithm Signature Use MD5 algorithm to encrypt the signature string, convert encrypted result to uppercase, get signature data and put signature data in HTTP Header - authorization.  | 
 **accessId** | **string** | access_id | 
 **market** | **string** | See &lt;API invocation description market&gt; | 
 **page** | **int64** | page | 

### Return type

[**InlineResponse20017**](InlineResponse20017.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AcquireFinishedOrderDetail

> InlineResponse2008 AcquireFinishedOrderDetail(ctx, id).Authorization(authorization).AccessId(accessId).Market(market).Page(page).Execute()

Acquire Finished Order Detail



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    authorization := "authorization_example" // string | Signature is required for Account API and trading API related interfaces. The signature data is placed in the authorization header of the HTTP header and authorization is the signature result string. No signature is required for market API related interfaces. Use 32-bit MD5 Algorithm Signature Use MD5 algorithm to encrypt the signature string, convert encrypted result to uppercase, get signature data and put signature data in HTTP Header - authorization. 
    id := "300034" // string | order no.
    accessId := "F0160AD40BAE4519A3836BFB057BD682" // string | access_id
    market := "market_example" // string | See <API invocation description market>
    page := int64(1) // int64 | page

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrderApi.AcquireFinishedOrderDetail(context.Background(), id).Authorization(authorization).AccessId(accessId).Market(market).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.AcquireFinishedOrderDetail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcquireFinishedOrderDetail`: InlineResponse2008
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.AcquireFinishedOrderDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | order no. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcquireFinishedOrderDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Signature is required for Account API and trading API related interfaces. The signature data is placed in the authorization header of the HTTP header and authorization is the signature result string. No signature is required for market API related interfaces. Use 32-bit MD5 Algorithm Signature Use MD5 algorithm to encrypt the signature string, convert encrypted result to uppercase, get signature data and put signature data in HTTP Header - authorization.  | 

 **accessId** | **string** | access_id | 
 **market** | **string** | See &lt;API invocation description market&gt; | 
 **page** | **int64** | page | 

### Return type

[**InlineResponse2008**](InlineResponse2008.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AcquireMiningDifficulty

> InlineResponse2007 AcquireMiningDifficulty(ctx).Authorization(authorization).Execute()

Acquire Mining Difficulty



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    authorization := "authorization_example" // string | Signature is required for Account API and trading API related interfaces. The signature data is placed in the authorization header of the HTTP header and authorization is the signature result string. No signature is required for market API related interfaces. Use 32-bit MD5 Algorithm Signature Use MD5 algorithm to encrypt the signature string, convert encrypted result to uppercase, get signature data and put signature data in HTTP Header - authorization. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrderApi.AcquireMiningDifficulty(context.Background()).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.AcquireMiningDifficulty``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcquireMiningDifficulty`: InlineResponse2007
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.AcquireMiningDifficulty`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAcquireMiningDifficultyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Signature is required for Account API and trading API related interfaces. The signature data is placed in the authorization header of the HTTP header and authorization is the signature result string. No signature is required for market API related interfaces. Use 32-bit MD5 Algorithm Signature Use MD5 algorithm to encrypt the signature string, convert encrypted result to uppercase, get signature data and put signature data in HTTP Header - authorization.  | 

### Return type

[**InlineResponse2007**](InlineResponse2007.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AcquireMultipleOrders

> InlineResponse20010 AcquireMultipleOrders(ctx).Authorization(authorization).AccessId(accessId).BatchIds(batchIds).Market(market).Tonce(tonce).Execute()

Acquire Multiple Orders



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    authorization := "authorization_example" // string | Signature is required for Account API and trading API related interfaces. The signature data is placed in the authorization header of the HTTP header and authorization is the signature result string. No signature is required for market API related interfaces. Use 32-bit MD5 Algorithm Signature Use MD5 algorithm to encrypt the signature string, convert encrypted result to uppercase, get signature data and put signature data in HTTP Header - authorization. 
    accessId := "F0160AD40BAE4519A3836BFB057BD682" // string | access_id
    batchIds := "300021,300022,300033" // string | multiple order nos string
    market := "BTCBCH" // string | See <API invocation description market>
    tonce := int64(1513746038205) // int64 | Tonce is a timestamp with a positive Interger that represents the number of milliseconds from Unix epoch to the current time. Error between tonce and server time can not exceed plus or minus 60s

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrderApi.AcquireMultipleOrders(context.Background()).Authorization(authorization).AccessId(accessId).BatchIds(batchIds).Market(market).Tonce(tonce).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.AcquireMultipleOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcquireMultipleOrders`: InlineResponse20010
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.AcquireMultipleOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAcquireMultipleOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Signature is required for Account API and trading API related interfaces. The signature data is placed in the authorization header of the HTTP header and authorization is the signature result string. No signature is required for market API related interfaces. Use 32-bit MD5 Algorithm Signature Use MD5 algorithm to encrypt the signature string, convert encrypted result to uppercase, get signature data and put signature data in HTTP Header - authorization.  | 
 **accessId** | **string** | access_id | 
 **batchIds** | **string** | multiple order nos string | 
 **market** | **string** | See &lt;API invocation description market&gt; | 
 **tonce** | **int64** | Tonce is a timestamp with a positive Interger that represents the number of milliseconds from Unix epoch to the current time. Error between tonce and server time can not exceed plus or minus 60s | 

### Return type

[**InlineResponse20010**](InlineResponse20010.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AcquireOrderStatus

> InlineResponse20018 AcquireOrderStatus(ctx).Authorization(authorization).AccessId(accessId).Id(id).Market(market).Tonce(tonce).Execute()

Acquire Order Status 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    authorization := "authorization_example" // string | Signature is required for Account API and trading API related interfaces. The signature data is placed in the authorization header of the HTTP header and authorization is the signature result string. No signature is required for market API related interfaces. Use 32-bit MD5 Algorithm Signature Use MD5 algorithm to encrypt the signature string, convert encrypted result to uppercase, get signature data and put signature data in HTTP Header - authorization. 
    accessId := "F0160AD40BAE4519A3836BFB057BD682" // string | access_id
    id := int64(300021) // int64 | order no
    market := "btcbch" // string | See <API invocation description market>
    tonce := int64(1513746038205) // int64 | Tonce is a timestamp with a positive Interger that represents the number of milliseconds from Unix epoch to the current time. Error between tonce and server time can not exceed plus or minus 60s

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrderApi.AcquireOrderStatus(context.Background()).Authorization(authorization).AccessId(accessId).Id(id).Market(market).Tonce(tonce).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.AcquireOrderStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcquireOrderStatus`: InlineResponse20018
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.AcquireOrderStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAcquireOrderStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Signature is required for Account API and trading API related interfaces. The signature data is placed in the authorization header of the HTTP header and authorization is the signature result string. No signature is required for market API related interfaces. Use 32-bit MD5 Algorithm Signature Use MD5 algorithm to encrypt the signature string, convert encrypted result to uppercase, get signature data and put signature data in HTTP Header - authorization.  | 
 **accessId** | **string** | access_id | 
 **id** | **int64** | order no | 
 **market** | **string** | See &lt;API invocation description market&gt; | 
 **tonce** | **int64** | Tonce is a timestamp with a positive Interger that represents the number of milliseconds from Unix epoch to the current time. Error between tonce and server time can not exceed plus or minus 60s | 

### Return type

[**InlineResponse20018**](InlineResponse20018.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AcquireUnexecutedOrderList

> InlineResponse20012 AcquireUnexecutedOrderList(ctx).Authorization(authorization).AccessId(accessId).Market(market).Page(page).Execute()

Acquire Unexecuted Order List



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    authorization := "authorization_example" // string | Signature is required for Account API and trading API related interfaces. The signature data is placed in the authorization header of the HTTP header and authorization is the signature result string. No signature is required for market API related interfaces. Use 32-bit MD5 Algorithm Signature Use MD5 algorithm to encrypt the signature string, convert encrypted result to uppercase, get signature data and put signature data in HTTP Header - authorization. 
    accessId := "BFFA64957AA240F6BBEA26F4E07EC0D9" // string | access_id
    market := "BTCBCH" // string | See <API invocation description market>
    page := int64(1) // int64 | Page

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrderApi.AcquireUnexecutedOrderList(context.Background()).Authorization(authorization).AccessId(accessId).Market(market).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.AcquireUnexecutedOrderList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcquireUnexecutedOrderList`: InlineResponse20012
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.AcquireUnexecutedOrderList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAcquireUnexecutedOrderListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Signature is required for Account API and trading API related interfaces. The signature data is placed in the authorization header of the HTTP header and authorization is the signature result string. No signature is required for market API related interfaces. Use 32-bit MD5 Algorithm Signature Use MD5 algorithm to encrypt the signature string, convert encrypted result to uppercase, get signature data and put signature data in HTTP Header - authorization.  | 
 **accessId** | **string** | access_id | 
 **market** | **string** | See &lt;API invocation description market&gt; | 
 **page** | **int64** | Page | 

### Return type

[**InlineResponse20012**](InlineResponse20012.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AcquireUserDeals

> InlineResponse20019 AcquireUserDeals(ctx).Authorization(authorization).AccessId(accessId).Market(market).Page(page).Execute()

Acquire User Deals



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    authorization := "authorization_example" // string | Signature is required for Account API and trading API related interfaces. The signature data is placed in the authorization header of the HTTP header and authorization is the signature result string. No signature is required for market API related interfaces. Use 32-bit MD5 Algorithm Signature Use MD5 algorithm to encrypt the signature string, convert encrypted result to uppercase, get signature data and put signature data in HTTP Header - authorization. 
    accessId := "F0160AD40BAE4519A3836BFB057BD682" // string | access_id
    market := "btcbch" // string | See <API invocation description market>
    page := int64(1) // int64 | page

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrderApi.AcquireUserDeals(context.Background()).Authorization(authorization).AccessId(accessId).Market(market).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.AcquireUserDeals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcquireUserDeals`: InlineResponse20019
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.AcquireUserDeals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAcquireUserDealsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Signature is required for Account API and trading API related interfaces. The signature data is placed in the authorization header of the HTTP header and authorization is the signature result string. No signature is required for market API related interfaces. Use 32-bit MD5 Algorithm Signature Use MD5 algorithm to encrypt the signature string, convert encrypted result to uppercase, get signature data and put signature data in HTTP Header - authorization.  | 
 **accessId** | **string** | access_id | 
 **market** | **string** | See &lt;API invocation description market&gt; | 
 **page** | **int64** | page | 

### Return type

[**InlineResponse20019**](InlineResponse20019.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelAllOrder

> InlineResponse20013 CancelAllOrder(ctx).Authorization(authorization).AccessId(accessId).AccountId(accountId).Execute()

Cancel All Order



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    authorization := "authorization_example" // string | Signature is required for Account API and trading API related interfaces. The signature data is placed in the authorization header of the HTTP header and authorization is the signature result string. No signature is required for market API related interfaces. Use 32-bit MD5 Algorithm Signature Use MD5 algorithm to encrypt the signature string, convert encrypted result to uppercase, get signature data and put signature data in HTTP Header - authorization. 
    accessId := "BFFA64957AA240F6BBEA26F4E07EC0D9" // string | access_id
    accountId := int64(0) // int64 | main account ID: 0

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrderApi.CancelAllOrder(context.Background()).Authorization(authorization).AccessId(accessId).AccountId(accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.CancelAllOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelAllOrder`: InlineResponse20013
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.CancelAllOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelAllOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Signature is required for Account API and trading API related interfaces. The signature data is placed in the authorization header of the HTTP header and authorization is the signature result string. No signature is required for market API related interfaces. Use 32-bit MD5 Algorithm Signature Use MD5 algorithm to encrypt the signature string, convert encrypted result to uppercase, get signature data and put signature data in HTTP Header - authorization.  | 
 **accessId** | **string** | access_id | 
 **accountId** | **int64** | main account ID: 0 | 

### Return type

[**InlineResponse20013**](InlineResponse20013.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelMultipleOrders

> InlineResponse20011 CancelMultipleOrders(ctx).Authorization(authorization).AccessId(accessId).BatchIds(batchIds).Market(market).Tonce(tonce).AccountId(accountId).Execute()

Cancel Multiple Orders



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    authorization := "authorization_example" // string | Signature is required for Account API and trading API related interfaces. The signature data is placed in the authorization header of the HTTP header and authorization is the signature result string. No signature is required for market API related interfaces. Use 32-bit MD5 Algorithm Signature Use MD5 algorithm to encrypt the signature string, convert encrypted result to uppercase, get signature data and put signature data in HTTP Header - authorization. 
    accessId := "BFFA64957AA240F6BBEA26F4E07EC0D9" // string | access_id
    batchIds := "123,234,345" // string | Unexecuted multiple order Nos string
    market := "BTCBCH" // string | See <API invocation description market>
    tonce := int64(1513746038205) // int64 | Tonce is a timestamp with a positive Interger that represents the number of milliseconds since Unix epoch to the current time. Error between tonce and server time can not exceed plus or minus 60s
    accountId := int64(789) // int64 | main account ID: 0 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrderApi.CancelMultipleOrders(context.Background()).Authorization(authorization).AccessId(accessId).BatchIds(batchIds).Market(market).Tonce(tonce).AccountId(accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.CancelMultipleOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelMultipleOrders`: InlineResponse20011
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.CancelMultipleOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelMultipleOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Signature is required for Account API and trading API related interfaces. The signature data is placed in the authorization header of the HTTP header and authorization is the signature result string. No signature is required for market API related interfaces. Use 32-bit MD5 Algorithm Signature Use MD5 algorithm to encrypt the signature string, convert encrypted result to uppercase, get signature data and put signature data in HTTP Header - authorization.  | 
 **accessId** | **string** | access_id | 
 **batchIds** | **string** | Unexecuted multiple order Nos string | 
 **market** | **string** | See &lt;API invocation description market&gt; | 
 **tonce** | **int64** | Tonce is a timestamp with a positive Interger that represents the number of milliseconds since Unix epoch to the current time. Error between tonce and server time can not exceed plus or minus 60s | 
 **accountId** | **int64** | main account ID: 0 | 

### Return type

[**InlineResponse20011**](InlineResponse20011.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlaceIocOrder

> InlineResponse20016 PlaceIocOrder(ctx).Authorization(authorization).PlaceIocOrderRequest(placeIocOrderRequest).Execute()

Place IOC Order



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    authorization := "authorization_example" // string | Signature is required for Account API and trading API related interfaces. The signature data is placed in the authorization header of the HTTP header and authorization is the signature result string. No signature is required for market API related interfaces. Use 32-bit MD5 Algorithm Signature Use MD5 algorithm to encrypt the signature string, convert encrypted result to uppercase, get signature data and put signature data in HTTP Header - authorization. 
    placeIocOrderRequest := *openapiclient.NewPlaceIocOrderRequest("AccessId_example", "Market_example", "Type_example") // PlaceIocOrderRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrderApi.PlaceIocOrder(context.Background()).Authorization(authorization).PlaceIocOrderRequest(placeIocOrderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.PlaceIocOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlaceIocOrder`: InlineResponse20016
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.PlaceIocOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlaceIocOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Signature is required for Account API and trading API related interfaces. The signature data is placed in the authorization header of the HTTP header and authorization is the signature result string. No signature is required for market API related interfaces. Use 32-bit MD5 Algorithm Signature Use MD5 algorithm to encrypt the signature string, convert encrypted result to uppercase, get signature data and put signature data in HTTP Header - authorization.  | 
 **placeIocOrderRequest** | [**PlaceIocOrderRequest**](PlaceIocOrderRequest.md) |  | 

### Return type

[**InlineResponse20016**](InlineResponse20016.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlaceLimitOrder

> InlineResponse20014 PlaceLimitOrder(ctx).Authorization(authorization).PlaceLimitOrderRequest(placeLimitOrderRequest).Execute()

Place Limit Order



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    authorization := "authorization_example" // string | Signature is required for Account API and trading API related interfaces. The signature data is placed in the authorization header of the HTTP header and authorization is the signature result string. No signature is required for market API related interfaces. Use 32-bit MD5 Algorithm Signature Use MD5 algorithm to encrypt the signature string, convert encrypted result to uppercase, get signature data and put signature data in HTTP Header - authorization. 
    placeLimitOrderRequest := *openapiclient.NewPlaceLimitOrderRequest("AccessId_example", "Market_example", "Type_example") // PlaceLimitOrderRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrderApi.PlaceLimitOrder(context.Background()).Authorization(authorization).PlaceLimitOrderRequest(placeLimitOrderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.PlaceLimitOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlaceLimitOrder`: InlineResponse20014
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.PlaceLimitOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlaceLimitOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Signature is required for Account API and trading API related interfaces. The signature data is placed in the authorization header of the HTTP header and authorization is the signature result string. No signature is required for market API related interfaces. Use 32-bit MD5 Algorithm Signature Use MD5 algorithm to encrypt the signature string, convert encrypted result to uppercase, get signature data and put signature data in HTTP Header - authorization.  | 
 **placeLimitOrderRequest** | [**PlaceLimitOrderRequest**](PlaceLimitOrderRequest.md) |  | 

### Return type

[**InlineResponse20014**](InlineResponse20014.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlaceMarketOrder

> InlineResponse20015 PlaceMarketOrder(ctx).Authorization(authorization).PlaceMarketOrderRequest(placeMarketOrderRequest).Execute()

Place Market Order



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    authorization := "authorization_example" // string | Signature is required for Account API and trading API related interfaces. The signature data is placed in the authorization header of the HTTP header and authorization is the signature result string. No signature is required for market API related interfaces. Use 32-bit MD5 Algorithm Signature Use MD5 algorithm to encrypt the signature string, convert encrypted result to uppercase, get signature data and put signature data in HTTP Header - authorization. 
    placeMarketOrderRequest := *openapiclient.NewPlaceMarketOrderRequest("AccessId_example", "Market_example", "Type_example") // PlaceMarketOrderRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrderApi.PlaceMarketOrder(context.Background()).Authorization(authorization).PlaceMarketOrderRequest(placeMarketOrderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.PlaceMarketOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlaceMarketOrder`: InlineResponse20015
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.PlaceMarketOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlaceMarketOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Signature is required for Account API and trading API related interfaces. The signature data is placed in the authorization header of the HTTP header and authorization is the signature result string. No signature is required for market API related interfaces. Use 32-bit MD5 Algorithm Signature Use MD5 algorithm to encrypt the signature string, convert encrypted result to uppercase, get signature data and put signature data in HTTP Header - authorization.  | 
 **placeMarketOrderRequest** | [**PlaceMarketOrderRequest**](PlaceMarketOrderRequest.md) |  | 

### Return type

[**InlineResponse20015**](InlineResponse20015.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlaceMultipleLimitOrders

> InlineResponse2009 PlaceMultipleLimitOrders(ctx).Authorization(authorization).PlaceMultipleLimitOrdersRequest(placeMultipleLimitOrdersRequest).Execute()

Place Multiple Limit Orders



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    authorization := "authorization_example" // string | Signature is required for Account API and trading API related interfaces. The signature data is placed in the authorization header of the HTTP header and authorization is the signature result string. No signature is required for market API related interfaces. Use 32-bit MD5 Algorithm Signature Use MD5 algorithm to encrypt the signature string, convert encrypted result to uppercase, get signature data and put signature data in HTTP Header - authorization. 
    placeMultipleLimitOrdersRequest := *openapiclient.NewPlaceMultipleLimitOrdersRequest("AccessId_example") // PlaceMultipleLimitOrdersRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrderApi.PlaceMultipleLimitOrders(context.Background()).Authorization(authorization).PlaceMultipleLimitOrdersRequest(placeMultipleLimitOrdersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.PlaceMultipleLimitOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlaceMultipleLimitOrders`: InlineResponse2009
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.PlaceMultipleLimitOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlaceMultipleLimitOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Signature is required for Account API and trading API related interfaces. The signature data is placed in the authorization header of the HTTP header and authorization is the signature result string. No signature is required for market API related interfaces. Use 32-bit MD5 Algorithm Signature Use MD5 algorithm to encrypt the signature string, convert encrypted result to uppercase, get signature data and put signature data in HTTP Header - authorization.  | 
 **placeMultipleLimitOrdersRequest** | [**PlaceMultipleLimitOrdersRequest**](PlaceMultipleLimitOrdersRequest.md) |  | 

### Return type

[**InlineResponse2009**](InlineResponse2009.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

