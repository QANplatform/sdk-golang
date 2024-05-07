# QAN Golang SDK

This repository is guaranteed up-to-date with the upstream QAN API definitions, and leverages OpenAPI technology to stay consistent.

Versioning is based on SEMVER, meaning:

- Stable releases guarantee backwards compatibility for the same major versions.
- Minor releases will not contain breaking changes.
- Patch releases only focus on fixing issues.

## Documentation for API Endpoints

All URIs are relative to *https://rpc-testnet.qanplatform.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultAPI* | [**QanAccounts**](docs/DefaultAPI.md#qanaccounts) | **Get** /accounts/ | 
*DefaultAPI* | [**QanBlobBaseFee**](docs/DefaultAPI.md#qanblobbasefee) | **Get** /blobBaseFee/ | 
*DefaultAPI* | [**QanBlockNumber**](docs/DefaultAPI.md#qanblocknumber) | **Get** /blockNumber/ | 
*DefaultAPI* | [**QanCall**](docs/DefaultAPI.md#qancall) | **Post** /call/ | 
*DefaultAPI* | [**QanChainId**](docs/DefaultAPI.md#qanchainid) | **Get** /chainId/ | 
*DefaultAPI* | [**QanEstimateGas**](docs/DefaultAPI.md#qanestimategas) | **Post** /estimateGas/ | 
*DefaultAPI* | [**QanFeeHistory**](docs/DefaultAPI.md#qanfeehistory) | **Post** /feeHistory/ | 
*DefaultAPI* | [**QanGasPrice**](docs/DefaultAPI.md#qangasprice) | **Get** /gasPrice/ | 
*DefaultAPI* | [**QanGetAccount**](docs/DefaultAPI.md#qangetaccount) | **Get** /getAccount/{Address}/{BlockReference}/ | 
*DefaultAPI* | [**QanGetBalance**](docs/DefaultAPI.md#qangetbalance) | **Get** /getBalance/{Address}/ | 
*DefaultAPI* | [**QanGetBlockByHash**](docs/DefaultAPI.md#qangetblockbyhash) | **Get** /getBlockByHash/{Hash}/{TransactionDetailFlag}/ | 
*DefaultAPI* | [**QanGetBlockByNumber**](docs/DefaultAPI.md#qangetblockbynumber) | **Get** /getBlockByNumber/{BlockNumber}/{TransactionDetailFlag}/ | 
*DefaultAPI* | [**QanGetBlockReceipts**](docs/DefaultAPI.md#qangetblockreceipts) | **Get** /getBlockReceipts/{BlockNumber}/ | 
*DefaultAPI* | [**QanGetBlockTransactionCountByHash**](docs/DefaultAPI.md#qangetblocktransactioncountbyhash) | **Get** /getBlockTransactionCountByHash/{Hash}/ | 
*DefaultAPI* | [**QanGetBlockTransactionCountByNumber**](docs/DefaultAPI.md#qangetblocktransactioncountbynumber) | **Get** /getBlockTransactionCountByNumber/{BlockNumber}/ | 
*DefaultAPI* | [**QanGetCode**](docs/DefaultAPI.md#qangetcode) | **Get** /getCode/{Address}/ | 
*DefaultAPI* | [**QanGetFilterChanges**](docs/DefaultAPI.md#qangetfilterchanges) | **Get** /getFilterChanges/{FilterId}/ | 
*DefaultAPI* | [**QanGetFilterLogs**](docs/DefaultAPI.md#qangetfilterlogs) | **Get** /getFilterLogs/{Id}/ | 
*DefaultAPI* | [**QanGetLogs**](docs/DefaultAPI.md#qangetlogs) | **Post** /getLogs/ | 
*DefaultAPI* | [**QanGetProof**](docs/DefaultAPI.md#qangetproof) | **Post** /getProof/ | 
*DefaultAPI* | [**QanGetStorageAt**](docs/DefaultAPI.md#qangetstorageat) | **Post** /getStorageAt/ | 
*DefaultAPI* | [**QanGetTransactionByBlockHashAndIndex**](docs/DefaultAPI.md#qangettransactionbyblockhashandindex) | **Get** /getTransactionByBlockHashAndIndex/{blockHash}/{index}/ | 
*DefaultAPI* | [**QanGetTransactionByBlockNumberAndIndex**](docs/DefaultAPI.md#qangettransactionbyblocknumberandindex) | **Get** /getTransactionByBlockNumberAndIndex/{blockNumber}/{index}/ | 
*DefaultAPI* | [**QanGetTransactionByHash**](docs/DefaultAPI.md#qangettransactionbyhash) | **Get** /getTransactionByHash/{hash}/ | 
*DefaultAPI* | [**QanGetTransactionCount**](docs/DefaultAPI.md#qangettransactioncount) | **Get** /getTransactionCount/{Address}/{BlockNumber}/ | 
*DefaultAPI* | [**QanGetTransactionReceipt**](docs/DefaultAPI.md#qangettransactionreceipt) | **Get** /getTransactionReceipt/{Hash}/ | 
*DefaultAPI* | [**QanGetUncleCountByBlockHash**](docs/DefaultAPI.md#qangetunclecountbyblockhash) | **Get** /getUncleCountByBlockHash/{Hash}/ | 
*DefaultAPI* | [**QanGetUncleCountByBlockNumber**](docs/DefaultAPI.md#qangetunclecountbyblocknumber) | **Get** /getUncleCountByBlockNumber/{BlockNumber}/ | 
*DefaultAPI* | [**QanMaxPriorityFeePerGas**](docs/DefaultAPI.md#qanmaxpriorityfeepergas) | **Get** /maxPriorityFeePerGas/ | 
*DefaultAPI* | [**QanNewBlockFilter**](docs/DefaultAPI.md#qannewblockfilter) | **Get** /newBlockFilter/ | 
*DefaultAPI* | [**QanNewFilter**](docs/DefaultAPI.md#qannewfilter) | **Post** /newFilter/ | 
*DefaultAPI* | [**QanNewPendingTransactionFilter**](docs/DefaultAPI.md#qannewpendingtransactionfilter) | **Get** /newPendingTransactionFilter/ | 
*DefaultAPI* | [**QanSendRawTransaction**](docs/DefaultAPI.md#qansendrawtransaction) | **Post** /sendRawTransaction/ | 
*DefaultAPI* | [**QanSubscribe**](docs/DefaultAPI.md#qansubscribe) | **Post** /subscribe/ | 
*DefaultAPI* | [**QanSyncing**](docs/DefaultAPI.md#qansyncing) | **Get** /syncing/ | 
*DefaultAPI* | [**QanUninstallFilter**](docs/DefaultAPI.md#qanuninstallfilter) | **Get** /uninstallFilter/{FilterId}/ | 
*DefaultAPI* | [**QanUnsubscribe**](docs/DefaultAPI.md#qanunsubscribe) | **Get** /unsubscribe/{SubscriptionId}/ | 
*DefaultAPI* | [**QanXlinkValid**](docs/DefaultAPI.md#qanxlinkvalid) | **Get** /xlinkValid/ | 


## Documentation For Models

 - [ErrorDetail](docs/ErrorDetail.md)
 - [ErrorModel](docs/ErrorModel.md)
 - [EstimateGasObject](docs/EstimateGasObject.md)
 - [FilterObject](docs/FilterObject.md)
 - [InputCall](docs/InputCall.md)
 - [InputEstimateGas](docs/InputEstimateGas.md)
 - [InputFeeHistory](docs/InputFeeHistory.md)
 - [InputGetLogs](docs/InputGetLogs.md)
 - [InputGetProof](docs/InputGetProof.md)
 - [InputGetStorageAt](docs/InputGetStorageAt.md)
 - [InputNewFilter](docs/InputNewFilter.md)
 - [InputSendRawTransaction](docs/InputSendRawTransaction.md)
 - [InputSubscribe](docs/InputSubscribe.md)
 - [OutputAccounts](docs/OutputAccounts.md)
 - [OutputBlobBaseFee](docs/OutputBlobBaseFee.md)
 - [OutputBlockNumber](docs/OutputBlockNumber.md)
 - [OutputCall](docs/OutputCall.md)
 - [OutputChainId](docs/OutputChainId.md)
 - [OutputEstimateGas](docs/OutputEstimateGas.md)
 - [OutputFeeHistory](docs/OutputFeeHistory.md)
 - [OutputGasPrice](docs/OutputGasPrice.md)
 - [OutputGetAccount](docs/OutputGetAccount.md)
 - [OutputGetBalance](docs/OutputGetBalance.md)
 - [OutputGetBlockByHash](docs/OutputGetBlockByHash.md)
 - [OutputGetBlockByNumber](docs/OutputGetBlockByNumber.md)
 - [OutputGetBlockReceipts](docs/OutputGetBlockReceipts.md)
 - [OutputGetBlockTransactionCountByHash](docs/OutputGetBlockTransactionCountByHash.md)
 - [OutputGetBlockTransactionCountByNumber](docs/OutputGetBlockTransactionCountByNumber.md)
 - [OutputGetCode](docs/OutputGetCode.md)
 - [OutputGetFilterChanges](docs/OutputGetFilterChanges.md)
 - [OutputGetFilterLogs](docs/OutputGetFilterLogs.md)
 - [OutputGetLogs](docs/OutputGetLogs.md)
 - [OutputGetProof](docs/OutputGetProof.md)
 - [OutputGetStorageAt](docs/OutputGetStorageAt.md)
 - [OutputGetTransactionByBlockHashAndIndex](docs/OutputGetTransactionByBlockHashAndIndex.md)
 - [OutputGetTransactionByBlockNumberAndIndex](docs/OutputGetTransactionByBlockNumberAndIndex.md)
 - [OutputGetTransactionByHash](docs/OutputGetTransactionByHash.md)
 - [OutputGetTransactionCount](docs/OutputGetTransactionCount.md)
 - [OutputGetTransactionReceipt](docs/OutputGetTransactionReceipt.md)
 - [OutputGetUncleCountByBlockHash](docs/OutputGetUncleCountByBlockHash.md)
 - [OutputGetUncleCountByBlockNumber](docs/OutputGetUncleCountByBlockNumber.md)
 - [OutputMaxPriorityFeePerGas](docs/OutputMaxPriorityFeePerGas.md)
 - [OutputNewBlockFilter](docs/OutputNewBlockFilter.md)
 - [OutputNewFilter](docs/OutputNewFilter.md)
 - [OutputNewPendingTransactionFilter](docs/OutputNewPendingTransactionFilter.md)
 - [OutputSendRawTransaction](docs/OutputSendRawTransaction.md)
 - [OutputSubscribe](docs/OutputSubscribe.md)
 - [OutputSyncing](docs/OutputSyncing.md)
 - [OutputUninstallFilter](docs/OutputUninstallFilter.md)
 - [OutputUnsubscribe](docs/OutputUnsubscribe.md)
 - [ParamsTransaction](docs/ParamsTransaction.md)
 - [QanXlinkValidResponse](docs/QanXlinkValidResponse.md)
 - [ResponseBlock](docs/ResponseBlock.md)
 - [ResponseLog](docs/ResponseLog.md)
 - [ResponseStorageEntry](docs/ResponseStorageEntry.md)
 - [ResponseTransaction](docs/ResponseTransaction.md)
 - [ResponseTransactionReceipt](docs/ResponseTransactionReceipt.md)
 - [ResponseWithdrawals](docs/ResponseWithdrawals.md)
 - [SyncStatus](docs/SyncStatus.md)

## Acknowledgements

We would like to thank Smartbear and OpenAPITools tech for making building declarative APIs possible.
A huge benefit for the whole industry!
