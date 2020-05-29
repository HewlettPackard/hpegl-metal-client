// Copyright (c) 2016-2020 Hewlett Packard Enterprise Development LP.

package model

// NewNetwork is used with the create network REST call to
// create a new project private network
type NewNetwork struct {
	Name        string
	LocationID  string
	Description string
}

// Network is used to describe the return value for get, list, create, or update
// network calls and as body to update network calls
type Network struct {
	ResourceBase `json:",inline" `
	LocationID   string
	Description  string
	Kind         NetworkKindEnum
	HostUse      HostUseEnum
}

// AvailableNetwork lists information about available networks as part of
// the AvailableResources struct
type AvailableNetwork struct {
	ID          string
	Name        string
	LocationID  string
	Description string
	Kind        NetworkKindEnum
	HostUse     HostUseEnum
}

// NetworkKindEnum describes the basic types of network
type NetworkKindEnum string

// NetworkKindEnum values
const (
	NetworkKindShared  NetworkKindEnum = "Shared"
	NetworkKindPrivate NetworkKindEnum = "Private"
	NetworkKindCustom  NetworkKindEnum = "Custom"
)

// HostUseEnum identifies whether a host is forced to use
// a specific network or whether it is optional.  Default is optional
// but will typically show up by default in GUI view.
type HostUseEnum string

// HostUseEnum values
const (
	HostUseRequired    HostUseEnum = "Required"
	HostUseRecommended HostUseEnum = "Default"
	HostUseOptional    HostUseEnum = "Optional"
)
