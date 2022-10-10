/*
Calling Extensions API

Provides a way for apps to add custom calling options to a contact record. This works in conjunction with the [Calling SDK](#), which is used to build your phone/calling UI. The endpoints here allow your service to appear as an option to HubSpot users when they access the *Call* action on a contact record. Once accessed, your custom phone/calling UI will be displayed in an iframe at the specified URL with the specified dimensions on that record.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package calling

import (
	"encoding/json"
)

// SettingsRequest Settings create request
type SettingsRequest struct {
	// The name of your calling service to display to users.
	Name string `json:"name"`
	// The URL to your phone/calling UI, built with the [Calling SDK](#).
	Url string `json:"url"`
	// The target height of the iframe that will contain your phone/calling UI.
	Height *int32 `json:"height,omitempty"`
	// The target width of the iframe that will contain your phone/calling UI.
	Width *int32 `json:"width,omitempty"`
	// When true, your service will appear as an option under the *Call* action in contact records of connected accounts.
	IsReady *bool `json:"isReady,omitempty"`
	// When true, you are indicating that your service is compatible with engagement v2 service and can be used with custom objects.
	SupportsCustomObjects *bool `json:"supportsCustomObjects,omitempty"`
}

// NewSettingsRequest instantiates a new SettingsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettingsRequest(name string, url string) *SettingsRequest {
	this := SettingsRequest{}
	this.Name = name
	this.Url = url
	return &this
}

// NewSettingsRequestWithDefaults instantiates a new SettingsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettingsRequestWithDefaults() *SettingsRequest {
	this := SettingsRequest{}
	return &this
}

// GetName returns the Name field value
func (o *SettingsRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SettingsRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SettingsRequest) SetName(v string) {
	o.Name = v
}

// GetUrl returns the Url field value
func (o *SettingsRequest) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *SettingsRequest) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *SettingsRequest) SetUrl(v string) {
	o.Url = v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *SettingsRequest) GetHeight() int32 {
	if o == nil || o.Height == nil {
		var ret int32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsRequest) GetHeightOk() (*int32, bool) {
	if o == nil || o.Height == nil {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *SettingsRequest) HasHeight() bool {
	if o != nil && o.Height != nil {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int32 and assigns it to the Height field.
func (o *SettingsRequest) SetHeight(v int32) {
	o.Height = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *SettingsRequest) GetWidth() int32 {
	if o == nil || o.Width == nil {
		var ret int32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsRequest) GetWidthOk() (*int32, bool) {
	if o == nil || o.Width == nil {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *SettingsRequest) HasWidth() bool {
	if o != nil && o.Width != nil {
		return true
	}

	return false
}

// SetWidth gets a reference to the given int32 and assigns it to the Width field.
func (o *SettingsRequest) SetWidth(v int32) {
	o.Width = &v
}

// GetIsReady returns the IsReady field value if set, zero value otherwise.
func (o *SettingsRequest) GetIsReady() bool {
	if o == nil || o.IsReady == nil {
		var ret bool
		return ret
	}
	return *o.IsReady
}

// GetIsReadyOk returns a tuple with the IsReady field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsRequest) GetIsReadyOk() (*bool, bool) {
	if o == nil || o.IsReady == nil {
		return nil, false
	}
	return o.IsReady, true
}

// HasIsReady returns a boolean if a field has been set.
func (o *SettingsRequest) HasIsReady() bool {
	if o != nil && o.IsReady != nil {
		return true
	}

	return false
}

// SetIsReady gets a reference to the given bool and assigns it to the IsReady field.
func (o *SettingsRequest) SetIsReady(v bool) {
	o.IsReady = &v
}

// GetSupportsCustomObjects returns the SupportsCustomObjects field value if set, zero value otherwise.
func (o *SettingsRequest) GetSupportsCustomObjects() bool {
	if o == nil || o.SupportsCustomObjects == nil {
		var ret bool
		return ret
	}
	return *o.SupportsCustomObjects
}

// GetSupportsCustomObjectsOk returns a tuple with the SupportsCustomObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsRequest) GetSupportsCustomObjectsOk() (*bool, bool) {
	if o == nil || o.SupportsCustomObjects == nil {
		return nil, false
	}
	return o.SupportsCustomObjects, true
}

// HasSupportsCustomObjects returns a boolean if a field has been set.
func (o *SettingsRequest) HasSupportsCustomObjects() bool {
	if o != nil && o.SupportsCustomObjects != nil {
		return true
	}

	return false
}

// SetSupportsCustomObjects gets a reference to the given bool and assigns it to the SupportsCustomObjects field.
func (o *SettingsRequest) SetSupportsCustomObjects(v bool) {
	o.SupportsCustomObjects = &v
}

func (o SettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if o.Height != nil {
		toSerialize["height"] = o.Height
	}
	if o.Width != nil {
		toSerialize["width"] = o.Width
	}
	if o.IsReady != nil {
		toSerialize["isReady"] = o.IsReady
	}
	if o.SupportsCustomObjects != nil {
		toSerialize["supportsCustomObjects"] = o.SupportsCustomObjects
	}
	return json.Marshal(toSerialize)
}

type NullableSettingsRequest struct {
	value *SettingsRequest
	isSet bool
}

func (v NullableSettingsRequest) Get() *SettingsRequest {
	return v.value
}

func (v *NullableSettingsRequest) Set(val *SettingsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSettingsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSettingsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettingsRequest(val *SettingsRequest) *NullableSettingsRequest {
	return &NullableSettingsRequest{value: val, isSet: true}
}

func (v NullableSettingsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettingsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}