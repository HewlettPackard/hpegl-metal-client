# Partition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**PartitionType**](PartitionType.md) |  | 
**Formula** | [**PartitionFormula**](PartitionFormula.md) |  | 
**Name** | **string** | Name is an optional name for the partition, only usable with GPT partition table types. | 
**MountPoint** | **string** | Where to mount the partition, it must be set for partitions that are filesystem types. | 
**PartitionID** | **int64** | A partitionID may be explicitly set for the partition and this allows for out-of-order partitions in the table (like VMWare).  If zero, sequential partitions are created starting at 1. Gaps in the partition ID numbering is also allowable (VMWare). | 
**Bootable** | **bool** | Identify a bootable partition. | 
**MakeOptions** | **[]string** | Provide options for the partition&#39;s creation. | 
**MountOptions** | **string** | Provide options for the patition mount. If empty, \&quot;defaults\&quot; will be used for filesystem partition types. M MountOptions will be placed in the \&quot;options\&quot; field for the partition in /etc/fstab. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


