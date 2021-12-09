# NewVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | **string** |  | [optional] 
**FlavorID** | **string** | Adds a new volume to the project.  This object requires the LocationID and is used when a new volume is created independently from the host creation therefore requiring a specified location. | 
**Capacity** | **int64** | The size of the volume in GiB | 
**LocationID** | **string** | The location of the volume (and the storage array) LocationID is one of those listed by the LocationInfo array returned as part of the get /available-resources call. Any volumes must be in the same location as their attached Host. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


