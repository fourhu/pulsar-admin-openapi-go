/*
Pulsar Admin REST API

This provides the REST API for admin operations

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// X509Certificate struct for X509Certificate
type X509Certificate struct {
	Type *string `json:"type,omitempty"`
	SubjectX500Principal *X500Principal `json:"subjectX500Principal,omitempty"`
	IssuerX500Principal *X500Principal `json:"issuerX500Principal,omitempty"`
	SubjectAlternativeNames *[][]map[string]interface{} `json:"subjectAlternativeNames,omitempty"`
	IssuerAlternativeNames *[][]map[string]interface{} `json:"issuerAlternativeNames,omitempty"`
	KeyUsage *[]bool `json:"keyUsage,omitempty"`
	Version *int32 `json:"version,omitempty"`
	IssuerDN *Principal `json:"issuerDN,omitempty"`
	SubjectDN *Principal `json:"subjectDN,omitempty"`
	Signature *[]string `json:"signature,omitempty"`
	BasicConstraints *int32 `json:"basicConstraints,omitempty"`
	SigAlgName *string `json:"sigAlgName,omitempty"`
	SerialNumber *int32 `json:"serialNumber,omitempty"`
	NotBefore *time.Time `json:"notBefore,omitempty"`
	NotAfter *time.Time `json:"notAfter,omitempty"`
	Tbscertificate *[]string `json:"tbscertificate,omitempty"`
	SigAlgOID *string `json:"sigAlgOID,omitempty"`
	SigAlgParams *[]string `json:"sigAlgParams,omitempty"`
	IssuerUniqueID *[]bool `json:"issuerUniqueID,omitempty"`
	SubjectUniqueID *[]bool `json:"subjectUniqueID,omitempty"`
	ExtendedKeyUsage *[]string `json:"extendedKeyUsage,omitempty"`
	CriticalExtensionOIDs *[]string `json:"criticalExtensionOIDs,omitempty"`
	NonCriticalExtensionOIDs *[]string `json:"nonCriticalExtensionOIDs,omitempty"`
	Encoded *[]string `json:"encoded,omitempty"`
	PublicKey *PublicKey `json:"publicKey,omitempty"`
}

// NewX509Certificate instantiates a new X509Certificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewX509Certificate() *X509Certificate {
	this := X509Certificate{}
	return &this
}

// NewX509CertificateWithDefaults instantiates a new X509Certificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewX509CertificateWithDefaults() *X509Certificate {
	this := X509Certificate{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *X509Certificate) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509Certificate) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *X509Certificate) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *X509Certificate) SetType(v string) {
	o.Type = &v
}

// GetSubjectX500Principal returns the SubjectX500Principal field value if set, zero value otherwise.
func (o *X509Certificate) GetSubjectX500Principal() X500Principal {
	if o == nil || o.SubjectX500Principal == nil {
		var ret X500Principal
		return ret
	}
	return *o.SubjectX500Principal
}

// GetSubjectX500PrincipalOk returns a tuple with the SubjectX500Principal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509Certificate) GetSubjectX500PrincipalOk() (*X500Principal, bool) {
	if o == nil || o.SubjectX500Principal == nil {
		return nil, false
	}
	return o.SubjectX500Principal, true
}

// HasSubjectX500Principal returns a boolean if a field has been set.
func (o *X509Certificate) HasSubjectX500Principal() bool {
	if o != nil && o.SubjectX500Principal != nil {
		return true
	}

	return false
}

// SetSubjectX500Principal gets a reference to the given X500Principal and assigns it to the SubjectX500Principal field.
func (o *X509Certificate) SetSubjectX500Principal(v X500Principal) {
	o.SubjectX500Principal = &v
}

// GetIssuerX500Principal returns the IssuerX500Principal field value if set, zero value otherwise.
func (o *X509Certificate) GetIssuerX500Principal() X500Principal {
	if o == nil || o.IssuerX500Principal == nil {
		var ret X500Principal
		return ret
	}
	return *o.IssuerX500Principal
}

// GetIssuerX500PrincipalOk returns a tuple with the IssuerX500Principal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509Certificate) GetIssuerX500PrincipalOk() (*X500Principal, bool) {
	if o == nil || o.IssuerX500Principal == nil {
		return nil, false
	}
	return o.IssuerX500Principal, true
}

// HasIssuerX500Principal returns a boolean if a field has been set.
func (o *X509Certificate) HasIssuerX500Principal() bool {
	if o != nil && o.IssuerX500Principal != nil {
		return true
	}

	return false
}

// SetIssuerX500Principal gets a reference to the given X500Principal and assigns it to the IssuerX500Principal field.
func (o *X509Certificate) SetIssuerX500Principal(v X500Principal) {
	o.IssuerX500Principal = &v
}

// GetSubjectAlternativeNames returns the SubjectAlternativeNames field value if set, zero value otherwise.
func (o *X509Certificate) GetSubjectAlternativeNames() [][]map[string]interface{} {
	if o == nil || o.SubjectAlternativeNames == nil {
		var ret [][]map[string]interface{}
		return ret
	}
	return *o.SubjectAlternativeNames
}

// GetSubjectAlternativeNamesOk returns a tuple with the SubjectAlternativeNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509Certificate) GetSubjectAlternativeNamesOk() (*[][]map[string]interface{}, bool) {
	if o == nil || o.SubjectAlternativeNames == nil {
		return nil, false
	}
	return o.SubjectAlternativeNames, true
}

// HasSubjectAlternativeNames returns a boolean if a field has been set.
func (o *X509Certificate) HasSubjectAlternativeNames() bool {
	if o != nil && o.SubjectAlternativeNames != nil {
		return true
	}

	return false
}

// SetSubjectAlternativeNames gets a reference to the given [][]map[string]interface{} and assigns it to the SubjectAlternativeNames field.
func (o *X509Certificate) SetSubjectAlternativeNames(v [][]map[string]interface{}) {
	o.SubjectAlternativeNames = &v
}

// GetIssuerAlternativeNames returns the IssuerAlternativeNames field value if set, zero value otherwise.
func (o *X509Certificate) GetIssuerAlternativeNames() [][]map[string]interface{} {
	if o == nil || o.IssuerAlternativeNames == nil {
		var ret [][]map[string]interface{}
		return ret
	}
	return *o.IssuerAlternativeNames
}

// GetIssuerAlternativeNamesOk returns a tuple with the IssuerAlternativeNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509Certificate) GetIssuerAlternativeNamesOk() (*[][]map[string]interface{}, bool) {
	if o == nil || o.IssuerAlternativeNames == nil {
		return nil, false
	}
	return o.IssuerAlternativeNames, true
}

// HasIssuerAlternativeNames returns a boolean if a field has been set.
func (o *X509Certificate) HasIssuerAlternativeNames() bool {
	if o != nil && o.IssuerAlternativeNames != nil {
		return true
	}

	return false
}

// SetIssuerAlternativeNames gets a reference to the given [][]map[string]interface{} and assigns it to the IssuerAlternativeNames field.
func (o *X509Certificate) SetIssuerAlternativeNames(v [][]map[string]interface{}) {
	o.IssuerAlternativeNames = &v
}

// GetKeyUsage returns the KeyUsage field value if set, zero value otherwise.
func (o *X509Certificate) GetKeyUsage() []bool {
	if o == nil || o.KeyUsage == nil {
		var ret []bool
		return ret
	}
	return *o.KeyUsage
}

// GetKeyUsageOk returns a tuple with the KeyUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509Certificate) GetKeyUsageOk() (*[]bool, bool) {
	if o == nil || o.KeyUsage == nil {
		return nil, false
	}
	return o.KeyUsage, true
}

// HasKeyUsage returns a boolean if a field has been set.
func (o *X509Certificate) HasKeyUsage() bool {
	if o != nil && o.KeyUsage != nil {
		return true
	}

	return false
}

// SetKeyUsage gets a reference to the given []bool and assigns it to the KeyUsage field.
func (o *X509Certificate) SetKeyUsage(v []bool) {
	o.KeyUsage = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *X509Certificate) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509Certificate) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *X509Certificate) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *X509Certificate) SetVersion(v int32) {
	o.Version = &v
}

// GetIssuerDN returns the IssuerDN field value if set, zero value otherwise.
func (o *X509Certificate) GetIssuerDN() Principal {
	if o == nil || o.IssuerDN == nil {
		var ret Principal
		return ret
	}
	return *o.IssuerDN
}

// GetIssuerDNOk returns a tuple with the IssuerDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509Certificate) GetIssuerDNOk() (*Principal, bool) {
	if o == nil || o.IssuerDN == nil {
		return nil, false
	}
	return o.IssuerDN, true
}

// HasIssuerDN returns a boolean if a field has been set.
func (o *X509Certificate) HasIssuerDN() bool {
	if o != nil && o.IssuerDN != nil {
		return true
	}

	return false
}

// SetIssuerDN gets a reference to the given Principal and assigns it to the IssuerDN field.
func (o *X509Certificate) SetIssuerDN(v Principal) {
	o.IssuerDN = &v
}

// GetSubjectDN returns the SubjectDN field value if set, zero value otherwise.
func (o *X509Certificate) GetSubjectDN() Principal {
	if o == nil || o.SubjectDN == nil {
		var ret Principal
		return ret
	}
	return *o.SubjectDN
}

// GetSubjectDNOk returns a tuple with the SubjectDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509Certificate) GetSubjectDNOk() (*Principal, bool) {
	if o == nil || o.SubjectDN == nil {
		return nil, false
	}
	return o.SubjectDN, true
}

// HasSubjectDN returns a boolean if a field has been set.
func (o *X509Certificate) HasSubjectDN() bool {
	if o != nil && o.SubjectDN != nil {
		return true
	}

	return false
}

// SetSubjectDN gets a reference to the given Principal and assigns it to the SubjectDN field.
func (o *X509Certificate) SetSubjectDN(v Principal) {
	o.SubjectDN = &v
}

// GetSignature returns the Signature field value if set, zero value otherwise.
func (o *X509Certificate) GetSignature() []string {
	if o == nil || o.Signature == nil {
		var ret []string
		return ret
	}
	return *o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509Certificate) GetSignatureOk() (*[]string, bool) {
	if o == nil || o.Signature == nil {
		return nil, false
	}
	return o.Signature, true
}

// HasSignature returns a boolean if a field has been set.
func (o *X509Certificate) HasSignature() bool {
	if o != nil && o.Signature != nil {
		return true
	}

	return false
}

// SetSignature gets a reference to the given []string and assigns it to the Signature field.
func (o *X509Certificate) SetSignature(v []string) {
	o.Signature = &v
}

// GetBasicConstraints returns the BasicConstraints field value if set, zero value otherwise.
func (o *X509Certificate) GetBasicConstraints() int32 {
	if o == nil || o.BasicConstraints == nil {
		var ret int32
		return ret
	}
	return *o.BasicConstraints
}

// GetBasicConstraintsOk returns a tuple with the BasicConstraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509Certificate) GetBasicConstraintsOk() (*int32, bool) {
	if o == nil || o.BasicConstraints == nil {
		return nil, false
	}
	return o.BasicConstraints, true
}

// HasBasicConstraints returns a boolean if a field has been set.
func (o *X509Certificate) HasBasicConstraints() bool {
	if o != nil && o.BasicConstraints != nil {
		return true
	}

	return false
}

// SetBasicConstraints gets a reference to the given int32 and assigns it to the BasicConstraints field.
func (o *X509Certificate) SetBasicConstraints(v int32) {
	o.BasicConstraints = &v
}

// GetSigAlgName returns the SigAlgName field value if set, zero value otherwise.
func (o *X509Certificate) GetSigAlgName() string {
	if o == nil || o.SigAlgName == nil {
		var ret string
		return ret
	}
	return *o.SigAlgName
}

// GetSigAlgNameOk returns a tuple with the SigAlgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509Certificate) GetSigAlgNameOk() (*string, bool) {
	if o == nil || o.SigAlgName == nil {
		return nil, false
	}
	return o.SigAlgName, true
}

// HasSigAlgName returns a boolean if a field has been set.
func (o *X509Certificate) HasSigAlgName() bool {
	if o != nil && o.SigAlgName != nil {
		return true
	}

	return false
}

// SetSigAlgName gets a reference to the given string and assigns it to the SigAlgName field.
func (o *X509Certificate) SetSigAlgName(v string) {
	o.SigAlgName = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *X509Certificate) GetSerialNumber() int32 {
	if o == nil || o.SerialNumber == nil {
		var ret int32
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509Certificate) GetSerialNumberOk() (*int32, bool) {
	if o == nil || o.SerialNumber == nil {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *X509Certificate) HasSerialNumber() bool {
	if o != nil && o.SerialNumber != nil {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given int32 and assigns it to the SerialNumber field.
func (o *X509Certificate) SetSerialNumber(v int32) {
	o.SerialNumber = &v
}

// GetNotBefore returns the NotBefore field value if set, zero value otherwise.
func (o *X509Certificate) GetNotBefore() time.Time {
	if o == nil || o.NotBefore == nil {
		var ret time.Time
		return ret
	}
	return *o.NotBefore
}

// GetNotBeforeOk returns a tuple with the NotBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509Certificate) GetNotBeforeOk() (*time.Time, bool) {
	if o == nil || o.NotBefore == nil {
		return nil, false
	}
	return o.NotBefore, true
}

// HasNotBefore returns a boolean if a field has been set.
func (o *X509Certificate) HasNotBefore() bool {
	if o != nil && o.NotBefore != nil {
		return true
	}

	return false
}

// SetNotBefore gets a reference to the given time.Time and assigns it to the NotBefore field.
func (o *X509Certificate) SetNotBefore(v time.Time) {
	o.NotBefore = &v
}

// GetNotAfter returns the NotAfter field value if set, zero value otherwise.
func (o *X509Certificate) GetNotAfter() time.Time {
	if o == nil || o.NotAfter == nil {
		var ret time.Time
		return ret
	}
	return *o.NotAfter
}

// GetNotAfterOk returns a tuple with the NotAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509Certificate) GetNotAfterOk() (*time.Time, bool) {
	if o == nil || o.NotAfter == nil {
		return nil, false
	}
	return o.NotAfter, true
}

// HasNotAfter returns a boolean if a field has been set.
func (o *X509Certificate) HasNotAfter() bool {
	if o != nil && o.NotAfter != nil {
		return true
	}

	return false
}

// SetNotAfter gets a reference to the given time.Time and assigns it to the NotAfter field.
func (o *X509Certificate) SetNotAfter(v time.Time) {
	o.NotAfter = &v
}

// GetTbscertificate returns the Tbscertificate field value if set, zero value otherwise.
func (o *X509Certificate) GetTbscertificate() []string {
	if o == nil || o.Tbscertificate == nil {
		var ret []string
		return ret
	}
	return *o.Tbscertificate
}

// GetTbscertificateOk returns a tuple with the Tbscertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509Certificate) GetTbscertificateOk() (*[]string, bool) {
	if o == nil || o.Tbscertificate == nil {
		return nil, false
	}
	return o.Tbscertificate, true
}

// HasTbscertificate returns a boolean if a field has been set.
func (o *X509Certificate) HasTbscertificate() bool {
	if o != nil && o.Tbscertificate != nil {
		return true
	}

	return false
}

// SetTbscertificate gets a reference to the given []string and assigns it to the Tbscertificate field.
func (o *X509Certificate) SetTbscertificate(v []string) {
	o.Tbscertificate = &v
}

// GetSigAlgOID returns the SigAlgOID field value if set, zero value otherwise.
func (o *X509Certificate) GetSigAlgOID() string {
	if o == nil || o.SigAlgOID == nil {
		var ret string
		return ret
	}
	return *o.SigAlgOID
}

// GetSigAlgOIDOk returns a tuple with the SigAlgOID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509Certificate) GetSigAlgOIDOk() (*string, bool) {
	if o == nil || o.SigAlgOID == nil {
		return nil, false
	}
	return o.SigAlgOID, true
}

// HasSigAlgOID returns a boolean if a field has been set.
func (o *X509Certificate) HasSigAlgOID() bool {
	if o != nil && o.SigAlgOID != nil {
		return true
	}

	return false
}

// SetSigAlgOID gets a reference to the given string and assigns it to the SigAlgOID field.
func (o *X509Certificate) SetSigAlgOID(v string) {
	o.SigAlgOID = &v
}

// GetSigAlgParams returns the SigAlgParams field value if set, zero value otherwise.
func (o *X509Certificate) GetSigAlgParams() []string {
	if o == nil || o.SigAlgParams == nil {
		var ret []string
		return ret
	}
	return *o.SigAlgParams
}

// GetSigAlgParamsOk returns a tuple with the SigAlgParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509Certificate) GetSigAlgParamsOk() (*[]string, bool) {
	if o == nil || o.SigAlgParams == nil {
		return nil, false
	}
	return o.SigAlgParams, true
}

// HasSigAlgParams returns a boolean if a field has been set.
func (o *X509Certificate) HasSigAlgParams() bool {
	if o != nil && o.SigAlgParams != nil {
		return true
	}

	return false
}

// SetSigAlgParams gets a reference to the given []string and assigns it to the SigAlgParams field.
func (o *X509Certificate) SetSigAlgParams(v []string) {
	o.SigAlgParams = &v
}

// GetIssuerUniqueID returns the IssuerUniqueID field value if set, zero value otherwise.
func (o *X509Certificate) GetIssuerUniqueID() []bool {
	if o == nil || o.IssuerUniqueID == nil {
		var ret []bool
		return ret
	}
	return *o.IssuerUniqueID
}

// GetIssuerUniqueIDOk returns a tuple with the IssuerUniqueID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509Certificate) GetIssuerUniqueIDOk() (*[]bool, bool) {
	if o == nil || o.IssuerUniqueID == nil {
		return nil, false
	}
	return o.IssuerUniqueID, true
}

// HasIssuerUniqueID returns a boolean if a field has been set.
func (o *X509Certificate) HasIssuerUniqueID() bool {
	if o != nil && o.IssuerUniqueID != nil {
		return true
	}

	return false
}

// SetIssuerUniqueID gets a reference to the given []bool and assigns it to the IssuerUniqueID field.
func (o *X509Certificate) SetIssuerUniqueID(v []bool) {
	o.IssuerUniqueID = &v
}

// GetSubjectUniqueID returns the SubjectUniqueID field value if set, zero value otherwise.
func (o *X509Certificate) GetSubjectUniqueID() []bool {
	if o == nil || o.SubjectUniqueID == nil {
		var ret []bool
		return ret
	}
	return *o.SubjectUniqueID
}

// GetSubjectUniqueIDOk returns a tuple with the SubjectUniqueID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509Certificate) GetSubjectUniqueIDOk() (*[]bool, bool) {
	if o == nil || o.SubjectUniqueID == nil {
		return nil, false
	}
	return o.SubjectUniqueID, true
}

// HasSubjectUniqueID returns a boolean if a field has been set.
func (o *X509Certificate) HasSubjectUniqueID() bool {
	if o != nil && o.SubjectUniqueID != nil {
		return true
	}

	return false
}

// SetSubjectUniqueID gets a reference to the given []bool and assigns it to the SubjectUniqueID field.
func (o *X509Certificate) SetSubjectUniqueID(v []bool) {
	o.SubjectUniqueID = &v
}

// GetExtendedKeyUsage returns the ExtendedKeyUsage field value if set, zero value otherwise.
func (o *X509Certificate) GetExtendedKeyUsage() []string {
	if o == nil || o.ExtendedKeyUsage == nil {
		var ret []string
		return ret
	}
	return *o.ExtendedKeyUsage
}

// GetExtendedKeyUsageOk returns a tuple with the ExtendedKeyUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509Certificate) GetExtendedKeyUsageOk() (*[]string, bool) {
	if o == nil || o.ExtendedKeyUsage == nil {
		return nil, false
	}
	return o.ExtendedKeyUsage, true
}

// HasExtendedKeyUsage returns a boolean if a field has been set.
func (o *X509Certificate) HasExtendedKeyUsage() bool {
	if o != nil && o.ExtendedKeyUsage != nil {
		return true
	}

	return false
}

// SetExtendedKeyUsage gets a reference to the given []string and assigns it to the ExtendedKeyUsage field.
func (o *X509Certificate) SetExtendedKeyUsage(v []string) {
	o.ExtendedKeyUsage = &v
}

// GetCriticalExtensionOIDs returns the CriticalExtensionOIDs field value if set, zero value otherwise.
func (o *X509Certificate) GetCriticalExtensionOIDs() []string {
	if o == nil || o.CriticalExtensionOIDs == nil {
		var ret []string
		return ret
	}
	return *o.CriticalExtensionOIDs
}

// GetCriticalExtensionOIDsOk returns a tuple with the CriticalExtensionOIDs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509Certificate) GetCriticalExtensionOIDsOk() (*[]string, bool) {
	if o == nil || o.CriticalExtensionOIDs == nil {
		return nil, false
	}
	return o.CriticalExtensionOIDs, true
}

// HasCriticalExtensionOIDs returns a boolean if a field has been set.
func (o *X509Certificate) HasCriticalExtensionOIDs() bool {
	if o != nil && o.CriticalExtensionOIDs != nil {
		return true
	}

	return false
}

// SetCriticalExtensionOIDs gets a reference to the given []string and assigns it to the CriticalExtensionOIDs field.
func (o *X509Certificate) SetCriticalExtensionOIDs(v []string) {
	o.CriticalExtensionOIDs = &v
}

// GetNonCriticalExtensionOIDs returns the NonCriticalExtensionOIDs field value if set, zero value otherwise.
func (o *X509Certificate) GetNonCriticalExtensionOIDs() []string {
	if o == nil || o.NonCriticalExtensionOIDs == nil {
		var ret []string
		return ret
	}
	return *o.NonCriticalExtensionOIDs
}

// GetNonCriticalExtensionOIDsOk returns a tuple with the NonCriticalExtensionOIDs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509Certificate) GetNonCriticalExtensionOIDsOk() (*[]string, bool) {
	if o == nil || o.NonCriticalExtensionOIDs == nil {
		return nil, false
	}
	return o.NonCriticalExtensionOIDs, true
}

// HasNonCriticalExtensionOIDs returns a boolean if a field has been set.
func (o *X509Certificate) HasNonCriticalExtensionOIDs() bool {
	if o != nil && o.NonCriticalExtensionOIDs != nil {
		return true
	}

	return false
}

// SetNonCriticalExtensionOIDs gets a reference to the given []string and assigns it to the NonCriticalExtensionOIDs field.
func (o *X509Certificate) SetNonCriticalExtensionOIDs(v []string) {
	o.NonCriticalExtensionOIDs = &v
}

// GetEncoded returns the Encoded field value if set, zero value otherwise.
func (o *X509Certificate) GetEncoded() []string {
	if o == nil || o.Encoded == nil {
		var ret []string
		return ret
	}
	return *o.Encoded
}

// GetEncodedOk returns a tuple with the Encoded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509Certificate) GetEncodedOk() (*[]string, bool) {
	if o == nil || o.Encoded == nil {
		return nil, false
	}
	return o.Encoded, true
}

// HasEncoded returns a boolean if a field has been set.
func (o *X509Certificate) HasEncoded() bool {
	if o != nil && o.Encoded != nil {
		return true
	}

	return false
}

// SetEncoded gets a reference to the given []string and assigns it to the Encoded field.
func (o *X509Certificate) SetEncoded(v []string) {
	o.Encoded = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *X509Certificate) GetPublicKey() PublicKey {
	if o == nil || o.PublicKey == nil {
		var ret PublicKey
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509Certificate) GetPublicKeyOk() (*PublicKey, bool) {
	if o == nil || o.PublicKey == nil {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *X509Certificate) HasPublicKey() bool {
	if o != nil && o.PublicKey != nil {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given PublicKey and assigns it to the PublicKey field.
func (o *X509Certificate) SetPublicKey(v PublicKey) {
	o.PublicKey = &v
}

func (o X509Certificate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.SubjectX500Principal != nil {
		toSerialize["subjectX500Principal"] = o.SubjectX500Principal
	}
	if o.IssuerX500Principal != nil {
		toSerialize["issuerX500Principal"] = o.IssuerX500Principal
	}
	if o.SubjectAlternativeNames != nil {
		toSerialize["subjectAlternativeNames"] = o.SubjectAlternativeNames
	}
	if o.IssuerAlternativeNames != nil {
		toSerialize["issuerAlternativeNames"] = o.IssuerAlternativeNames
	}
	if o.KeyUsage != nil {
		toSerialize["keyUsage"] = o.KeyUsage
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.IssuerDN != nil {
		toSerialize["issuerDN"] = o.IssuerDN
	}
	if o.SubjectDN != nil {
		toSerialize["subjectDN"] = o.SubjectDN
	}
	if o.Signature != nil {
		toSerialize["signature"] = o.Signature
	}
	if o.BasicConstraints != nil {
		toSerialize["basicConstraints"] = o.BasicConstraints
	}
	if o.SigAlgName != nil {
		toSerialize["sigAlgName"] = o.SigAlgName
	}
	if o.SerialNumber != nil {
		toSerialize["serialNumber"] = o.SerialNumber
	}
	if o.NotBefore != nil {
		toSerialize["notBefore"] = o.NotBefore
	}
	if o.NotAfter != nil {
		toSerialize["notAfter"] = o.NotAfter
	}
	if o.Tbscertificate != nil {
		toSerialize["tbscertificate"] = o.Tbscertificate
	}
	if o.SigAlgOID != nil {
		toSerialize["sigAlgOID"] = o.SigAlgOID
	}
	if o.SigAlgParams != nil {
		toSerialize["sigAlgParams"] = o.SigAlgParams
	}
	if o.IssuerUniqueID != nil {
		toSerialize["issuerUniqueID"] = o.IssuerUniqueID
	}
	if o.SubjectUniqueID != nil {
		toSerialize["subjectUniqueID"] = o.SubjectUniqueID
	}
	if o.ExtendedKeyUsage != nil {
		toSerialize["extendedKeyUsage"] = o.ExtendedKeyUsage
	}
	if o.CriticalExtensionOIDs != nil {
		toSerialize["criticalExtensionOIDs"] = o.CriticalExtensionOIDs
	}
	if o.NonCriticalExtensionOIDs != nil {
		toSerialize["nonCriticalExtensionOIDs"] = o.NonCriticalExtensionOIDs
	}
	if o.Encoded != nil {
		toSerialize["encoded"] = o.Encoded
	}
	if o.PublicKey != nil {
		toSerialize["publicKey"] = o.PublicKey
	}
	return json.Marshal(toSerialize)
}

type NullableX509Certificate struct {
	value *X509Certificate
	isSet bool
}

func (v NullableX509Certificate) Get() *X509Certificate {
	return v.value
}

func (v *NullableX509Certificate) Set(val *X509Certificate) {
	v.value = val
	v.isSet = true
}

func (v NullableX509Certificate) IsSet() bool {
	return v.isSet
}

func (v *NullableX509Certificate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableX509Certificate(val *X509Certificate) *NullableX509Certificate {
	return &NullableX509Certificate{value: val, isSet: true}
}

func (v NullableX509Certificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableX509Certificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


