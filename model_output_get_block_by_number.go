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

// checks if the OutputGetBlockByNumber type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OutputGetBlockByNumber{}

// OutputGetBlockByNumber struct for OutputGetBlockByNumber
type OutputGetBlockByNumber struct {
	// A URL to the JSON Schema for this object.
	Schema *string `json:"$schema,omitempty"`
	// A block object, or null when no block was found
	Block ResponseBlock `json:"Block"`
}

type _OutputGetBlockByNumber OutputGetBlockByNumber

// NewOutputGetBlockByNumber instantiates a new OutputGetBlockByNumber object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutputGetBlockByNumber(block ResponseBlock) *OutputGetBlockByNumber {
	this := OutputGetBlockByNumber{}
	this.Block = block
	return &this
}

// NewOutputGetBlockByNumberWithDefaults instantiates a new OutputGetBlockByNumber object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutputGetBlockByNumberWithDefaults() *OutputGetBlockByNumber {
	this := OutputGetBlockByNumber{}
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *OutputGetBlockByNumber) GetSchema() string {
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputGetBlockByNumber) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *OutputGetBlockByNumber) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *OutputGetBlockByNumber) SetSchema(v string) {
	o.Schema = &v
}

// GetBlock returns the Block field value
func (o *OutputGetBlockByNumber) GetBlock() ResponseBlock {
	if o == nil {
		var ret ResponseBlock
		return ret
	}

	return o.Block
}

// GetBlockOk returns a tuple with the Block field value
// and a boolean to check if the value has been set.
func (o *OutputGetBlockByNumber) GetBlockOk() (*ResponseBlock, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Block, true
}

// SetBlock sets field value
func (o *OutputGetBlockByNumber) SetBlock(v ResponseBlock) {
	o.Block = v
}

func (o OutputGetBlockByNumber) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OutputGetBlockByNumber) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schema) {
		toSerialize["$schema"] = o.Schema
	}
	toSerialize["Block"] = o.Block
	return toSerialize, nil
}

func (o *OutputGetBlockByNumber) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Block",
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

	varOutputGetBlockByNumber := _OutputGetBlockByNumber{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOutputGetBlockByNumber)

	if err != nil {
		return err
	}

	*o = OutputGetBlockByNumber(varOutputGetBlockByNumber)

	return err
}

type NullableOutputGetBlockByNumber struct {
	value *OutputGetBlockByNumber
	isSet bool
}

func (v NullableOutputGetBlockByNumber) Get() *OutputGetBlockByNumber {
	return v.value
}

func (v *NullableOutputGetBlockByNumber) Set(val *OutputGetBlockByNumber) {
	v.value = val
	v.isSet = true
}

func (v NullableOutputGetBlockByNumber) IsSet() bool {
	return v.isSet
}

func (v *NullableOutputGetBlockByNumber) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutputGetBlockByNumber(val *OutputGetBlockByNumber) *NullableOutputGetBlockByNumber {
	return &NullableOutputGetBlockByNumber{value: val, isSet: true}
}

func (v NullableOutputGetBlockByNumber) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutputGetBlockByNumber) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

