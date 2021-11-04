/*
Pulsar Admin REST API

This provides the REST API for admin operations

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// ResourceQuotasApiService ResourceQuotasApi service
type ResourceQuotasApiService service

type ApiGetDefaultResourceQuotaRequest struct {
	ctx _context.Context
	ApiService *ResourceQuotasApiService
}


func (r ApiGetDefaultResourceQuotaRequest) Execute() ([]string, *_nethttp.Response, error) {
	return r.ApiService.GetDefaultResourceQuotaExecute(r)
}

/*
GetDefaultResourceQuota Get the default quota

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetDefaultResourceQuotaRequest
*/
func (a *ResourceQuotasApiService) GetDefaultResourceQuota(ctx _context.Context) ApiGetDefaultResourceQuotaRequest {
	return ApiGetDefaultResourceQuotaRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []string
func (a *ResourceQuotasApiService) GetDefaultResourceQuotaExecute(r ApiGetDefaultResourceQuotaRequest) ([]string, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourceQuotasApiService.GetDefaultResourceQuota")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resource-quotas"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetNamespaceBundleResourceQuotaRequest struct {
	ctx _context.Context
	ApiService *ResourceQuotasApiService
	tenant string
	namespace string
	bundle string
}


func (r ApiGetNamespaceBundleResourceQuotaRequest) Execute() (ResourceQuota, *_nethttp.Response, error) {
	return r.ApiService.GetNamespaceBundleResourceQuotaExecute(r)
}

/*
GetNamespaceBundleResourceQuota Get resource quota of a namespace bundle.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tenant Tenant name
 @param namespace Namespace name within the specified tenant
 @param bundle Namespace bundle range
 @return ApiGetNamespaceBundleResourceQuotaRequest
*/
func (a *ResourceQuotasApiService) GetNamespaceBundleResourceQuota(ctx _context.Context, tenant string, namespace string, bundle string) ApiGetNamespaceBundleResourceQuotaRequest {
	return ApiGetNamespaceBundleResourceQuotaRequest{
		ApiService: a,
		ctx: ctx,
		tenant: tenant,
		namespace: namespace,
		bundle: bundle,
	}
}

// Execute executes the request
//  @return ResourceQuota
func (a *ResourceQuotasApiService) GetNamespaceBundleResourceQuotaExecute(r ApiGetNamespaceBundleResourceQuotaRequest) (ResourceQuota, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ResourceQuota
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourceQuotasApiService.GetNamespaceBundleResourceQuota")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resource-quotas/{tenant}/{namespace}/{bundle}"
	localVarPath = strings.Replace(localVarPath, "{"+"tenant"+"}", _neturl.PathEscape(parameterToString(r.tenant, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"namespace"+"}", _neturl.PathEscape(parameterToString(r.namespace, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bundle"+"}", _neturl.PathEscape(parameterToString(r.bundle, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiRemoveNamespaceBundleResourceQuotaRequest struct {
	ctx _context.Context
	ApiService *ResourceQuotasApiService
	tenant string
	namespace string
	bundle string
}


func (r ApiRemoveNamespaceBundleResourceQuotaRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.RemoveNamespaceBundleResourceQuotaExecute(r)
}

/*
RemoveNamespaceBundleResourceQuota Remove resource quota for a namespace.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tenant Tenant name
 @param namespace Namespace name within the specified tenant
 @param bundle Namespace bundle range
 @return ApiRemoveNamespaceBundleResourceQuotaRequest
*/
func (a *ResourceQuotasApiService) RemoveNamespaceBundleResourceQuota(ctx _context.Context, tenant string, namespace string, bundle string) ApiRemoveNamespaceBundleResourceQuotaRequest {
	return ApiRemoveNamespaceBundleResourceQuotaRequest{
		ApiService: a,
		ctx: ctx,
		tenant: tenant,
		namespace: namespace,
		bundle: bundle,
	}
}

// Execute executes the request
func (a *ResourceQuotasApiService) RemoveNamespaceBundleResourceQuotaExecute(r ApiRemoveNamespaceBundleResourceQuotaRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourceQuotasApiService.RemoveNamespaceBundleResourceQuota")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resource-quotas/{tenant}/{namespace}/{bundle}"
	localVarPath = strings.Replace(localVarPath, "{"+"tenant"+"}", _neturl.PathEscape(parameterToString(r.tenant, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"namespace"+"}", _neturl.PathEscape(parameterToString(r.namespace, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bundle"+"}", _neturl.PathEscape(parameterToString(r.bundle, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiSetDefaultResourceQuotaRequest struct {
	ctx _context.Context
	ApiService *ResourceQuotasApiService
	body *ResourceQuota
}

// Default resource quota
func (r ApiSetDefaultResourceQuotaRequest) Body(body ResourceQuota) ApiSetDefaultResourceQuotaRequest {
	r.body = &body
	return r
}

func (r ApiSetDefaultResourceQuotaRequest) Execute() ([]string, *_nethttp.Response, error) {
	return r.ApiService.SetDefaultResourceQuotaExecute(r)
}

/*
SetDefaultResourceQuota Set the default quota

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSetDefaultResourceQuotaRequest
*/
func (a *ResourceQuotasApiService) SetDefaultResourceQuota(ctx _context.Context) ApiSetDefaultResourceQuotaRequest {
	return ApiSetDefaultResourceQuotaRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []string
func (a *ResourceQuotasApiService) SetDefaultResourceQuotaExecute(r ApiSetDefaultResourceQuotaRequest) ([]string, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourceQuotasApiService.SetDefaultResourceQuota")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resource-quotas"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSetNamespaceBundleResourceQuotaRequest struct {
	ctx _context.Context
	ApiService *ResourceQuotasApiService
	tenant string
	namespace string
	bundle string
	body *ResourceQuota
}

// Resource quota for the specified namespace
func (r ApiSetNamespaceBundleResourceQuotaRequest) Body(body ResourceQuota) ApiSetNamespaceBundleResourceQuotaRequest {
	r.body = &body
	return r
}

func (r ApiSetNamespaceBundleResourceQuotaRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.SetNamespaceBundleResourceQuotaExecute(r)
}

/*
SetNamespaceBundleResourceQuota Set resource quota on a namespace.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tenant Tenant name
 @param namespace Namespace name within the specified tenant
 @param bundle Namespace bundle range
 @return ApiSetNamespaceBundleResourceQuotaRequest
*/
func (a *ResourceQuotasApiService) SetNamespaceBundleResourceQuota(ctx _context.Context, tenant string, namespace string, bundle string) ApiSetNamespaceBundleResourceQuotaRequest {
	return ApiSetNamespaceBundleResourceQuotaRequest{
		ApiService: a,
		ctx: ctx,
		tenant: tenant,
		namespace: namespace,
		bundle: bundle,
	}
}

// Execute executes the request
func (a *ResourceQuotasApiService) SetNamespaceBundleResourceQuotaExecute(r ApiSetNamespaceBundleResourceQuotaRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourceQuotasApiService.SetNamespaceBundleResourceQuota")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resource-quotas/{tenant}/{namespace}/{bundle}"
	localVarPath = strings.Replace(localVarPath, "{"+"tenant"+"}", _neturl.PathEscape(parameterToString(r.tenant, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"namespace"+"}", _neturl.PathEscape(parameterToString(r.namespace, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bundle"+"}", _neturl.PathEscape(parameterToString(r.bundle, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
