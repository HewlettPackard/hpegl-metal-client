// Copyright (c) 2016-2021 Hewlett Packard Enterprise Development LP.

package model

import (
	"github.com/hpe-hcss/quake-client/v1/pkg/model/enums/country"
)

// AvailableResources Lists the information needed for client to create host
type AvailableResources struct {
	VolumeFlavors    []VolumeFlavor
	Volumes          []VolumeInfo
	Images           []AvailableImage
	SSHKeys          []SSHKeyEntry
	MachineInventory []MachineInventory
	StorageInventory []StorageInventory
	MachineSizes     []MachineSize
	Locations        []LocationInfo
	Networks         []AvailableNetwork
}

// StorageInventory provides information about where volumes of
// specific flavor are available and the maximum volume that
// could be created for that flavor.
//
// NOTE: Multiple inventory entries may exist for the same FlavorID
// and LocationID to represent multiple capacity pools. Or the
// server could choose to only show the entry with the largest Capacity.
type StorageInventory struct {
	FlavorID   string // VolumeFlavor ID
	Capacity   uint64 // Capacity in MB - Capped by project limits
	LocationID string // Pod ID
}

// VolumeFlavor repesents a minimal description of a VolumeFlavor as needed
// by AvailableResources struct
type VolumeFlavor struct {
	ID      string
	Name    string
	Details FlavorDesc
}

// LocationInfo provides a mapping from LocationID to a more descriptive
// attributes about the pod or data center location
type LocationInfo struct {
	ID         string
	Country    country.Enum // must match country field
	Region     string
	DataCenter string
}

// AvailableImage lists the available imaging services
type AvailableImage struct {
	ID          string
	Category    string
	Flavor      string
	Version     string
	Description string
}

// MachineSize lists and describes an available MachineSize
type MachineSize struct {
	ID      string
	Name    string
	Details FlavorDesc
}

// FlavorDesc provides the detailed description of a machine size
// without any pricing information.
type FlavorDesc struct {
	Collection string
	Banner1    string
	Banner2    string
	Bullets    []string
	InfoLink   string
	Tooltip    string
}

// MachineInventory lists the number of available Machines of a certain size
// (MachineSize ID) at a specific location (Pod ID)
type MachineInventory struct {
	LocationID string
	SizeID     string
	Number     int32 // capped by project limits
}
