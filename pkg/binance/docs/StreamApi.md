# \StreamApi

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV3UserDataStreamDelete**](StreamApi.md#ApiV3UserDataStreamDelete) | **Delete** /api/v3/userDataStream | Close a ListenKey (USER_STREAM)
[**ApiV3UserDataStreamPost**](StreamApi.md#ApiV3UserDataStreamPost) | **Post** /api/v3/userDataStream | Create a ListenKey (USER_STREAM)
[**ApiV3UserDataStreamPut**](StreamApi.md#ApiV3UserDataStreamPut) | **Put** /api/v3/userDataStream | Ping/Keep-alive a ListenKey (USER_STREAM)



## ApiV3UserDataStreamDelete

> map[string]interface{} ApiV3UserDataStreamDelete(ctx).ListenKey(listenKey).Execute()

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
    resp, r, err := api_client.StreamApi.ApiV3UserDataStreamDelete(context.Background()).ListenKey(listenKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamApi.ApiV3UserDataStreamDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3UserDataStreamDelete`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `StreamApi.ApiV3UserDataStreamDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3UserDataStreamDeleteRequest struct via the builder pattern


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


## ApiV3UserDataStreamPost

> InlineResponse20026 ApiV3UserDataStreamPost(ctx).Execute()

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
    resp, r, err := api_client.StreamApi.ApiV3UserDataStreamPost(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamApi.ApiV3UserDataStreamPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3UserDataStreamPost`: InlineResponse20026
    fmt.Fprintf(os.Stdout, "Response from `StreamApi.ApiV3UserDataStreamPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3UserDataStreamPostRequest struct via the builder pattern


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


## ApiV3UserDataStreamPut

> map[string]interface{} ApiV3UserDataStreamPut(ctx).ListenKey(listenKey).Execute()

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
    resp, r, err := api_client.StreamApi.ApiV3UserDataStreamPut(context.Background()).ListenKey(listenKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamApi.ApiV3UserDataStreamPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3UserDataStreamPut`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `StreamApi.ApiV3UserDataStreamPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3UserDataStreamPutRequest struct via the builder pattern


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

