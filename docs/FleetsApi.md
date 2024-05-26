# \FleetsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FleetCreate**](FleetsApi.md#FleetCreate) | **Post** /v1/fleet | Create a Fleet
[**FleetDelete**](FleetsApi.md#FleetDelete) | **Delete** /v1/fleet/{fleet_name} | Delete a Fleet
[**FleetGet**](FleetsApi.md#FleetGet) | **Get** /v1/fleet/{fleet_name} | Get a Fleet
[**FleetLinkAppVersion**](FleetsApi.md#FleetLinkAppVersion) | **Put** /v1/fleet/{fleet_name}/app/{app_name}/version/{version_name} | Link an Application Version to a Fleet
[**FleetPoliciesCreate**](FleetsApi.md#FleetPoliciesCreate) | **Post** /v1/fleet/{fleet_name}/policies | Create a Fleet Policy
[**FleetPoliciesDelete**](FleetsApi.md#FleetPoliciesDelete) | **Delete** /v1/fleet/{fleet_name}/policies/{policy_name} | Delete a Policy
[**FleetPoliciesGet**](FleetsApi.md#FleetPoliciesGet) | **Get** /v1/fleet/{fleet_name}/policies/{policy_name} | Get a Policy
[**FleetPoliciesList**](FleetsApi.md#FleetPoliciesList) | **Get** /v1/fleet/{fleet_name}/policies | List All Policies of a Fleet
[**FleetPoliciesUpdate**](FleetsApi.md#FleetPoliciesUpdate) | **Patch** /v1/fleet/{fleet_name}/policies/{policy_name} | Update a Policy
[**FleetUnlinkAppVersion**](FleetsApi.md#FleetUnlinkAppVersion) | **Delete** /v1/fleet/{fleet_name}/app/{app_name}/version/{version_name} | Unlink an Application Version From a Fleet
[**FleetUpdate**](FleetsApi.md#FleetUpdate) | **Patch** /v1/fleet/{fleet_name} | Update a Fleet
[**Fleets**](FleetsApi.md#Fleets) | **Get** /v1/fleets | List All Fleets


# **FleetCreate**
> FleetPostResponse FleetCreate(ctx, payload)
Create a Fleet

Create a fleet. A fleet is a top-level object; you must create child resources to work properly.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payload** | [**FleetPostPayload**](FleetPostPayload.md)|  | 

### Return type

[**FleetPostResponse**](FleetPostResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FleetDelete**
> FleetDeleteResponse FleetDelete(ctx, fleetName)
Delete a Fleet

Delete a fleet, its policies and links between the application versions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fleetName** | **string**| The name of the fleet | 

### Return type

[**FleetDeleteResponse**](FleetDeleteResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FleetGet**
> FleetGetResponse FleetGet(ctx, fleetName)
Get a Fleet

Retrieve a fleet with its details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fleetName** | **string**| The name of the fleet | 

### Return type

[**FleetGetResponse**](FleetGetResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FleetLinkAppVersion**
> HorizontalScalerAppVersionLink FleetLinkAppVersion(ctx, fleetName, appName, versionName)
Link an Application Version to a Fleet

Link an application version to a fleet. By linking this version, the fleet will automatically create deployments of this version according to the fleet policies.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fleetName** | **string**| The name of the fleet | 
  **appName** | **string**| The name of the application to link | 
  **versionName** | **string**| The name of the application version to link | 

### Return type

[**HorizontalScalerAppVersionLink**](HorizontalScalerAppVersionLink.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FleetPoliciesCreate**
> FleetPoliciesPostResponse FleetPoliciesCreate(ctx, fleetName, payload)
Create a Fleet Policy

Create a fleet policy. Policies are conditions that the fleet must respect.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fleetName** | **string**| The name of the fleet | 
  **payload** | [**FleetPoliciesPostPayload**](FleetPoliciesPostPayload.md)|  | 

### Return type

[**FleetPoliciesPostResponse**](FleetPoliciesPostResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FleetPoliciesDelete**
> FleetPoliciesDelete(ctx, fleetName, policyName)
Delete a Policy

Delete a policy. It will not delete the fleet.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fleetName** | **string**| The name of the fleet | 
  **policyName** | **string**| The name of the policy to delete | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FleetPoliciesGet**
> FleetPoliciesGetResponse FleetPoliciesGet(ctx, fleetName, policyName)
Get a Policy

Retrieve a specific policy of a fleet.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fleetName** | **string**| The name of the fleet | 
  **policyName** | **string**| The name of the policy | 

### Return type

[**FleetPoliciesGetResponse**](FleetPoliciesGetResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FleetPoliciesList**
> HorizontalScalerConstraintList FleetPoliciesList(ctx, fleetName, optional)
List All Policies of a Fleet

List all the policies of a fleet.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fleetName** | **string**| The name of the fleet | 
 **optional** | ***FleetsApiFleetPoliciesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FleetsApiFleetPoliciesListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Page number for pagination | [default to 1]
 **limit** | **optional.Int32**| Limit of Fleet for each page | [default to 10]
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**HorizontalScalerConstraintList**](HorizontalScalerConstraintList.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FleetPoliciesUpdate**
> FleetPoliciesPatchPayload FleetPoliciesUpdate(ctx, fleetName, policyName, payload)
Update a Policy

Update a policy with new specifications.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fleetName** | **string**| The name of the fleet | 
  **policyName** | **string**| The name of the policy | 
  **payload** | [**FleetPoliciesPatchPayload**](FleetPoliciesPatchPayload.md)|  | 

### Return type

[**FleetPoliciesPatchPayload**](FleetPoliciesPatchPayload.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FleetUnlinkAppVersion**
> FleetUnlinkAppVersion(ctx, fleetName, appName, versionName)
Unlink an Application Version From a Fleet

Unlink an application version from a fleet. It will not delete the application version or the fleet

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fleetName** | **string**| The name of the fleet | 
  **appName** | **string**| The name of the application to link | 
  **versionName** | **string**| The name of the application version to link | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FleetUpdate**
> FleetPatchResponse FleetUpdate(ctx, fleetName, payload)
Update a Fleet

Update a fleet with new specifications.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fleetName** | **string**| The name of the fleet | 
  **payload** | [**FleetPatchPayload**](FleetPatchPayload.md)|  | 

### Return type

[**FleetPatchResponse**](FleetPatchResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Fleets**
> FleetList Fleets(ctx, optional)
List All Fleets

List all the fleets you own.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FleetsApiFleetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FleetsApiFleetsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Page number for pagination | [default to 1]
 **limit** | **optional.Int32**| Limit of Fleet for each page | [default to 10]

### Return type

[**FleetList**](FleetList.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

