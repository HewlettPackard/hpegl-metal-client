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
	   Returns an array of allocation information for each server instance type used by each PCE service. If siteID is present, the information returned is specific to that site ID, otherwise the allocation information for all sites is returned.
	     - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	     - @param optional nil or *AllocationApiGetBySiteOpts - Optional Parameters:
	     - @param "SiteID" (optional.String) -  site ID

	   @return []Allocation
	*/
	GetBySite(ctx _context.Context, localVarOptionals *AllocationApiGetBySiteOpts) ([]Allocation, *_nethttp.Response, error)
	/*
	   StorageGetBySite Get storage allocation
	   Returns an array of allocation information for each volume type used by each PCE service. If siteID is present, the information returned is specific to that site ID, otherwise the allocation information for all sites is returned.
	     - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	     - @param optional nil or *AllocationApiStorageGetBySiteOpts - Optional Parameters:
	     - @param "SiteID" (optional.String) -  site ID

	   @return []AllocationStorage
	*/
	StorageGetBySite(ctx _context.Context, localVarOptionals *AllocationApiStorageGetBySiteOpts) ([]AllocationStorage, *_nethttp.Response, error)
}
