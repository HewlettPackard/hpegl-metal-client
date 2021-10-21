// Copyright 2016-2021 Hewlett Packard Enterprise Development LP

package model

// VolumeAttachment represents an attachment between a host and an external volume.
type VolumeAttachment struct {
	ResourceBase
	VolumeID string
	HostID   string
	// IPV4 address of the iSCSI initiator (host) in dot notation or FQN.
	HostIPAddress string
	// IQN of the iSCSI initiator used for identification during iSCSI login.
	IQN string
	// Challenge Authentication Protocol secret to be shared between array and initiator.
	// If empty, no CHAP login is enabled; if set, it must be a string between 12 and 16 characters.
	CHAPSecret string
	// CHAP username to use for CHAP authentication
	CHAPUserName string
	// Logical Unit Number to be assigned to the volume on export.
	LUN int32
	// IQN of the iSCSI target managing the volume.
	VolumeTargetIQN string
	// IPV4 address of the iSCSI target managing the volume.
	VolumeTargetIPAddress string
	// State of the volumeattachment, managed by the portal volmon bot.
	State VaStateEnum
	// VolumeGroupID           string // VolumeGroupID may be empty of this volume is not part of a group.
	// ArrayNetworkInterfaceID string // ArrayNetworkInterfaceID this must be initialized at creation time; it must correspond to an arry netoworkID that is reported in // StorageNetwork.ArrayNetworkInterfaceID
	// ArrayVolumeAttachmentID string // Array parameters that may be passed to the driver with casts. These fields are populated by the storagemon bot and should be // unintialized at point of creation.
	// ArrayHostID             string // ArrayHostID is a reference to the array-side host entry.
}

// VaStateEnum defines the posible VolumeAttachment states.
type VaStateEnum string

// VaStateEnum values.
const (
	VaStateNull        VaStateEnum = ""            // Unknown state coming from previous schema versions.
	VaStateNew         VaStateEnum = "new"         // VA is created in the portal/database.
	VaStateExporting   VaStateEnum = "exporting"   // Volume is being exported on target.
	VaStateAttaching   VaStateEnum = "attaching"   // Volume is being attached on initiator: only for v2 lifecycle.
	VaStateReady       VaStateEnum = "ready"       // Volume is ready for traffic between initiator and target.
	VaStateDetaching   VaStateEnum = "detaching"   // Volume is being detached on initiator: only for v2 lifecycle.
	VaStateUnexporting VaStateEnum = "unexporting" // Volume is being unexported on target.
	VaStateDeleted     VaStateEnum = "deleted"     // VA has been deleted (might not be needed if object is removed from DB).
	VaStateFailed      VaStateEnum = "failed"      // VA has run into a problem.
)

// ISCSIParameters contains the parameters required to attach a volume using iSCSI.
type ISCSIParameters struct {
	// IPV4 address of the iSCSI initiator (host) in dot notation or FQN.
	HostIPAddress string
	// IQN of the iSCSI initiator used for identification during iSCSI login.
	InitiatorName string
	// Challenge Authentication Protocol secret to be shared between array and initiator.
	// If empty, no CHAP login is enabled; if set, it must be a string between 12 and 16 characters.
	CHAPSecret string
	// CHAP username to use for CHAP authentication.
	CHAPUserName string
}

// ProtocolEnum defines the supported attachment protocols.
type ProtocolEnum string

// Protocol Enum Value
const (
	ProtocolKindUnknown = "unknown"
	ProtocolKindISCSI   = "iscsi"
)

// ProtocolParameters contains the information to attach a volume using the
// specified protocol.
type ProtocolParameters struct {
	// Protocol specifies the protocol to be used to attach the volume. See the
	// protocol specific parameters below for legal values.
	Protocol ProtocolEnum
	// ISCSI contains the information required to attach the volume using ISCSI.
	// It must be specified if Protocol is ProtocolKindISCSI.
	ISCSI ISCSIParameters
}

// NewVolumeAttachment is used to create an attachment between a non-Quake
// controlled machine (i.e. a machine that does not have a HostID) and a volume.
type NewVolumeAttachment struct {
	// ResourceBase contains common elements of Quake Objects. The caller must set
	// Name and should set Description.
	ResourceBase
	// VolumeID identifies the volume attach.
	VolumeID string
	// Protocol specifies the protocol specific information for this attachment.
	Protocol ProtocolParameters
}

// VolumeAttachmentHostUuid represents the host that a volume is being attached to or
// detached from.
type VolumeAttachmentHostUuid struct {
	HostID string
}
