# \ApplicationsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppVersionDelete**](ApplicationsApi.md#AppVersionDelete) | **Delete** /v1/app/{app_name}/version/{version_name} | Delete an Application Version
[**AppVersionGet**](ApplicationsApi.md#AppVersionGet) | **Get** /v1/app/{app_name}/version/{version_name} | Get an Application Version
[**AppVersionPost**](ApplicationsApi.md#AppVersionPost) | **Post** /v1/app/{app_name}/version | Create a New Application Version
[**AppVersionWhitelistEntryDelete**](ApplicationsApi.md#AppVersionWhitelistEntryDelete) | **Delete** /v1/app/{app_name}/version/{version_name}/whitelist/{entry_id} | Delete an ACL Entry
[**AppVersionWhitelistEntryGet**](ApplicationsApi.md#AppVersionWhitelistEntryGet) | **Get** /v1/app/{app_name}/version/{version_name}/whitelist/{entry_id} | Get an ACL Entry
[**AppVersionWhitelistGet**](ApplicationsApi.md#AppVersionWhitelistGet) | **Get** /v1/app/{app_name}/version/{version_name}/whitelist | List All ACL Entries for an Application Version
[**AppVersionWhitelistPost**](ApplicationsApi.md#AppVersionWhitelistPost) | **Post** /v1/app/{app_name}/version/{version_name}/whitelist | Create an ACL Entry
[**AppVersionsGet**](ApplicationsApi.md#AppVersionsGet) | **Get** /v1/app/{app_name}/versions | List All Versions for an Application
[**AppVersionsPatch**](ApplicationsApi.md#AppVersionsPatch) | **Patch** /v1/app/{app_name}/version/{version_name} | Update an Application Version
[**ApplicationDelete**](ApplicationsApi.md#ApplicationDelete) | **Delete** /v1/app/{app_name} | Delete an Application
[**ApplicationGet**](ApplicationsApi.md#ApplicationGet) | **Get** /v1/app/{app_name} | Get an Application
[**ApplicationPatch**](ApplicationsApi.md#ApplicationPatch) | **Patch** /v1/app/{app_name} | Update an Application
[**ApplicationPost**](ApplicationsApi.md#ApplicationPost) | **Post** /v1/app | Create a New Application
[**ApplicationsGet**](ApplicationsApi.md#ApplicationsGet) | **Get** /v1/apps | List All Applications


# **AppVersionDelete**
> AppVersionDelete AppVersionDelete(ctx, appName, versionName)
Delete an Application Version

Delete a specific version of an application.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appName** | **string**| The name of the application | 
  **versionName** | **string**| The name of the application version | 

### Return type

[**AppVersionDelete**](AppVersionDelete.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppVersionGet**
> AppVersionPayload AppVersionGet(ctx, appName, versionName)
Get an Application Version

Retrieve the specifications of an application version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appName** | **string**| The name of the application | 
  **versionName** | **string**| The name of the application version | 

### Return type

[**AppVersionPayload**](AppVersionPayload.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppVersionPost**
> AppVersionCreateResponse AppVersionPost(ctx, appName, payload)
Create a New Application Version

Create an application version associated with an application. The version contains all the specifications to create a deployment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appName** | **string**| The name of the application associated | 
  **payload** | [**AppVersionPayload**](AppVersionPayload.md)|  | 

### Return type

[**AppVersionCreateResponse**](AppVersionCreateResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppVersionWhitelistEntryDelete**
> AppVersionWhitelistEntrySuccess AppVersionWhitelistEntryDelete(ctx, appName, versionName, entryId)
Delete an ACL Entry

Delete an access control list entry for a specific application version

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appName** | **string**| The name of the application | 
  **versionName** | **string**| The name of the application version | 
  **entryId** | **string**| The unique ID of the entry | 

### Return type

[**AppVersionWhitelistEntrySuccess**](AppVersionWhitelistEntrySuccess.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppVersionWhitelistEntryGet**
> AppVersionWhitelistEntry AppVersionWhitelistEntryGet(ctx, appName, versionName, entryId)
Get an ACL Entry

Retrieve a specific access control list entry for an application version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appName** | **string**| The name of the application | 
  **versionName** | **string**| The name of the application version | 
  **entryId** | **string**| The unique ID of the entry | 

### Return type

[**AppVersionWhitelistEntry**](AppVersionWhitelistEntry.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppVersionWhitelistGet**
> AppVersionWhitelistResponse AppVersionWhitelistGet(ctx, appName, versionName)
List All ACL Entries for an Application Version

List all the access control list entries for a specific application version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appName** | **string**| The name of the application | 
  **versionName** | **string**| The name of the application version | 

### Return type

[**AppVersionWhitelistResponse**](AppVersionWhitelistResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppVersionWhitelistPost**
> AppVersionWhitelistEntrySuccess AppVersionWhitelistPost(ctx, appName, versionName, payload)
Create an ACL Entry

Create an access control list entry for an app version. This will allow the specified CIDR to connect to the deployment. The option ```whitelisting_active``` must be activated in the application version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appName** | **string**| The name of the application | 
  **versionName** | **string**| The name of the application version | 
  **payload** | [**AppVersionWhitelistEntryPayload**](AppVersionWhitelistEntryPayload.md)|  | 

### Return type

[**AppVersionWhitelistEntrySuccess**](AppVersionWhitelistEntrySuccess.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppVersionsGet**
> AppVersionList AppVersionsGet(ctx, appName)
List All Versions for an Application

List all versions of a specific application.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appName** | **string**| The name of the application | 

### Return type

[**AppVersionList**](AppVersionList.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppVersionsPatch**
> AppVersionUpdateResponse AppVersionsPatch(ctx, appName, versionName, payload)
Update an Application Version

Update an application version with new specifications.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appName** | **string**| The name of the application | 
  **versionName** | **string**| The name of the application version | 
  **payload** | [**AppVersionUpdatePayload**](AppVersionUpdatePayload.md)|  | 

### Return type

[**AppVersionUpdateResponse**](AppVersionUpdateResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationDelete**
> ApplicationDelete(ctx, appName)
Delete an Application

Delete an application and all its current versions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appName** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationGet**
> Application ApplicationGet(ctx, appName)
Get an Application

Retrieve an application and its information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appName** | **string**|  | 

### Return type

[**Application**](Application.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationPatch**
> Application ApplicationPatch(ctx, appName, payload)
Update an Application

Update an application with new information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appName** | **string**|  | 
  **payload** | [**ApplicationPatch**](ApplicationPatch.md)|  | 

### Return type

[**Application**](Application.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationPost**
> Application ApplicationPost(ctx, payload)
Create a New Application

Create an application that will regroup application versions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payload** | [**ApplicationPost**](ApplicationPost.md)|  | 

### Return type

[**Application**](Application.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationsGet**
> Applications ApplicationsGet(ctx, )
List All Applications

List all the applications that you own.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Applications**](Applications.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

