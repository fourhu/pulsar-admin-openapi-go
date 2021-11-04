# \BookiesApi

All URIs are relative to *http://localhost/admin/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteBookieRackInfo**](BookiesApi.md#DeleteBookieRackInfo) | **Delete** /bookies/racks-info/{bookie} | Removed the rack placement information for a specific bookie in the cluster
[**GetBookieRackInfo**](BookiesApi.md#GetBookieRackInfo) | **Get** /bookies/racks-info/{bookie} | Gets the rack placement information for a specific bookie in the cluster
[**GetBookiesRackInfo**](BookiesApi.md#GetBookiesRackInfo) | **Get** /bookies/racks-info | Gets the rack placement information for all the bookies in the cluster
[**UpdateBookieRackInfo**](BookiesApi.md#UpdateBookieRackInfo) | **Post** /bookies/racks-info/{bookie} | Updates the rack placement information for a specific bookie in the cluster (note. bookie address format:&#x60;address:port&#x60;)



## DeleteBookieRackInfo

> DeleteBookieRackInfo(ctx, bookie).Execute()

Removed the rack placement information for a specific bookie in the cluster

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
    bookie := "bookie_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BookiesApi.DeleteBookieRackInfo(context.Background(), bookie).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookiesApi.DeleteBookieRackInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bookie** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBookieRackInfoRequest struct via the builder pattern


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


## GetBookieRackInfo

> BookieInfo GetBookieRackInfo(ctx, bookie).Execute()

Gets the rack placement information for a specific bookie in the cluster

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
    bookie := "bookie_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BookiesApi.GetBookieRackInfo(context.Background(), bookie).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookiesApi.GetBookieRackInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBookieRackInfo`: BookieInfo
    fmt.Fprintf(os.Stdout, "Response from `BookiesApi.GetBookieRackInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bookie** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBookieRackInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BookieInfo**](BookieInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBookiesRackInfo

> map[string]map[string]BookieInfo GetBookiesRackInfo(ctx).Execute()

Gets the rack placement information for all the bookies in the cluster

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
    resp, r, err := api_client.BookiesApi.GetBookiesRackInfo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookiesApi.GetBookiesRackInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBookiesRackInfo`: map[string]map[string]BookieInfo
    fmt.Fprintf(os.Stdout, "Response from `BookiesApi.GetBookiesRackInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBookiesRackInfoRequest struct via the builder pattern


### Return type

[**map[string]map[string]BookieInfo**](map.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBookieRackInfo

> UpdateBookieRackInfo(ctx, bookie).Group(group).Execute()

Updates the rack placement information for a specific bookie in the cluster (note. bookie address format:`address:port`)

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
    bookie := "bookie_example" // string | 
    group := "group_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BookiesApi.UpdateBookieRackInfo(context.Background(), bookie).Group(group).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookiesApi.UpdateBookieRackInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bookie** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBookieRackInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **group** | **string** |  | 

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

