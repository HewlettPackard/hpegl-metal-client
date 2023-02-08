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
// VolumeInfo Description of the details about a pre-existing volume
type VolumeInfo struct {
	// Unique ID for the volume.  Referenced if the volume is to be attached to a new Host 
	ID string `json:"ID,omitempty"`
	// User-provided name for the volume
	Name string `json:"Name,omitempty"`
	// User-provided description of the volume
	Description string `json:"Description,omitempty"`
	// The ID of the volume flavor requested when creating the volume
	FlavorID string `json:"FlavorID,omitempty"`
	// The size of the volume in KiB
	Capacity int64 `json:"Capacity,omitempty"`
	// Indicates if the volume can be attached to multiple hosts
	Shareable bool `json:"Shareable,omitempty"`
	// The location ID of the data center where the volume exists
	LocationID string `json:"LocationID,omitempty"`
	// The discovery IP for the iSCSI volume
	DiscoveryIP string `json:"DiscoveryIP,omitempty"`
	// The target IQN for the iSCSI volume
	TargetIQN string `json:"TargetIQN,omitempty"`
	State VolumeState `json:"State,omitempty"`
	Status VolumeStatus `json:"Status,omitempty"`
}
