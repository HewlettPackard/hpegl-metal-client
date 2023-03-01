# PartitionFormula

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Equation** | **string** | Represents a partition size in MiB.  As a special case a size of 0 can be specified that will equate to the remaining space on the disk. | 
**MinSize** | **string** | It causes the value computed by formula execution to be rounded up to MinSize; if MinSize is non-zero. Multipliers like M and G maybe used, e.g. 1M &#x3D;&#x3D; (1024*1024)bytes. An empty string means zero bytes.  If no multiplier is specified M is assumed. | 
**MaxSize** | **string** | MaxSize causes the value computed by formula execution to be rounded up to MaxSize; if MaxSize is non-zero. Multipliers like M, G and T maybe used, e.g. 1M &#x3D;&#x3D; (1024*1024)bytes. An empty string means zero bytes.  If no multiplier is specified M is assumed. Equations may also be used e.g. 25%disk means MaxSize will cause any partition  to never exceed 25% of device capacity. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


