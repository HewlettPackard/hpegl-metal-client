// Copyright (c) 2016-2020 Hewlett Packard Enterprise Development LP.

package model

// NewProject is to create a new Project within the Quake service
type NewProject struct {
	Name    string // Name of the project
	Profile Profile
	Limits  ProjectLimits
}

// Project defines an isolated space for creating Hosts, Volumes,
// and private Networks.  A project is often aligned to a specific
// team within an organization or a cluster
type Project struct {
	ResourceBase
	Profile Profile
	Limits  ProjectLimits
}

// ProjectLimits place general constraints on a project in terms of
// server and storage.
type ProjectLimits struct {
	Hosts          uint32
	Volumes        uint32
	VolumeCapacity uint64
}

// Profile is a generalized description of a team
type Profile struct {
	TeamName      string
	TeamDesc      string
	Company       string
	Address       string
	Email         string
	EmailVerified bool
	PhoneNumber   string
	PhoneVerified bool
}
