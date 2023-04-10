// (C) Copyright 2021-2023 Hewlett Packard Enterprise Development LP

// Code generated by ifacemaker; DO NOT EDIT.

package client

import (
	_context "context"
	_nethttp "net/http"
	"os"
)

// ServiceAPI defines the client functions provided for Service.
type ServiceAPI interface {
	/*
	   Add Create a new OS service image
	   Adds a new OS service image that can be referenced during host creation. If GreenLake IAM issued token is used for authentication, then it is required  to pass either &#39;spaceid&#39; header or &#39;Space&#39; header.  Note that Hoster or BMaaS Access Owner role is required for this operation.
	     - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	     - @param fileName
	     - @param optional nil or *ServiceApiAddOpts - Optional Parameters:
	     - @param "Spaceid" (optional.String) -  GreenLake space ID

	   @return OsServiceImage
	*/
	Add(ctx _context.Context, fileName *os.File, localVarOptionals *ServiceApiAddOpts) (OsServiceImage, *_nethttp.Response, error)
}
