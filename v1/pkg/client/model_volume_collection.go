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
// VolumeCollection Volume collection within a data center location
type VolumeCollection struct {
	// The ID of the volume collection resource that can be used while creating a volume
	ID string `json:"ID"`
	// The name of the volume collection resource
	Name string `json:"Name"`
	// The location ID of the data center where the volume collection resides
	LocationID string `json:"LocationID"`
	// The description of the volume collection resource
	Description string `json:"Description"`
	// IDs of the valid storage pools that are valid for this volume collection as listed in the Storage Pools array
	StoragePoolIDs []string `json:"StoragePoolIDs,omitempty"`
}
