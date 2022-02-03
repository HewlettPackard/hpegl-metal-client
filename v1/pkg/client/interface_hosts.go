// (C) Copyright 2021-2022 Hewlett Packard Enterprise Development LP

//

package client

import (
	_context "context"
	_nethttp "net/http"
)

// HostsAPI defines the client functions provided for Hosts.
type HostsAPI interface {
	/*
	   Add Create a new Host
	   Creates a new host object which kicks off the provisioning of a physical server in accordance to the attributes provided for the Host object.  Most values for these options must be selected from the set of options provided by the get available-resources API call. The SvcFlavor, SvcVersion, LocationID, SSHKeyIDs, and Network attribute must all be set with appropriate ID values from the available-resources call.
	    * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    * @param newHost Defines the configuration of the desired host. See the schema for descriptions of individual attributes.
	   @return Host
	*/
	Add(ctx _context.Context, newHost NewHost) (Host, *_nethttp.Response, error)
	/*
	   Delete Delete a Host
	   Deletes the Host with the matching ID.  A host in the &#39;Ready&#39; state must first be powered-off before a delete will be permitted.  Deletes to hosts in other states is permitted regardless of the power state
	    * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    * @param hostId ID of Host to delete
	*/
	Delete(ctx _context.Context, hostId string) (*_nethttp.Response, error)
	/*
	   GetByID Retrieve Host by ID
	   Returns a single Host with matching ID
	    * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    * @param hostId ID of Host to return
	   @return Host
	*/
	GetByID(ctx _context.Context, hostId string) (Host, *_nethttp.Response, error)
	/*
	   List List all Hosts in project
	   Returns an array of all Host objects defined within the project.
	    * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    * @param optional nil or *HostsApiListOpts - Optional Parameters:
	    * @param "All" (optional.String) -  Includes deleted Host objects in the response when set to \"true\".
	   @return []Host
	*/
	List(ctx _context.Context, localVarOptionals *HostsApiListOpts) ([]Host, *_nethttp.Response, error)
	/*
	   PowerOff Power off Host by ID
	   Powers off a single Host with matching ID
	    * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    * @param hostId ID of Host to power off
	   @return Host
	*/
	PowerOff(ctx _context.Context, hostId string) (Host, *_nethttp.Response, error)
	/*
	   PowerOn Power on Host by ID
	   Powers on a single Host with matching ID
	    * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    * @param hostId ID of Host to power on
	   @return Host
	*/
	PowerOn(ctx _context.Context, hostId string) (Host, *_nethttp.Response, error)
	/*
	   Update Update an existing Host
	   Updates the Host with the matching ID.  Update is permitted only if the host is in the &#39;Ready&#39; or &#39;Connection Updating Failed&#39; state.  Only the Host &#39;Description&#39;, &#39;Networks&#39;, and &#39;NetworkForDefaultRoute&#39; can be updated.
	    * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    * @param hostId ID of host to update
	    * @param host Updated Host
	   @return Host
	*/
	Update(ctx _context.Context, hostId string, host Host) (Host, *_nethttp.Response, error)
}
