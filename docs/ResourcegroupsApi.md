# \ResourcegroupsApi

All URIs are relative to *http://localhost/admin/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrUpdateResourceGroup**](ResourcegroupsApi.md#CreateOrUpdateResourceGroup) | **Put** /resourcegroups/{resourcegroup} | Creates a new resourcegroup with the specified rate limiters
[**DeleteResourceGroup**](ResourcegroupsApi.md#DeleteResourceGroup) | **Delete** /resourcegroups/{resourcegroup} | Delete a resourcegroup.
[**GetResourceGroup**](ResourcegroupsApi.md#GetResourceGroup) | **Get** /resourcegroups/{resourcegroup} | Get the rate limiters specified for a resourcegroup.
[**GetResourceGroups**](ResourcegroupsApi.md#GetResourceGroups) | **Get** /resourcegroups | Get the list of all the resourcegroups.



## CreateOrUpdateResourceGroup

> CreateOrUpdateResourceGroup(ctx, resourcegroup).Body(body).Execute()

Creates a new resourcegroup with the specified rate limiters

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
    resourcegroup := "resourcegroup_example" // string | 
    body := *openapiclient.NewResourceGroup() // ResourceGroup | Rate limiters for the resourcegroup (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourcegroupsApi.CreateOrUpdateResourceGroup(context.Background(), resourcegroup).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcegroupsApi.CreateOrUpdateResourceGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourcegroup** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateResourceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ResourceGroup**](ResourceGroup.md) | Rate limiters for the resourcegroup | 

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


## DeleteResourceGroup

> DeleteResourceGroup(ctx, resourcegroup).Execute()

Delete a resourcegroup.

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
    resourcegroup := "resourcegroup_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourcegroupsApi.DeleteResourceGroup(context.Background(), resourcegroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcegroupsApi.DeleteResourceGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourcegroup** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteResourceGroupRequest struct via the builder pattern


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


## GetResourceGroup

> ResourceGroup GetResourceGroup(ctx, resourcegroup).Execute()

Get the rate limiters specified for a resourcegroup.

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
    resourcegroup := "resourcegroup_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourcegroupsApi.GetResourceGroup(context.Background(), resourcegroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcegroupsApi.GetResourceGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetResourceGroup`: ResourceGroup
    fmt.Fprintf(os.Stdout, "Response from `ResourcegroupsApi.GetResourceGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourcegroup** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResourceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResourceGroup**](ResourceGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResourceGroups

> []string GetResourceGroups(ctx).Execute()

Get the list of all the resourcegroups.

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
    resp, r, err := api_client.ResourcegroupsApi.GetResourceGroups(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcegroupsApi.GetResourceGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetResourceGroups`: []string
    fmt.Fprintf(os.Stdout, "Response from `ResourcegroupsApi.GetResourceGroups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetResourceGroupsRequest struct via the builder pattern


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

