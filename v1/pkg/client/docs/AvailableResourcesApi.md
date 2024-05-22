# \AvailableResourcesApi

All URIs are relative to *https://client.greenlake.hpe.com/api/metal/rest/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**List**](AvailableResourcesApi.md#List) | **Get** /available-resources | Get lists of available resources for creating hosts and volumes



## List

> AvailableResources List(ctx, optional)

Get lists of available resources for creating hosts and volumes

Used to get lists of options that are used when creating hosts and volumes. A get /available-resources will return an object that includes the following arrays: * Images - A list of image service IDs along with their category (Linux),    flavor (ubuntu), and version (18.04)  * MachineSizes - A list of machine size IDs along with the machine size    names and detailed descriptions  * Locations - A list of location IDs along with their country, region,    and data center.  * Networks - A list of available Network IDs along with the network name,   location ID, network kind, and host usage (Required, Default, Optional)  * MachineInventory - Information about the available inventory of machines    based on location ID and machine size ID.  While this information may    change rapidly, it can be used by GUIs and systems to restrict host   creates to locations with the desired machine size.  * StorageInventory - Information about the current available storage capacity    for a specific volume flavor by site.   * VolumeFlavors - A list of volume flavor IDs along with their name and    detailed description.  * Volumes - A list of current, existing volumes.  If the volume is in the   the right state, it could be attached to a new Host.  If GreenLake Platform IAM issued token is used for authentication, then it is required to pass  'X-Role' and 'X-Workspaceid' headers. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRole** | **optional.String**| GreenLake Platform role name | 
 **xWorkspaceid** | **optional.String**| GreenLake Platform workspace ID | 

### Return type

[**AvailableResources**](AvailableResources.md)

### Authorization

[BearerAuth](../README.md#BearerAuth), [Membership](../README.md#Membership), [Project](../README.md#Project)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

