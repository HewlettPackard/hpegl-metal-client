// (C) Copyright 2021-2022 Hewlett Packard Enterprise Development LP

/*
 * HPE GreenLake for bare metal API
 *
 * This Metal Client REST API provides access to bare metal as-a-service (BMaaS) within a single project context. Clients are able to create fully-provisioned hosts, storage volumes, and project-specific private networks in an isolated project environment. Project-owned resources that can be accessed via this API include... Host, Volume, VolumeAttachment, Network (project private), and SSH Key.    Each API call is done within a single project context. The specific Project identifier must be provided within the header of each REST call. The server will validate that the provided authentication credentials (JWTs) are valid for the referenced project before any operation is performed.  If a single credential is valid for multiple projects, the client must still reference a single project in the header for each API call.    Clients can also access information about available services and resources through the AvailableResources object. This object provides detailed  information about the OS imaging options, the machine size options, the storage volume options, and data center locations which are needed when creating hosts and volumes.  
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
import (
	"time"
)
// Host struct for Host
type Host struct {
	// Unique ID for the resource instance as generated by the Metal service
	ID string `json:"ID"`
	// Used to determine whether the DB entry has changed since it was last read. This value is updated each time the resource is updated.  Client must send this value unchanged for any update operation.
	ETag string `json:"ETag"`
	// Common name for the resource instance. Must be 128 or fewer printable characters
	Name string `json:"Name"`
	// Time when the resource was created in the database
	Created time.Time `json:"Created,omitempty"`
	// Time when the resource was last modified in the database
	Modified time.Time `json:"Modified,omitempty"`
	Description string `json:"Description"`
	// The image service identifier used to image the server. ServiceID is one of those listed by the Images array returned as part of the get /available-resources call.
	ServiceID string `json:"ServiceID,omitempty"`
	// Overall flavor of server image used to image the server
	ServiceFlavor string `json:"ServiceFlavor,omitempty"`
	// Version of the ServiceFlavor used to image the server
	ServiceVersion string `json:"ServiceVersion,omitempty"`
	// The location of the machine assigned to the host.  LocationID is one of those listed by the LocationInfo array returned as part of the get /available-resources call.
	LocationID string `json:"LocationID,omitempty"`
	// Name of the machine size used to identify and select the machine assigned to the host.  MachineSizes are described by the MachineSize array returned by the get /available-resources call.
	MachineSizeName string `json:"MachineSizeName,omitempty"`
	// UniqueID referring to the machine size used to identify and select the machine assigned to the host.  MachineSizes are described by the MachineSize array returned by the get /available-resources call.
	MachineSizeID string `json:"MachineSizeID,omitempty"`
	// UniqueID referring to the machine on which this host is running.
	MachineID string `json:"MachineID,omitempty"`
	// IDs of SSH Keys used when configuring the Host
	SSHKeyIDs []string `json:"SSHKeyIDs,omitempty"`
	// Specific SSH keys that were when configuring the host.
	SSHAuthorizedKeys []string `json:"SSHAuthorizedKeys,omitempty"`
	// The list of IDs corresponding to the networks that were provisioned to the host. These networks are among those listed in the Networks array returned by the get /available-resources call.
	NetworkIDs []string `json:"NetworkIDs"`
	// The host's default network ID. This needs to be one of the values in the  \"NetworkIDs\" list.
	NetworkForDefaultRoute string `json:"NetworkForDefaultRoute"`
	// ID of the network selected to be untagged. This needs to be one of the values in the  \"NetworkIDs\" list.
	NetworkUntagged string `json:"NetworkUntagged"`
	// The list of pre-allocated IP addresses corresponding to the list of NetworkIDs. Pre-allocated IP addresses are optional, but required when updating a host containing Pre-allocated IP addresses.
	PreAllocatedIPs []string `json:"PreAllocatedIPs,omitempty"`
	// The map of Service Network (Provider) ID to Provider MAC address.   The Service Network must be a provider network provisioned to this host. Any Service Networks not included here will default to the physical MAC learned during machine discovery.
	ServiceNetsProviderMAC map[string]string `json:"ServiceNetsProviderMAC,omitempty"`
	// User-provided data attached to the image configuration data when the host was provisioned
	UserData string `json:"UserData,omitempty"`
	// User-provided data to represent the identity of the host within an application environment. For example, this could be set to represent the Kubernetes node ID if the host is provisioned as a Kubernetes node.
	NodeID string `json:"NodeID,omitempty"`
	ISCSIConfig *HostIscsiConfig `json:"ISCSIConfig,omitempty"`
	// Details describing host network connections
	Connections []HostConnection `json:"Connections,omitempty"`
	// True if the Host has been deleted.
	Deleted bool `json:"Deleted,omitempty"`
	// Describes if the portal is in active communication to the device
	PortalCommOkay bool `json:"PortalCommOkay,omitempty"`
	PowerStatus HostPowerState `json:"PowerStatus,omitempty"`
	State HostState `json:"State,omitempty"`
	Substate HostSubstate `json:"Substate,omitempty"`
	StateTime time.Time `json:"StateTime,omitempty"`
	SubstateTime time.Time `json:"SubstateTime,omitempty"`
	Progress int64 `json:"Progress,omitempty"`
	Alert bool `json:"Alert,omitempty"`
	AlertInfo []HostAlertInfo `json:"AlertInfo,omitempty"`
	// The current workflow the host is in
	Workflow string `json:"Workflow,omitempty"`
	SummaryStatus HealthStatus `json:"SummaryStatus,omitempty"`
	// The map of label name to label value for the host.
	Labels map[string]string `json:"Labels,omitempty"`
}
