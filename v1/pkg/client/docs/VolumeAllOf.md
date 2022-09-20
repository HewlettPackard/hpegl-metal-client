# VolumeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | [optional] 
**FlavorID** | **string** | The VolumeFlavorID matching an entry in the VolumeFlavors array returned as part of the get /available-resources call | [optional] 
**Capacity** | **int64** | The size of the volume in KiB | [optional] 
**Shareable** | **bool** | Indicates if the volume can be attached to multiple hosts | [optional] 
**LocationID** | **string** | The location of the volume (and the storage array) LocationID is one of those listed by the LocationInfo array returned as part of the get /available-resources call. Any volumes must be in the same location as their attached Host. | [optional] 
**State** | [**VolumeState**](VolumeState.md) |  | [optional] 
**SubState** | [**VolumeSubState**](VolumeSubState.md) |  | [optional] 
**Status** | [**VolumeStatus**](VolumeStatus.md) |  | [optional] 
**Labels** | **map[string]string** | The map of label name to label value for the volume. | [optional] 
**WWN** | **string** | Serial number of the volume. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


