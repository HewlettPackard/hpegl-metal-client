// Copyright (c) 2016-2020 Hewlett Packard Enterprise Development LP.

package model

// NewProject is to create a new Project within the Quake service.
type NewProject struct {
	Name    string // Name of the project
	Profile Profile
	Limits  ProjectLimits
}

// Project defines an isolated space for creating Hosts, Volumes,
// and private Networks.  A project is often aligned to a specific
// team within an organization or a cluster.
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

// Profile is a generalized description of a team.
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

// ProjectsInfo for information about projects.
type ProjectsInfo struct {
	Projects      []ProjectInfo      `json:"projects"`
	MachineSizes  []MachineSizeInfo  `json:"machine_sizes"`
	VolumeFlavors []VolumeFlavorInfo `json:"volume_flavors"`
}

// ProjectInfo for information about a project.
type ProjectInfo struct {
	ID             string            `json:"id"`
	Name           string            `json:"name"`
	Description    string            `json:"description"`
	NumHosts       int               `json:"num_hosts"`
	NumVolumes     int               `json:"num_volumes"`
	TotalStorageGB int               `json:"total_storage"` //GiB
	Status         ProjectStatusEnum `json:"status"`
}

// ProjectStatusEnum indicates the status of a project.
// TODO Needs to be refined once the project status is finalized.
type ProjectStatusEnum string

const (
	// StatusEnabled indicates that the function or resource is enabled.
	StatusEnabled ProjectStatusEnum = "Enabled"
	// StatusDisabled indicates that the function or resource is enabled.
	StatusDisabled ProjectStatusEnum = "Disabled"
)

// MachineSizeInfo for information about machine sizes.
type MachineSizeInfo struct {
	Name         string `json:"name"`
	ProjectsUsed int    `json:"projects_used"`
	OtherUsed    int    `json:"other_used"`
	Available    int    `json:"available"`
}

// VolumeFlavorInfo for information about volume flavor.
type VolumeFlavorInfo struct {
	Name         string `json:"name"`
	ProjectsUsed int    `json:"projects_used"` //GiB
	OtherUsed    int    `json:"other_used"`    //GiB
	Available    int    `json:"available"`     //GiB
}
