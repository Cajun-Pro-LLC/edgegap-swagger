# \CustomSessionsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCustomSession**](CustomSessionsApi.md#DeleteCustomSession) | **Delete** /v1/deployment/{request_id}/custom/session/{custom_id} | Delete a Custom Session
[**DeleteCustomSessions**](CustomSessionsApi.md#DeleteCustomSessions) | **Delete** /v1/deployment/{request_id}/custom/sessions | Delete Custom Sessions in Bulk
[**GetCustomSession**](CustomSessionsApi.md#GetCustomSession) | **Get** /v1/deployment/{request_id}/custom/session/{custom_id} | Get a Custom Session
[**GetCustomSessions**](CustomSessionsApi.md#GetCustomSessions) | **Get** /v1/deployment/{request_id}/custom/sessions | List All Custom Session of a Deployment
[**PatchCustomSession**](CustomSessionsApi.md#PatchCustomSession) | **Patch** /v1/deployment/{request_id}/custom/session/{custom_id} | Update a Custom Session
[**PostCustomSession**](CustomSessionsApi.md#PostCustomSession) | **Post** /v1/deployment/{request_id}/custom/session/{custom_id} | Create a New Custom Session
[**PostCustomSessions**](CustomSessionsApi.md#PostCustomSessions) | **Post** /v1/deployment/{request_id}/custom/sessions | Create Custom Sessions in Bulk


# **DeleteCustomSession**
> SessionDelete DeleteCustomSession(ctx, customId, requestId)
Delete a Custom Session

Delete a custom session. Once deleted, a custom session is no more accessible and does not have a history. The deployment associated will not be deleted.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customId** | **string**| Custom ID Managed by you | 
  **requestId** | **string**| Deployment Request ID | 

### Return type

[**SessionDelete**](SessionDelete.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCustomSessions**
> BulkSessionDelete DeleteCustomSessions(ctx, requestId, payload)
Delete Custom Sessions in Bulk

Delete multiple custom sessions from a specific deployment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **requestId** | **string**| Deployment Request ID | 
  **payload** | [**CustomSessionDeleteModel**](CustomSessionDeleteModel.md)|  | 

### Return type

[**BulkSessionDelete**](BulkSessionDelete.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomSession**
> SessionGet GetCustomSession(ctx, customId, requestId)
Get a Custom Session

Retrieve the information of a custom session.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customId** | **string**| Custom ID Managed by you | 
  **requestId** | **string**| Deployment Request ID | 

### Return type

[**SessionGet**](SessionGet.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomSessions**
> Sessions GetCustomSessions(ctx, requestId)
List All Custom Session of a Deployment

List all custom sessions of a deployment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **requestId** | **string**| Deployment Request ID | 

### Return type

[**Sessions**](Sessions.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchCustomSession**
> SessionGet PatchCustomSession(ctx, customId, requestId, payload)
Update a Custom Session

Update a custom session with new specifications.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customId** | **string**| Custom ID Managed by you | 
  **requestId** | **string**| Deployment Request ID | 
  **payload** | [**CustomSessionModel**](CustomSessionModel.md)|  | 

### Return type

[**SessionGet**](SessionGet.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCustomSession**
> SessionRequest PostCustomSession(ctx, customId, requestId, payload)
Create a New Custom Session

Create a custom session with users. You must specify a custom ID and a deployment request ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customId** | **string**| Custom ID Managed by you | 
  **requestId** | **string**| Deployment Request ID | 
  **payload** | [**CustomSessionModel**](CustomSessionModel.md)|  | 

### Return type

[**SessionRequest**](SessionRequest.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCustomSessions**
> BulkSessionPost PostCustomSessions(ctx, requestId, payload)
Create Custom Sessions in Bulk

Create multiple custom sessions in a deployment. You must specify a custom ID for each.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **requestId** | **string**| Deployment Request ID | 
  **payload** | [**CustomBulkSessionsModel**](CustomBulkSessionsModel.md)|  | 

### Return type

[**BulkSessionPost**](BulkSessionPost.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

