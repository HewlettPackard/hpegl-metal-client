# \VolumeAttachmentsApi

All URIs are relative to *https://client.greenlake.hpe.com/api/metal/rest/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Add**](VolumeAttachmentsApi.md#Add) | **Post** /volume-attachments | Create a new VolumeAttachment
[**Delete**](VolumeAttachmentsApi.md#Delete) | **Delete** /volume-attachments/{attachmentId} | Delete a VolumeAttachment
[**GetByID**](VolumeAttachmentsApi.md#GetByID) | **Get** /volume-attachments/{attachmentId} | Retrieve volume attachment by ID
[**List**](VolumeAttachmentsApi.md#List) | **Get** /volume-attachments | List all volume attachments in project



## Add

> VolumeAttachment Add(ctx, newVolumeAttachment)

Create a new VolumeAttachment

Adds a new VolumeAttachment which enables a machine to use a volume. Note that this API is for creation of a VolumeAttachment for non-Hosts. A VolumeAttachment resource must be deleted using the 'DELETE /volume-attachments/{attachmentId}' API. For regular hosts, a VolumeAttachment is created via the 'POST /volumes/{volumeId}/attach' API.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**newVolumeAttachment** | [**NewVolumeAttachment**](NewVolumeAttachment.md)| NewVolumeAttachement parameters to create a new VolumeAttachment. | 

### Return type

[**VolumeAttachment**](VolumeAttachment.md)

### Authorization

[BearerAuth](../README.md#BearerAuth), [Membership](../README.md#Membership), [Project](../README.md#Project)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(ctx, attachmentId)

Delete a VolumeAttachment

Deletes the VolumeAttachment with the matching ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attachmentId** | **string**| ID of VolumeAttachment to delete | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth), [Membership](../README.md#Membership), [Project](../README.md#Project)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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

[BearerAuth](../README.md#BearerAuth), [Membership](../README.md#Membership), [Project](../README.md#Project)

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

[BearerAuth](../README.md#BearerAuth), [Membership](../README.md#Membership), [Project](../README.md#Project)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

