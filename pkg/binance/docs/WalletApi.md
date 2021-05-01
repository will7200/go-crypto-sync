# \WalletApi

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SapiV1AccountDisableFastWithdrawSwitchPost**](WalletApi.md#SapiV1AccountDisableFastWithdrawSwitchPost) | **Post** /sapi/v1/account/disableFastWithdrawSwitch | Disable Fast Withdraw Switch (USER_DATA)
[**SapiV1AccountEnableFastWithdrawSwitchPost**](WalletApi.md#SapiV1AccountEnableFastWithdrawSwitchPost) | **Post** /sapi/v1/account/enableFastWithdrawSwitch | Enable Fast Withdraw Switch (USER_DATA)
[**SapiV1AccountSnapshotGet**](WalletApi.md#SapiV1AccountSnapshotGet) | **Get** /sapi/v1/accountSnapshot | Daily Account Snapshot (USER_DATA)
[**SapiV1AssetAssetDividendGet**](WalletApi.md#SapiV1AssetAssetDividendGet) | **Get** /sapi/v1/asset/assetDividend | Dust Transfer
[**SapiV1AssetDustPost**](WalletApi.md#SapiV1AssetDustPost) | **Post** /sapi/v1/asset/dust | Dust Transfer
[**SapiV1CapitalConfigGetallGet**](WalletApi.md#SapiV1CapitalConfigGetallGet) | **Get** /sapi/v1/capital/config/getall | All Coins&#39; Information (USER_DATA)
[**SapiV1CapitalDepositAddressGet**](WalletApi.md#SapiV1CapitalDepositAddressGet) | **Get** /sapi/v1/capital/deposit/address | Withdraw History (supporting network) (USER_DATA)
[**SapiV1CapitalDepositHisrecGet**](WalletApi.md#SapiV1CapitalDepositHisrecGet) | **Get** /sapi/v1/capital/deposit/hisrec | Deposit History（supporting network） (USER_DATA)
[**SapiV1CapitalWithdrawApplyPost**](WalletApi.md#SapiV1CapitalWithdrawApplyPost) | **Post** /sapi/v1/capital/withdraw/apply | Withdraw
[**SapiV1CapitalWithdrawHistoryGet**](WalletApi.md#SapiV1CapitalWithdrawHistoryGet) | **Get** /sapi/v1/capital/withdraw/history | Withdraw History (supporting network) (USER_DATA)
[**WapiV3AccountStatusHtmlGet**](WalletApi.md#WapiV3AccountStatusHtmlGet) | **Get** /wapi/v3/accountStatus.html | Account Status (USER_DATA)
[**WapiV3ApiTradingStatusHtmlGet**](WalletApi.md#WapiV3ApiTradingStatusHtmlGet) | **Get** /wapi/v3/apiTradingStatus.html | Account API Trading Status (USER_DATA)
[**WapiV3AssetDetailHtmlGet**](WalletApi.md#WapiV3AssetDetailHtmlGet) | **Get** /wapi/v3/assetDetail.html | Asset Detail (USER_DATA)
[**WapiV3SystemStatusHtmlGet**](WalletApi.md#WapiV3SystemStatusHtmlGet) | **Get** /wapi/v3/systemStatus.html | System Status (System)
[**WapiV3TradeFeeHtmlGet**](WalletApi.md#WapiV3TradeFeeHtmlGet) | **Get** /wapi/v3/tradeFee.html | Trade Fee (USER_DATA)
[**WapiV3UserAssetDribbletLogHtmlGet**](WalletApi.md#WapiV3UserAssetDribbletLogHtmlGet) | **Get** /wapi/v3/userAssetDribbletLog.html | DustLog (USER_DATA)



## SapiV1AccountDisableFastWithdrawSwitchPost

> SapiV1AccountDisableFastWithdrawSwitchPost(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Signature(signature).Execute()

Disable Fast Withdraw Switch (USER_DATA)



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
    timestamp := int32(56) // int32 | UTC timestamp (optional)
    recvWindow := int32(5000) // int32 | The value cannot be greater than 60000 (optional)
    signature := "signature_example" // string | signature (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WalletApi.SapiV1AccountDisableFastWithdrawSwitchPost(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Signature(signature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.SapiV1AccountDisableFastWithdrawSwitchPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1AccountDisableFastWithdrawSwitchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int32** | UTC timestamp | 
 **recvWindow** | **int32** | The value cannot be greater than 60000 | 
 **signature** | **string** | signature | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1AccountEnableFastWithdrawSwitchPost

> SapiV1AccountEnableFastWithdrawSwitchPost(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Signature(signature).Execute()

Enable Fast Withdraw Switch (USER_DATA)



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
    timestamp := int32(56) // int32 | UTC timestamp (optional)
    recvWindow := int32(5000) // int32 | The value cannot be greater than 60000 (optional)
    signature := "signature_example" // string | signature (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WalletApi.SapiV1AccountEnableFastWithdrawSwitchPost(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Signature(signature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.SapiV1AccountEnableFastWithdrawSwitchPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1AccountEnableFastWithdrawSwitchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int32** | UTC timestamp | 
 **recvWindow** | **int32** | The value cannot be greater than 60000 | 
 **signature** | **string** | signature | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1AccountSnapshotGet

> InlineResponse20014 SapiV1AccountSnapshotGet(ctx).Type_(type_).Limit(limit).StartTime(startTime).EndTime(endTime).Timestamp(timestamp).RecvWindow(recvWindow).Signature(signature).Execute()

Daily Account Snapshot (USER_DATA)



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
    type_ := "type__example" // string | 
    limit := int32(56) // int32 |  (default to 5)
    startTime := int32(56) // int32 | Timestamp in ms (optional)
    endTime := int32(56) // int32 | Timestamp in ms (optional)
    timestamp := int32(56) // int32 | UTC timestamp (optional)
    recvWindow := int32(5000) // int32 | The value cannot be greater than 60000 (optional)
    signature := "signature_example" // string | signature (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WalletApi.SapiV1AccountSnapshotGet(context.Background()).Type_(type_).Limit(limit).StartTime(startTime).EndTime(endTime).Timestamp(timestamp).RecvWindow(recvWindow).Signature(signature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.SapiV1AccountSnapshotGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1AccountSnapshotGet`: InlineResponse20014
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.SapiV1AccountSnapshotGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1AccountSnapshotGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** |  | 
 **limit** | **int32** |  | [default to 5]
 **startTime** | **int32** | Timestamp in ms | 
 **endTime** | **int32** | Timestamp in ms | 
 **timestamp** | **int32** | UTC timestamp | 
 **recvWindow** | **int32** | The value cannot be greater than 60000 | 
 **signature** | **string** | signature | 

### Return type

[**InlineResponse20014**](InlineResponse20014.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1AssetAssetDividendGet

> InlineResponse20023 SapiV1AssetAssetDividendGet(ctx).Limit(limit).Asset(asset).StartTime(startTime).EndTime(endTime).Timestamp(timestamp).RecvWindow(recvWindow).Signature(signature).Execute()

Dust Transfer



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
    limit := "limit_example" // string |  (default to "20")
    asset := "BNB" // string | The asset being transferred, e.g., BTC (optional)
    startTime := int32(56) // int32 | Timestamp in ms (optional)
    endTime := int32(56) // int32 | Timestamp in ms (optional)
    timestamp := int32(56) // int32 | UTC timestamp (optional)
    recvWindow := int32(5000) // int32 | The value cannot be greater than 60000 (optional)
    signature := "signature_example" // string | signature (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WalletApi.SapiV1AssetAssetDividendGet(context.Background()).Limit(limit).Asset(asset).StartTime(startTime).EndTime(endTime).Timestamp(timestamp).RecvWindow(recvWindow).Signature(signature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.SapiV1AssetAssetDividendGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1AssetAssetDividendGet`: InlineResponse20023
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.SapiV1AssetAssetDividendGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1AssetAssetDividendGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **string** |  | [default to &quot;20&quot;]
 **asset** | **string** | The asset being transferred, e.g., BTC | 
 **startTime** | **int32** | Timestamp in ms | 
 **endTime** | **int32** | Timestamp in ms | 
 **timestamp** | **int32** | UTC timestamp | 
 **recvWindow** | **int32** | The value cannot be greater than 60000 | 
 **signature** | **string** | signature | 

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


## SapiV1AssetDustPost

> InlineResponse20022 SapiV1AssetDustPost(ctx).Asset(asset).Timestamp(timestamp).RecvWindow(recvWindow).Signature(signature).Execute()

Dust Transfer



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
    asset := "asset=BTC&asset=USDT" // string | The asset being converted. For example, asset=BTC&asset=USDT
    timestamp := int32(56) // int32 | UTC timestamp (optional)
    recvWindow := int32(5000) // int32 | The value cannot be greater than 60000 (optional)
    signature := "signature_example" // string | signature (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WalletApi.SapiV1AssetDustPost(context.Background()).Asset(asset).Timestamp(timestamp).RecvWindow(recvWindow).Signature(signature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.SapiV1AssetDustPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1AssetDustPost`: InlineResponse20022
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.SapiV1AssetDustPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1AssetDustPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** | The asset being converted. For example, asset&#x3D;BTC&amp;asset&#x3D;USDT | 
 **timestamp** | **int32** | UTC timestamp | 
 **recvWindow** | **int32** | The value cannot be greater than 60000 | 
 **signature** | **string** | signature | 

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


## SapiV1CapitalConfigGetallGet

> InlineResponse20013 SapiV1CapitalConfigGetallGet(ctx).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()

All Coins' Information (USER_DATA)



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
    recvWindow := int32(5000) // int32 | The value cannot be greater than 60000 (optional)
    timestamp := int32(56) // int32 | UTC timestamp (optional)
    signature := "signature_example" // string | signature (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WalletApi.SapiV1CapitalConfigGetallGet(context.Background()).RecvWindow(recvWindow).Timestamp(timestamp).Signature(signature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.SapiV1CapitalConfigGetallGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1CapitalConfigGetallGet`: InlineResponse20013
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.SapiV1CapitalConfigGetallGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1CapitalConfigGetallGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **recvWindow** | **int32** | The value cannot be greater than 60000 | 
 **timestamp** | **int32** | UTC timestamp | 
 **signature** | **string** | signature | 

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


## SapiV1CapitalDepositAddressGet

> InlineResponse20018 SapiV1CapitalDepositAddressGet(ctx).Coin(coin).Network(network).Timestamp(timestamp).RecvWindow(recvWindow).Signature(signature).Execute()

Withdraw History (supporting network) (USER_DATA)



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
    coin := "BNB" // string | coin name
    network := "ETH" // string |  (optional)
    timestamp := int32(56) // int32 | UTC timestamp (optional)
    recvWindow := int32(5000) // int32 | The value cannot be greater than 60000 (optional)
    signature := "signature_example" // string | signature (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WalletApi.SapiV1CapitalDepositAddressGet(context.Background()).Coin(coin).Network(network).Timestamp(timestamp).RecvWindow(recvWindow).Signature(signature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.SapiV1CapitalDepositAddressGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1CapitalDepositAddressGet`: InlineResponse20018
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.SapiV1CapitalDepositAddressGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1CapitalDepositAddressGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **coin** | **string** | coin name | 
 **network** | **string** |  | 
 **timestamp** | **int32** | UTC timestamp | 
 **recvWindow** | **int32** | The value cannot be greater than 60000 | 
 **signature** | **string** | signature | 

### Return type

[**InlineResponse20018**](InlineResponse20018.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1CapitalDepositHisrecGet

> []InlineResponse20016 SapiV1CapitalDepositHisrecGet(ctx).Coin(coin).Status(status).StartTime(startTime).EndTime(endTime).Offset(offset).Limit(limit).Timestamp(timestamp).RecvWindow(recvWindow).Signature(signature).Execute()

Deposit History（supporting network） (USER_DATA)



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
    coin := "BNB" // string |  (optional)
    status := int32(56) // int32 | 0 -> pending\\ 6 -> credited but cannot withdraw\\ 1 -> success (optional)
    startTime := int32(56) // int32 | Timestamp in ms (optional)
    endTime := int32(56) // int32 | Timestamp in ms (optional)
    offset := int32(56) // int32 |  (optional)
    limit := int32(500) // int32 | Default 500; max 1000. (optional)
    timestamp := int32(56) // int32 | UTC timestamp (optional)
    recvWindow := int32(5000) // int32 | The value cannot be greater than 60000 (optional)
    signature := "signature_example" // string | signature (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WalletApi.SapiV1CapitalDepositHisrecGet(context.Background()).Coin(coin).Status(status).StartTime(startTime).EndTime(endTime).Offset(offset).Limit(limit).Timestamp(timestamp).RecvWindow(recvWindow).Signature(signature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.SapiV1CapitalDepositHisrecGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1CapitalDepositHisrecGet`: []InlineResponse20016
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.SapiV1CapitalDepositHisrecGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1CapitalDepositHisrecGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **coin** | **string** |  | 
 **status** | **int32** | 0 -&gt; pending\\ 6 -&gt; credited but cannot withdraw\\ 1 -&gt; success | 
 **startTime** | **int32** | Timestamp in ms | 
 **endTime** | **int32** | Timestamp in ms | 
 **offset** | **int32** |  | 
 **limit** | **int32** | Default 500; max 1000. | 
 **timestamp** | **int32** | UTC timestamp | 
 **recvWindow** | **int32** | The value cannot be greater than 60000 | 
 **signature** | **string** | signature | 

### Return type

[**[]InlineResponse20016**](InlineResponse20016.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1CapitalWithdrawApplyPost

> InlineResponse20015 SapiV1CapitalWithdrawApplyPost(ctx).Coin(coin).Address(address).Amount(amount).WithdrawOrderId(withdrawOrderId).Network(network).AddressTag(addressTag).TransactionFeeFlag(transactionFeeFlag).Name(name).Timestamp(timestamp).RecvWindow(recvWindow).Signature(signature).Execute()

Withdraw



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
    coin := "BNB" // string | coin name
    address := "address_example" // string | 
    amount := float32(3.4) // float32 | 
    withdrawOrderId := "withdrawOrderId_example" // string | client id for withdraw (optional)
    network := "network_example" // string | getting value from `GET /sapi/v1/capital/config/getall` (optional)
    addressTag := "addressTag_example" // string | Secondary address identifier for coins like XRP,XMR etc. (optional)
    transactionFeeFlag := true // bool | When making internal transfer - `true` ->  returning the fee to the destination account;  - `false` -> returning the fee back to the departure account.  (optional) (default to false)
    name := "name_example" // string |  (optional)
    timestamp := int32(56) // int32 | UTC timestamp (optional)
    recvWindow := int32(5000) // int32 | The value cannot be greater than 60000 (optional)
    signature := "signature_example" // string | signature (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WalletApi.SapiV1CapitalWithdrawApplyPost(context.Background()).Coin(coin).Address(address).Amount(amount).WithdrawOrderId(withdrawOrderId).Network(network).AddressTag(addressTag).TransactionFeeFlag(transactionFeeFlag).Name(name).Timestamp(timestamp).RecvWindow(recvWindow).Signature(signature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.SapiV1CapitalWithdrawApplyPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1CapitalWithdrawApplyPost`: InlineResponse20015
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.SapiV1CapitalWithdrawApplyPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1CapitalWithdrawApplyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **coin** | **string** | coin name | 
 **address** | **string** |  | 
 **amount** | **float32** |  | 
 **withdrawOrderId** | **string** | client id for withdraw | 
 **network** | **string** | getting value from &#x60;GET /sapi/v1/capital/config/getall&#x60; | 
 **addressTag** | **string** | Secondary address identifier for coins like XRP,XMR etc. | 
 **transactionFeeFlag** | **bool** | When making internal transfer - &#x60;true&#x60; -&gt;  returning the fee to the destination account;  - &#x60;false&#x60; -&gt; returning the fee back to the departure account.  | [default to false]
 **name** | **string** |  | 
 **timestamp** | **int32** | UTC timestamp | 
 **recvWindow** | **int32** | The value cannot be greater than 60000 | 
 **signature** | **string** | signature | 

### Return type

[**InlineResponse20015**](InlineResponse20015.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1CapitalWithdrawHistoryGet

> []InlineResponse20017 SapiV1CapitalWithdrawHistoryGet(ctx).Coin(coin).Status(status).StartTime(startTime).EndTime(endTime).Offset(offset).Limit(limit).Timestamp(timestamp).RecvWindow(recvWindow).Signature(signature).Execute()

Withdraw History (supporting network) (USER_DATA)



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
    coin := "BNB" // string |  (optional)
    status := int32(56) // int32 | 0:Email Sent 1:Cancelled 2:Awaiting Approval  3:Rejected  4:Processing  5:Failure  6:Completed (optional)
    startTime := int32(56) // int32 | Timestamp in ms (optional)
    endTime := int32(56) // int32 | Timestamp in ms (optional)
    offset := int32(56) // int32 |  (optional)
    limit := int32(500) // int32 | Default 500; max 1000. (optional)
    timestamp := int32(56) // int32 | UTC timestamp (optional)
    recvWindow := int32(5000) // int32 | The value cannot be greater than 60000 (optional)
    signature := "signature_example" // string | signature (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WalletApi.SapiV1CapitalWithdrawHistoryGet(context.Background()).Coin(coin).Status(status).StartTime(startTime).EndTime(endTime).Offset(offset).Limit(limit).Timestamp(timestamp).RecvWindow(recvWindow).Signature(signature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.SapiV1CapitalWithdrawHistoryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1CapitalWithdrawHistoryGet`: []InlineResponse20017
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.SapiV1CapitalWithdrawHistoryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1CapitalWithdrawHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **coin** | **string** |  | 
 **status** | **int32** | 0:Email Sent 1:Cancelled 2:Awaiting Approval  3:Rejected  4:Processing  5:Failure  6:Completed | 
 **startTime** | **int32** | Timestamp in ms | 
 **endTime** | **int32** | Timestamp in ms | 
 **offset** | **int32** |  | 
 **limit** | **int32** | Default 500; max 1000. | 
 **timestamp** | **int32** | UTC timestamp | 
 **recvWindow** | **int32** | The value cannot be greater than 60000 | 
 **signature** | **string** | signature | 

### Return type

[**[]InlineResponse20017**](InlineResponse20017.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WapiV3AccountStatusHtmlGet

> InlineResponse20019 WapiV3AccountStatusHtmlGet(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Signature(signature).Execute()

Account Status (USER_DATA)



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
    timestamp := int32(56) // int32 | UTC timestamp (optional)
    recvWindow := int32(5000) // int32 | The value cannot be greater than 60000 (optional)
    signature := "signature_example" // string | signature (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WalletApi.WapiV3AccountStatusHtmlGet(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Signature(signature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.WapiV3AccountStatusHtmlGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WapiV3AccountStatusHtmlGet`: InlineResponse20019
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.WapiV3AccountStatusHtmlGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWapiV3AccountStatusHtmlGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int32** | UTC timestamp | 
 **recvWindow** | **int32** | The value cannot be greater than 60000 | 
 **signature** | **string** | signature | 

### Return type

[**InlineResponse20019**](InlineResponse20019.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WapiV3ApiTradingStatusHtmlGet

> InlineResponse20020 WapiV3ApiTradingStatusHtmlGet(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Signature(signature).Execute()

Account API Trading Status (USER_DATA)



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
    timestamp := int32(56) // int32 | UTC timestamp (optional)
    recvWindow := int32(5000) // int32 | The value cannot be greater than 60000 (optional)
    signature := "signature_example" // string | signature (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WalletApi.WapiV3ApiTradingStatusHtmlGet(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Signature(signature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.WapiV3ApiTradingStatusHtmlGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WapiV3ApiTradingStatusHtmlGet`: InlineResponse20020
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.WapiV3ApiTradingStatusHtmlGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWapiV3ApiTradingStatusHtmlGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int32** | UTC timestamp | 
 **recvWindow** | **int32** | The value cannot be greater than 60000 | 
 **signature** | **string** | signature | 

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


## WapiV3AssetDetailHtmlGet

> InlineResponse20024 WapiV3AssetDetailHtmlGet(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Signature(signature).Execute()

Asset Detail (USER_DATA)



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
    timestamp := int32(56) // int32 | UTC timestamp (optional)
    recvWindow := int32(5000) // int32 | The value cannot be greater than 60000 (optional)
    signature := "signature_example" // string | signature (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WalletApi.WapiV3AssetDetailHtmlGet(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Signature(signature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.WapiV3AssetDetailHtmlGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WapiV3AssetDetailHtmlGet`: InlineResponse20024
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.WapiV3AssetDetailHtmlGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWapiV3AssetDetailHtmlGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int32** | UTC timestamp | 
 **recvWindow** | **int32** | The value cannot be greater than 60000 | 
 **signature** | **string** | signature | 

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


## WapiV3SystemStatusHtmlGet

> map[string]interface{} WapiV3SystemStatusHtmlGet(ctx).Execute()

System Status (System)



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
    resp, r, err := api_client.WalletApi.WapiV3SystemStatusHtmlGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.WapiV3SystemStatusHtmlGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WapiV3SystemStatusHtmlGet`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.WapiV3SystemStatusHtmlGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWapiV3SystemStatusHtmlGetRequest struct via the builder pattern


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


## WapiV3TradeFeeHtmlGet

> InlineResponse20025 WapiV3TradeFeeHtmlGet(ctx).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Signature(signature).Execute()

Trade Fee (USER_DATA)



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
    symbol := "BNBUSDT" // string | trading symbol, e.g. BNBUSDT (optional)
    timestamp := int32(56) // int32 | UTC timestamp (optional)
    recvWindow := int32(5000) // int32 | The value cannot be greater than 60000 (optional)
    signature := "signature_example" // string | signature (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WalletApi.WapiV3TradeFeeHtmlGet(context.Background()).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Signature(signature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.WapiV3TradeFeeHtmlGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WapiV3TradeFeeHtmlGet`: InlineResponse20025
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.WapiV3TradeFeeHtmlGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWapiV3TradeFeeHtmlGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | trading symbol, e.g. BNBUSDT | 
 **timestamp** | **int32** | UTC timestamp | 
 **recvWindow** | **int32** | The value cannot be greater than 60000 | 
 **signature** | **string** | signature | 

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


## WapiV3UserAssetDribbletLogHtmlGet

> InlineResponse20021 WapiV3UserAssetDribbletLogHtmlGet(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Signature(signature).Execute()

DustLog (USER_DATA)



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
    timestamp := int32(56) // int32 | UTC timestamp (optional)
    recvWindow := int32(5000) // int32 | The value cannot be greater than 60000 (optional)
    signature := "signature_example" // string | signature (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WalletApi.WapiV3UserAssetDribbletLogHtmlGet(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Signature(signature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.WapiV3UserAssetDribbletLogHtmlGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WapiV3UserAssetDribbletLogHtmlGet`: InlineResponse20021
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.WapiV3UserAssetDribbletLogHtmlGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWapiV3UserAssetDribbletLogHtmlGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int32** | UTC timestamp | 
 **recvWindow** | **int32** | The value cannot be greater than 60000 | 
 **signature** | **string** | signature | 

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

