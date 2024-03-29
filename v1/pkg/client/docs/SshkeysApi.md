# \SshkeysApi

All URIs are relative to *https://client.greenlake.hpe.com/api/metal/rest/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Add**](SshkeysApi.md#Add) | **Post** /sshkeys | Add a new SSH Key
[**Delete**](SshkeysApi.md#Delete) | **Delete** /sshkeys/{sshkeyId} | Delete an SSH key
[**GetByID**](SshkeysApi.md#GetByID) | **Get** /sshkeys/{sshkeyId} | Retrieve SSH Key by ID
[**List**](SshkeysApi.md#List) | **Get** /sshkeys | List all sshkeys in project
[**Update**](SshkeysApi.md#Update) | **Put** /sshkeys/{sshkeyId} | Update an existing SSH Key.  Only &#39;Name&#39; or &#39;Key&#39; fields can be changed.



## Add

> SshKey Add(ctx, newSshKey)

Add a new SSH Key

Adds a new SSH Key that can be referenced when creating a Host

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**newSshKey** | [**NewSshKey**](NewSshKey.md)| SSH Key that is to be added to the project | 

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

> Delete(ctx, sshkeyId)

Delete an SSH key

Deletes the SSH key with the matching ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sshkeyId** | **string**| ID of sshkey to delete | 

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

> SshKey GetByID(ctx, sshkeyId)

Retrieve SSH Key by ID

Returns a single SSH key with matching ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sshkeyId** | **string**| ID of sshkey to return | 

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

> []SshKey List(ctx, )

List all sshkeys in project

Returns an array of all SSHKey objects defined within the project.

### Required Parameters

This endpoint does not need any parameter.

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

> SshKey Update(ctx, sshkeyId, updateSshKey)

Update an existing SSH Key.  Only 'Name' or 'Key' fields can be changed.



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sshkeyId** | **string**| ID of sshkey to update | 
**updateSshKey** | [**UpdateSshKey**](UpdateSshKey.md)| Updated SSH key | 

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

