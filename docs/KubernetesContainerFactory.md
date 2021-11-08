# KubernetesContainerFactory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**K8Uri** | Pointer to **string** |  | [optional] 
**JobNamespace** | Pointer to **string** |  | [optional] 
**JobName** | Pointer to **string** |  | [optional] 
**PulsarDockerImageName** | Pointer to **string** |  | [optional] 
**FunctionDockerImages** | Pointer to **map[string]string** |  | [optional] 
**ImagePullPolicy** | Pointer to **string** |  | [optional] 
**PulsarRootDir** | Pointer to **string** |  | [optional] 
**ConfigAdminCLI** | Pointer to **string** |  | [optional] 
**SubmittingInsidePod** | Pointer to **bool** |  | [optional] 
**PulsarServiceUrl** | Pointer to **string** |  | [optional] 
**PulsarAdminUrl** | Pointer to **string** |  | [optional] 
**InstallUserCodeDependencies** | Pointer to **bool** |  | [optional] 
**PythonDependencyRepository** | Pointer to **string** |  | [optional] 
**PythonExtraDependencyRepository** | Pointer to **string** |  | [optional] 
**ExtraFunctionDependenciesDir** | Pointer to **string** |  | [optional] 
**CustomLabels** | Pointer to **map[string]string** |  | [optional] 
**ExpectedMetricsCollectionInterval** | Pointer to **int32** |  | [optional] 
**ChangeConfigMap** | Pointer to **string** |  | [optional] 
**ChangeConfigMapNamespace** | Pointer to **string** |  | [optional] 
**PercentMemoryPadding** | Pointer to **int32** |  | [optional] 
**CpuOverCommitRatio** | Pointer to **float64** |  | [optional] 
**MemoryOverCommitRatio** | Pointer to **float64** |  | [optional] 
**GrpcPort** | Pointer to **int32** |  | [optional] 
**MetricsPort** | Pointer to **int32** |  | [optional] 
**NarExtractionDirectory** | Pointer to **string** |  | [optional] 
**FunctionInstanceClassPath** | Pointer to **string** |  | [optional] 

## Methods

### NewKubernetesContainerFactory

`func NewKubernetesContainerFactory() *KubernetesContainerFactory`

NewKubernetesContainerFactory instantiates a new KubernetesContainerFactory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesContainerFactoryWithDefaults

`func NewKubernetesContainerFactoryWithDefaults() *KubernetesContainerFactory`

NewKubernetesContainerFactoryWithDefaults instantiates a new KubernetesContainerFactory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetK8Uri

`func (o *KubernetesContainerFactory) GetK8Uri() string`

GetK8Uri returns the K8Uri field if non-nil, zero value otherwise.

### GetK8UriOk

`func (o *KubernetesContainerFactory) GetK8UriOk() (*string, bool)`

GetK8UriOk returns a tuple with the K8Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8Uri

`func (o *KubernetesContainerFactory) SetK8Uri(v string)`

SetK8Uri sets K8Uri field to given value.

### HasK8Uri

`func (o *KubernetesContainerFactory) HasK8Uri() bool`

HasK8Uri returns a boolean if a field has been set.

### GetJobNamespace

`func (o *KubernetesContainerFactory) GetJobNamespace() string`

GetJobNamespace returns the JobNamespace field if non-nil, zero value otherwise.

### GetJobNamespaceOk

`func (o *KubernetesContainerFactory) GetJobNamespaceOk() (*string, bool)`

GetJobNamespaceOk returns a tuple with the JobNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobNamespace

`func (o *KubernetesContainerFactory) SetJobNamespace(v string)`

SetJobNamespace sets JobNamespace field to given value.

### HasJobNamespace

`func (o *KubernetesContainerFactory) HasJobNamespace() bool`

HasJobNamespace returns a boolean if a field has been set.

### GetJobName

`func (o *KubernetesContainerFactory) GetJobName() string`

