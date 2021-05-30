# \MarginApi

All URIs are relative to *https://api.coinex.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcquireLoanList**](MarginApi.md#AcquireLoanList) | **Get** /margin/loan/history | Acquire Loan List
[**AcquireMarginAccountSettings**](MarginApi.md#AcquireMarginAccountSettings) | **Get** /margin/config | Acquire Margin Account Settings
[**InquireMarginAccountInfo**](MarginApi.md#InquireMarginAccountInfo) | **Get** /margin/account | Inquire Margin Account Info
[**PlaceFlat**](MarginApi.md#PlaceFlat) | **Post** /margin/flat | Place Flat
[**PlaceLoan**](MarginApi.md#PlaceLoan) | **Post** /margin/loan | Place Loan



## AcquireLoanList

> InlineResponse20028 AcquireLoanList(ctx).Authorization(authorization).AccessId(accessId).Tonce(tonce).Market(market).Status(status).Page(page).Execute()

Acquire Loan List



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
    accessId := "4DA36FFC61334695A66F8D29020EB589" // string | access_id
    tonce := int64(1513746038205) // int64 | Tonce is a timestamp with a positive Interger that represents the number of milliseconds from Unix epoch to the current time. Error between tonce and server time can not exceed plus or minus 60s
    market := "market_example" // string | See<API invocation description market> (optional)
    status := "status_example" // string | status (optional)
    page := int64(789) // int64 | Page (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginApi.AcquireLoanList(context.Background()).Authorization(authorization).AccessId(accessId).Tonce(tonce).Market(market).Status(status).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.AcquireLoanList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcquireLoanList`: InlineResponse20028
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.AcquireLoanList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAcquireLoanListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Signature is required for Account API and trading API related interfaces. The signature data is placed in the authorization header of the HTTP header and authorization is the signature result string. No signature is required for market API related interfaces. Use 32-bit MD5 Algorithm Signature Use MD5 algorithm to encrypt the signature string, convert encrypted result to uppercase, get signature data and put signature data in HTTP Header - authorization.  | 
 **accessId** | **string** | access_id | 
 **tonce** | **int64** | Tonce is a timestamp with a positive Interger that represents the number of milliseconds from Unix epoch to the current time. Error between tonce and server time can not exceed plus or minus 60s | 
 **market** | **string** | See&lt;API invocation description market&gt; | 
 **status** | **string** | status | 
 **page** | **int64** | Page | 

### Return type

[**InlineResponse20028**](InlineResponse20028.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AcquireMarginAccountSettings

> InlineResponse20027 AcquireMarginAccountSettings(ctx).Authorization(authorization).AccessId(accessId).Tonce(tonce).Execute()

Acquire Margin Account Settings



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
    accessId := "4DA36FFC61334695A66F8D29020EB589" // string | access_id
    tonce := int64(1513746038205) // int64 | Tonce is a timestamp with a positive Interger that represents the number of milliseconds from Unix epoch to the current time. Error between tonce and server time can not exceed plus or minus 60s

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginApi.AcquireMarginAccountSettings(context.Background()).Authorization(authorization).AccessId(accessId).Tonce(tonce).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.AcquireMarginAccountSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcquireMarginAccountSettings`: InlineResponse20027
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.AcquireMarginAccountSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAcquireMarginAccountSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Signature is required for Account API and trading API related interfaces. The signature data is placed in the authorization header of the HTTP header and authorization is the signature result string. No signature is required for market API related interfaces. Use 32-bit MD5 Algorithm Signature Use MD5 algorithm to encrypt the signature string, convert encrypted result to uppercase, get signature data and put signature data in HTTP Header - authorization.  | 
 **accessId** | **string** | access_id | 
 **tonce** | **int64** | Tonce is a timestamp with a positive Interger that represents the number of milliseconds from Unix epoch to the current time. Error between tonce and server time can not exceed plus or minus 60s | 

### Return type

[**InlineResponse20027**](InlineResponse20027.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InquireMarginAccountInfo

> InlineResponse20026 InquireMarginAccountInfo(ctx).Authorization(authorization).AccessId(accessId).Tonce(tonce).Market(market).Execute()

Inquire Margin Account Info



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
    accessId := "4DA36FFC61334695A66F8D29020EB589" // string | access_id
    tonce := int64(1513746038205) // int64 | Tonce is a timestamp with a positive Interger that represents the number of milliseconds from Unix epoch to the current time. Error between tonce and server time can not exceed plus or minus 60s
    market := "market_example" // string | Market Name

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginApi.InquireMarginAccountInfo(context.Background()).Authorization(authorization).AccessId(accessId).Tonce(tonce).Market(market).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.InquireMarginAccountInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InquireMarginAccountInfo`: InlineResponse20026
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.InquireMarginAccountInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInquireMarginAccountInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Signature is required for Account API and trading API related interfaces. The signature data is placed in the authorization header of the HTTP header and authorization is the signature result string. No signature is required for market API related interfaces. Use 32-bit MD5 Algorithm Signature Use MD5 algorithm to encrypt the signature string, convert encrypted result to uppercase, get signature data and put signature data in HTTP Header - authorization.  | 
 **accessId** | **string** | access_id | 
 **tonce** | **int64** | Tonce is a timestamp with a positive Interger that represents the number of milliseconds from Unix epoch to the current time. Error between tonce and server time can not exceed plus or minus 60s | 
 **market** | **string** | Market Name | 

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


## PlaceFlat

> InlineResponse20013 PlaceFlat(ctx).Authorization(authorization).AccessId(accessId).Tonce(tonce).Market(market).CoinType(coinType).Execute()

Place Flat



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
    accessId := "4DA36FFC61334695A66F8D29020EB589" // string | access_id
    tonce := int64(1513746038205) // int64 | Tonce is a timestamp with a positive Interger that represents the number of milliseconds from Unix epoch to the current time. Error between tonce and server time can not exceed plus or minus 60s
    market := "market_example" // string | See<API invocation description market>
    coinType := "coinType_example" // string | Coin type

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginApi.PlaceFlat(context.Background()).Authorization(authorization).AccessId(accessId).Tonce(tonce).Market(market).CoinType(coinType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.PlaceFlat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlaceFlat`: InlineResponse20013
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.PlaceFlat`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlaceFlatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Signature is required for Account API and trading API related interfaces. The signature data is placed in the authorization header of the HTTP header and authorization is the signature result string. No signature is required for market API related interfaces. Use 32-bit MD5 Algorithm Signature Use MD5 algorithm to encrypt the signature string, convert encrypted result to uppercase, get signature data and put signature data in HTTP Header - authorization.  | 
 **accessId** | **string** | access_id | 
 **tonce** | **int64** | Tonce is a timestamp with a positive Interger that represents the number of milliseconds from Unix epoch to the current time. Error between tonce and server time can not exceed plus or minus 60s | 
 **market** | **string** | See&lt;API invocation description market&gt; | 
 **coinType** | **string** | Coin type | 

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


## PlaceLoan

> InlineResponse20029 PlaceLoan(ctx).Authorization(authorization).AccessId(accessId).Tonce(tonce).Market(market).CoinType(coinType).Execute()

Place Loan



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
    accessId := "4DA36FFC61334695A66F8D29020EB589" // string | access_id
    tonce := int64(1513746038205) // int64 | Tonce is a timestamp with a positive Interger that represents the number of milliseconds from Unix epoch to the current time. Error between tonce and server time can not exceed plus or minus 60s
    market := "market_example" // string | See<API invocation description market>
    coinType := "coinType_example" // string | Coin type

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginApi.PlaceLoan(context.Background()).Authorization(authorization).AccessId(accessId).Tonce(tonce).Market(market).CoinType(coinType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.PlaceLoan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlaceLoan`: InlineResponse20029
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.PlaceLoan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlaceLoanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Signature is required for Account API and trading API related interfaces. The signature data is placed in the authorization header of the HTTP header and authorization is the signature result string. No signature is required for market API related interfaces. Use 32-bit MD5 Algorithm Signature Use MD5 algorithm to encrypt the signature string, convert encrypted result to uppercase, get signature data and put signature data in HTTP Header - authorization.  | 
 **accessId** | **string** | access_id | 
 **tonce** | **int64** | Tonce is a timestamp with a positive Interger that represents the number of milliseconds from Unix epoch to the current time. Error between tonce and server time can not exceed plus or minus 60s | 
 **market** | **string** | See&lt;API invocation description market&gt; | 
 **coinType** | **string** | Coin type | 

### Return type

[**InlineResponse20029**](InlineResponse20029.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

