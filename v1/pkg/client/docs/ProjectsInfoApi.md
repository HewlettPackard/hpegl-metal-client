# \ProjectsInfoApi

All URIs are relative to *https://client.greenlake.hpe.com/api/metal/rest/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**List**](ProjectsInfoApi.md#List) | **Get** /projects-info | List of all projects info within an organization or cluster for which user is authorized.



## List

> ProjectsInfo List(ctx, optional)

List of all projects info within an organization or cluster for which user is authorized.

Returns an object with information on projects, machine sizes, and volume flavors.  The 'Projects' list includes projects authorized for a user, and the 'MachineSizes' and  'VolumeFlavors' list include only those machine sizes and volume flavors permitted for projects. When GreenLake Cloud Service IAM issued token is used for authentication, it is required to  pass either 'Space' or 'spaceid' header. When both are set, 'Space' header is ignored. If GreenLake Platform IAM issued token is used for authentication, then it is required to pass  'X-Role' and 'X-Workspaceid' headers.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **space** | **optional.String**| GreenLake Cloud Services space name | 
 **spaceid** | **optional.String**| GreenLake Cloud Services space ID | 
 **siteid** | **optional.String**| GreenLake site ID | 
 **xRole** | **optional.String**| GreenLake Platform role name | 
 **xWorkspaceid** | **optional.String**| GreenLake Platform workspace ID | 

### Return type

[**ProjectsInfo**](ProjectsInfo.md)

### Authorization

[BearerAuth](../README.md#BearerAuth), [Membership](../README.md#Membership)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

