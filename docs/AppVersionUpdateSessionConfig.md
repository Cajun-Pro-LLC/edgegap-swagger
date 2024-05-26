# AppVersionUpdateSessionConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | The kind of session to create. If &#39;Default&#39; if chosen, the application will be handled like a normal application. The kind of session must be: Default, Seat, Match | [optional] [default to null]
**Sockets** | **int32** | The number of game slots on each deployment of this app version. | [optional] [default to null]
**Autodeploy** | **bool** | If a deployment should be made autonomously if there is not enough sockets open on a new session. | [optional] [default to null]
**EmptyTtl** | **int32** | The number of minutes a deployment of this app version can spend with no session connected before being terminated. | [optional] [default to null]
**SessionMaxDuration** | **int32** | The number of minutes after a session-type deployment has been terminated to remove all the session information connected to your deployment. Minimum and default value is set to 60 minutes so you can manage your session termination before it is removed. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


