# IscsiParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostIPAddress** | **string** | The IP address of the host in dotted notation. | 
**InitiatorName** | **string** | The full initiator name to be created. The name must be at least 12 characters in length and begin with \&quot;iqn.\&quot;. | 
**CHAPSecret** | **string** | CHAPSecret is the Challenge Authentication Protocol secret to be shared between array and initiator. If empty, no CHAP login is enabled; if set it must be a string between 12 and 16 characters. | [optional] 
**CHAPUserName** | **string** | CHAPUserName is the CHAP username to use for CHAP authentication. If CHAPSecret is specified, CHAPUserName must also be specified. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


