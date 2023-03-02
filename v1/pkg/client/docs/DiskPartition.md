# DiskPartition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | **string** | Indicates the device name of the disk of the form /dev/sdX as known by the os e.g. /dev/sda This is the default device that will be used to deploy the OS Image to the host. The default value may be overridden when creating the host by a UUID/NAA/EUI  logical volume ID (BootDeviceID). | 
**Description** | **string** |  | 
**StartOffset** | **int64** | StartOffset is the offset from the start of the disk to the first partition,  if zero a default of 1 MiB (2048 sectors) will be used, units bytes. | 
**TableType** | [**PartitionTable**](PartitionTable.md) |  | 
**Partitions** | [**Partition**](Partition.md) |  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


