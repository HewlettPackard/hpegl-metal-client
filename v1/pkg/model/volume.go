// Copyright (c) 2016-2020 Hewlett Packard Enterprise Development LP.

package model

// NewVolume is used as request body during Volume create only
type NewVolume struct {
	Name        string
	Description string // Description is some human readable inforamtion.
	FlavorID    string // Flavor of volume created
	Capacity    uint64 // Requested size of the volume in MB.
	LocationID  string // ID of the associated pod / data center
}

// AddVolume is used as part of NewHost to allow creation of Volumes
// at the same time a Host is created
type AddVolume struct {
	Name        string
	Description string // Description is some human readable inforamtion.
	FlavorID    string // Flavor of volume created
	Capacity    uint64 // Requested size of the volume in MB.
}

// Volume represents an external array-based disk volume.
// Return body for create, get, list, update and as update request body
type Volume struct {
	ResourceBase
	Description string       // Description is some human readable inforamtion.
	LocationID  string       // PodID is the ID of the associated pod
	Capacity    uint64       // Capacity is the requeste size of the volume in MB.
	FlavorID    string       // Flavor of volume created
	State       VolumeState  // State is the state of the volume - updated by the storage controller(?? Db load ??).
	Status      VolumeStatus // Status of the volume updated by the storage controller.
}

// VolumeInfo repesents a minimal description of an already existing Volume for
// use by AvailableResources struct.  This is to allow a user to attach an
// existing volume to a new host.
type VolumeInfo struct {
	ID          string
	Name        string
	Description string
	FlavorID    string
	Capacity    uint64 // Volume size in MB
	LocationID  string
	State       VolumeState  // State is the state of the volume - updated by the storage controller(?? Db load ??).
	Status      VolumeStatus // Status of the volume updated by the storage controller.
}

// VolumeState all posible volume states stored in the  model.
type VolumeState string

// VolumeState enum values
const (
	VolumeStateNew       VolumeState = "new"       // VolumeStateNew a new volume is being requested in the portal
	VolumeStateAllocated VolumeState = "allocated" // VolumeStateAllocated volume is created and ready for an attachment
	VolumeStateVisible   VolumeState = "visible"   // VolumeStateVisible volume is visible from the array to an initiator.
	VolumeStateDeleted   VolumeState = "deleted"   // VolumeStateDeleted is no longer on the array.
	VolumeStateFailed    VolumeState = "failed"    // VolumeStateFailed the volume is not useable and needs to be deleted.
)

// VolumeStatus status of the volume stored in the model.
type VolumeStatus string

// VolumeStatus enum values
const (
	VolumeStatusOK       VolumeStatus = "ok"       // VolumeStatusOK the array detects no issues with the volume.
	VolumeStatusError    VolumeStatus = "error"    // VolumeStatusError the volume is inaccessible.
	VolumeStatusDegraged VolumeStatus = "degraded" // VolumeStatusDegraged the volume is degraded but still accessible.
)
