# \IppoolsApi

All URIs are relative to *https://client.greenlake.hpe.com/api/metal/rest/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllocateIPs**](IppoolsApi.md#AllocateIPs) | **Post** /ippools/{ippoolId}/allocation | Allocate IPs from the pool
[**GetByID**](IppoolsApi.md#GetByID) | **Get** /ippools/{ippoolId} | Retrieve IP pool by ID
[**List**](IppoolsApi.md#List) | **Get** /ippools | List all ip pools in project
[**ReturnIPs**](IppoolsApi.md#ReturnIPs) | **Post** /ippools/{ippoolId}/return | Return IPs to the pool
[**Update**](IppoolsApi.md#Update) | **Put** /ippools/{ippoolId} | Update IP pool by ID



## AllocateIPs

> IpPool AllocateIPs(ctx, ippoolId, iPAllocation, optional)

Allocate IPs from the pool

Allocate IPs from the pool If GreenLake Platform IAM issued token is used for authentication, then it is required to pass  'X-Role' and 'X-Workspaceid' headers.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ippoolId** | **string**| ID of IP pool to allocate IPs | 
**iPAllocation** | [**[]IpAllocation**](IPAllocation.md)| IPs being requested starting from an optional base IP and their usage | 
 **optional** | ***AllocateIPsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AllocateIPsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRole** | **optional.String**| GreenLake Platform role | 
 **xWorkspaceid** | **optional.String**| GreenLake Platform workspace ID | 

### Return type

[**IpPool**](IPPool.md)

### Authorization

[BearerAuth](../README.md#BearerAuth), [Membership](../README.md#Membership), [Project](../README.md#Project), [Role](../README.md#Role), [Workspace](../README.md#Workspace)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetByID

> IpPool GetByID(ctx, ippoolId, optional)

Retrieve IP pool by ID

Returns a single ip pool with matching imaged. If GreenLake Platform IAM issued token is used for authentication, then it is required to pass  'X-Role' and 'X-Workspaceid' headers.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ippoolId** | **string**| ID of IP pool to return | 
 **optional** | ***GetByIDOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetByIDOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRole** | **optional.String**| GreenLake Platform role | 
 **xWorkspaceid** | **optional.String**| GreenLake Platform workspace ID | 

### Return type

[**IpPool**](IPPool.md)

### Authorization

[BearerAuth](../README.md#BearerAuth), [Membership](../README.md#Membership), [Project](../README.md#Project), [Role](../README.md#Role), [Workspace](../README.md#Workspace)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> []IpPool List(ctx, optional)

List all ip pools in project

Returns an array of all ip pool objects defined within the project. If GreenLake Platform IAM issued token is used for authentication, then it is required to pass  'X-Role' and 'X-Workspaceid' headers.

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

[**[]IpPool**](IPPool.md)

### Authorization

[BearerAuth](../README.md#BearerAuth), [Membership](../README.md#Membership), [Project](../README.md#Project), [Role](../README.md#Role), [Workspace](../README.md#Workspace)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnIPs

> IpPool ReturnIPs(ctx, ippoolId, requestBody, optional)

Return IPs to the pool

Return IPs to the pool. If GreenLake Platform IAM issued token is used for authentication, then it is required to pass  'X-Role' and 'X-Workspaceid' headers.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ippoolId** | **string**| ID of IP pool to return IPs | 
**requestBody** | [**[]string**](string.md)| IP returned to the pool | 
 **optional** | ***ReturnIPsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReturnIPsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRole** | **optional.String**| GreenLake Platform role | 
 **xWorkspaceid** | **optional.String**| GreenLake Platform workspace ID | 

### Return type

[**IpPool**](IPPool.md)

### Authorization

[BearerAuth](../README.md#BearerAuth), [Membership](../README.md#Membership), [Project](../README.md#Project), [Role](../README.md#Role), [Workspace](../README.md#Workspace)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> IpPool Update(ctx, ippoolId, updateIpPool, optional)

Update IP pool by ID

Update a single ip pool with matching ID. 'DefaultRoute' can only be updated if ip pool is not currently in-use. If GreenLake Platform IAM issued token is used for authentication, then it is required to pass  'X-Role' and 'X-Workspaceid' headers.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ippoolId** | **string**| ID of IP pool to update | 
**updateIpPool** | [**UpdateIpPool**](UpdateIpPool.md)| Update IPPool | 
 **optional** | ***UpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRole** | **optional.String**| GreenLake Platform role | 
 **xWorkspaceid** | **optional.String**| GreenLake Platform workspace ID | 

### Return type

[**IpPool**](IPPool.md)

### Authorization

[BearerAuth](../README.md#BearerAuth), [Membership](../README.md#Membership), [Project](../README.md#Project), [Role](../README.md#Role), [Workspace](../README.md#Workspace)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

