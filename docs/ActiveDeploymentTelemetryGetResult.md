# ActiveDeploymentTelemetryGetResult

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RetrievalKey** | **string** | Unique retrieval key to get the telemetry response. | [default to null]
**Scores** | [**[]ActiveDeploymentTelemetryScore**](ActiveDeploymentTelemetryScore.md) | Result sorted by best score. Index 0 is the best one. | [default to null]
**PartialResult** | **bool** | If the score list is incomplete and missing request IDs. Can occur if you request the results before we receive telemetry from every deployment. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


