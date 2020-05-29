// Copyright (c) 2016-2020 Hewlett Packard Enterprise Development LP.

package model

import (
	"time"

	"github.com/quattronetworks/quake-client/v1/pkg/model/enums/netmask"
)

// NewHost defines the "post /hosts" body when creating a Host.  Field IDs
// noted below are determined using the arrays returned during the "get
// /available-resources" call. Where applicable, the AvailableResources
// attribute that provides the IDs is shown in []
type NewHost struct {
	Name          string      // Host name in operating system
	Description   string      // General description of hot
	ServiceID     string      // ID of the imaging service to use   [Images]
	LocationID    string      // ID of the data center (Pod)        [Locations]
	MachineSizeID string      // ID of the desired machine size     [MachineSizes]
	SSHKeyIDs     []string    // IDs of SSHKeys to be added         [SSHKeys]
	NetworkIDs    []string    // IDs of networks connected to host  [Networks]
	VolumeIDs     []string    // IDs of volumes to attach           [Volumes]
	NewVolumes    []AddVolume // Information on any new volumes to be created
	UserData      string      // UserData is copied directly to CloudInit
	NodeID        string      // Reference ID for application software (optional)
}

// Host defines get, list, create, and update response body as well
// as the update request body
//
// NOTE: Attached volumes are not shown here as they are only reference in
// the VolumeAttachment object.
type Host struct {
	ResourceBase

	Description       string          // General description of hot
	ServiceID         string          // ID of the imaging service used   [Images]
	ServiceFlavor     string          // Flavor of server image
	ServiceVersion    string          // Version of server image
	LocationID        string          // ID of the data center (Pod)        [Locations]
	MachineSizeName   string          // Name of the machine size     [MachineSizes]
	MachineSizeID     string          // ID of the machine size     [MachineSizes]
	SSHKeyIDs         []string        // IDs of SSHKeys added when imaged     [SSHKeys]
	SSHAuthorizedKeys []string        // Direct input of Add'l SSH keys
	NetworkIDs        []string        // IDs of networks connected to host  [Networks]
	UserData          string          // UserData is copied directly to CloudInit
	NodeID            string          // Optional reference ID for application software
	ISCSIConfig       ISCSIConfig     // TODO: Is this needed?  iSCSI related information; shared with any external volumes
	Connections       []Connection    // Details describing host network connections
	Deleted           bool            // TODO: Should we be able to return info on deleted hosts?
	PortalCommOkay    bool            // Describes if the portal is in active communication to device
	PowerStatus       PowerStatusEnum // Current machine power status

	// Attributes for host state machine work:
	State        HostStateEnum    // updated to correspond to new state machine major states
	Substate     HostSubstateEnum // substates or steps within the major host states
	StateTime    time.Time        // timestamp of when the current State was entered
	SubstateTime time.Time        // timestamp of when the current Substate was entered
	Progress     int64            // Progress on the overall host provisioning
	Alert        bool             // flags when an issue has been detected
	AlertInfo    []HostAlertInfo  // details about a specific alert condition
	// Events            []*HostEvent
	// LogFileID    string           // provides the file ID for the console log snapped after deploy
	// RetryCount   uint32           // is the number of times a substate retry has been attempted
}

// PowerStatusEnum is used to report the current host PowerStatus
type PowerStatusEnum string

// PowerStatusEnum values
const (
	PowerStatusOn      PowerStatusEnum = "ON"
	PowerStatusOff     PowerStatusEnum = "OFF"
	PowerStatusUnknown PowerStatusEnum = "UNKNOWN"
)

// ISCSIConfig encapsulates iSCSI configuration informatino
type ISCSIConfig struct {
	InitiatorName         string // Fully qualified iSCSI intiator name of this host. This is generated when a host is created and the information is pushed to the new host via cloudInit.
	CHAPSecret            string // CHAPSecret is also generated and pushed
	CHAPUser              string // CHAPUser is also generated and pushed
	ISCSIDiscoveryAddress string // populated by volume attachment operations and can be populated and pushed if the attachament is known a priori.
}

// HostAlertInfo gives information related to host state machine errors
type HostAlertInfo struct {
	Alert    HostAlertEnum
	State    HostStateEnum
	Substate HostSubstateEnum
	Message  string
	Time     time.Time
	Ack      bool
}

// Connection defines a Host network connection with some details being
// embedded in the NetworkConnection struct.
type Connection struct {
	Name     string
	Ports    []ServerPort
	HA       bool
	Speed    PortSpeedEnum
	Networks []NetworkConnection
}

// PortSpeedEnum indicates the speed of the Ethernet port
type PortSpeedEnum string

