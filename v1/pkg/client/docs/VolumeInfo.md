# VolumeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | **string** | Unique ID for the volume.  Referenced if the volume is to be attached to a new Host  | [optional] 
**Name** | **string** | User-provided name for the volume | [optional] 
**Description** | **string** | User-provided description of the volume | [optional] 
**FlavorID** | **string** | The ID of the volume flavor requested when creating the volume | [optional] 
**Capacity** | **int64** | The size of the volume in KiB | [optional] 
**Shareable** | **bool** | Indicates if the volume can be attached to multiple hosts | [optional] 
**LocationID** | **string** | The location ID of the data center where the volume was created | [optional] 
**State** | [**VolumeState**](VolumeState.md) |  | [optional] 
**Status** | [**VolumeStatus**](VolumeStatus.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


