// Copyright (c) 2016-2020 Hewlett Packard Enterprise Development LP.

package model

import "time"

// BillingHourFraction represents the fraction of an hour over which the usage
// will be rounded up to the next hour.  A little bit of grace time is provided,
// but not much.
const BillingHourFraction = 0.05

// UsageReportRequest is used in the requestBody to get a usage report
type UsageReportRequest struct {
	StartTime time.Time
	EndTime   time.Time
	Include   []ReportItemEnum
}

// ReportItemEnum is used to indicate what parts of the report are to be
// returned
type ReportItemEnum string

// ReportItemEnum values
const (
	HostReporting   ReportItemEnum = "hosts"
	VolumeReporting ReportItemEnum = "volumes"
)

// UsageReport is the summary response to a UsageReportRequest
type UsageReport struct {
	Hosts   []HostUsageEntry
	Volumes []VolumeUsageEntry
}

// HostUsageEntry defines usage of a single machine during the reporting
// window
type HostUsageEntry struct {
	ProjectID  string    // Project ID that contained the host
	LocationID string    // Unique ID of the data center location
	Allocated  time.Time // Timestamp of when a machine was allocated for use by host
	Freed      time.Time // Timestamp of when the host was deleted and machine freed
	Ready      time.Time // Timestamp of when the the host transitioned to ready state
	UsageStart time.Time // The start of usage reporting window
	UsageEnd   time.Time // The end of the the usage reporting window
	UsageHours uint32    // The total number of hours used (rounded up)
	Error      string    // Error reporting if usage could not be calculated

	HostName        string // Host name
	HostID          string // Host unique ID
	MachineSizeName string // Name of the MachineSize requested when host was created
	MachineSizeID   string // Unique ID of the MachineSize requested when host was created
}

// VolumeUsageEntry provides usage of a single volume during the reporting
// window.  Actual usage cost must take into account the volume Capacity
// and the UsageHours
type VolumeUsageEntry struct {
	ProjectID  string    // Project ID that contained the host
	LocationID string    // Unique ID of the data center location
	Allocated  time.Time // Timestamp of when a machine was allocated for use by host
	Ready      time.Time // Timestamp of when the storage was available for use
	Freed      time.Time // Timestamp of when the storage was freed
	UsageStart time.Time // The start of usage reporting window
	UsageEnd   time.Time // The end of the the usage reporting window
	UsageHours uint32    // The total number of hours used (rounded up)
	Error      string    // Error reporting if usage could not be calculated

	VolumeName string // Host name
	VolumeID   string // Host unique ID
	FlavorName string // Name of the MachineSize requested when host was created
	FlavorID   string // Unique ID of the MachineSize requested when host was created
	Capacity   uint64 // Volume size in MB
}