GetJobName returns the JobName field if non-nil, zero value otherwise.

### GetJobNameOk

`func (o *KubernetesContainerFactory) GetJobNameOk() (*string, bool)`

GetJobNameOk returns a tuple with the JobName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobName

`func (o *KubernetesContainerFactory) SetJobName(v string)`

SetJobName sets JobName field to given value.

### HasJobName

`func (o *KubernetesContainerFactory) HasJobName() bool`

HasJobName returns a boolean if a field has been set.

### GetPulsarDockerImageName

`func (o *KubernetesContainerFactory) GetPulsarDockerImageName() string`

GetPulsarDockerImageName returns the PulsarDockerImageName field if non-nil, zero value otherwise.

### GetPulsarDockerImageNameOk

`func (o *KubernetesContainerFactory) GetPulsarDockerImageNameOk() (*string, bool)`

GetPulsarDockerImageNameOk returns a tuple with the PulsarDockerImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulsarDockerImageName

`func (o *KubernetesContainerFactory) SetPulsarDockerImageName(v string)`

SetPulsarDockerImageName sets PulsarDockerImageName field to given value.

### HasPulsarDockerImageName

`func (o *KubernetesContainerFactory) HasPulsarDockerImageName() bool`

HasPulsarDockerImageName returns a boolean if a field has been set.

### GetFunctionDockerImages

`func (o *KubernetesContainerFactory) GetFunctionDockerImages() map[string]string`

GetFunctionDockerImages returns the FunctionDockerImages field if non-nil, zero value otherwise.

### GetFunctionDockerImagesOk

`func (o *KubernetesContainerFactory) GetFunctionDockerImagesOk() (*map[string]string, bool)`

GetFunctionDockerImagesOk returns a tuple with the FunctionDockerImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionDockerImages

`func (o *KubernetesContainerFactory) SetFunctionDockerImages(v map[string]string)`

SetFunctionDockerImages sets FunctionDockerImages field to given value.

### HasFunctionDockerImages

`func (o *KubernetesContainerFactory) HasFunctionDockerImages() bool`

HasFunctionDockerImages returns a boolean if a field has been set.

### GetImagePullPolicy

`func (o *KubernetesContainerFactory) GetImagePullPolicy() string`

GetImagePullPolicy returns the ImagePullPolicy field if non-nil, zero value otherwise.

### GetImagePullPolicyOk

`func (o *KubernetesContainerFactory) GetImagePullPolicyOk() (*string, bool)`

GetImagePullPolicyOk returns a tuple with the ImagePullPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagePullPolicy

`func (o *KubernetesContainerFactory) SetImagePullPolicy(v string)`

SetImagePullPolicy sets ImagePullPolicy field to given value.

### HasImagePullPolicy

`func (o *KubernetesContainerFactory) HasImagePullPolicy() bool`

HasImagePullPolicy returns a boolean if a field has been set.

### GetPulsarRootDir

`func (o *KubernetesContainerFactory) GetPulsarRootDir() string`

GetPulsarRootDir returns the PulsarRootDir field if non-nil, zero value otherwise.

### GetPulsarRootDirOk

`func (o *KubernetesContainerFactory) GetPulsarRootDirOk() (*string, bool)`

GetPulsarRootDirOk returns a tuple with the PulsarRootDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulsarRootDir

`func (o *KubernetesContainerFactory) SetPulsarRootDir(v string)`

SetPulsarRootDir sets PulsarRootDir field to given value.

### HasPulsarRootDir

`func (o *KubernetesContainerFactory) HasPulsarRootDir() bool`

HasPulsarRootDir returns a boolean if a field has been set.

### GetConfigAdminCLI

`func (o *KubernetesContainerFactory) GetConfigAdminCLI() string`

GetConfigAdminCLI returns the ConfigAdminCLI field if non-nil, zero value otherwise.

### GetConfigAdminCLIOk

