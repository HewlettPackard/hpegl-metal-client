// (C) Copyright 2021-2023 Hewlett Packard Enterprise Development LP

/*
 * HPE GreenLake for bare metal API
 *
 * This Metal Client REST API provides access to bare metal as-a-service (BMaaS) within a single project context. Clients are able to create fully-provisioned hosts, storage volumes, and project-specific private networks in an isolated project environment.   Project-owned resources that can be accessed via this API include - Host,  Volume, VolumeAttachment, Network (project private), and SSH Key. Each API  call is done within a single project context. The specific Project identifier  must be provided within the header of for each REST call. The server will  validate that the provided authentication credentials (JWTs) are valid  for  the referenced project before any operation is performed. If a single credential is valid for multiple projects, the client must still reference a single project  in the header for each API call.    Clients can also access information about available services and resources through the AvailableResources object. This object provides detailed  information about the OS imaging options, the machine size options, the storage volume options, and data center locations which are needed when creating hosts and volumes.  
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// OsServiceImageAllOf struct for OsServiceImageAllOf
type OsServiceImageAllOf struct {
	Description string `json:"Description"`
	// A high level classification of the service. 
	Category string `json:"Category"`
	// An high level description of the OS image type. 
	Flavor string `json:"Flavor"`
	// A specific version of the service flavor.
	Version string `json:"Version"`
	Origin OsServiceImageOrigin `json:"Origin"`
	// Maximum amount of time in seconds for the service to complete its work.  This time includes the amount of time to download the service image if necessary.
	Timeout int64 `json:"Timeout"`
	// Device partitioning information.
	DeviceLayouts []DiskPartitions `json:"DeviceLayouts"`
	// List of projects allowed to use this OS service image.
	PermittedProjects []string `json:"PermittedProjects"`
	// Files and relative information for this OS service image.
	Files []FileInfo `json:"Files"`
	// Defines a list of embedded contents to be attached to this OS service image in form of files.
	Info []PassedInfo `json:"Info"`
	Approach OsServiceImageApproach `json:"Approach"`
	AssumedBootMethod BootMethod `json:"AssumedBootMethod"`
	// Indicates if a LAG needs to be created on the switch connecting the host with this OS service image.
	NoSwitchLAG bool `json:"NoSwitchLAG"`
	BondMode BondMode `json:"BondMode"`
	// Unique id of the firmware baseline to be applied by the firmware update operation  in Maintenance Steps and Imaging Prep Steps.
	FWBaselineID string `json:"FWBaselineID"`
	UserDefinedSteps UserDefinedSteps `json:"UserDefinedSteps"`
	// Allows setting of restricted use for this service by machine attributes.  If rules are not defined, this OS service image is available to be installed on any machine.
	Classifiers []MachineClassifier `json:"Classifiers"`
}
