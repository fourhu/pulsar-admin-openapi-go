# JavaType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JavaLangObject** | Pointer to **bool** |  | [optional] 
**ArrayType** | Pointer to **bool** |  | [optional] 
**Throwable** | Pointer to **bool** |  | [optional] 
**CollectionLikeType** | Pointer to **bool** |  | [optional] 
**MapLikeType** | Pointer to **bool** |  | [optional] 
**ReferencedType** | Pointer to [**JavaType**](JavaType.md) |  | [optional] 
**ValueHandler** | Pointer to **map[string]interface{}** |  | [optional] 
**TypeHandler** | Pointer to **map[string]interface{}** |  | [optional] 
**ContentValueHandler** | Pointer to **map[string]interface{}** |  | [optional] 
**ContentTypeHandler** | Pointer to **map[string]interface{}** |  | [optional] 
**ErasedSignature** | Pointer to **string** |  | [optional] 
**ContainerType** | Pointer to **bool** |  | [optional] 
**Bindings** | Pointer to [**TypeBindings**](TypeBindings.md) |  | [optional] 
**KeyType** | Pointer to [**JavaType**](JavaType.md) |  | [optional] 
**Concrete** | Pointer to **bool** |  | [optional] 
**ContentType** | Pointer to [**JavaType**](JavaType.md) |  | [optional] 
**EnumType** | Pointer to **bool** |  | [optional] 
**Interface** | Pointer to **bool** |  | [optional] 
**Primitive** | Pointer to **bool** |  | [optional] 
**Interfaces** | Pointer to [**[]JavaType**](JavaType.md) |  | [optional] 
**GenericSignature** | Pointer to **string** |  | [optional] 
**Final** | Pointer to **bool** |  | [optional] 
**Abstract** | Pointer to **bool** |  | [optional] 
**SuperClass** | Pointer to [**JavaType**](JavaType.md) |  | [optional] 
**TypeName** | Pointer to **string** |  | [optional] 
**ReferenceType** | Pointer to **bool** |  | [optional] 

## Methods

### NewJavaType

`func NewJavaType() *JavaType`

NewJavaType instantiates a new JavaType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJavaTypeWithDefaults

`func NewJavaTypeWithDefaults() *JavaType`

NewJavaTypeWithDefaults instantiates a new JavaType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJavaLangObject

`func (o *JavaType) GetJavaLangObject() bool`

GetJavaLangObject returns the JavaLangObject field if non-nil, zero value otherwise.

### GetJavaLangObjectOk

`func (o *JavaType) GetJavaLangObjectOk() (*bool, bool)`

GetJavaLangObjectOk returns a tuple with the JavaLangObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJavaLangObject

`func (o *JavaType) SetJavaLangObject(v bool)`

SetJavaLangObject sets JavaLangObject field to given value.

### HasJavaLangObject

`func (o *JavaType) HasJavaLangObject() bool`

HasJavaLangObject returns a boolean if a field has been set.

### GetArrayType

`func (o *JavaType) GetArrayType() bool`

GetArrayType returns the ArrayType field if non-nil, zero value otherwise.

### GetArrayTypeOk

`func (o *JavaType) GetArrayTypeOk() (*bool, bool)`

GetArrayTypeOk returns a tuple with the ArrayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayType

`func (o *JavaType) SetArrayType(v bool)`

SetArrayType sets ArrayType field to given value.

### HasArrayType

`func (o *JavaType) HasArrayType() bool`

HasArrayType returns a boolean if a field has been set.

### GetThrowable

`func (o *JavaType) GetThrowable() bool`

GetThrowable returns the Throwable field if non-nil, zero value otherwise.

### GetThrowableOk

`func (o *JavaType) GetThrowableOk() (*bool, bool)`

GetThrowableOk returns a tuple with the Throwable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrowable

`func (o *JavaType) SetThrowable(v bool)`

SetThrowable sets Throwable field to given value.

### HasThrowable

