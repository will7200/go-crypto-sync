# \MarginStreamApi

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SapiV1UserDataStreamDelete**](MarginStreamApi.md#SapiV1UserDataStreamDelete) | **Delete** /sapi/v1/userDataStream | Close a ListenKey (USER_STREAM)
[**SapiV1UserDataStreamPost**](MarginStreamApi.md#SapiV1UserDataStreamPost) | **Post** /sapi/v1/userDataStream | Create a ListenKey (USER_STREAM)
[**SapiV1UserDataStreamPut**](MarginStreamApi.md#SapiV1UserDataStreamPut) | **Put** /sapi/v1/userDataStream | Ping/Keep-alive a ListenKey (USER_STREAM)



## SapiV1UserDataStreamDelete

> map[string]interface{} SapiV1UserDataStreamDelete(ctx).ListenKey(listenKey).Execute()

Close a ListenKey (USER_STREAM)



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
    listenKey := "pqia91ma19a5s61cv6a81va65sdf19v8a65a1a5s61cv6a81va65sdf19v8a65a1" // string | User websocket listen key (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginStreamApi.SapiV1UserDataStreamDelete(context.Background()).ListenKey(listenKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginStreamApi.SapiV1UserDataStreamDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1UserDataStreamDelete`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MarginStreamApi.SapiV1UserDataStreamDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1UserDataStreamDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listenKey** | **string** | User websocket listen key | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1UserDataStreamPost

> InlineResponse20026 SapiV1UserDataStreamPost(ctx).Execute()

Create a ListenKey (USER_STREAM)



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
    resp, r, err := api_client.MarginStreamApi.SapiV1UserDataStreamPost(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginStreamApi.SapiV1UserDataStreamPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1UserDataStreamPost`: InlineResponse20026
    fmt.Fprintf(os.Stdout, "Response from `MarginStreamApi.SapiV1UserDataStreamPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1UserDataStreamPostRequest struct via the builder pattern


### Return type

[**InlineResponse20026**](InlineResponse20026.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1UserDataStreamPut

> map[string]interface{} SapiV1UserDataStreamPut(ctx).ListenKey(listenKey).Execute()

Ping/Keep-alive a ListenKey (USER_STREAM)



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
    listenKey := "pqia91ma19a5s61cv6a81va65sdf19v8a65a1a5s61cv6a81va65sdf19v8a65a1" // string | User websocket listen key (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginStreamApi.SapiV1UserDataStreamPut(context.Background()).ListenKey(listenKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginStreamApi.SapiV1UserDataStreamPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1UserDataStreamPut`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MarginStreamApi.SapiV1UserDataStreamPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1UserDataStreamPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listenKey** | **string** | User websocket listen key | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

