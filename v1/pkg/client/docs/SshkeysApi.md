# \SshkeysApi

All URIs are relative to *https://client.greenlake.hpe.com/api/metal/rest/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Add**](SshkeysApi.md#Add) | **Post** /sshkeys | Add a new SSH Key
[**Delete**](SshkeysApi.md#Delete) | **Delete** /sshkeys/{sshkeyId} | Delete an SSH key by ID.
[**GetByID**](SshkeysApi.md#GetByID) | **Get** /sshkeys/{sshkeyId} | Retrieve SSH Key by ID
[**List**](SshkeysApi.md#List) | **Get** /sshkeys | List all sshkeys in project
[**Update**](SshkeysApi.md#Update) | **Put** /sshkeys/{sshkeyId} | Update an existing SSH Key by ID.



## Add

> SshKey Add(ctx, newSshKey, optional)

Add a new SSH Key

Adds a new SSH Key that can be referenced when creating a Host. If GreenLake Platform IAM issued token is used for authentication, then it is required to pass  'X-Role' and 'X-Workspaceid' headers.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**newSshKey** | [**NewSshKey**](NewSshKey.md)| SSH Key that is to be added to the project | 
 **optional** | ***AddOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRole** | **optional.String**| GreenLake Platform role | 
 **xWorkspaceid** | **optional.String**| GreenLake Platform workspace ID | 

### Return type

[**SshKey**](SSHKey.md)

### Authorization

[BearerAuth](../README.md#BearerAuth), [Membership](../README.md#Membership), [Project](../README.md#Project)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(ctx, sshkeyId, optional)

Delete an SSH key by ID.

Deletes the SSH key with the matching ID. If GreenLake Platform IAM issued token is used for authentication, then it is required to pass  'X-Role' and 'X-Workspaceid' headers.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sshkeyId** | **string**| ID of sshkey to delete | 
 **optional** | ***DeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRole** | **optional.String**| GreenLake Platform role | 
 **xWorkspaceid** | **optional.String**| GreenLake Platform workspace ID | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth), [Membership](../README.md#Membership), [Project](../README.md#Project)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetByID

> SshKey GetByID(ctx, sshkeyId, optional)

Retrieve SSH Key by ID

Returns a single SSH key with matching ID. If GreenLake Platform IAM issued token is used for authentication, then it is required to pass  'X-Role' and 'X-Workspaceid' headers.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sshkeyId** | **string**| ID of sshkey to return | 
 **optional** | ***GetByIDOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetByIDOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRole** | **optional.String**| GreenLake Platform role | 
 **xWorkspaceid** | **optional.String**| GreenLake Platform workspace ID | 

### Return type

[**SshKey**](SSHKey.md)

### Authorization

[BearerAuth](../README.md#BearerAuth), [Membership](../README.md#Membership), [Project](../README.md#Project)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> []SshKey List(ctx, optional)

List all sshkeys in project

Returns an array of all SSHKey objects defined within the project. If GreenLake Platform IAM issued token is used for authentication, then it is required to pass  'X-Role' and 'X-Workspaceid' headers.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRole** | **optional.String**| GreenLake Platform role | 
 **xWorkspaceid** | **optional.String**| GreenLake Platform workspace ID | 

### Return type

[**[]SshKey**](SSHKey.md)

### Authorization

[BearerAuth](../README.md#BearerAuth), [Membership](../README.md#Membership), [Project](../README.md#Project)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> SshKey Update(ctx, sshkeyId, updateSshKey, optional)

Update an existing SSH Key by ID.

Update a single SSH key with matching ID. Only 'Name' or 'Key' fields can be changed. If GreenLake Platform IAM issued token is used for authentication, then it is required to pass  'X-Role' and 'X-Workspaceid' headers.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sshkeyId** | **string**| ID of sshkey to update | 
**updateSshKey** | [**UpdateSshKey**](UpdateSshKey.md)| Updated SSH key | 
 **optional** | ***UpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRole** | **optional.String**| GreenLake Platform role | 
 **xWorkspaceid** | **optional.String**| GreenLake Platform workspace ID | 

### Return type

[**SshKey**](SSHKey.md)

### Authorization

[BearerAuth](../README.md#BearerAuth), [Membership](../README.md#Membership), [Project](../README.md#Project)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

