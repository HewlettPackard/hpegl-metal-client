# \UsageReportsApi

All URIs are relative to *http://localhost:8080/rest/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](UsageReportsApi.md#Get) | **Get** /usage-reports | Get a usage report



## Get

> UsageReport Get(ctx, start, optional)

Get a usage report

Creates and returns a usage report based on the parameters passed in the request body 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**start** | **string**| Start of the billing period | 
 **optional** | ***GetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **end** | **optional.String**| End of the billing period, default to now if omitted | 

### Return type

[**UsageReport**](UsageReport.md)

### Authorization

[quake_auth](../README.md#quake_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

