# \ProjectsInfoApi

All URIs are relative to *https://quake.dev.hpehcss.net/rest/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**List**](ProjectsInfoApi.md#List) | **Get** /projects-info | List of all projects info within an organization or cluster for which user is authorized.



## List

> ProjectsInfo List(ctx, )

List of all projects info within an organization or cluster for which user is authorized.

Returns an object with info on projects that have been created. This includes information on machine sizes and volumes falvors used by the projects.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ProjectsInfo**](ProjectsInfo.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

