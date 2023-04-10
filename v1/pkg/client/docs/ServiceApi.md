# \ServiceApi

All URIs are relative to *https://client.greenlake.hpe.com/api/metal/rest/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Add**](ServiceApi.md#Add) | **Post** /services | Create a new OS service image



## Add

> OsServiceImage Add(ctx, fileName, optional)

Create a new OS service image

Adds a new OS service image that can be referenced during host creation. If GreenLake IAM issued token is used for authentication, then it is required  to pass either 'spaceid' header or 'Space' header.  Note that Hoster or BMaaS Access Owner role is required for this operation.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileName** | ***os.File*****os.File**|  | 
 **optional** | ***AddOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **spaceid** | **optional.String**| GreenLake space ID | 

### Return type

[**OsServiceImage**](OSServiceImage.md)

### Authorization

[BearerAuth](../README.md#BearerAuth), [Membership](../README.md#Membership)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

