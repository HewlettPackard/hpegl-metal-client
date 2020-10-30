# \NetworksApi

All URIs are relative to *http://repurpose for client api version*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Add**](NetworksApi.md#Add) | **Post** /networks | Add a new network
[**Delete**](NetworksApi.md#Delete) | **Delete** /networks/{networkId} | Delete a network
[**GetByID**](NetworksApi.md#GetByID) | **Get** /networks/{networkId} | Retrieve network by ID
[**List**](NetworksApi.md#List) | **Get** /networks | List all networks in project
[**Update**](NetworksApi.md#Update) | **Put** /networks/{networkId} | Update an existing network.  NOT YET SUPPORTED



## Add

> Network Add(ctx, newNetwork)

Add a new network

Adds a new network that can be referenced when creating a Host

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**newNetwork** | [**NewNetwork**](NewNetwork.md)| Network that is to be added to the project | 

### Return type

[**Network**](Network.md)

### Authorization

[quake_auth](../README.md#quake_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(ctx, networkId)

Delete a network

Deletes the network with the matching ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string**| ID of network to delete | 

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

> Network GetByID(ctx, networkId)

Retrieve network by ID

Returns a single network with matching ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string**| ID of network to return | 

### Return type

[**Network**](Network.md)

### Authorization

[quake_auth](../README.md#quake_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> []Network List(ctx, )

List all networks in project

Returns an array of all network objects defined within the project.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]Network**](Network.md)

### Authorization

[quake_auth](../README.md#quake_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Network Update(ctx, networkId, network)

Update an existing network.  NOT YET SUPPORTED

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string**| ID of network to update | 
**network** | [**Network**](Network.md)| Updated network | 

### Return type

[**Network**](Network.md)

### Authorization

[quake_auth](../README.md#quake_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

