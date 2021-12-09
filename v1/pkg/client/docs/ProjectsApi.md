# \ProjectsApi

All URIs are relative to *https://quake.dev.hpehcss.net/rest/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Add**](ProjectsApi.md#Add) | **Post** /projects | Create a new project
[**Delete**](ProjectsApi.md#Delete) | **Delete** /projects/{projectId} | Delete a Project
[**GetByID**](ProjectsApi.md#GetByID) | **Get** /projects/{projectId} | Retrieve a project by its ID
[**List**](ProjectsApi.md#List) | **Get** /projects | List of all Projects within an organization or cluster
[**Update**](ProjectsApi.md#Update) | **Put** /projects/{projectId} | Update a project by its ID



## Add

> Project Add(ctx, newProject, optional)

Create a new project

Adds a new Project which creates an isolated space for creating Hosts, Volumes, and private Networks. A project is often aligned to a specific team within an organization or a cluster. If GreenLake IAM issued token is used for authentication, then it is required to pass either 'Space' or 'spaceid' header. When both are set, 'Space' header is ignored. Note that Hoster or BMaaS Access Owner role is required for this operation.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**newProject** | [**NewProject**](NewProject.md)| NewProject parameters to create a new Project | 
 **optional** | ***AddOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **space** | **optional.String**| GreenLake space name | 
 **spaceid** | **optional.String**| GreenLake space ID | 

### Return type

[**Project**](Project.md)

### Authorization

[BearerAuth](../README.md#BearerAuth), [Membership](../README.md#Membership)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(ctx, projectId)

Delete a Project

Deletes the Project with the matching ID. Note that Hoster or BMaaS Access Owner role is required for this operation.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| ID of project to delete | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth), [Membership](../README.md#Membership)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetByID

> Project GetByID(ctx, projectId)

Retrieve a project by its ID

Returns a single Project object with its matching ID This includes profile information for the project and project limits on resouces like hosts, private networks, volumes, and volume capacity.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| ID of project to return | 

### Return type

[**Project**](Project.md)

### Authorization

[BearerAuth](../README.md#BearerAuth), [Membership](../README.md#Membership)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> []Project List(ctx, optional)

List of all Projects within an organization or cluster

Returns an array of all Project objects that have been created. This includes profile information for the project and project limits on resouces like hosts, private networks, volumes, and volume capacity. If GreenLake IAM issued token is used for authentication, then it is required to pass either 'Space' or 'spaceid' header. When both are set, 'Space' header is ignored.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **spaceid** | **optional.String**| GreenLake space ID | 
 **space** | **optional.String**| GreenLake space name | 

### Return type

[**[]Project**](Project.md)

### Authorization

[BearerAuth](../README.md#BearerAuth), [Membership](../README.md#Membership)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Project Update(ctx, projectId, project)

Update a project by its ID

Updates a project with a matching ID. Only Project 'Name', 'Profile' and 'Limits' can be updated with this operation. Note that Hoster or BMaaS Access Owner role is required for this operation.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| ID of project to update | 
**project** | [**Project**](Project.md)| Project parameters to update an existing Project | 

### Return type

[**Project**](Project.md)

### Authorization

[BearerAuth](../README.md#BearerAuth), [Membership](../README.md#Membership)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

