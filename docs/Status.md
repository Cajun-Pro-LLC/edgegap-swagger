# Status

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | The Unique ID of the Deployment&#39;s request | [default to null]
**Fqdn** | **string** | The FQDN that allow to connect to your Deployment | [default to null]
**AppName** | **string** | The name of the deployed App | [default to null]
**AppVersion** | **string** | The version of the deployed App | [default to null]
**CurrentStatus** | **string** | The current status of the Deployment | [default to null]
**Running** | **bool** | True if the current Deployment is ready to be connected and running | [default to null]
**WhitelistingActive** | **bool** | True if the current Deployment is ACL protected | [default to null]
**StartTime** | **string** | Timestamp of the Deployment when it is up and running | [default to null]
**RemovalTime** | **string** | Timestamp of the end of the Deployment | [optional] [default to null]
**ElapsedTime** | **int32** | Time since the Deployment is up and running in seconds | [default to null]
**LastStatus** | **string** | The last status of the Deployment | [optional] [default to null]
**Error_** | **bool** | True if there is an error with the Deployment | [default to null]
**ErrorDetail** | **string** | The error details of the Deployment | [optional] [default to null]
**Ports** | [**map[string]PortMapping**](PortMapping.md) |  | [optional] [default to null]
**PublicIp** | **string** | The public IP | [default to null]
**Sessions** | [**[]DeploymentSessionContext**](DeploymentSessionContext.md) | List of Active Sessions if Deployment App is Session Based | [optional] [default to null]
**Location** | [***DeploymentLocation**](DeploymentLocation.md) | Location related information | [optional] [default to null]
**Tags** | **[]string** | List of tags associated with the deployment | [optional] [default to null]
**Sockets** | **int32** | The Capacity of the Deployment | [optional] [default to null]
**SocketsUsage** | **int32** | The Capacity Usage of the Deployment | [optional] [default to null]
**Command** | **string** | The command to use in the container, null mean it will take the default of the container | [optional] [default to null]
**Arguments** | **string** | The arguments to use in the container, null mean it will take the default of the container | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


