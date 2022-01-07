// (C) Copyright 2021-2022 Hewlett Packard Enterprise Development LP

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
	   Returns an object with information on projects, machine sizes, and volume flavors.  The &#39;Projects&#39; list includes projects authorized for a user, and the &#39;MachineSizes&#39; and  &#39;VolumeFlavors&#39; list include only those machine sizes and volume flavors permitted for projects.  When GreenLake IAM issued token is used for authentication, it is required to  pass either &#39;Space&#39; or &#39;spaceid&#39; header. When both are set, &#39;Space&#39; header is ignored.
	    * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    * @param optional nil or *ProjectsInfoApiListOpts - Optional Parameters:
	    * @param "Space" (optional.String) -  GreenLake space name
	    * @param "Spaceid" (optional.String) -  GreenLake space ID
	   @return ProjectsInfo
	*/
	List(ctx _context.Context, localVarOptionals *ProjectsInfoApiListOpts) (ProjectsInfo, *_nethttp.Response, error)
}
