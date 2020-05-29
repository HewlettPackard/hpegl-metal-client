// Copyright (c) 2020 Hewlett Packard Enterprise Development LP.

package model

import "time"

// ResourceBase is a type that implements Resource
type ResourceBase struct {
	ID       string    // globally unique id of this resource
	ETag     string    // ETag version uuid
	Name     string    // Unique name of this resource, typically used for display name as well
	Created  time.Time // timestamp when resource was created
	Modified time.Time // timestamp when resource was modified
}
