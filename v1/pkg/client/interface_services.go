// (C) Copyright 2021-2023 Hewlett Packard Enterprise Development LP

// Code generated by ifacemaker; DO NOT EDIT.

package client

import (
	_context "context"
	_nethttp "net/http"
	"os"
)

// ServicesAPI defines the client functions provided for Services.
type ServicesAPI interface {
	/*
	   Add Create a new OS service image
	   Adds a new OS service image that can be referenced during host creation. If GreenLake IAM issued token is used for authentication, then it is required  to pass either &#39;spaceid&#39; header or &#39;Space&#39; header.  Note that Hoster or BMaaS Access Owner role is required for this operation.
	     - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	     - @param fileName
	     - @param optional nil or *ServicesApiAddOpts - Optional Parameters:
	     - @param "Spaceid" (optional.String) -  GreenLake space ID
	     - @param "Space" (optional.String) -  GreenLake space name

	   @return OsServiceImage
	*/
	Add(ctx _context.Context, fileName *os.File, localVarOptionals *ServicesApiAddOpts) (OsServiceImage, *_nethttp.Response, error)
	/*
	   Delete Delete an OS service image
	   Deletes the OS service image with the matching ID. Note that Hoster or BMaaS Access Owner role is required for this operation.
	     - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	     - @param serviceId ID of OS service image to delete
	*/
	Delete(ctx _context.Context, serviceId string) (*_nethttp.Response, error)
	/*
	   GetByID Retrieve an OS service image
	   Returns a single OS service image object with its matching ID.
	     - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	     - @param serviceId ID of OS service image to return

	   @return OsServiceImage
	*/
	GetByID(ctx _context.Context, serviceId string) (OsServiceImage, *_nethttp.Response, error)
	/*
	   List List of all OS service images within an tenant
	   Returns an array of all OS service images objects that have been created. If GreenLake IAM issued token is used for authentication,  then it is required to pass &#39;spaceid&#39; header
	     - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	     - @param optional nil or *ServicesApiListOpts - Optional Parameters:
	     - @param "Spaceid" (optional.String) -  GreenLake space ID
	     - @param "Space" (optional.String) -  GreenLake space name

	   @return []OsServiceImage
	*/
	List(ctx _context.Context, localVarOptionals *ServicesApiListOpts) ([]OsServiceImage, *_nethttp.Response, error)
	/*
	   Update Update an OS service image by its ID
	   Updates an OS service image with a matching ID. Note that Hoster or BMaaS Access Owner role is required for this operation.
	     - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	     - @param serviceId ID of OS service image to update
	     - @param fileName

	   @return OsServiceImage
	*/
	Update(ctx _context.Context, serviceId string, fileName *os.File) (OsServiceImage, *_nethttp.Response, error)
}
