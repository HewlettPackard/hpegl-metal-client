# UpdateHostAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**NetworkIDs** | **[]string** | The list of IDs corresponding to the networks that were provisioned to the host. These networks are among those listed in the Networks array returned by the get /available-resources call. | 
**NetworkForDefaultRoute** | **string** | The host&#39;s default network ID. This needs to be one of the values in the  \&quot;NetworkIDs\&quot; list. | 
**NetworkUntagged** | **string** | ID of the network selected to be untagged. This needs to be one of the values in the  \&quot;NetworkIDs\&quot; list. | 
**ISCSIConfig** | Pointer to [**UpdateHostIscsiConfig**](UpdateHostISCSIConfig.md) |  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


