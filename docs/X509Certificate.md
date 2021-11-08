# X509Certificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**SubjectX500Principal** | Pointer to [**X500Principal**](X500Principal.md) |  | [optional] 
**IssuerX500Principal** | Pointer to [**X500Principal**](X500Principal.md) |  | [optional] 
**SubjectAlternativeNames** | Pointer to **[][]map[string]interface{}** |  | [optional] 
**IssuerAlternativeNames** | Pointer to **[][]map[string]interface{}** |  | [optional] 
**KeyUsage** | Pointer to **[]bool** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 
**IssuerDN** | Pointer to [**Principal**](Principal.md) |  | [optional] 
**SubjectDN** | Pointer to [**Principal**](Principal.md) |  | [optional] 
**Signature** | Pointer to **[]string** |  | [optional] 
**BasicConstraints** | Pointer to **int32** |  | [optional] 
**SigAlgName** | Pointer to **string** |  | [optional] 
**SerialNumber** | Pointer to **int32** |  | [optional] 
**NotBefore** | Pointer to **time.Time** |  | [optional] 
**NotAfter** | Pointer to **time.Time** |  | [optional] 
**Tbscertificate** | Pointer to **[]string** |  | [optional] 
**SigAlgOID** | Pointer to **string** |  | [optional] 
**SigAlgParams** | Pointer to **[]string** |  | [optional] 
**IssuerUniqueID** | Pointer to **[]bool** |  | [optional] 
**SubjectUniqueID** | Pointer to **[]bool** |  | [optional] 
**ExtendedKeyUsage** | Pointer to **[]string** |  | [optional] 
**CriticalExtensionOIDs** | Pointer to **[]string** |  | [optional] 
**NonCriticalExtensionOIDs** | Pointer to **[]string** |  | [optional] 
**Encoded** | Pointer to **[]string** |  | [optional] 
**PublicKey** | Pointer to [**PublicKey**](PublicKey.md) |  | [optional] 

## Methods

### NewX509Certificate

`func NewX509Certificate() *X509Certificate`

NewX509Certificate instantiates a new X509Certificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewX509CertificateWithDefaults

`func NewX509CertificateWithDefaults() *X509Certificate`

NewX509CertificateWithDefaults instantiates a new X509Certificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *X509Certificate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *X509Certificate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *X509Certificate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *X509Certificate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSubjectX500Principal

`func (o *X509Certificate) GetSubjectX500Principal() X500Principal`

GetSubjectX500Principal returns the SubjectX500Principal field if non-nil, zero value otherwise.

### GetSubjectX500PrincipalOk

`func (o *X509Certificate) GetSubjectX500PrincipalOk() (*X500Principal, bool)`

GetSubjectX500PrincipalOk returns a tuple with the SubjectX500Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectX500Principal

`func (o *X509Certificate) SetSubjectX500Principal(v X500Principal)`

SetSubjectX500Principal sets SubjectX500Principal field to given value.

### HasSubjectX500Principal

`func (o *X509Certificate) HasSubjectX500Principal() bool`

HasSubjectX500Principal returns a boolean if a field has been set.

### GetIssuerX500Principal

`func (o *X509Certificate) GetIssuerX500Principal() X500Principal`

GetIssuerX500Principal returns the IssuerX500Principal field if non-nil, zero value otherwise.

### GetIssuerX500PrincipalOk

`func (o *X509Certificate) GetIssuerX500PrincipalOk() (*X500Principal, bool)`

GetIssuerX500PrincipalOk returns a tuple with the IssuerX500Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerX500Principal

`func (o *X509Certificate) SetIssuerX500Principal(v X500Principal)`

SetIssuerX500Principal sets IssuerX500Principal field to given value.

### HasIssuerX500Principal

`func (o *X509Certificate) HasIssuerX500Principal() bool`

HasIssuerX500Principal returns a boolean if a field has been set.

### GetSubjectAlternativeNames

`func (o *X509Certificate) GetSubjectAlternativeNames() [][]map[string]interface{}`

GetSubjectAlternativeNames returns the SubjectAlternativeNames field if non-nil, zero value otherwise.

### GetSubjectAlternativeNamesOk

`func (o *X509Certificate) GetSubjectAlternativeNamesOk() (*[][]map[string]interface{}, bool)`

