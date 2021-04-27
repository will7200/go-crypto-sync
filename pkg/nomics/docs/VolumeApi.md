# \VolumeApi

All URIs are relative to *https://api.nomics.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVolumeHistory**](VolumeApi.md#GetVolumeHistory) | **Get** /volume/history | Global Volume History



## GetVolumeHistory

> []InlineResponse20016 GetVolumeHistory(ctx).Start(start).End(end).Convert(convert).Format(format).IncludeTransparency(includeTransparency).Execute()

Global Volume History



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
    start := "start_example" // string | Start time of the interval in RFC3339 (URI escaped) (optional)
    end := "end_example" // string | End time of the interval in RFC3339 (URI escaped). If not provided, the current time is used.  (optional)
    convert := "convert_example" // string | Currency to quote volume values. May be a Fiat Currency or Cryptocurrency. Default is `USD`. (optional)
    format := "format_example" // string | Format of the response. Defaults to JSON when blank. (optional)
    includeTransparency := false // bool | Whether to include [Transparent Volume](https://blog.nomics.com/essays/transparent-volume/) and transparent market cap information in the response. Default is `false`. This option is only available to customers of our paid API plans.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VolumeApi.GetVolumeHistory(context.Background(), ).Start(start).End(end).Convert(convert).Format(format).IncludeTransparency(includeTransparency).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumeApi.GetVolumeHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVolumeHistory`: []InlineResponse20016
    fmt.Fprintf(os.Stdout, "Response from `VolumeApi.GetVolumeHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVolumeHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** | Start time of the interval in RFC3339 (URI escaped) | 
 **end** | **string** | End time of the interval in RFC3339 (URI escaped). If not provided, the current time is used.  | 
 **convert** | **string** | Currency to quote volume values. May be a Fiat Currency or Cryptocurrency. Default is &#x60;USD&#x60;. | 
 **format** | **string** | Format of the response. Defaults to JSON when blank. | 
 **includeTransparency** | **bool** | Whether to include [Transparent Volume](https://blog.nomics.com/essays/transparent-volume/) and transparent market cap information in the response. Default is &#x60;false&#x60;. This option is only available to customers of our paid API plans.  | 

### Return type

[**[]InlineResponse20016**](inline_response_200_16.md)

### Authorization

[Key](../README.md#Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, test/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

