# \EndpointStorageApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EndpointCreate**](EndpointStorageApi.md#EndpointCreate) | **Post** /v1/storage/endpoint | Create a New Endpoint Storage
[**EndpointDelete**](EndpointStorageApi.md#EndpointDelete) | **Delete** /v1/storage/endpoint/{endpoint_name} | Delete an Endpoint Storage
[**EndpointGet**](EndpointStorageApi.md#EndpointGet) | **Get** /v1/storage/endpoint/{endpoint_name} | Get an Endpoint Storage
[**EndpointUpdate**](EndpointStorageApi.md#EndpointUpdate) | **Patch** /v1/storage/endpoint/{endpoint_name} | Update an Endpoint Storage
[**EndpointsList**](EndpointStorageApi.md#EndpointsList) | **Get** /v1/storage/endpoints | List All Endpoint Storage
[**PullProfileCreate**](EndpointStorageApi.md#PullProfileCreate) | **Post** /v1/storage/endpoint/{endpoint_name}/pull-profile | Create a New Pull Profile
[**PullProfileDelete**](EndpointStorageApi.md#PullProfileDelete) | **Delete** /v1/storage/endpoint/{endpoint_name}/pull-profile/{pull_profile_name} | Delete a Pull Profile
[**PullProfileGet**](EndpointStorageApi.md#PullProfileGet) | **Get** /v1/storage/endpoint/{endpoint_name}/pull-profile/{pull_profile_name} | Get a Pull Profile
[**PullProfileLinkAppVersion**](EndpointStorageApi.md#PullProfileLinkAppVersion) | **Put** /v1/storage/endpoint/{endpoint_name}/pull-profile/{pull_profile_name}/app/{app_name}/version/{version_name} | Link a Pull Profile to an Application Version
[**PullProfileList**](EndpointStorageApi.md#PullProfileList) | **Get** /v1/storage/endpoint/{endpoint_name}/pull-profiles | List All Pull Profile of an Endpoint Storage
[**PullProfileUnlinkAppVersion**](EndpointStorageApi.md#PullProfileUnlinkAppVersion) | **Delete** /v1/storage/endpoint/{endpoint_name}/pull-profile/{pull_profile_name}/app/{app_name}/version/{version_name} | Unlink a Pull Profile From an Application Version
[**PullProfileUpdate**](EndpointStorageApi.md#PullProfileUpdate) | **Patch** /v1/storage/endpoint/{endpoint_name}/pull-profile/{pull_profile_name} | Update a Pull Profile


# **EndpointCreate**
> EndpointStoragePostResponse EndpointCreate(ctx, payload)
Create a New Endpoint Storage

Create an endpoint storage to store your container logs at the end of a deployment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payload** | [**EndpointStoragePostPayload**](EndpointStoragePostPayload.md)|  | 

### Return type

[**EndpointStoragePostResponse**](EndpointStoragePostResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EndpointDelete**
> EndpointStorageDeleteResponse EndpointDelete(ctx, endpointName)
Delete an Endpoint Storage

Delete an endpoint storage. All the application versions linked to it won't be able to store logs anymore.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **endpointName** | **string**|  | 

### Return type

[**EndpointStorageDeleteResponse**](EndpointStorageDeleteResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EndpointGet**
> EndpointStorageGetResponse EndpointGet(ctx, endpointName)
Get an Endpoint Storage

Retrieve an endpoint storage. The ```secret_access_key``` won't be displayed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **endpointName** | **string**|  | 

### Return type

[**EndpointStorageGetResponse**](EndpointStorageGetResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EndpointUpdate**
> EndpointStoragePatchResponse EndpointUpdate(ctx, endpointName, payload)
Update an Endpoint Storage

Update an Endpoint Storage with new specifications.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **endpointName** | **string**|  | 
  **payload** | [**EndpointStoragePatchPayload**](EndpointStoragePatchPayload.md)|  | 

### Return type

[**EndpointStoragePatchResponse**](EndpointStoragePatchResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EndpointsList**
> EndpointStorageListResponse EndpointsList(ctx, optional)
List All Endpoint Storage

List all endpoint storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EndpointStorageApiEndpointsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EndpointStorageApiEndpointsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Page number for pagination | [default to 1]
 **limit** | **optional.Int32**| Limit of Fleet for each page | [default to 10]
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**EndpointStorageListResponse**](EndpointStorageListResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PullProfileCreate**
> PullProfilePostResponse PullProfileCreate(ctx, endpointName, payload)
Create a New Pull Profile

Create a pull profile. Pull profile will upload data from an endpoint storage to a deployment container on boot. You must link the application version to the pull profile first.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **endpointName** | **string**|  | 
  **payload** | [**PullProfilePostPayload**](PullProfilePostPayload.md)|  | 

### Return type

[**PullProfilePostResponse**](PullProfilePostResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PullProfileDelete**
> PullProfileDelete(ctx, endpointName, pullProfileName)
Delete a Pull Profile

Delete a pull profile. All the application versions linked won't receive the data upload anymore. It will not delete your endpoint storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **endpointName** | **string**|  | 
  **pullProfileName** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PullProfileGet**
> PullProfileGetResponse PullProfileGet(ctx, endpointName, pullProfileName)
Get a Pull Profile

Retrieve a pull profile and its specifications.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **endpointName** | **string**|  | 
  **pullProfileName** | **string**|  | 

### Return type

[**PullProfileGetResponse**](PullProfileGetResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PullProfileLinkAppVersion**
> PullProfileAppVersionLinkResponse PullProfileLinkAppVersion(ctx, endpointName, pullProfileName, appName, versionName)
Link a Pull Profile to an Application Version

Link a pull profile to an app version. Without a link, the pull profile by itself will do nothing.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **endpointName** | **string**|  | 
  **pullProfileName** | **string**|  | 
  **appName** | **string**|  | 
  **versionName** | **string**|  | 

### Return type

[**PullProfileAppVersionLinkResponse**](PullProfileAppVersionLinkResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PullProfileList**
> PullProfilesListResponse PullProfileList(ctx, endpointName, optional)
List All Pull Profile of an Endpoint Storage

List all pull profiles of an endpoint storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **endpointName** | **string**|  | 
 **optional** | ***EndpointStorageApiPullProfileListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EndpointStorageApiPullProfileListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Page number for pagination | [default to 1]
 **limit** | **optional.Int32**| Limit of pull profiles for each page | [default to 10]

### Return type

[**PullProfilesListResponse**](PullProfilesListResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PullProfileUnlinkAppVersion**
> PullProfileUnlinkAppVersion(ctx, endpointName, pullProfileName, appName, versionName)
Unlink a Pull Profile From an Application Version

Unlink a pull profile from an app version. It will not delete the pull profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **endpointName** | **string**|  | 
  **pullProfileName** | **string**|  | 
  **appName** | **string**|  | 
  **versionName** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PullProfileUpdate**
> PulloProfilePatchResponse PullProfileUpdate(ctx, endpointName, pullProfileName, payload)
Update a Pull Profile

Update a pull profile with new specifications.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **endpointName** | **string**|  | 
  **pullProfileName** | **string**|  | 
  **payload** | [**PullProfilePatchPayload**](PullProfilePatchPayload.md)|  | 

### Return type

[**PulloProfilePatchResponse**](PulloProfilePatchResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

