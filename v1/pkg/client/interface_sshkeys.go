// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package client

import (
	_context "context"
	_nethttp "net/http"
)

// SshkeysAPI defines the functions provided by the client for Sshkeys.
type SshkeysAPI interface {
	/*
	   Add Add a new SSH Key
	   Adds a new SSH Key that can be referenced when creating a Host
	    * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    * @param newSshKey SSH Key that is to be added to the project
	   @return SshKey
	*/
	Add(ctx _context.Context, newSshKey NewSshKey) (SshKey, *_nethttp.Response, error)
	/*
	   Delete Delete an SSH key
	   Deletes the SSH key with the matching ID
	    * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    * @param sshkeyId ID of sshkey to delete
	*/
	Delete(ctx _context.Context, sshkeyId string) (*_nethttp.Response, error)
	/*
	   GetByID Retrieve SSH Key by ID
	   Returns a single SSH key with matching ID
	    * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    * @param sshkeyId ID of sshkey to return
	   @return SshKey
	*/
	GetByID(ctx _context.Context, sshkeyId string) (SshKey, *_nethttp.Response, error)
	/*
	   List List all sshkeys in project
	   Returns an array of all SSHKey objects defined within the project. This does not include any SSH keys that are only defined as part of creating host.
	    * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	   @return []SshKey
	*/
	List(ctx _context.Context) ([]SshKey, *_nethttp.Response, error)
	/*
	   Update Update an existing SSH Key.  Only the name or key fields can be changed.
	    * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    * @param sshkeyId ID of sshkey to update
	    * @param sshKey Updated SSH key
	   @return SshKey
	*/
	Update(ctx _context.Context, sshkeyId string, sshKey SshKey) (SshKey, *_nethttp.Response, error)
}
