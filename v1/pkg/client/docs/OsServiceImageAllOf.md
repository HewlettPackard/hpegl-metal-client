# OsServiceImageAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**Category** | **string** | A high level classification of the service.  | 
**Flavor** | **string** | An high level description of the OS image type.  | 
**Version** | **string** | A specific version of the service flavor. | 
**Origin** | [**OsServiceImageOrigin**](OSServiceImageOrigin.md) |  | 
**Timeout** | **int64** | Maximum amount of time in seconds for the service to complete its work.  This time includes the amount of time to download the service image if necessary. | 
**DeviceLayouts** | [**[]DiskPartition**](DiskPartition.md) | Device partitioning information. | 
**PermittedProjects** | **[]string** | List of projects allowed to use this OS service image. | 
**Files** | [**[]FileInfo**](FileInfo.md) | Files and relative information for this OS service image. | 
**Info** | [**[]PassedInfo**](PassedInfo.md) | Defines a list of embedded contents to be attached to this OS service image if form of files. | 
**Approach** | [**OsServiceImageApproach**](OSServiceImageApproach.md) |  | 
**AssumedBootMethod** | [**BootMethod**](BootMethod.md) |  | 
**NoSwitchLAG** | **bool** | Indicates if a LAG needs to be created on the switch connecting the host with this OS service image. | 
**BondMode** | [**BondMode**](BondMode.md) |  | 
**FWBaselineID** | **string** | Unique id of the firmware baseline to be applied by the firmware update operation  in Maintenance Steps and Imaging Prep Steps. | 
**UserDefinedSteps** | [**UserDefinedSteps**](UserDefinedSteps.md) |  | 
**Classifiers** | [**[]MachineClassifier**](MachineClassifier.md) | Allows setting of restricted use for this service by machine attributes.  If rules are not defined, this OS service image is available to be installed on any machine. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