`func (o *JavaType) HasThrowable() bool`

HasThrowable returns a boolean if a field has been set.

### GetCollectionLikeType

`func (o *JavaType) GetCollectionLikeType() bool`

GetCollectionLikeType returns the CollectionLikeType field if non-nil, zero value otherwise.

### GetCollectionLikeTypeOk

`func (o *JavaType) GetCollectionLikeTypeOk() (*bool, bool)`

GetCollectionLikeTypeOk returns a tuple with the CollectionLikeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionLikeType

`func (o *JavaType) SetCollectionLikeType(v bool)`

SetCollectionLikeType sets CollectionLikeType field to given value.

### HasCollectionLikeType

`func (o *JavaType) HasCollectionLikeType() bool`

HasCollectionLikeType returns a boolean if a field has been set.

### GetMapLikeType

`func (o *JavaType) GetMapLikeType() bool`

GetMapLikeType returns the MapLikeType field if non-nil, zero value otherwise.

### GetMapLikeTypeOk

`func (o *JavaType) GetMapLikeTypeOk() (*bool, bool)`

GetMapLikeTypeOk returns a tuple with the MapLikeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapLikeType

`func (o *JavaType) SetMapLikeType(v bool)`

SetMapLikeType sets MapLikeType field to given value.

### HasMapLikeType

`func (o *JavaType) HasMapLikeType() bool`

HasMapLikeType returns a boolean if a field has been set.

### GetReferencedType

`func (o *JavaType) GetReferencedType() JavaType`

GetReferencedType returns the ReferencedType field if non-nil, zero value otherwise.

### GetReferencedTypeOk

`func (o *JavaType) GetReferencedTypeOk() (*JavaType, bool)`

GetReferencedTypeOk returns a tuple with the ReferencedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedType

`func (o *JavaType) SetReferencedType(v JavaType)`

SetReferencedType sets ReferencedType field to given value.

### HasReferencedType

`func (o *JavaType) HasReferencedType() bool`

HasReferencedType returns a boolean if a field has been set.

### GetValueHandler

`func (o *JavaType) GetValueHandler() map[string]interface{}`

GetValueHandler returns the ValueHandler field if non-nil, zero value otherwise.

### GetValueHandlerOk

`func (o *JavaType) GetValueHandlerOk() (*map[string]interface{}, bool)`

GetValueHandlerOk returns a tuple with the ValueHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueHandler

`func (o *JavaType) SetValueHandler(v map[string]interface{})`

SetValueHandler sets ValueHandler field to given value.

### HasValueHandler

`func (o *JavaType) HasValueHandler() bool`

HasValueHandler returns a boolean if a field has been set.

### GetTypeHandler

`func (o *JavaType) GetTypeHandler() map[string]interface{}`

GetTypeHandler returns the TypeHandler field if non-nil, zero value otherwise.

### GetTypeHandlerOk

`func (o *JavaType) GetTypeHandlerOk() (*map[string]interface{}, bool)`

GetTypeHandlerOk returns a tuple with the TypeHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeHandler

`func (o *JavaType) SetTypeHandler(v map[string]interface{})`

SetTypeHandler sets TypeHandler field to given value.

### HasTypeHandler

`func (o *JavaType) HasTypeHandler() bool`

HasTypeHandler returns a boolean if a field has been set.

### GetContentValueHandler

`func (o *JavaType) GetContentValueHandler() map[string]interface{}`

GetContentValueHandler returns the ContentValueHandler field if non-nil, zero value otherwise.

### GetContentValueHandlerOk

`func (o *JavaType) GetContentValueHandlerOk() (*map[string]interface{}, bool)`

GetContentValueHandlerOk returns a tuple with the ContentValueHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentValueHandler

`func (o *JavaType) SetContentValueHandler(v map[string]interface{})`

SetContentValueHandler sets ContentValueHandler field to given value.

### HasContentValueHandler

`func (o *JavaType) HasContentValueHandler() bool`

