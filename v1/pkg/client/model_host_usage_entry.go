// (C) Copyright 2021-2022 Hewlett Packard Enterprise Development LP

/*
 * HPE GreenLake Metal Client API
 *
 * This Metal Client REST API provides access to bare metal as-a-service (BMaaS) within a single project context.  Clients are able to create fully-provisioned hosts, storage volumes, and project-specific private networks in an isolated project environment.  Project-owned resources that can be accessed via this API include... Host, Volume, VolumeAttachment, Network (project private), and SSH Key.    Each API call is done within a single project context.  The specific Project identifier must be provided within the header of each REST call. The server will validate that the provided authentication credentials (JWTs) are valid for the referenced project before any operation is performed.  If a single credential is valid for multiple projects, the client must still reference a single project in the header each API call.  Clients can also access information about available services and resources through the AvailableResources object.  This object provides detailed  information about the OS imaging options, the machine size options, the storage volume options, data center locations, and such that are needed when creating hosts and volumes.    Note: All URIs are relative to metal_service_url/rest/v1 
 *
 * API version: 1.3.7
 * Contact: quake-core@hpe.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
import (
	"time"
)
// HostUsageEntry struct for HostUsageEntry
type HostUsageEntry struct {
	// Project ID that contained the host
	ProjectID string `json:"ProjectID,omitempty"`
	// The location ID is the data center location of the resource.  The LocationID must be one of those provided by the LocationInfo array returned as part of the get /available-resources call.  The locations are typically described by country, region, and data center.
	LocationID string `json:"LocationID,omitempty"`
	// Timestamp of when resource (machine or storage) was allocated
	Allocated time.Time `json:"Allocated,omitempty"`
	// Timestamp of when resource (host or volume) was ready for use
	Ready time.Time `json:"Ready,omitempty"`
	// Timestamp of when resource (machine or storage) was freed
	Freed time.Time `json:"Freed,omitempty"`
	// The start of the usage reporting window or when the resource was allocated
	UsageStart time.Time `json:"UsageStart,omitempty"`
	// The end of the usage reporting window or when the resource was freed
	UsageEnd time.Time `json:"UsageEnd,omitempty"`
	// The difference between the UsageEnd and UsageStart rounded up to the UsageHours
	UsageHours int64 `json:"UsageHours,omitempty"`
	// Description of error that affected the usage reporting
	Error string `json:"Error,omitempty"`
	// Name of the MachineSize requested when host was created
	MachineSizeName string `json:"MachineSizeName,omitempty"`
	// Unique ID of the MachineSize requested when host was created
	MachineSizeID string `json:"MachineSizeID,omitempty"`
	// Name of the associated Host
	HostName string `json:"HostName,omitempty"`
	// Unique ID of the associated Host
	HostID string `json:"HostID,omitempty"`
}
