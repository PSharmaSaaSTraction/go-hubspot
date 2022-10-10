/*
Marketing Events Extension

These APIs allow you to interact with HubSpot's Marketing Events Extension. It allows you to: * Create, Read or update Marketing Event information in HubSpot * Specify whether a HubSpot contact has registered, attended or cancelled a registration to a Marketing Event. * Specify a URL that can be called to get the details of a Marketing Event.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package marketing_events_beta

import (
	"encoding/json"
)

// MarketingEventEmailSubscriber struct for MarketingEventEmailSubscriber
type MarketingEventEmailSubscriber struct {
	// The date and time at which the contact subscribed to the event.
	InteractionDateTime int64              `json:"interactionDateTime"`
	Properties          *map[string]string `json:"properties,omitempty"`
	// The email address of the contact in HubSpot to associate with the event. Note that the contact must already exist in HubSpot; a contact will not be created.
	Email             string             `json:"email"`
	ContactProperties *map[string]string `json:"contactProperties,omitempty"`
}

// NewMarketingEventEmailSubscriber instantiates a new MarketingEventEmailSubscriber object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketingEventEmailSubscriber(interactionDateTime int64, email string) *MarketingEventEmailSubscriber {
	this := MarketingEventEmailSubscriber{}
	this.InteractionDateTime = interactionDateTime
	this.Email = email
	return &this
}

// NewMarketingEventEmailSubscriberWithDefaults instantiates a new MarketingEventEmailSubscriber object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketingEventEmailSubscriberWithDefaults() *MarketingEventEmailSubscriber {
	this := MarketingEventEmailSubscriber{}
	return &this
}

// GetInteractionDateTime returns the InteractionDateTime field value
func (o *MarketingEventEmailSubscriber) GetInteractionDateTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.InteractionDateTime
}

// GetInteractionDateTimeOk returns a tuple with the InteractionDateTime field value
// and a boolean to check if the value has been set.
func (o *MarketingEventEmailSubscriber) GetInteractionDateTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InteractionDateTime, true
}

// SetInteractionDateTime sets field value
func (o *MarketingEventEmailSubscriber) SetInteractionDateTime(v int64) {
	o.InteractionDateTime = v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *MarketingEventEmailSubscriber) GetProperties() map[string]string {
	if o == nil || o.Properties == nil {
		var ret map[string]string
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingEventEmailSubscriber) GetPropertiesOk() (*map[string]string, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *MarketingEventEmailSubscriber) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]string and assigns it to the Properties field.
func (o *MarketingEventEmailSubscriber) SetProperties(v map[string]string) {
	o.Properties = &v
}

// GetEmail returns the Email field value
func (o *MarketingEventEmailSubscriber) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *MarketingEventEmailSubscriber) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *MarketingEventEmailSubscriber) SetEmail(v string) {
	o.Email = v
}

// GetContactProperties returns the ContactProperties field value if set, zero value otherwise.
func (o *MarketingEventEmailSubscriber) GetContactProperties() map[string]string {
	if o == nil || o.ContactProperties == nil {
		var ret map[string]string
		return ret
	}
	return *o.ContactProperties
}

// GetContactPropertiesOk returns a tuple with the ContactProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingEventEmailSubscriber) GetContactPropertiesOk() (*map[string]string, bool) {
	if o == nil || o.ContactProperties == nil {
		return nil, false
	}
	return o.ContactProperties, true
}

// HasContactProperties returns a boolean if a field has been set.
func (o *MarketingEventEmailSubscriber) HasContactProperties() bool {
	if o != nil && o.ContactProperties != nil {
		return true
	}

	return false
}

// SetContactProperties gets a reference to the given map[string]string and assigns it to the ContactProperties field.
func (o *MarketingEventEmailSubscriber) SetContactProperties(v map[string]string) {
	o.ContactProperties = &v
}

func (o MarketingEventEmailSubscriber) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["interactionDateTime"] = o.InteractionDateTime
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if true {
		toSerialize["email"] = o.Email
	}
	if o.ContactProperties != nil {
		toSerialize["contactProperties"] = o.ContactProperties
	}
	return json.Marshal(toSerialize)
}

type NullableMarketingEventEmailSubscriber struct {
	value *MarketingEventEmailSubscriber
	isSet bool
}

func (v NullableMarketingEventEmailSubscriber) Get() *MarketingEventEmailSubscriber {
	return v.value
}

func (v *NullableMarketingEventEmailSubscriber) Set(val *MarketingEventEmailSubscriber) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketingEventEmailSubscriber) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketingEventEmailSubscriber) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketingEventEmailSubscriber(val *MarketingEventEmailSubscriber) *NullableMarketingEventEmailSubscriber {
	return &NullableMarketingEventEmailSubscriber{value: val, isSet: true}
}

func (v NullableMarketingEventEmailSubscriber) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketingEventEmailSubscriber) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}