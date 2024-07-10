# FileInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | **string** | Unique ID of this FileInfo. | 
**Path** | **string** | Depending upon the service approach, this is either the mount-path where the file should be placed or  the relative URL where it should be served. | 
**FileSize** | **int32** | Size of the files in bytes. | 
**DownloadTimeout** | **int64** | Maximum amount of time in seconds to download the Service image. | 
**Signature** | **string** | The signature (checksum) of the image to download.  This ensure the integrity and authenticity of the images downloaded. | 
**Algorithm** | [**Algorithm**](Algorithm.md) |  | 
**Expand** | **bool** | Indicates if the file is compressed and should be expanded  using the filename suffix in the Path as a guide. | 
**DisplayURL** | **string** | URL of the file that should be returned in REST response or used for display purpose. The file is downloaded from this URL only when SecureURL is not set. | 
**SecureURL** | **string** | URL of the file that should be kept secret.  If this field is set, it will be used for accessing the file and DisplayURL will be ignored. | 
**SkipSslVerify** | **bool** | Indicates if the web server the file is being downloaded from should have the SSL certificate validation bypassed. Useful for downloading from an internal webserver with either self-signed or internal CA issued SSL certificate. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


