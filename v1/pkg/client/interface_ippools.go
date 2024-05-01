// (C) Copyright 2021-2023 Hewlett Packard Enterprise Development LP

// Code generated by ifacemaker; DO NOT EDIT.

package client

import (
	_context "context"
	_nethttp "net/http"
)

// IPPoolsAPI defines the client functions provided for Ippools.
type IPPoolsAPI interface {
	/*
	   AllocateIPs Allocate IPs from the pool
	   Allocate IPs from the pool If GreenLake Platform IAM issued token is used for authentication, then it is required to pass  &#39;X-Role&#39; and &#39;X-Workspaceid&#39; headers.
	     - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	     - @param ippoolId ID of IP pool to allocate IPs
	     - @param iPAllocation IPs being requested starting from an optional base IP and their usage
	     - @param optional nil or *IppoolsApiAllocateIPsOpts - Optional Parameters:
	     - @param "XRole" (optional.String) -  GreenLake Platform role
	     - @param "XWorkspaceid" (optional.String) -  GreenLake Platform workspace ID

	   @return IpPool
	*/
	AllocateIPs(ctx _context.Context, ippoolId string, iPAllocation []IpAllocation, localVarOptionals *IppoolsApiAllocateIPsOpts) (IpPool, *_nethttp.Response, error)
	/*
	   GetByID Retrieve IP pool by ID
	   Returns a single ip pool with matching imaged. If GreenLake Platform IAM issued token is used for authentication, then it is required to pass  &#39;X-Role&#39; and &#39;X-Workspaceid&#39; headers.
	     - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	     - @param ippoolId ID of IP pool to return
	     - @param optional nil or *IppoolsApiGetByIDOpts - Optional Parameters:
	     - @param "XRole" (optional.String) -  GreenLake Platform role
	     - @param "XWorkspaceid" (optional.String) -  GreenLake Platform workspace ID

	   @return IpPool
	*/
	GetByID(ctx _context.Context, ippoolId string, localVarOptionals *IppoolsApiGetByIDOpts) (IpPool, *_nethttp.Response, error)
	/*
	   List List all ip pools in project
	   Returns an array of all ip pool objects defined within the project. If GreenLake Platform IAM issued token is used for authentication, then it is required to pass  &#39;X-Role&#39; and &#39;X-Workspaceid&#39; headers.
	     - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	     - @param optional nil or *IppoolsApiListOpts - Optional Parameters:
	     - @param "XRole" (optional.String) -  GreenLake Platform role
	     - @param "XWorkspaceid" (optional.String) -  GreenLake Platform workspace ID

	   @return []IpPool
	*/
	List(ctx _context.Context, localVarOptionals *IppoolsApiListOpts) ([]IpPool, *_nethttp.Response, error)
	/*
	   ReturnIPs Return IPs to the pool
	   Return IPs to the pool. If GreenLake Platform IAM issued token is used for authentication, then it is required to pass  &#39;X-Role&#39; and &#39;X-Workspaceid&#39; headers.
	     - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	     - @param ippoolId ID of IP pool to return IPs
	     - @param requestBody IP returned to the pool
	     - @param optional nil or *IppoolsApiReturnIPsOpts - Optional Parameters:
	     - @param "XRole" (optional.String) -  GreenLake Platform role
	     - @param "XWorkspaceid" (optional.String) -  GreenLake Platform workspace ID

	   @return IpPool
	*/
	ReturnIPs(ctx _context.Context, ippoolId string, requestBody []string, localVarOptionals *IppoolsApiReturnIPsOpts) (IpPool, *_nethttp.Response, error)
	/*
	   Update Update IP pool by ID
	   Update a single ip pool with matching ID. &#39;DefaultRoute&#39; can only be updated if ip pool is not currently in-use. If GreenLake Platform IAM issued token is used for authentication, then it is required to pass  &#39;X-Role&#39; and &#39;X-Workspaceid&#39; headers.
	     - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	     - @param ippoolId ID of IP pool to update
	     - @param updateIpPool Update IPPool
	     - @param optional nil or *IppoolsApiUpdateOpts - Optional Parameters:
	     - @param "XRole" (optional.String) -  GreenLake Platform role
	     - @param "XWorkspaceid" (optional.String) -  GreenLake Platform workspace ID

	   @return IpPool
	*/
	Update(ctx _context.Context, ippoolId string, updateIpPool UpdateIpPool, localVarOptionals *IppoolsApiUpdateOpts) (IpPool, *_nethttp.Response, error)
}
