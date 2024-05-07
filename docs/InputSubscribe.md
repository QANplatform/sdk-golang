# InputSubscribe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**Data** | **interface{}** |  | 
**Flag** | **bool** | If true, method will return the full transaction data, otherwise only the transaction hash | [default to false]
**SubscriptionName** | **string** | The type of event you want to subscribe to (i.e., newHeads, logs, newPendingTransactions) | 

## Methods

### NewInputSubscribe

`func NewInputSubscribe(data interface{}, flag bool, subscriptionName string, ) *InputSubscribe`

NewInputSubscribe instantiates a new InputSubscribe object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputSubscribeWithDefaults

`func NewInputSubscribeWithDefaults() *InputSubscribe`

NewInputSubscribeWithDefaults instantiates a new InputSubscribe object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *InputSubscribe) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *InputSubscribe) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *InputSubscribe) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *InputSubscribe) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetData

`func (o *InputSubscribe) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InputSubscribe) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InputSubscribe) SetData(v interface{})`

SetData sets Data field to given value.


### SetDataNil

`func (o *InputSubscribe) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *InputSubscribe) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetFlag

`func (o *InputSubscribe) GetFlag() bool`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *InputSubscribe) GetFlagOk() (*bool, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *InputSubscribe) SetFlag(v bool)`

SetFlag sets Flag field to given value.


### GetSubscriptionName

`func (o *InputSubscribe) GetSubscriptionName() string`

GetSubscriptionName returns the SubscriptionName field if non-nil, zero value otherwise.

### GetSubscriptionNameOk

`func (o *InputSubscribe) GetSubscriptionNameOk() (*string, bool)`

GetSubscriptionNameOk returns a tuple with the SubscriptionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionName

`func (o *InputSubscribe) SetSubscriptionName(v string)`

SetSubscriptionName sets SubscriptionName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


