# \ResourceQuotasApi

All URIs are relative to *http://localhost/admin/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDefaultResourceQuota**](ResourceQuotasApi.md#GetDefaultResourceQuota) | **Get** /resource-quotas | Get the default quota
[**GetNamespaceBundleResourceQuota**](ResourceQuotasApi.md#GetNamespaceBundleResourceQuota) | **Get** /resource-quotas/{tenant}/{namespace}/{bundle} | Get resource quota of a namespace bundle.
[**RemoveNamespaceBundleResourceQuota**](ResourceQuotasApi.md#RemoveNamespaceBundleResourceQuota) | **Delete** /resource-quotas/{tenant}/{namespace}/{bundle} | Remove resource quota for a namespace.
[**SetDefaultResourceQuota**](ResourceQuotasApi.md#SetDefaultResourceQuota) | **Post** /resource-quotas | Set the default quota
[**SetNamespaceBundleResourceQuota**](ResourceQuotasApi.md#SetNamespaceBundleResourceQuota) | **Post** /resource-quotas/{tenant}/{namespace}/{bundle} | Set resource quota on a namespace.



## GetDefaultResourceQuota

> []string GetDefaultResourceQuota(ctx).Execute()

Get the default quota

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourceQuotasApi.GetDefaultResourceQuota(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceQuotasApi.GetDefaultResourceQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDefaultResourceQuota`: []string
    fmt.Fprintf(os.Stdout, "Response from `ResourceQuotasApi.GetDefaultResourceQuota`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultResourceQuotaRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamespaceBundleResourceQuota

> ResourceQuota GetNamespaceBundleResourceQuota(ctx, tenant, namespace, bundle).Execute()

Get resource quota of a namespace bundle.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | Tenant name
    namespace := "namespace_example" // string | Namespace name within the specified tenant
    bundle := "bundle_example" // string | Namespace bundle range

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourceQuotasApi.GetNamespaceBundleResourceQuota(context.Background(), tenant, namespace, bundle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceQuotasApi.GetNamespaceBundleResourceQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNamespaceBundleResourceQuota`: ResourceQuota
    fmt.Fprintf(os.Stdout, "Response from `ResourceQuotasApi.GetNamespaceBundleResourceQuota`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** | Tenant name | 
**namespace** | **string** | Namespace name within the specified tenant | 
**bundle** | **string** | Namespace bundle range | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamespaceBundleResourceQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ResourceQuota**](ResourceQuota.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveNamespaceBundleResourceQuota

> RemoveNamespaceBundleResourceQuota(ctx, tenant, namespace, bundle).Execute()

Remove resource quota for a namespace.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | Tenant name
    namespace := "namespace_example" // string | Namespace name within the specified tenant
    bundle := "bundle_example" // string | Namespace bundle range

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourceQuotasApi.RemoveNamespaceBundleResourceQuota(context.Background(), tenant, namespace, bundle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceQuotasApi.RemoveNamespaceBundleResourceQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** | Tenant name | 
**namespace** | **string** | Namespace name within the specified tenant | 
**bundle** | **string** | Namespace bundle range | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveNamespaceBundleResourceQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDefaultResourceQuota

> []string SetDefaultResourceQuota(ctx).Body(body).Execute()

Set the default quota

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewResourceQuota() // ResourceQuota | Default resource quota (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourceQuotasApi.SetDefaultResourceQuota(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceQuotasApi.SetDefaultResourceQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetDefaultResourceQuota`: []string
    fmt.Fprintf(os.Stdout, "Response from `ResourceQuotasApi.SetDefaultResourceQuota`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetDefaultResourceQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ResourceQuota**](ResourceQuota.md) | Default resource quota | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetNamespaceBundleResourceQuota

> SetNamespaceBundleResourceQuota(ctx, tenant, namespace, bundle).Body(body).Execute()

Set resource quota on a namespace.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | Tenant name
    namespace := "namespace_example" // string | Namespace name within the specified tenant
    bundle := "bundle_example" // string | Namespace bundle range
    body := *openapiclient.NewResourceQuota() // ResourceQuota | Resource quota for the specified namespace (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourceQuotasApi.SetNamespaceBundleResourceQuota(context.Background(), tenant, namespace, bundle).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceQuotasApi.SetNamespaceBundleResourceQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** | Tenant name | 
**namespace** | **string** | Namespace name within the specified tenant | 
**bundle** | **string** | Namespace bundle range | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetNamespaceBundleResourceQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**ResourceQuota**](ResourceQuota.md) | Resource quota for the specified namespace | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

