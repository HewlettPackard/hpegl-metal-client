# NewIpPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name for the IP pool | [optional] 
**Description** | **string** | Description for the IP pool | [optional] 
**IPVersion** | [**IpVer**](IPVer.md) |  | 
**BaseIP** | **string** | Base address of the IP pool | 
**Netmask** | [**Netmask**](Netmask.md) |  | 
**DefaultRoute** | **string** | Default route associated with the IP pool | [optional] 
**Sources** | [**[]IpSource**](IPSource.md) |  | [optional] 
**DNS** | **[]string** | List of DNS servers for the IP pool | [optional] 
**Proxy** | **string** | Optional web-proxy for external internet access should the pool actually be behind a firewall | [optional] 
**NoProxy** | **string** | Addresses or CIDRs for which proxy requests are not made | [optional] 
**NTP** | **[]string** | List of NTP servers for the IP pool | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


