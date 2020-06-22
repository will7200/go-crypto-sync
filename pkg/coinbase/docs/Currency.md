# Currency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | 
**Name** | Pointer to **string** |  | 
**Color** | Pointer to **string** |  | 
**SortIndex** | Pointer to **float32** |  | 
**Exponent** | Pointer to **float32** |  | 
**Type** | Pointer to **string** |  | 
**AddressRegex** | Pointer to **string** |  | [optional] 
**AssetId** | Pointer to **string** |  | 
**Slug** | Pointer to **string** |  | 

## Methods

### NewCurrency

`func NewCurrency(code string, name string, color string, sortIndex float32, exponent float32, type_ string, assetId string, slug string, ) *Currency`

NewCurrency instantiates a new Currency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrencyWithDefaults

`func NewCurrencyWithDefaults() *Currency`

NewCurrencyWithDefaults instantiates a new Currency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *Currency) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Currency) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Currency) SetCode(v string)`

SetCode sets Code field to given value.


### GetName

`func (o *Currency) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Currency) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Currency) SetName(v string)`

SetName sets Name field to given value.


### GetColor

`func (o *Currency) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *Currency) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *Currency) SetColor(v string)`

SetColor sets Color field to given value.


### GetSortIndex

`func (o *Currency) GetSortIndex() float32`

GetSortIndex returns the SortIndex field if non-nil, zero value otherwise.

### GetSortIndexOk

`func (o *Currency) GetSortIndexOk() (*float32, bool)`

GetSortIndexOk returns a tuple with the SortIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortIndex

`func (o *Currency) SetSortIndex(v float32)`

SetSortIndex sets SortIndex field to given value.


### GetExponent

`func (o *Currency) GetExponent() float32`

GetExponent returns the Exponent field if non-nil, zero value otherwise.

### GetExponentOk

`func (o *Currency) GetExponentOk() (*float32, bool)`

GetExponentOk returns a tuple with the Exponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExponent

`func (o *Currency) SetExponent(v float32)`

SetExponent sets Exponent field to given value.


### GetType

`func (o *Currency) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Currency) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Currency) SetType(v string)`

SetType sets Type field to given value.


### GetAddressRegex

`func (o *Currency) GetAddressRegex() string`

GetAddressRegex returns the AddressRegex field if non-nil, zero value otherwise.

### GetAddressRegexOk

`func (o *Currency) GetAddressRegexOk() (*string, bool)`

GetAddressRegexOk returns a tuple with the AddressRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressRegex

`func (o *Currency) SetAddressRegex(v string)`

SetAddressRegex sets AddressRegex field to given value.

### HasAddressRegex

`func (o *Currency) HasAddressRegex() bool`

HasAddressRegex returns a boolean if a field has been set.

### GetAssetId

`func (o *Currency) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *Currency) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *Currency) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetSlug

`func (o *Currency) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *Currency) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *Currency) SetSlug(v string)`

SetSlug sets Slug field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


