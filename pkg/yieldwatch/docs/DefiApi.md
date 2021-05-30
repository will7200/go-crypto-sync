# \DefiApi

All URIs are relative to *https://www.yieldwatch.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDefiPortfolioAll**](DefiApi.md#GetDefiPortfolioAll) | **Get** /all/{wallet_address} | Get all providers for the waller address



## GetDefiPortfolioAll

> InlineResponse200 GetDefiPortfolioAll(ctx, walletAddress).Platforms(platforms).Execute()

Get all providers for the waller address



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
    walletAddress := "walletAddress_example" // string | wallet address to return values for
    platforms := "pancake,hyperjump,blizzard" // string | comma seperated list of platforms to speed up request (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefiApi.GetDefiPortfolioAll(context.Background(), walletAddress).Platforms(platforms).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefiApi.GetDefiPortfolioAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDefiPortfolioAll`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `DefiApi.GetDefiPortfolioAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletAddress** | **string** | wallet address to return values for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefiPortfolioAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **platforms** | **string** | comma seperated list of platforms to speed up request | 

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

