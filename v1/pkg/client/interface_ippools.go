// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

//

package client

import (
	_context "context"
	_nethttp "net/http"
)

// IPPoolsAPI defines the client functions provided for Ippools.
type IPPoolsAPI interface {
	/*
	   AllocateIPs Allocate IPs from the pool
	   Allocate IPs from the pool
	    * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    * @param ippoolId ID of IP pool to allocate IPs
	    * @param iPAllocation IPs being requested starting from an optional base IP and their usage
	   @return IpPool
	*/
	AllocateIPs(ctx _context.Context, ippoolId string, iPAllocation []IpAllocation) (IpPool, *_nethttp.Response, error)
	/*
	   GetByID Retrieve IP pool by ID
	   Returns a single ip pool with matching ID
	    * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    * @param ippoolId ID of IP pool to return
	   @return IpPool
	*/
	GetByID(ctx _context.Context, ippoolId string) (IpPool, *_nethttp.Response, error)
	/*
	   List List all ip pools in project
	   Returns an array of all ip pool objects defined within the project.
	    * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	   @return []IpPool
	*/
	List(ctx _context.Context) ([]IpPool, *_nethttp.Response, error)
	/*
	   ReturnIPs Return IPs to the pool
	   Return IPs to the pool
	    * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    * @param ippoolId ID of IP pool to return IPs
	    * @param requestBody IP returned to the pool
	   @return IpPool
	*/
	ReturnIPs(ctx _context.Context, ippoolId string, requestBody []string) (IpPool, *_nethttp.Response, error)
	/*
	   Update Update IP pool by ID
	   Update a single ip pool with matching ID. &#39;DefaultRoute&#39; can only be updated if ip pool is not currently in-use.
	    * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    * @param ippoolId ID of IP pool to update
	    * @param ipPool Update IPPool
	   @return IpPool
	*/
	Update(ctx _context.Context, ippoolId string, ipPool IpPool) (IpPool, *_nethttp.Response, error)
}
