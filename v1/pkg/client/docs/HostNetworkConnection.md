# HostNetworkConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the network connection | 
**NetworkID** | **string** | Unique ID corresponding to a network in the local data center | 
**IP** | **string** | IP address for the network connection | 
**Subnet** | **string** | The IP subnet address | 
**Netmask** | **string** |  | 
**Gateway** | **string** | The IP subnet gateway address | 
**DNS** | **[]string** | List of DNS servers for the IP subnet | 
**VLAN** | **int32** | VLAN ID of the network | 
**VNI** | **int32** | VNI ID of the network | 
**Untagged** | **bool** | True if the network is untagged | 
**Proxy** | **string** | Optional web-proxy for external internet access should the IP subnet actually be behind a firewall | 
**NoProxy** | **string** | Addresses or CIDRs for which proxy requests are not made | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


