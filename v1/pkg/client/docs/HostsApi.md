# \HostsApi

All URIs are relative to *http://localhost:8080/rest/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Add**](HostsApi.md#Add) | **Post** /hosts | Create a new Host
[**Delete**](HostsApi.md#Delete) | **Delete** /hosts/{hostId} | Delete a Host
[**GetByID**](HostsApi.md#GetByID) | **Get** /hosts/{hostId} | Retrieve Host by ID
[**List**](HostsApi.md#List) | **Get** /hosts | List all Hosts in project
[**PowerOff**](HostsApi.md#PowerOff) | **Post** /hosts/{hostId}/poweroff | Power off Host by ID
[**PowerOn**](HostsApi.md#PowerOn) | **Post** /hosts/{hostId}/poweron | Power on Host by ID
[**Update**](HostsApi.md#Update) | **Put** /hosts/{hostId} | Update an existing Host -- NOT SUPPORTED



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

[quake_auth](../README.md#quake_auth)

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

[quake_auth](../README.md#quake_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

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

[quake_auth](../README.md#quake_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> []Host List(ctx, )

List all Hosts in project

Returns an array of all Host objects defined within the project.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]Host**](Host.md)

### Authorization

[quake_auth](../README.md#quake_auth)

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

[quake_auth](../README.md#quake_auth)

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

[quake_auth](../README.md#quake_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Host Update(ctx, hostId, host)

Update an existing Host -- NOT SUPPORTED

NOT CURRENTLY SUPPORTED.  This call will (eventually) allow users to update a limited number of fields associated with the host.  Since most of this information is used when initially provisioning the host, supporting later changes would require careful coordination with host-based agents.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **string**| ID of host to update | 
**host** | [**Host**](Host.md)| Updated Host | 

### Return type

[**Host**](Host.md)

### Authorization

[quake_auth](../README.md#quake_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

