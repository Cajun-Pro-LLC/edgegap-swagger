# \MatchmakerApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMatchmaker**](MatchmakerApi.md#DeleteMatchmaker) | **Delete** /v1/aom/matchmaker/{matchmaker_name} | Delete a Matchmaker
[**DeleteMatchmakerComponent**](MatchmakerApi.md#DeleteMatchmakerComponent) | **Delete** /v1/aom/component/{component_name} | Delete a Matchmaker Component
[**DeleteMatchmakerComponentEnv**](MatchmakerApi.md#DeleteMatchmakerComponentEnv) | **Delete** /v1/aom/component/{component_name}/env/{env_key} | Delete a Matchmaker Component ENV
[**DeleteMatchmakerManagedRelease**](MatchmakerApi.md#DeleteMatchmakerManagedRelease) | **Delete** /v1/aom/matchmaker/{matchmaker_name}/release/managed/{release_version} | Delete a Matchmaker Managed Release
[**DeleteMatchmakerRelease**](MatchmakerApi.md#DeleteMatchmakerRelease) | **Delete** /v1/aom/matchmaker/{matchmaker_name}/release/{release_version} | Delete a Matchmaker Release
[**DeleteMatchmakerReleaseConfig**](MatchmakerApi.md#DeleteMatchmakerReleaseConfig) | **Delete** /v1/aom/release/config/{config_name} | Delete a Matchmaker Release Config
[**GetComponentList**](MatchmakerApi.md#GetComponentList) | **Get** /v1/aom/components | List All Matchmaker Components
[**GetEnvsList**](MatchmakerApi.md#GetEnvsList) | **Get** /v1/aom/component/{component_name}/envs | List All Matchmaker Component ENVs
[**GetMatchmaker**](MatchmakerApi.md#GetMatchmaker) | **Get** /v1/aom/matchmaker/{matchmaker_name} | Get a Matchmaker
[**GetMatchmakerComponent**](MatchmakerApi.md#GetMatchmakerComponent) | **Get** /v1/aom/component/{component_name} | Get a Matchmaker Component
[**GetMatchmakerComponentEnv**](MatchmakerApi.md#GetMatchmakerComponentEnv) | **Get** /v1/aom/component/{component_name}/env/{env_key} | Get a Matchmaker Component ENV
[**GetMatchmakerList**](MatchmakerApi.md#GetMatchmakerList) | **Get** /v1/aom/matchmakers | List All Matchmakers
[**GetMatchmakerManagedRelease**](MatchmakerApi.md#GetMatchmakerManagedRelease) | **Get** /v1/aom/matchmaker/{matchmaker_name}/release/managed/{release_version} | Get a Matchmaker Managed Release
[**GetMatchmakerRelease**](MatchmakerApi.md#GetMatchmakerRelease) | **Get** /v1/aom/matchmaker/{matchmaker_name}/release/{release_version} | Get a Matchmaker Release
[**GetMatchmakerReleaseConfig**](MatchmakerApi.md#GetMatchmakerReleaseConfig) | **Get** /v1/aom/release/config/{config_name} | Get a Matchmaker Release Config
[**GetReleaseConfigsList**](MatchmakerApi.md#GetReleaseConfigsList) | **Get** /v1/aom/release/configs | List All Matchmaker Release Configs
[**GetReleaseList**](MatchmakerApi.md#GetReleaseList) | **Get** /v1/aom/matchmaker/{matchmaker_name}/releases | List All Matchmaker Releases
[**PatchMatchmaker**](MatchmakerApi.md#PatchMatchmaker) | **Patch** /v1/aom/matchmaker/{matchmaker_name} | Update a Matchmaker
[**PatchMatchmakerComponent**](MatchmakerApi.md#PatchMatchmakerComponent) | **Patch** /v1/aom/component/{component_name} | Update a Matchmaker Component
[**PatchMatchmakerComponentEnv**](MatchmakerApi.md#PatchMatchmakerComponentEnv) | **Patch** /v1/aom/component/{component_name}/env/{env_key} | Update a Matchmaker Component ENV
[**PatchMatchmakerManagedRelease**](MatchmakerApi.md#PatchMatchmakerManagedRelease) | **Patch** /v1/aom/matchmaker/{matchmaker_name}/release/managed/{release_version} | Update a Matchmaker Managed Release
[**PatchMatchmakerRelease**](MatchmakerApi.md#PatchMatchmakerRelease) | **Patch** /v1/aom/matchmaker/{matchmaker_name}/release/{release_version} | Update a Matchmaker Release
[**PatchMatchmakerReleaseConfig**](MatchmakerApi.md#PatchMatchmakerReleaseConfig) | **Patch** /v1/aom/release/config/{config_name} | Update a Matchmaker Release Config
[**PostMatchmaker**](MatchmakerApi.md#PostMatchmaker) | **Post** /v1/aom/matchmaker | Create a Matchmaker
[**PostMatchmakerComponent**](MatchmakerApi.md#PostMatchmakerComponent) | **Post** /v1/aom/component | Create a Matchmaker Component
[**PostMatchmakerComponentEnv**](MatchmakerApi.md#PostMatchmakerComponentEnv) | **Post** /v1/aom/component/{component_name}/env | Create a Matchmaker Component ENV
[**PostMatchmakerManagedRelease**](MatchmakerApi.md#PostMatchmakerManagedRelease) | **Post** /v1/aom/matchmaker/{matchmaker_name}/release/managed | Create a Matchmaker Managed Release
[**PostMatchmakerRelease**](MatchmakerApi.md#PostMatchmakerRelease) | **Post** /v1/aom/matchmaker/{matchmaker_name}/release | Create a Matchmaker Release
[**PostMatchmakerReleaseConfig**](MatchmakerApi.md#PostMatchmakerReleaseConfig) | **Post** /v1/aom/release/config | Create a Matchmaker Release Config


# **DeleteMatchmaker**
> DeleteMatchmaker(ctx, matchmakerName)
Delete a Matchmaker

Delete a matchmaker.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **matchmakerName** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMatchmakerComponent**
> DeleteMatchmakerComponent(ctx, componentName)
Delete a Matchmaker Component

Delete a matchmaker component. It will not delete the matchmaker.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **componentName** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMatchmakerComponentEnv**
> DeleteMatchmakerComponentEnv(ctx, componentName, envKey)
Delete a Matchmaker Component ENV

Delete a matchmaker component ENV. It will not delete the component or the matchmaker.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **componentName** | **string**|  | 
  **envKey** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMatchmakerManagedRelease**
> DeleteMatchmakerManagedRelease(ctx, matchmakerName, releaseVersion)
Delete a Matchmaker Managed Release

Delete a matchmaker managed release. It will not delete the matchmaker.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **matchmakerName** | **string**|  | 
  **releaseVersion** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMatchmakerRelease**
> DeleteMatchmakerRelease(ctx, matchmakerName, releaseVersion)
Delete a Matchmaker Release

Delete a matchmaker release.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **matchmakerName** | **string**|  | 
  **releaseVersion** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMatchmakerReleaseConfig**
> DeleteMatchmakerReleaseConfig(ctx, configName)
Delete a Matchmaker Release Config

Delete a matchmaker release config.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **configName** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetComponentList**
> MatchmakerComponentListResponse GetComponentList(ctx, )
List All Matchmaker Components

List all components for a specific matchmaker.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**MatchmakerComponentListResponse**](MatchmakerComponentListResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEnvsList**
> MatchmakerComponentEnvListResponse GetEnvsList(ctx, componentName)
List All Matchmaker Component ENVs

List all ENVs for a specific matchmaker component.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **componentName** | **string**|  | 

### Return type

[**MatchmakerComponentEnvListResponse**](MatchmakerComponentEnvListResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMatchmaker**
> MatchmakerResponse GetMatchmaker(ctx, matchmakerName)
Get a Matchmaker

Retrieve a matchmaker.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **matchmakerName** | **string**|  | 

### Return type

[**MatchmakerResponse**](MatchmakerResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMatchmakerComponent**
> MatchmakerComponentResponse GetMatchmakerComponent(ctx, componentName)
Get a Matchmaker Component

Retrieve a matchmaker component.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **componentName** | **string**|  | 

### Return type

[**MatchmakerComponentResponse**](MatchmakerComponentResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMatchmakerComponentEnv**
> MatchmakerComponentEnvsResponse GetMatchmakerComponentEnv(ctx, componentName, envKey)
Get a Matchmaker Component ENV

Retrieve a matchmaker component ENV.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **componentName** | **string**|  | 
  **envKey** | **string**|  | 

### Return type

[**MatchmakerComponentEnvsResponse**](MatchmakerComponentEnvsResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMatchmakerList**
> MatchmakerListResponse GetMatchmakerList(ctx, )
List All Matchmakers

List all matchmakers.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**MatchmakerListResponse**](MatchmakerListResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMatchmakerManagedRelease**
> MatchmakerManagedReleaseResponse GetMatchmakerManagedRelease(ctx, matchmakerName, releaseVersion)
Get a Matchmaker Managed Release

Retrieve a matchmaker managed release.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **matchmakerName** | **string**|  | 
  **releaseVersion** | **string**|  | 

### Return type

[**MatchmakerManagedReleaseResponse**](MatchmakerManagedReleaseResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMatchmakerRelease**
> MatchmakerReleaseResponse GetMatchmakerRelease(ctx, matchmakerName, releaseVersion)
Get a Matchmaker Release

Retrieve a matchmaker release.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **matchmakerName** | **string**|  | 
  **releaseVersion** | **string**|  | 

### Return type

[**MatchmakerReleaseResponse**](MatchmakerReleaseResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMatchmakerReleaseConfig**
> MatchmakerReleaseConfigResponse GetMatchmakerReleaseConfig(ctx, configName)
Get a Matchmaker Release Config

Get a matchmaker release config.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **configName** | **string**|  | 

### Return type

[**MatchmakerReleaseConfigResponse**](MatchmakerReleaseConfigResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReleaseConfigsList**
> MatchmakerComponentEnvListResponse GetReleaseConfigsList(ctx, )
List All Matchmaker Release Configs

List all configs for a specific matchmaker release.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**MatchmakerComponentEnvListResponse**](MatchmakerComponentEnvListResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReleaseList**
> MatchmakerComponentEnvListResponse GetReleaseList(ctx, matchmakerName)
List All Matchmaker Releases

List all releases of a specific matchmaker.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **matchmakerName** | **string**|  | 

### Return type

[**MatchmakerComponentEnvListResponse**](MatchmakerComponentEnvListResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchMatchmaker**
> MatchmakerResponse PatchMatchmaker(ctx, matchmakerName, payload)
Update a Matchmaker

Update a matchmaker with new specifications.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **matchmakerName** | **string**|  | 
  **payload** | [**MatchmakerUpdate**](MatchmakerUpdate.md)|  | 

### Return type

[**MatchmakerResponse**](MatchmakerResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchMatchmakerComponent**
> MatchmakerComponentResponse PatchMatchmakerComponent(ctx, componentName, payload)
Update a Matchmaker Component

Update a matchmaker component with new specifications.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **componentName** | **string**|  | 
  **payload** | [**MatchmakerComponentUpdate**](MatchmakerComponentUpdate.md)|  | 

### Return type

[**MatchmakerComponentResponse**](MatchmakerComponentResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchMatchmakerComponentEnv**
> MatchmakerComponentEnvsResponse PatchMatchmakerComponentEnv(ctx, componentName, envKey, payload)
Update a Matchmaker Component ENV

Update a matchmaker component ENV.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **componentName** | **string**|  | 
  **envKey** | **string**|  | 
  **payload** | [**MatchmakerComponentEnvsUpdate**](MatchmakerComponentEnvsUpdate.md)|  | 

### Return type

[**MatchmakerComponentEnvsResponse**](MatchmakerComponentEnvsResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchMatchmakerManagedRelease**
> MatchmakerManagedReleaseResponse PatchMatchmakerManagedRelease(ctx, matchmakerName, releaseVersion, payload)
Update a Matchmaker Managed Release

Update a matchmaker managed release.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **matchmakerName** | **string**|  | 
  **releaseVersion** | **string**|  | 
  **payload** | [**MatchmakerManagedReleaseUpdate**](MatchmakerManagedReleaseUpdate.md)|  | 

### Return type

[**MatchmakerManagedReleaseResponse**](MatchmakerManagedReleaseResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchMatchmakerRelease**
> MatchmakerReleaseResponse PatchMatchmakerRelease(ctx, matchmakerName, releaseVersion, payload)
Update a Matchmaker Release

Update a matchmaker release.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **matchmakerName** | **string**|  | 
  **releaseVersion** | **string**|  | 
  **payload** | [**MatchmakerReleaseUpdate**](MatchmakerReleaseUpdate.md)|  | 

### Return type

[**MatchmakerReleaseResponse**](MatchmakerReleaseResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchMatchmakerReleaseConfig**
> MatchmakerReleaseConfigResponse PatchMatchmakerReleaseConfig(ctx, configName, payload)
Update a Matchmaker Release Config

Update a matchmaker release config.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **configName** | **string**|  | 
  **payload** | [**MatchmakerReleaseConfigUpdate**](MatchmakerReleaseConfigUpdate.md)|  | 

### Return type

[**MatchmakerReleaseConfigResponse**](MatchmakerReleaseConfigResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostMatchmaker**
> MatchmakerResponse PostMatchmaker(ctx, payload)
Create a Matchmaker

Create a new matchmaker. A matchmaker is a top-level object; you must create child resources to work properly.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payload** | [**MatchmakerCreate**](MatchmakerCreate.md)|  | 

### Return type

[**MatchmakerResponse**](MatchmakerResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostMatchmakerComponent**
> MatchmakerComponentResponse PostMatchmakerComponent(ctx, payload)
Create a Matchmaker Component

Create a new matchmaker component.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payload** | [**MatchmakerComponentCreate**](MatchmakerComponentCreate.md)|  | 

### Return type

[**MatchmakerComponentResponse**](MatchmakerComponentResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostMatchmakerComponentEnv**
> MatchmakerComponentEnvsResponse PostMatchmakerComponentEnv(ctx, componentName, payload)
Create a Matchmaker Component ENV

Create a new matchmaker component ENV.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **componentName** | **string**|  | 
  **payload** | [**MatchmakerComponentEnvsCreate**](MatchmakerComponentEnvsCreate.md)|  | 

### Return type

[**MatchmakerComponentEnvsResponse**](MatchmakerComponentEnvsResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostMatchmakerManagedRelease**
> MatchmakerManagedReleaseResponse PostMatchmakerManagedRelease(ctx, matchmakerName, payload)
Create a Matchmaker Managed Release

Create a matchmaker managed release.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **matchmakerName** | **string**|  | 
  **payload** | [**MatchmakerManagedReleaseCreate**](MatchmakerManagedReleaseCreate.md)|  | 

### Return type

[**MatchmakerManagedReleaseResponse**](MatchmakerManagedReleaseResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostMatchmakerRelease**
> MatchmakerReleaseResponse PostMatchmakerRelease(ctx, matchmakerName, payload)
Create a Matchmaker Release

Create a matchmaker release.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **matchmakerName** | **string**|  | 
  **payload** | [**MatchmakerReleaseCreate**](MatchmakerReleaseCreate.md)|  | 

### Return type

[**MatchmakerReleaseResponse**](MatchmakerReleaseResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostMatchmakerReleaseConfig**
> MatchmakerReleaseConfigResponse PostMatchmakerReleaseConfig(ctx, payload)
Create a Matchmaker Release Config

Create a matchmaker release config.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payload** | [**MatchmakerReleaseConfigCreate**](MatchmakerReleaseConfigCreate.md)|  | 

### Return type

[**MatchmakerReleaseConfigResponse**](MatchmakerReleaseConfigResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

