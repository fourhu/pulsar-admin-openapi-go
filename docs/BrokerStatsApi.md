# \BrokerStatsApi

All URIs are relative to *http://localhost/admin/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllocatorStats**](BrokerStatsApi.md#GetAllocatorStats) | **Get** /broker-stats/allocator-stats/{allocator} | Get the stats for the Netty allocator. Available allocators are &#39;default&#39; and &#39;ml-cache&#39;
[**GetBrokerResourceAvailability**](BrokerStatsApi.md#GetBrokerResourceAvailability) | **Get** /broker-stats/broker-resource-availability/{tenant}/{namespace} | Broker availability report
[**GetLoadReport**](BrokerStatsApi.md#GetLoadReport) | **Get** /broker-stats/load-report | Get Load for this broker
[**GetMBeans**](BrokerStatsApi.md#GetMBeans) | **Get** /broker-stats/mbeans | Get all the mbean details of this broker JVM
[**GetMetrics**](BrokerStatsApi.md#GetMetrics) | **Get** /broker-stats/metrics | Gets the metrics for Monitoring
[**GetPendingBookieOpsStats**](BrokerStatsApi.md#GetPendingBookieOpsStats) | **Get** /broker-stats/bookieops | Get pending bookie client op stats by namesapce
[**GetTopics2**](BrokerStatsApi.md#GetTopics2) | **Get** /broker-stats/topics | Get all the topic stats by namespace



## GetAllocatorStats

> AllocatorStats GetAllocatorStats(ctx, allocator).Execute()

Get the stats for the Netty allocator. Available allocators are 'default' and 'ml-cache'

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
    allocator := "allocator_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BrokerStatsApi.GetAllocatorStats(context.Background(), allocator).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrokerStatsApi.GetAllocatorStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllocatorStats`: AllocatorStats
    fmt.Fprintf(os.Stdout, "Response from `BrokerStatsApi.GetAllocatorStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**allocator** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllocatorStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AllocatorStats**](AllocatorStats.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBrokerResourceAvailability

> map[string]ResourceUnit GetBrokerResourceAvailability(ctx, tenant, namespace).Execute()

Broker availability report



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BrokerStatsApi.GetBrokerResourceAvailability(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrokerStatsApi.GetBrokerResourceAvailability``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBrokerResourceAvailability`: map[string]ResourceUnit
    fmt.Fprintf(os.Stdout, "Response from `BrokerStatsApi.GetBrokerResourceAvailability`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBrokerResourceAvailabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**map[string]ResourceUnit**](ResourceUnit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoadReport

> LoadReport GetLoadReport(ctx).Execute()

Get Load for this broker



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
    resp, r, err := api_client.BrokerStatsApi.GetLoadReport(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrokerStatsApi.GetLoadReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLoadReport`: LoadReport
    fmt.Fprintf(os.Stdout, "Response from `BrokerStatsApi.GetLoadReport`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoadReportRequest struct via the builder pattern


### Return type

[**LoadReport**](LoadReport.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMBeans

> []Metrics GetMBeans(ctx).Execute()

Get all the mbean details of this broker JVM

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
    resp, r, err := api_client.BrokerStatsApi.GetMBeans(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrokerStatsApi.GetMBeans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMBeans`: []Metrics
    fmt.Fprintf(os.Stdout, "Response from `BrokerStatsApi.GetMBeans`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMBeansRequest struct via the builder pattern


### Return type

[**[]Metrics**](Metrics.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetrics

> []Metrics GetMetrics(ctx).Execute()

Gets the metrics for Monitoring



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
    resp, r, err := api_client.BrokerStatsApi.GetMetrics(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrokerStatsApi.GetMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMetrics`: []Metrics
    fmt.Fprintf(os.Stdout, "Response from `BrokerStatsApi.GetMetrics`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricsRequest struct via the builder pattern


### Return type

[**[]Metrics**](Metrics.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPendingBookieOpsStats

> map[string]PendingBookieOpsStats GetPendingBookieOpsStats(ctx).Execute()

Get pending bookie client op stats by namesapce

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
    resp, r, err := api_client.BrokerStatsApi.GetPendingBookieOpsStats(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrokerStatsApi.GetPendingBookieOpsStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPendingBookieOpsStats`: map[string]PendingBookieOpsStats
    fmt.Fprintf(os.Stdout, "Response from `BrokerStatsApi.GetPendingBookieOpsStats`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPendingBookieOpsStatsRequest struct via the builder pattern


### Return type

[**map[string]PendingBookieOpsStats**](PendingBookieOpsStats.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTopics2

> map[string]interface{} GetTopics2(ctx).Execute()

Get all the topic stats by namespace

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
    resp, r, err := api_client.BrokerStatsApi.GetTopics2(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrokerStatsApi.GetTopics2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTopics2`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BrokerStatsApi.GetTopics2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTopics2Request struct via the builder pattern


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

