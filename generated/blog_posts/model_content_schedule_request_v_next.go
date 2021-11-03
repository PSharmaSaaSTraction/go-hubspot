/*
Blog Post endpoints

\"Use these endpoints for interacting with Blog Posts, Blog Authors, and Blog Tags\"

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package blog_posts

import (
	"encoding/json"
	"time"
)

// ContentScheduleRequestVNext struct for ContentScheduleRequestVNext
type ContentScheduleRequestVNext struct {
	// The ID of the object to be scheduled.
	Id string `json:"id"`
	// The date the object should transition from scheduled to published.
	PublishDate time.Time `json:"publishDate"`
}

// NewContentScheduleRequestVNext instantiates a new ContentScheduleRequestVNext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentScheduleRequestVNext(id string, publishDate time.Time) *ContentScheduleRequestVNext {
	this := ContentScheduleRequestVNext{}
	this.Id = id
	this.PublishDate = publishDate
	return &this
}

// NewContentScheduleRequestVNextWithDefaults instantiates a new ContentScheduleRequestVNext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentScheduleRequestVNextWithDefaults() *ContentScheduleRequestVNext {
	this := ContentScheduleRequestVNext{}
	return &this
}

// GetId returns the Id field value
func (o *ContentScheduleRequestVNext) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ContentScheduleRequestVNext) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ContentScheduleRequestVNext) SetId(v string) {
	o.Id = v
}

// GetPublishDate returns the PublishDate field value
func (o *ContentScheduleRequestVNext) GetPublishDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.PublishDate
}

// GetPublishDateOk returns a tuple with the PublishDate field value
// and a boolean to check if the value has been set.
func (o *ContentScheduleRequestVNext) GetPublishDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublishDate, true
}

// SetPublishDate sets field value
func (o *ContentScheduleRequestVNext) SetPublishDate(v time.Time) {
	o.PublishDate = v
}

func (o ContentScheduleRequestVNext) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["publishDate"] = o.PublishDate
	}
	return json.Marshal(toSerialize)
}

type NullableContentScheduleRequestVNext struct {
	value *ContentScheduleRequestVNext
	isSet bool
}

func (v NullableContentScheduleRequestVNext) Get() *ContentScheduleRequestVNext {
	return v.value
}

func (v *NullableContentScheduleRequestVNext) Set(val *ContentScheduleRequestVNext) {
	v.value = val
	v.isSet = true
}

func (v NullableContentScheduleRequestVNext) IsSet() bool {
	return v.isSet
}

func (v *NullableContentScheduleRequestVNext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentScheduleRequestVNext(val *ContentScheduleRequestVNext) *NullableContentScheduleRequestVNext {
	return &NullableContentScheduleRequestVNext{value: val, isSet: true}
}

func (v NullableContentScheduleRequestVNext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentScheduleRequestVNext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}