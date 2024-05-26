# \DeploymentTagsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeploymentsTagsCreate**](DeploymentTagsApi.md#DeploymentsTagsCreate) | **Post** /v1/deployments/{request_id}/tags | Create Tag for a Deployment
[**DeploymentsTagsDelete**](DeploymentTagsApi.md#DeploymentsTagsDelete) | **Delete** /v1/deployments/{request_id}/tags/{tag_name} | Delete Tag for a Deployment
[**DeploymentsTagsList**](DeploymentTagsApi.md#DeploymentsTagsList) | **Get** /v1/deployments/{request_id}/tags | List tags for a Deployment
[**DeploymentsTagsRead**](DeploymentTagsApi.md#DeploymentsTagsRead) | **Get** /v1/deployments/{request_id}/tags/{tag_name} | Get tag for a Deployment
[**DeploymentsTagsUpdate**](DeploymentTagsApi.md#DeploymentsTagsUpdate) | **Patch** /v1/deployments/{request_id}/tags/{tag_name} | Update Tag for a Deployment


# **DeploymentsTagsCreate**
> DeploymentTagResponse DeploymentsTagsCreate(ctx, requestId, payload)
Create Tag for a Deployment

Create a tag for a deployment. The tag will however not be injected into a running container.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **requestId** | **string**|  | 
  **payload** | [**DeploymentTagPayload**](DeploymentTagPayload.md)|  | 

### Return type

[**DeploymentTagResponse**](DeploymentTagResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeploymentsTagsDelete**
> DeploymentsTagsDelete(ctx, requestId, tagName)
Delete Tag for a Deployment

Delete a tag for a deployment. The tag will however not be removed from the environment of a running container.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **requestId** | **string**|  | 
  **tagName** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeploymentsTagsList**
> DeploymentTagListResponse DeploymentsTagsList(ctx, requestId)
List tags for a Deployment

List tags for a deployment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **requestId** | **string**|  | 

### Return type

[**DeploymentTagListResponse**](DeploymentTagListResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeploymentsTagsRead**
> DeploymentTagResponse DeploymentsTagsRead(ctx, requestId, tagName)
Get tag for a Deployment

Get tag for a deployment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **requestId** | **string**|  | 
  **tagName** | **string**|  | 

### Return type

[**DeploymentTagResponse**](DeploymentTagResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeploymentsTagsUpdate**
> DeploymentTagResponse DeploymentsTagsUpdate(ctx, requestId, tagName, payload)
Update Tag for a Deployment

Update a tag for a deployment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **requestId** | **string**|  | 
  **tagName** | **string**|  | 
  **payload** | [**DeploymentTagPayload**](DeploymentTagPayload.md)|  | 

### Return type

[**DeploymentTagResponse**](DeploymentTagResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