GetSubjectAlternativeNamesOk returns a tuple with the SubjectAlternativeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAlternativeNames

`func (o *X509Certificate) SetSubjectAlternativeNames(v [][]map[string]interface{})`

SetSubjectAlternativeNames sets SubjectAlternativeNames field to given value.

### HasSubjectAlternativeNames

`func (o *X509Certificate) HasSubjectAlternativeNames() bool`

HasSubjectAlternativeNames returns a boolean if a field has been set.

### GetIssuerAlternativeNames

`func (o *X509Certificate) GetIssuerAlternativeNames() [][]map[string]interface{}`

GetIssuerAlternativeNames returns the IssuerAlternativeNames field if non-nil, zero value otherwise.

### GetIssuerAlternativeNamesOk

`func (o *X509Certificate) GetIssuerAlternativeNamesOk() (*[][]map[string]interface{}, bool)`

GetIssuerAlternativeNamesOk returns a tuple with the IssuerAlternativeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerAlternativeNames

`func (o *X509Certificate) SetIssuerAlternativeNames(v [][]map[string]interface{})`

SetIssuerAlternativeNames sets IssuerAlternativeNames field to given value.

### HasIssuerAlternativeNames

`func (o *X509Certificate) HasIssuerAlternativeNames() bool`

HasIssuerAlternativeNames returns a boolean if a field has been set.

### GetKeyUsage

`func (o *X509Certificate) GetKeyUsage() []bool`

GetKeyUsage returns the KeyUsage field if non-nil, zero value otherwise.

### GetKeyUsageOk

`func (o *X509Certificate) GetKeyUsageOk() (*[]bool, bool)`

GetKeyUsageOk returns a tuple with the KeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyUsage

`func (o *X509Certificate) SetKeyUsage(v []bool)`

SetKeyUsage sets KeyUsage field to given value.

### HasKeyUsage

`func (o *X509Certificate) HasKeyUsage() bool`

HasKeyUsage returns a boolean if a field has been set.

### GetVersion

`func (o *X509Certificate) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *X509Certificate) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *X509Certificate) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *X509Certificate) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetIssuerDN

`func (o *X509Certificate) GetIssuerDN() Principal`

GetIssuerDN returns the IssuerDN field if non-nil, zero value otherwise.

### GetIssuerDNOk

`func (o *X509Certificate) GetIssuerDNOk() (*Principal, bool)`

GetIssuerDNOk returns a tuple with the IssuerDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerDN

`func (o *X509Certificate) SetIssuerDN(v Principal)`

SetIssuerDN sets IssuerDN field to given value.

### HasIssuerDN

`func (o *X509Certificate) HasIssuerDN() bool`

HasIssuerDN returns a boolean if a field has been set.

### GetSubjectDN

`func (o *X509Certificate) GetSubjectDN() Principal`

GetSubjectDN returns the SubjectDN field if non-nil, zero value otherwise.

### GetSubjectDNOk

`func (o *X509Certificate) GetSubjectDNOk() (*Principal, bool)`

GetSubjectDNOk returns a tuple with the SubjectDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectDN

`func (o *X509Certificate) SetSubjectDN(v Principal)`

SetSubjectDN sets SubjectDN field to given value.

### HasSubjectDN

`func (o *X509Certificate) HasSubjectDN() bool`

HasSubjectDN returns a boolean if a field has been set.

### GetSignature

`func (o *X509Certificate) GetSignature() []string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *X509Certificate) GetSignatureOk() (*[]string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *X509Certificate) SetSignature(v []string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *X509Certificate) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetBasicConstraints

`func (o *X509Certificate) GetBasicConstraints() int32`

GetBasicConstraints returns the BasicConstraints field if non-nil, zero value otherwise.

### GetBasicConstraintsOk

`func (o *X509Certificate) GetBasicConstraintsOk() (*int32, bool)`

GetBasicConstraintsOk returns a tuple with the BasicConstraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicConstraints

`func (o *X509Certificate) SetBasicConstraints(v int32)`

SetBasicConstraints sets BasicConstraints field to given value.

### HasBasicConstraints

`func (o *X509Certificate) HasBasicConstraints() bool`

HasBasicConstraints returns a boolean if a field has been set.

