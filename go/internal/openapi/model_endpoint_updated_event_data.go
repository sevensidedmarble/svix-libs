/*
 * Svix API
 *
 * Welcome to the Svix API documentation!  Useful links: [Homepage](https://www.svix.com) | [Support email](mailto:support+docs@svix.com) | [Blog](https://www.svix.com/blog/) | [Slack Community](https://www.svix.com/slack/)  # Introduction  This is the reference documentation and schemas for the [Svix webhook service](https://www.svix.com) API. For tutorials and other documentation please refer to [the documentation](https://docs.svix.com).  ## Main concepts  In Svix you have four important entities you will be interacting with:  - `messages`: these are the webhooks being sent. They can have contents and a few other properties. - `application`: this is where `messages` are sent to. Usually you want to create one application for each of your users. - `endpoint`: endpoints are the URLs messages will be sent to. Each application can have multiple `endpoints` and each message sent to that application will be sent to all of them (unless they are not subscribed to the sent event type). - `event-type`: event types are identifiers denoting the type of the message being sent. Event types are primarily used to decide which events are sent to which endpoint.   ## Authentication  Get your authentication token (`AUTH_TOKEN`) from the [Svix dashboard](https://dashboard.svix.com) and use it as part of the `Authorization` header as such: `Authorization: Bearer ${AUTH_TOKEN}`.  <SecurityDefinitions />   ## Code samples  The code samples assume you already have the respective libraries installed and you know how to use them. For the latest information on how to do that, please refer to [the documentation](https://docs.svix.com/).   ## Cross-Origin Resource Sharing  This API features Cross-Origin Resource Sharing (CORS) implemented in compliance with [W3C spec](https://www.w3.org/TR/cors/). And that allows cross-domain communication from the browser. All responses have a wildcard same-origin which makes them completely public and accessible to everyone, including any code on any site. 
 *
 * API version: 1.4
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// EndpointUpdatedEventData struct for EndpointUpdatedEventData
type EndpointUpdatedEventData struct {
	AppId string `json:"appId"`
	// Optional unique identifier for the application
	AppUid *string `json:"appUid,omitempty"`
	EndpointId string `json:"endpointId"`
}

// NewEndpointUpdatedEventData instantiates a new EndpointUpdatedEventData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointUpdatedEventData(appId string, endpointId string) *EndpointUpdatedEventData {
	this := EndpointUpdatedEventData{}
	this.AppId = appId
	this.EndpointId = endpointId
	return &this
}

// NewEndpointUpdatedEventDataWithDefaults instantiates a new EndpointUpdatedEventData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointUpdatedEventDataWithDefaults() *EndpointUpdatedEventData {
	this := EndpointUpdatedEventData{}
	return &this
}

// GetAppId returns the AppId field value
func (o *EndpointUpdatedEventData) GetAppId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value
// and a boolean to check if the value has been set.
func (o *EndpointUpdatedEventData) GetAppIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AppId, true
}

// SetAppId sets field value
func (o *EndpointUpdatedEventData) SetAppId(v string) {
	o.AppId = v
}

// GetAppUid returns the AppUid field value if set, zero value otherwise.
func (o *EndpointUpdatedEventData) GetAppUid() string {
	if o == nil || o.AppUid == nil {
		var ret string
		return ret
	}
	return *o.AppUid
}

// GetAppUidOk returns a tuple with the AppUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointUpdatedEventData) GetAppUidOk() (*string, bool) {
	if o == nil || o.AppUid == nil {
		return nil, false
	}
	return o.AppUid, true
}

// HasAppUid returns a boolean if a field has been set.
func (o *EndpointUpdatedEventData) HasAppUid() bool {
	if o != nil && o.AppUid != nil {
		return true
	}

	return false
}

// SetAppUid gets a reference to the given string and assigns it to the AppUid field.
func (o *EndpointUpdatedEventData) SetAppUid(v string) {
	o.AppUid = &v
}

// GetEndpointId returns the EndpointId field value
func (o *EndpointUpdatedEventData) GetEndpointId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndpointId
}

// GetEndpointIdOk returns a tuple with the EndpointId field value
// and a boolean to check if the value has been set.
func (o *EndpointUpdatedEventData) GetEndpointIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EndpointId, true
}

// SetEndpointId sets field value
func (o *EndpointUpdatedEventData) SetEndpointId(v string) {
	o.EndpointId = v
}

func (o EndpointUpdatedEventData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["appId"] = o.AppId
	}
	if o.AppUid != nil {
		toSerialize["appUid"] = o.AppUid
	}
	if true {
		toSerialize["endpointId"] = o.EndpointId
	}
	return json.Marshal(toSerialize)
}

type NullableEndpointUpdatedEventData struct {
	value *EndpointUpdatedEventData
	isSet bool
}

func (v NullableEndpointUpdatedEventData) Get() *EndpointUpdatedEventData {
	return v.value
}

func (v *NullableEndpointUpdatedEventData) Set(val *EndpointUpdatedEventData) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointUpdatedEventData) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointUpdatedEventData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointUpdatedEventData(val *EndpointUpdatedEventData) *NullableEndpointUpdatedEventData {
	return &NullableEndpointUpdatedEventData{value: val, isSet: true}
}

func (v NullableEndpointUpdatedEventData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointUpdatedEventData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


