// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

//

package client

import (
	_context "context"
	_nethttp "net/http"
)

// ProjectsAPI defines the client functions provided for Projects.
type ProjectsAPI interface {
	/*
	   Add Create a new project
	   Adds a new Project which creates an isolated space for creating Hosts, Volumes, and private Networks. A project is often aligned to a specific team within an organization or a cluster
	    * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    * @param newProject NewProject parameters to create a new Project
	   @return Project
	*/
	Add(ctx _context.Context, newProject NewProject) (Project, *_nethttp.Response, error)
	/*
	   Delete Delete a Project
	   Deletes the Project with the matching ID
	    * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    * @param projectId ID of project to return
	*/
	Delete(ctx _context.Context, projectId string) (*_nethttp.Response, error)
	/*
	   GetByID Retrieve a project by its ID
	   Returns a single Project object with its matching ID This includes profile information for the project and project limits on resouces like hosts, private networks, volumes, and volume capacity.
	    * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    * @param projectId ID of project to return
	   @return Project
	*/
	GetByID(ctx _context.Context, projectId string) (Project, *_nethttp.Response, error)
	/*
	   List List of all Projects within an organization or cluster
	   Returns an array of all Project objects that have been created. This includes profile information for the project and project limits on resouces like hosts, private networks, volumes, and volume capacity.
	    * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	   @return []Project
	*/
	List(ctx _context.Context) ([]Project, *_nethttp.Response, error)
	/*
	   Update Update a project by its ID
	   Updates a project with a matching ID. Project profile limits can be updated with this operation.
	    * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    * @param projectId ID of project to return
	    * @param project Project parameters to update an existing Project
	   @return Project
	*/
	Update(ctx _context.Context, projectId string, project Project) (Project, *_nethttp.Response, error)
}
