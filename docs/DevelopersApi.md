# \DevelopersApi

All URIs are relative to *https://virtserver.swaggerhub.com/aHisayoshiSuehiro/ssl_manage_server/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchInventory**](DevelopersApi.md#SearchInventory) | **Get** /inventory | searches inventory


# **SearchInventory**
> []InventoryItem SearchInventory($searchString, $skip, $limit)

searches inventory

By passing in the appropriate options, you can search for available inventory in the system 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchString** | **string**| pass an optional search string for looking up inventory | [optional] 
 **skip** | **int32**| number of records to skip for pagination | [optional] 
 **limit** | **int32**| maximum number of records to return | [optional] 

### Return type

[**[]InventoryItem**](InventoryItem.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

