# \PredictionsApi

All URIs are relative to *https://api.nomics.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCurrenciesPredictionsHistory**](PredictionsApi.md#GetCurrenciesPredictionsHistory) | **Get** /currencies/predictions/history | Currency Predictions History
[**GetCurrenciesPredictionsTicker**](PredictionsApi.md#GetCurrenciesPredictionsTicker) | **Get** /currencies/predictions/ticker | Currency Predictions Ticker



## GetCurrenciesPredictionsHistory

> InlineResponse20022 GetCurrenciesPredictionsHistory(ctx).Id(id).Interval(interval).Execute()

Currency Predictions History



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/will7200/go-crypto-sync/pkg/nomics"
)

func main() {
    id := "id_example" // string | Nomics currency ID (optional)
    interval := "interval_example" // string | Prediction interval (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PredictionsApi.GetCurrenciesPredictionsHistory(context.Background(), ).Id(id).Interval(interval).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PredictionsApi.GetCurrenciesPredictionsHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrenciesPredictionsHistory`: InlineResponse20022
    fmt.Fprintf(os.Stdout, "Response from `PredictionsApi.GetCurrenciesPredictionsHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrenciesPredictionsHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Nomics currency ID | 
 **interval** | **string** | Prediction interval | 

### Return type

[**InlineResponse20022**](inline_response_200_22.md)

### Authorization

[Key](../README.md#Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrenciesPredictionsTicker

> []InlineResponse20021 GetCurrenciesPredictionsTicker(ctx).Ids(ids).Execute()

Currency Predictions Ticker



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/will7200/go-crypto-sync/pkg/nomics"
)

func main() {
    ids := "ids_example" // string | Comma separated list of Nomics currency IDs (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PredictionsApi.GetCurrenciesPredictionsTicker(context.Background(), ).Ids(ids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PredictionsApi.GetCurrenciesPredictionsTicker``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrenciesPredictionsTicker`: []InlineResponse20021
    fmt.Fprintf(os.Stdout, "Response from `PredictionsApi.GetCurrenciesPredictionsTicker`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrenciesPredictionsTickerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** | Comma separated list of Nomics currency IDs | 

### Return type

[**[]InlineResponse20021**](inline_response_200_21.md)

### Authorization

[Key](../README.md#Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

