# \RelaysApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RelaySessionCreate**](RelaysApi.md#RelaySessionCreate) | **Post** /v1/relays/sessions | Create a Relay Session
[**RelaySessionDelete**](RelaysApi.md#RelaySessionDelete) | **Delete** /v1/relays/sessions/{session_id} | Delete a Relay Session
[**RelaySessionGet**](RelaysApi.md#RelaySessionGet) | **Get** /v1/relays/sessions/{session_id} | Get a Relay Session
[**RelaySessionList**](RelaysApi.md#RelaySessionList) | **Get** /v1/relays/sessions | List all Relay Sessions
[**RelayUserAuthorize**](RelaysApi.md#RelayUserAuthorize) | **Post** /v1/relays/sessions:authorize-user | Authorize a user on a Relay Session
[**RelayUserRevoke**](RelaysApi.md#RelayUserRevoke) | **Post** /v1/relays/sessions:revoke-user | Remove a user on a Relay Session


# **RelaySessionCreate**
> RelaySessionBaseResponse RelaySessionCreate(ctx, payload)
Create a Relay Session

Create a relay session with users.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payload** | [**RelaySessionCreatePayload**](RelaySessionCreatePayload.md)|  | 

### Return type

[**RelaySessionBaseResponse**](RelaySessionBaseResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RelaySessionDelete**
> RelaySessionDelete(ctx, sessionId)
Delete a Relay Session

Delete a relay session.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sessionId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RelaySessionGet**
> RelaySessionBaseResponse RelaySessionGet(ctx, sessionId)
Get a Relay Session

Retrieve the information for a relay session.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sessionId** | **string**|  | 

### Return type

[**RelaySessionBaseResponse**](RelaySessionBaseResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RelaySessionList**
> RelaySessionListResponse RelaySessionList(ctx, )
List all Relay Sessions

List all the active relay sessions.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**RelaySessionListResponse**](RelaySessionListResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RelayUserAuthorize**
> RelaySessionUserBaseResponse RelayUserAuthorize(ctx, payload)
Authorize a user on a Relay Session

Authorize a user on a Relay Session

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payload** | [**RelayUserAuthorizePayload**](RelayUserAuthorizePayload.md)|  | 

### Return type

[**RelaySessionUserBaseResponse**](RelaySessionUserBaseResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RelayUserRevoke**
> RelaySessionBaseResponse RelayUserRevoke(ctx, payload)
Remove a user on a Relay Session

Authorize a user on a Relay Session

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payload** | [**RelayUserRevokePayload**](RelayUserRevokePayload.md)|  | 

### Return type

[**RelaySessionBaseResponse**](RelaySessionBaseResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

