// (C) Copyright 2021-2022 Hewlett Packard Enterprise Development LP

// Code generated by ifacemaker; DO NOT EDIT.

package client

import (
	_context "context"
	_nethttp "net/http"
)

// AllocationAPI defines the client functions provided for Allocation.
type AllocationAPI interface {
	/*
	   GetBySite Get servers allocation
	   Returns allocation information for each server instance type used by each PCE service. If siteID is present, the information returned is specific to that site ID, otherwise the allocation information for all sites is returned.
	    * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    * @param siteID site ID
	   @return Allocation
	*/
	GetBySite(ctx _context.Context, siteID string) (Allocation, *_nethttp.Response, error)
	/*
	   StorageGetBySite Get storage allocation
	   Returns allocation information for each volume type used by each PCE service. If siteID is present, the information returned is specific to that site ID, otherwise the allocation information for all sites is returned.
	    * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    * @param siteID site ID
	   @return Allocation
	*/
	StorageGetBySite(ctx _context.Context, siteID string) (Allocation, *_nethttp.Response, error)
}
