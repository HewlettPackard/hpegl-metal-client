// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

/*
 * Quake Project Client API
 *
 * This Quake client REST API provides access to bare metal as-a-service (BMaaS) within a single project context.  Clients are able to create fully-provisioned hosts, storage volumes, and project-specific private networks in an isolated project environment.  Project-owned resources that can be accessed via this API include... Host, Volume, VolumeAttachment, Network (project private), and SSH Key.    Each API call is done within a single project context.  The specific Project identifier must be provided within the header of each REST call. The server will validate that the provided authentication credentials (JWTs) are valid for the referenced project before any operation is performed.  If a single credential is valid for multiple projects, the client must still reference a single project in the header each API call.  Clients can also access information about available services and resources through the AvailableResources object.  This object provides detailed  information about the OS imaging options, the machine size options, the storage volume options, data center locations, and such that are needed when creating hosts and volumes. 
 *
 * API version: 1.3.1
 * Contact: chuck.hudson@hpe.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
import (
	"time"
)
// IpPoolUpdate struct for IpPoolUpdate
type IpPoolUpdate struct {
	// Unique ID for the resource instance as generated by the Quake service
	ID string `json:"ID,omitempty"`
	// Used to determine whether the DB entry has changed since it was last read.  should be returned unchanged on any update operation.
	ETag string `json:"ETag,omitempty"`
	// Common name for the resource instance
	Name string `json:"Name,omitempty"`
	// Time when the resource was created in the database
	Created time.Time `json:"Created,omitempty"`
	// Time when the resource was last modified in the database
	Modified time.Time `json:"Modified,omitempty"`
	// Description for the IP pool
	Description string `json:"Description,omitempty"`
	// Default route for the IP pool
	DefaultRoute string `json:"DefaultRoute,omitempty"`
	// List of DNS servers for the IP pool
	DNS []string `json:"DNS,omitempty"`
	// Optional web-proxy for external internet access should the pool actually be behind a firewall
	Proxy string `json:"Proxy,omitempty"`
	// Addresses or CIDRs for which proxy requests are not made
	NoProxy string `json:"NoProxy,omitempty"`
	// List of NTP servers for the IP pool
	NTP []string `json:"NTP,omitempty"`
}
