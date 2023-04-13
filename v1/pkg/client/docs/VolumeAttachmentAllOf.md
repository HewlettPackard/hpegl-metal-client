# VolumeAttachmentAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VolumeID** | **string** | Unique ID of the volume attached to the host | 
**HostID** | **string** | Unique ID of the host connected to the volume | 
**HostIPAddress** | **string** | Host IP address for the network connection that connects to the storage array | 
**IQN** | **string** | IQN is the full initiator name used for identification during iSCSI login | 
**CHAPSecret** | **string** | CHAPSecret is the Challenge Authentication Protocol secret to be shared between array and initiator. | 
**CHAPUserName** | **string** | CHAPUserName is the CHAP username to use for CHAP authentication | 
**LUN** | **int32** | LUN is the Logical Unit Number to be assigned to the volume on export | 
**VolumeTargetIQN** | **string** | VolumeTargetIQN is the iQN for the volume, assigned by the array corresponding to the volume | 
**VolumeTargetIPAddress** | **string** | VolumeTargetIPAddress is the IPV4 address of the iSCSI volume export | 
**State** | [**VaStateEnum**](VaStateEnum.md) |  | 
**FSConfig** | Pointer to [**VafsConfig**](VAFSConfig.md) | File share specific configuration parameters | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


