# \ClustersApi

All URIs are relative to *http://localhost/admin/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCluster**](ClustersApi.md#CreateCluster) | **Put** /clusters/{cluster} | Create a new cluster.
[**DeleteCluster**](ClustersApi.md#DeleteCluster) | **Delete** /clusters/{cluster} | Delete an existing cluster.
[**DeleteFailureDomain**](ClustersApi.md#DeleteFailureDomain) | **Delete** /clusters/{cluster}/failureDomains/{domainName} | Delete the failure domain of the cluster
[**DeleteNamespaceIsolationPolicy**](ClustersApi.md#DeleteNamespaceIsolationPolicy) | **Delete** /clusters/{cluster}/namespaceIsolationPolicies/{policyName} | Delete namespace isolation policy.
[**GetBrokerWithNamespaceIsolationPolicy**](ClustersApi.md#GetBrokerWithNamespaceIsolationPolicy) | **Get** /clusters/{cluster}/namespaceIsolationPolicies/brokers/{broker} | Get a broker with namespace-isolation policies attached to it.
[**GetBrokersWithNamespaceIsolationPolicy**](ClustersApi.md#GetBrokersWithNamespaceIsolationPolicy) | **Get** /clusters/{cluster}/namespaceIsolationPolicies/brokers | Get list of brokers with namespace-isolation policies attached to them.
[**GetCluster**](ClustersApi.md#GetCluster) | **Get** /clusters/{cluster} | Get the configuration for the specified cluster.
[**GetClusters**](ClustersApi.md#GetClusters) | **Get** /clusters | Get the list of all the Pulsar clusters.
[**GetDomain**](ClustersApi.md#GetDomain) | **Get** /clusters/{cluster}/failureDomains/{domainName} | Get a domain in a cluster
[**GetFailureDomains**](ClustersApi.md#GetFailureDomains) | **Get** /clusters/{cluster}/failureDomains | Get the cluster failure domains.
[**GetNamespaceIsolationPolicies**](ClustersApi.md#GetNamespaceIsolationPolicies) | **Get** /clusters/{cluster}/namespaceIsolationPolicies | Get the namespace isolation policies assigned to the cluster.
[**GetNamespaceIsolationPolicy**](ClustersApi.md#GetNamespaceIsolationPolicy) | **Get** /clusters/{cluster}/namespaceIsolationPolicies/{policyName} | Get the single namespace isolation policy assigned to the cluster.
[**GetPeerCluster**](ClustersApi.md#GetPeerCluster) | **Get** /clusters/{cluster}/peers | Get the peer-cluster data for the specified cluster.
[**SetFailureDomain**](ClustersApi.md#SetFailureDomain) | **Post** /clusters/{cluster}/failureDomains/{domainName} | Set the failure domain of the cluster.
[**SetNamespaceIsolationPolicy**](ClustersApi.md#SetNamespaceIsolationPolicy) | **Post** /clusters/{cluster}/namespaceIsolationPolicies/{policyName} | Set namespace isolation policy.
[**SetPeerClusterNames**](ClustersApi.md#SetPeerClusterNames) | **Post** /clusters/{cluster}/peers | Update peer-cluster-list for a cluster.
[**UpdateCluster**](ClustersApi.md#UpdateCluster) | **Post** /clusters/{cluster} | Update the configuration for a cluster.



## CreateCluster

> CreateCluster(ctx, cluster).Body(body).Execute()

Create a new cluster.



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
    cluster := "cluster_example" // string | The cluster name
    body := *openapiclient.NewClusterData() // ClusterData | The cluster data

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.CreateCluster(context.Background(), cluster).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.CreateCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | **string** | The cluster name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ClusterData**](ClusterData.md) | The cluster data | 

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


## DeleteCluster

> DeleteCluster(ctx, cluster).Execute()

Delete an existing cluster.



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
    cluster := "cluster_example" // string | The cluster name

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.DeleteCluster(context.Background(), cluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.DeleteCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | **string** | The cluster name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClusterRequest struct via the builder pattern


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


## DeleteFailureDomain

> DeleteFailureDomain(ctx, cluster, domainName).Execute()

Delete the failure domain of the cluster



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
    cluster := "cluster_example" // string | The cluster name
    domainName := "domainName_example" // string | The failure domain name

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.DeleteFailureDomain(context.Background(), cluster, domainName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.DeleteFailureDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | **string** | The cluster name | 
**domainName** | **string** | The failure domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFailureDomainRequest struct via the builder pattern


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


## DeleteNamespaceIsolationPolicy

> DeleteNamespaceIsolationPolicy(ctx, cluster, policyName).Execute()

Delete namespace isolation policy.



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
    cluster := "cluster_example" // string | The cluster name
    policyName := "policyName_example" // string | The namespace isolation policy name

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.DeleteNamespaceIsolationPolicy(context.Background(), cluster, policyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.DeleteNamespaceIsolationPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | **string** | The cluster name | 
**policyName** | **string** | The namespace isolation policy name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNamespaceIsolationPolicyRequest struct via the builder pattern


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


## GetBrokerWithNamespaceIsolationPolicy

> BrokerNamespaceIsolationData GetBrokerWithNamespaceIsolationPolicy(ctx, cluster, broker).Execute()

Get a broker with namespace-isolation policies attached to it.



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
    cluster := "cluster_example" // string | The cluster name
    broker := "broker1:8080" // string | The broker name (<broker-hostname>:<web-service-port>)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.GetBrokerWithNamespaceIsolationPolicy(context.Background(), cluster, broker).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.GetBrokerWithNamespaceIsolationPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBrokerWithNamespaceIsolationPolicy`: BrokerNamespaceIsolationData
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.GetBrokerWithNamespaceIsolationPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | **string** | The cluster name | 
**broker** | **string** | The broker name (&lt;broker-hostname&gt;:&lt;web-service-port&gt;) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBrokerWithNamespaceIsolationPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BrokerNamespaceIsolationData**](BrokerNamespaceIsolationData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBrokersWithNamespaceIsolationPolicy

> []BrokerNamespaceIsolationData GetBrokersWithNamespaceIsolationPolicy(ctx, cluster).Execute()

Get list of brokers with namespace-isolation policies attached to them.



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
    cluster := "cluster_example" // string | The cluster name

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.GetBrokersWithNamespaceIsolationPolicy(context.Background(), cluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.GetBrokersWithNamespaceIsolationPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBrokersWithNamespaceIsolationPolicy`: []BrokerNamespaceIsolationData
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.GetBrokersWithNamespaceIsolationPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | **string** | The cluster name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBrokersWithNamespaceIsolationPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]BrokerNamespaceIsolationData**](BrokerNamespaceIsolationData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCluster

> ClusterData GetCluster(ctx, cluster).Execute()

Get the configuration for the specified cluster.



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
    cluster := "cluster_example" // string | The cluster name

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.GetCluster(context.Background(), cluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.GetCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCluster`: ClusterData
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.GetCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | **string** | The cluster name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClusterData**](ClusterData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusters

> []string GetClusters(ctx).Execute()

Get the list of all the Pulsar clusters.

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
    resp, r, err := api_client.ClustersApi.GetClusters(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.GetClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusters`: []string
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.GetClusters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersRequest struct via the builder pattern


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


## GetDomain

> FailureDomain GetDomain(ctx, cluster, domainName).Execute()

Get a domain in a cluster



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
    cluster := "cluster_example" // string | The cluster name
    domainName := "domainName_example" // string | The failure domain name

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.GetDomain(context.Background(), cluster, domainName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.GetDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDomain`: FailureDomain
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.GetDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | **string** | The cluster name | 
**domainName** | **string** | The failure domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FailureDomain**](FailureDomain.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFailureDomains

> map[string]FailureDomain GetFailureDomains(ctx, cluster).Execute()

Get the cluster failure domains.



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
    cluster := "cluster_example" // string | The cluster name

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.GetFailureDomains(context.Background(), cluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.GetFailureDomains``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFailureDomains`: map[string]FailureDomain
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.GetFailureDomains`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | **string** | The cluster name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFailureDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**map[string]FailureDomain**](FailureDomain.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamespaceIsolationPolicies

> map[string]NamespaceIsolationData GetNamespaceIsolationPolicies(ctx, cluster).Execute()

Get the namespace isolation policies assigned to the cluster.



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
    cluster := "cluster_example" // string | The cluster name

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.GetNamespaceIsolationPolicies(context.Background(), cluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.GetNamespaceIsolationPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNamespaceIsolationPolicies`: map[string]NamespaceIsolationData
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.GetNamespaceIsolationPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | **string** | The cluster name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamespaceIsolationPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**map[string]NamespaceIsolationData**](NamespaceIsolationData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamespaceIsolationPolicy

> NamespaceIsolationData GetNamespaceIsolationPolicy(ctx, cluster, policyName).Execute()

Get the single namespace isolation policy assigned to the cluster.



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
    cluster := "cluster_example" // string | The cluster name
    policyName := "policyName_example" // string | The name of the namespace isolation policy

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.GetNamespaceIsolationPolicy(context.Background(), cluster, policyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.GetNamespaceIsolationPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNamespaceIsolationPolicy`: NamespaceIsolationData
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.GetNamespaceIsolationPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | **string** | The cluster name | 
**policyName** | **string** | The name of the namespace isolation policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamespaceIsolationPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NamespaceIsolationData**](NamespaceIsolationData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPeerCluster

> []string GetPeerCluster(ctx, cluster).Execute()

Get the peer-cluster data for the specified cluster.



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
    cluster := "cluster_example" // string | The cluster name

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.GetPeerCluster(context.Background(), cluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.GetPeerCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPeerCluster`: []string
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.GetPeerCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | **string** | The cluster name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPeerClusterRequest struct via the builder pattern


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


## SetFailureDomain

> SetFailureDomain(ctx, cluster, domainName).Body(body).Execute()

Set the failure domain of the cluster.



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
    cluster := "cluster_example" // string | The cluster name
    domainName := "domainName_example" // string | The failure domain name
    body := *openapiclient.NewFailureDomain() // FailureDomain | The configuration data of a failure domain

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.SetFailureDomain(context.Background(), cluster, domainName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.SetFailureDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | **string** | The cluster name | 
**domainName** | **string** | The failure domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetFailureDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**FailureDomain**](FailureDomain.md) | The configuration data of a failure domain | 

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


## SetNamespaceIsolationPolicy

> SetNamespaceIsolationPolicy(ctx, cluster, policyName).Body(body).Execute()

Set namespace isolation policy.



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
    cluster := "cluster_example" // string | The cluster name
    policyName := "policyName_example" // string | The namespace isolation policy name
    body := *openapiclient.NewNamespaceIsolationData() // NamespaceIsolationData | The namespace isolation policy data

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.SetNamespaceIsolationPolicy(context.Background(), cluster, policyName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.SetNamespaceIsolationPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | **string** | The cluster name | 
**policyName** | **string** | The namespace isolation policy name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetNamespaceIsolationPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**NamespaceIsolationData**](NamespaceIsolationData.md) | The namespace isolation policy data | 

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


## SetPeerClusterNames

> SetPeerClusterNames(ctx, cluster).Body(body).Execute()

Update peer-cluster-list for a cluster.



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
    cluster := "cluster_example" // string | The cluster name
    body := []string{"Property_example"} // []string | The list of peer cluster names

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.SetPeerClusterNames(context.Background(), cluster).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.SetPeerClusterNames``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | **string** | The cluster name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetPeerClusterNamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **[]string** | The list of peer cluster names | 

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


## UpdateCluster

> UpdateCluster(ctx, cluster).Body(body).Execute()

Update the configuration for a cluster.



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
    cluster := "cluster_example" // string | The cluster name
    body := *openapiclient.NewClusterData() // ClusterData | The cluster data

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.UpdateCluster(context.Background(), cluster).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.UpdateCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | **string** | The cluster name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ClusterData**](ClusterData.md) | The cluster data | 

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

