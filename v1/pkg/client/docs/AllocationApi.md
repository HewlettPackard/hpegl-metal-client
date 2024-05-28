# \AllocationApi

All URIs are relative to *https://client.greenlake.hpe.com/api/metal/rest/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBySite**](AllocationApi.md#GetBySite) | **Get** /allocation/servers | Get servers allocation
[**StorageGetBySite**](AllocationApi.md#StorageGetBySite) | **Get** /allocation/storage | Get storage allocation



## GetBySite

> []Allocation GetBySite(ctx, optional)

Get servers allocation

Returns an array of allocation information for each server instance type. Each instance type's allocation information is further grouped by service type. When using a Metal token, the value in the 'Membership' header determines the scope of the response, i.e., Project or Hoster. However, the default scope is Tenant/Hoster when using the GL IAM token.  If the allocation data is needed for a specific project, then the 'Project' header must be present in the request. If the 'siteID' query parameter is present, the information returned is specific to that site ID. If GreenLake Platform IAM issued token is used for authentication, then it is required to pass  'X-Role' and 'X-Worksapceid' headers. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetBySiteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetBySiteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **siteID** | **optional.String**| site ID | 
 **xRole** | **optional.String**| GreenLake Platform role name | 
 **xWorkspaceid** | **optional.String**| GreenLake Platform workspace ID | 

### Return type

[**[]Allocation**](Allocation.md)

### Authorization

[BearerAuth](../README.md#BearerAuth), [Membership](../README.md#Membership), [Project](../README.md#Project)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageGetBySite

> []AllocationStorage StorageGetBySite(ctx, optional)

Get storage allocation

Returns an array of allocation information for each server instance type. Each instance type's allocation information is further grouped by service type. When using a Metal token, the value in the 'Membership' header determines the scope of the response, i.e., Project or Hoster. However, the default scope is Tenant/Hoster when using the GL IAM token.  If the allocation data is needed for a specific project, then the 'Project' header must be present in the request. If the 'siteID' query parameter is present, the information returned is specific to that site ID. If GreenLake Platform IAM issued token is used for authentication, then it is required to pass  'X-Role' and 'X-Worksapceid' headers. 

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
 **xRole** | **optional.String**| GreenLake Platform role name | 
 **xWorkspaceid** | **optional.String**| GreenLake Platform workspace ID | 

### Return type

[**[]AllocationStorage**](AllocationStorage.md)

### Authorization

[BearerAuth](../README.md#BearerAuth), [Membership](../README.md#Membership), [Project](../README.md#Project)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

