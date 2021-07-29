# \AutocompleteApi

All URIs are relative to *https://api.aylien.com/news*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAutocompletes**](AutocompleteApi.md#ListAutocompletes) | **Get** /autocompletes | List autocompletes



## ListAutocompletes

> Autocompletes ListAutocompletes(ctx, type_, term, optional)
List autocompletes

The autocompletes endpoint a string of characters provided to it, and then returns suggested terms that are the most likely full words or strings. The terms returned by the News API are based on parameters the type parameters you can see below. For example, if you provide the autocompletes endpoint with the term `New York C` and select the type `dbpedia_resources`, the API will return links to the DBpedia resources `New_York_City`, `New_York_City_Subway`, `New_York_City_Police_Department`, and so on. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string**| This parameter is used for defining the type of autocompletes.  | 
**term** | **string**| This parameter is used for finding autocomplete objects that contain the specified value.  | 
 **optional** | ***ListAutocompletesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAutocompletesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **language** | **optional.String**| This parameter is used for autocompletes whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes.  | [default to en]
 **perPage** | **optional.Int32**| This parameter is used for specifying number of items in each page.  | [default to 3]

### Return type

[**Autocompletes**](Autocompletes.md)

### Authorization

[app_id](../README.md#app_id), [app_key](../README.md#app_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

