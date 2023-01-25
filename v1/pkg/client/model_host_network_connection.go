// (C) Copyright 2021-2022 Hewlett Packard Enterprise Development LP

/*
 * test HPE GreenLake for bare metal API
 *
 * This Metal Client REST API provides access to bare metal as-a-service (BMaaS) within a single project context. Clients are able to create fully-provisioned hosts, storage volumes, and project-specific private networks in an isolated project environment.   Project-owned resources that can be accessed via this API include - Host,  Volume, VolumeAttachment, Network (project private), and SSH Key. Each API  call is done within a single project context. The specific Project identifier  must be provided within the header of for each REST call. The server will  validate that the provided authentication credentials (JWTs) are valid  for  the referenced project before any operation is performed. If a single credential is valid for multiple projects, the client must still reference a single project  in the header for each API call.    Clients can also access information about available services and resources through the AvailableResources object. This object provides detailed  information about the OS imaging options, the machine size options, the storage volume options, and data center locations which are needed when creating hosts and volumes.  
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// HostNetworkConnection struct for HostNetworkConnection
type HostNetworkConnection struct {
	// Name of the network connection
	Name string `json:"Name,omitempty"`
	// Unique ID corresponding to a network in the local data center
	NetworkID string `json:"NetworkID,omitempty"`
	// IP address for the network connection
	IP string `json:"IP,omitempty"`
	// The IP subnet address
	Subnet string `json:"Subnet,omitempty"`
	Netmask string `json:"Netmask,omitempty"`
	// The IP subnet gateway address
	Gateway string `json:"Gateway,omitempty"`
	// List of DNS servers for the IP subnet
	DNS []string `json:"DNS,omitempty"`
	// VLAN ID of the network
	VLAN int32 `json:"VLAN,omitempty"`
	// VNI ID of the network
	VNI int32 `json:"VNI,omitempty"`
	// True if the network is untagged
	Untagged bool `json:"Untagged,omitempty"`
	// Optional web-proxy for external internet access should the IP subnet actually be behind a firewall
	Proxy string `json:"Proxy,omitempty"`
	// Addresses or CIDRs for which proxy requests are not made
	NoProxy string `json:"NoProxy,omitempty"`
}
