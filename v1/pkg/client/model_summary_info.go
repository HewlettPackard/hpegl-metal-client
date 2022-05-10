// (C) Copyright 2021-2022 Hewlett Packard Enterprise Development LP

/*
 * HPE GreenLake Metal Client API
 *
 * This Metal Client REST API provides access to bare metal as-a-service (BMaaS) within a single project context.  Clients are able to create fully-provisioned hosts, storage volumes, and project-specific private networks in an isolated project environment.  Project-owned resources that can be accessed via this API include... Host, Volume, VolumeAttachment, Network (project private), and SSH Key.    Each API call is done within a single project context.  The specific Project identifier must be provided within the header of each REST call. The server will validate that the provided authentication credentials (JWTs) are valid for the referenced project before any operation is performed.  If a single credential is valid for multiple projects, the client must still reference a single project in the header each API call.  Clients can also access information about available services and resources through the AvailableResources object.  This object provides detailed  information about the OS imaging options, the machine size options, the storage volume options, data center locations, and such that are needed when creating hosts and volumes.    Note: All URIs are relative to metal_service_url/rest/v1 
 *
 * API version: 1.3.8
 * Contact: quake-core@hpe.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// SummaryInfo struct for SummaryInfo
type SummaryInfo struct {
	// Locations where resources exist for all listed projects
	Locations []LocationInfo `json:"Locations,omitempty"`
	// Number of hosts for all listed projects
	NumHosts int32 `json:"NumHosts,omitempty"`
	// Number of volumes for all listed projects
	NumVolumes int32 `json:"NumVolumes,omitempty"`
	// Total storage for all listed projects
	TotalStorageGB int32 `json:"TotalStorageGB,omitempty"`
	// Number of projects with OK health summary status
	NumOK int32 `json:"NumOK,omitempty"`
	// Number of projects with Warning health summary status
	NumWarning int32 `json:"NumWarning,omitempty"`
	// Number of projects with Critical health summary status
	NumCritical int32 `json:"NumCritical,omitempty"`
	// Number of projects with Unknown health summary status
	NumUnknown int32 `json:"NumUnknown,omitempty"`
}