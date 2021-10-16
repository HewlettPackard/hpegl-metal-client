# \VolumesApi

All URIs are relative to *https://quake.dev.hpehcss.net/rest/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Add**](VolumesApi.md#Add) | **Post** /volumes | Add a new volume
[**Attach**](VolumesApi.md#Attach) | **Post** /volumes/{volumeId}/attach | Attach existing volume to Host
[**Delete**](VolumesApi.md#Delete) | **Delete** /volumes/{volumeId} | Delete a volume
[**Detach**](VolumesApi.md#Detach) | **Post** /volumes/{volumeId}/detach | Detach existing volume from Host
[**GetByID**](VolumesApi.md#GetByID) | **Get** /volumes/{volumeId} | Retrieve volume by ID
[**List**](VolumesApi.md#List) | **Get** /volumes | List all volumes in project
[**Update**](VolumesApi.md#Update) | **Put** /volumes/{volumeId} | Update an existing volume.  NOT SUPPORTED!!



## Add

> Volume Add(ctx, newVolume)

Add a new volume

Adds a new volume to the project.  Volumes may be created separately and then referenced in the create Host call; or volumes may be created directly within the create Host call. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**newVolume** | [**NewVolume**](NewVolume.md)| Volume that is to be added to the project | 

### Return type

[**Volume**](Volume.md)

### Authorization

[quake_auth](../README.md#quake_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Attach

> VolumeAttachment Attach(ctx, volumeId, volumeAttachHostUuid)

Attach existing volume to Host

Attaches the indicated volume to a host identified in the requestBody.   This attachment will create a VolumeAttachment object that contains  details about the connection of the volume and will update the Host  with iSCSI configuration information. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string**| ID of volume to attach | 
**volumeAttachHostUuid** | [**VolumeAttachHostUuid**](VolumeAttachHostUuid.md)| Unique ID of the Host to which the volume will be attached | 

### Return type

[**VolumeAttachment**](VolumeAttachment.md)

### Authorization

[quake_auth](../README.md#quake_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(ctx, volumeId)

Delete a volume

Deletes the volume with the matching ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string**| ID of volume to delete | 

### Return type

 (empty response body)

### Authorization

[quake_auth](../README.md#quake_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Detach

> Detach(ctx, volumeId, volumeAttachHostUuid)

Detach existing volume from Host

Detaches the indicated volume from the host identified in the requestBody.   This detachment will delete the VolumeAttachment object that contains  details about the connection of the volume and will update the Host  to remove selected iSCSI configuration information. Note that the HostID is required in the body of the request to ensure that the operation is well understood and that a volume is not accidently being removed from the wrong host. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string**| ID of volume to attach | 
**volumeAttachHostUuid** | [**VolumeAttachHostUuid**](VolumeAttachHostUuid.md)| Unique ID of the Host from which a volume will be detached | 

### Return type

 (empty response body)

### Authorization

[quake_auth](../README.md#quake_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetByID

> Volume GetByID(ctx, volumeId)

Retrieve volume by ID

Returns a single volume with matching ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string**| ID of volume to return | 

### Return type

[**Volume**](Volume.md)

### Authorization

[quake_auth](../README.md#quake_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> []Volume List(ctx, )

List all volumes in project

Returns an array of all volumes defined within the project.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]Volume**](Volume.md)

### Authorization

[quake_auth](../README.md#quake_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Volume Update(ctx, volumeId, volume)

Update an existing volume.  NOT SUPPORTED!!

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string**| ID of volume to update | 
**volume** | [**Volume**](Volume.md)| Updated volume | 

### Return type

[**Volume**](Volume.md)

### Authorization

[quake_auth](../README.md#quake_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

