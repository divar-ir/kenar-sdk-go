/*
Kenar API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kenarapi

import (
	"encoding/json"
	"time"
)

// checks if the PaymentTransaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentTransaction{}

// PaymentTransaction struct for PaymentTransaction
type PaymentTransaction struct {
	// هزینه تراکنش به ریال برای اپلیکیشن شما
	CostRials *string `json:"cost_rials,omitempty"`
	// The time when the transaction was created
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// همان جزئیات اضافی که در درخواست ارسال کردید
	ExtraDetails *string `json:"extra_details,omitempty"`
	// همان uuid هنگام ایجاد تراکنش
	Id *string `json:"id,omitempty"`
	State *PaymentTransactionState `json:"state,omitempty"`
	Type *PaymentTransactionType `json:"type,omitempty"`
}

// NewPaymentTransaction instantiates a new PaymentTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentTransaction() *PaymentTransaction {
	this := PaymentTransaction{}
	return &this
}

// NewPaymentTransactionWithDefaults instantiates a new PaymentTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentTransactionWithDefaults() *PaymentTransaction {
	this := PaymentTransaction{}
	return &this
}

// GetCostRials returns the CostRials field value if set, zero value otherwise.
func (o *PaymentTransaction) GetCostRials() string {
	if o == nil || IsNil(o.CostRials) {
		var ret string
		return ret
	}
	return *o.CostRials
}

// GetCostRialsOk returns a tuple with the CostRials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentTransaction) GetCostRialsOk() (*string, bool) {
	if o == nil || IsNil(o.CostRials) {
		return nil, false
	}
	return o.CostRials, true
}

// HasCostRials returns a boolean if a field has been set.
func (o *PaymentTransaction) HasCostRials() bool {
	if o != nil && !IsNil(o.CostRials) {
		return true
	}

	return false
}

// SetCostRials gets a reference to the given string and assigns it to the CostRials field.
func (o *PaymentTransaction) SetCostRials(v string) {
	o.CostRials = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PaymentTransaction) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentTransaction) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PaymentTransaction) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PaymentTransaction) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetExtraDetails returns the ExtraDetails field value if set, zero value otherwise.
func (o *PaymentTransaction) GetExtraDetails() string {
	if o == nil || IsNil(o.ExtraDetails) {
		var ret string
		return ret
	}
	return *o.ExtraDetails
}

// GetExtraDetailsOk returns a tuple with the ExtraDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentTransaction) GetExtraDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.ExtraDetails) {
		return nil, false
	}
	return o.ExtraDetails, true
}

// HasExtraDetails returns a boolean if a field has been set.
func (o *PaymentTransaction) HasExtraDetails() bool {
	if o != nil && !IsNil(o.ExtraDetails) {
		return true
	}

	return false
}

// SetExtraDetails gets a reference to the given string and assigns it to the ExtraDetails field.
func (o *PaymentTransaction) SetExtraDetails(v string) {
	o.ExtraDetails = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PaymentTransaction) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentTransaction) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PaymentTransaction) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PaymentTransaction) SetId(v string) {
	o.Id = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *PaymentTransaction) GetState() PaymentTransactionState {
	if o == nil || IsNil(o.State) {
		var ret PaymentTransactionState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentTransaction) GetStateOk() (*PaymentTransactionState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *PaymentTransaction) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given PaymentTransactionState and assigns it to the State field.
func (o *PaymentTransaction) SetState(v PaymentTransactionState) {
	o.State = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PaymentTransaction) GetType() PaymentTransactionType {
	if o == nil || IsNil(o.Type) {
		var ret PaymentTransactionType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentTransaction) GetTypeOk() (*PaymentTransactionType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PaymentTransaction) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given PaymentTransactionType and assigns it to the Type field.
func (o *PaymentTransaction) SetType(v PaymentTransactionType) {
	o.Type = &v
}

func (o PaymentTransaction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentTransaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CostRials) {
		toSerialize["cost_rials"] = o.CostRials
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.ExtraDetails) {
		toSerialize["extra_details"] = o.ExtraDetails
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullablePaymentTransaction struct {
	value *PaymentTransaction
	isSet bool
}

func (v NullablePaymentTransaction) Get() *PaymentTransaction {
	return v.value
}

func (v *NullablePaymentTransaction) Set(val *PaymentTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentTransaction(val *PaymentTransaction) *NullablePaymentTransaction {
	return &NullablePaymentTransaction{value: val, isSet: true}
}

func (v NullablePaymentTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


