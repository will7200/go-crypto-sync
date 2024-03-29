# \AccountsApi

All URIs are relative to *https://api.coinbase.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccount**](AccountsApi.md#GetAccount) | **Get** /accounts/{account_id} | Retrieve Account
[**ListAccounts**](AccountsApi.md#ListAccounts) | **Get** /accounts | List Accounts
[**ListTransactions**](AccountsApi.md#ListTransactions) | **Get** /accounts/{account_id}/transactions | List transactions



## GetAccount

> Account GetAccount(ctx, accountId).CBACCESSSIGN(cBACCESSSIGN).CBACCESSTIMESTAMP(cBACCESSTIMESTAMP).Execute()

Retrieve Account



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/will7200/go-crypto-sync/pkg/coinbase"
)

func main() {
    accountId := "accountId_example" // string | The id of the account
    cBACCESSSIGN := "cBACCESSSIGN_example" // string | The user generated message signature (see below) The CB-ACCESS-SIGN header is generated by creating a sha256 HMAC using the secret key on the prehash string timestamp + method + requestPath + body (where + represents string concatenation).  (optional) (default to "auto")
    cBACCESSTIMESTAMP := int32(56) // int32 | A timestamp for your request (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.GetAccount(context.Background(), accountId).CBACCESSSIGN(cBACCESSSIGN).CBACCESSTIMESTAMP(cBACCESSTIMESTAMP).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.GetAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccount`: Account
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.GetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The id of the account | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cBACCESSSIGN** | **string** | The user generated message signature (see below) The CB-ACCESS-SIGN header is generated by creating a sha256 HMAC using the secret key on the prehash string timestamp + method + requestPath + body (where + represents string concatenation).  | [default to &quot;auto&quot;]
 **cBACCESSTIMESTAMP** | **int32** | A timestamp for your request | 

### Return type

[**Account**](Account.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccounts

> InlineResponse200 ListAccounts(ctx).CBACCESSSIGN(cBACCESSSIGN).CBACCESSTIMESTAMP(cBACCESSTIMESTAMP).EndingBefore(endingBefore).StartingAfter(startingAfter).Limit(limit).Order(order).PreviousUri(previousUri).NextUri(nextUri).Execute()

List Accounts



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/will7200/go-crypto-sync/pkg/coinbase"
)

func main() {
    cBACCESSSIGN := "cBACCESSSIGN_example" // string | The user generated message signature (see below) The CB-ACCESS-SIGN header is generated by creating a sha256 HMAC using the secret key on the prehash string timestamp + method + requestPath + body (where + represents string concatenation).  (optional) (default to "auto")
    cBACCESSTIMESTAMP := int32(56) // int32 | A timestamp for your request (optional)
    endingBefore := "endingBefore_example" // string |  (optional)
    startingAfter := "startingAfter_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional) (default to 25)
    order := "order_example" // string |  (optional) (default to "desc")
    previousUri := "previousUri_example" // string |  (optional)
    nextUri := "nextUri_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.ListAccounts(context.Background(), ).CBACCESSSIGN(cBACCESSSIGN).CBACCESSTIMESTAMP(cBACCESSTIMESTAMP).EndingBefore(endingBefore).StartingAfter(startingAfter).Limit(limit).Order(order).PreviousUri(previousUri).NextUri(nextUri).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.ListAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccounts`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.ListAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cBACCESSSIGN** | **string** | The user generated message signature (see below) The CB-ACCESS-SIGN header is generated by creating a sha256 HMAC using the secret key on the prehash string timestamp + method + requestPath + body (where + represents string concatenation).  | [default to &quot;auto&quot;]
 **cBACCESSTIMESTAMP** | **int32** | A timestamp for your request | 
 **endingBefore** | **string** |  | 
 **startingAfter** | **string** |  | 
 **limit** | **int32** |  | [default to 25]
 **order** | **string** |  | [default to &quot;desc&quot;]
 **previousUri** | **string** |  | 
 **nextUri** | **string** |  | 

### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTransactions

> InlineResponse2001 ListTransactions(ctx, accountId).CBACCESSSIGN(cBACCESSSIGN).CBACCESSTIMESTAMP(cBACCESSTIMESTAMP).Execute()

List transactions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/will7200/go-crypto-sync/pkg/coinbase"
)

func main() {
    accountId := "accountId_example" // string | The id of the account
    cBACCESSSIGN := "cBACCESSSIGN_example" // string | The user generated message signature (see below) The CB-ACCESS-SIGN header is generated by creating a sha256 HMAC using the secret key on the prehash string timestamp + method + requestPath + body (where + represents string concatenation).  (optional) (default to "auto")
    cBACCESSTIMESTAMP := int32(56) // int32 | A timestamp for your request (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.ListTransactions(context.Background(), accountId).CBACCESSSIGN(cBACCESSSIGN).CBACCESSTIMESTAMP(cBACCESSTIMESTAMP).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.ListTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTransactions`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.ListTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The id of the account | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cBACCESSSIGN** | **string** | The user generated message signature (see below) The CB-ACCESS-SIGN header is generated by creating a sha256 HMAC using the secret key on the prehash string timestamp + method + requestPath + body (where + represents string concatenation).  | [default to &quot;auto&quot;]
 **cBACCESSTIMESTAMP** | **int32** | A timestamp for your request | 

### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

