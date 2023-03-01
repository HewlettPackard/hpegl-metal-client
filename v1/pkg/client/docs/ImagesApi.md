# \ImagesApi

All URIs are relative to *https://client.greenlake.hpe.com/api/metal/rest/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Add**](ImagesApi.md#Add) | **Post** /images | Create a new OS image
[**Delete**](ImagesApi.md#Delete) | **Delete** /images/{imageId} | Delete an OS image
[**GetByID**](ImagesApi.md#GetByID) | **Get** /images/{imageId} | Retrieve an OS image its ID
[**List**](ImagesApi.md#List) | **Get** /images | List of all OS Images within an tenant
[**Update**](ImagesApi.md#Update) | **Post** /images/{imageId} | Update an OS image by its ID



## Add

> OsServiceImage Add(ctx, fileName, optional)

Create a new OS image

Adds a new OS Image that can be referenced during host creation. If GreenLake IAM issued token is used for authentication, then it is required  to pass either 'spaceid' header. Note that Hoster or BMaaS Access Owner role is  required for this operation.

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


## Delete

> Delete(ctx, imageId)

Delete an OS image

Deletes the OS image with the matching ID. Note that Hoster or BMaaS Access Owner role is required for this operation.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string**| ID of OS image to delete | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth), [Membership](../README.md#Membership)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetByID

> OsServiceImage GetByID(ctx, imageId)

Retrieve an OS image its ID

Returns a single Os Image object with its matching ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string**| ID of OS image to return | 

### Return type

[**OsServiceImage**](OSServiceImage.md)

### Authorization

[BearerAuth](../README.md#BearerAuth), [Membership](../README.md#Membership)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> []OsServiceImage List(ctx, optional)

List of all OS Images within an tenant

Returns an array of all OS images objects that have been created. If GreenLake IAM issued token is used for authentication,  then it is required to pass 'spaceid' header

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **spaceid** | **optional.String**| GreenLake space ID | 

### Return type

[**[]OsServiceImage**](OSServiceImage.md)

### Authorization

[BearerAuth](../README.md#BearerAuth), [Membership](../README.md#Membership)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> OsServiceImage Update(ctx, imageId, fileName)

Update an OS image by its ID

Updates an OS Image with a matching ID. Note that Hoster or BMaaS Access Owner role is required for this operation.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string**| ID of OS image to update | 
**fileName** | ***os.File*****os.File**|  | 

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

