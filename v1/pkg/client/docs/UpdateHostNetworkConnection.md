# UpdateHostNetworkConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the network connection | [optional] 
**NetworkID** | **string** | Unique ID corresponding to a network in the local data center | [optional] 
**IP** | **string** | IP address for the network connection | [optional] 
**Subnet** | **string** | The IP subnet address | [optional] 
**Netmask** | **string** |  | [optional] 
**Gateway** | **string** | The IP subnet gateway address | [optional] 
**DNS** | **[]string** | List of DNS servers for the IP subnet | [optional] 
**VLAN** | **int32** | VLAN ID of the network | [optional] 
**VNI** | **int32** | VNI ID of the network | [optional] 
**Untagged** | **bool** | True if the network is untagged | [optional] 
**Proxy** | **string** | Optional web-proxy for external internet access should the IP subnet actually be behind a firewall | [optional] 
**NoProxy** | **string** | Addresses or CIDRs for which proxy requests are not made | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


