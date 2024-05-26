# \ContextApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContextCreateDeploymentTag**](ContextApi.md#ContextCreateDeploymentTag) | **Post** /v1/context/{request_id}/{security_number}/tags | Create a Tag for a Running Deployment
[**ContextDeleteDeploymentTag**](ContextApi.md#ContextDeleteDeploymentTag) | **Delete** /v1/context/{request_id}/{security_number}/tags/{tag_name} | Delete a Tag from a Running Deployment
[**ContextGet**](ContextApi.md#ContextGet) | **Get** /v1/context/{request_id}/{security_number} | Get the Context of a Deployment


# **ContextCreateDeploymentTag**
> ContextDeploymentTagResponse ContextCreateDeploymentTag(ctx, requestId, securityNumber, authorization, payload)
Create a Tag for a Running Deployment

Create a tag for a running deployment. You should use this URL inside your deployment container. The URL is injected in your deployment and accessible via the environment variable ARBITRIUM_CONTEXT_URL and you need to append \"/tags\" at the end of this URL.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **requestId** | **string**| Unique Identifier to keep track of your request across all Arbitrium ecosystem.  | 
  **securityNumber** | **int32**| Random Security number generate to validate the request id. | 
  **authorization** | **string**| Auto Generated token. This token is injected in your deployment and can be found via the environment variable named ARBITRIUM_CONTEXT_TOKEN  | 
  **payload** | [**Payload**](Payload.md)|  | 

### Return type

[**ContextDeploymentTagResponse**](ContextDeploymentTagResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContextDeleteDeploymentTag**
> ContextDeleteDeploymentTag(ctx, tagName, requestId, securityNumber, authorization)
Delete a Tag from a Running Deployment

Delete a tag from a running deployment. You should use this URL inside your deployment container. The URL is injected in your deployment and accessible via the environment variable ARBITRIUM_CONTEXT_URL and you need to append \"/tags/{tag_name}\" at the end of this URL.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tagName** | **string**|  | 
  **requestId** | **string**| Unique Identifier to keep track of your request across all Arbitrium ecosystem.  | 
  **securityNumber** | **int32**| Random Security number generate to validate the request id. | 
  **authorization** | **string**| Auto Generated token. This token is injected in your deployment and can be found via the environment variable named ARBITRIUM_CONTEXT_TOKEN  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContextGet**
> Deployment ContextGet(ctx, requestId, securityNumber, authorization)
Get the Context of a Deployment

Retrieve contextual information about your live deployment. You should use this URL inside your deployment container. The URL is injected in your deployment and accessible via the environment variable ARBITRIUM_CONTEXT_URL.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **requestId** | **string**| Unique Identifier to keep track of your request across all Arbitrium ecosystem.  | 
  **securityNumber** | **int32**| Random Security number generate to validate the request id. | 
  **authorization** | **string**| Auto Generated token. This token is injected in your deployment and can be found via the environment variable named ARBITRIUM_CONTEXT_TOKEN  | 

### Return type

[**Deployment**](Deployment.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

