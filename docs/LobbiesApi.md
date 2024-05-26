# \LobbiesApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LobbyCreate**](LobbiesApi.md#LobbyCreate) | **Post** /v1/lobbies | Create a Lobby
[**LobbyDeploy**](LobbiesApi.md#LobbyDeploy) | **Post** /v1/lobbies:deploy | Deploy a Lobby
[**LobbyGet**](LobbiesApi.md#LobbyGet) | **Get** /v1/lobbies/{lobby_name} | Get a Lobby
[**LobbyTerminate**](LobbiesApi.md#LobbyTerminate) | **Post** /v1/lobbies:terminate | Terminate a Lobby


# **LobbyCreate**
> LobbyReadResponse LobbyCreate(ctx, payload)
Create a Lobby

Create a named lobby.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payload** | [**LobbyCreatePayload**](LobbyCreatePayload.md)|  | 

### Return type

[**LobbyReadResponse**](LobbyReadResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LobbyDeploy**
> LobbyReadResponse LobbyDeploy(ctx, payload)
Deploy a Lobby

Deploy the lobby with the given name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payload** | [**LobbyDeployPayload**](LobbyDeployPayload.md)|  | 

### Return type

[**LobbyReadResponse**](LobbyReadResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LobbyGet**
> LobbyReadResponse LobbyGet(ctx, lobbyName)
Get a Lobby

Get a named lobby.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lobbyName** | **string**|  | 

### Return type

[**LobbyReadResponse**](LobbyReadResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LobbyTerminate**
> LobbyReadResponse LobbyTerminate(ctx, payload)
Terminate a Lobby

Terminate the lobby with the given name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payload** | [**LobbyTerminatePayload**](LobbyTerminatePayload.md)|  | 

### Return type

[**LobbyReadResponse**](LobbyReadResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

