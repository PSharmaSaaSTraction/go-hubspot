/*
HubDB endpoints

HubDB is a relational data store that presents data as rows, columns, and cells in a table, much like a spreadsheet. HubDB tables can be added or modified [in the HubSpot CMS](https://knowledge.hubspot.com/cos-general/how-to-edit-hubdb-tables), but you can also use the API endpoints documented here. For more information on HubDB tables and using their data on a HubSpot site, see the [CMS developers site](https://designers.hubspot.com/docs/tools/hubdb). You can also see the [documentation for dynamic pages](https://designers.hubspot.com/docs/tutorials/how-to-build-dynamic-pages-with-hubdb) for more details about the `useForPages` field.  HubDB tables support `draft` and `published` versions. This allows you to update data in the table, either for testing or to allow for a manual approval process, without affecting any live pages using the existing data. Draft data can be reviewed, and published by a user working in HubSpot or published via the API. Draft data can also be discarded, allowing users to go back to the published version of the data without disrupting it. If a table is set to be `allowed for public access`, you can access the published version of the table and rows without any authentication by specifying the portal id via the query parameter `portalId`.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hubdb

import (
	"encoding/json"
	"time"
)

// HubDbTableRowV3 struct for HubDbTableRowV3
type HubDbTableRowV3 struct {
	// The id of the table row
	Id *string `json:"id,omitempty"`
	// Specifies the value for `hs_path` column, which will be used as slug in the dynamic pages
	Path *string `json:"path,omitempty"`
	// Specifies the value for `hs_name` column, which will be used as title in the dynamic pages
	Name *string `json:"name,omitempty"`
	// Timestamp at which the row is created
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Timestamp at which the row is updated last time
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// Specifies the value for the column child table id
	ChildTableId *string `json:"childTableId,omitempty"`
	// List of key value pairs with the column name and column value
	Values map[string]map[string]interface{} `json:"values"`
}

// NewHubDbTableRowV3 instantiates a new HubDbTableRowV3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHubDbTableRowV3(values map[string]map[string]interface{}) *HubDbTableRowV3 {
	this := HubDbTableRowV3{}
	this.Values = values
	return &this
}

// NewHubDbTableRowV3WithDefaults instantiates a new HubDbTableRowV3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHubDbTableRowV3WithDefaults() *HubDbTableRowV3 {
	this := HubDbTableRowV3{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *HubDbTableRowV3) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HubDbTableRowV3) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *HubDbTableRowV3) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *HubDbTableRowV3) SetId(v string) {
	o.Id = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *HubDbTableRowV3) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HubDbTableRowV3) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *HubDbTableRowV3) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *HubDbTableRowV3) SetPath(v string) {
	o.Path = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HubDbTableRowV3) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HubDbTableRowV3) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HubDbTableRowV3) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HubDbTableRowV3) SetName(v string) {
	o.Name = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *HubDbTableRowV3) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HubDbTableRowV3) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *HubDbTableRowV3) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *HubDbTableRowV3) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *HubDbTableRowV3) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HubDbTableRowV3) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *HubDbTableRowV3) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *HubDbTableRowV3) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetChildTableId returns the ChildTableId field value if set, zero value otherwise.
func (o *HubDbTableRowV3) GetChildTableId() string {
	if o == nil || o.ChildTableId == nil {
		var ret string
		return ret
	}
	return *o.ChildTableId
}

// GetChildTableIdOk returns a tuple with the ChildTableId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HubDbTableRowV3) GetChildTableIdOk() (*string, bool) {
	if o == nil || o.ChildTableId == nil {
		return nil, false
	}
	return o.ChildTableId, true
}

// HasChildTableId returns a boolean if a field has been set.
func (o *HubDbTableRowV3) HasChildTableId() bool {
	if o != nil && o.ChildTableId != nil {
		return true
	}

	return false
}

// SetChildTableId gets a reference to the given string and assigns it to the ChildTableId field.
func (o *HubDbTableRowV3) SetChildTableId(v string) {
	o.ChildTableId = &v
}

// GetValues returns the Values field value
func (o *HubDbTableRowV3) GetValues() map[string]map[string]interface{} {
	if o == nil {
		var ret map[string]map[string]interface{}
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *HubDbTableRowV3) GetValuesOk() (*map[string]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Values, true
}

// SetValues sets field value
func (o *HubDbTableRowV3) SetValues(v map[string]map[string]interface{}) {
	o.Values = v
}

func (o HubDbTableRowV3) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.ChildTableId != nil {
		toSerialize["childTableId"] = o.ChildTableId
	}
	if true {
		toSerialize["values"] = o.Values
	}
	return json.Marshal(toSerialize)
}

type NullableHubDbTableRowV3 struct {
	value *HubDbTableRowV3
	isSet bool
}

func (v NullableHubDbTableRowV3) Get() *HubDbTableRowV3 {
	return v.value
}

func (v *NullableHubDbTableRowV3) Set(val *HubDbTableRowV3) {
	v.value = val
	v.isSet = true
}

func (v NullableHubDbTableRowV3) IsSet() bool {
	return v.isSet
}

func (v *NullableHubDbTableRowV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHubDbTableRowV3(val *HubDbTableRowV3) *NullableHubDbTableRowV3 {
	return &NullableHubDbTableRowV3{value: val, isSet: true}
}

func (v NullableHubDbTableRowV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHubDbTableRowV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}