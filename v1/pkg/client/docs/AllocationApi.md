# \AllocationApi

All URIs are relative to *https://quake.dev.hpehcss.net/rest/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBySite**](AllocationApi.md#GetBySite) | **Get** /allocation/servers/{siteID} | Get allocation for servers
[**StorageGetBySite**](AllocationApi.md#StorageGetBySite) | **Get** /allocation/storage/{siteID} | Get allocation for storage



## GetBySite

> Allocation GetBySite(ctx, siteID)

Get allocation for servers

Returns allocation information, for each server instance type used by each PCE service.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteID** | **string**| site ID | 

### Return type

[**Allocation**](Allocation.md)

### Authorization

[BearerAuth](../README.md#BearerAuth), [Membership](../README.md#Membership), [Project](../README.md#Project)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageGetBySite

> Allocation StorageGetBySite(ctx, siteID)

Get allocation for storage

Returns allocation information, for each volume type used by each PCE service.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteID** | **string**| site ID | 

### Return type

[**Allocation**](Allocation.md)

### Authorization

[BearerAuth](../README.md#BearerAuth), [Membership](../README.md#Membership), [Project](../README.md#Project)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