HasContentValueHandler returns a boolean if a field has been set.

### GetContentTypeHandler

`func (o *JavaType) GetContentTypeHandler() map[string]interface{}`

GetContentTypeHandler returns the ContentTypeHandler field if non-nil, zero value otherwise.

### GetContentTypeHandlerOk

`func (o *JavaType) GetContentTypeHandlerOk() (*map[string]interface{}, bool)`

GetContentTypeHandlerOk returns a tuple with the ContentTypeHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentTypeHandler

`func (o *JavaType) SetContentTypeHandler(v map[string]interface{})`

SetContentTypeHandler sets ContentTypeHandler field to given value.

### HasContentTypeHandler

`func (o *JavaType) HasContentTypeHandler() bool`

HasContentTypeHandler returns a boolean if a field has been set.

### GetErasedSignature

`func (o *JavaType) GetErasedSignature() string`

GetErasedSignature returns the ErasedSignature field if non-nil, zero value otherwise.

### GetErasedSignatureOk

`func (o *JavaType) GetErasedSignatureOk() (*string, bool)`

GetErasedSignatureOk returns a tuple with the ErasedSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErasedSignature

`func (o *JavaType) SetErasedSignature(v string)`

SetErasedSignature sets ErasedSignature field to given value.

### HasErasedSignature

`func (o *JavaType) HasErasedSignature() bool`

HasErasedSignature returns a boolean if a field has been set.

### GetContainerType

`func (o *JavaType) GetContainerType() bool`

GetContainerType returns the ContainerType field if non-nil, zero value otherwise.

### GetContainerTypeOk

`func (o *JavaType) GetContainerTypeOk() (*bool, bool)`

GetContainerTypeOk returns a tuple with the ContainerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerType

`func (o *JavaType) SetContainerType(v bool)`

SetContainerType sets ContainerType field to given value.

### HasContainerType

`func (o *JavaType) HasContainerType() bool`

HasContainerType returns a boolean if a field has been set.

### GetBindings

`func (o *JavaType) GetBindings() TypeBindings`

GetBindings returns the Bindings field if non-nil, zero value otherwise.

### GetBindingsOk

`func (o *JavaType) GetBindingsOk() (*TypeBindings, bool)`

GetBindingsOk returns a tuple with the Bindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindings

`func (o *JavaType) SetBindings(v TypeBindings)`

SetBindings sets Bindings field to given value.

### HasBindings

`func (o *JavaType) HasBindings() bool`

HasBindings returns a boolean if a field has been set.

### GetKeyType

`func (o *JavaType) GetKeyType() JavaType`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *JavaType) GetKeyTypeOk() (*JavaType, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *JavaType) SetKeyType(v JavaType)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *JavaType) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### GetConcrete

`func (o *JavaType) GetConcrete() bool`

GetConcrete returns the Concrete field if non-nil, zero value otherwise.

### GetConcreteOk

`func (o *JavaType) GetConcreteOk() (*bool, bool)`

GetConcreteOk returns a tuple with the Concrete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcrete

`func (o *JavaType) SetConcrete(v bool)`

SetConcrete sets Concrete field to given value.

### HasConcrete

`func (o *JavaType) HasConcrete() bool`

HasConcrete returns a boolean if a field has been set.

### GetContentType

`func (o *JavaType) GetContentType() JavaType`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *JavaType) GetContentTypeOk() (*JavaType, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *JavaType) SetContentType(v JavaType)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *JavaType) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetEnumType

`func (o *JavaType) GetEnumType() bool`

GetEnumType returns the EnumType field if non-nil, zero value otherwise.

### GetEnumTypeOk

`func (o *JavaType) GetEnumTypeOk() (*bool, bool)`

GetEnumTypeOk returns a tuple with the EnumType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnumType

`func (o *JavaType) SetEnumType(v bool)`

SetEnumType sets EnumType field to given value.

### HasEnumType

