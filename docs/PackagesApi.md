# \PackagesApi

All URIs are relative to *http://localhost/admin/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Delete**](PackagesApi.md#Delete) | **Delete** /packages/{type}/{tenant}/{namespace}/{packageName}/{version} | Delete a package with the package name.
[**Download**](PackagesApi.md#Download) | **Get** /packages/{type}/{tenant}/{namespace}/{packageName}/{version} | Download a package with the package name.
[**GetMeta**](PackagesApi.md#GetMeta) | **Get** /packages/{type}/{tenant}/{namespace}/{packageName}/{version}/metadata | Get the metadata of a package.
[**ListPackageVersion**](PackagesApi.md#ListPackageVersion) | **Get** /packages/{type}/{tenant}/{namespace}/{packageName} | Get all the versions of a package.
[**ListPackages**](PackagesApi.md#ListPackages) | **Get** /packages/{type}/{tenant}/{namespace} | Get all the specified type packages in a namespace.
[**UpdateMeta**](PackagesApi.md#UpdateMeta) | **Put** /packages/{type}/{tenant}/{namespace}/{packageName}/{version}/metadata | Update the metadata of a package.
[**Upload**](PackagesApi.md#Upload) | **Post** /packages/{type}/{tenant}/{namespace}/{packageName}/{version} | Upload a package.



## Delete

> Delete(ctx, type_, tenant, namespace, packageName, version).Execute()

Delete a package with the package name.

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
    type_ := "type__example" // string | 
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    packageName := "packageName_example" // string | 
    version := "version_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PackagesApi.Delete(context.Background(), type_, tenant, namespace, packageName, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** |  | 
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**packageName** | **string** |  | 
**version** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


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


## Download

> map[string]interface{} Download(ctx, type_, tenant, namespace, packageName, version).Execute()

Download a package with the package name.

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
    type_ := "type__example" // string | 
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    packageName := "packageName_example" // string | 
    version := "version_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PackagesApi.Download(context.Background(), type_, tenant, namespace, packageName, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.Download``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Download`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.Download`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** |  | 
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**packageName** | **string** |  | 
**version** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMeta

> PackageMetadata GetMeta(ctx, type_, tenant, namespace, packageName, version).Execute()

Get the metadata of a package.

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
    type_ := "type__example" // string | 
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    packageName := "packageName_example" // string | 
    version := "version_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PackagesApi.GetMeta(context.Background(), type_, tenant, namespace, packageName, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.GetMeta``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMeta`: PackageMetadata
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.GetMeta`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** |  | 
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**packageName** | **string** |  | 
**version** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






### Return type

[**PackageMetadata**](PackageMetadata.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPackageVersion

> []string ListPackageVersion(ctx, type_, tenant, namespace, packageName).Execute()

Get all the versions of a package.

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
    type_ := "type__example" // string | 
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    packageName := "packageName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PackagesApi.ListPackageVersion(context.Background(), type_, tenant, namespace, packageName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.ListPackageVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPackageVersion`: []string
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.ListPackageVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** |  | 
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**packageName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPackageVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





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


## ListPackages

> PackageMetadata ListPackages(ctx, type_, tenant, namespace).Execute()

Get all the specified type packages in a namespace.

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
    type_ := "type__example" // string | 
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PackagesApi.ListPackages(context.Background(), type_, tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.ListPackages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPackages`: PackageMetadata
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.ListPackages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** |  | 
**tenant** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPackagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**PackageMetadata**](PackageMetadata.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMeta

> UpdateMeta(ctx, type_, tenant, namespace, packageName, version).Execute()

Update the metadata of a package.

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
    type_ := "type__example" // string | 
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    packageName := "packageName_example" // string | 
    version := "version_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PackagesApi.UpdateMeta(context.Background(), type_, tenant, namespace, packageName, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.UpdateMeta``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** |  | 
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**packageName** | **string** |  | 
**version** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMetaRequest struct via the builder pattern


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


## Upload

> Upload(ctx, type_, tenant, namespace, packageName, version).Execute()

Upload a package.

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
    type_ := "type__example" // string | 
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    packageName := "packageName_example" // string | 
    version := "version_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PackagesApi.Upload(context.Background(), type_, tenant, namespace, packageName, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.Upload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** |  | 
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**packageName** | **string** |  | 
**version** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadRequest struct via the builder pattern


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

