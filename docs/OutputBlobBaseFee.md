# OutputBlobBaseFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**BaseFee** | **string** | The expected base fee in wei, represented in decimal format | 

## Methods

### NewOutputBlobBaseFee

`func NewOutputBlobBaseFee(baseFee string, ) *OutputBlobBaseFee`

NewOutputBlobBaseFee instantiates a new OutputBlobBaseFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputBlobBaseFeeWithDefaults

`func NewOutputBlobBaseFeeWithDefaults() *OutputBlobBaseFee`

NewOutputBlobBaseFeeWithDefaults instantiates a new OutputBlobBaseFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *OutputBlobBaseFee) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *OutputBlobBaseFee) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *OutputBlobBaseFee) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *OutputBlobBaseFee) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetBaseFee

`func (o *OutputBlobBaseFee) GetBaseFee() string`

GetBaseFee returns the BaseFee field if non-nil, zero value otherwise.

### GetBaseFeeOk

`func (o *OutputBlobBaseFee) GetBaseFeeOk() (*string, bool)`

GetBaseFeeOk returns a tuple with the BaseFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseFee

`func (o *OutputBlobBaseFee) SetBaseFee(v string)`

SetBaseFee sets BaseFee field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


