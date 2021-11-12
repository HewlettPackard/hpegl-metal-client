# HostAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**ServiceID** | **string** | The image service identifier used to image the server. ServiceID is one of those listed by the Images array returned as part of the get /available-resources call. | [optional] 
**ServiceFlavor** | **string** | Overall flavor of server image used to image the server | [optional] 
**ServiceVersion** | **string** | Version of the ServiceFlavor used to image the server | [optional] 
**LocationID** | **string** | The location of the machine assigned to the host.  LocationID is one of those listed by the LocationInfo array returned as part of the get /available-resources call. | [optional] 
**MachineSizeName** | **string** | Name of the machine size used to identify and select the machine assigned to the host.  MachineSizes are described by the MachineSize array returned by the get /available-resources call. | [optional] 
**MachineSizeID** | **string** | UniqueID referring to the machine size used to identify and select the machine assigned to the host.  MachineSizes are described by the MachineSize array returned by the get /available-resources call. | [optional] 
**MachineID** | **string** | UniqueID referring to the machine on which this host is running. | [optional] 
**SSHKeyIDs** | **[]string** | IDs of SSH Keys used when configuring the Host | [optional] 
**SSHAuthorizedKeys** | **[]string** | Specific SSH keys that were when configuring the host. | [optional] 
**NetworkIDs** | **[]string** | The list of IDs corresponding to the networks that were provisioned to the host. These networks are among those listed in the Networks array returned by the get /available-resources call. | 
**NetworkForDefaultRoute** | **string** | The host default network ID | 
**UserData** | **string** | User-provided data attached to the image configuration data when the host was provisioned | [optional] 
**NodeID** | **string** | User-provided data to represent the identity of the host within an application environment. For example, this could be set to represent the Kubernetes node ID if the host is provisioned as a Kubernetes node. | [optional] 
**ISCSIConfig** | [**HostIscsiConfig**](HostISCSIConfig.md) |  | [optional] 
**Connections** | [**[]HostConnection**](HostConnection.md) | Details describing host network connections | [optional] 
**Deleted** | **bool** | True if the Host has been deleted. | [optional] 
**PortalCommOkay** | **bool** | Describes if the portal is in active communication to the device | [optional] 
**PowerStatus** | [**HostPowerState**](HostPowerState.md) |  | [optional] 
**State** | [**HostState**](HostState.md) |  | [optional] 
**Substate** | [**HostSubstate**](HostSubstate.md) |  | [optional] 
**StateTime** | [**time.Time**](time.Time.md) |  | [optional] 
**SubstateTime** | [**time.Time**](time.Time.md) |  | [optional] 
**Progress** | **int64** |  | [optional] 
**Alert** | **bool** |  | [optional] 
**AlertInfo** | [**[]HostAlertInfo**](HostAlertInfo.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


