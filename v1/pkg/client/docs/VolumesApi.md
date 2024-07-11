# \VolumesApi

All URIs are relative to *https://client.greenlake.hpe.com/api/metal/rest/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Add**](VolumesApi.md#Add) | **Post** /volumes | Add a new volume
[**Attach**](VolumesApi.md#Attach) | **Post** /volumes/{volumeId}/attach | Attach existing volume to Host
[**Delete**](VolumesApi.md#Delete) | **Delete** /volumes/{volumeId} | Delete a volume
[**Detach**](VolumesApi.md#Detach) | **Post** /volumes/{volumeId}/detach | Detach existing volume from Host
[**GetByID**](VolumesApi.md#GetByID) | **Get** /volumes/{volumeId} | Retrieve volume by ID
[**List**](VolumesApi.md#List) | **Get** /volumes | List all volumes in project
[**Update**](VolumesApi.md#Update) | **Put** /volumes/{volumeId} | Update an existing volume



## Add

> Volume Add(ctx, newVolume, optional)

Add a new volume

Adds a new volume to the project. Volumes may be created separately and then referenced in the create Host call; or volumes may be created directly within the create Host call. If GreenLake Platform IAM issued token is used for authentication, then it is required to pass  'X-Role' and 'X-Workspaceid' headers. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**newVolume** | [**NewVolume**](NewVolume.md)| Volume that is to be added to the project | 
 **optional** | ***AddOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRole** | **optional.String**| GreenLake Platform role name | 
 **xWorkspaceid** | **optional.String**| GreenLake Platform workspace ID | 

### Return type

[**Volume**](Volume.md)

### Authorization

[BearerAuth](../README.md#BearerAuth), [Membership](../README.md#Membership), [Project](../README.md#Project)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Attach

> VolumeAttachment Attach(ctx, volumeId, volumeAttachHostUuid, optional)

Attach existing volume to Host

Attaches the indicated volume to a host identified in the requestBody.   This attachment will create a VolumeAttachment object that contains  details about the connection of the volume and will update the Host  with iSCSI configuration information. If GreenLake Platform IAM issued token is used for authentication, then it is required to pass  'X-Role' and 'X-Workspaceid' headers. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string**| ID of volume to attach | 
**volumeAttachHostUuid** | [**VolumeAttachHostUuid**](VolumeAttachHostUuid.md)| Unique ID of the Host to which the volume will be attached | 
 **optional** | ***AttachOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AttachOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRole** | **optional.String**| GreenLake Platform role name | 
 **xWorkspaceid** | **optional.String**| GreenLake Platform workspace ID | 

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

> Delete(ctx, volumeId, optional)

Delete a volume

Deletes the volume with the matching ID. If GreenLake Platform IAM issued token is used for authentication, then it is required to pass  'X-Role' and 'X-Workspaceid' headers.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string**| ID of volume to delete | 
 **optional** | ***DeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRole** | **optional.String**| GreenLake Platform role name | 
 **xWorkspaceid** | **optional.String**| GreenLake Platform workspace ID | 

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


## Detach

> Detach(ctx, volumeId, volumeAttachHostUuid, optional)

Detach existing volume from Host

Detaches the indicated volume from the host identified in the requestBody.   This detachment will delete the VolumeAttachment object that contains  details about the connection of the volume and will update the Host  to remove selected iSCSI configuration information. Note that the HostID is required in the body of the request to ensure that the operation is well understood and that a volume is not accidently being removed from the wrong host. If GreenLake Platform IAM issued token is used for authentication, then it is required to pass  'X-Role' and 'X-Workspaceid' headers. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string**| ID of volume to attach | 
**volumeAttachHostUuid** | [**VolumeAttachHostUuid**](VolumeAttachHostUuid.md)| Unique ID of the Host from which a volume will be detached | 
 **optional** | ***DetachOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DetachOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRole** | **optional.String**| GreenLake Platform role name | 
 **xWorkspaceid** | **optional.String**| GreenLake Platform workspace ID | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth), [Membership](../README.md#Membership), [Project](../README.md#Project)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetByID

> Volume GetByID(ctx, volumeId, optional)

Retrieve volume by ID

Returns a single volume with matching imaged. If GreenLake Platform IAM issued token is used for authentication, then it is required to pass  'X-Role' and 'X-Workspaceid' headers.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string**| ID of volume to return | 
 **optional** | ***GetByIDOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetByIDOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRole** | **optional.String**| GreenLake Platform role name | 
 **xWorkspaceid** | **optional.String**| GreenLake Platform workspace ID | 

### Return type

[**Volume**](Volume.md)

### Authorization

[BearerAuth](../README.md#BearerAuth), [Membership](../README.md#Membership), [Project](../README.md#Project)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> []Volume List(ctx, optional)

List all volumes in project

Returns an array of all volumes defined within the project. If GreenLake Platform IAM issued token is used for authentication, then it is required to pass  'X-Role' and 'X-Workspaceid' headers.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRole** | **optional.String**| GreenLake Platform role name | 
 **xWorkspaceid** | **optional.String**| GreenLake Platform workspace ID | 

### Return type

[**[]Volume**](Volume.md)

### Authorization

[BearerAuth](../README.md#BearerAuth), [Membership](../README.md#Membership), [Project](../README.md#Project)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Volume Update(ctx, volumeId, updateVolume, optional)

Update an existing volume

Updates volume with matching ID. Update is permitted only when volume is in 'Allocated' or 'Visible' state. Only the Volume 'Capacity' can be updated with a value greater than the existing one to expand the volume.  If GreenLake Platform IAM issued token is used for authentication, then it is required to pass  'X-Role' and 'X-Workspaceid' headers. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string**| ID of volume to return | 
**updateVolume** | [**UpdateVolume**](UpdateVolume.md)| Volume object with its ID and Capacity in GiB indicating the expanded size to be specified. | 
 **optional** | ***UpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRole** | **optional.String**| GreenLake Platform role name | 
 **xWorkspaceid** | **optional.String**| GreenLake Platform workspace ID | 

### Return type

[**Volume**](Volume.md)

### Authorization

[BearerAuth](../README.md#BearerAuth), [Membership](../README.md#Membership), [Project](../README.md#Project)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

