# \NetworksApi

All URIs are relative to *https://client.greenlake.hpe.com/api/metal/rest/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Add**](NetworksApi.md#Add) | **Post** /networks | Add a new network
[**Delete**](NetworksApi.md#Delete) | **Delete** /networks/{networkId} | Delete a network
[**GetByID**](NetworksApi.md#GetByID) | **Get** /networks/{networkId} | Retrieve network by ID.
[**List**](NetworksApi.md#List) | **Get** /networks | List all networks in project
[**Update**](NetworksApi.md#Update) | **Put** /networks/{networkId} | Update an existing network by ID.



## Add

> Network Add(ctx, newNetwork, optional)

Add a new network

Adds a new network that can be referenced when creating a Host. If GreenLake Platform IAM issued token is used for authentication, then it is required to pass  'X-Role' and 'X-Workspaceid' headers.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**newNetwork** | [**NewNetwork**](NewNetwork.md)| Network that is to be added to the project | 
 **optional** | ***AddOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRole** | **optional.String**| GreenLake Platform role | 
 **xWorkspaceid** | **optional.String**| GreenLake Platform workspace ID | 

### Return type

[**Network**](Network.md)

### Authorization

[BearerAuth](../README.md#BearerAuth), [Membership](../README.md#Membership), [Project](../README.md#Project), [Role](../README.md#Role), [Workspace](../README.md#Workspace)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(ctx, networkId, optional)

Delete a network

Deletes the network with the matching ID. If GreenLake Platform IAM issued token is used for authentication, then it is required to pass  'X-Role' and 'X-Workspaceid' headers.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string**| ID of network to delete | 
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

[BearerAuth](../README.md#BearerAuth), [Membership](../README.md#Membership), [Project](../README.md#Project), [Role](../README.md#Role), [Workspace](../README.md#Workspace)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetByID

> Network GetByID(ctx, networkId, optional)

Retrieve network by ID.

Returns a single network with matching ID. If GreenLake Platform IAM issued token is used for authentication, then it is required to pass  'X-Role' and 'X-Workspaceid' headers.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string**| ID of network to return | 
 **optional** | ***GetByIDOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetByIDOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRole** | **optional.String**| GreenLake Platform role | 
 **xWorkspaceid** | **optional.String**| GreenLake Platform workspace ID | 

### Return type

[**Network**](Network.md)

### Authorization

[BearerAuth](../README.md#BearerAuth), [Membership](../README.md#Membership), [Project](../README.md#Project), [Role](../README.md#Role), [Workspace](../README.md#Workspace)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> []Network List(ctx, optional)

List all networks in project

Returns an array of all network objects defined within the project. If GreenLake Platform IAM issued token is used for authentication, then it is required to pass  'X-Role' and 'X-Workspaceid' headers.

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

[**[]Network**](Network.md)

### Authorization

[BearerAuth](../README.md#BearerAuth), [Membership](../README.md#Membership), [Project](../README.md#Project), [Role](../README.md#Role), [Workspace](../README.md#Workspace)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Network Update(ctx, networkId, updateNetwork, optional)

Update an existing network by ID.

Update an existing network by ID. If GreenLake Platform IAM issued token is used for authentication, then it is required to pass  'X-Role' and 'X-Workspaceid' headers.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string**| ID of network to update | 
**updateNetwork** | [**UpdateNetwork**](UpdateNetwork.md)| Updated network | 
 **optional** | ***UpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRole** | **optional.String**| GreenLake Platform role | 
 **xWorkspaceid** | **optional.String**| GreenLake Platform workspace ID | 

### Return type

[**Network**](Network.md)

### Authorization

[BearerAuth](../README.md#BearerAuth), [Membership](../README.md#Membership), [Project](../README.md#Project), [Role](../README.md#Role), [Workspace](../README.md#Workspace)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

