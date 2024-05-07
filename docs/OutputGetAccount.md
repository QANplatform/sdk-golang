# OutputGetAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**Balance** | **string** | The current balance of the account in wei | 
**CodeHash** | **string** | A 32 byte hash of the code of the account | 
**Nonce** | **string** | The transaction count of an account | 
**StorageRoot** | **string** | The hash of the account&#39;s storage data | 

## Methods

### NewOutputGetAccount

`func NewOutputGetAccount(balance string, codeHash string, nonce string, storageRoot string, ) *OutputGetAccount`

NewOutputGetAccount instantiates a new OutputGetAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputGetAccountWithDefaults

`func NewOutputGetAccountWithDefaults() *OutputGetAccount`

NewOutputGetAccountWithDefaults instantiates a new OutputGetAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *OutputGetAccount) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *OutputGetAccount) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *OutputGetAccount) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *OutputGetAccount) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetBalance

`func (o *OutputGetAccount) GetBalance() string`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *OutputGetAccount) GetBalanceOk() (*string, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *OutputGetAccount) SetBalance(v string)`

SetBalance sets Balance field to given value.


### GetCodeHash

`func (o *OutputGetAccount) GetCodeHash() string`

GetCodeHash returns the CodeHash field if non-nil, zero value otherwise.

### GetCodeHashOk

`func (o *OutputGetAccount) GetCodeHashOk() (*string, bool)`

GetCodeHashOk returns a tuple with the CodeHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeHash

`func (o *OutputGetAccount) SetCodeHash(v string)`

SetCodeHash sets CodeHash field to given value.


### GetNonce

`func (o *OutputGetAccount) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *OutputGetAccount) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *OutputGetAccount) SetNonce(v string)`

SetNonce sets Nonce field to given value.


### GetStorageRoot

`func (o *OutputGetAccount) GetStorageRoot() string`

GetStorageRoot returns the StorageRoot field if non-nil, zero value otherwise.

### GetStorageRootOk

`func (o *OutputGetAccount) GetStorageRootOk() (*string, bool)`

GetStorageRootOk returns a tuple with the StorageRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageRoot

`func (o *OutputGetAccount) SetStorageRoot(v string)`

SetStorageRoot sets StorageRoot field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