`func (o *KubernetesContainerFactory) GetConfigAdminCLIOk() (*string, bool)`

GetConfigAdminCLIOk returns a tuple with the ConfigAdminCLI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigAdminCLI

`func (o *KubernetesContainerFactory) SetConfigAdminCLI(v string)`

SetConfigAdminCLI sets ConfigAdminCLI field to given value.

### HasConfigAdminCLI

`func (o *KubernetesContainerFactory) HasConfigAdminCLI() bool`

HasConfigAdminCLI returns a boolean if a field has been set.

### GetSubmittingInsidePod

`func (o *KubernetesContainerFactory) GetSubmittingInsidePod() bool`

GetSubmittingInsidePod returns the SubmittingInsidePod field if non-nil, zero value otherwise.

### GetSubmittingInsidePodOk

`func (o *KubernetesContainerFactory) GetSubmittingInsidePodOk() (*bool, bool)`

GetSubmittingInsidePodOk returns a tuple with the SubmittingInsidePod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittingInsidePod

`func (o *KubernetesContainerFactory) SetSubmittingInsidePod(v bool)`

SetSubmittingInsidePod sets SubmittingInsidePod field to given value.

### HasSubmittingInsidePod

`func (o *KubernetesContainerFactory) HasSubmittingInsidePod() bool`

HasSubmittingInsidePod returns a boolean if a field has been set.

### GetPulsarServiceUrl

`func (o *KubernetesContainerFactory) GetPulsarServiceUrl() string`

GetPulsarServiceUrl returns the PulsarServiceUrl field if non-nil, zero value otherwise.

### GetPulsarServiceUrlOk

`func (o *KubernetesContainerFactory) GetPulsarServiceUrlOk() (*string, bool)`

GetPulsarServiceUrlOk returns a tuple with the PulsarServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulsarServiceUrl

`func (o *KubernetesContainerFactory) SetPulsarServiceUrl(v string)`

SetPulsarServiceUrl sets PulsarServiceUrl field to given value.

### HasPulsarServiceUrl

`func (o *KubernetesContainerFactory) HasPulsarServiceUrl() bool`

HasPulsarServiceUrl returns a boolean if a field has been set.

### GetPulsarAdminUrl

`func (o *KubernetesContainerFactory) GetPulsarAdminUrl() string`

GetPulsarAdminUrl returns the PulsarAdminUrl field if non-nil, zero value otherwise.

### GetPulsarAdminUrlOk

`func (o *KubernetesContainerFactory) GetPulsarAdminUrlOk() (*string, bool)`

GetPulsarAdminUrlOk returns a tuple with the PulsarAdminUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulsarAdminUrl

`func (o *KubernetesContainerFactory) SetPulsarAdminUrl(v string)`

SetPulsarAdminUrl sets PulsarAdminUrl field to given value.

### HasPulsarAdminUrl

`func (o *KubernetesContainerFactory) HasPulsarAdminUrl() bool`

HasPulsarAdminUrl returns a boolean if a field has been set.

### GetInstallUserCodeDependencies

`func (o *KubernetesContainerFactory) GetInstallUserCodeDependencies() bool`

GetInstallUserCodeDependencies returns the InstallUserCodeDependencies field if non-nil, zero value otherwise.

### GetInstallUserCodeDependenciesOk

`func (o *KubernetesContainerFactory) GetInstallUserCodeDependenciesOk() (*bool, bool)`

GetInstallUserCodeDependenciesOk returns a tuple with the InstallUserCodeDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallUserCodeDependencies

`func (o *KubernetesContainerFactory) SetInstallUserCodeDependencies(v bool)`

SetInstallUserCodeDependencies sets InstallUserCodeDependencies field to given value.

### HasInstallUserCodeDependencies

`func (o *KubernetesContainerFactory) HasInstallUserCodeDependencies() bool`

HasInstallUserCodeDependencies returns a boolean if a field has been set.

### GetPythonDependencyRepository

