# JsonSerializerObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnwrappingSerializer** | Pointer to **bool** |  | [optional] 
**Delegatee** | Pointer to [**JsonSerializerObject**](JsonSerializerObject.md) |  | [optional] 

## Methods

### NewJsonSerializerObject

`func NewJsonSerializerObject() *JsonSerializerObject`

NewJsonSerializerObject instantiates a new JsonSerializerObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonSerializerObjectWithDefaults

`func NewJsonSerializerObjectWithDefaults() *JsonSerializerObject`

NewJsonSerializerObjectWithDefaults instantiates a new JsonSerializerObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnwrappingSerializer

`func (o *JsonSerializerObject) GetUnwrappingSerializer() bool`

GetUnwrappingSerializer returns the UnwrappingSerializer field if non-nil, zero value otherwise.

### GetUnwrappingSerializerOk

`func (o *JsonSerializerObject) GetUnwrappingSerializerOk() (*bool, bool)`

GetUnwrappingSerializerOk returns a tuple with the UnwrappingSerializer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnwrappingSerializer

`func (o *JsonSerializerObject) SetUnwrappingSerializer(v bool)`

SetUnwrappingSerializer sets UnwrappingSerializer field to given value.

### HasUnwrappingSerializer

`func (o *JsonSerializerObject) HasUnwrappingSerializer() bool`

HasUnwrappingSerializer returns a boolean if a field has been set.

### GetDelegatee

`func (o *JsonSerializerObject) GetDelegatee() JsonSerializerObject`

GetDelegatee returns the Delegatee field if non-nil, zero value otherwise.

### GetDelegateeOk

`func (o *JsonSerializerObject) GetDelegateeOk() (*JsonSerializerObject, bool)`

GetDelegateeOk returns a tuple with the Delegatee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatee

`func (o *JsonSerializerObject) SetDelegatee(v JsonSerializerObject)`

SetDelegatee sets Delegatee field to given value.

### HasDelegatee

`func (o *JsonSerializerObject) HasDelegatee() bool`

HasDelegatee returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


