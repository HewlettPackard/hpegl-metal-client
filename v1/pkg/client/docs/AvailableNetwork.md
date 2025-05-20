# AvailableNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | **string** | ID of the network | 
**Name** | **string** | The name of the network | 
**LocationID** | **string** | The location ID is the data center location of the resource.  The LocationID must be one of those provided by the LocationInfo array returned as part of the get /available-resources call.  The locations are typically described by country, region, and data center. | 
**Description** | **string** |  | 
**HostUse** | [**NetworkHostUse**](NetworkHostUse.md) |  | 
**Purpose** | [**NetworkPurpose**](NetworkPurpose.md) |  | 
**IPPoolID** | **string** |  | 
**NoIPPool** | **bool** | True if the Network does not have an associated IP Pool. | 
**VLAN** | **int32** | VLAN ID of the network | 
**VNI** | **int32** | VNI ID of the network | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


