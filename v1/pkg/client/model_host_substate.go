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
	HOSTSUBSTATE_CONFIRM_OFF HostSubstate = "Confirm Off"
	HOSTSUBSTATE_CONFIRM_ON HostSubstate = "Confirm On"
	HOSTSUBSTATE_CONNECT HostSubstate = "Connect"
	HOSTSUBSTATE_CONNECT_PROVISIONING HostSubstate = "Connect Provisioning"
	HOSTSUBSTATE_DEPLOY HostSubstate = "Deploy"
	HOSTSUBSTATE_DETACH_VOLUMES HostSubstate = "Detach Volumes"
	HOSTSUBSTATE_ERROR_RECOVERY HostSubstate = "Error Recovery"
	HOSTSUBSTATE_FAIL_CLEANUP HostSubstate = "Fail Cleanup"
	HOSTSUBSTATE_FAILED HostSubstate = "Failed"
	HOSTSUBSTATE_INIT HostSubstate = "Init"
	HOSTSUBSTATE_INIT_OFF HostSubstate = "Init Off"
	HOSTSUBSTATE_ISOLATE HostSubstate = "Isolate"
	HOSTSUBSTATE_IN_MAINTENANCE HostSubstate = "In Maintenance"
	HOSTSUBSTATE_POWER_OFF HostSubstate = "Power Off"
	HOSTSUBSTATE_POWER_ON HostSubstate = "Power On"
	HOSTSUBSTATE_RELEASE HostSubstate = "Release"
	HOSTSUBSTATE_RELEASE_WITH_PROBLEM HostSubstate = "Release With Problem"
	HOSTSUBSTATE_SET_BOOT_DISK HostSubstate = "Set Boot Disk"
	HOSTSUBSTATE_SNAP_LOG HostSubstate = "Snap Log"
	HOSTSUBSTATE_UPDATE HostSubstate = "Update"
)
