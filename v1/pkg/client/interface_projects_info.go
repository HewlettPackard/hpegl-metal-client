// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

//

package client

import (
	_context "context"
	_nethttp "net/http"
)

// ProjectsInfoAPI defines the client functions provided for ProjectsInfo.
type ProjectsInfoAPI interface {
	/*
	   List List of all projects info within an organization or cluster for which user is authorized.
	   Returns an object with info on projects that have been created. This includes information on machine sizes and volumes falvors used by the projects.
	    * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	   @return ProjectsInfo
	*/
	List(ctx _context.Context) (ProjectsInfo, *_nethttp.Response, error)
}
