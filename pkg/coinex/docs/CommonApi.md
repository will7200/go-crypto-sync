# \CommonApi

All URIs are relative to *https://api.coinex.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcquireAssetConfig**](CommonApi.md#AcquireAssetConfig) | **Get** /common/asset/config | Acquire Asset Config
[**AcquireCurrencyRate**](CommonApi.md#AcquireCurrencyRate) | **Get** /common/currency/rate | Acquire Currency Rate



## AcquireAssetConfig

> InlineResponse20025 AcquireAssetConfig(ctx).CoinType(coinType).Execute()

Acquire Asset Config



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
    coinType := "coinType_example" // string | Coin type (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommonApi.AcquireAssetConfig(context.Background()).CoinType(coinType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommonApi.AcquireAssetConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcquireAssetConfig`: InlineResponse20025
    fmt.Fprintf(os.Stdout, "Response from `CommonApi.AcquireAssetConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAcquireAssetConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **coinType** | **string** | Coin type | 

### Return type

[**InlineResponse20025**](InlineResponse20025.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AcquireCurrencyRate

> InlineResponse20024 AcquireCurrencyRate(ctx).Execute()

Acquire Currency Rate



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
    resp, r, err := api_client.CommonApi.AcquireCurrencyRate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommonApi.AcquireCurrencyRate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcquireCurrencyRate`: InlineResponse20024
    fmt.Fprintf(os.Stdout, "Response from `CommonApi.AcquireCurrencyRate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAcquireCurrencyRateRequest struct via the builder pattern


### Return type

[**InlineResponse20024**](InlineResponse20024.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

