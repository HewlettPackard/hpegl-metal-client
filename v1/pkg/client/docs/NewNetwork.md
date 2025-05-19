# NewNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the network | 
**LocationID** | **string** | The location ID is the data center location of the resource.  The LocationID must be one of those provided by the LocationInfo array returned as part of the get /available-resources call.  The locations are typically described by country, region, and data center. | 
**Description** | **string** |  | [optional] 
**HostUse** | [**NetworkHostUse**](NetworkHostUse.md) |  | [optional] 
**Purpose** | [**NetworkPurpose**](NetworkPurpose.md) |  | [optional] 
**NewIPPool** | Pointer to [**NewIpPool**](NewIPPool.md) |  | [optional] 
**NoIPPool** | **bool** | True if the Network does not have an associated IP Pool. | [optional] 
**VLAN** | **int32** | Optional VLAN ID of the network. If not specified, it is allocated from reserved pool. | [optional] 
**VNI** | **int32** | Optional VNI ID of the network. If not specified, it is allocated from reserved pool if required. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


