# \LocationsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LocationBeaconList**](LocationsApi.md#LocationBeaconList) | **Get** /v1/locations/beacons | List All Location Beacons
[**LocationsGet**](LocationsApi.md#LocationsGet) | **Get** /v1/locations | List All Locations


# **LocationBeaconList**
> LocationBeaconList LocationBeaconList(ctx, )
List All Location Beacons

List all the active location beacons. They can be used to ping them for your matchmaking system. You cannot deploy on beacons.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**LocationBeaconList**](LocationBeaconList.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationsGet**
> Locations LocationsGet(ctx, optional)
List All Locations

List all the locations available to deploy on. You can specify an application and a version to filter out the locations that don’t have enough resources to deploy this application version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LocationsApiLocationsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocationsApiLocationsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app** | **optional.String**| The App Name you want to filter with capacity | 
 **version** | **optional.String**| The Version Name you want to filter with capacity | 
 **type_** | **optional.String**| The type of the location | 
 **tags** | **optional.String**| Gets locations with tags. Set to: \&quot;true\&quot; to have the tags | 

### Return type

[**Locations**](Locations.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

