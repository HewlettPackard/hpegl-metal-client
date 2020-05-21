# AvailableResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Images** | [**[]AvailableImage**](AvailableImage.md) | Array listing the available host imaging services | [optional] 
**Locations** | [**[]LocationInfo**](LocationInfo.md) | Array listing the data center locations with available resources | [optional] 
**Networks** | [**[]AvailableNetwork**](AvailableNetwork.md) | Array listing the networks available for host connections | [optional] 
**MachineSizes** | [**[]MachineSize**](MachineSize.md) | Array listing the available machine (server) sizes | [optional] 
**VolumeFlavors** | [**[]VolumeFlavor**](VolumeFlavor.md) | Array listing the available volume flavors | [optional] 
**Volumes** | [**[]VolumeInfo**](VolumeInfo.md) | Array listing the existing project volumes that could be attached to a host | [optional] 
**MachineInventory** | [**[]MachineInventory**](MachineInventory.md) | Array listing the number of machines of each size in each location | [optional] 
**StorageInventory** | [**[]StorageInventory**](StorageInventory.md) | Array providing information on the amount of available storage of each flavor in each location | [optional] 
**SSHKeys** | [**[]SshKeyEntry**](SSHKeyEntry.md) | Array listing pre-defined SSH keys that could be referenced when creating a Host | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


