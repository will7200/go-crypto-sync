/*
 * Nomics Cryptocurrency & Bitcoin API
 *
 * # Introduction  Welcome to the Nomics Cryptocurrency & Bitcoin API. To sign up for an API key please [go here](https://p.nomics.com/cryptocurrency-bitcoin-api/).  [nomics.com](https://nomics.com) is built entirely with the Nomics API. Everything we've done on [nomics.com](https://nomics.com) you can do with our API. There are no internal API endpoints.  If you need support, reach out to use at our [forums](https://forums.nomics.com/).  # General  ## API Server URL  The Nomics API runs at `https://api.nomics.com/v1`. All requests should be prefixed by the server URL.  ## JSON and CSV Support  By default, all endpoints serve data as JSON. However, by passing `format=csv` in the URL, some endpoints will return CSV data. This can be used in Google Sheets via the `IMPORTDATA` function.  CSV responses will not contain a header row, this is so that data can be easily concatenated from multiple requests. The fields will be rendered in the same order as the JSON fields. See the endpoint's documentation for an example.  Not all endpoints support CSV. Endpoints that support CSV will have the `format` parameter in the parameters section.  ## Errors  The Nomics API uses standard HTTP status codes to indicate success or failure. 200 represents success, 4xx represents a user error (such as a problem with your key), and 5xx represents a problem with our API.  ## Versioning  We follow Semantic Versioning. That means our API is versioned as Major.Minor.Patch. For example, Version 1.2.3 has major version 1, minor version 2, and patch version 3.  Major version changes indicate that we have altered the API significantly and it is no longer compatible with a previous version. Major versions are also used as the API URL prefix.  When we update the major version, we will not remove the previous version without notice to API customers and a deprecation period to allow everyone to smoothly update to the new version.  Minor version changes indicate that we have added new functionality without breaking any existing functionality. An API client is compatible with future minor versions. Note that a minor version update may add a new field to an existing API endpoint's response. Your API client must ignore fields it does not understand in order to be compatible with future minor versions.  Patch version changes indicate we fixed a bug or security vulnerability. Patch versions don't add new functionality.  ## Cross Origin Resource Sharing (CORS)  This API supports Cross Origin Resource Sharing, which allows you to make API requests directly from your user's browser.  To use CORS, you must provide Nomics with the domains on which your application will run so that we can whitelist them for CORS access.  Requests from `localhost`, `127.0.0.1`, and `0.0.0.0` will always succeed to aid in development.  ## Demo Application  A demo application using the Nomics API, CORS, and React is available on Glitch.com. This can help you get started using the Nomics API. Keep in mind it uses the demo key, which is rotated frequently. You should get your own API key before deploying an app to production. Check it out:  <div class=\"glitch-embed-wrap\" style=\"height: 420px; width: 100%;\">   <iframe src=\"https://glitch.com/embed/#!/embed/nomics-api-demo?path=README.md\" alt=\"nomics-api-demo on glitch\" style=\"height: 100%; width: 100%; border: 0;\"></iframe> </div>  ## Demo Spreadsheet  Here is a demo of using the Nomics API with Google Sheets.  <iframe width=\"100%\" height=\"400px\" src=\"https://docs.google.com/spreadsheets/d/e/2PACX-1vShn2iWjvqQ0ueBa9l9g1UBYVM92OZSgZ4nmp0rWuykvHPrvyMyMeSN4r0Orj0ACEIIKdCz6cc5abCw/pubhtml?widget=true&amp;headers=false\"></iframe>  ### Formulas  * A2: `=IMPORTDATA(\"https://api.nomics.com/v1/prices?key=your-key-here&format=csv\")` * Column F: `=LOOKUP(D2,A:A,B:B)` finds D2 (BTC) in column A and pulls the price from column B * Column G: `=E2*F2` * Column H: `=G2/I$2` * Column I: `=SUM(G:G)`  # SDKs and Libraries  ## By Nomics - [Nomics JavaScript Client](https://github.com/nomics-crypto/nomics-javascript)  ## Community Submissions - [Nomics.com Swift SDK](https://forums.nomics.com/t/swift-sdk-supporting-ios-macos-tvos-and-watchos/) by Nick DiZazzo - [Nomics Node.js Library](https://forums.nomics.com/t/i-made-a-library-for-node-js/) by mikunimaru - [Nomics Python Wrapper](https://forums.nomics.com/t/python-package-for-nomics-api/119) by Taylor Facen - [Python Wrapper for Nomics](https://github.com/AviFelman/py-nomics) by Avi Felman  We love watching developers explore new use-cases with our API. Whether you're tinkering on a small side project or building an open-source resource, please share what you're up to in our [forums](https://forums.nomics.com/).  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package nomics

import (
	"encoding/json"
	"time"
)

// PtrBool is a helper routine that returns a pointer to given integer value.
func PtrBool(v bool) *bool { return &v }

// PtrInt is a helper routine that returns a pointer to given integer value.
func PtrInt(v int) *int { return &v }

// PtrInt32 is a helper routine that returns a pointer to given integer value.
func PtrInt32(v int32) *int32 { return &v }

// PtrInt64 is a helper routine that returns a pointer to given integer value.
func PtrInt64(v int64) *int64 { return &v }

// PtrFloat32 is a helper routine that returns a pointer to given float value.
func PtrFloat32(v float32) *float32 { return &v }

// PtrFloat64 is a helper routine that returns a pointer to given float value.
func PtrFloat64(v float64) *float64 { return &v }

// PtrString is a helper routine that returns a pointer to given string value.
func PtrString(v string) *string { return &v }

// PtrTime is helper routine that returns a pointer to given Time value.
func PtrTime(v time.Time) *time.Time { return &v }

type NullableBool struct {
	value *bool
	isSet bool
}

func (v NullableBool) Get() *bool {
	return v.value
}

func (v *NullableBool) Set(val *bool) {
	v.value = val
	v.isSet = true
}

func (v NullableBool) IsSet() bool {
	return v.isSet
}

func (v *NullableBool) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBool(val *bool) *NullableBool {
	return &NullableBool{value: val, isSet: true}
}

func (v NullableBool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type NullableInt struct {
	value *int
	isSet bool
}

func (v NullableInt) Get() *int {
	return v.value
}

func (v *NullableInt) Set(val *int) {
	v.value = val
	v.isSet = true
}

func (v NullableInt) IsSet() bool {
	return v.isSet
}

func (v *NullableInt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInt(val *int) *NullableInt {
	return &NullableInt{value: val, isSet: true}
}

func (v NullableInt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type NullableInt32 struct {
	value *int32
	isSet bool
}

func (v NullableInt32) Get() *int32 {
	return v.value
}

func (v *NullableInt32) Set(val *int32) {
	v.value = val
	v.isSet = true
}

func (v NullableInt32) IsSet() bool {
	return v.isSet
}

func (v *NullableInt32) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInt32(val *int32) *NullableInt32 {
	return &NullableInt32{value: val, isSet: true}
}

func (v NullableInt32) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInt32) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type NullableInt64 struct {
	value *int64
	isSet bool
}

func (v NullableInt64) Get() *int64 {
	return v.value
}

func (v *NullableInt64) Set(val *int64) {
	v.value = val
	v.isSet = true
}

func (v NullableInt64) IsSet() bool {
	return v.isSet
}

func (v *NullableInt64) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInt64(val *int64) *NullableInt64 {
	return &NullableInt64{value: val, isSet: true}
}

func (v NullableInt64) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInt64) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type NullableFloat32 struct {
	value *float32
	isSet bool
}

func (v NullableFloat32) Get() *float32 {
	return v.value
}

func (v *NullableFloat32) Set(val *float32) {
	v.value = val
	v.isSet = true
}

func (v NullableFloat32) IsSet() bool {
	return v.isSet
}

func (v *NullableFloat32) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFloat32(val *float32) *NullableFloat32 {
	return &NullableFloat32{value: val, isSet: true}
}

func (v NullableFloat32) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFloat32) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type NullableFloat64 struct {
	value *float64
	isSet bool
}

func (v NullableFloat64) Get() *float64 {
	return v.value
}

func (v *NullableFloat64) Set(val *float64) {
	v.value = val
	v.isSet = true
}

func (v NullableFloat64) IsSet() bool {
	return v.isSet
}

func (v *NullableFloat64) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFloat64(val *float64) *NullableFloat64 {
	return &NullableFloat64{value: val, isSet: true}
}

func (v NullableFloat64) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFloat64) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type NullableString struct {
	value *string
	isSet bool
}

func (v NullableString) Get() *string {
	return v.value
}

func (v *NullableString) Set(val *string) {
	v.value = val
	v.isSet = true
}

func (v NullableString) IsSet() bool {
	return v.isSet
}

func (v *NullableString) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableString(val *string) *NullableString {
	return &NullableString{value: val, isSet: true}
}

func (v NullableString) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableString) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type NullableTime struct {
	value *time.Time
	isSet bool
}

func (v NullableTime) Get() *time.Time {
	return v.value
}

func (v *NullableTime) Set(val *time.Time) {
	v.value = val
	v.isSet = true
}

func (v NullableTime) IsSet() bool {
	return v.isSet
}

func (v *NullableTime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTime(val *time.Time) *NullableTime {
	return &NullableTime{value: val, isSet: true}
}

func (v NullableTime) MarshalJSON() ([]byte, error) {
	return v.value.MarshalJSON()
}

func (v *NullableTime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
