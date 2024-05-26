# \IPLookupApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IP**](IPLookupApi.md#IP) | **Get** /v1/ip | Get Your Public IP
[**IPLookup**](IPLookupApi.md#IPLookup) | **Get** /v1/ip/{ip}/lookup | Get an IP&#39;s information
[**IPsLookup**](IPLookupApi.md#IPsLookup) | **Post** /v1/ips/lookup | Get IPs Information in Bulk


# **IP**
> IpAddressResponse IP(ctx, )
Get Your Public IP

Retrieve your public IP address.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**IpAddressResponse**](IPAddressResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IPLookup**
> IpAddressLookupResponse IPLookup(ctx, ip)
Get an IP's information

Lookup an IP address and return the associated information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ip** | **string**|  | 

### Return type

[**IpAddressLookupResponse**](IPAddressLookupResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IPsLookup**
> IpAddressesLookupResponse IPsLookup(ctx, payload)
Get IPs Information in Bulk

Lookup IP addresses and return the associated information. Maximum of 20 IPs.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payload** | [**IpAddressesLookupPayload**](IpAddressesLookupPayload.md)|  | 

### Return type

[**IpAddressesLookupResponse**](IPAddressesLookupResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

