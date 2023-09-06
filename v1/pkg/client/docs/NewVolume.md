# NewVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | **string** |  | [optional] 
**FlavorID** | **string** | Adds a new volume to the project.  This object requires the LocationID and is used when a new volume is created independently from the host creation therefore requiring a specified location. | 
**Capacity** | **int64** | The size of the volume in GiB | 
**Shareable** | **bool** | Indicates if the volume can be attached to multiple hosts | [optional] 
**LocationID** | **string** | The location of the volume (and the storage array) LocationID is one of those listed by the LocationInfo array returned as part of the get /available-resources call. Any volumes must be in the same location as their attached Host. | 
**Labels** | **map[string]string** | The map of service/user specified label name to label value for this volume. Setting service labels is restricted by role. | [optional] 
**StoragePoolID** | **string** | The storage pool is one of those listed by the StoragePools array returned as part of the get /available-resources call that are available to create volumes of the specified flavor and location. | [optional] 
**VolumeCollectionID** | **string** | The  optional volume collection is one of those listed by the VolumeCollections  array returned as part of the get /available-resources call that are available to create volumes of the specified flavor and location. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


