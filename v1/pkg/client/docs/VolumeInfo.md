# VolumeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | **string** | Unique ID for the volume.  Referenced if the volume is to be attached to a new Host  | 
**Name** | **string** | User-provided name for the volume | 
**Description** | **string** | User-provided description of the volume | 
**FlavorID** | **string** | The ID of the volume flavor requested when creating the volume | 
**StoragePoolID** | **string** | The ID of the storage pool where volume resides | 
**Capacity** | **int64** | The size of the volume in KiB | 
**Shareable** | **bool** | Indicates if the volume can be attached to multiple hosts | 
**LocationID** | **string** | The location ID of the data center where the volume exists | 
**DiscoveryIP** | **string** | The discovery IP for the iSCSI volume | 
**TargetIQN** | **string** | The target IQN for the iSCSI volume | 
**State** | [**VolumeState**](VolumeState.md) |  | 
**Status** | [**VolumeStatus**](VolumeStatus.md) |  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


