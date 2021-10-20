# VolumeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | [optional] 
**FlavorID** | **string** | The VolumeFlavorID matching an entry in the VolumeFlavors array returned as part of the get /available-resources call | [optional] 
**Capacity** | **uint64** | The size of the volume in MB | [optional] 
**LocationID** | **string** | The location of the volume (and the storage array) LocationID is one of those listed by the LocationInfo array returned as part of the get /available-resources call. Any volumes must be in the same location as their attached Host. | [optional] 
**State** | [**VolumeState**](VolumeState.md) |  | [optional] 
**Status** | [**VolumeStatus**](VolumeStatus.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


