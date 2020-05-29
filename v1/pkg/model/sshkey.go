// Copyright (c) 2016-2020 Hewlett Packard Enterprise Development LP.

package model

// NewSSHKey is used when creating a new SSHKey via project API
type NewSSHKey struct {
	Name string
	Key  string
}

// SSHKey value returned by list, get, create, update and as sent
// during an update
type SSHKey struct {
	ResourceBase
	Key string
}

// SSHKeyEntry describes an existing SSHKey so it can be added/referenced during
// a create Host operation.  Part of the AvailableResources struct.
type SSHKeyEntry struct {
	ID   string
	Name string
	Key  string
}
