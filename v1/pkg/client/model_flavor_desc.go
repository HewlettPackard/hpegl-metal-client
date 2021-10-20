// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

/*
 * HPE GreenLake Metal Client API
 *
 * This Metal Client REST API provides access to bare metal as-a-service (BMaaS) within a single project context.  Clients are able to create fully-provisioned hosts, storage volumes, and project-specific private networks in an isolated project environment.  Project-owned resources that can be accessed via this API include... Host, Volume, VolumeAttachment, Network (project private), and SSH Key.    Each API call is done within a single project context.  The specific Project identifier must be provided within the header of each REST call. The server will validate that the provided authentication credentials (JWTs) are valid for the referenced project before any operation is performed.  If a single credential is valid for multiple projects, the client must still reference a single project in the header each API call.  Clients can also access information about available services and resources through the AvailableResources object.  This object provides detailed  information about the OS imaging options, the machine size options, the storage volume options, data center locations, and such that are needed when creating hosts and volumes. 
 *
 * API version: 1.3.3
 * Contact: quake-core@hpe.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// FlavorDesc Describes the details of machine size or volume flavor
type FlavorDesc struct {
	// Groups flavors together to make finding a specific flavor easier
	Collection string `json:"Collection,omitempty"`
	Banner1 string `json:"Banner1,omitempty"`
	Banner2 string `json:"Banner2,omitempty"`
	Bullets []string `json:"Bullets,omitempty"`
	// URI to more information about the specific machine size or volume flavor
	InfoLink string `json:"InfoLink,omitempty"`
	// Supplemental tooltip text to use in GUIs
	Tooltip string `json:"Tooltip,omitempty"`
}
