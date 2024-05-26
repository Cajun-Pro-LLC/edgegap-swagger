# \TelemetryApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActiveDeploymentTelemetryGet**](TelemetryApi.md#ActiveDeploymentTelemetryGet) | **Get** /v1/telemetry/active-deployments/{retrieval_key} | Get the Result of an Active Deployment Telemetry Request
[**ActiveDeploymentTelemetryPost**](TelemetryApi.md#ActiveDeploymentTelemetryPost) | **Post** /v1/telemetry/active-deployments | Create a New Active Deployment Telemetry Request


# **ActiveDeploymentTelemetryGet**
> ActiveDeploymentTelemetryGetResult ActiveDeploymentTelemetryGet(ctx, retrievalKey)
Get the Result of an Active Deployment Telemetry Request

Retrieve the results of a telemetry request on active deployment(s) for given IP(s). The score array is sorted from the best to the worse deployment. You can use this to add players on a running deployment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **retrievalKey** | **string**|  | 

### Return type

[**ActiveDeploymentTelemetryGetResult**](ActiveDeploymentTelemetryGetResult.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ActiveDeploymentTelemetryPost**
> ActiveDeploymentTelemetryResponse ActiveDeploymentTelemetryPost(ctx, payload)
Create a New Active Deployment Telemetry Request

Create a telemetry request to get the best deployment(s) for given IP(s). You can use this to add players on a running deployment. If you set a webhook URL, the result will be sent to it.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payload** | [**ActiveDeploymentTelemetryRequest**](ActiveDeploymentTelemetryRequest.md)|  | 

### Return type

[**ActiveDeploymentTelemetryResponse**](ActiveDeploymentTelemetryResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

