# FleetPoliciesPatchPayload

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the Policy | [optional] [default to null]
**Enabled** | **bool** | If the Policy is enabled. Defaults to true. | [optional] [default to null]
**Type_** | **string** | Type of the Policy. See our documentation for more information on the type and it&#39;s data. | [optional] [default to null]
**Minimum** | **int32** | Minimum deployments at all time | [optional] [default to null]
**Maximum** | **int32** | Maximum deployments for the provided type. | [optional] [default to null]
**Threshold** | **float32** | Based on the number of sockets connected, how filled should a session deployment be before initiating a scale-up deployment. Float between 0.1 and 1. | [optional] [default to null]
**Data** | **interface{}** | JSON object for your filters. See our documentation for more information. | [optional] [default to null]
**CreateTime** | **string** | UTC time of policy creation | [optional] [default to null]
**LastUpdated** | **string** | UTC time of policy last update | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


