# BookieInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rack** | Pointer to **string** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 

## Methods

### NewBookieInfo

`func NewBookieInfo() *BookieInfo`

NewBookieInfo instantiates a new BookieInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookieInfoWithDefaults

`func NewBookieInfoWithDefaults() *BookieInfo`

NewBookieInfoWithDefaults instantiates a new BookieInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRack

`func (o *BookieInfo) GetRack() string`

GetRack returns the Rack field if non-nil, zero value otherwise.

### GetRackOk

`func (o *BookieInfo) GetRackOk() (*string, bool)`

GetRackOk returns a tuple with the Rack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRack

`func (o *BookieInfo) SetRack(v string)`

SetRack sets Rack field to given value.

### HasRack

`func (o *BookieInfo) HasRack() bool`

HasRack returns a boolean if a field has been set.

### GetHostname

`func (o *BookieInfo) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *BookieInfo) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *BookieInfo) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *BookieInfo) HasHostname() bool`

HasHostname returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


