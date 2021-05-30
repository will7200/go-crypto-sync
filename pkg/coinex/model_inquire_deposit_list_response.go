/*
 * CoinEx API
 *
 * Coinex doesn't have a openapi docs sadville  Open and simple, CoinEx API makes sure that you can build your own trading tools to achieve a more effective trading strategy. CoinEx API is now available for these features:
 *
 * API version: 2021-05-29
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package coinex

import (
	"encoding/json"
)

// InquireDepositListResponse InquireDepositList Response Value
type InquireDepositListResponse struct {
	// actual deposit amount
	ActualAmount *map[string]interface{} `json:"actual_amount,omitempty"`
	// display of actual deposit amount
	ActualAmountDisplay *map[string]interface{} `json:"actual_amount_display,omitempty"`
	// Depositor
	AddExplorer *map[string]interface{} `json:"add_explorer,omitempty"`
	// Amount
	Amount *map[string]interface{} `json:"amount,omitempty"`
	// Amount displayed
	AmountDisplay *map[string]interface{} `json:"amount_display,omitempty"`
	// Deposit add
	CoinAddress *map[string]interface{} `json:"coin_address,omitempty"`
	// Deposit add displayed
	CoinAddressDisplay *map[string]interface{} `json:"coin_address_display,omitempty"`
	// Deposit ID
	CoinDepositId *map[string]interface{} `json:"coin_deposit_id,omitempty"`
	// Coin type
	CoinType *map[string]interface{} `json:"coin_type,omitempty"`
	// confirmations
	Confirmations *map[string]interface{} `json:"confirmations,omitempty"`
	// create time
	CreateTime *map[string]interface{} `json:"create_time,omitempty"`
	// explorer
	Explorer *map[string]interface{} `json:"explorer,omitempty"`
	// remark
	Remark *map[string]interface{} `json:"remark,omitempty"`
	// processing，confirming，cancel，finish
	Status *map[string]interface{} `json:"status,omitempty"`
	// Status displayed
	StatusDisplay *map[string]interface{} `json:"status_display,omitempty"`
	// transfer method
	TransferMethod *map[string]interface{} `json:"transfer_method,omitempty"`
	// tx id
	TxId *map[string]interface{} `json:"tx_id,omitempty"`
	// tx id display
	TxIdDisplay *map[string]interface{} `json:"tx_id_display,omitempty"`
}

// NewInquireDepositListResponse instantiates a new InquireDepositListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInquireDepositListResponse() *InquireDepositListResponse {
	this := InquireDepositListResponse{}
	return &this
}

// NewInquireDepositListResponseWithDefaults instantiates a new InquireDepositListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInquireDepositListResponseWithDefaults() *InquireDepositListResponse {
	this := InquireDepositListResponse{}
	return &this
}

// GetActualAmount returns the ActualAmount field value if set, zero value otherwise.
func (o *InquireDepositListResponse) GetActualAmount() map[string]interface{} {
	if o == nil || o.ActualAmount == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.ActualAmount
}

// GetActualAmountOk returns a tuple with the ActualAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InquireDepositListResponse) GetActualAmountOk() (*map[string]interface{}, bool) {
	if o == nil || o.ActualAmount == nil {
		return nil, false
	}
	return o.ActualAmount, true
}

// HasActualAmount returns a boolean if a field has been set.
func (o *InquireDepositListResponse) HasActualAmount() bool {
	if o != nil && o.ActualAmount != nil {
		return true
	}

	return false
}

// SetActualAmount gets a reference to the given map[string]interface{} and assigns it to the ActualAmount field.
func (o *InquireDepositListResponse) SetActualAmount(v map[string]interface{}) {
	o.ActualAmount = &v
}

// GetActualAmountDisplay returns the ActualAmountDisplay field value if set, zero value otherwise.
func (o *InquireDepositListResponse) GetActualAmountDisplay() map[string]interface{} {
	if o == nil || o.ActualAmountDisplay == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.ActualAmountDisplay
}

// GetActualAmountDisplayOk returns a tuple with the ActualAmountDisplay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InquireDepositListResponse) GetActualAmountDisplayOk() (*map[string]interface{}, bool) {
	if o == nil || o.ActualAmountDisplay == nil {
		return nil, false
	}
	return o.ActualAmountDisplay, true
}

// HasActualAmountDisplay returns a boolean if a field has been set.
func (o *InquireDepositListResponse) HasActualAmountDisplay() bool {
	if o != nil && o.ActualAmountDisplay != nil {
		return true
	}

	return false
}

// SetActualAmountDisplay gets a reference to the given map[string]interface{} and assigns it to the ActualAmountDisplay field.
func (o *InquireDepositListResponse) SetActualAmountDisplay(v map[string]interface{}) {
	o.ActualAmountDisplay = &v
}

// GetAddExplorer returns the AddExplorer field value if set, zero value otherwise.
func (o *InquireDepositListResponse) GetAddExplorer() map[string]interface{} {
	if o == nil || o.AddExplorer == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.AddExplorer
}

// GetAddExplorerOk returns a tuple with the AddExplorer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InquireDepositListResponse) GetAddExplorerOk() (*map[string]interface{}, bool) {
	if o == nil || o.AddExplorer == nil {
		return nil, false
	}
	return o.AddExplorer, true
}

// HasAddExplorer returns a boolean if a field has been set.
func (o *InquireDepositListResponse) HasAddExplorer() bool {
	if o != nil && o.AddExplorer != nil {
		return true
	}

	return false
}

// SetAddExplorer gets a reference to the given map[string]interface{} and assigns it to the AddExplorer field.
func (o *InquireDepositListResponse) SetAddExplorer(v map[string]interface{}) {
	o.AddExplorer = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *InquireDepositListResponse) GetAmount() map[string]interface{} {
	if o == nil || o.Amount == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InquireDepositListResponse) GetAmountOk() (*map[string]interface{}, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *InquireDepositListResponse) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given map[string]interface{} and assigns it to the Amount field.
func (o *InquireDepositListResponse) SetAmount(v map[string]interface{}) {
	o.Amount = &v
}

// GetAmountDisplay returns the AmountDisplay field value if set, zero value otherwise.
func (o *InquireDepositListResponse) GetAmountDisplay() map[string]interface{} {
	if o == nil || o.AmountDisplay == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.AmountDisplay
}

// GetAmountDisplayOk returns a tuple with the AmountDisplay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InquireDepositListResponse) GetAmountDisplayOk() (*map[string]interface{}, bool) {
	if o == nil || o.AmountDisplay == nil {
		return nil, false
	}
	return o.AmountDisplay, true
}

// HasAmountDisplay returns a boolean if a field has been set.
func (o *InquireDepositListResponse) HasAmountDisplay() bool {
	if o != nil && o.AmountDisplay != nil {
		return true
	}

	return false
}

// SetAmountDisplay gets a reference to the given map[string]interface{} and assigns it to the AmountDisplay field.
func (o *InquireDepositListResponse) SetAmountDisplay(v map[string]interface{}) {
	o.AmountDisplay = &v
}

// GetCoinAddress returns the CoinAddress field value if set, zero value otherwise.
func (o *InquireDepositListResponse) GetCoinAddress() map[string]interface{} {
	if o == nil || o.CoinAddress == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.CoinAddress
}

// GetCoinAddressOk returns a tuple with the CoinAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InquireDepositListResponse) GetCoinAddressOk() (*map[string]interface{}, bool) {
	if o == nil || o.CoinAddress == nil {
		return nil, false
	}
	return o.CoinAddress, true
}

// HasCoinAddress returns a boolean if a field has been set.
func (o *InquireDepositListResponse) HasCoinAddress() bool {
	if o != nil && o.CoinAddress != nil {
		return true
	}

	return false
}

// SetCoinAddress gets a reference to the given map[string]interface{} and assigns it to the CoinAddress field.
func (o *InquireDepositListResponse) SetCoinAddress(v map[string]interface{}) {
	o.CoinAddress = &v
}

// GetCoinAddressDisplay returns the CoinAddressDisplay field value if set, zero value otherwise.
func (o *InquireDepositListResponse) GetCoinAddressDisplay() map[string]interface{} {
	if o == nil || o.CoinAddressDisplay == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.CoinAddressDisplay
}

// GetCoinAddressDisplayOk returns a tuple with the CoinAddressDisplay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InquireDepositListResponse) GetCoinAddressDisplayOk() (*map[string]interface{}, bool) {
	if o == nil || o.CoinAddressDisplay == nil {
		return nil, false
	}
	return o.CoinAddressDisplay, true
}

// HasCoinAddressDisplay returns a boolean if a field has been set.
func (o *InquireDepositListResponse) HasCoinAddressDisplay() bool {
	if o != nil && o.CoinAddressDisplay != nil {
		return true
	}

	return false
}

// SetCoinAddressDisplay gets a reference to the given map[string]interface{} and assigns it to the CoinAddressDisplay field.
func (o *InquireDepositListResponse) SetCoinAddressDisplay(v map[string]interface{}) {
	o.CoinAddressDisplay = &v
}

// GetCoinDepositId returns the CoinDepositId field value if set, zero value otherwise.
func (o *InquireDepositListResponse) GetCoinDepositId() map[string]interface{} {
	if o == nil || o.CoinDepositId == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.CoinDepositId
}

// GetCoinDepositIdOk returns a tuple with the CoinDepositId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InquireDepositListResponse) GetCoinDepositIdOk() (*map[string]interface{}, bool) {
	if o == nil || o.CoinDepositId == nil {
		return nil, false
	}
	return o.CoinDepositId, true
}

// HasCoinDepositId returns a boolean if a field has been set.
func (o *InquireDepositListResponse) HasCoinDepositId() bool {
	if o != nil && o.CoinDepositId != nil {
		return true
	}

	return false
}

// SetCoinDepositId gets a reference to the given map[string]interface{} and assigns it to the CoinDepositId field.
func (o *InquireDepositListResponse) SetCoinDepositId(v map[string]interface{}) {
	o.CoinDepositId = &v
}

// GetCoinType returns the CoinType field value if set, zero value otherwise.
func (o *InquireDepositListResponse) GetCoinType() map[string]interface{} {
	if o == nil || o.CoinType == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.CoinType
}

// GetCoinTypeOk returns a tuple with the CoinType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InquireDepositListResponse) GetCoinTypeOk() (*map[string]interface{}, bool) {
	if o == nil || o.CoinType == nil {
		return nil, false
	}
	return o.CoinType, true
}

// HasCoinType returns a boolean if a field has been set.
func (o *InquireDepositListResponse) HasCoinType() bool {
	if o != nil && o.CoinType != nil {
		return true
	}

	return false
}

// SetCoinType gets a reference to the given map[string]interface{} and assigns it to the CoinType field.
func (o *InquireDepositListResponse) SetCoinType(v map[string]interface{}) {
	o.CoinType = &v
}

// GetConfirmations returns the Confirmations field value if set, zero value otherwise.
func (o *InquireDepositListResponse) GetConfirmations() map[string]interface{} {
	if o == nil || o.Confirmations == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Confirmations
}

// GetConfirmationsOk returns a tuple with the Confirmations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InquireDepositListResponse) GetConfirmationsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Confirmations == nil {
		return nil, false
	}
	return o.Confirmations, true
}

// HasConfirmations returns a boolean if a field has been set.
func (o *InquireDepositListResponse) HasConfirmations() bool {
	if o != nil && o.Confirmations != nil {
		return true
	}

	return false
}

// SetConfirmations gets a reference to the given map[string]interface{} and assigns it to the Confirmations field.
func (o *InquireDepositListResponse) SetConfirmations(v map[string]interface{}) {
	o.Confirmations = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *InquireDepositListResponse) GetCreateTime() map[string]interface{} {
	if o == nil || o.CreateTime == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InquireDepositListResponse) GetCreateTimeOk() (*map[string]interface{}, bool) {
	if o == nil || o.CreateTime == nil {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *InquireDepositListResponse) HasCreateTime() bool {
	if o != nil && o.CreateTime != nil {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given map[string]interface{} and assigns it to the CreateTime field.
func (o *InquireDepositListResponse) SetCreateTime(v map[string]interface{}) {
	o.CreateTime = &v
}

// GetExplorer returns the Explorer field value if set, zero value otherwise.
func (o *InquireDepositListResponse) GetExplorer() map[string]interface{} {
	if o == nil || o.Explorer == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Explorer
}

// GetExplorerOk returns a tuple with the Explorer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InquireDepositListResponse) GetExplorerOk() (*map[string]interface{}, bool) {
	if o == nil || o.Explorer == nil {
		return nil, false
	}
	return o.Explorer, true
}

// HasExplorer returns a boolean if a field has been set.
func (o *InquireDepositListResponse) HasExplorer() bool {
	if o != nil && o.Explorer != nil {
		return true
	}

	return false
}

// SetExplorer gets a reference to the given map[string]interface{} and assigns it to the Explorer field.
func (o *InquireDepositListResponse) SetExplorer(v map[string]interface{}) {
	o.Explorer = &v
}

// GetRemark returns the Remark field value if set, zero value otherwise.
func (o *InquireDepositListResponse) GetRemark() map[string]interface{} {
	if o == nil || o.Remark == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Remark
}

// GetRemarkOk returns a tuple with the Remark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InquireDepositListResponse) GetRemarkOk() (*map[string]interface{}, bool) {
	if o == nil || o.Remark == nil {
		return nil, false
	}
	return o.Remark, true
}

// HasRemark returns a boolean if a field has been set.
func (o *InquireDepositListResponse) HasRemark() bool {
	if o != nil && o.Remark != nil {
		return true
	}

	return false
}

// SetRemark gets a reference to the given map[string]interface{} and assigns it to the Remark field.
func (o *InquireDepositListResponse) SetRemark(v map[string]interface{}) {
	o.Remark = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InquireDepositListResponse) GetStatus() map[string]interface{} {
	if o == nil || o.Status == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InquireDepositListResponse) GetStatusOk() (*map[string]interface{}, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InquireDepositListResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given map[string]interface{} and assigns it to the Status field.
func (o *InquireDepositListResponse) SetStatus(v map[string]interface{}) {
	o.Status = &v
}

// GetStatusDisplay returns the StatusDisplay field value if set, zero value otherwise.
func (o *InquireDepositListResponse) GetStatusDisplay() map[string]interface{} {
	if o == nil || o.StatusDisplay == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.StatusDisplay
}

// GetStatusDisplayOk returns a tuple with the StatusDisplay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InquireDepositListResponse) GetStatusDisplayOk() (*map[string]interface{}, bool) {
	if o == nil || o.StatusDisplay == nil {
		return nil, false
	}
	return o.StatusDisplay, true
}

// HasStatusDisplay returns a boolean if a field has been set.
func (o *InquireDepositListResponse) HasStatusDisplay() bool {
	if o != nil && o.StatusDisplay != nil {
		return true
	}

	return false
}

// SetStatusDisplay gets a reference to the given map[string]interface{} and assigns it to the StatusDisplay field.
func (o *InquireDepositListResponse) SetStatusDisplay(v map[string]interface{}) {
	o.StatusDisplay = &v
}

// GetTransferMethod returns the TransferMethod field value if set, zero value otherwise.
func (o *InquireDepositListResponse) GetTransferMethod() map[string]interface{} {
	if o == nil || o.TransferMethod == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.TransferMethod
}

// GetTransferMethodOk returns a tuple with the TransferMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InquireDepositListResponse) GetTransferMethodOk() (*map[string]interface{}, bool) {
	if o == nil || o.TransferMethod == nil {
		return nil, false
	}
	return o.TransferMethod, true
}

// HasTransferMethod returns a boolean if a field has been set.
func (o *InquireDepositListResponse) HasTransferMethod() bool {
	if o != nil && o.TransferMethod != nil {
		return true
	}

	return false
}

// SetTransferMethod gets a reference to the given map[string]interface{} and assigns it to the TransferMethod field.
func (o *InquireDepositListResponse) SetTransferMethod(v map[string]interface{}) {
	o.TransferMethod = &v
}

// GetTxId returns the TxId field value if set, zero value otherwise.
func (o *InquireDepositListResponse) GetTxId() map[string]interface{} {
	if o == nil || o.TxId == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InquireDepositListResponse) GetTxIdOk() (*map[string]interface{}, bool) {
	if o == nil || o.TxId == nil {
		return nil, false
	}
	return o.TxId, true
}

// HasTxId returns a boolean if a field has been set.
func (o *InquireDepositListResponse) HasTxId() bool {
	if o != nil && o.TxId != nil {
		return true
	}

	return false
}

// SetTxId gets a reference to the given map[string]interface{} and assigns it to the TxId field.
func (o *InquireDepositListResponse) SetTxId(v map[string]interface{}) {
	o.TxId = &v
}

// GetTxIdDisplay returns the TxIdDisplay field value if set, zero value otherwise.
func (o *InquireDepositListResponse) GetTxIdDisplay() map[string]interface{} {
	if o == nil || o.TxIdDisplay == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.TxIdDisplay
}

// GetTxIdDisplayOk returns a tuple with the TxIdDisplay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InquireDepositListResponse) GetTxIdDisplayOk() (*map[string]interface{}, bool) {
	if o == nil || o.TxIdDisplay == nil {
		return nil, false
	}
	return o.TxIdDisplay, true
}

// HasTxIdDisplay returns a boolean if a field has been set.
func (o *InquireDepositListResponse) HasTxIdDisplay() bool {
	if o != nil && o.TxIdDisplay != nil {
		return true
	}

	return false
}

// SetTxIdDisplay gets a reference to the given map[string]interface{} and assigns it to the TxIdDisplay field.
func (o *InquireDepositListResponse) SetTxIdDisplay(v map[string]interface{}) {
	o.TxIdDisplay = &v
}

func (o InquireDepositListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ActualAmount != nil {
		toSerialize["actual_amount"] = o.ActualAmount
	}
	if o.ActualAmountDisplay != nil {
		toSerialize["actual_amount_display"] = o.ActualAmountDisplay
	}
	if o.AddExplorer != nil {
		toSerialize["add_explorer"] = o.AddExplorer
	}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.AmountDisplay != nil {
		toSerialize["amount_display"] = o.AmountDisplay
	}
	if o.CoinAddress != nil {
		toSerialize["coin_address"] = o.CoinAddress
	}
	if o.CoinAddressDisplay != nil {
		toSerialize["coin_address_display"] = o.CoinAddressDisplay
	}
	if o.CoinDepositId != nil {
		toSerialize["coin_deposit_id"] = o.CoinDepositId
	}
	if o.CoinType != nil {
		toSerialize["coin_type"] = o.CoinType
	}
	if o.Confirmations != nil {
		toSerialize["confirmations"] = o.Confirmations
	}
	if o.CreateTime != nil {
		toSerialize["create_time"] = o.CreateTime
	}
	if o.Explorer != nil {
		toSerialize["explorer"] = o.Explorer
	}
	if o.Remark != nil {
		toSerialize["remark"] = o.Remark
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.StatusDisplay != nil {
		toSerialize["status_display"] = o.StatusDisplay
	}
	if o.TransferMethod != nil {
		toSerialize["transfer_method"] = o.TransferMethod
	}
	if o.TxId != nil {
		toSerialize["tx_id"] = o.TxId
	}
	if o.TxIdDisplay != nil {
		toSerialize["tx_id_display"] = o.TxIdDisplay
	}
	return json.Marshal(toSerialize)
}

type NullableInquireDepositListResponse struct {
	value *InquireDepositListResponse
	isSet bool
}

func (v NullableInquireDepositListResponse) Get() *InquireDepositListResponse {
	return v.value
}

func (v *NullableInquireDepositListResponse) Set(val *InquireDepositListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInquireDepositListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInquireDepositListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInquireDepositListResponse(val *InquireDepositListResponse) *NullableInquireDepositListResponse {
	return &NullableInquireDepositListResponse{value: val, isSet: true}
}

func (v NullableInquireDepositListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInquireDepositListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}