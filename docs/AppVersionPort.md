# AppVersionPort

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | **int32** | The Port to Expose your service. Port 0 reserved for one-to-one port mapping. See our doc for more information. | [default to null]
**Protocol** | **string** | Available protocols: TCP, UDP, TCP/UDP, HTTP, HTTPS, WS or WSS | [default to null]
**ToCheck** | **bool** | If the port must be verified by our port validations | [optional] [default to null]
**TlsUpgrade** | **bool** | Enabling with HTTP or WS will inject a sidecar proxy that upgrades the connection with TLS | [optional] [default to null]
**Name** | **string** | An optional name for the port for easier handling. Mandatory if using port 0 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


