# UpdateHostAlertInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alert** | **string** | Identifies the type of alert | [optional] 
**State** | [**HostState**](HostState.md) |  | [optional] 
**Substate** | [**HostSubstate**](HostSubstate.md) |  | [optional] 
**Message** | **string** | Provides some detailed description about the Alert | [optional] 
**Time** | [**time.Time**](time.Time.md) |  | [optional] 
**Ack** | **bool** | Used to acknowledge the alert so that the UI can list only unacknowledged alerts | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


