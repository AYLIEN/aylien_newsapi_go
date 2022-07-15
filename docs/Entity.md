# Entity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique ID of the entity | [optional] 
**Links** | [**EntityLinks**](EntityLinks.md) |  | [optional] 
**StockTickers** | **[]string** | The stock tickers of the entity (might be empty) | [optional] 
**Types** | **[]string** | An array of the entity types | [optional] 
**OverallSentiment** | [**EntitySentiment**](EntitySentiment.md) |  | [optional] 
**OverallProminence** | **float32** | Describes how relevant an entity is to the article | [optional] 
**OverallFrequency** | **int32** | Amount of entity surface form mentions in the article | [optional] 
**Body** | [**EntityInText**](EntityInText.md) |  | [optional] 
**Title** | [**EntityInText**](EntityInText.md) |  | [optional] 
**ExternalIds** | [**ExternalIds**](ExternalIds.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


