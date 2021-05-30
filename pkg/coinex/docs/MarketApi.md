# \MarketApi

All URIs are relative to *https://api.coinex.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcquireAllMarketData**](MarketApi.md#AcquireAllMarketData) | **Get** /market/ticker/all | Acquire All Market Data
[**AcquireKLineData**](MarketApi.md#AcquireKLineData) | **Get** /market/kline | Acquire K-Line Data
[**AcquireLatestTransactionData**](MarketApi.md#AcquireLatestTransactionData) | **Get** /market/deals | Acquire Latest Transaction Data
[**AcquireMarketDepth**](MarketApi.md#AcquireMarketDepth) | **Get** /market/depth  | Acquire Market Depth
[**AcquireMarketInformation**](MarketApi.md#AcquireMarketInformation) | **Get** /market/info  | Acquire Market Information
[**AcquireMarketList**](MarketApi.md#AcquireMarketList) | **Get** /market/list | Acquire Market List
[**AcquireSingleMarketInformation**](MarketApi.md#AcquireSingleMarketInformation) | **Get** /market/detail  | Acquire Single Market Information



## AcquireAllMarketData

> InlineResponse2001 AcquireAllMarketData(ctx).Execute()

Acquire All Market Data



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarketApi.AcquireAllMarketData(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketApi.AcquireAllMarketData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcquireAllMarketData`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `MarketApi.AcquireAllMarketData`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAcquireAllMarketDataRequest struct via the builder pattern


### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AcquireKLineData

> InlineResponse2004 AcquireKLineData(ctx).Market(market).Limit(limit).Execute()

Acquire K-Line Data



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
    market := "bchbtc" // string | See<API invocation description market>
    limit := int64(789) // int64 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarketApi.AcquireKLineData(context.Background()).Market(market).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketApi.AcquireKLineData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcquireKLineData`: InlineResponse2004
    fmt.Fprintf(os.Stdout, "Response from `MarketApi.AcquireKLineData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAcquireKLineDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **market** | **string** | See&lt;API invocation description market&gt; | 
 **limit** | **int64** |  | 

### Return type

[**InlineResponse2004**](InlineResponse2004.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AcquireLatestTransactionData

> InlineResponse2003 AcquireLatestTransactionData(ctx).Market(market).LastId(lastId).Execute()

Acquire Latest Transaction Data



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
    market := "BCHBTC" // string | See<API invocation description market>
    lastId := int64(2) // int64 | Transaction history id (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarketApi.AcquireLatestTransactionData(context.Background()).Market(market).LastId(lastId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketApi.AcquireLatestTransactionData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcquireLatestTransactionData`: InlineResponse2003
    fmt.Fprintf(os.Stdout, "Response from `MarketApi.AcquireLatestTransactionData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAcquireLatestTransactionDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **market** | **string** | See&lt;API invocation description market&gt; | 
 **lastId** | **int64** | Transaction history id | 

### Return type

[**InlineResponse2003**](InlineResponse2003.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AcquireMarketDepth

> InlineResponse2002 AcquireMarketDepth(ctx).Market(market).Merge(merge).Execute()

Acquire Market Depth



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
    market := "bchbtc" // string | See<API invocation description market>
    merge := "0" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarketApi.AcquireMarketDepth(context.Background()).Market(market).Merge(merge).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketApi.AcquireMarketDepth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcquireMarketDepth`: InlineResponse2002
    fmt.Fprintf(os.Stdout, "Response from `MarketApi.AcquireMarketDepth`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAcquireMarketDepthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **market** | **string** | See&lt;API invocation description market&gt; | 
 **merge** | **string** |  | 

### Return type

[**InlineResponse2002**](InlineResponse2002.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AcquireMarketInformation

> InlineResponse2005 AcquireMarketInformation(ctx).Execute()

Acquire Market Information



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarketApi.AcquireMarketInformation(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketApi.AcquireMarketInformation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcquireMarketInformation`: InlineResponse2005
    fmt.Fprintf(os.Stdout, "Response from `MarketApi.AcquireMarketInformation`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAcquireMarketInformationRequest struct via the builder pattern


### Return type

[**InlineResponse2005**](InlineResponse2005.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AcquireMarketList

> InlineResponse200 AcquireMarketList(ctx).Execute()

Acquire Market List



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarketApi.AcquireMarketList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketApi.AcquireMarketList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcquireMarketList`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `MarketApi.AcquireMarketList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAcquireMarketListRequest struct via the builder pattern


### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AcquireSingleMarketInformation

> InlineResponse2006 AcquireSingleMarketInformation(ctx).Market(market).Execute()

Acquire Single Market Information



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
    market := "BTCUSDT" // string | See<API invocation description market>

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarketApi.AcquireSingleMarketInformation(context.Background()).Market(market).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketApi.AcquireSingleMarketInformation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcquireSingleMarketInformation`: InlineResponse2006
    fmt.Fprintf(os.Stdout, "Response from `MarketApi.AcquireSingleMarketInformation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAcquireSingleMarketInformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **market** | **string** | See&lt;API invocation description market&gt; | 

### Return type

[**InlineResponse2006**](InlineResponse2006.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

