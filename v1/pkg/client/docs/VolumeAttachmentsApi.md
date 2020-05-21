# \VolumeAttachmentsApi

All URIs are relative to *http://localhost:8080/rest/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetByID**](VolumeAttachmentsApi.md#GetByID) | **Get** /volume-attachments/{attachmentId} | Retrieve volume attachment by ID
[**List**](VolumeAttachmentsApi.md#List) | **Get** /volume-attachments | List all volume attachments in project



## GetByID

> VolumeAttachment GetByID(ctx, attachmentId)

Retrieve volume attachment by ID

Returns a single volume attachment with matching ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attachmentId** | **string**| ID of volume attachment to return | 

### Return type

[**VolumeAttachment**](VolumeAttachment.md)

### Authorization

[quake_auth](../README.md#quake_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> []VolumeAttachment List(ctx, )

List all volume attachments in project

Returns an array of all VolumeAttachments defined within the project.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]VolumeAttachment**](VolumeAttachment.md)

### Authorization

[quake_auth](../README.md#quake_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

