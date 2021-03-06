/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
)

// WidgetEventSize Size to use to display an event.
type WidgetEventSize string

// List of WidgetEventSize
const (
	WIDGETEVENTSIZE_SMALL WidgetEventSize = "s"
	WIDGETEVENTSIZE_LARGE WidgetEventSize = "l"
)

func (v *WidgetEventSize) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WidgetEventSize(value)
	for _, existing := range []WidgetEventSize{"s", "l"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WidgetEventSize", value)
}

// Ptr returns reference to WidgetEventSize value
func (v WidgetEventSize) Ptr() *WidgetEventSize {
	return &v
}

type NullableWidgetEventSize struct {
	value *WidgetEventSize
	isSet bool
}

func (v NullableWidgetEventSize) Get() *WidgetEventSize {
	return v.value
}

func (v *NullableWidgetEventSize) Set(val *WidgetEventSize) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetEventSize) IsSet() bool {
	return v.isSet
}

func (v *NullableWidgetEventSize) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetEventSize(val *WidgetEventSize) *NullableWidgetEventSize {
	return &NullableWidgetEventSize{value: val, isSet: true}
}

func (v NullableWidgetEventSize) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetEventSize) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
