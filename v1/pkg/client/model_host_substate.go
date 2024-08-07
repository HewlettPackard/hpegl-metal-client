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
// HostSubstate Host substate within HostState
type HostSubstate string

// List of HostSubstate
const (
	HOSTSUBSTATE_EMPTY HostSubstate = ""
	HOSTSUBSTATE_ABORT_DEPLOY HostSubstate = "Abort Deploy"
	HOSTSUBSTATE_ALLOCATE HostSubstate = "Allocate"
	HOSTSUBSTATE_ATTACH_VOLUMES HostSubstate = "Attach Volumes"
	HOSTSUBSTATE_BOOT_SERVICE_OS HostSubstate = "Boot Service-OS"
	HOSTSUBSTATE_CLEAR_LOG HostSubstate = "Clear Log"
	HOSTSUBSTATE_COMPLETE HostSubstate = "Complete"
	HOSTSUBSTATE_CONFIRM_POST_COMPLETE HostSubstate = "Confirm Post Complete"
	HOSTSUBSTATE_CONNECT HostSubstate = "Connect"
	HOSTSUBSTATE_CONNECT_PROVISIONING HostSubstate = "Connect Provisioning"
	HOSTSUBSTATE_DNS_ADD HostSubstate = "DNS Add"
	HOSTSUBSTATE_DNS_ADD_INIT HostSubstate = "DNS Add Init"
	HOSTSUBSTATE_DNS_DELETE_INIT HostSubstate = "DNS Delete Init"
	HOSTSUBSTATE_DNS_DELETE HostSubstate = "DNS Delete"
	HOSTSUBSTATE_DEPLOY HostSubstate = "Deploy"
	HOSTSUBSTATE_DETACH_VOLUMES HostSubstate = "Detach Volumes"
	HOSTSUBSTATE_FAILED HostSubstate = "Failed"
	HOSTSUBSTATE_IN_MAINTENANCE HostSubstate = "In Maintenance"
	HOSTSUBSTATE_IN_IMAGING_PREP HostSubstate = "In Imaging Prep"
	HOSTSUBSTATE_IN_IMAGING_COMPLETE HostSubstate = "In Imaging Complete"
	HOSTSUBSTATE_INIT HostSubstate = "Init"
	HOSTSUBSTATE_INIT_ATTACH_VOLUMES HostSubstate = "Init Attach Volumes"
	HOSTSUBSTATE_INIT_DETACH_VOLUMES HostSubstate = "Init Detach Volumes"
	HOSTSUBSTATE_INIT_MAINTENANCE HostSubstate = "Init Maintenance"
	HOSTSUBSTATE_INIT_IMAGING_PREP HostSubstate = "Init Imaging Prep"
	HOSTSUBSTATE_INIT_IMAGING_COMPLETE HostSubstate = "Init Imaging Complete"
	HOSTSUBSTATE_INIT_OFF HostSubstate = "Init Off"
	HOSTSUBSTATE_ISOLATE HostSubstate = "Isolate"
	HOSTSUBSTATE_POWER_OFF HostSubstate = "Power Off"
	HOSTSUBSTATE_POWER_ON HostSubstate = "Power On"
	HOSTSUBSTATE_RELEASE HostSubstate = "Release"
	HOSTSUBSTATE_RELEASE_WITH_PROBLEM HostSubstate = "Release With Problem"
	HOSTSUBSTATE_SET_BOOT_DISK HostSubstate = "Set Boot Disk"
	HOSTSUBSTATE_SNAP_LOG HostSubstate = "Snap Log"
	HOSTSUBSTATE_SNAP_LOG_OF_FAILURE HostSubstate = "Snap Log of Failure"
	HOSTSUBSTATE_UPDATE HostSubstate = "Update"
)
