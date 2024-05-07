# OutputAccounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**Accounts** | **[]string** | An array of addresses owned by the client | 

## Methods

### NewOutputAccounts

`func NewOutputAccounts(accounts []string, ) *OutputAccounts`

NewOutputAccounts instantiates a new OutputAccounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputAccountsWithDefaults

`func NewOutputAccountsWithDefaults() *OutputAccounts`

NewOutputAccountsWithDefaults instantiates a new OutputAccounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *OutputAccounts) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *OutputAccounts) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *OutputAccounts) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *OutputAccounts) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetAccounts

`func (o *OutputAccounts) GetAccounts() []string`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *OutputAccounts) GetAccountsOk() (*[]string, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *OutputAccounts) SetAccounts(v []string)`

SetAccounts sets Accounts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


