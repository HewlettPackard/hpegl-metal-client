// (C) Copyright 2021-2024 Hewlett Packard Enterprise Development LP

/*
 * HPE GreenLake for bare metal API
 *
 * This Metal Client REST API provides access to bare metal as-a-service (BMaaS) within a single project context. Clients are able to create fully-provisioned hosts, storage volumes, and project-specific private networks in an isolated project environment.   Project-owned resources that can be accessed via this API include - Host,  Volume, VolumeAttachment, Network (project private), and SSH Key. Each API  call is done within a single project context. The specific Project identifier  must be provided within the header of for each REST call. The server will  validate that the provided authentication credentials (JWTs) are valid  for  the referenced project before any operation is performed. If a single credential is valid for multiple projects, the client must still reference a single project  in the header for each API call.    Clients can also access information about available services and resources through the AvailableResources object. This object provides detailed  information about the OS imaging options, the machine size options, the storage volume options, and data center locations which are needed when creating hosts and volumes.  
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// Limits struct for Limits
type Limits struct {
	// Maximum number of hosts to allow
	Hosts int32 `json:"Hosts"`
	// Maximum number of volumes to allow
	Volumes int32 `json:"Volumes"`
	// Maximum capacity to allow in TiB
	VolumeCapacity int64 `json:"VolumeCapacity"`
	// Maximum number of private networks to allow
	PrivateNetworks int32 `json:"PrivateNetworks"`
	// Map of instance type ID to maximum number of hosts that can be created with that instance type
	InstanceTypes map[string]int32 `json:"InstanceTypes"`
}
