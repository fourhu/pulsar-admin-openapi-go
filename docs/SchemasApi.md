# \SchemasApi

All URIs are relative to *http://localhost/admin/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSchema**](SchemasApi.md#DeleteSchema) | **Delete** /schemas/{tenant}/{namespace}/{topic}/schema | Delete the schema of a topic
[**GetAllSchemas**](SchemasApi.md#GetAllSchemas) | **Get** /schemas/{tenant}/{namespace}/{topic}/schemas | Get the all schemas of a topic
[**GetSchema**](SchemasApi.md#GetSchema) | **Get** /schemas/{tenant}/{namespace}/{topic}/schema | Get the schema of a topic
[**GetSchemaVersion**](SchemasApi.md#GetSchemaVersion) | **Get** /schemas/{tenant}/{namespace}/{topic}/schema/{version} | Get the schema of a topic at a given version
[**GetVersionBySchema**](SchemasApi.md#GetVersionBySchema) | **Post** /schemas/{tenant}/{namespace}/{topic}/version | get the version of the schema
[**PostSchema**](SchemasApi.md#PostSchema) | **Post** /schemas/{tenant}/{namespace}/{topic}/schema | Update the schema of a topic
[**TestCompatibility**](SchemasApi.md#TestCompatibility) | **Post** /schemas/{tenant}/{namespace}/{topic}/compatibility | test the schema compatibility



## DeleteSchema

> DeleteSchemaResponse DeleteSchema(ctx, tenant, namespace, topic).Authoritative(authoritative).Execute()

Delete the schema of a topic

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
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SchemasApi.DeleteSchema(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.DeleteSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSchema`: DeleteSchemaResponse
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.DeleteSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** |  | [default to false]

### Return type

[**DeleteSchemaResponse**](DeleteSchemaResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllSchemas

> GetAllVersionsSchemaResponse GetAllSchemas(ctx, tenant, namespace, topic).Authoritative(authoritative).Execute()

Get the all schemas of a topic

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
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SchemasApi.GetAllSchemas(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.GetAllSchemas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllSchemas`: GetAllVersionsSchemaResponse
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.GetAllSchemas`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllSchemasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** |  | [default to false]

### Return type

[**GetAllVersionsSchemaResponse**](GetAllVersionsSchemaResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSchema

> GetSchemaResponse GetSchema(ctx, tenant, namespace, topic).Authoritative(authoritative).Execute()

Get the schema of a topic

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
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SchemasApi.GetSchema(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.GetSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSchema`: GetSchemaResponse
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.GetSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** |  | [default to false]

### Return type

[**GetSchemaResponse**](GetSchemaResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSchemaVersion

> GetSchemaResponse GetSchemaVersion(ctx, tenant, namespace, topic, version).Authoritative(authoritative).Execute()

Get the schema of a topic at a given version

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
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    version := "version_example" // string | 
    authoritative := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SchemasApi.GetSchemaVersion(context.Background(), tenant, namespace, topic, version).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.GetSchemaVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSchemaVersion`: GetSchemaResponse
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.GetSchemaVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 
**version** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSchemaVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **authoritative** | **bool** |  | [default to false]

### Return type

[**GetSchemaResponse**](GetSchemaResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVersionBySchema

> LongSchemaVersion GetVersionBySchema(ctx, tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()

get the version of the schema

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
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool |  (optional) (default to false)
    body := *openapiclient.NewPostSchemaPayload() // PostSchemaPayload | A JSON value presenting a schema playload. An example of the expected schema can be found down here. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SchemasApi.GetVersionBySchema(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.GetVersionBySchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVersionBySchema`: LongSchemaVersion
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.GetVersionBySchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVersionBySchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** |  | [default to false]
 **body** | [**PostSchemaPayload**](PostSchemaPayload.md) | A JSON value presenting a schema playload. An example of the expected schema can be found down here. | 

### Return type

[**LongSchemaVersion**](LongSchemaVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSchema

> PostSchemaResponse PostSchema(ctx, tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()

Update the schema of a topic

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
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool |  (optional) (default to false)
    body := *openapiclient.NewPostSchemaPayload() // PostSchemaPayload | A JSON value presenting a schema playload. An example of the expected schema can be found down here. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SchemasApi.PostSchema(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.PostSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostSchema`: PostSchemaResponse
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.PostSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** |  | [default to false]
 **body** | [**PostSchemaPayload**](PostSchemaPayload.md) | A JSON value presenting a schema playload. An example of the expected schema can be found down here. | 

### Return type

[**PostSchemaResponse**](PostSchemaResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestCompatibility

> IsCompatibilityResponse TestCompatibility(ctx, tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()

test the schema compatibility

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
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool |  (optional) (default to false)
    body := *openapiclient.NewPostSchemaPayload() // PostSchemaPayload | A JSON value presenting a schema playload. An example of the expected schema can be found down here. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SchemasApi.TestCompatibility(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.TestCompatibility``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestCompatibility`: IsCompatibilityResponse
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.TestCompatibility`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestCompatibilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** |  | [default to false]
 **body** | [**PostSchemaPayload**](PostSchemaPayload.md) | A JSON value presenting a schema playload. An example of the expected schema can be found down here. | 

### Return type

[**IsCompatibilityResponse**](IsCompatibilityResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

