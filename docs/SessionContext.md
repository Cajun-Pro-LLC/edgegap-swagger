# SessionContext

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionId** | **string** | Unique UUID | [default to null]
**CustomId** | **string** | Custom ID if Available | [optional] [default to null]
**Status** | **string** | Current status of the session | [default to null]
**Ready** | **bool** | If the session is linked to a Ready deployment | [default to null]
**Linked** | **bool** | If the session is linked to a deployment | [default to null]
**Kind** | **string** | Type of session created | [default to null]
**UserCount** | **int32** | Count of user this session currently have | [default to null]
**DeploymentRequestId** | **string** | Unique UUID | [optional] [default to null]
**WebhookUrl** | **string** | When your Session is Linked, Unprocessable or in Error, we will POST the session&#39;s details on the webhook_url  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


