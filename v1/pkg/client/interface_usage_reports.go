// (C) Copyright 2021-2022 Hewlett Packard Enterprise Development LP

//

package client

import (
	_context "context"
	_nethttp "net/http"
)

// UsageReportsAPI defines the client functions provided for UsageReports.
type UsageReportsAPI interface {
	/*
	   Get Get a usage report
	   Creates and returns a usage report based on the parameters passed in the request body
	    * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    * @param start Start of the billing period
	    * @param optional nil or *UsageReportsApiGetOpts - Optional Parameters:
	    * @param "End" (optional.String) -  End of the billing period, default to now if omitted
	   @return UsageReport
	*/
	Get(ctx _context.Context, start string, localVarOptionals *UsageReportsApiGetOpts) (UsageReport, *_nethttp.Response, error)
}
