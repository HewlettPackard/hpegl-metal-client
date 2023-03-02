# UpdateHostAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**NetworkIDs** | **[]string** | The list of IDs corresponding to the networks that were provisioned to the host. These networks are among those listed in the Networks array returned by the get /available-resources call. | 
**NetworkForDefaultRoute** | **string** | The host&#39;s default network ID. This needs to be one of the values in the  \&quot;NetworkIDs\&quot; list. | 
**NetworkUntagged** | **string** | ID of the network selected to be untagged. This needs to be one of the values in the  \&quot;NetworkIDs\&quot; list. | 
**ISCSIConfig** | Pointer to [**UpdateHostIscsiConfig**](UpdateHostISCSIConfig.md) |  | [optional] 
**PreAllocatedIPs** | **[]string** | The list of pre-allocated IP addresses corresponding to the list of NetworkIDs. Pre-allocated IP addresses are optional, but required when updating a host containing Pre-allocated IP addresses. | 
**ServiceNetsProviderMAC** | **map[string]string** | The map of Service Network (Provider) ID to Provider MAC address.   The Service Network must be a provider network provisioned to this host. Any Service Networks not included here will default to the physical MAC learned during machine discovery. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