### GetSigAlgName

`func (o *X509Certificate) GetSigAlgName() string`

GetSigAlgName returns the SigAlgName field if non-nil, zero value otherwise.

### GetSigAlgNameOk

`func (o *X509Certificate) GetSigAlgNameOk() (*string, bool)`

GetSigAlgNameOk returns a tuple with the SigAlgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigAlgName

`func (o *X509Certificate) SetSigAlgName(v string)`

SetSigAlgName sets SigAlgName field to given value.

### HasSigAlgName

`func (o *X509Certificate) HasSigAlgName() bool`

HasSigAlgName returns a boolean if a field has been set.

### GetSerialNumber

`func (o *X509Certificate) GetSerialNumber() int32`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *X509Certificate) GetSerialNumberOk() (*int32, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *X509Certificate) SetSerialNumber(v int32)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *X509Certificate) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetNotBefore

`func (o *X509Certificate) GetNotBefore() time.Time`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *X509Certificate) GetNotBeforeOk() (*time.Time, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *X509Certificate) SetNotBefore(v time.Time)`

SetNotBefore sets NotBefore field to given value.

### HasNotBefore

`func (o *X509Certificate) HasNotBefore() bool`

HasNotBefore returns a boolean if a field has been set.

### GetNotAfter

`func (o *X509Certificate) GetNotAfter() time.Time`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *X509Certificate) GetNotAfterOk() (*time.Time, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *X509Certificate) SetNotAfter(v time.Time)`

SetNotAfter sets NotAfter field to given value.

### HasNotAfter

