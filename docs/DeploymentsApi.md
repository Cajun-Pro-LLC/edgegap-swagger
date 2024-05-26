# \DeploymentsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Deploy**](DeploymentsApi.md#Deploy) | **Post** /v1/deploy | Create a Deployment
[**DeploymentDelete**](DeploymentsApi.md#DeploymentDelete) | **Delete** /v1/stop/{request_id} | Delete a Deployment
[**DeploymentGetLogs**](DeploymentsApi.md#DeploymentGetLogs) | **Get** /v1/deployment/{request_id}/container-logs | Get Deployment Container Logs
[**DeploymentStatusGet**](DeploymentsApi.md#DeploymentStatusGet) | **Get** /v1/status/{request_id} | Get a Deployment Status and Information
[**DeploymentUpdate**](DeploymentsApi.md#DeploymentUpdate) | **Patch** /v1/deployments/{request_id} | Updates properties of a deployment
[**DeploymentsBulkDelete**](DeploymentsApi.md#DeploymentsBulkDelete) | **Post** /v1/deployments/bulk-stop | Delete Deployments in Bulk
[**DeploymentsGet**](DeploymentsApi.md#DeploymentsGet) | **Get** /v1/deployments | List All Deployments
[**SelfDeploymentDelete**](DeploymentsApi.md#SelfDeploymentDelete) | **Delete** /v1/self/stop/{request_id}/{access_point_id} | Delete a Deployment from inside the container


# **Deploy**
> Request Deploy(ctx, payload)
Create a Deployment

Create a new deployment. Deployment is a server instance of your application version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payload** | [**DeployModel**](DeployModel.md)|  | 

### Return type

[**Request**](Request.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeploymentDelete**
> Delete DeploymentDelete(ctx, requestId, optional)
Delete a Deployment

Delete an instance of deployment. It will stop the running container and all its games.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **requestId** | **string**| Unique Identifier to keep track of your request across all Arbitrium ecosystem. It&#39;s included in the response of the app deploy, example:    93924761ccde | 
 **optional** | ***DeploymentsApiDeploymentDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeploymentsApiDeploymentDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **containerLogStorage** | **optional.String**| If you want to enable the container log storage for the deployment. You can put &#39;true&#39; if you already have endpoint storage associated with your deployment&#39;s app version. You can put &#39;false&#39; if it is enabled by default and you want to disable it for this specific request. Or you can put the name of your endpoint storage and if it is valid we will store the container logs. | 

### Return type

[**Delete**](Delete.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeploymentGetLogs**
> ApiModelContainerlogs DeploymentGetLogs(ctx, requestId)
Get Deployment Container Logs

Retrieve the logs of your container. Logs are not available when your deployment is terminated. You can set up an endpoint storage to save your logs. <a target='_blank' href='https://docs.edgegap.com/docs/deployment/endpoint-storage'>Endpoint Storage Documentation</a>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **requestId** | **string**|  | 

### Return type

[**ApiModelContainerlogs**](api-model-containerlogs.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeploymentStatusGet**
> Status DeploymentStatusGet(ctx, requestId)
Get a Deployment Status and Information

Retrieve the information for a deployment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **requestId** | **string**| Unique Identifier to keep track of your request across all Arbitrium ecosystem. It&#39;s included in the response of the app deploy, example:    93924761ccde | 

### Return type

[**Status**](Status.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeploymentUpdate**
> DeploymentUpdateResponse DeploymentUpdate(ctx, requestId, payload)
Updates properties of a deployment

Updates properties of a deployment. Currently only the `is_joinable_by_session` property can be updated.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **requestId** | **string**|  | 
  **payload** | [**DeploymentUpdatePayload**](DeploymentUpdatePayload.md)|  | 

### Return type

[**DeploymentUpdateResponse**](DeploymentUpdateResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeploymentsBulkDelete**
> DeploymentBulkStopResponse DeploymentsBulkDelete(ctx, payload)
Delete Deployments in Bulk

Make a bulk delete of deployments using filters. All the deployments matching the given filters will be permanently deleted.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payload** | [**DeploymentBulkStopPayload**](DeploymentBulkStopPayload.md)|  | 

### Return type

[**DeploymentBulkStopResponse**](DeploymentBulkStopResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeploymentsGet**
> Deployments DeploymentsGet(ctx, )
List All Deployments

List all deployments.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Deployments**](Deployments.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SelfDeploymentDelete**
> Delete SelfDeploymentDelete(ctx, requestId, accessPointId, authorization, optional)
Delete a Deployment from inside the container

Delete a deployment from the inside of a container. You should use this URL inside your deployment. The URL is injected in your deployment and can be found via the environment variable ARBITRIUM_DELETE_URL.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **requestId** | **string**| Unique Identifier to keep track of your request across all Arbitrium ecosystem. It&#39;s included in the response of the app deploy, example:    93924761ccde | 
  **accessPointId** | **int32**| Access Point Number provided by our system | 
  **authorization** | **string**| Auto Generated token. This token is injected in your deployment and can be found via the environment variable named ARBITRIUM_DELETE_TOKEN | 
 **optional** | ***DeploymentsApiSelfDeploymentDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeploymentsApiSelfDeploymentDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **containerLogStorage** | **optional.String**| If you want to enable the container log storage for the deployment. You can put &#39;true&#39; if you already have endpoint storage associated with your deployment&#39;s app version. You can put &#39;false&#39; if it is enabled by default and you want to disable it for this specific request. Or you can put the name of your endpoint storage and if it is valid we will store the container logs. | 

### Return type

[**Delete**](Delete.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

