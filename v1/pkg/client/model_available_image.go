// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

/*
 * HPE GreenLake Metal Client API
 *
 * This Metal Client REST API provides access to bare metal as-a-service (BMaaS) within a single project context.  Clients are able to create fully-provisioned hosts, storage volumes, and project-specific private networks in an isolated project environment.  Project-owned resources that can be accessed via this API include... Host, Volume, VolumeAttachment, Network (project private), and SSH Key.    Each API call is done within a single project context.  The specific Project identifier must be provided within the header of each REST call. The server will validate that the provided authentication credentials (JWTs) are valid for the referenced project before any operation is performed.  If a single credential is valid for multiple projects, the client must still reference a single project in the header each API call.  Clients can also access information about available services and resources through the AvailableResources object.  This object provides detailed  information about the OS imaging options, the machine size options, the storage volume options, data center locations, and such that are needed when creating hosts and volumes.    Note: All URIs are relative to metal_service_url/rest/v1 
 *
 * API version: 1.3.6
 * Contact: quake-core@hpe.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// AvailableImage Entry describing an available imaging service
type AvailableImage struct {
	// Unique ID for imaging service
	ID string `json:"ID,omitempty"`
	// Top-level grouping of imaging services that may reference an OS or application type
	Category string `json:"Category,omitempty"`
	// Second-level grouping of imaging services.  Typically references a specific OS or application.
	Flavor string `json:"Flavor,omitempty"`
	// Specific version of a imaging service flavor.
	Version string `json:"Version,omitempty"`
	// Additional image information for additional services added to the OS
	Description string `json:"Description,omitempty"`
}
