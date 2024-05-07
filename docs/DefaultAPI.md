# \DefaultAPI

All URIs are relative to *https://rpc-testnet.qanplatform.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QanAccounts**](DefaultAPI.md#QanAccounts) | **Get** /accounts/ | 
[**QanBlobBaseFee**](DefaultAPI.md#QanBlobBaseFee) | **Get** /blobBaseFee/ | 
[**QanBlockNumber**](DefaultAPI.md#QanBlockNumber) | **Get** /blockNumber/ | 
[**QanCall**](DefaultAPI.md#QanCall) | **Post** /call/ | 
[**QanChainId**](DefaultAPI.md#QanChainId) | **Get** /chainId/ | 
[**QanEstimateGas**](DefaultAPI.md#QanEstimateGas) | **Post** /estimateGas/ | 
[**QanFeeHistory**](DefaultAPI.md#QanFeeHistory) | **Post** /feeHistory/ | 
[**QanGasPrice**](DefaultAPI.md#QanGasPrice) | **Get** /gasPrice/ | 
[**QanGetAccount**](DefaultAPI.md#QanGetAccount) | **Get** /getAccount/{Address}/{BlockReference}/ | 
[**QanGetBalance**](DefaultAPI.md#QanGetBalance) | **Get** /getBalance/{Address}/ | 
[**QanGetBlockByHash**](DefaultAPI.md#QanGetBlockByHash) | **Get** /getBlockByHash/{Hash}/{TransactionDetailFlag}/ | 
[**QanGetBlockByNumber**](DefaultAPI.md#QanGetBlockByNumber) | **Get** /getBlockByNumber/{BlockNumber}/{TransactionDetailFlag}/ | 
[**QanGetBlockReceipts**](DefaultAPI.md#QanGetBlockReceipts) | **Get** /getBlockReceipts/{BlockNumber}/ | 
[**QanGetBlockTransactionCountByHash**](DefaultAPI.md#QanGetBlockTransactionCountByHash) | **Get** /getBlockTransactionCountByHash/{Hash}/ | 
[**QanGetBlockTransactionCountByNumber**](DefaultAPI.md#QanGetBlockTransactionCountByNumber) | **Get** /getBlockTransactionCountByNumber/{BlockNumber}/ | 
[**QanGetCode**](DefaultAPI.md#QanGetCode) | **Get** /getCode/{Address}/ | 
[**QanGetFilterChanges**](DefaultAPI.md#QanGetFilterChanges) | **Get** /getFilterChanges/{FilterId}/ | 
[**QanGetFilterLogs**](DefaultAPI.md#QanGetFilterLogs) | **Get** /getFilterLogs/{Id}/ | 
[**QanGetLogs**](DefaultAPI.md#QanGetLogs) | **Post** /getLogs/ | 
[**QanGetProof**](DefaultAPI.md#QanGetProof) | **Post** /getProof/ | 
[**QanGetStorageAt**](DefaultAPI.md#QanGetStorageAt) | **Post** /getStorageAt/ | 
[**QanGetTransactionByBlockHashAndIndex**](DefaultAPI.md#QanGetTransactionByBlockHashAndIndex) | **Get** /getTransactionByBlockHashAndIndex/{blockHash}/{index}/ | 
[**QanGetTransactionByBlockNumberAndIndex**](DefaultAPI.md#QanGetTransactionByBlockNumberAndIndex) | **Get** /getTransactionByBlockNumberAndIndex/{blockNumber}/{index}/ | 
[**QanGetTransactionByHash**](DefaultAPI.md#QanGetTransactionByHash) | **Get** /getTransactionByHash/{hash}/ | 
[**QanGetTransactionCount**](DefaultAPI.md#QanGetTransactionCount) | **Get** /getTransactionCount/{Address}/{BlockNumber}/ | 
[**QanGetTransactionReceipt**](DefaultAPI.md#QanGetTransactionReceipt) | **Get** /getTransactionReceipt/{Hash}/ | 
[**QanGetUncleCountByBlockHash**](DefaultAPI.md#QanGetUncleCountByBlockHash) | **Get** /getUncleCountByBlockHash/{Hash}/ | 
[**QanGetUncleCountByBlockNumber**](DefaultAPI.md#QanGetUncleCountByBlockNumber) | **Get** /getUncleCountByBlockNumber/{BlockNumber}/ | 
[**QanMaxPriorityFeePerGas**](DefaultAPI.md#QanMaxPriorityFeePerGas) | **Get** /maxPriorityFeePerGas/ | 
[**QanNewBlockFilter**](DefaultAPI.md#QanNewBlockFilter) | **Get** /newBlockFilter/ | 
[**QanNewFilter**](DefaultAPI.md#QanNewFilter) | **Post** /newFilter/ | 
[**QanNewPendingTransactionFilter**](DefaultAPI.md#QanNewPendingTransactionFilter) | **Get** /newPendingTransactionFilter/ | 
[**QanSendRawTransaction**](DefaultAPI.md#QanSendRawTransaction) | **Post** /sendRawTransaction/ | 
[**QanSubscribe**](DefaultAPI.md#QanSubscribe) | **Post** /subscribe/ | 
[**QanSyncing**](DefaultAPI.md#QanSyncing) | **Get** /syncing/ | 
[**QanUninstallFilter**](DefaultAPI.md#QanUninstallFilter) | **Get** /uninstallFilter/{FilterId}/ | 
[**QanUnsubscribe**](DefaultAPI.md#QanUnsubscribe) | **Get** /unsubscribe/{SubscriptionId}/ | 
[**QanXlinkValid**](DefaultAPI.md#QanXlinkValid) | **Get** /xlinkValid/ | 



## QanAccounts

> OutputAccounts QanAccounts(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qanplatform/sdk-golang"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.QanAccounts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.QanAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanAccounts`: OutputAccounts
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.QanAccounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiQanAccountsRequest struct via the builder pattern


### Return type

[**OutputAccounts**](OutputAccounts.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanBlobBaseFee

> OutputBlobBaseFee QanBlobBaseFee(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qanplatform/sdk-golang"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.QanBlobBaseFee(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.QanBlobBaseFee``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanBlobBaseFee`: OutputBlobBaseFee
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.QanBlobBaseFee`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiQanBlobBaseFeeRequest struct via the builder pattern


### Return type

[**OutputBlobBaseFee**](OutputBlobBaseFee.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanBlockNumber

> OutputBlockNumber QanBlockNumber(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qanplatform/sdk-golang"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.QanBlockNumber(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.QanBlockNumber``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanBlockNumber`: OutputBlockNumber
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.QanBlockNumber`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiQanBlockNumberRequest struct via the builder pattern


### Return type

[**OutputBlockNumber**](OutputBlockNumber.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanCall

> OutputCall QanCall(ctx).InputCall(inputCall).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qanplatform/sdk-golang"
)

func main() {
	inputCall := *openapiclient.NewInputCall("BlockNumber_example", *openapiclient.NewParamsTransaction("To_example")) // InputCall | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.QanCall(context.Background()).InputCall(inputCall).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.QanCall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanCall`: OutputCall
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.QanCall`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQanCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputCall** | [**InputCall**](InputCall.md) |  | 

### Return type

[**OutputCall**](OutputCall.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanChainId

> OutputChainId QanChainId(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qanplatform/sdk-golang"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.QanChainId(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.QanChainId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanChainId`: OutputChainId
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.QanChainId`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiQanChainIdRequest struct via the builder pattern


### Return type

[**OutputChainId**](OutputChainId.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanEstimateGas

> OutputEstimateGas QanEstimateGas(ctx).InputEstimateGas(inputEstimateGas).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qanplatform/sdk-golang"
)

func main() {
	inputEstimateGas := *openapiclient.NewInputEstimateGas(*openapiclient.NewParamsTransaction("To_example")) // InputEstimateGas | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.QanEstimateGas(context.Background()).InputEstimateGas(inputEstimateGas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.QanEstimateGas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanEstimateGas`: OutputEstimateGas
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.QanEstimateGas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQanEstimateGasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputEstimateGas** | [**InputEstimateGas**](InputEstimateGas.md) |  | 

### Return type

[**OutputEstimateGas**](OutputEstimateGas.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanFeeHistory

> OutputFeeHistory QanFeeHistory(ctx).InputFeeHistory(inputFeeHistory).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qanplatform/sdk-golang"
)

func main() {
	inputFeeHistory := *openapiclient.NewInputFeeHistory(int32(123), "NewestBlock_example", []int32{int32(123)}) // InputFeeHistory | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.QanFeeHistory(context.Background()).InputFeeHistory(inputFeeHistory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.QanFeeHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanFeeHistory`: OutputFeeHistory
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.QanFeeHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQanFeeHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputFeeHistory** | [**InputFeeHistory**](InputFeeHistory.md) |  | 

### Return type

[**OutputFeeHistory**](OutputFeeHistory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanGasPrice

> OutputGasPrice QanGasPrice(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qanplatform/sdk-golang"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.QanGasPrice(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.QanGasPrice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanGasPrice`: OutputGasPrice
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.QanGasPrice`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiQanGasPriceRequest struct via the builder pattern


### Return type

[**OutputGasPrice**](OutputGasPrice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanGetAccount

> OutputGetAccount QanGetAccount(ctx, address, blockReference).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qanplatform/sdk-golang"
)

func main() {
	address := "0xa1e4380a3b1f749673e270229993ee55f35663b4" // string | The account address for which the information is to be retrieved
	blockReference := "latest" // string | The block number in hexadecimal or decimal format or the string latest, earliest, pending, see the default block parameter description in the official Ethereum documentation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.QanGetAccount(context.Background(), address, blockReference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.QanGetAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanGetAccount`: OutputGetAccount
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.QanGetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | The account address for which the information is to be retrieved | 
**blockReference** | **string** | The block number in hexadecimal or decimal format or the string latest, earliest, pending, see the default block parameter description in the official Ethereum documentation | 

### Other Parameters

Other parameters are passed through a pointer to a apiQanGetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OutputGetAccount**](OutputGetAccount.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanGetBalance

> OutputGetBalance QanGetBalance(ctx, address).BlockNumber(blockNumber).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qanplatform/sdk-golang"
)

func main() {
	address := "0xa1e4380a3b1f749673e270229993ee55f35663b4" // string | A 20 bytes long hexadecimal value representing an Ethereum address
	blockNumber := "blockNumber_example" // string | The block number in hexadecimal or decimal format or the string latest, earliest, pending, see the default block parameter description in the official Ethereum documentation (optional) (default to "latest")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.QanGetBalance(context.Background(), address).BlockNumber(blockNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.QanGetBalance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanGetBalance`: OutputGetBalance
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.QanGetBalance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | A 20 bytes long hexadecimal value representing an Ethereum address | 

### Other Parameters

Other parameters are passed through a pointer to a apiQanGetBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blockNumber** | **string** | The block number in hexadecimal or decimal format or the string latest, earliest, pending, see the default block parameter description in the official Ethereum documentation | [default to &quot;latest&quot;]

### Return type

[**OutputGetBalance**](OutputGetBalance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanGetBlockByHash

> OutputGetBlockByHash QanGetBlockByHash(ctx, hash, transactionDetailFlag).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qanplatform/sdk-golang"
)

func main() {
	hash := "0x4e3a3754410177e6937ef1f84bba68ea139e8d1a2258c5f85db9f1cd715a1bdd" // string | The hash (32 bytes) of the block
	transactionDetailFlag := true // bool | The method returns the full transaction objects when this value is true otherwise, it returns only the hashes of the transactions (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.QanGetBlockByHash(context.Background(), hash, transactionDetailFlag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.QanGetBlockByHash``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanGetBlockByHash`: OutputGetBlockByHash
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.QanGetBlockByHash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string** | The hash (32 bytes) of the block | 
**transactionDetailFlag** | **bool** | The method returns the full transaction objects when this value is true otherwise, it returns only the hashes of the transactions | [default to false]

### Other Parameters

Other parameters are passed through a pointer to a apiQanGetBlockByHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OutputGetBlockByHash**](OutputGetBlockByHash.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanGetBlockByNumber

> OutputGetBlockByNumber QanGetBlockByNumber(ctx, blockNumber, transactionDetailFlag).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qanplatform/sdk-golang"
)

func main() {
	blockNumber := "blockNumber_example" // string | The block number in hexadecimal or decimal format or the string latest, earliest, pending, see the default block parameter description in the official Ethereum documentation (default to "latest")
	transactionDetailFlag := true // bool | The method returns the full transaction objects when this value is true otherwise, it returns only the hashes of the transactions (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.QanGetBlockByNumber(context.Background(), blockNumber, transactionDetailFlag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.QanGetBlockByNumber``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanGetBlockByNumber`: OutputGetBlockByNumber
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.QanGetBlockByNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockNumber** | **string** | The block number in hexadecimal or decimal format or the string latest, earliest, pending, see the default block parameter description in the official Ethereum documentation | [default to &quot;latest&quot;]
**transactionDetailFlag** | **bool** | The method returns the full transaction objects when this value is true otherwise, it returns only the hashes of the transactions | [default to false]

### Other Parameters

Other parameters are passed through a pointer to a apiQanGetBlockByNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OutputGetBlockByNumber**](OutputGetBlockByNumber.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanGetBlockReceipts

> OutputGetBlockReceipts QanGetBlockReceipts(ctx, blockNumber).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qanplatform/sdk-golang"
)

func main() {
	blockNumber := "blockNumber_example" // string | The block number in hexadecimal or decimal format or the string latest, earliest, pending, see the default block parameter description in the official Ethereum documentation (default to "latest")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.QanGetBlockReceipts(context.Background(), blockNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.QanGetBlockReceipts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanGetBlockReceipts`: OutputGetBlockReceipts
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.QanGetBlockReceipts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockNumber** | **string** | The block number in hexadecimal or decimal format or the string latest, earliest, pending, see the default block parameter description in the official Ethereum documentation | [default to &quot;latest&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiQanGetBlockReceiptsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OutputGetBlockReceipts**](OutputGetBlockReceipts.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanGetBlockTransactionCountByHash

> OutputGetBlockTransactionCountByHash QanGetBlockTransactionCountByHash(ctx, hash).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qanplatform/sdk-golang"
)

func main() {
	hash := "0x4e3a3754410177e6937ef1f84bba68ea139e8d1a2258c5f85db9f1cd715a1bdd" // string | The hash of the block

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.QanGetBlockTransactionCountByHash(context.Background(), hash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.QanGetBlockTransactionCountByHash``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanGetBlockTransactionCountByHash`: OutputGetBlockTransactionCountByHash
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.QanGetBlockTransactionCountByHash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string** | The hash of the block | 

### Other Parameters

Other parameters are passed through a pointer to a apiQanGetBlockTransactionCountByHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OutputGetBlockTransactionCountByHash**](OutputGetBlockTransactionCountByHash.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanGetBlockTransactionCountByNumber

> OutputGetBlockTransactionCountByNumber QanGetBlockTransactionCountByNumber(ctx, blockNumber).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qanplatform/sdk-golang"
)

func main() {
	blockNumber := "latest" // string | The block number in hexadecimal or decimal format or the string latest, earliest, pending, see the default block parameter description in the official Ethereum documentation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.QanGetBlockTransactionCountByNumber(context.Background(), blockNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.QanGetBlockTransactionCountByNumber``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanGetBlockTransactionCountByNumber`: OutputGetBlockTransactionCountByNumber
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.QanGetBlockTransactionCountByNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockNumber** | **string** | The block number in hexadecimal or decimal format or the string latest, earliest, pending, see the default block parameter description in the official Ethereum documentation | 

### Other Parameters

Other parameters are passed through a pointer to a apiQanGetBlockTransactionCountByNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OutputGetBlockTransactionCountByNumber**](OutputGetBlockTransactionCountByNumber.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanGetCode

> OutputGetCode QanGetCode(ctx, address).BlockNumber(blockNumber).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qanplatform/sdk-golang"
)

func main() {
	address := "0xa1e4380a3b1f749673e270229993ee55f35663b4" // string | The address of the smart contract from which the bytecode will be obtained
	blockNumber := "latest" // string | The block number in hexadecimal or decimal format or the string latest, earliest, pending, see the default block parameter description in the official Ethereum documentation (optional) (default to "latest")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.QanGetCode(context.Background(), address).BlockNumber(blockNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.QanGetCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanGetCode`: OutputGetCode
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.QanGetCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | The address of the smart contract from which the bytecode will be obtained | 

### Other Parameters

Other parameters are passed through a pointer to a apiQanGetCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blockNumber** | **string** | The block number in hexadecimal or decimal format or the string latest, earliest, pending, see the default block parameter description in the official Ethereum documentation | [default to &quot;latest&quot;]

### Return type

[**OutputGetCode**](OutputGetCode.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanGetFilterChanges

> OutputGetFilterChanges QanGetFilterChanges(ctx, filterId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qanplatform/sdk-golang"
)

func main() {
	filterId := "filterId_example" // string | The filter id that is returned from eth_newFilter, eth_newBlockFilter or eth_newPendingTransactionFilter

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.QanGetFilterChanges(context.Background(), filterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.QanGetFilterChanges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanGetFilterChanges`: OutputGetFilterChanges
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.QanGetFilterChanges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filterId** | **string** | The filter id that is returned from eth_newFilter, eth_newBlockFilter or eth_newPendingTransactionFilter | 

### Other Parameters

Other parameters are passed through a pointer to a apiQanGetFilterChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OutputGetFilterChanges**](OutputGetFilterChanges.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanGetFilterLogs

> OutputGetFilterLogs QanGetFilterLogs(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qanplatform/sdk-golang"
)

func main() {
	id := "id_example" // string | The filter ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.QanGetFilterLogs(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.QanGetFilterLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanGetFilterLogs`: OutputGetFilterLogs
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.QanGetFilterLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The filter ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiQanGetFilterLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OutputGetFilterLogs**](OutputGetFilterLogs.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanGetLogs

> OutputGetLogs QanGetLogs(ctx).InputGetLogs(inputGetLogs).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qanplatform/sdk-golang"
)

func main() {
	inputGetLogs := *openapiclient.NewInputGetLogs() // InputGetLogs | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.QanGetLogs(context.Background()).InputGetLogs(inputGetLogs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.QanGetLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanGetLogs`: OutputGetLogs
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.QanGetLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQanGetLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputGetLogs** | [**InputGetLogs**](InputGetLogs.md) |  | 

### Return type

[**OutputGetLogs**](OutputGetLogs.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanGetProof

> OutputGetProof QanGetProof(ctx).InputGetProof(inputGetProof).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qanplatform/sdk-golang"
)

func main() {
	inputGetProof := *openapiclient.NewInputGetProof("Address_example", []string{"StorageKeys_example"}) // InputGetProof | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.QanGetProof(context.Background()).InputGetProof(inputGetProof).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.QanGetProof``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanGetProof`: OutputGetProof
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.QanGetProof`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQanGetProofRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputGetProof** | [**InputGetProof**](InputGetProof.md) |  | 

### Return type

[**OutputGetProof**](OutputGetProof.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanGetStorageAt

> OutputGetStorageAt QanGetStorageAt(ctx).InputGetStorageAt(inputGetStorageAt).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qanplatform/sdk-golang"
)

func main() {
	inputGetStorageAt := *openapiclient.NewInputGetStorageAt("Address_example", "BlockNumber_example", "Position_example") // InputGetStorageAt | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.QanGetStorageAt(context.Background()).InputGetStorageAt(inputGetStorageAt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.QanGetStorageAt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanGetStorageAt`: OutputGetStorageAt
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.QanGetStorageAt`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQanGetStorageAtRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputGetStorageAt** | [**InputGetStorageAt**](InputGetStorageAt.md) |  | 

### Return type

[**OutputGetStorageAt**](OutputGetStorageAt.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanGetTransactionByBlockHashAndIndex

> OutputGetTransactionByBlockHashAndIndex QanGetTransactionByBlockHashAndIndex(ctx, blockHash, index).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qanplatform/sdk-golang"
)

func main() {
	blockHash := "0x4e3a3754410177e6937ef1f84bba68ea139e8d1a2258c5f85db9f1cd715a1bdd" // string | 
	index := "0" // string | An integer of the transaction index position

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.QanGetTransactionByBlockHashAndIndex(context.Background(), blockHash, index).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.QanGetTransactionByBlockHashAndIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanGetTransactionByBlockHashAndIndex`: OutputGetTransactionByBlockHashAndIndex
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.QanGetTransactionByBlockHashAndIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockHash** | **string** |  | 
**index** | **string** | An integer of the transaction index position | 

### Other Parameters

Other parameters are passed through a pointer to a apiQanGetTransactionByBlockHashAndIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OutputGetTransactionByBlockHashAndIndex**](OutputGetTransactionByBlockHashAndIndex.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanGetTransactionByBlockNumberAndIndex

> OutputGetTransactionByBlockNumberAndIndex QanGetTransactionByBlockNumberAndIndex(ctx, blockNumber, index).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qanplatform/sdk-golang"
)

func main() {
	blockNumber := "latest" // string | The block number in hexadecimal or decimal format or the string latest, earliest, pending, see the default block parameter description in the official Ethereum documentation
	index := "0" // string | An integer of the transaction index position

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.QanGetTransactionByBlockNumberAndIndex(context.Background(), blockNumber, index).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.QanGetTransactionByBlockNumberAndIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanGetTransactionByBlockNumberAndIndex`: OutputGetTransactionByBlockNumberAndIndex
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.QanGetTransactionByBlockNumberAndIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockNumber** | **string** | The block number in hexadecimal or decimal format or the string latest, earliest, pending, see the default block parameter description in the official Ethereum documentation | 
**index** | **string** | An integer of the transaction index position | 

### Other Parameters

Other parameters are passed through a pointer to a apiQanGetTransactionByBlockNumberAndIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OutputGetTransactionByBlockNumberAndIndex**](OutputGetTransactionByBlockNumberAndIndex.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanGetTransactionByHash

> OutputGetTransactionByHash QanGetTransactionByHash(ctx, hash).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qanplatform/sdk-golang"
)

func main() {
	hash := "0x5c504ed432cb51138bcf09aa5e8a410dd4a1e204ef84bfed1be16dfba1b22060" // string | The hash of a transaction

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.QanGetTransactionByHash(context.Background(), hash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.QanGetTransactionByHash``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanGetTransactionByHash`: OutputGetTransactionByHash
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.QanGetTransactionByHash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string** | The hash of a transaction | 

### Other Parameters

Other parameters are passed through a pointer to a apiQanGetTransactionByHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OutputGetTransactionByHash**](OutputGetTransactionByHash.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanGetTransactionCount

> OutputGetTransactionCount QanGetTransactionCount(ctx, address, blockNumber).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qanplatform/sdk-golang"
)

func main() {
	address := "0xa1e4380a3b1f749673e270229993ee55f35663b4" // string | The address from which the transaction count to be checked
	blockNumber := "latest" // string | The block number in hexadecimal or decimal format or the string latest, earliest, pending, see the default block parameter description in the official Ethereum documentation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.QanGetTransactionCount(context.Background(), address, blockNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.QanGetTransactionCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanGetTransactionCount`: OutputGetTransactionCount
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.QanGetTransactionCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | The address from which the transaction count to be checked | 
**blockNumber** | **string** | The block number in hexadecimal or decimal format or the string latest, earliest, pending, see the default block parameter description in the official Ethereum documentation | 

### Other Parameters

Other parameters are passed through a pointer to a apiQanGetTransactionCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OutputGetTransactionCount**](OutputGetTransactionCount.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanGetTransactionReceipt

> OutputGetTransactionReceipt QanGetTransactionReceipt(ctx, hash).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qanplatform/sdk-golang"
)

func main() {
	hash := "0x4e3a3754410177e6937ef1f84bba68ea139e8d1a2258c5f85db9f1cd715a1bdd" // string | The hash of a transaction

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.QanGetTransactionReceipt(context.Background(), hash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.QanGetTransactionReceipt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanGetTransactionReceipt`: OutputGetTransactionReceipt
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.QanGetTransactionReceipt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string** | The hash of a transaction | 

### Other Parameters

Other parameters are passed through a pointer to a apiQanGetTransactionReceiptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OutputGetTransactionReceipt**](OutputGetTransactionReceipt.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanGetUncleCountByBlockHash

> OutputGetUncleCountByBlockHash QanGetUncleCountByBlockHash(ctx, hash).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qanplatform/sdk-golang"
)

func main() {
	hash := "0x4e216c95f527e9ba0f1161a1c4609b893302c704f05a520da8141ca91878f63e" // string | The hash of the block to get uncles for

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.QanGetUncleCountByBlockHash(context.Background(), hash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.QanGetUncleCountByBlockHash``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanGetUncleCountByBlockHash`: OutputGetUncleCountByBlockHash
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.QanGetUncleCountByBlockHash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string** | The hash of the block to get uncles for | 

### Other Parameters

Other parameters are passed through a pointer to a apiQanGetUncleCountByBlockHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OutputGetUncleCountByBlockHash**](OutputGetUncleCountByBlockHash.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanGetUncleCountByBlockNumber

> OutputGetUncleCountByBlockNumber QanGetUncleCountByBlockNumber(ctx, blockNumber).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qanplatform/sdk-golang"
)

func main() {
	blockNumber := "15537381" // string | The block number in hexadecimal or decimal format or the string latest, earliest, pending, see the default block parameter description in the official Ethereum documentation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.QanGetUncleCountByBlockNumber(context.Background(), blockNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.QanGetUncleCountByBlockNumber``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanGetUncleCountByBlockNumber`: OutputGetUncleCountByBlockNumber
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.QanGetUncleCountByBlockNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockNumber** | **string** | The block number in hexadecimal or decimal format or the string latest, earliest, pending, see the default block parameter description in the official Ethereum documentation | 

### Other Parameters

Other parameters are passed through a pointer to a apiQanGetUncleCountByBlockNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OutputGetUncleCountByBlockNumber**](OutputGetUncleCountByBlockNumber.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanMaxPriorityFeePerGas

> OutputMaxPriorityFeePerGas QanMaxPriorityFeePerGas(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qanplatform/sdk-golang"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.QanMaxPriorityFeePerGas(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.QanMaxPriorityFeePerGas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanMaxPriorityFeePerGas`: OutputMaxPriorityFeePerGas
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.QanMaxPriorityFeePerGas`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiQanMaxPriorityFeePerGasRequest struct via the builder pattern


### Return type

[**OutputMaxPriorityFeePerGas**](OutputMaxPriorityFeePerGas.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanNewBlockFilter

> OutputNewBlockFilter QanNewBlockFilter(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qanplatform/sdk-golang"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.QanNewBlockFilter(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.QanNewBlockFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanNewBlockFilter`: OutputNewBlockFilter
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.QanNewBlockFilter`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiQanNewBlockFilterRequest struct via the builder pattern


### Return type

[**OutputNewBlockFilter**](OutputNewBlockFilter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanNewFilter

> OutputNewFilter QanNewFilter(ctx).InputNewFilter(inputNewFilter).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qanplatform/sdk-golang"
)

func main() {
	inputNewFilter := *openapiclient.NewInputNewFilter(*openapiclient.NewFilterObject("Address_example", "FromBlock_example", "ToBlock_example", []string{"Topics_example"})) // InputNewFilter | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.QanNewFilter(context.Background()).InputNewFilter(inputNewFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.QanNewFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanNewFilter`: OutputNewFilter
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.QanNewFilter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQanNewFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputNewFilter** | [**InputNewFilter**](InputNewFilter.md) |  | 

### Return type

[**OutputNewFilter**](OutputNewFilter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanNewPendingTransactionFilter

> OutputNewPendingTransactionFilter QanNewPendingTransactionFilter(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qanplatform/sdk-golang"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.QanNewPendingTransactionFilter(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.QanNewPendingTransactionFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanNewPendingTransactionFilter`: OutputNewPendingTransactionFilter
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.QanNewPendingTransactionFilter`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiQanNewPendingTransactionFilterRequest struct via the builder pattern


### Return type

[**OutputNewPendingTransactionFilter**](OutputNewPendingTransactionFilter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanSendRawTransaction

> OutputSendRawTransaction QanSendRawTransaction(ctx).InputSendRawTransaction(inputSendRawTransaction).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qanplatform/sdk-golang"
)

func main() {
	inputSendRawTransaction := *openapiclient.NewInputSendRawTransaction("Data_example") // InputSendRawTransaction | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.QanSendRawTransaction(context.Background()).InputSendRawTransaction(inputSendRawTransaction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.QanSendRawTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanSendRawTransaction`: OutputSendRawTransaction
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.QanSendRawTransaction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQanSendRawTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputSendRawTransaction** | [**InputSendRawTransaction**](InputSendRawTransaction.md) |  | 

### Return type

[**OutputSendRawTransaction**](OutputSendRawTransaction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanSubscribe

> OutputSubscribe QanSubscribe(ctx).InputSubscribe(inputSubscribe).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qanplatform/sdk-golang"
)

func main() {
	inputSubscribe := *openapiclient.NewInputSubscribe(interface{}(123), false, "SubscriptionName_example") // InputSubscribe | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.QanSubscribe(context.Background()).InputSubscribe(inputSubscribe).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.QanSubscribe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanSubscribe`: OutputSubscribe
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.QanSubscribe`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQanSubscribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputSubscribe** | [**InputSubscribe**](InputSubscribe.md) |  | 

### Return type

[**OutputSubscribe**](OutputSubscribe.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanSyncing

> OutputSyncing QanSyncing(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qanplatform/sdk-golang"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.QanSyncing(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.QanSyncing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanSyncing`: OutputSyncing
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.QanSyncing`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiQanSyncingRequest struct via the builder pattern


### Return type

[**OutputSyncing**](OutputSyncing.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanUninstallFilter

> OutputUninstallFilter QanUninstallFilter(ctx, filterId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qanplatform/sdk-golang"
)

func main() {
	filterId := "filterId_example" // string | The filter ID that needs to be uninstalled. It should always be called when watch is no longer needed. Additionally, Filters timeout when they aren't requested with eth_getFilterChanges for a period of time

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.QanUninstallFilter(context.Background(), filterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.QanUninstallFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanUninstallFilter`: OutputUninstallFilter
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.QanUninstallFilter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filterId** | **string** | The filter ID that needs to be uninstalled. It should always be called when watch is no longer needed. Additionally, Filters timeout when they aren&#39;t requested with eth_getFilterChanges for a period of time | 

### Other Parameters

Other parameters are passed through a pointer to a apiQanUninstallFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OutputUninstallFilter**](OutputUninstallFilter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanUnsubscribe

> OutputUnsubscribe QanUnsubscribe(ctx, subscriptionId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qanplatform/sdk-golang"
)

func main() {
	subscriptionId := "subscriptionId_example" // string | A subscription ID that was previously generated in a eth_subscribe RPC request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.QanUnsubscribe(context.Background(), subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.QanUnsubscribe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanUnsubscribe`: OutputUnsubscribe
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.QanUnsubscribe`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | A subscription ID that was previously generated in a eth_subscribe RPC request | 

### Other Parameters

Other parameters are passed through a pointer to a apiQanUnsubscribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OutputUnsubscribe**](OutputUnsubscribe.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanXlinkValid

> QanXlinkValidResponse QanXlinkValid(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qanplatform/sdk-golang"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.QanXlinkValid(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.QanXlinkValid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanXlinkValid`: QanXlinkValidResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.QanXlinkValid`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiQanXlinkValidRequest struct via the builder pattern


### Return type

[**QanXlinkValidResponse**](QanXlinkValidResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

