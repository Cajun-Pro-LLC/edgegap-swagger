# PullProfilePostResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the pull profile | [default to null]
**Source** | **string** | Source in the S3 bucket to fetch from | [default to null]
**SourceType** | **string** | If the source is a File or a Directory | [default to null]
**Destination** | **string** | Destination path where your source will be uploaded in your container. Make sure to avoid protected destinations, such as &#x60;/etc/&#x60;, as this will prevent the files from being copied to your deployment, and will make your deployment fail. Make sure a normal user can write to the destination folder. | [default to null]
**CreateTime** | **string** | UTC time of pull profile creation | [optional] [default to null]
**LastUpdated** | **string** | UTC time of pull profile last update | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


