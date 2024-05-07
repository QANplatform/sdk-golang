# OutputSubscribe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**SubscriptionId** | **string** | The hex encoded subscription ID. This ID will be attached to all received events and can also be used to cancel the subscription using eth_unsubscribe | 

## Methods

### NewOutputSubscribe

`func NewOutputSubscribe(subscriptionId string, ) *OutputSubscribe`

NewOutputSubscribe instantiates a new OutputSubscribe object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputSubscribeWithDefaults

`func NewOutputSubscribeWithDefaults() *OutputSubscribe`

NewOutputSubscribeWithDefaults instantiates a new OutputSubscribe object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *OutputSubscribe) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *OutputSubscribe) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *OutputSubscribe) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *OutputSubscribe) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *OutputSubscribe) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *OutputSubscribe) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *OutputSubscribe) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