`func (o *JavaType) HasEnumType() bool`

HasEnumType returns a boolean if a field has been set.

### GetInterface

`func (o *JavaType) GetInterface() bool`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *JavaType) GetInterfaceOk() (*bool, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *JavaType) SetInterface(v bool)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *JavaType) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetPrimitive

`func (o *JavaType) GetPrimitive() bool`

GetPrimitive returns the Primitive field if non-nil, zero value otherwise.

### GetPrimitiveOk

`func (o *JavaType) GetPrimitiveOk() (*bool, bool)`

GetPrimitiveOk returns a tuple with the Primitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimitive

`func (o *JavaType) SetPrimitive(v bool)`

SetPrimitive sets Primitive field to given value.

### HasPrimitive

`func (o *JavaType) HasPrimitive() bool`

HasPrimitive returns a boolean if a field has been set.

### GetInterfaces

`func (o *JavaType) GetInterfaces() []JavaType`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *JavaType) GetInterfacesOk() (*[]JavaType, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *JavaType) SetInterfaces(v []JavaType)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *JavaType) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetGenericSignature

`func (o *JavaType) GetGenericSignature() string`

GetGenericSignature returns the GenericSignature field if non-nil, zero value otherwise.

### GetGenericSignatureOk

`func (o *JavaType) GetGenericSignatureOk() (*string, bool)`

GetGenericSignatureOk returns a tuple with the GenericSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericSignature

`func (o *JavaType) SetGenericSignature(v string)`

SetGenericSignature sets GenericSignature field to given value.

### HasGenericSignature

`func (o *JavaType) HasGenericSignature() bool`

HasGenericSignature returns a boolean if a field has been set.

### GetFinal

`func (o *JavaType) GetFinal() bool`

GetFinal returns the Final field if non-nil, zero value otherwise.

### GetFinalOk

`func (o *JavaType) GetFinalOk() (*bool, bool)`

GetFinalOk returns a tuple with the Final field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinal

`func (o *JavaType) SetFinal(v bool)`

SetFinal sets Final field to given value.

### HasFinal

`func (o *JavaType) HasFinal() bool`

HasFinal returns a boolean if a field has been set.

### GetAbstract

`func (o *JavaType) GetAbstract() bool`

GetAbstract returns the Abstract field if non-nil, zero value otherwise.

### GetAbstractOk

`func (o *JavaType) GetAbstractOk() (*bool, bool)`

GetAbstractOk returns a tuple with the Abstract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbstract

`func (o *JavaType) SetAbstract(v bool)`

SetAbstract sets Abstract field to given value.

### HasAbstract

`func (o *JavaType) HasAbstract() bool`

HasAbstract returns a boolean if a field has been set.

### GetSuperClass

`func (o *JavaType) GetSuperClass() JavaType`

GetSuperClass returns the SuperClass field if non-nil, zero value otherwise.

### GetSuperClassOk

`func (o *JavaType) GetSuperClassOk() (*JavaType, bool)`

GetSuperClassOk returns a tuple with the SuperClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuperClass

`func (o *JavaType) SetSuperClass(v JavaType)`

SetSuperClass sets SuperClass field to given value.

### HasSuperClass

`func (o *JavaType) HasSuperClass() bool`

HasSuperClass returns a boolean if a field has been set.

### GetTypeName

`func (o *JavaType) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *JavaType) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *JavaType) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *JavaType) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.

### GetReferenceType

`func (o *JavaType) GetReferenceType() bool`

GetReferenceType returns the ReferenceType field if non-nil, zero value otherwise.

### GetReferenceTypeOk

`func (o *JavaType) GetReferenceTypeOk() (*bool, bool)`

GetReferenceTypeOk returns a tuple with the ReferenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceType

`func (o *JavaType) SetReferenceType(v bool)`

SetReferenceType sets ReferenceType field to given value.

### HasReferenceType

`func (o *JavaType) HasReferenceType() bool`

HasReferenceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


