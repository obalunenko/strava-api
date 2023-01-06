# {{classname}}

All URIs are relative to *https://www.strava.com/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUpload**](UploadsApi.md#CreateUpload) | **Post** /uploads | Upload Activity
[**GetUploadById**](UploadsApi.md#GetUploadById) | **Get** /uploads/{uploadId} | Get Upload

# **CreateUpload**
> Upload CreateUpload(ctx, optional)
Upload Activity

Uploads a new data file to create an activity from. Requires activity:write scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UploadsApiCreateUploadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UploadsApiCreateUploadOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | **optional.Interface of *os.File****optional.**|  | 
 **name** | **optional.**|  | 
 **description** | **optional.**|  | 
 **trainer** | **optional.**|  | 
 **commute** | **optional.**|  | 
 **dataType** | **optional.**|  | 
 **externalId** | **optional.**|  | 

### Return type

[**Upload**](Upload.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUploadById**
> Upload GetUploadById(ctx, uploadId)
Get Upload

Returns an upload for a given identifier. Requires activity:write scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uploadId** | **int64**| The identifier of the upload. | 

### Return type

[**Upload**](Upload.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

