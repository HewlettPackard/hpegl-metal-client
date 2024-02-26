# IpPoolAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**IPVersion** | [**IpVer**](IPVer.md) |  | 
**NetworkID** | **string** | Unique ID of the network associated with the IP pool | 
**BaseIP** | **string** | Base address of the IP pool | 
**Netmask** | [**Netmask**](Netmask.md) |  | 
**DefaultRoute** | **string** | Default route associated with the IP pool | 
**Sources** | [**[]IpSource**](IPSource.md) |  | 
**UseRecords** | [**[]UseRecord**](UseRecord.md) |  | 
**DNS** | **[]string** | List of DNS servers for the IP pool | 
**Proxy** | **string** | Optional web-proxy for external internet access should the pool actually be behind a firewall | 
**NoProxy** | **string** | Addresses or CIDRs for which proxy requests are not made | 
**NTP** | **[]string** | List of NTP servers for the IP pool | 
**Pool** | [**Pool**](Pool.md) |  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


