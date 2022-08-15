# Allocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | **string** | Type of the resource the allocation information is listed. It is either Server or Volume | [optional] 
**Total** | **int32** | If the resource type is Server, Total is total number of servers.  If the resource type is Volume, Total  is the total storage capacity in TB | [optional] 
**InstanceFamily** | **string** | Instance type family name. | [optional] 
**InstanceType** | **string** | Server instance type or volume type | [optional] 
**Available** | **int32** | If the resource type is Servers, it is available number of servers.  If resource type is volume, it is available storage space in TB | [optional] 
**AllocationInfo** | [**[]AllocationInfo**](AllocationInfo.md) | Array listing the allocation information for each service type | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


