# VolumeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**FlavorID** | **string** | The VolumeFlavorID matching an entry in the VolumeFlavors array returned as part of the get /available-resources call | 
**StoragePoolID** | **string** | The storage pool ID matching an entry in the StoragePools array returned as part of the get /available-resources call | 
**StoragePoolName** | **string** | Name of the storage pool from where the volume is allocated. | 
**Capacity** | **int64** | The size of the volume in KiB | 
**CapacityUsed** | **int64** | The amount of the volume currently used as reported by the array in KiB | 
**Shareable** | **bool** | Indicates if the volume can be attached to multiple hosts | 
**LocationID** | **string** | The location of the volume (and the storage array) LocationID is one of those listed by the LocationInfo array returned as part of the get /available-resources call. Any volumes must be in the same location as their attached Host. | 
**DataCenterName** | **string** | Name of the data center where the volume is created on the storage array. | 
**VolumeCollectionID** | **string** | The optional volume collection ID matching an entry in the VolumeCollections array returned as part of the get /available-resources call | [optional] 
**State** | [**VolumeState**](VolumeState.md) |  | 
**SubState** | [**VolumeSubState**](VolumeSubState.md) |  | 
**Status** | [**VolumeStatus**](VolumeStatus.md) |  | 
**Labels** | **map[string]string** | The map of label name to label value for the resource. | 
**WWN** | **string** | Serial number of the volume. | 
**ReplicationEnabled** | **bool** | Indicates whether replication is enabled for this volume. | 
**UnmanagedVolume** | **bool** | Indicates whether the volume is a native Metal created one or an external one. | 
**ActiveSite** | **string** | The site where the remote copy role for the volume is  Primary at the time of most recent import. | 
**CreatedSite** | **string** | The site where the volume was originally created.       | 
**ExportCount** | **int32** | The number of active exports for this volume | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


