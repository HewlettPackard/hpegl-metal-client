# HostAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**ServiceID** | **string** | The image service identifier used to image the server. ServiceID is one of those listed by the Images array returned as part of the get /available-resources call. | 
**ServiceFlavor** | **string** | Overall flavor of server image used to image the server | 
**ServiceVersion** | **string** | Version of the ServiceFlavor used to image the server | 
**LocationID** | **string** | The location of the machine assigned to the host.  LocationID is one of those listed by the LocationInfo array returned as part of the get /available-resources call. | 
**MachineSizeName** | **string** | Name of the machine size used to identify and select the machine assigned to the host.  MachineSizes are described by the MachineSize array returned by the get /available-resources call. | 
**MachineSizeID** | **string** | UniqueID referring to the machine size used to identify and select the machine assigned to the host.  MachineSizes are described by the MachineSize array returned by the get /available-resources call. | 
**MachineID** | **string** | UniqueID referring to the machine on which this host is running. | 
**SSHKeyIDs** | **[]string** | IDs of SSH Keys used when configuring the Host | 
**SSHAuthorizedKeys** | **[]string** | Specific SSH keys that were when configuring the host. | 
**NetworkIDs** | **[]string** | The list of IDs corresponding to the networks that were provisioned to the host. These networks are among those listed in the Networks array returned by the get /available-resources call. | 
**NetworkForDefaultRoute** | **string** | The host&#39;s default network ID. This needs to be one of the values in the  \&quot;NetworkIDs\&quot; list. | 
**NetworkUntagged** | **string** | ID of the network selected to be untagged. This needs to be one of the values in the  \&quot;NetworkIDs\&quot; list. | 
**PreAllocatedIPs** | **[]string** | The list of pre-allocated IP addresses corresponding to the list of NetworkIDs. Pre-allocated IP addresses are optional, but required when updating a host containing Pre-allocated IP addresses. | 
**ServiceNetsProviderMAC** | **map[string]string** | The map of Service Network (Provider) ID to Provider MAC address.   The Service Network must be a provider network provisioned to this host. Any Service Networks not included here will default to the physical MAC learned during machine discovery. | 
**UserData** | **string** | User-provided data attached to the image configuration data when the host was provisioned | 
**NodeID** | **string** | User-provided data to represent the identity of the host within an application environment. For example, this could be set to represent the Kubernetes node ID if the host is provisioned as a Kubernetes node. | 
**ISCSIConfig** | Pointer to [**HostIscsiConfig**](HostISCSIConfig.md) |  | 
**Connections** | [**[]HostConnection**](HostConnection.md) | Details describing host network connections | 
**Deleted** | **bool** | True if the Host has been deleted. | 
**PortalCommOkay** | **bool** | Describes if the portal is in active communication to the device | 
**PowerStatus** | [**HostPowerState**](HostPowerState.md) |  | 
**State** | [**HostState**](HostState.md) |  | 
**Substate** | [**HostSubstate**](HostSubstate.md) |  | 
**StateTime** | [**time.Time**](time.Time.md) |  | 
**SubstateTime** | [**time.Time**](time.Time.md) |  | 
**Progress** | **int64** |  | 
**Alert** | **bool** |  | 
**AlertInfo** | [**[]HostAlertInfo**](HostAlertInfo.md) |  | 
**Workflow** | **string** | The current workflow the host is in | 
**SummaryStatus** | [**HealthStatus**](HealthStatus.md) |  | 
**Labels** | **map[string]string** | The map of label name to label value for the resource. | 
**WWPNs** | **[]string** | FC HBA world wide port names | 
**FWBaselineID** | **string** | The ID of the firmware baseline that is installed on the host. | 
**FWBaselineVersion** | **string** | The version of the firmware baseline that is installed on the host. | 
**AvailableFWBaselineID** | **string** | The ID of the firmware baseline that is available for the host. This field is only populated when there is a new firmware baseline available. | 
**AvailableFWBaselineVersion** | **string** | The version of the firmware baseline that is available for the host. This field is only populated when there is a new firmware baseline available. | 
**SerialNumber** | **string** | The serial number of the host. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


