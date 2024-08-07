// (C) Copyright 2021-2024 Hewlett Packard Enterprise Development LP

/*
 * HPE GreenLake for bare metal API
 *
 * This Metal Client REST API provides access to bare metal as-a-service (BMaaS) within a single project context. Clients are able to create fully-provisioned hosts, storage volumes, and project-specific private networks in an isolated project environment.   Project-owned resources that can be accessed via this API include - Host,  Volume, VolumeAttachment, Network (project private), and SSH Key. Each API  call is done within a single project context. The specific Project identifier  must be provided within the header of for each REST call. The server will  validate that the provided authentication credentials (JWTs) are valid  for  the referenced project before any operation is performed. If a single credential is valid for multiple projects, the client must still reference a single project  in the header for each API call.    Clients can also access information about available services and resources through the AvailableResources object. This object provides detailed  information about the OS imaging options, the machine size options, the storage volume options, and data center locations which are needed when creating hosts and volumes.  
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// PartitionFormula It provide partition size information.
type PartitionFormula struct {
	// Represents a partition size in MiB.  As a special case a size of 0 can be specified that will equate to the remaining space on the disk.
	Equation string `json:"Equation"`
	// It causes the value computed by formula execution to be rounded up to MinSize; if MinSize is non-zero. Multipliers like M and G maybe used, e.g. 1M == (1024*1024)bytes. An empty string means zero bytes.  If no multiplier is specified M is assumed.
	MinSize string `json:"MinSize"`
	// MaxSize causes the value computed by formula execution to be rounded up to MaxSize; if MaxSize is non-zero. Multipliers like M, G and T maybe used, e.g. 1M == (1024*1024)bytes. An empty string means zero bytes.  If no multiplier is specified M is assumed. Equations may also be used e.g. 25%disk means MaxSize will cause any partition  to never exceed 25% of device capacity.
	MaxSize string `json:"MaxSize"`
}
