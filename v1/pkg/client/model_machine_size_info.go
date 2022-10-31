// (C) Copyright 2021-2022 Hewlett Packard Enterprise Development LP

/*
 * HPE GreenLake for bare metal API
 *
 * This Metal Client REST API provides access to bare metal as-a-service (BMaaS) within a single project context.  Clients are able to create fully-provisioned hosts, storage volumes, and project-specific private networks in an isolated project environment.  Project-owned resources that can be accessed via this API include... Host, Volume, VolumeAttachment, Network (project private), and SSH Key.   Each API call is done within a single project context.  The specific Project identifier must be provided within the header of each REST call. The server will validate that the provided authentication credentials (JWTs) are valid for the referenced project before any operation is performed.  If a single credential is valid for multiple projects, the client must still reference a single project in the header each API call. Clients can also access information about available services and resources through the AvailableResources object.  This object provides detailed  information about the OS imaging options, the machine size options, the storage volume options, data center locations, and such that are needed when creating hosts and volumes.  
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// MachineSizeInfo struct for MachineSizeInfo
type MachineSizeInfo struct {
	Name string `json:"Name,omitempty"`
	ProjectsUsed int32 `json:"ProjectsUsed,omitempty"`
	OtherUsed int32 `json:"OtherUsed,omitempty"`
	Available int32 `json:"Available,omitempty"`
	ID string `json:"ID,omitempty"`
	// True if the usage of this machine size is restricted  to selected projects, false if the machine size is available to anyone. 
	Restricted bool `json:"Restricted,omitempty"`
}
