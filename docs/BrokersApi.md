# \BrokersApi

All URIs are relative to *http://localhost/admin/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BacklogQuotaCheck**](BrokersApi.md#BacklogQuotaCheck) | **Get** /brokers/backlog-quota-check | An REST endpoint to trigger backlogQuotaCheck
[**DeleteDynamicConfiguration**](BrokersApi.md#DeleteDynamicConfiguration) | **Delete** /brokers/configuration/{configName} | Delete dynamic serviceconfiguration into zk only. This operation requires Pulsar super-user privileges.
[**GetActiveBrokers**](BrokersApi.md#GetActiveBrokers) | **Get** /brokers/{cluster} | Get the list of active brokers (web service addresses) in the cluster.If authorization is not enabled, any cluster name is valid.
[**GetAllDynamicConfigurations**](BrokersApi.md#GetAllDynamicConfigurations) | **Get** /brokers/configuration/values | Get value of all dynamic configurations&#39; value overridden on local config
[**GetDynamicConfigurationName**](BrokersApi.md#GetDynamicConfigurationName) | **Get** /brokers/configuration | Get all updatable dynamic configurations&#39;s name
[**GetInternalConfigurationData**](BrokersApi.md#GetInternalConfigurationData) | **Get** /brokers/internal-configuration | Get the internal configuration data
[**GetLeaderBroker**](BrokersApi.md#GetLeaderBroker) | **Get** /brokers/leaderBroker | Get the information of the leader broker.
[**GetOwnedNamespaces**](BrokersApi.md#GetOwnedNamespaces) | **Get** /brokers/{clusterName}/{broker-webserviceurl}/ownedNamespaces | Get the list of namespaces served by the specific broker
[**GetRuntimeConfiguration**](BrokersApi.md#GetRuntimeConfiguration) | **Get** /brokers/configuration/runtime | Get all runtime configurations. This operation requires Pulsar super-user privileges.
[**Healthcheck**](BrokersApi.md#Healthcheck) | **Get** /brokers/health | Run a healthcheck against the broker
[**IsReady**](BrokersApi.md#IsReady) | **Get** /brokers/ready | Check if the broker is fully initialized
[**UpdateDynamicConfiguration**](BrokersApi.md#UpdateDynamicConfiguration) | **Post** /brokers/configuration/{configName}/{configValue} | Update dynamic serviceconfiguration into zk only. This operation requires Pulsar super-user privileges.
[**Version**](BrokersApi.md#Version) | **Get** /brokers/version | Get version of current broker



## BacklogQuotaCheck

> BacklogQuotaCheck(ctx).Execute()

An REST endpoint to trigger backlogQuotaCheck

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
    resp, r, err := api_client.BrokersApi.BacklogQuotaCheck(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrokersApi.BacklogQuotaCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiBacklogQuotaCheckRequest struct via the builder pattern


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


## DeleteDynamicConfiguration

> DeleteDynamicConfiguration(ctx, configName).Execute()

Delete dynamic serviceconfiguration into zk only. This operation requires Pulsar super-user privileges.

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
    configName := "configName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BrokersApi.DeleteDynamicConfiguration(context.Background(), configName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrokersApi.DeleteDynamicConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDynamicConfigurationRequest struct via the builder pattern


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


## GetActiveBrokers

> []string GetActiveBrokers(ctx, cluster).Execute()

Get the list of active brokers (web service addresses) in the cluster.If authorization is not enabled, any cluster name is valid.

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
    cluster := "cluster_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BrokersApi.GetActiveBrokers(context.Background(), cluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrokersApi.GetActiveBrokers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetActiveBrokers`: []string
    fmt.Fprintf(os.Stdout, "Response from `BrokersApi.GetActiveBrokers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetActiveBrokersRequest struct via the builder pattern


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


## GetAllDynamicConfigurations

> map[string]string GetAllDynamicConfigurations(ctx).Execute()

Get value of all dynamic configurations' value overridden on local config

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
    resp, r, err := api_client.BrokersApi.GetAllDynamicConfigurations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrokersApi.GetAllDynamicConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllDynamicConfigurations`: map[string]string
    fmt.Fprintf(os.Stdout, "Response from `BrokersApi.GetAllDynamicConfigurations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllDynamicConfigurationsRequest struct via the builder pattern


### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDynamicConfigurationName

> []string GetDynamicConfigurationName(ctx).Execute()

Get all updatable dynamic configurations's name

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
    resp, r, err := api_client.BrokersApi.GetDynamicConfigurationName(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrokersApi.GetDynamicConfigurationName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDynamicConfigurationName`: []string
    fmt.Fprintf(os.Stdout, "Response from `BrokersApi.GetDynamicConfigurationName`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDynamicConfigurationNameRequest struct via the builder pattern


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


## GetInternalConfigurationData

> InternalConfigurationData GetInternalConfigurationData(ctx).Execute()

Get the internal configuration data

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
    resp, r, err := api_client.BrokersApi.GetInternalConfigurationData(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrokersApi.GetInternalConfigurationData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInternalConfigurationData`: InternalConfigurationData
    fmt.Fprintf(os.Stdout, "Response from `BrokersApi.GetInternalConfigurationData`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetInternalConfigurationDataRequest struct via the builder pattern


### Return type

[**InternalConfigurationData**](InternalConfigurationData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLeaderBroker

> BrokerInfo GetLeaderBroker(ctx).Execute()

Get the information of the leader broker.

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
    resp, r, err := api_client.BrokersApi.GetLeaderBroker(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrokersApi.GetLeaderBroker``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLeaderBroker`: BrokerInfo
    fmt.Fprintf(os.Stdout, "Response from `BrokersApi.GetLeaderBroker`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLeaderBrokerRequest struct via the builder pattern


### Return type

[**BrokerInfo**](BrokerInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOwnedNamespaces

> map[string]NamespaceOwnershipStatus GetOwnedNamespaces(ctx, clusterName, brokerWebserviceurl).Execute()

Get the list of namespaces served by the specific broker

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
    clusterName := "clusterName_example" // string | 
    brokerWebserviceurl := "brokerWebserviceurl_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BrokersApi.GetOwnedNamespaces(context.Background(), clusterName, brokerWebserviceurl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrokersApi.GetOwnedNamespaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOwnedNamespaces`: map[string]NamespaceOwnershipStatus
    fmt.Fprintf(os.Stdout, "Response from `BrokersApi.GetOwnedNamespaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** |  | 
**brokerWebserviceurl** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOwnedNamespacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**map[string]NamespaceOwnershipStatus**](NamespaceOwnershipStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRuntimeConfiguration

> map[string]string GetRuntimeConfiguration(ctx).Execute()

Get all runtime configurations. This operation requires Pulsar super-user privileges.

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
    resp, r, err := api_client.BrokersApi.GetRuntimeConfiguration(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrokersApi.GetRuntimeConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRuntimeConfiguration`: map[string]string
    fmt.Fprintf(os.Stdout, "Response from `BrokersApi.GetRuntimeConfiguration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuntimeConfigurationRequest struct via the builder pattern


### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Healthcheck

> Healthcheck(ctx).Execute()

Run a healthcheck against the broker

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
    resp, r, err := api_client.BrokersApi.Healthcheck(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrokersApi.Healthcheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHealthcheckRequest struct via the builder pattern


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


## IsReady

> IsReady(ctx).Execute()

Check if the broker is fully initialized

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
    resp, r, err := api_client.BrokersApi.IsReady(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrokersApi.IsReady``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIsReadyRequest struct via the builder pattern


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


## UpdateDynamicConfiguration

> UpdateDynamicConfiguration(ctx, configName, configValue).Execute()

Update dynamic serviceconfiguration into zk only. This operation requires Pulsar super-user privileges.

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
    configName := "configName_example" // string | 
    configValue := "configValue_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BrokersApi.UpdateDynamicConfiguration(context.Background(), configName, configValue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrokersApi.UpdateDynamicConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configName** | **string** |  | 
**configValue** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDynamicConfigurationRequest struct via the builder pattern


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


## Version

> string Version(ctx).Execute()

Get version of current broker

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
    resp, r, err := api_client.BrokersApi.Version(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrokersApi.Version``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Version`: string
    fmt.Fprintf(os.Stdout, "Response from `BrokersApi.Version`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiVersionRequest struct via the builder pattern


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

