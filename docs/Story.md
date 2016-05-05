# Story

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | ID of the story which is unique identification | [optional] [default to null]
**Title** | **string** | Title of the story | [optional] [default to null]
**Body** | **string** | Body of the story | [optional] [default to null]
**Summary** | [**Summary**](Summary.md) | The suggested story summary | [optional] [default to null]
**Source** | [**Source**](Source.md) | The story source | [optional] [default to null]
**Author** | [**Author**](Author.md) | The story author | [optional] [default to null]
**Entities** | [**Entities**](Entities.md) | Extracted entities from the story title or body | [optional] [default to null]
**Keywords** | **[]string** | Extracted keywords mentioned in the story title or body | [optional] [default to null]
**Hashtags** | **[]string** | An array of suggested Story hashtags | [optional] [default to null]
**CharactersCount** | **int32** | Characters count of the story body | [optional] [default to null]
**WordsCount** | **int32** | Words count of the story body | [optional] [default to null]
**SentencesCount** | **int32** | Sentences count of the story body | [optional] [default to null]
**ParagraphsCount** | **int32** | Paragraphs count of the story body | [optional] [default to null]
**Categories** | [**[]Category**](Category.md) | Suggested categories for the story | [optional] [default to null]
**SocialSharesCount** | [**ShareCounts**](ShareCounts.md) | Social shares count for the story | [optional] [default to null]
**Media** | [**[]Media**](Media.md) | An array of extracted media such as images and videos | [optional] [default to null]
**Sentiment** | [**Sentiments**](Sentiments.md) | Suggested sentiments for the story title or body | [optional] [default to null]
**Language** | **string** | Language of the story | [optional] [default to null]
**PublishedAt** | [**time.Time**](time.Time.md) | Published date of the story | [optional] [default to null]
**Links** | [**StoryLinks**](StoryLinks.md) | Links which is related to the story | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


