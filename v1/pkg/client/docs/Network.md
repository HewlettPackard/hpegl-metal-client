# Network

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | **string** | Unique ID for the resource instance as generated by the Metal service | 
**ETag** | **string** | Used to determine whether the DB entry has changed since it was last read. This value is updated each time the resource is updated.  Client must send this value unchanged for any update operation. | 
**Name** | **string** | Common name for the resource instance. Must be 128 or fewer printable characters | 
**Created** | [**time.Time**](time.Time.md) | Time when the resource was created in the database | 
**Modified** | [**time.Time**](time.Time.md) | Time when the resource was last modified in the database | 
**LocationID** | **string** | The location ID is the data center location of the resource.  The LocationID must be one of those provided by the LocationInfo array returned as part of the get /available-resources call.  The locations are typically described by country, region, and data center. | 
**Description** | **string** |  | 
**HostUse** | [**NetworkHostUse**](NetworkHostUse.md) |  | 
**Purpose** | [**NetworkPurpose**](NetworkPurpose.md) |  | 
**IPPoolID** | **string** |  | 
**NoIPPool** | **bool** | True if the Network does not have an associated IP Pool. | 
**VLAN** | **int32** | VLAN ID of the network | 
**VNI** | **int32** | VNI ID of the network | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


