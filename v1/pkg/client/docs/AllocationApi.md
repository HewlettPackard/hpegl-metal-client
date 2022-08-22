# \AllocationApi

All URIs are relative to *https://quake.dev.hpehcss.net/rest/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBySite**](AllocationApi.md#GetBySite) | **Get** /allocation/servers/{siteID} | Get servers allocation
[**StorageGetBySite**](AllocationApi.md#StorageGetBySite) | **Get** /allocation/storage/{siteID} | Get storage allocation



## GetBySite

> Allocation GetBySite(ctx, siteID)

Get servers allocation

Returns allocation information for each server instance type used by each PCE service. If siteID is present, the information returned is specific to that site ID, otherwise the allocation information for all sites is returned.

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

> Allocation StorageGetBySite(ctx, optional)

Get storage allocation

Returns allocation information for each volume type used by each PCE service. If siteID is present, the information returned is specific to that site ID, otherwise the allocation information for all sites is returned.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StorageGetBySiteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StorageGetBySiteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **siteID** | **optional.String**| site ID | 

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

