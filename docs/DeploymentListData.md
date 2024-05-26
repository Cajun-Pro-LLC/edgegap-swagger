# DeploymentListData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Unique UUID | [default to null]
**Fqdn** | **string** | The FQDN that allow to connect to your deployment | [default to null]
**StartTime** | **string** | Timestamp of the deployment when it is up and running | [default to null]
**Ready** | **bool** | If the deployment is ready | [default to null]
**PublicIp** | **string** | The public IP | [default to null]
**Ports** | [**map[string]PortMapping**](PortMapping.md) |  | [optional] [default to null]
**Tags** | **[]string** | List of tags associated with the deployment | [optional] [default to null]
**Sockets** | **int32** | The capacity of the deployment | [optional] [default to null]
**SocketsUsage** | **int32** | The capacity usage of the deployment | [optional] [default to null]
**IsJoinableBySession** | **bool** | If the deployment is joinable by sessions | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


