/*
Launcher API

The Launcher API is the execution layer for the Capsules framework.  It handles all the details of launching and monitoring runtime environments.

API version: 3.2.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package launcher

import (
	"encoding/json"
)

// WarehouseMetadata struct for WarehouseMetadata
type WarehouseMetadata struct {
	Owner *string `json:"owner,omitempty"`
	Created *string `json:"created,omitempty"`
	Modified *string `json:"modified,omitempty"`
	Version *string `json:"version,omitempty"`
	Tags *[]string `json:"tags,omitempty"`
	Server *string `json:"server,omitempty"`
	AdditionalPropertiesField *map[string]interface{} `json:"additionalProperties,omitempty"`
}

// NewWarehouseMetadata instantiates a new WarehouseMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWarehouseMetadata() *WarehouseMetadata {
	this := WarehouseMetadata{}
	return &this
}

// NewWarehouseMetadataWithDefaults instantiates a new WarehouseMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWarehouseMetadataWithDefaults() *WarehouseMetadata {
	this := WarehouseMetadata{}
	return &this
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *WarehouseMetadata) GetOwner() string {
	if o == nil || o.Owner == nil {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarehouseMetadata) GetOwnerOk() (*string, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *WarehouseMetadata) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *WarehouseMetadata) SetOwner(v string) {
	o.Owner = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *WarehouseMetadata) GetCreated() string {
	if o == nil || o.Created == nil {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarehouseMetadata) GetCreatedOk() (*string, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *WarehouseMetadata) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *WarehouseMetadata) SetCreated(v string) {
	o.Created = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *WarehouseMetadata) GetModified() string {
	if o == nil || o.Modified == nil {
		var ret string
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarehouseMetadata) GetModifiedOk() (*string, bool) {
	if o == nil || o.Modified == nil {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *WarehouseMetadata) HasModified() bool {
	if o != nil && o.Modified != nil {
		return true
	}

	return false
}

// SetModified gets a reference to the given string and assigns it to the Modified field.
func (o *WarehouseMetadata) SetModified(v string) {
	o.Modified = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *WarehouseMetadata) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarehouseMetadata) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *WarehouseMetadata) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *WarehouseMetadata) SetVersion(v string) {
	o.Version = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WarehouseMetadata) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarehouseMetadata) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WarehouseMetadata) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *WarehouseMetadata) SetTags(v []string) {
	o.Tags = &v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *WarehouseMetadata) GetServer() string {
	if o == nil || o.Server == nil {
		var ret string
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarehouseMetadata) GetServerOk() (*string, bool) {
	if o == nil || o.Server == nil {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *WarehouseMetadata) HasServer() bool {
	if o != nil && o.Server != nil {
		return true
	}

	return false
}

// SetServer gets a reference to the given string and assigns it to the Server field.
func (o *WarehouseMetadata) SetServer(v string) {
	o.Server = &v
}

// GetAdditionalPropertiesField returns the AdditionalPropertiesField field value if set, zero value otherwise.
func (o *WarehouseMetadata) GetAdditionalPropertiesField() map[string]interface{} {
	if o == nil || o.AdditionalPropertiesField == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.AdditionalPropertiesField
}

// GetAdditionalPropertiesFieldOk returns a tuple with the AdditionalPropertiesField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarehouseMetadata) GetAdditionalPropertiesFieldOk() (*map[string]interface{}, bool) {
	if o == nil || o.AdditionalPropertiesField == nil {
		return nil, false
	}
	return o.AdditionalPropertiesField, true
}

// HasAdditionalPropertiesField returns a boolean if a field has been set.
func (o *WarehouseMetadata) HasAdditionalPropertiesField() bool {
	if o != nil && o.AdditionalPropertiesField != nil {
		return true
	}

	return false
}

// SetAdditionalPropertiesField gets a reference to the given map[string]interface{} and assigns it to the AdditionalPropertiesField field.
func (o *WarehouseMetadata) SetAdditionalPropertiesField(v map[string]interface{}) {
	o.AdditionalPropertiesField = &v
}

func (o WarehouseMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Modified != nil {
		toSerialize["modified"] = o.Modified
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Server != nil {
		toSerialize["server"] = o.Server
	}
	if o.AdditionalPropertiesField != nil {
		toSerialize["additionalProperties"] = o.AdditionalPropertiesField
	}
	return json.Marshal(toSerialize)
}

type NullableWarehouseMetadata struct {
	value *WarehouseMetadata
	isSet bool
}

func (v NullableWarehouseMetadata) Get() *WarehouseMetadata {
	return v.value
}

func (v *NullableWarehouseMetadata) Set(val *WarehouseMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableWarehouseMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableWarehouseMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWarehouseMetadata(val *WarehouseMetadata) *NullableWarehouseMetadata {
	return &NullableWarehouseMetadata{value: val, isSet: true}
}

func (v NullableWarehouseMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWarehouseMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


