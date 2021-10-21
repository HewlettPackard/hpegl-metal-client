# NewHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The host name used when installing the host operating system.  Note that some OS implementations may require the name to be formatted as a fully qualified domain name. | 
**Description** | **string** |  | [optional] 
**ServiceID** | **string** | The image service identifier must be identifer of one of the available imaging services provided by the AvailableImage array returned as part of the get /available-resources call. Images are typically described by category (e.g. Linux), flavor (e.g. ubuntu) and version (e.g. 18.04-20190807) | 
**LocationID** | **string** | The location ID must be one of those provided by the LocationInfo array returned as part of the get /available-resources call.  The locations are typically described by country, region, and data center. The location ID must also be one that has sufficient inventory for the selected MachineSize.  See the MachineInventory array as returned by the get /available-resources call to select a location that has machines available with the correct machine size. | 
**MachineSizeID** | **string** | The machine size ID must be one of those listed in the MachineSizes array returned as part of the get /available-resources call.  This array provides the name and detailed description for each machine size.  Use the MachineInventory array provided by the get /available-resources call to find a location with an adequate inventory of machines with the desired machine size. | 
**MachineID** | **string** | The machine ID is normally not used. | [optional] 
**SSHKeyIDs** | **[]string** | These IDs must correspond to the IDs for SSH keys already created within the context of the project.  These SSH keys and those included as SSHAuthorizedKeys will be used together to provision SSH keys when the machine is imaged. | [optional] 
**NetworkIDs** | **[]string** | The list of IDs corresponding to the networks that will be provisioned to the host. These networks must be among those listed in the Networks array returned by the get /available-resources call.  This list must include the ID for any network identified by the HostUseEnum as Required. | 
**PreAllocatedIPs** | **[]string** | A list of pre-allocated IP addresses to be used for corresponding networks. This array of IP addresses, if present, is in one-to-one correspondence with Networks. | [optional] 
**NetworkForDefaultRoute** | **string** | The host default network ID | 
**VolumeIDs** | **[]string** | The list of IDs corresponding to existing, unattached volumes that should be attached to the new host.  The volume must be one of those listed in the Volumes array returned by the get /available-resources call.  In addition, the volume must be in the visible state and not a part of any current VolumeAttachment (see get volume-attachments) | [optional] 
**NewVolumes** | [**[]AddVolume**](AddVolume.md) | New volumes may be created and connected to the Host when the host is provisioned. The information provided here to create a host is the same as required when doing a post /volumes call | [optional] 
**UserData** | **string** | User-provided data to be attached to the image configuration data. | [optional] 
**NodeID** | **string** | User-provided data to represent the identity of the host within an application environment. For example, this could be set to represent the Kubernetes node ID if the host is provisioned as a Kubernetes node. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


