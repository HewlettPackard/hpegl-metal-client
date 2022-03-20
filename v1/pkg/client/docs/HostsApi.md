# \HostsApi

All URIs are relative to *https://quake.dev.hpehcss.net/rest/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Add**](HostsApi.md#Add) | **Post** /hosts | Create a new Host
[**Delete**](HostsApi.md#Delete) | **Delete** /hosts/{hostId} | Delete a Host
[**GetByID**](HostsApi.md#GetByID) | **Get** /hosts/{hostId} | Retrieve Host by ID
[**List**](HostsApi.md#List) | **Get** /hosts | List all Hosts in project
[**Maintenance**](HostsApi.md#Maintenance) | **Post** /hosts/{hostId}/maintenance | Do maintenance on a Host by ID
[**PowerOff**](HostsApi.md#PowerOff) | **Post** /hosts/{hostId}/poweroff | Power off Host by ID
[**PowerOn**](HostsApi.md#PowerOn) | **Post** /hosts/{hostId}/poweron | Power on Host by ID
[**Replace**](HostsApi.md#Replace) | **Post** /hosts/{hostId}/replace | Replace Host by ID
[**Update**](HostsApi.md#Update) | **Put** /hosts/{hostId} | Update an existing Host



## Add

> Host Add(ctx, newHost)

Create a new Host

Creates a new host object which kicks off the provisioning of a physical server in accordance to the attributes provided for the Host object.  Most values for these options must be selected from the set of options provided by the get available-resources API call. The SvcFlavor, SvcVersion, LocationID, SSHKeyIDs, and Network attribute must all be set with appropriate ID values from the available-resources call.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**newHost** | [**NewHost**](NewHost.md)| Defines the configuration of the desired host. See the schema for descriptions of individual attributes. | 

### Return type

[**Host**](Host.md)

### Authorization

[BearerAuth](../README.md#BearerAuth), [Membership](../README.md#Membership), [Project](../README.md#Project)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(ctx, hostId)

Delete a Host

Deletes the Host with the matching ID.  A host in the 'Ready' state must first be powered-off before a delete will be permitted.  Deletes to hosts in other states is permitted regardless of the power state

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **string**| ID of Host to delete | 

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

> Host GetByID(ctx, hostId)

Retrieve Host by ID

Returns a single Host with matching ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **string**| ID of Host to return | 

### Return type

[**Host**](Host.md)

### Authorization

[BearerAuth](../README.md#BearerAuth), [Membership](../README.md#Membership), [Project](../README.md#Project)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> []Host List(ctx, optional)

List all Hosts in project

Returns an array of all Host objects defined within the project.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **all** | **optional.String**| Includes deleted Host objects in the response when set to \&quot;true\&quot;. | 

### Return type

[**[]Host**](Host.md)

### Authorization

[BearerAuth](../README.md#BearerAuth), [Membership](../README.md#Membership), [Project](../README.md#Project)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Maintenance

> Host Maintenance(ctx, hostId)

Do maintenance on a Host by ID

Do maintenance on a host by executing pre-defined operations. The host must be powered off.  The host must also be in the Ready state or in the Failed state and in the Maintenance workflow.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **string**| ID of Host to do maintenance on | 

### Return type

[**Host**](Host.md)

### Authorization

[BearerAuth](../README.md#BearerAuth), [Membership](../README.md#Membership), [Project](../README.md#Project)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PowerOff

> Host PowerOff(ctx, hostId)

Power off Host by ID

Powers off a single Host with matching ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **string**| ID of Host to power off | 

### Return type

[**Host**](Host.md)

### Authorization

[BearerAuth](../README.md#BearerAuth), [Membership](../README.md#Membership), [Project](../README.md#Project)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PowerOn

> Host PowerOn(ctx, hostId)

Power on Host by ID

Powers on a single Host with matching ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **string**| ID of Host to power on | 

### Return type

[**Host**](Host.md)

### Authorization

[BearerAuth](../README.md#BearerAuth), [Membership](../README.md#Membership), [Project](../README.md#Project)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Replace

> Host Replace(ctx, hostId)

Replace Host by ID

Re-deploys a host with a new machine that satisfies the current host settings. Only the machine is replaced, IP addresses, volumes, etc are not changed. The host must be powered off.  The host must also be in the Ready state or in the Failed state and in the Replace or Maintenace workflow.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **string**| ID of Host to replace | 

### Return type

[**Host**](Host.md)

### Authorization

[BearerAuth](../README.md#BearerAuth), [Membership](../README.md#Membership), [Project](../README.md#Project)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Host Update(ctx, hostId, host)

Update an existing Host

Updates the Host with the matching ID.  Update is permitted only if the host is in the 'Ready' or 'Connection Updating Failed' state.  Only the Host 'Description', 'Networks', and 'NetworkForDefaultRoute' can be updated.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **string**| ID of host to update | 
**host** | [**Host**](Host.md)| Updated Host | 

### Return type

[**Host**](Host.md)

### Authorization

[BearerAuth](../README.md#BearerAuth), [Membership](../README.md#Membership), [Project](../README.md#Project)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