`func (o *X509Certificate) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.

### GetTbscertificate

`func (o *X509Certificate) GetTbscertificate() []string`

GetTbscertificate returns the Tbscertificate field if non-nil, zero value otherwise.

### GetTbscertificateOk

`func (o *X509Certificate) GetTbscertificateOk() (*[]string, bool)`

GetTbscertificateOk returns a tuple with the Tbscertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTbscertificate

`func (o *X509Certificate) SetTbscertificate(v []string)`

SetTbscertificate sets Tbscertificate field to given value.

### HasTbscertificate

`func (o *X509Certificate) HasTbscertificate() bool`

HasTbscertificate returns a boolean if a field has been set.

### GetSigAlgOID

`func (o *X509Certificate) GetSigAlgOID() string`

GetSigAlgOID returns the SigAlgOID field if non-nil, zero value otherwise.

### GetSigAlgOIDOk

`func (o *X509Certificate) GetSigAlgOIDOk() (*string, bool)`

GetSigAlgOIDOk returns a tuple with the SigAlgOID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigAlgOID

`func (o *X509Certificate) SetSigAlgOID(v string)`

SetSigAlgOID sets SigAlgOID field to given value.

### HasSigAlgOID

`func (o *X509Certificate) HasSigAlgOID() bool`

HasSigAlgOID returns a boolean if a field has been set.

### GetSigAlgParams

`func (o *X509Certificate) GetSigAlgParams() []string`

GetSigAlgParams returns the SigAlgParams field if non-nil, zero value otherwise.

### GetSigAlgParamsOk

`func (o *X509Certificate) GetSigAlgParamsOk() (*[]string, bool)`

GetSigAlgParamsOk returns a tuple with the SigAlgParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigAlgParams

`func (o *X509Certificate) SetSigAlgParams(v []string)`

SetSigAlgParams sets SigAlgParams field to given value.

### HasSigAlgParams

`func (o *X509Certificate) HasSigAlgParams() bool`

HasSigAlgParams returns a boolean if a field has been set.

### GetIssuerUniqueID

`func (o *X509Certificate) GetIssuerUniqueID() []bool`

GetIssuerUniqueID returns the IssuerUniqueID field if non-nil, zero value otherwise.

### GetIssuerUniqueIDOk

`func (o *X509Certificate) GetIssuerUniqueIDOk() (*[]bool, bool)`

GetIssuerUniqueIDOk returns a tuple with the IssuerUniqueID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerUniqueID

`func (o *X509Certificate) SetIssuerUniqueID(v []bool)`

SetIssuerUniqueID sets IssuerUniqueID field to given value.

### HasIssuerUniqueID

`func (o *X509Certificate) HasIssuerUniqueID() bool`

HasIssuerUniqueID returns a boolean if a field has been set.

### GetSubjectUniqueID

`func (o *X509Certificate) GetSubjectUniqueID() []bool`

GetSubjectUniqueID returns the SubjectUniqueID field if non-nil, zero value otherwise.

### GetSubjectUniqueIDOk

`func (o *X509Certificate) GetSubjectUniqueIDOk() (*[]bool, bool)`

GetSubjectUniqueIDOk returns a tuple with the SubjectUniqueID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectUniqueID

`func (o *X509Certificate) SetSubjectUniqueID(v []bool)`

SetSubjectUniqueID sets SubjectUniqueID field to given value.

### HasSubjectUniqueID

`func (o *X509Certificate) HasSubjectUniqueID() bool`

HasSubjectUniqueID returns a boolean if a field has been set.

### GetExtendedKeyUsage

`func (o *X509Certificate) GetExtendedKeyUsage() []string`

GetExtendedKeyUsage returns the ExtendedKeyUsage field if non-nil, zero value otherwise.

### GetExtendedKeyUsageOk

`func (o *X509Certificate) GetExtendedKeyUsageOk() (*[]string, bool)`

GetExtendedKeyUsageOk returns a tuple with the ExtendedKeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedKeyUsage

`func (o *X509Certificate) SetExtendedKeyUsage(v []string)`

SetExtendedKeyUsage sets ExtendedKeyUsage field to given value.

### HasExtendedKeyUsage

`func (o *X509Certificate) HasExtendedKeyUsage() bool`

HasExtendedKeyUsage returns a boolean if a field has been set.

### GetCriticalExtensionOIDs

`func (o *X509Certificate) GetCriticalExtensionOIDs() []string`

GetCriticalExtensionOIDs returns the CriticalExtensionOIDs field if non-nil, zero value otherwise.

### GetCriticalExtensionOIDsOk

`func (o *X509Certificate) GetCriticalExtensionOIDsOk() (*[]string, bool)`

GetCriticalExtensionOIDsOk returns a tuple with the CriticalExtensionOIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalExtensionOIDs

`func (o *X509Certificate) SetCriticalExtensionOIDs(v []string)`

SetCriticalExtensionOIDs sets CriticalExtensionOIDs field to given value.

### HasCriticalExtensionOIDs

`func (o *X509Certificate) HasCriticalExtensionOIDs() bool`

HasCriticalExtensionOIDs returns a boolean if a field has been set.

### GetNonCriticalExtensionOIDs

`func (o *X509Certificate) GetNonCriticalExtensionOIDs() []string`

GetNonCriticalExtensionOIDs returns the NonCriticalExtensionOIDs field if non-nil, zero value otherwise.

### GetNonCriticalExtensionOIDsOk

`func (o *X509Certificate) GetNonCriticalExtensionOIDsOk() (*[]string, bool)`

GetNonCriticalExtensionOIDsOk returns a tuple with the NonCriticalExtensionOIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonCriticalExtensionOIDs

`func (o *X509Certificate) SetNonCriticalExtensionOIDs(v []string)`

SetNonCriticalExtensionOIDs sets NonCriticalExtensionOIDs field to given value.

### HasNonCriticalExtensionOIDs

`func (o *X509Certificate) HasNonCriticalExtensionOIDs() bool`

HasNonCriticalExtensionOIDs returns a boolean if a field has been set.

### GetEncoded

`func (o *X509Certificate) GetEncoded() []string`

GetEncoded returns the Encoded field if non-nil, zero value otherwise.

### GetEncodedOk

`func (o *X509Certificate) GetEncodedOk() (*[]string, bool)`

GetEncodedOk returns a tuple with the Encoded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoded

`func (o *X509Certificate) SetEncoded(v []string)`

SetEncoded sets Encoded field to given value.

### HasEncoded

`func (o *X509Certificate) HasEncoded() bool`

HasEncoded returns a boolean if a field has been set.

### GetPublicKey

`func (o *X509Certificate) GetPublicKey() PublicKey`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *X509Certificate) GetPublicKeyOk() (*PublicKey, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *X509Certificate) SetPublicKey(v PublicKey)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *X509Certificate) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


