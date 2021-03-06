/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// IncidentTeamUpdateData Incident Team data for an update request.
type IncidentTeamUpdateData struct {
	Attributes *IncidentTeamUpdateAttributes `json:"attributes,omitempty"`
	// The incident team's ID.
	Id            string                     `json:"id"`
	Relationships *IncidentTeamRelationships `json:"relationships,omitempty"`
	Type          IncidentTeamType           `json:"type"`
}

// NewIncidentTeamUpdateData instantiates a new IncidentTeamUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncidentTeamUpdateData(id string, type_ IncidentTeamType) *IncidentTeamUpdateData {
	this := IncidentTeamUpdateData{}
	this.Id = id
	this.Type = type_
	return &this
}

// NewIncidentTeamUpdateDataWithDefaults instantiates a new IncidentTeamUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncidentTeamUpdateDataWithDefaults() *IncidentTeamUpdateData {
	this := IncidentTeamUpdateData{}
	var type_ IncidentTeamType = "teams"
	this.Type = type_
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *IncidentTeamUpdateData) GetAttributes() IncidentTeamUpdateAttributes {
	if o == nil || o.Attributes == nil {
		var ret IncidentTeamUpdateAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTeamUpdateData) GetAttributesOk() (*IncidentTeamUpdateAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *IncidentTeamUpdateData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given IncidentTeamUpdateAttributes and assigns it to the Attributes field.
func (o *IncidentTeamUpdateData) SetAttributes(v IncidentTeamUpdateAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value
func (o *IncidentTeamUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IncidentTeamUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *IncidentTeamUpdateData) SetId(v string) {
	o.Id = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *IncidentTeamUpdateData) GetRelationships() IncidentTeamRelationships {
	if o == nil || o.Relationships == nil {
		var ret IncidentTeamRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTeamUpdateData) GetRelationshipsOk() (*IncidentTeamRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *IncidentTeamUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given IncidentTeamRelationships and assigns it to the Relationships field.
func (o *IncidentTeamUpdateData) SetRelationships(v IncidentTeamRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value
func (o *IncidentTeamUpdateData) GetType() IncidentTeamType {
	if o == nil {
		var ret IncidentTeamType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IncidentTeamUpdateData) GetTypeOk() (*IncidentTeamType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *IncidentTeamUpdateData) SetType(v IncidentTeamType) {
	o.Type = v
}

func (o IncidentTeamUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableIncidentTeamUpdateData struct {
	value *IncidentTeamUpdateData
	isSet bool
}

func (v NullableIncidentTeamUpdateData) Get() *IncidentTeamUpdateData {
	return v.value
}

func (v *NullableIncidentTeamUpdateData) Set(val *IncidentTeamUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableIncidentTeamUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableIncidentTeamUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncidentTeamUpdateData(val *IncidentTeamUpdateData) *NullableIncidentTeamUpdateData {
	return &NullableIncidentTeamUpdateData{value: val, isSet: true}
}

func (v NullableIncidentTeamUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncidentTeamUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