`func (o *KubernetesContainerFactory) GetPythonDependencyRepository() string`

GetPythonDependencyRepository returns the PythonDependencyRepository field if non-nil, zero value otherwise.

### GetPythonDependencyRepositoryOk

`func (o *KubernetesContainerFactory) GetPythonDependencyRepositoryOk() (*string, bool)`

GetPythonDependencyRepositoryOk returns a tuple with the PythonDependencyRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPythonDependencyRepository

`func (o *KubernetesContainerFactory) SetPythonDependencyRepository(v string)`

SetPythonDependencyRepository sets PythonDependencyRepository field to given value.

### HasPythonDependencyRepository

`func (o *KubernetesContainerFactory) HasPythonDependencyRepository() bool`

HasPythonDependencyRepository returns a boolean if a field has been set.

### GetPythonExtraDependencyRepository

`func (o *KubernetesContainerFactory) GetPythonExtraDependencyRepository() string`

GetPythonExtraDependencyRepository returns the PythonExtraDependencyRepository field if non-nil, zero value otherwise.

### GetPythonExtraDependencyRepositoryOk

`func (o *KubernetesContainerFactory) GetPythonExtraDependencyRepositoryOk() (*string, bool)`

GetPythonExtraDependencyRepositoryOk returns a tuple with the PythonExtraDependencyRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPythonExtraDependencyRepository

`func (o *KubernetesContainerFactory) SetPythonExtraDependencyRepository(v string)`

SetPythonExtraDependencyRepository sets PythonExtraDependencyRepository field to given value.

### HasPythonExtraDependencyRepository

`func (o *KubernetesContainerFactory) HasPythonExtraDependencyRepository() bool`

HasPythonExtraDependencyRepository returns a boolean if a field has been set.

### GetExtraFunctionDependenciesDir

`func (o *KubernetesContainerFactory) GetExtraFunctionDependenciesDir() string`

GetExtraFunctionDependenciesDir returns the ExtraFunctionDependenciesDir field if non-nil, zero value otherwise.

### GetExtraFunctionDependenciesDirOk

`func (o *KubernetesContainerFactory) GetExtraFunctionDependenciesDirOk() (*string, bool)`

GetExtraFunctionDependenciesDirOk returns a tuple with the ExtraFunctionDependenciesDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraFunctionDependenciesDir

`func (o *KubernetesContainerFactory) SetExtraFunctionDependenciesDir(v string)`

SetExtraFunctionDependenciesDir sets ExtraFunctionDependenciesDir field to given value.

### HasExtraFunctionDependenciesDir

`func (o *KubernetesContainerFactory) HasExtraFunctionDependenciesDir() bool`

HasExtraFunctionDependenciesDir returns a boolean if a field has been set.

### GetCustomLabels

`func (o *KubernetesContainerFactory) GetCustomLabels() map[string]string`

GetCustomLabels returns the CustomLabels field if non-nil, zero value otherwise.

### GetCustomLabelsOk

`func (o *KubernetesContainerFactory) GetCustomLabelsOk() (*map[string]string, bool)`

GetCustomLabelsOk returns a tuple with the CustomLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomLabels

`func (o *KubernetesContainerFactory) SetCustomLabels(v map[string]string)`

SetCustomLabels sets CustomLabels field to given value.

### HasCustomLabels

`func (o *KubernetesContainerFactory) HasCustomLabels() bool`

HasCustomLabels returns a boolean if a field has been set.

### GetExpectedMetricsCollectionInterval

`func (o *KubernetesContainerFactory) GetExpectedMetricsCollectionInterval() int32`

GetExpectedMetricsCollectionInterval returns the ExpectedMetricsCollectionInterval field if non-nil, zero value otherwise.

### GetExpectedMetricsCollectionIntervalOk

`func (o *KubernetesContainerFactory) GetExpectedMetricsCollectionIntervalOk() (*int32, bool)`

