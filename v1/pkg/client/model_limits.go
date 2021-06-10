/*
 * Quake Project Client API
 *
 * This Quake client REST API provides access to bare metal as-a-service (BMaaS) within a single project context.  Clients are able to create fully-provisioned hosts, storage volumes, and project-specific private networks in an isolated project environment.  Project-owned resources that can be accessed via this API include... Host, Volume, VolumeAttachment, Network (project private), and SSH Key.    Each API call is done within a single project context.  The specific Project identifier must be provided within the header of each REST call. The server will validate that the provided authentication credentials (JWTs) are valid for the referenced project before any operation is performed.  If a single credential is valid for multiple projects, the client must still reference a single project in the header each API call.  Clients can also access information about available services and resources through the AvailableResources object.  This object provides detailed  information about the OS imaging options, the machine size options, the storage volume options, data center locations, and such that are needed when creating hosts and volumes. 
 *
 * API version: 1.2.1
 * Contact: chuck.hudson@hpe.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// Limits struct for Limits
type Limits struct {
	// Maximum number of hosts to allow
	Hosts uint32 `json:"Hosts,omitempty"`
	// Maximum number of volumes to allow
	Volumes uint32 `json:"Volumes,omitempty"`
	// Maximum capacity to allow in GiB
	VolumeCapacity uint64 `json:"VolumeCapacity,omitempty"`
}
