// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package client

import (
	_context "context"
	_nethttp "net/http"
)

// VersionAPI defines the functions provided by the client for Versions.
type VersionAPI interface {
	/*
	   Get Get api server build version
	    * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	   @return Version
	*/
	Get(ctx _context.Context) (Version, *_nethttp.Response, error)
}
