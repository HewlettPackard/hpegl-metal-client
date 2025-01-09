// (C) Copyright 2021-2024 Hewlett Packard Enterprise Development LP

/*
 * HPE GreenLake for bare metal API
 *
 * This Metal Client REST API provides access to bare metal as-a-service (BMaaS) within a single project context. Clients are able to create fully-provisioned hosts, storage volumes, and project-specific private networks in an isolated project environment.   Project-owned resources that can be accessed via this API include - Host,  Volume, VolumeAttachment, Network (project private), and SSH Key. Each API  call is done within a single project context. The specific Project identifier  must be provided within the header of for each REST call. The server will  validate that the provided authentication credentials (JWTs) are valid  for  the referenced project before any operation is performed. If a single credential is valid for multiple projects, the client must still reference a single project  in the header for each API call.    Clients can also access information about available services and resources through the AvailableResources object. This object provides detailed  information about the OS imaging options, the machine size options, the storage volume options, and data center locations which are needed when creating hosts and volumes.  
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// HostState Overall host state
type HostState string

// List of HostState
const (
	HOSTSTATE_EMPTY HostState = ""
	HOSTSTATE_NEW HostState = "New"
	HOSTSTATE_DELETING HostState = "Deleting"
	HOSTSTATE_DELETED HostState = "Deleted"
	HOSTSTATE_FAILED HostState = "Failed"
	HOSTSTATE_FORCE_DELETING HostState = "Force Deleting"
	HOSTSTATE_UPDATING_CONNECTIONS HostState = "Updating Connections"
	HOSTSTATE_IMAGING HostState = "Imaging"
	HOSTSTATE_IMAGING_PREP HostState = "Imaging Prep"
	HOSTSTATE_CONNECTING HostState = "Connecting"
	HOSTSTATE_BOOTING HostState = "Booting"
	HOSTSTATE_READY HostState = "Ready"
	HOSTSTATE_OFFLINE HostState = "Offline"
	HOSTSTATE_REIMAGING_PREP HostState = "Reimaging Prep"
	HOSTSTATE_REPLACING HostState = "Replacing"
	HOSTSTATE_RELEASING HostState = "Releasing"
	HOSTSTATE_ALLOCATING HostState = "Allocating"
	HOSTSTATE_MAINTENANCE HostState = "Maintenance"
	HOSTSTATE_ATTACHING HostState = "Attaching"
	HOSTSTATE_DETACHING HostState = "Detaching"
	HOSTSTATE_ISCSI_ATTACHING HostState = "ISCSI Attaching"
	HOSTSTATE_ALL_DETACHING HostState = "All Detaching"
)
