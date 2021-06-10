# IpPoolAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | [optional] 
**IPVersion** | [**IpVer**](IPVer.md) |  | [optional] 
**NetworkID** | **string** | Unique ID of the network associated with the IP pool | [optional] 
**BaseIP** | **string** | Base address of the IP pool | [optional] 
**Netmask** | [**Netmask**](Netmask.md) |  | [optional] 
**DefaultRoute** | **string** | Default route associated with the IP pool | [optional] 
**Sources** | [**[]IpSource**](IPSource.md) |  | [optional] 
**UseRecords** | [**[]UseRecord**](UseRecord.md) |  | [optional] 
**DNS** | **[]string** | List of DNS servers for the IP pool | [optional] 
**Proxy** | **string** | Optional web-proxy for external internet access should the pool actually be behind a firewall | [optional] 
**NoProxy** | **string** | Addresses or CIDRs for which proxy requests are not made | [optional] 
**NTP** | **[]string** | List of NTP servers for the IP pool | [optional] 
**Pool** | [**Pool**](Pool.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


