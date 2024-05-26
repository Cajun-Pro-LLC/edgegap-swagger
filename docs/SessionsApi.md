# \SessionsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUsersSession**](SessionsApi.md#DeleteUsersSession) | **Delete** /v1/session/{session_id}/users | Delete Users From a Session
[**GetSession**](SessionsApi.md#GetSession) | **Get** /v1/session/{session_id} | Get a Session
[**GetUsersSession**](SessionsApi.md#GetUsersSession) | **Get** /v1/session/{session_id}/users | List Users of a Session
[**ListSessions**](SessionsApi.md#ListSessions) | **Get** /v1/sessions | List All Sessions
[**PutUsersSession**](SessionsApi.md#PutUsersSession) | **Put** /v1/session/{session_id}/users | Put Users in a Session
[**SessionDelete**](SessionsApi.md#SessionDelete) | **Delete** /v1/session/{session_id} | Delete a Session
[**SessionPost**](SessionsApi.md#SessionPost) | **Post** /v1/session | Create a Session
[**SessionsBulkStop**](SessionsApi.md#SessionsBulkStop) | **Post** /v1/sessions/bulk-stop | Delete Sessions in Bulk


# **DeleteUsersSession**
> SessionUserContext DeleteUsersSession(ctx, sessionId, payload)
Delete Users From a Session

Remove specified users from a session.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sessionId** | **string**|  | 
  **payload** | [**PatchSessionModel**](PatchSessionModel.md)|  | 

### Return type

[**SessionUserContext**](SessionUserContext.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSession**
> SessionGet GetSession(ctx, sessionId)
Get a Session

Retrieve the information for a session.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sessionId** | **string**|  | 

### Return type

[**SessionGet**](SessionGet.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsersSession**
> SessionUserContext GetUsersSession(ctx, sessionId)
List Users of a Session

List all the users of session.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sessionId** | **string**|  | 

### Return type

[**SessionUserContext**](SessionUserContext.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSessions**
> Sessions ListSessions(ctx, )
List All Sessions

List all the active sessions.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Sessions**](Sessions.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutUsersSession**
> SessionUserContext PutUsersSession(ctx, sessionId, payload)
Put Users in a Session

Add specified users to a session.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sessionId** | **string**|  | 
  **payload** | [**PatchSessionModel**](PatchSessionModel.md)|  | 

### Return type

[**SessionUserContext**](SessionUserContext.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SessionDelete**
> SessionDelete SessionDelete(ctx, sessionId)
Delete a Session

Delete a session. Once deleted, a session is no more accessible and does not have a history. The deployment associated will not be deleted.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sessionId** | **string**|  | 

### Return type

[**SessionDelete**](SessionDelete.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SessionPost**
> SessionRequest SessionPost(ctx, payload)
Create a Session

Create a session with users. Sessions are linked to a deployment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payload** | [**SessionModel**](SessionModel.md)|  | 

### Return type

[**SessionRequest**](SessionRequest.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SessionsBulkStop**
> SessionBulkStopResponse SessionsBulkStop(ctx, payload)
Delete Sessions in Bulk

Make a bulk delete of sessions using filters. All the sessions matching the given filters will be permanently deleted.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payload** | [**SessionBulkStopPayload**](SessionBulkStopPayload.md)|  | 

### Return type

[**SessionBulkStopResponse**](SessionBulkStopResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

