# RelaySessionUserBaseResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionId** | **string** | The session ID | [default to null]
**AuthorizationToken** | **int32** | The authorization token for the session | [optional] [default to null]
**Status** | **string** | The status of the session | [default to null]
**Ready** | **bool** | If the session is ready to be used | [default to null]
**Linked** | **bool** | If the session is linked to a relay | [default to null]
**Error_** | **string** | The error message if the session failed | [optional] [default to null]
**SessionUser** | [***RelaySessionUserBaseResponseSessionUser**](RelaySessionUserBaseResponse_session_user.md) |  | [optional] [default to null]
**Relay** | [***RelaySessionBaseResponseRelay**](RelaySessionBaseResponse_relay.md) |  | [optional] [default to null]
**WebhookUrl** | **string** | The webhook URL that we will call once the session is ready | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