GetExpectedMetricsCollectionIntervalOk returns a tuple with the ExpectedMetricsCollectionInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedMetricsCollectionInterval

`func (o *KubernetesContainerFactory) SetExpectedMetricsCollectionInterval(v int32)`

SetExpectedMetricsCollectionInterval sets ExpectedMetricsCollectionInterval field to given value.

### HasExpectedMetricsCollectionInterval

`func (o *KubernetesContainerFactory) HasExpectedMetricsCollectionInterval() bool`

HasExpectedMetricsCollectionInterval returns a boolean if a field has been set.

### GetChangeConfigMap

`func (o *KubernetesContainerFactory) GetChangeConfigMap() string`

GetChangeConfigMap returns the ChangeConfigMap field if non-nil, zero value otherwise.

### GetChangeConfigMapOk

`func (o *KubernetesContainerFactory) GetChangeConfigMapOk() (*string, bool)`

GetChangeConfigMapOk returns a tuple with the ChangeConfigMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeConfigMap

`func (o *KubernetesContainerFactory) SetChangeConfigMap(v string)`

SetChangeConfigMap sets ChangeConfigMap field to given value.

### HasChangeConfigMap

`func (o *KubernetesContainerFactory) HasChangeConfigMap() bool`

HasChangeConfigMap returns a boolean if a field has been set.

### GetChangeConfigMapNamespace

`func (o *KubernetesContainerFactory) GetChangeConfigMapNamespace() string`

GetChangeConfigMapNamespace returns the ChangeConfigMapNamespace field if non-nil, zero value otherwise.

### GetChangeConfigMapNamespaceOk

`func (o *KubernetesContainerFactory) GetChangeConfigMapNamespaceOk() (*string, bool)`

GetChangeConfigMapNamespaceOk returns a tuple with the ChangeConfigMapNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeConfigMapNamespace

`func (o *KubernetesContainerFactory) SetChangeConfigMapNamespace(v string)`

SetChangeConfigMapNamespace sets ChangeConfigMapNamespace field to given value.

### HasChangeConfigMapNamespace

`func (o *KubernetesContainerFactory) HasChangeConfigMapNamespace() bool`

HasChangeConfigMapNamespace returns a boolean if a field has been set.

### GetPercentMemoryPadding

`func (o *KubernetesContainerFactory) GetPercentMemoryPadding() int32`

GetPercentMemoryPadding returns the PercentMemoryPadding field if non-nil, zero value otherwise.

### GetPercentMemoryPaddingOk

`func (o *KubernetesContainerFactory) GetPercentMemoryPaddingOk() (*int32, bool)`

GetPercentMemoryPaddingOk returns a tuple with the PercentMemoryPadding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentMemoryPadding

`func (o *KubernetesContainerFactory) SetPercentMemoryPadding(v int32)`

SetPercentMemoryPadding sets PercentMemoryPadding field to given value.

### HasPercentMemoryPadding

`func (o *KubernetesContainerFactory) HasPercentMemoryPadding() bool`

HasPercentMemoryPadding returns a boolean if a field has been set.

### GetCpuOverCommitRatio

`func (o *KubernetesContainerFactory) GetCpuOverCommitRatio() float64`

GetCpuOverCommitRatio returns the CpuOverCommitRatio field if non-nil, zero value otherwise.

### GetCpuOverCommitRatioOk

`func (o *KubernetesContainerFactory) GetCpuOverCommitRatioOk() (*float64, bool)`

GetCpuOverCommitRatioOk returns a tuple with the CpuOverCommitRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuOverCommitRatio

`func (o *KubernetesContainerFactory) SetCpuOverCommitRatio(v float64)`

SetCpuOverCommitRatio sets CpuOverCommitRatio field to given value.

### HasCpuOverCommitRatio

`func (o *KubernetesContainerFactory) HasCpuOverCommitRatio() bool`

HasCpuOverCommitRatio returns a boolean if a field has been set.

### GetMemoryOverCommitRatio

