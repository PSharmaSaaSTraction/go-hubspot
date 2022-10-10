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

// SettingsPatchRequest Settings update request
type SettingsPatchRequest struct {
	// The name of your calling service to display to users.
	Name *string `json:"name,omitempty"`
	// The URL to your phone/calling UI, built with the [Calling SDK](#).
	Url *string `json:"url,omitempty"`
	// The target height of the iframe that will contain your phone/calling UI.
	Height *int32 `json:"height,omitempty"`
	// The target width of the iframe that will contain your phone/calling UI.
	Width *int32 `json:"width,omitempty"`
	// When true, your service will appear as an option under the *Call* action in contact records of connected accounts.
	IsReady *bool `json:"isReady,omitempty"`
	// When true, you are indicating that your service is compatible with engagement v2 service and can be used with custom objects.
	SupportsCustomObjects *bool `json:"supportsCustomObjects,omitempty"`
}

// NewSettingsPatchRequest instantiates a new SettingsPatchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettingsPatchRequest() *SettingsPatchRequest {
	this := SettingsPatchRequest{}
	return &this
}

// NewSettingsPatchRequestWithDefaults instantiates a new SettingsPatchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettingsPatchRequestWithDefaults() *SettingsPatchRequest {
	this := SettingsPatchRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SettingsPatchRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsPatchRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SettingsPatchRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SettingsPatchRequest) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *SettingsPatchRequest) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsPatchRequest) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *SettingsPatchRequest) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *SettingsPatchRequest) SetUrl(v string) {
	o.Url = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *SettingsPatchRequest) GetHeight() int32 {
	if o == nil || o.Height == nil {
		var ret int32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsPatchRequest) GetHeightOk() (*int32, bool) {
	if o == nil || o.Height == nil {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *SettingsPatchRequest) HasHeight() bool {
	if o != nil && o.Height != nil {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int32 and assigns it to the Height field.
func (o *SettingsPatchRequest) SetHeight(v int32) {
	o.Height = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *SettingsPatchRequest) GetWidth() int32 {
	if o == nil || o.Width == nil {
		var ret int32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsPatchRequest) GetWidthOk() (*int32, bool) {
	if o == nil || o.Width == nil {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *SettingsPatchRequest) HasWidth() bool {
	if o != nil && o.Width != nil {
		return true
	}

	return false
}

// SetWidth gets a reference to the given int32 and assigns it to the Width field.
func (o *SettingsPatchRequest) SetWidth(v int32) {
	o.Width = &v
}

// GetIsReady returns the IsReady field value if set, zero value otherwise.
func (o *SettingsPatchRequest) GetIsReady() bool {
	if o == nil || o.IsReady == nil {
		var ret bool
		return ret
	}
	return *o.IsReady
}

// GetIsReadyOk returns a tuple with the IsReady field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsPatchRequest) GetIsReadyOk() (*bool, bool) {
	if o == nil || o.IsReady == nil {
		return nil, false
	}
	return o.IsReady, true
}

// HasIsReady returns a boolean if a field has been set.
func (o *SettingsPatchRequest) HasIsReady() bool {
	if o != nil && o.IsReady != nil {
		return true
	}

	return false
}

// SetIsReady gets a reference to the given bool and assigns it to the IsReady field.
func (o *SettingsPatchRequest) SetIsReady(v bool) {
	o.IsReady = &v
}

// GetSupportsCustomObjects returns the SupportsCustomObjects field value if set, zero value otherwise.
func (o *SettingsPatchRequest) GetSupportsCustomObjects() bool {
	if o == nil || o.SupportsCustomObjects == nil {
		var ret bool
		return ret
	}
	return *o.SupportsCustomObjects
}

// GetSupportsCustomObjectsOk returns a tuple with the SupportsCustomObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsPatchRequest) GetSupportsCustomObjectsOk() (*bool, bool) {
	if o == nil || o.SupportsCustomObjects == nil {
		return nil, false
	}
	return o.SupportsCustomObjects, true
}

// HasSupportsCustomObjects returns a boolean if a field has been set.
func (o *SettingsPatchRequest) HasSupportsCustomObjects() bool {
	if o != nil && o.SupportsCustomObjects != nil {
		return true
	}

	return false
}

// SetSupportsCustomObjects gets a reference to the given bool and assigns it to the SupportsCustomObjects field.
func (o *SettingsPatchRequest) SetSupportsCustomObjects(v bool) {
	o.SupportsCustomObjects = &v
}

func (o SettingsPatchRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Url != nil {
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

type NullableSettingsPatchRequest struct {
	value *SettingsPatchRequest
	isSet bool
}

func (v NullableSettingsPatchRequest) Get() *SettingsPatchRequest {
	return v.value
}

func (v *NullableSettingsPatchRequest) Set(val *SettingsPatchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSettingsPatchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSettingsPatchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettingsPatchRequest(val *SettingsPatchRequest) *NullableSettingsPatchRequest {
	return &NullableSettingsPatchRequest{value: val, isSet: true}
}

func (v NullableSettingsPatchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettingsPatchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}