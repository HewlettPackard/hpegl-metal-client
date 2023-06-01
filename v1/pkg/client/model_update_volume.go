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
// UpdateVolume struct for UpdateVolume
type UpdateVolume struct {
	// Unique ID for the resource instance as generated by the Metal service
	ID string `json:"ID"`
	// Used to determine whether the DB entry has changed since it was last read. This value is updated each time the resource is updated.  Client must send this value unchanged for any update operation.
	ETag string `json:"ETag"`
	// Common name for the resource instance. Must be 128 or fewer printable characters
	Name string `json:"Name"`
	// The size of the volume in KiB
	Capacity int64 `json:"Capacity"`
	// The map of service/user specified label name to label value for this volume. Setting this field is restricted by role. At minimum, this field should be set to the values specified during create.
	Labels map[string]string `json:"Labels"`
}
