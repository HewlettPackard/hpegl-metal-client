// Copyright (c) 2016-2020 Hewlett Packard Enterprise Development LP.

package model

// VolumeAttachment represents an attachment between a host and an external volume.
type VolumeAttachment struct {
	ResourceBase
	VolumeID              string
	HostID                string //
	HostIPAddress         string // HostIPAddress is the IPV4 IPAddress of the iSCSI initiator (host) in dot notation or FQN.
	IQN                   string // IQN is the full initiator name used for identification during iSCSI login.
	CHAPSecret            string // CHAPSecret is the Challenge Authentication Protocol secret to be shared between array and initiator. If empty, no CHAP // login is enabled; if set it must be a string between 12 and 16 characters.
	CHAPUserName          string // CHAPUserName is the CHAP username to use for CHAP authentication
	LUN                   int32  // LUN is the Logical Unit Number to be assigned to the volume on export.
	VolumeTargetIQN       string // VolumeTargetIQN is the iQN for the volume, assigned by the array correspnding to the volume.
	VolumeTargetIPAddress string // VolumeTargetIPAddress is the IPV4 address of the iSCSI volume export.
	// VolumeGroupID           string // VolumeGroupID may be empty of this volume is not part of a group.
	// ArrayNetworkInterfaceID string // ArrayNetworkInterfaceID this must be initialized at creation time; it must correspond to an arry netoworkID that is reported in // StorageNetwork.ArrayNetworkInterfaceID
	// ArrayVolumeAttachmentID string // Array parameters that may be passed to the driver with casts. These fields are populated by the storagemon bot and should be // unintialized at point of creation.
	// ArrayHostID             string // ArrayHostID is a reference to the array-side host entry.
}

// VolumeAttachmentHostUuid represents the host that a volume is being attached to or
// detached from.
type VolumeAttachmentHostUuid struct {
	HostID string
}
