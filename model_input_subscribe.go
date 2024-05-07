/*
QAN AutoApi

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the InputSubscribe type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InputSubscribe{}

// InputSubscribe struct for InputSubscribe
type InputSubscribe struct {
	// A URL to the JSON Schema for this object.
	Schema *string `json:"$schema,omitempty"`
	Data interface{} `json:"Data"`
	// If true, method will return the full transaction data, otherwise only the transaction hash
	Flag bool `json:"Flag"`
	// The type of event you want to subscribe to (i.e., newHeads, logs, newPendingTransactions)
	SubscriptionName string `json:"SubscriptionName"`
}

type _InputSubscribe InputSubscribe

// NewInputSubscribe instantiates a new InputSubscribe object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputSubscribe(data interface{}, flag bool, subscriptionName string) *InputSubscribe {
	this := InputSubscribe{}
	this.Data = data
	this.Flag = flag
	this.SubscriptionName = subscriptionName
	return &this
}

// NewInputSubscribeWithDefaults instantiates a new InputSubscribe object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputSubscribeWithDefaults() *InputSubscribe {
	this := InputSubscribe{}
	var flag bool = false
	this.Flag = flag
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *InputSubscribe) GetSchema() string {
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputSubscribe) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *InputSubscribe) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *InputSubscribe) SetSchema(v string) {
	o.Schema = &v
}

// GetData returns the Data field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *InputSubscribe) GetData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InputSubscribe) GetDataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *InputSubscribe) SetData(v interface{}) {
	o.Data = v
}

// GetFlag returns the Flag field value
func (o *InputSubscribe) GetFlag() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Flag
}

// GetFlagOk returns a tuple with the Flag field value
// and a boolean to check if the value has been set.
func (o *InputSubscribe) GetFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Flag, true
}

// SetFlag sets field value
func (o *InputSubscribe) SetFlag(v bool) {
	o.Flag = v
}

// GetSubscriptionName returns the SubscriptionName field value
func (o *InputSubscribe) GetSubscriptionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionName
}

// GetSubscriptionNameOk returns a tuple with the SubscriptionName field value
// and a boolean to check if the value has been set.
func (o *InputSubscribe) GetSubscriptionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionName, true
}

// SetSubscriptionName sets field value
func (o *InputSubscribe) SetSubscriptionName(v string) {
	o.SubscriptionName = v
}

func (o InputSubscribe) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InputSubscribe) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schema) {
		toSerialize["$schema"] = o.Schema
	}
	if o.Data != nil {
		toSerialize["Data"] = o.Data
	}
	toSerialize["Flag"] = o.Flag
	toSerialize["SubscriptionName"] = o.SubscriptionName
	return toSerialize, nil
}

func (o *InputSubscribe) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Data",
		"Flag",
		"SubscriptionName",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varInputSubscribe := _InputSubscribe{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInputSubscribe)

	if err != nil {
		return err
	}

	*o = InputSubscribe(varInputSubscribe)

	return err
}

type NullableInputSubscribe struct {
	value *InputSubscribe
	isSet bool
}

func (v NullableInputSubscribe) Get() *InputSubscribe {
	return v.value
}

func (v *NullableInputSubscribe) Set(val *InputSubscribe) {
	v.value = val
	v.isSet = true
}

func (v NullableInputSubscribe) IsSet() bool {
	return v.isSet
}

func (v *NullableInputSubscribe) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputSubscribe(val *InputSubscribe) *NullableInputSubscribe {
	return &NullableInputSubscribe{value: val, isSet: true}
}

func (v NullableInputSubscribe) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputSubscribe) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

