# SessionRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionId** | **string** | The Unique Identifier of the Session | [default to null]
**CustomId** | **string** | Custom ID if Available | [optional] [default to null]
**App** | **string** | The Name of the App you requested | [default to null]
**Version** | **string** | The name of the App Version you requested | [default to null]
**DeploymentRequestId** | **string** | Unique UUID | [optional] [default to null]
**Selectors** | [**[]SelectorModel**](SelectorModel.md) | List of Selectors to filter potential Deployment to link and tag the Session | [optional] [default to null]
**WebhookUrl** | **string** | When your Session is Linked, Unprocessable or in Error, we will POST the session&#39;s details on the webhook_url  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


