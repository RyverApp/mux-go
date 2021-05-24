/*
 * Mux API
 *
 * Mux is how developers build online video. This API encompasses both Mux Video and Mux Data functionality to help you build your video-related projects better and faster than ever before. 
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package muxgo

import (
	"encoding/json"
)

// ListInsightsResponse struct for ListInsightsResponse
type ListInsightsResponse struct {
	Data *[]Insight `json:"data,omitempty"`
	TotalRowCount *int64 `json:"total_row_count,omitempty"`
	Timeframe *[]int64 `json:"timeframe,omitempty"`
}

// NewListInsightsResponse instantiates a new ListInsightsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListInsightsResponse() *ListInsightsResponse {
	this := ListInsightsResponse{}
	return &this
}

// NewListInsightsResponseWithDefaults instantiates a new ListInsightsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListInsightsResponseWithDefaults() *ListInsightsResponse {
	this := ListInsightsResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListInsightsResponse) GetData() []Insight {
	if o == nil || o.Data == nil {
		var ret []Insight
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInsightsResponse) GetDataOk() (*[]Insight, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListInsightsResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []Insight and assigns it to the Data field.
func (o *ListInsightsResponse) SetData(v []Insight) {
	o.Data = &v
}

// GetTotalRowCount returns the TotalRowCount field value if set, zero value otherwise.
func (o *ListInsightsResponse) GetTotalRowCount() int64 {
	if o == nil || o.TotalRowCount == nil {
		var ret int64
		return ret
	}
	return *o.TotalRowCount
}

// GetTotalRowCountOk returns a tuple with the TotalRowCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInsightsResponse) GetTotalRowCountOk() (*int64, bool) {
	if o == nil || o.TotalRowCount == nil {
		return nil, false
	}
	return o.TotalRowCount, true
}

// HasTotalRowCount returns a boolean if a field has been set.
func (o *ListInsightsResponse) HasTotalRowCount() bool {
	if o != nil && o.TotalRowCount != nil {
		return true
	}

	return false
}

// SetTotalRowCount gets a reference to the given int64 and assigns it to the TotalRowCount field.
func (o *ListInsightsResponse) SetTotalRowCount(v int64) {
	o.TotalRowCount = &v
}

// GetTimeframe returns the Timeframe field value if set, zero value otherwise.
func (o *ListInsightsResponse) GetTimeframe() []int64 {
	if o == nil || o.Timeframe == nil {
		var ret []int64
		return ret
	}
	return *o.Timeframe
}

// GetTimeframeOk returns a tuple with the Timeframe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInsightsResponse) GetTimeframeOk() (*[]int64, bool) {
	if o == nil || o.Timeframe == nil {
		return nil, false
	}
	return o.Timeframe, true
}

// HasTimeframe returns a boolean if a field has been set.
func (o *ListInsightsResponse) HasTimeframe() bool {
	if o != nil && o.Timeframe != nil {
		return true
	}

	return false
}

// SetTimeframe gets a reference to the given []int64 and assigns it to the Timeframe field.
func (o *ListInsightsResponse) SetTimeframe(v []int64) {
	o.Timeframe = &v
}

func (o ListInsightsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.TotalRowCount != nil {
		toSerialize["total_row_count"] = o.TotalRowCount
	}
	if o.Timeframe != nil {
		toSerialize["timeframe"] = o.Timeframe
	}
	return json.Marshal(toSerialize)
}

type NullableListInsightsResponse struct {
	value *ListInsightsResponse
	isSet bool
}

func (v NullableListInsightsResponse) Get() *ListInsightsResponse {
	return v.value
}

func (v *NullableListInsightsResponse) Set(val *ListInsightsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListInsightsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListInsightsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListInsightsResponse(val *ListInsightsResponse) *NullableListInsightsResponse {
	return &NullableListInsightsResponse{value: val, isSet: true}
}

func (v NullableListInsightsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListInsightsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


