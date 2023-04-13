// (C) Copyright 2021-2023 Hewlett Packard Enterprise Development LP

/*
 * HPE GreenLake for bare metal API
 *
 * This Metal Client REST API provides access to bare metal as-a-service (BMaaS) within a single project context. Clients are able to create fully-provisioned hosts, storage volumes, and project-specific private networks in an isolated project environment.   Project-owned resources that can be accessed via this API include - Host,  Volume, VolumeAttachment, Network (project private), and SSH Key. Each API  call is done within a single project context. The specific Project identifier  must be provided within the header of for each REST call. The server will  validate that the provided authentication credentials (JWTs) are valid  for  the referenced project before any operation is performed. If a single credential is valid for multiple projects, the client must still reference a single project  in the header for each API call.    Clients can also access information about available services and resources through the AvailableResources object. This object provides detailed  information about the OS imaging options, the machine size options, the storage volume options, and data center locations which are needed when creating hosts and volumes.  
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// UserOpEnum Defines user operation for a machine.
type UserOpEnum string

// List of UserOpEnum
const (
	USEROPENUM_RESET_SERVER UserOpEnum = "Reset Server"
	USEROPENUM_POWER_ON UserOpEnum = "Power On"
	USEROPENUM_POWER_OFF UserOpEnum = "Power Off"
	USEROPENUM_BOOT_SERVICE_OS UserOpEnum = "Boot Service OS"
	USEROPENUM_RE_DISCOVER UserOpEnum = "Re-Discover"
	USEROPENUM_BURN_IN_SERVER UserOpEnum = "Burn-In Server"
	USEROPENUM_ERASE_DRIVE UserOpEnum = "Erase Drive"
	USEROPENUM_UPDATE_FIRMWARE UserOpEnum = "Update Firmware"
	USEROPENUM_SET_BIOS_PARAMETERS UserOpEnum = "Set BIOS Parameters"
	USEROPENUM_SET_RAID0_BOOT_VOLUME UserOpEnum = "Set RAID0 Boot Volume"
	USEROPENUM_SET_RAID1_BOOT_VOLUME UserOpEnum = "Set RAID1 Boot Volume"
	USEROPENUM_CLEAR_RAID_CONFIGURATION UserOpEnum = "Clear RAID Configuration"
	USEROPENUM_PAUSE UserOpEnum = "Pause"
	USEROPENUM_UPDATE_BMC_FIRMWARE UserOpEnum = "Update BMC Firmware"
	USEROPENUM_DISCOVER UserOpEnum = "Discover"
	USEROPENUM_UPDATE_BMC_LICENSE UserOpEnum = "Update BMC License"
	USEROPENUM_UPDATE_SERVER_FW_FROM_ISO UserOpEnum = "Update Server FW from ISO"
	USEROPENUM_RESET_SERVER_BIOS_TO_FACTORY_DEFAULT UserOpEnum = "Reset Server BIOS to Factory Default"
	USEROPENUM_CONNECT_L3_UNDERLAY UserOpEnum = "Connect L3 Underlay"
	USEROPENUM_CONFIGURE_SMART_IO UserOpEnum = "Configure Smart IO"
	USEROPENUM_CLEAN_UP_CONTROL_PLANE UserOpEnum = "Clean up Control Plane"
	USEROPENUM_INIT_BMC UserOpEnum = "Init BMC"
	USEROPENUM_ENCRYPT_BOOT_VOLUME UserOpEnum = "Encrypt Boot Volume"
	USEROPENUM_CLEAR_BMC_LOGS UserOpEnum = "Clear BMC Logs"
)