`func (o *KubernetesContainerFactory) GetMemoryOverCommitRatio() float64`

GetMemoryOverCommitRatio returns the MemoryOverCommitRatio field if non-nil, zero value otherwise.

### GetMemoryOverCommitRatioOk

`func (o *KubernetesContainerFactory) GetMemoryOverCommitRatioOk() (*float64, bool)`

GetMemoryOverCommitRatioOk returns a tuple with the MemoryOverCommitRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryOverCommitRatio

`func (o *KubernetesContainerFactory) SetMemoryOverCommitRatio(v float64)`

SetMemoryOverCommitRatio sets MemoryOverCommitRatio field to given value.

### HasMemoryOverCommitRatio

`func (o *KubernetesContainerFactory) HasMemoryOverCommitRatio() bool`

HasMemoryOverCommitRatio returns a boolean if a field has been set.

### GetGrpcPort

`func (o *KubernetesContainerFactory) GetGrpcPort() int32`

GetGrpcPort returns the GrpcPort field if non-nil, zero value otherwise.

### GetGrpcPortOk

`func (o *KubernetesContainerFactory) GetGrpcPortOk() (*int32, bool)`

GetGrpcPortOk returns a tuple with the GrpcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrpcPort

`func (o *KubernetesContainerFactory) SetGrpcPort(v int32)`

SetGrpcPort sets GrpcPort field to given value.

### HasGrpcPort

`func (o *KubernetesContainerFactory) HasGrpcPort() bool`

HasGrpcPort returns a boolean if a field has been set.

### GetMetricsPort

`func (o *KubernetesContainerFactory) GetMetricsPort() int32`

GetMetricsPort returns the MetricsPort field if non-nil, zero value otherwise.

### GetMetricsPortOk

`func (o *KubernetesContainerFactory) GetMetricsPortOk() (*int32, bool)`

GetMetricsPortOk returns a tuple with the MetricsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsPort

`func (o *KubernetesContainerFactory) SetMetricsPort(v int32)`

SetMetricsPort sets MetricsPort field to given value.

### HasMetricsPort

`func (o *KubernetesContainerFactory) HasMetricsPort() bool`

HasMetricsPort returns a boolean if a field has been set.

### GetNarExtractionDirectory

`func (o *KubernetesContainerFactory) GetNarExtractionDirectory() string`

GetNarExtractionDirectory returns the NarExtractionDirectory field if non-nil, zero value otherwise.

### GetNarExtractionDirectoryOk

`func (o *KubernetesContainerFactory) GetNarExtractionDirectoryOk() (*string, bool)`

GetNarExtractionDirectoryOk returns a tuple with the NarExtractionDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNarExtractionDirectory

`func (o *KubernetesContainerFactory) SetNarExtractionDirectory(v string)`

SetNarExtractionDirectory sets NarExtractionDirectory field to given value.

### HasNarExtractionDirectory

`func (o *KubernetesContainerFactory) HasNarExtractionDirectory() bool`

HasNarExtractionDirectory returns a boolean if a field has been set.

### GetFunctionInstanceClassPath

`func (o *KubernetesContainerFactory) GetFunctionInstanceClassPath() string`

GetFunctionInstanceClassPath returns the FunctionInstanceClassPath field if non-nil, zero value otherwise.

### GetFunctionInstanceClassPathOk

`func (o *KubernetesContainerFactory) GetFunctionInstanceClassPathOk() (*string, bool)`

GetFunctionInstanceClassPathOk returns a tuple with the FunctionInstanceClassPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionInstanceClassPath

`func (o *KubernetesContainerFactory) SetFunctionInstanceClassPath(v string)`

SetFunctionInstanceClassPath sets FunctionInstanceClassPath field to given value.

### HasFunctionInstanceClassPath

`func (o *KubernetesContainerFactory) HasFunctionInstanceClassPath() bool`

HasFunctionInstanceClassPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


