# \IppoolsApi

All URIs are relative to *https://quake.dev.hpehcss.net/rest/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllocateIPs**](IppoolsApi.md#AllocateIPs) | **Post** /ippools/{ippoolId}/allocation | Allocate IPs from the pool
[**GetByID**](IppoolsApi.md#GetByID) | **Get** /ippools/{ippoolId} | Retrieve IP pool by ID
[**List**](IppoolsApi.md#List) | **Get** /ippools | List all ip pools in project
[**ReturnIPs**](IppoolsApi.md#ReturnIPs) | **Post** /ippools/{ippoolId}/return | Return IPs to the pool
[**Update**](IppoolsApi.md#Update) | **Put** /ippools/{ippoolId} | Update IP pool by ID



## AllocateIPs

> IpPool AllocateIPs(ctx, ippoolId, iPAllocation)

Allocate IPs from the pool

Allocate IPs from the pool

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ippoolId** | **string**| ID of IP pool to allocate IPs | 
**iPAllocation** | [**[]IpAllocation**](IPAllocation.md)| IPs being requested starting from an optional base IP and their usage | 

### Return type

[**IpPool**](IPPool.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetByID

> IpPool GetByID(ctx, ippoolId)

Retrieve IP pool by ID

Returns a single ip pool with matching ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ippoolId** | **string**| ID of IP pool to return | 

### Return type

[**IpPool**](IPPool.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> []IpPool List(ctx, )

List all ip pools in project

Returns an array of all ip pool objects defined within the project.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]IpPool**](IPPool.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnIPs

> IpPool ReturnIPs(ctx, ippoolId, requestBody)

Return IPs to the pool

Return IPs to the pool

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ippoolId** | **string**| ID of IP pool to return IPs | 
**requestBody** | [**[]string**](string.md)| IP returned to the pool | 

### Return type

[**IpPool**](IPPool.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> IpPool Update(ctx, ippoolId, ipPoolUpdate)

Update IP pool by ID

Update a single ip pool with matching ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ippoolId** | **string**| ID of IP pool to update | 
**ipPoolUpdate** | [**IpPoolUpdate**](IpPoolUpdate.md)| Update IPPool | 

### Return type

[**IpPool**](IPPool.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

