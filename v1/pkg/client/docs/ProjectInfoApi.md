# \ProjectInfoApi

All URIs are relative to *http://repurpose for client api version*

Method | HTTP request | Description
------------- | ------------- | -------------
[**List**](ProjectInfoApi.md#List) | **Get** /project-info | List of all Projects info within an organization or cluster for which user is authorized.



## List

> []ProjectsInfo List(ctx, )

List of all Projects info within an organization or cluster for which user is authorized.

Returns an object with info on projects that have been created. This includes information on machine sizes and volumes falvors used by the projects.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]ProjectsInfo**](ProjectsInfo.md)

### Authorization

[quake_auth](../README.md#quake_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

