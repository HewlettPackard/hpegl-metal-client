// (C) Copyright 2021-2023 Hewlett Packard Enterprise Development LP

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
	   Returns an array of allocation information for each server instance type. Each instance type&#39;s allocation information is further grouped by service type. When using a Metal token, the value in the &#39;Membership&#39; header determines the scope of the response, i.e., Project or Hoster. However, the default scope is Tenant/Hoster when using the GL IAM token.  If the allocation data is needed for a specific project, then the &#39;Project&#39; header must be present in the request. If the &#39;siteID&#39; query parameter is present, the information returned is specific to that site ID. If GreenLake Platform IAM issued token is used for authentication, then it is required to pass  &#39;X-Role&#39; header.
	     - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	     - @param optional nil or *AllocationApiGetBySiteOpts - Optional Parameters:
	     - @param "SiteID" (optional.String) -  site ID
	     - @param "XRole" (optional.String) -  GreenLake Platform role

	   @return []Allocation
	*/
	GetBySite(ctx _context.Context, localVarOptionals *AllocationApiGetBySiteOpts) ([]Allocation, *_nethttp.Response, error)
	/*
	   StorageGetBySite Get storage allocation
	   Returns an array of allocation information for each server instance type. Each instance type&#39;s allocation information is further grouped by service type. When using a Metal token, the value in the &#39;Membership&#39; header determines the scope of the response, i.e., Project or Hoster. However, the default scope is Tenant/Hoster when using the GL IAM token.  If the allocation data is needed for a specific project, then the &#39;Project&#39; header must be present in the request. If the &#39;siteID&#39; query parameter is present, the information returned is specific to that site ID. If GreenLake Platform IAM issued token is used for authentication, then it is required to pass  &#39;X-Role&#39; header.
	     - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	     - @param optional nil or *AllocationApiStorageGetBySiteOpts - Optional Parameters:
	     - @param "SiteID" (optional.String) -  site ID
	     - @param "XRole" (optional.String) -  GreenLake Platform role

	   @return []AllocationStorage
	*/
	StorageGetBySite(ctx _context.Context, localVarOptionals *AllocationApiStorageGetBySiteOpts) ([]AllocationStorage, *_nethttp.Response, error)
}