// Enum values
const (
	Speed100Mb PortSpeedEnum = "100Mb"
	Speed1Gb   PortSpeedEnum = "1Gb"
	Speed10Gb  PortSpeedEnum = "10Gb"
	Speed40Gb  PortSpeedEnum = "40Gb"
)

// ServerPort represents a server port.  Multiple server ports may
// be used in a single MachineNetworkConnection.
type ServerPort struct {
	Name   string
	HWAddr string
}

// NetworkConnection provides the all of the Host specifics to allow a network
// connection to be provisioned on a Host.
type NetworkConnection struct {
	Name      string
	NetworkID string
	IP        string
	Subnet    string
	Netmask   netmask.Enum
	Gateway   string
	DNS       []string
	VLAN      uint32
	VNI       uint32
	Proxy     string
	NoProxy   string
}

// HostStateEnum defines the possible states that an Host can be in
type HostStateEnum string

// HostStateEnum values
const (
	HostStateNew        HostStateEnum = "New"      // Host has been created
	HostStateDeleting   HostStateEnum = "Deleting" // Host is being deleted
	HostStateDeleted    HostStateEnum = "Deleted"  // Host has been deleted
	HostStateFailed     HostStateEnum = "Failed"   // Host configuration failed
	HostStateAllocating HostStateEnum = "Allocating"
	HostStateImaging    HostStateEnum = "Imaging"
	HostStateConnecting HostStateEnum = "Connecting"
	HostStateReady      HostStateEnum = "Ready"
)

// HostSubstateEnum identifies valid host substates
type HostSubstateEnum string

// HostSubstateEnum values
const (
	HostSubstateNil             HostSubstateEnum = ""                 // Used by multiple states
	HostSubstateInit            HostSubstateEnum = "Init"             //
	HostSubstateComplete        HostSubstateEnum = "Complete"         //
	HostSubstateFailed          HostSubstateEnum = "Failed"           //
	HostSubstateWorking         HostSubstateEnum = "Working"          // Allocating Substates
	HostSubstateNeedMachine     HostSubstateEnum = "Need-Machine"     //
	HostSubstateNeedIPs         HostSubstateEnum = "Need-IPs"         //
	HostSubstateInitOff         HostSubstateEnum = "Init-Off"         // Imaging Substates
	HostSubstateSettingPxe      HostSubstateEnum = "Setting-PXE"      //
	HostSubstateClearingLog     HostSubstateEnum = "Clearing-Log"     //
	HostSubstateStartingDeploy  HostSubstateEnum = "Starting-Deploy"  //
	HostSubstateBlocked         HostSubstateEnum = "Blocked"          //
	HostSubstateDeployPoweron   HostSubstateEnum = "Deploy-Poweron"   //
	HostSubstateDeploying       HostSubstateEnum = "Deploying"        //
	HostSubstateConfirmOff      HostSubstateEnum = "Confirm-Off"      //
	HostSubstateSnapLog         HostSubstateEnum = "SnapLog"          //
	HostSubstateFailSnapLog     HostSubstateEnum = "Fail-SnapLog"     //
	HostSubstateSettingDiskBoot HostSubstateEnum = "Setting-DiskBoot" //
	HostSubstateConnect         HostSubstateEnum = "Connect"          // Connecting Substates
	HostSubstateConnectPoweron  HostSubstateEnum = "Connect-Poweron"  //
	HostSubstateDisconnect      HostSubstateEnum = "Disconnect"       //Deleting Substates
	HostSubstateAbortDeploy     HostSubstateEnum = "Abort-Deploy"     //
	HostSubstateDeletePoweroff  HostSubstateEnum = "Delete-Poweroff"  //
	HostSubstateReleasing       HostSubstateEnum = "Releasing"        //
)

// HostAlertEnum string
type HostAlertEnum string

// HostAlertEnum values
const (
	HostAlertSubstateTimedOut HostAlertEnum = "substate-timed-out" // Time-out
	HostAlertOpFailed         HostAlertEnum = "op-failed"          // Op failed
	HostAlertUnknownState     HostAlertEnum = "unknown-state"      // Encountered an unknown host state/substate
)

/*
// HostEvent holds information relevant to billing for host time
type HostEvent struct {
	Time     Time
	Type     EventType
	UserID   string
	UserName string
}

// EventType is an enum for HostEvents
type EventType string

// HostEvent event types
const (
	EventTypeCreate EventType = "Create"
	EventTypeAlloc  EventType = "Allocate" // machine is allocated
	EventTypeDeploy EventType = "Deploy"   //
	EventTypeReady  EventType = "Ready"    // moved to active state
	EventTypeFree   EventType = "Free"
	EventTypeDelete EventType = "Delete"
	EventTypeFailed EventType = "Failed"
)
*/
