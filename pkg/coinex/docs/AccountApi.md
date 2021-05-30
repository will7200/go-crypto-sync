# \AccountApi

All URIs are relative to *https://api.coinex.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelWithdrawal**](AccountApi.md#CancelWithdrawal) | **Delete** /balance/coin/withdraw | Cancel Withdrawal
[**InquireAccountInfo**](AccountApi.md#InquireAccountInfo) | **Get** /balance/info | Inquire Account Info
[**InquireDepositList**](AccountApi.md#InquireDepositList) | **Get** /balance/coin/deposit | Inquire Deposit List 
[**InquireMarginAccountMarketInfo**](AccountApi.md#InquireMarginAccountMarketInfo) | **Get** /margin/market | Inquire Margin Account Market Info
[**InquireWithdrawalList**](AccountApi.md#InquireWithdrawalList) | **Get** /balance/coin/withdraw | Inquire Withdrawal List
[**SubmitWithdrawalOrder**](AccountApi.md#SubmitWithdrawalOrder) | **Post** /balance/coin/withdraw | Submit Withdrawal Order
[**TransferBetweenMainAccountAndMarginAccount**](AccountApi.md#TransferBetweenMainAccountAndMarginAccount) | **Post** /margin/transfer | Transfer between main account and margin account
[**TransferBetweenMainAccountAndSubAccount**](AccountApi.md#TransferBetweenMainAccountAndSubAccount) | **Post** /sub_account/transfer | Transfer between main account and sub account



## CancelWithdrawal

> InlineResponse20013 CancelWithdrawal(ctx).Authorization(authorization).AccessId(accessId).Tonce(tonce).CoinWithdrawId(coinWithdrawId).Execute()

Cancel Withdrawal



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
    coinWithdrawId := int64(206) // int64 | coin withdrawal id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.CancelWithdrawal(context.Background()).Authorization(authorization).AccessId(accessId).Tonce(tonce).CoinWithdrawId(coinWithdrawId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.CancelWithdrawal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelWithdrawal`: InlineResponse20013
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.CancelWithdrawal`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelWithdrawalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Signature is required for Account API and trading API related interfaces. The signature data is placed in the authorization header of the HTTP header and authorization is the signature result string. No signature is required for market API related interfaces. Use 32-bit MD5 Algorithm Signature Use MD5 algorithm to encrypt the signature string, convert encrypted result to uppercase, get signature data and put signature data in HTTP Header - authorization.  | 
 **accessId** | **string** | access_id | 
 **tonce** | **int64** | Tonce is a timestamp with a positive Interger that represents the number of milliseconds from Unix epoch to the current time. Error between tonce and server time can not exceed plus or minus 60s | 
 **coinWithdrawId** | **int64** | coin withdrawal id | 

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


## InquireAccountInfo

> InlineResponse20020 InquireAccountInfo(ctx).Authorization(authorization).AccessId(accessId).Tonce(tonce).Execute()

Inquire Account Info



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
    resp, r, err := api_client.AccountApi.InquireAccountInfo(context.Background()).Authorization(authorization).AccessId(accessId).Tonce(tonce).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.InquireAccountInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InquireAccountInfo`: InlineResponse20020
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.InquireAccountInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInquireAccountInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Signature is required for Account API and trading API related interfaces. The signature data is placed in the authorization header of the HTTP header and authorization is the signature result string. No signature is required for market API related interfaces. Use 32-bit MD5 Algorithm Signature Use MD5 algorithm to encrypt the signature string, convert encrypted result to uppercase, get signature data and put signature data in HTTP Header - authorization.  | 
 **accessId** | **string** | access_id | 
 **tonce** | **int64** | Tonce is a timestamp with a positive Interger that represents the number of milliseconds from Unix epoch to the current time. Error between tonce and server time can not exceed plus or minus 60s | 

### Return type

[**InlineResponse20020**](InlineResponse20020.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InquireDepositList

> InlineResponse20023 InquireDepositList(ctx).Authorization(authorization).AccessId(accessId).Tonce(tonce).CoinType(coinType).Execute()

Inquire Deposit List 



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
    accessId := "XXX6FXXXXXXXXXX902XXXXX89" // string | access_id
    tonce := int64(1513746038205) // int64 | Tonce is a timestamp with a positive Interger that represents the number of milliseconds from Unix epoch to the current time. Error between tonce and server time can not exceed plus or minus 60s
    coinType := "coinType_example" // string | Coin type (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.InquireDepositList(context.Background()).Authorization(authorization).AccessId(accessId).Tonce(tonce).CoinType(coinType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.InquireDepositList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InquireDepositList`: InlineResponse20023
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.InquireDepositList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInquireDepositListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Signature is required for Account API and trading API related interfaces. The signature data is placed in the authorization header of the HTTP header and authorization is the signature result string. No signature is required for market API related interfaces. Use 32-bit MD5 Algorithm Signature Use MD5 algorithm to encrypt the signature string, convert encrypted result to uppercase, get signature data and put signature data in HTTP Header - authorization.  | 
 **accessId** | **string** | access_id | 
 **tonce** | **int64** | Tonce is a timestamp with a positive Interger that represents the number of milliseconds from Unix epoch to the current time. Error between tonce and server time can not exceed plus or minus 60s | 
 **coinType** | **string** | Coin type | 

### Return type

[**InlineResponse20023**](InlineResponse20023.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InquireMarginAccountMarketInfo

> InlineResponse20013 InquireMarginAccountMarketInfo(ctx).Execute()

Inquire Margin Account Market Info



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
    resp, r, err := api_client.AccountApi.InquireMarginAccountMarketInfo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.InquireMarginAccountMarketInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InquireMarginAccountMarketInfo`: InlineResponse20013
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.InquireMarginAccountMarketInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInquireMarginAccountMarketInfoRequest struct via the builder pattern


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


## InquireWithdrawalList

> InlineResponse20021 InquireWithdrawalList(ctx).Authorization(authorization).AccessId(accessId).Tonce(tonce).CoinType(coinType).Execute()

Inquire Withdrawal List



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
    coinType := "BCH" // string | Coin type (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.InquireWithdrawalList(context.Background()).Authorization(authorization).AccessId(accessId).Tonce(tonce).CoinType(coinType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.InquireWithdrawalList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InquireWithdrawalList`: InlineResponse20021
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.InquireWithdrawalList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInquireWithdrawalListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Signature is required for Account API and trading API related interfaces. The signature data is placed in the authorization header of the HTTP header and authorization is the signature result string. No signature is required for market API related interfaces. Use 32-bit MD5 Algorithm Signature Use MD5 algorithm to encrypt the signature string, convert encrypted result to uppercase, get signature data and put signature data in HTTP Header - authorization.  | 
 **accessId** | **string** | access_id | 
 **tonce** | **int64** | Tonce is a timestamp with a positive Interger that represents the number of milliseconds from Unix epoch to the current time. Error between tonce and server time can not exceed plus or minus 60s | 
 **coinType** | **string** | Coin type | 

### Return type

[**InlineResponse20021**](InlineResponse20021.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitWithdrawalOrder

> InlineResponse20022 SubmitWithdrawalOrder(ctx).Authorization(authorization).AccessId(accessId).Tonce(tonce).CoinType(coinType).SmartContractName(smartContractName).Execute()

Submit Withdrawal Order



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
    coinType := "coinType_example" // string | coin_type
    smartContractName := "smartContractName_example" // string | Multi-protocol USDT parameter: ERC20 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.SubmitWithdrawalOrder(context.Background()).Authorization(authorization).AccessId(accessId).Tonce(tonce).CoinType(coinType).SmartContractName(smartContractName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.SubmitWithdrawalOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitWithdrawalOrder`: InlineResponse20022
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.SubmitWithdrawalOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitWithdrawalOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Signature is required for Account API and trading API related interfaces. The signature data is placed in the authorization header of the HTTP header and authorization is the signature result string. No signature is required for market API related interfaces. Use 32-bit MD5 Algorithm Signature Use MD5 algorithm to encrypt the signature string, convert encrypted result to uppercase, get signature data and put signature data in HTTP Header - authorization.  | 
 **accessId** | **string** | access_id | 
 **tonce** | **int64** | Tonce is a timestamp with a positive Interger that represents the number of milliseconds from Unix epoch to the current time. Error between tonce and server time can not exceed plus or minus 60s | 
 **coinType** | **string** | coin_type | 
 **smartContractName** | **string** | Multi-protocol USDT parameter: ERC20 | 

### Return type

[**InlineResponse20022**](InlineResponse20022.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransferBetweenMainAccountAndMarginAccount

> InlineResponse20013 TransferBetweenMainAccountAndMarginAccount(ctx).FromAccount(fromAccount).Execute()

Transfer between main account and margin account



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
    fromAccount := int64(789) // int64 | the remitting account ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.TransferBetweenMainAccountAndMarginAccount(context.Background()).FromAccount(fromAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.TransferBetweenMainAccountAndMarginAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransferBetweenMainAccountAndMarginAccount`: InlineResponse20013
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.TransferBetweenMainAccountAndMarginAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransferBetweenMainAccountAndMarginAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fromAccount** | **int64** | the remitting account ID | 

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


## TransferBetweenMainAccountAndSubAccount

> InlineResponse20013 TransferBetweenMainAccountAndSubAccount(ctx).AccessId(accessId).Tonce(tonce).TransferAccount(transferAccount).TransferSide(transferSide).Execute()

Transfer between main account and sub account



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
    accessId := "4DA36FFC61334695A66F8D29020EB589" // string | access_id
    tonce := int64(1513746038205) // int64 | Tonce is a timestamp with a positive Interger that represents the number of milliseconds from Unix epoch to the current time. Error between tonce and server time can not exceed plus or minus 60s
    transferAccount := "transferAccount_example" // string | the sub account name (optional)
    transferSide := "transferSide_example" // string | to post (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.TransferBetweenMainAccountAndSubAccount(context.Background()).AccessId(accessId).Tonce(tonce).TransferAccount(transferAccount).TransferSide(transferSide).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.TransferBetweenMainAccountAndSubAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransferBetweenMainAccountAndSubAccount`: InlineResponse20013
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.TransferBetweenMainAccountAndSubAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransferBetweenMainAccountAndSubAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessId** | **string** | access_id | 
 **tonce** | **int64** | Tonce is a timestamp with a positive Interger that represents the number of milliseconds from Unix epoch to the current time. Error between tonce and server time can not exceed plus or minus 60s | 
 **transferAccount** | **string** | the sub account name | 
 **transferSide** | **string** | to post | 

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

