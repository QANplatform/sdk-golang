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

// checks if the InputFeeHistory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InputFeeHistory{}

// InputFeeHistory struct for InputFeeHistory
type InputFeeHistory struct {
	// A URL to the JSON Schema for this object.
	Schema *string `json:"$schema,omitempty"`
	// The number of blocks in the requested range. Between 1 and 1024 blocks can be requested in a single query. It will return less than the requested range if not all blocks are available
	BlockCount int32 `json:"BlockCount"`
	// The highest number block of the requested range in hexadecimal format or tags. The supported tag values include earliest for the earliest/genesis block, latest for the latest mined block, pending for the pending state/transactions.
	NewestBlock string `json:"NewestBlock"`
	// A list of percentile values with a monotonic increase in value. The transactions will be ranked by effective tip per gas for each block in the requested range, and the corresponding effective tip for the percentile will be calculated while taking gas consumption into consideration
	RewardPercentiles []int32 `json:"RewardPercentiles"`
}

type _InputFeeHistory InputFeeHistory

// NewInputFeeHistory instantiates a new InputFeeHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputFeeHistory(blockCount int32, newestBlock string, rewardPercentiles []int32) *InputFeeHistory {
	this := InputFeeHistory{}
	this.BlockCount = blockCount
	this.NewestBlock = newestBlock
	this.RewardPercentiles = rewardPercentiles
	return &this
}

// NewInputFeeHistoryWithDefaults instantiates a new InputFeeHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputFeeHistoryWithDefaults() *InputFeeHistory {
	this := InputFeeHistory{}
	var blockCount int32 = 2
	this.BlockCount = blockCount
	var newestBlock string = "latest"
	this.NewestBlock = newestBlock
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *InputFeeHistory) GetSchema() string {
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputFeeHistory) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *InputFeeHistory) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *InputFeeHistory) SetSchema(v string) {
	o.Schema = &v
}

// GetBlockCount returns the BlockCount field value
func (o *InputFeeHistory) GetBlockCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BlockCount
}

// GetBlockCountOk returns a tuple with the BlockCount field value
// and a boolean to check if the value has been set.
func (o *InputFeeHistory) GetBlockCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockCount, true
}

// SetBlockCount sets field value
func (o *InputFeeHistory) SetBlockCount(v int32) {
	o.BlockCount = v
}

// GetNewestBlock returns the NewestBlock field value
func (o *InputFeeHistory) GetNewestBlock() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NewestBlock
}

// GetNewestBlockOk returns a tuple with the NewestBlock field value
// and a boolean to check if the value has been set.
func (o *InputFeeHistory) GetNewestBlockOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewestBlock, true
}

// SetNewestBlock sets field value
func (o *InputFeeHistory) SetNewestBlock(v string) {
	o.NewestBlock = v
}

// GetRewardPercentiles returns the RewardPercentiles field value
func (o *InputFeeHistory) GetRewardPercentiles() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.RewardPercentiles
}

// GetRewardPercentilesOk returns a tuple with the RewardPercentiles field value
// and a boolean to check if the value has been set.
func (o *InputFeeHistory) GetRewardPercentilesOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RewardPercentiles, true
}

// SetRewardPercentiles sets field value
func (o *InputFeeHistory) SetRewardPercentiles(v []int32) {
	o.RewardPercentiles = v
}

func (o InputFeeHistory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InputFeeHistory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schema) {
		toSerialize["$schema"] = o.Schema
	}
	toSerialize["BlockCount"] = o.BlockCount
	toSerialize["NewestBlock"] = o.NewestBlock
	toSerialize["RewardPercentiles"] = o.RewardPercentiles
	return toSerialize, nil
}

func (o *InputFeeHistory) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"BlockCount",
		"NewestBlock",
		"RewardPercentiles",
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

	varInputFeeHistory := _InputFeeHistory{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInputFeeHistory)

	if err != nil {
		return err
	}

	*o = InputFeeHistory(varInputFeeHistory)

	return err
}

type NullableInputFeeHistory struct {
	value *InputFeeHistory
	isSet bool
}

func (v NullableInputFeeHistory) Get() *InputFeeHistory {
	return v.value
}

func (v *NullableInputFeeHistory) Set(val *InputFeeHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableInputFeeHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableInputFeeHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputFeeHistory(val *InputFeeHistory) *NullableInputFeeHistory {
	return &NullableInputFeeHistory{value: val, isSet: true}
}

func (v NullableInputFeeHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputFeeHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

