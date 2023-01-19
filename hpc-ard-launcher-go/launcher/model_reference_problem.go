/*
Launcher API

The Launcher API is the execution layer for the Capsules framework.  It handles all the details of launching and monitoring runtime environments.

API version: 3.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package launcher

import (
	"encoding/json"
)

// ReferenceProblem struct for ReferenceProblem
type ReferenceProblem struct {
	// An absolute URI that identifies the problem type.  When dereferenced, it SHOULD provide human-readable documentation for the problem type (e.g., using HTML). 
	Type *string `json:"type,omitempty"`
	// A short, summary of the problem type. Written in english and readable for engineers (usually not suited for non technical stakeholders and not localized); example: Service Unavailable 
	Title *string `json:"title,omitempty"`
	// The HTTP status code generated by the origin server for this occurrence of the problem. 
	Status *int32 `json:"status,omitempty"`
	// A human readable explanation specific to this occurrence of the problem. 
	Detail *string `json:"detail,omitempty"`
	// An absolute URI that identifies the specific occurrence of the problem. It may or may not yield further information if dereferenced. 
	Instance *string `json:"instance,omitempty"`
	// The value of the reference resulting in the problem
	Reference *string `json:"reference,omitempty"`
}

// NewReferenceProblem instantiates a new ReferenceProblem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReferenceProblem() *ReferenceProblem {
	this := ReferenceProblem{}
	var type_ string = "about:blank"
	this.Type = &type_
	return &this
}

// NewReferenceProblemWithDefaults instantiates a new ReferenceProblem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReferenceProblemWithDefaults() *ReferenceProblem {
	this := ReferenceProblem{}
	var type_ string = "about:blank"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ReferenceProblem) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceProblem) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ReferenceProblem) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ReferenceProblem) SetType(v string) {
	o.Type = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ReferenceProblem) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceProblem) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ReferenceProblem) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ReferenceProblem) SetTitle(v string) {
	o.Title = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ReferenceProblem) GetStatus() int32 {
	if o == nil || o.Status == nil {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceProblem) GetStatusOk() (*int32, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ReferenceProblem) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *ReferenceProblem) SetStatus(v int32) {
	o.Status = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *ReferenceProblem) GetDetail() string {
	if o == nil || o.Detail == nil {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceProblem) GetDetailOk() (*string, bool) {
	if o == nil || o.Detail == nil {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *ReferenceProblem) HasDetail() bool {
	if o != nil && o.Detail != nil {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *ReferenceProblem) SetDetail(v string) {
	o.Detail = &v
}

// GetInstance returns the Instance field value if set, zero value otherwise.
func (o *ReferenceProblem) GetInstance() string {
	if o == nil || o.Instance == nil {
		var ret string
		return ret
	}
	return *o.Instance
}

// GetInstanceOk returns a tuple with the Instance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceProblem) GetInstanceOk() (*string, bool) {
	if o == nil || o.Instance == nil {
		return nil, false
	}
	return o.Instance, true
}

// HasInstance returns a boolean if a field has been set.
func (o *ReferenceProblem) HasInstance() bool {
	if o != nil && o.Instance != nil {
		return true
	}

	return false
}

// SetInstance gets a reference to the given string and assigns it to the Instance field.
func (o *ReferenceProblem) SetInstance(v string) {
	o.Instance = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *ReferenceProblem) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceProblem) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *ReferenceProblem) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *ReferenceProblem) SetReference(v string) {
	o.Reference = &v
}

func (o ReferenceProblem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Detail != nil {
		toSerialize["detail"] = o.Detail
	}
	if o.Instance != nil {
		toSerialize["instance"] = o.Instance
	}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	return json.Marshal(toSerialize)
}

type NullableReferenceProblem struct {
	value *ReferenceProblem
	isSet bool
}

func (v NullableReferenceProblem) Get() *ReferenceProblem {
	return v.value
}

func (v *NullableReferenceProblem) Set(val *ReferenceProblem) {
	v.value = val
	v.isSet = true
}

func (v NullableReferenceProblem) IsSet() bool {
	return v.isSet
}

func (v *NullableReferenceProblem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReferenceProblem(val *ReferenceProblem) *NullableReferenceProblem {
	return &NullableReferenceProblem{value: val, isSet: true}
}

func (v NullableReferenceProblem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReferenceProblem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


