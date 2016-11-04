# \DefaultApi

All URIs are relative to *https://api.newsapi.aylien.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAutocompletes**](DefaultApi.md#ListAutocompletes) | **Get** /autocompletes | List autocompletes
[**ListCoverages**](DefaultApi.md#ListCoverages) | **Post** /coverages | List coverages
[**ListHistograms**](DefaultApi.md#ListHistograms) | **Get** /histograms | List histograms
[**ListRelatedStories**](DefaultApi.md#ListRelatedStories) | **Post** /related_stories | List related stories
[**ListStories**](DefaultApi.md#ListStories) | **Get** /stories | List Stories
[**ListTimeSeries**](DefaultApi.md#ListTimeSeries) | **Get** /time_series | List time series
[**ListTrends**](DefaultApi.md#ListTrends) | **Get** /trends | List trends


# **ListAutocompletes**
> Autocompletes ListAutocompletes($type_, $term, $language, $perPage)

List autocompletes

This endpoint is used for getting list of autocompletes by providing a specific term and type.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Type** | **string**| This parameter is used for defining the type of autocompletes. | 
 **Term** | **string**| This parameter is used for finding autocomplete objects that contain the specified value. | 
 **Language** | **string**| This parameter is used for autocompletes whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes. | [optional] [default to en]
 **PerPage** | **int32**| This parameter is used for specifying number of items in each page. | [optional] [default to 3]

### Return type

[**Autocompletes**](Autocompletes.md)

### Authorization

[app_key](../README.md#app_key), [app_id](../README.md#app_id)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCoverages**
> Coverages ListCoverages($id, $title, $body, $text, $language, $publishedAtStart, $publishedAtEnd, $categoriesTaxonomy, $categoriesConfident, $categoriesId, $categoriesLevel, $entitiesTitleText, $entitiesTitleType, $entitiesTitleLinksDbpedia, $entitiesBodyText, $entitiesBodyType, $entitiesBodyLinksDbpedia, $sentimentTitlePolarity, $sentimentBodyPolarity, $mediaImagesCountMin, $mediaImagesCountMax, $mediaImagesWidthMin, $mediaImagesWidthMax, $mediaImagesHeightMin, $mediaImagesHeightMax, $mediaImagesContentLengthMin, $mediaImagesContentLengthMax, $mediaImagesFormat, $mediaVideosCountMin, $mediaVideosCountMax, $authorId, $authorName, $sourceId, $sourceName, $sourceDomain, $sourceLocationsCountry, $sourceLocationsState, $sourceLocationsCity, $sourceScopesCountry, $sourceScopesState, $sourceScopesCity, $sourceScopesLevel, $sourceLinksInCountMin, $sourceLinksInCountMax, $sourceRankingsAlexaRankMin, $sourceRankingsAlexaRankMax, $sourceRankingsAlexaCountry, $socialSharesCountFacebookMin, $socialSharesCountFacebookMax, $socialSharesCountGooglePlusMin, $socialSharesCountGooglePlusMax, $socialSharesCountLinkedinMin, $socialSharesCountLinkedinMax, $socialSharesCountRedditMin, $socialSharesCountRedditMax, $cluster, $clusterAlgorithm, $return_, $storyId, $storyUrl, $storyTitle, $storyBody, $storyPublishedAt, $storyLanguage, $perPage)

List coverages

This endpoint is used for finding story coverages based on the parameters provided. The maximum number of related stories returned is 100.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | [**[]int64**](int64.md)| This parameter is used for finding stories by story id. | [optional] 
 **Title** | **string**| This parameter is used for finding stories whose title contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Body** | **string**| This parameter is used for finding stories whose body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Text** | **string**| This parameter is used for finding stories whose title or body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Language** | [**[]string**](string.md)| This parameter is used for finding stories whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes. | [optional] 
 **PublishedAtStart** | **string**| This parameter is used for finding stories whose published at time is greater than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates). | [optional] 
 **PublishedAtEnd** | **string**| This parameter is used for finding stories whose published at time is less than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates). | [optional] 
 **CategoriesTaxonomy** | **string**| This parameter is used for defining the type of the taxonomy for the rest of the categories queries. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] 
 **CategoriesConfident** | **bool**| This parameter is used for finding stories whose categories are confident. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] [default to true]
 **CategoriesId** | [**[]string**](string.md)| This parameter is used for finding stories by categories id. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] 
 **CategoriesLevel** | [**[]int32**](int32.md)| This parameter is used for finding stories by categories level. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] 
 **EntitiesTitleText** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;text&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesTitleType** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;type&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesTitleLinksDbpedia** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities dbpedia URL in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesBodyText** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;text&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesBodyType** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;type&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesBodyLinksDbpedia** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities dbpedia URL in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **SentimentTitlePolarity** | **string**| This parameter is used for finding stories whose title sentiment is the specified value. | [optional] 
 **SentimentBodyPolarity** | **string**| This parameter is used for finding stories whose body sentiment is the specified value. | [optional] 
 **MediaImagesCountMin** | **string**| This parameter is used for finding stories whose number of images is greater than or equal to the specified value. | [optional] 
 **MediaImagesCountMax** | **string**| This parameter is used for finding stories whose number of images is less than or equal to the specified value. | [optional] 
 **MediaImagesWidthMin** | **string**| This parameter is used for finding stories whose width of images are greater than or equal to the specified value. | [optional] 
 **MediaImagesWidthMax** | **string**| This parameter is used for finding stories whose width of images are less than or equal to the specified value. | [optional] 
 **MediaImagesHeightMin** | **string**| This parameter is used for finding stories whose height of images are greater than or equal to the specified value. | [optional] 
 **MediaImagesHeightMax** | **string**| This parameter is used for finding stories whose height of images are less than or equal to the specified value. | [optional] 
 **MediaImagesContentLengthMin** | **string**| This parameter is used for finding stories whose images content length are greater than or equal to the specified value. | [optional] 
 **MediaImagesContentLengthMax** | **string**| This parameter is used for finding stories whose images content length are less than or equal to the specified value. | [optional] 
 **MediaImagesFormat** | [**[]string**](string.md)| This parameter is used for finding stories whose images format are the specified value. | [optional] 
 **MediaVideosCountMin** | **string**| This parameter is used for finding stories whose number of videos is greater than or equal to the specified value. | [optional] 
 **MediaVideosCountMax** | **string**| This parameter is used for finding stories whose number of videos is less than or equal to the specified value. | [optional] 
 **AuthorId** | [**[]int32**](int32.md)| This parameter is used for finding stories whose author id is the specified value. | [optional] 
 **AuthorName** | **string**| This parameter is used for finding stories whose author full name contains the specified value. | [optional] 
 **SourceId** | [**[]int32**](int32.md)| This parameter is used for finding stories whose source id is the specified value. | [optional] 
 **SourceName** | [**[]string**](string.md)| This parameter is used for finding stories whose source name contains the specified value. | [optional] 
 **SourceDomain** | [**[]string**](string.md)| This parameter is used for finding stories whose source domain is the specified value. | [optional] 
 **SourceLocationsCountry** | [**[]string**](string.md)| This parameter is used for finding stories whose source country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceLocationsState** | [**[]string**](string.md)| This parameter is used for finding stories whose source state/province is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceLocationsCity** | [**[]string**](string.md)| This parameter is used for finding stories whose source city is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceScopesCountry** | [**[]string**](string.md)| This parameter is used for finding stories whose source scopes  is the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceScopesState** | [**[]string**](string.md)| This parameter is used for finding stories whose source scopes is the specified state/province value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceScopesCity** | [**[]string**](string.md)| This parameter is used for finding stories whose source scopes is the specified city value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceScopesLevel** | [**[]string**](string.md)| This parameter is used for finding stories whose source scopes  is the specified level value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceLinksInCountMin** | **int32**| This parameter is used for finding stories from sources whose Links in count is greater than or equal to the specified value. You can read more about working with Links in count [here](https://newsapi.aylien.com/docs/working-with-links-in-count). | [optional] 
 **SourceLinksInCountMax** | **int32**| This parameter is used for finding stories from sources whose Links in count is less than or equal to the specified value. You can read more about working with Links in count [here](https://newsapi.aylien.com/docs/working-with-links-in-count). | [optional] 
 **SourceRankingsAlexaRankMin** | **int32**| This parameter is used for finding stories from sources whose Alexa rank is greater than or equal to the specified value. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
 **SourceRankingsAlexaRankMax** | **int32**| This parameter is used for finding stories from sources whose Alexa rank is less than or equal to the specified value. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
 **SourceRankingsAlexaCountry** | [**[]string**](string.md)| This parameter is used for finding stories from sources whose Alexa rank is in the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
 **SocialSharesCountFacebookMin** | **string**| This parameter is used for finding stories whose Facebook social shares count is greater than or equal to the specified value. | [optional] 
 **SocialSharesCountFacebookMax** | **string**| This parameter is used for finding stories whose Facebook social shares count is less than or equal to the specified value. | [optional] 
 **SocialSharesCountGooglePlusMin** | **string**| This parameter is used for finding stories whose Google+ social shares count is greater than or equal to the specified value. | [optional] 
 **SocialSharesCountGooglePlusMax** | **string**| This parameter is used for finding stories whose Google+ social shares count is less than or equal to the specified value. | [optional] 
 **SocialSharesCountLinkedinMin** | **string**| This parameter is used for finding stories whose LinkedIn social shares count is greater than or equal to the specified value. | [optional] 
 **SocialSharesCountLinkedinMax** | **string**| This parameter is used for finding stories whose LinkedIn social shares count is less than or equal to the specified value. | [optional] 
 **SocialSharesCountRedditMin** | **string**| This parameter is used for finding stories whose Reddit social shares count is greater than or equal to the specified value. | [optional] 
 **SocialSharesCountRedditMax** | **string**| This parameter is used for finding stories whose Reddit social shares count is less than or equal to the specified value. | [optional] 
 **Cluster** | **bool**| This parameter enables clustering for the returned stories. You can read more about working with clustering [here](https://newsapi.aylien.com/docs/working-with-clustering). | [optional] [default to false]
 **ClusterAlgorithm** | **string**| This parameter is used for specifying the clustering algorithm you wish to use. It supprts STC, Lingo and [k-means](https://en.wikipedia.org/wiki/K-means_clustering) algorithms. You can read more about working with clustering [here](https://newsapi.aylien.com/docs/working-with-clustering). | [optional] [default to lingo]
 **Return** | [**[]string**](string.md)| This parameter is used for specifying return fields. | [optional] 
 **StoryId** | **int64**| A story id | [optional] 
 **StoryUrl** | **string**| An article or webpage | [optional] 
 **StoryTitle** | **string**| Title of the article | [optional] 
 **StoryBody** | **string**| Body of the article | [optional] 
 **StoryPublishedAt** | **time.Time**| Publish date of the article. If you use a url or title and body of an article for getting coverages, this parameter is required. The format used is a restricted form of the canonical representation of dateTime in the [XML Schema specification (ISO 8601)](https://www.w3.org/TR/xmlschema-2/#dateTime). | [optional] 
 **StoryLanguage** | **string**| This parameter is used for setting the language of the story. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes. | [optional] [default to auto]
 **PerPage** | **int32**| This parameter is used for specifying number of items in each page. | [optional] [default to 3]

### Return type

[**Coverages**](Coverages.md)

### Authorization

[app_key](../README.md#app_key), [app_id](../README.md#app_id)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListHistograms**
> Histograms ListHistograms($id, $title, $body, $text, $language, $publishedAtStart, $publishedAtEnd, $categoriesTaxonomy, $categoriesConfident, $categoriesId, $categoriesLevel, $entitiesTitleText, $entitiesTitleType, $entitiesTitleLinksDbpedia, $entitiesBodyText, $entitiesBodyType, $entitiesBodyLinksDbpedia, $sentimentTitlePolarity, $sentimentBodyPolarity, $mediaImagesCountMin, $mediaImagesCountMax, $mediaImagesWidthMin, $mediaImagesWidthMax, $mediaImagesHeightMin, $mediaImagesHeightMax, $mediaImagesContentLengthMin, $mediaImagesContentLengthMax, $mediaImagesFormat, $mediaVideosCountMin, $mediaVideosCountMax, $authorId, $authorName, $sourceId, $sourceName, $sourceDomain, $sourceLocationsCountry, $sourceLocationsState, $sourceLocationsCity, $sourceScopesCountry, $sourceScopesState, $sourceScopesCity, $sourceScopesLevel, $sourceLinksInCountMin, $sourceLinksInCountMax, $sourceRankingsAlexaRankMin, $sourceRankingsAlexaRankMax, $sourceRankingsAlexaCountry, $socialSharesCountFacebookMin, $socialSharesCountFacebookMax, $socialSharesCountGooglePlusMin, $socialSharesCountGooglePlusMax, $socialSharesCountLinkedinMin, $socialSharesCountLinkedinMax, $socialSharesCountRedditMin, $socialSharesCountRedditMax, $intervalStart, $intervalEnd, $intervalWidth, $field)

List histograms

This endpoint is used for getting histograms based on the `field` parameter passed to the API.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | [**[]int64**](int64.md)| This parameter is used for finding stories by story id. | [optional] 
 **Title** | **string**| This parameter is used for finding stories whose title contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Body** | **string**| This parameter is used for finding stories whose body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Text** | **string**| This parameter is used for finding stories whose title or body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Language** | [**[]string**](string.md)| This parameter is used for finding stories whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes. | [optional] 
 **PublishedAtStart** | **string**| This parameter is used for finding stories whose published at time is greater than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates). | [optional] 
 **PublishedAtEnd** | **string**| This parameter is used for finding stories whose published at time is less than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates). | [optional] 
 **CategoriesTaxonomy** | **string**| This parameter is used for defining the type of the taxonomy for the rest of the categories queries. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] 
 **CategoriesConfident** | **bool**| This parameter is used for finding stories whose categories are confident. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] [default to true]
 **CategoriesId** | [**[]string**](string.md)| This parameter is used for finding stories by categories id. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] 
 **CategoriesLevel** | [**[]int32**](int32.md)| This parameter is used for finding stories by categories level. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] 
 **EntitiesTitleText** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;text&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesTitleType** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;type&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesTitleLinksDbpedia** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities dbpedia URL in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesBodyText** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;text&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesBodyType** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;type&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesBodyLinksDbpedia** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities dbpedia URL in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **SentimentTitlePolarity** | **string**| This parameter is used for finding stories whose title sentiment is the specified value. | [optional] 
 **SentimentBodyPolarity** | **string**| This parameter is used for finding stories whose body sentiment is the specified value. | [optional] 
 **MediaImagesCountMin** | **string**| This parameter is used for finding stories whose number of images is greater than or equal to the specified value. | [optional] 
 **MediaImagesCountMax** | **string**| This parameter is used for finding stories whose number of images is less than or equal to the specified value. | [optional] 
 **MediaImagesWidthMin** | **string**| This parameter is used for finding stories whose width of images are greater than or equal to the specified value. | [optional] 
 **MediaImagesWidthMax** | **string**| This parameter is used for finding stories whose width of images are less than or equal to the specified value. | [optional] 
 **MediaImagesHeightMin** | **string**| This parameter is used for finding stories whose height of images are greater than or equal to the specified value. | [optional] 
 **MediaImagesHeightMax** | **string**| This parameter is used for finding stories whose height of images are less than or equal to the specified value. | [optional] 
 **MediaImagesContentLengthMin** | **string**| This parameter is used for finding stories whose images content length are greater than or equal to the specified value. | [optional] 
 **MediaImagesContentLengthMax** | **string**| This parameter is used for finding stories whose images content length are less than or equal to the specified value. | [optional] 
 **MediaImagesFormat** | [**[]string**](string.md)| This parameter is used for finding stories whose images format are the specified value. | [optional] 
 **MediaVideosCountMin** | **string**| This parameter is used for finding stories whose number of videos is greater than or equal to the specified value. | [optional] 
 **MediaVideosCountMax** | **string**| This parameter is used for finding stories whose number of videos is less than or equal to the specified value. | [optional] 
 **AuthorId** | [**[]int32**](int32.md)| This parameter is used for finding stories whose author id is the specified value. | [optional] 
 **AuthorName** | **string**| This parameter is used for finding stories whose author full name contains the specified value. | [optional] 
 **SourceId** | [**[]int32**](int32.md)| This parameter is used for finding stories whose source id is the specified value. | [optional] 
 **SourceName** | [**[]string**](string.md)| This parameter is used for finding stories whose source name contains the specified value. | [optional] 
 **SourceDomain** | [**[]string**](string.md)| This parameter is used for finding stories whose source domain is the specified value. | [optional] 
 **SourceLocationsCountry** | [**[]string**](string.md)| This parameter is used for finding stories whose source country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceLocationsState** | [**[]string**](string.md)| This parameter is used for finding stories whose source state/province is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceLocationsCity** | [**[]string**](string.md)| This parameter is used for finding stories whose source city is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceScopesCountry** | [**[]string**](string.md)| This parameter is used for finding stories whose source scopes is the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceScopesState** | [**[]string**](string.md)| This parameter is used for finding stories whose source scopes is the specified state/province value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceScopesCity** | [**[]string**](string.md)| This parameter is used for finding stories whose source scopes is the specified city value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceScopesLevel** | [**[]string**](string.md)| This parameter is used for finding stories whose source scopes is the specified level value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceLinksInCountMin** | **int32**| This parameter is used for finding stories from sources whose Links in count is greater than or equal to the specified value. You can read more about working with Links in count [here](https://newsapi.aylien.com/docs/working-with-links-in-count). | [optional] 
 **SourceLinksInCountMax** | **int32**| This parameter is used for finding stories from sources whose Links in count is less than or equal to the specified value. You can read more about working with Links in count [here](https://newsapi.aylien.com/docs/working-with-links-in-count). | [optional] 
 **SourceRankingsAlexaRankMin** | **int32**| This parameter is used for finding stories from sources whose Alexa rank is greater than or equal to the specified value. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
 **SourceRankingsAlexaRankMax** | **int32**| This parameter is used for finding stories from sources whose Alexa rank is less than or equal to the specified value. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
 **SourceRankingsAlexaCountry** | [**[]string**](string.md)| This parameter is used for finding stories from sources whose Alexa rank is in the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
 **SocialSharesCountFacebookMin** | **string**| This parameter is used for finding stories whose Facebook social shares count is greater than or equal to the specified value. | [optional] 
 **SocialSharesCountFacebookMax** | **string**| This parameter is used for finding stories whose Facebook social shares count is less than or equal to the specified value. | [optional] 
 **SocialSharesCountGooglePlusMin** | **string**| This parameter is used for finding stories whose Google+ social shares count is greater than or equal to the specified value. | [optional] 
 **SocialSharesCountGooglePlusMax** | **string**| This parameter is used for finding stories whose Google+ social shares count is less than or equal to the specified value. | [optional] 
 **SocialSharesCountLinkedinMin** | **string**| This parameter is used for finding stories whose LinkedIn social shares count is greater than or equal to the specified value. | [optional] 
 **SocialSharesCountLinkedinMax** | **string**| This parameter is used for finding stories whose LinkedIn social shares count is less than or equal to the specified value. | [optional] 
 **SocialSharesCountRedditMin** | **string**| This parameter is used for finding stories whose Reddit social shares count is greater than or equal to the specified value. | [optional] 
 **SocialSharesCountRedditMax** | **string**| This parameter is used for finding stories whose Reddit social shares count is less than or equal to the specified value. | [optional] 
 **IntervalStart** | **int32**| This parameter is used for setting the start data point of histogram intervals. | [optional] 
 **IntervalEnd** | **int32**| This parameter is used for setting the end data point of histogram intervals. | [optional] 
 **IntervalWidth** | **int32**| This parameter is used for setting the width of histogram intervals. | [optional] 
 **Field** | **string**| This parameter is used for specifying the y-axis variable for the histogram. | [optional] [default to social_shares_count]

### Return type

[**Histograms**](Histograms.md)

### Authorization

[app_key](../README.md#app_key), [app_id](../README.md#app_id)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRelatedStories**
> RelatedStories ListRelatedStories($id, $title, $body, $text, $language, $publishedAtStart, $publishedAtEnd, $categoriesTaxonomy, $categoriesConfident, $categoriesId, $categoriesLevel, $entitiesTitleText, $entitiesTitleType, $entitiesTitleLinksDbpedia, $entitiesBodyText, $entitiesBodyType, $entitiesBodyLinksDbpedia, $sentimentTitlePolarity, $sentimentBodyPolarity, $mediaImagesCountMin, $mediaImagesCountMax, $mediaImagesWidthMin, $mediaImagesWidthMax, $mediaImagesHeightMin, $mediaImagesHeightMax, $mediaImagesContentLengthMin, $mediaImagesContentLengthMax, $mediaImagesFormat, $mediaVideosCountMin, $mediaVideosCountMax, $authorId, $authorName, $sourceId, $sourceName, $sourceDomain, $sourceLocationsCountry, $sourceLocationsState, $sourceLocationsCity, $sourceScopesCountry, $sourceScopesState, $sourceScopesCity, $sourceScopesLevel, $sourceLinksInCountMin, $sourceLinksInCountMax, $sourceRankingsAlexaRankMin, $sourceRankingsAlexaRankMax, $sourceRankingsAlexaCountry, $socialSharesCountFacebookMin, $socialSharesCountFacebookMax, $socialSharesCountGooglePlusMin, $socialSharesCountGooglePlusMax, $socialSharesCountLinkedinMin, $socialSharesCountLinkedinMax, $socialSharesCountRedditMin, $socialSharesCountRedditMax, $cluster, $clusterAlgorithm, $return_, $storyId, $storyUrl, $storyTitle, $storyBody, $boostBy, $storyLanguage, $perPage)

List related stories

This endpoint is used for finding related stories based on the parameters provided. The maximum number of related stories returned is 100.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | [**[]int64**](int64.md)| This parameter is used for finding stories by story id. | [optional] 
 **Title** | **string**| This parameter is used for finding stories whose title contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Body** | **string**| This parameter is used for finding stories whose body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Text** | **string**| This parameter is used for finding stories whose title or body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Language** | [**[]string**](string.md)| This parameter is used for finding stories whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes. | [optional] 
 **PublishedAtStart** | **string**| This parameter is used for finding stories whose published at time is greater than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates). | [optional] 
 **PublishedAtEnd** | **string**| This parameter is used for finding stories whose published at time is less than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates). | [optional] 
 **CategoriesTaxonomy** | **string**| This parameter is used for defining the type of the taxonomy for the rest of the categories queries. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] 
 **CategoriesConfident** | **bool**| This parameter is used for finding stories whose categories are confident. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] [default to true]
 **CategoriesId** | [**[]string**](string.md)| This parameter is used for finding stories by categories id. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] 
 **CategoriesLevel** | [**[]int32**](int32.md)| This parameter is used for finding stories by categories level. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] 
 **EntitiesTitleText** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;text&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesTitleType** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;type&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesTitleLinksDbpedia** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities dbpedia URL in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesBodyText** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;text&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesBodyType** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;type&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesBodyLinksDbpedia** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities dbpedia URL in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **SentimentTitlePolarity** | **string**| This parameter is used for finding stories whose title sentiment is the specified value. | [optional] 
 **SentimentBodyPolarity** | **string**| This parameter is used for finding stories whose body sentiment is the specified value. | [optional] 
 **MediaImagesCountMin** | **string**| This parameter is used for finding stories whose number of images is greater than or equal to the specified value. | [optional] 
 **MediaImagesCountMax** | **string**| This parameter is used for finding stories whose number of images is less than or equal to the specified value. | [optional] 
 **MediaImagesWidthMin** | **string**| This parameter is used for finding stories whose width of images are greater than or equal to the specified value. | [optional] 
 **MediaImagesWidthMax** | **string**| This parameter is used for finding stories whose width of images are less than or equal to the specified value. | [optional] 
 **MediaImagesHeightMin** | **string**| This parameter is used for finding stories whose height of images are greater than or equal to the specified value. | [optional] 
 **MediaImagesHeightMax** | **string**| This parameter is used for finding stories whose height of images are less than or equal to the specified value. | [optional] 
 **MediaImagesContentLengthMin** | **string**| This parameter is used for finding stories whose images content length are greater than or equal to the specified value. | [optional] 
 **MediaImagesContentLengthMax** | **string**| This parameter is used for finding stories whose images content length are less than or equal to the specified value. | [optional] 
 **MediaImagesFormat** | [**[]string**](string.md)| This parameter is used for finding stories whose images format are the specified value. | [optional] 
 **MediaVideosCountMin** | **string**| This parameter is used for finding stories whose number of videos is greater than or equal to the specified value. | [optional] 
 **MediaVideosCountMax** | **string**| This parameter is used for finding stories whose number of videos is less than or equal to the specified value. | [optional] 
 **AuthorId** | [**[]int32**](int32.md)| This parameter is used for finding stories whose author id is the specified value. | [optional] 
 **AuthorName** | **string**| This parameter is used for finding stories whose author full name contains the specified value. | [optional] 
 **SourceId** | [**[]int32**](int32.md)| This parameter is used for finding stories whose source id is the specified value. | [optional] 
 **SourceName** | [**[]string**](string.md)| This parameter is used for finding stories whose source name contains the specified value. | [optional] 
 **SourceDomain** | [**[]string**](string.md)| This parameter is used for finding stories whose source domain is the specified value. | [optional] 
 **SourceLocationsCountry** | [**[]string**](string.md)| This parameter is used for finding stories whose source country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceLocationsState** | [**[]string**](string.md)| This parameter is used for finding stories whose source state/province is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceLocationsCity** | [**[]string**](string.md)| This parameter is used for finding stories whose source city is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceScopesCountry** | [**[]string**](string.md)| This parameter is used for finding stories whose source scopes  is the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceScopesState** | [**[]string**](string.md)| This parameter is used for finding stories whose source scopes is the specified state/province value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceScopesCity** | [**[]string**](string.md)| This parameter is used for finding stories whose source scopes is the specified city value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceScopesLevel** | [**[]string**](string.md)| This parameter is used for finding stories whose source scopes  is the specified level value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceLinksInCountMin** | **int32**| This parameter is used for finding stories from sources whose Links in count is greater than or equal to the specified value. You can read more about working with Links in count [here](https://newsapi.aylien.com/docs/working-with-links-in-count). | [optional] 
 **SourceLinksInCountMax** | **int32**| This parameter is used for finding stories from sources whose Links in count is less than or equal to the specified value. You can read more about working with Links in count [here](https://newsapi.aylien.com/docs/working-with-links-in-count). | [optional] 
 **SourceRankingsAlexaRankMin** | **int32**| This parameter is used for finding stories from sources whose Alexa rank is greater than or equal to the specified value. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
 **SourceRankingsAlexaRankMax** | **int32**| This parameter is used for finding stories from sources whose Alexa rank is less than or equal to the specified value. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
 **SourceRankingsAlexaCountry** | [**[]string**](string.md)| This parameter is used for finding stories from sources whose Alexa rank is in the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
 **SocialSharesCountFacebookMin** | **string**| This parameter is used for finding stories whose Facebook social shares count is greater than or equal to the specified value. | [optional] 
 **SocialSharesCountFacebookMax** | **string**| This parameter is used for finding stories whose Facebook social shares count is less than or equal to the specified value. | [optional] 
 **SocialSharesCountGooglePlusMin** | **string**| This parameter is used for finding stories whose Google+ social shares count is greater than or equal to the specified value. | [optional] 
 **SocialSharesCountGooglePlusMax** | **string**| This parameter is used for finding stories whose Google+ social shares count is less than or equal to the specified value. | [optional] 
 **SocialSharesCountLinkedinMin** | **string**| This parameter is used for finding stories whose LinkedIn social shares count is greater than or equal to the specified value. | [optional] 
 **SocialSharesCountLinkedinMax** | **string**| This parameter is used for finding stories whose LinkedIn social shares count is less than or equal to the specified value. | [optional] 
 **SocialSharesCountRedditMin** | **string**| This parameter is used for finding stories whose Reddit social shares count is greater than or equal to the specified value. | [optional] 
 **SocialSharesCountRedditMax** | **string**| This parameter is used for finding stories whose Reddit social shares count is less than or equal to the specified value. | [optional] 
 **Cluster** | **bool**| This parameter enables clustering for the returned stories. You can read more about working with clustering [here](https://newsapi.aylien.com/docs/working-with-clustering). | [optional] [default to false]
 **ClusterAlgorithm** | **string**| This parameter is used for specifying the clustering algorithm you wish to use. It supprts STC, Lingo and [k-means](https://en.wikipedia.org/wiki/K-means_clustering) algorithms. You can read more about working with clustering [here](https://newsapi.aylien.com/docs/working-with-clustering). | [optional] [default to lingo]
 **Return** | [**[]string**](string.md)| This parameter is used for specifying return fields. | [optional] 
 **StoryId** | **int64**| A story id | [optional] 
 **StoryUrl** | **string**| An article or webpage | [optional] 
 **StoryTitle** | **string**| Title of the article | [optional] 
 **StoryBody** | **string**| Body of the article | [optional] 
 **BoostBy** | **string**| This parameter is used for boosting the result by the specified value. | [optional] 
 **StoryLanguage** | **string**| This parameter is used for setting the language of the story. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes. | [optional] [default to auto]
 **PerPage** | **int32**| This parameter is used for specifying number of items in each page. | [optional] [default to 3]

### Return type

[**RelatedStories**](RelatedStories.md)

### Authorization

[app_key](../README.md#app_key), [app_id](../README.md#app_id)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListStories**
> Stories ListStories($id, $title, $body, $text, $language, $publishedAtStart, $publishedAtEnd, $categoriesTaxonomy, $categoriesConfident, $categoriesId, $categoriesLevel, $entitiesTitleText, $entitiesTitleType, $entitiesTitleLinksDbpedia, $entitiesBodyText, $entitiesBodyType, $entitiesBodyLinksDbpedia, $sentimentTitlePolarity, $sentimentBodyPolarity, $mediaImagesCountMin, $mediaImagesCountMax, $mediaImagesWidthMin, $mediaImagesWidthMax, $mediaImagesHeightMin, $mediaImagesHeightMax, $mediaImagesContentLengthMin, $mediaImagesContentLengthMax, $mediaImagesFormat, $mediaVideosCountMin, $mediaVideosCountMax, $authorId, $authorName, $sourceId, $sourceName, $sourceDomain, $sourceLocationsCountry, $sourceLocationsState, $sourceLocationsCity, $sourceScopesCountry, $sourceScopesState, $sourceScopesCity, $sourceScopesLevel, $sourceLinksInCountMin, $sourceLinksInCountMax, $sourceRankingsAlexaRankMin, $sourceRankingsAlexaRankMax, $sourceRankingsAlexaCountry, $socialSharesCountFacebookMin, $socialSharesCountFacebookMax, $socialSharesCountGooglePlusMin, $socialSharesCountGooglePlusMax, $socialSharesCountLinkedinMin, $socialSharesCountLinkedinMax, $socialSharesCountRedditMin, $socialSharesCountRedditMax, $cluster, $clusterAlgorithm, $return_, $sortBy, $sortDirection, $cursor, $perPage)

List Stories

This endpoint is used for getting a list of stories.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | [**[]int64**](int64.md)| This parameter is used for finding stories by story id. | [optional] 
 **Title** | **string**| This parameter is used for finding stories whose title contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Body** | **string**| This parameter is used for finding stories whose body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Text** | **string**| This parameter is used for finding stories whose title or body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Language** | [**[]string**](string.md)| This parameter is used for finding stories whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes. | [optional] 
 **PublishedAtStart** | **string**| This parameter is used for finding stories whose published at time is greater than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates). | [optional] 
 **PublishedAtEnd** | **string**| This parameter is used for finding stories whose published at time is less than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates). | [optional] 
 **CategoriesTaxonomy** | **string**| This parameter is used for defining the type of the taxonomy for the rest of the categories queries. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] 
 **CategoriesConfident** | **bool**| This parameter is used for finding stories whose categories are confident. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] [default to true]
 **CategoriesId** | [**[]string**](string.md)| This parameter is used for finding stories by categories id. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] 
 **CategoriesLevel** | [**[]int32**](int32.md)| This parameter is used for finding stories by categories level. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] 
 **EntitiesTitleText** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;text&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesTitleType** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;type&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesTitleLinksDbpedia** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities dbpedia URL in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesBodyText** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;text&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesBodyType** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;type&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesBodyLinksDbpedia** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities dbpedia URL in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **SentimentTitlePolarity** | **string**| This parameter is used for finding stories whose title sentiment is the specified value. | [optional] 
 **SentimentBodyPolarity** | **string**| This parameter is used for finding stories whose body sentiment is the specified value. | [optional] 
 **MediaImagesCountMin** | **string**| This parameter is used for finding stories whose number of images is greater than or equal to the specified value. | [optional] 
 **MediaImagesCountMax** | **string**| This parameter is used for finding stories whose number of images is less than or equal to the specified value. | [optional] 
 **MediaImagesWidthMin** | **string**| This parameter is used for finding stories whose width of images are greater than or equal to the specified value. | [optional] 
 **MediaImagesWidthMax** | **string**| This parameter is used for finding stories whose width of images are less than or equal to the specified value. | [optional] 
 **MediaImagesHeightMin** | **string**| This parameter is used for finding stories whose height of images are greater than or equal to the specified value. | [optional] 
 **MediaImagesHeightMax** | **string**| This parameter is used for finding stories whose height of images are less than or equal to the specified value. | [optional] 
 **MediaImagesContentLengthMin** | **string**| This parameter is used for finding stories whose images content length are greater than or equal to the specified value. | [optional] 
 **MediaImagesContentLengthMax** | **string**| This parameter is used for finding stories whose images content length are less than or equal to the specified value. | [optional] 
 **MediaImagesFormat** | [**[]string**](string.md)| This parameter is used for finding stories whose images format are the specified value. | [optional] 
 **MediaVideosCountMin** | **string**| This parameter is used for finding stories whose number of videos is greater than or equal to the specified value. | [optional] 
 **MediaVideosCountMax** | **string**| This parameter is used for finding stories whose number of videos is less than or equal to the specified value. | [optional] 
 **AuthorId** | [**[]int32**](int32.md)| This parameter is used for finding stories whose author id is the specified value. | [optional] 
 **AuthorName** | **string**| This parameter is used for finding stories whose author full name contains the specified value. | [optional] 
 **SourceId** | [**[]int32**](int32.md)| This parameter is used for finding stories whose source id is the specified value. | [optional] 
 **SourceName** | [**[]string**](string.md)| This parameter is used for finding stories whose source name contains the specified value. | [optional] 
 **SourceDomain** | [**[]string**](string.md)| This parameter is used for finding stories whose source domain is the specified value. | [optional] 
 **SourceLocationsCountry** | [**[]string**](string.md)| This parameter is used for finding stories whose source country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceLocationsState** | [**[]string**](string.md)| This parameter is used for finding stories whose source state/province is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceLocationsCity** | [**[]string**](string.md)| This parameter is used for finding stories whose source city is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceScopesCountry** | [**[]string**](string.md)| This parameter is used for finding stories whose source scopes is the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceScopesState** | [**[]string**](string.md)| This parameter is used for finding stories whose source scopes is the specified state/province value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceScopesCity** | [**[]string**](string.md)| This parameter is used for finding stories whose source scopes is the specified city value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceScopesLevel** | [**[]string**](string.md)| This parameter is used for finding stories whose source scopes is the specified level value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceLinksInCountMin** | **int32**| This parameter is used for finding stories from sources whose Links in count is greater than or equal to the specified value. You can read more about working with Links in count [here](https://newsapi.aylien.com/docs/working-with-links-in-count). | [optional] 
 **SourceLinksInCountMax** | **int32**| This parameter is used for finding stories from sources whose Links in count is less than or equal to the specified value. You can read more about working with Links in count [here](https://newsapi.aylien.com/docs/working-with-links-in-count). | [optional] 
 **SourceRankingsAlexaRankMin** | **int32**| This parameter is used for finding stories from sources whose Alexa rank is greater than or equal to the specified value. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
 **SourceRankingsAlexaRankMax** | **int32**| This parameter is used for finding stories from sources whose Alexa rank is less than or equal to the specified value. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
 **SourceRankingsAlexaCountry** | [**[]string**](string.md)| This parameter is used for finding stories from sources whose Alexa rank is in the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
 **SocialSharesCountFacebookMin** | **string**| This parameter is used for finding stories whose Facebook social shares count is greater than or equal to the specified value. | [optional] 
 **SocialSharesCountFacebookMax** | **string**| This parameter is used for finding stories whose Facebook social shares count is less than or equal to the specified value. | [optional] 
 **SocialSharesCountGooglePlusMin** | **string**| This parameter is used for finding stories whose Google+ social shares count is greater than or equal to the specified value. | [optional] 
 **SocialSharesCountGooglePlusMax** | **string**| This parameter is used for finding stories whose Google+ social shares count is less than or equal to the specified value. | [optional] 
 **SocialSharesCountLinkedinMin** | **string**| This parameter is used for finding stories whose LinkedIn social shares count is greater than or equal to the specified value. | [optional] 
 **SocialSharesCountLinkedinMax** | **string**| This parameter is used for finding stories whose LinkedIn social shares count is less than or equal to the specified value. | [optional] 
 **SocialSharesCountRedditMin** | **string**| This parameter is used for finding stories whose Reddit social shares count is greater than or equal to the specified value. | [optional] 
 **SocialSharesCountRedditMax** | **string**| This parameter is used for finding stories whose Reddit social shares count is less than or equal to the specified value. | [optional] 
 **Cluster** | **bool**| This parameter enables clustering for the returned stories. You can read more about working with clustering [here](https://newsapi.aylien.com/docs/working-with-clustering). | [optional] [default to false]
 **ClusterAlgorithm** | **string**| This parameter is used for specifying the clustering algorithm you wish to use. It supprts STC, Lingo and [k-means](https://en.wikipedia.org/wiki/K-means_clustering) algorithms. You can read more about working with clustering [here](https://newsapi.aylien.com/docs/working-with-clustering). | [optional] [default to lingo]
 **Return** | [**[]string**](string.md)| This parameter is used for specifying return fields. | [optional] 
 **SortBy** | **string**| This parameter is used for changing the order column of the results. | [optional] [default to published_at]
 **SortDirection** | **string**| This parameter is used for changing the order direction of the result. | [optional] [default to desc]
 **Cursor** | **string**| This parameter is used for finding a specific page. You can read more about pagination of results [here](https://newsapi.aylien.com/docs/pagination-of-results). | [optional] [default to *]
 **PerPage** | **int32**| This parameter is used for specifying number of items in each page You can read more about pagination of results [here](https://newsapi.aylien.com/docs/pagination-of-results) | [optional] [default to 10]

### Return type

[**Stories**](Stories.md)

### Authorization

[app_key](../README.md#app_key), [app_id](../README.md#app_id)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTimeSeries**
> TimeSeriesList ListTimeSeries($id, $title, $body, $text, $language, $categoriesTaxonomy, $categoriesConfident, $categoriesId, $categoriesLevel, $entitiesTitleText, $entitiesTitleType, $entitiesTitleLinksDbpedia, $entitiesBodyText, $entitiesBodyType, $entitiesBodyLinksDbpedia, $sentimentTitlePolarity, $sentimentBodyPolarity, $mediaImagesCountMin, $mediaImagesCountMax, $mediaImagesWidthMin, $mediaImagesWidthMax, $mediaImagesHeightMin, $mediaImagesHeightMax, $mediaImagesContentLengthMin, $mediaImagesContentLengthMax, $mediaImagesFormat, $mediaVideosCountMin, $mediaVideosCountMax, $authorId, $authorName, $sourceId, $sourceName, $sourceDomain, $sourceLocationsCountry, $sourceLocationsState, $sourceLocationsCity, $sourceScopesCountry, $sourceScopesState, $sourceScopesCity, $sourceScopesLevel, $sourceLinksInCountMin, $sourceLinksInCountMax, $sourceRankingsAlexaRankMin, $sourceRankingsAlexaRankMax, $sourceRankingsAlexaCountry, $socialSharesCountFacebookMin, $socialSharesCountFacebookMax, $socialSharesCountGooglePlusMin, $socialSharesCountGooglePlusMax, $socialSharesCountLinkedinMin, $socialSharesCountLinkedinMax, $socialSharesCountRedditMin, $socialSharesCountRedditMax, $publishedAtStart, $publishedAtEnd, $period)

List time series

This endpoint is used for getting time series by stories.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | [**[]int64**](int64.md)| This parameter is used for finding stories by story id. | [optional] 
 **Title** | **string**| This parameter is used for finding stories whose title contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Body** | **string**| This parameter is used for finding stories whose body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Text** | **string**| This parameter is used for finding stories whose title or body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Language** | [**[]string**](string.md)| This parameter is used for finding stories whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes. | [optional] 
 **CategoriesTaxonomy** | **string**| This parameter is used for defining the type of the taxonomy for the rest of the categories queries. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] 
 **CategoriesConfident** | **bool**| This parameter is used for finding stories whose categories are confident. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] [default to true]
 **CategoriesId** | [**[]string**](string.md)| This parameter is used for finding stories by categories id. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] 
 **CategoriesLevel** | [**[]int32**](int32.md)| This parameter is used for finding stories by categories level. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] 
 **EntitiesTitleText** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;text&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesTitleType** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;type&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesTitleLinksDbpedia** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities dbpedia URL in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesBodyText** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;text&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesBodyType** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;type&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesBodyLinksDbpedia** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities dbpedia URL in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **SentimentTitlePolarity** | **string**| This parameter is used for finding stories whose title sentiment is the specified value. | [optional] 
 **SentimentBodyPolarity** | **string**| This parameter is used for finding stories whose body sentiment is the specified value. | [optional] 
 **MediaImagesCountMin** | **string**| This parameter is used for finding stories whose number of images is greater than or equal to the specified value. | [optional] 
 **MediaImagesCountMax** | **string**| This parameter is used for finding stories whose number of images is less than or equal to the specified value. | [optional] 
 **MediaImagesWidthMin** | **string**| This parameter is used for finding stories whose width of images are greater than or equal to the specified value. | [optional] 
 **MediaImagesWidthMax** | **string**| This parameter is used for finding stories whose width of images are less than or equal to the specified value. | [optional] 
 **MediaImagesHeightMin** | **string**| This parameter is used for finding stories whose height of images are greater than or equal to the specified value. | [optional] 
 **MediaImagesHeightMax** | **string**| This parameter is used for finding stories whose height of images are less than or equal to the specified value. | [optional] 
 **MediaImagesContentLengthMin** | **string**| This parameter is used for finding stories whose images content length are greater than or equal to the specified value. | [optional] 
 **MediaImagesContentLengthMax** | **string**| This parameter is used for finding stories whose images content length are less than or equal to the specified value. | [optional] 
 **MediaImagesFormat** | [**[]string**](string.md)| This parameter is used for finding stories whose images format are the specified value. | [optional] 
 **MediaVideosCountMin** | **string**| This parameter is used for finding stories whose number of videos is greater than or equal to the specified value. | [optional] 
 **MediaVideosCountMax** | **string**| This parameter is used for finding stories whose number of videos is less than or equal to the specified value. | [optional] 
 **AuthorId** | [**[]int32**](int32.md)| This parameter is used for finding stories whose author id is the specified value. | [optional] 
 **AuthorName** | **string**| This parameter is used for finding stories whose author full name contains the specified value. | [optional] 
 **SourceId** | [**[]int32**](int32.md)| This parameter is used for finding stories whose source id is the specified value. | [optional] 
 **SourceName** | [**[]string**](string.md)| This parameter is used for finding stories whose source name contains the specified value. | [optional] 
 **SourceDomain** | [**[]string**](string.md)| This parameter is used for finding stories whose source domain is the specified value. | [optional] 
 **SourceLocationsCountry** | [**[]string**](string.md)| This parameter is used for finding stories whose source country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceLocationsState** | [**[]string**](string.md)| This parameter is used for finding stories whose source state/province is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceLocationsCity** | [**[]string**](string.md)| This parameter is used for finding stories whose source city is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceScopesCountry** | [**[]string**](string.md)| This parameter is used for finding stories whose source scopes is the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceScopesState** | [**[]string**](string.md)| This parameter is used for finding stories whose source scopes is the specified state/province value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceScopesCity** | [**[]string**](string.md)| This parameter is used for finding stories whose source scopes is the specified city value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceScopesLevel** | [**[]string**](string.md)| This parameter is used for finding stories whose source scopes is the specified level value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceLinksInCountMin** | **int32**| This parameter is used for finding stories from sources whose Links in count is greater than or equal to the specified value. You can read more about working with Links in count [here](https://newsapi.aylien.com/docs/working-with-links-in-count). | [optional] 
 **SourceLinksInCountMax** | **int32**| This parameter is used for finding stories from sources whose Links in count is less than or equal to the specified value. You can read more about working with Links in count [here](https://newsapi.aylien.com/docs/working-with-links-in-count). | [optional] 
 **SourceRankingsAlexaRankMin** | **int32**| This parameter is used for finding stories from sources whose Alexa rank is greater than or equal to the specified value. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
 **SourceRankingsAlexaRankMax** | **int32**| This parameter is used for finding stories from sources whose Alexa rank is less than or equal to the specified value. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
 **SourceRankingsAlexaCountry** | [**[]string**](string.md)| This parameter is used for finding stories from sources whose Alexa rank is in the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
 **SocialSharesCountFacebookMin** | **string**| This parameter is used for finding stories whose Facebook social shares count is greater than or equal to the specified value. | [optional] 
 **SocialSharesCountFacebookMax** | **string**| This parameter is used for finding stories whose Facebook social shares count is less than or equal to the specified value. | [optional] 
 **SocialSharesCountGooglePlusMin** | **string**| This parameter is used for finding stories whose Google+ social shares count is greater than or equal to the specified value. | [optional] 
 **SocialSharesCountGooglePlusMax** | **string**| This parameter is used for finding stories whose Google+ social shares count is less than or equal to the specified value. | [optional] 
 **SocialSharesCountLinkedinMin** | **string**| This parameter is used for finding stories whose LinkedIn social shares count is greater than or equal to the specified value. | [optional] 
 **SocialSharesCountLinkedinMax** | **string**| This parameter is used for finding stories whose LinkedIn social shares count is less than or equal to the specified value. | [optional] 
 **SocialSharesCountRedditMin** | **string**| This parameter is used for finding stories whose Reddit social shares count is greater than or equal to the specified value. | [optional] 
 **SocialSharesCountRedditMax** | **string**| This parameter is used for finding stories whose Reddit social shares count is less than or equal to the specified value. | [optional] 
 **PublishedAtStart** | **string**| This parameter is used for finding stories whose published at time is less than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates). | [optional] [default to NOW-7DAYS/DAY]
 **PublishedAtEnd** | **string**| This parameter is used for finding stories whose published at time is greater than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates). | [optional] [default to NOW/DAY]
 **Period** | **string**| The size of each date range is expressed as an interval to be added to the lower bound. It supports Date Math Syntax. Valid options are &#x60;+&#x60; following an integer number greater than 0 and one of the Date Math keywords. e.g. &#x60;+1DAY&#x60;, &#x60;+2MINUTES&#x60; and &#x60;+1MONTH&#x60;. Here are [Supported keywords](https://newsapi.aylien.com/docs/working-with-dates#date-math). | [optional] [default to +1DAY]

### Return type

[**TimeSeriesList**](TimeSeriesList.md)

### Authorization

[app_key](../README.md#app_key), [app_id](../README.md#app_id)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTrends**
> Trends ListTrends($id, $title, $body, $text, $language, $publishedAtStart, $publishedAtEnd, $categoriesTaxonomy, $categoriesConfident, $categoriesId, $categoriesLevel, $entitiesTitleText, $entitiesTitleType, $entitiesTitleLinksDbpedia, $entitiesBodyText, $entitiesBodyType, $entitiesBodyLinksDbpedia, $sentimentTitlePolarity, $sentimentBodyPolarity, $mediaImagesCountMin, $mediaImagesCountMax, $mediaImagesWidthMin, $mediaImagesWidthMax, $mediaImagesHeightMin, $mediaImagesHeightMax, $mediaImagesContentLengthMin, $mediaImagesContentLengthMax, $mediaImagesFormat, $mediaVideosCountMin, $mediaVideosCountMax, $authorId, $authorName, $sourceId, $sourceName, $sourceDomain, $sourceLocationsCountry, $sourceLocationsState, $sourceLocationsCity, $sourceScopesCountry, $sourceScopesState, $sourceScopesCity, $sourceScopesLevel, $sourceLinksInCountMin, $sourceLinksInCountMax, $sourceRankingsAlexaRankMin, $sourceRankingsAlexaRankMax, $sourceRankingsAlexaCountry, $socialSharesCountFacebookMin, $socialSharesCountFacebookMax, $socialSharesCountGooglePlusMin, $socialSharesCountGooglePlusMax, $socialSharesCountLinkedinMin, $socialSharesCountLinkedinMax, $socialSharesCountRedditMin, $socialSharesCountRedditMax, $field)

List trends

This endpoint is used for finding trends based on stories.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | [**[]int64**](int64.md)| This parameter is used for finding stories by story id. | [optional] 
 **Title** | **string**| This parameter is used for finding stories whose title contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Body** | **string**| This parameter is used for finding stories whose body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Text** | **string**| This parameter is used for finding stories whose title or body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Language** | [**[]string**](string.md)| This parameter is used for finding stories whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes. | [optional] 
 **PublishedAtStart** | **string**| This parameter is used for finding stories whose published at time is greater than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates). | [optional] 
 **PublishedAtEnd** | **string**| This parameter is used for finding stories whose published at time is less than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates). | [optional] 
 **CategoriesTaxonomy** | **string**| This parameter is used for defining the type of the taxonomy for the rest of the categories queries. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] 
 **CategoriesConfident** | **bool**| This parameter is used for finding stories whose categories are confident. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] [default to true]
 **CategoriesId** | [**[]string**](string.md)| This parameter is used for finding stories by categories id. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] 
 **CategoriesLevel** | [**[]int32**](int32.md)| This parameter is used for finding stories by categories level. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] 
 **EntitiesTitleText** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;text&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesTitleType** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;type&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesTitleLinksDbpedia** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities dbpedia URL in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesBodyText** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;text&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesBodyType** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;type&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesBodyLinksDbpedia** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities dbpedia URL in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **SentimentTitlePolarity** | **string**| This parameter is used for finding stories whose title sentiment is the specified value. | [optional] 
 **SentimentBodyPolarity** | **string**| This parameter is used for finding stories whose body sentiment is the specified value. | [optional] 
 **MediaImagesCountMin** | **string**| This parameter is used for finding stories whose number of images is greater than or equal to the specified value. | [optional] 
 **MediaImagesCountMax** | **string**| This parameter is used for finding stories whose number of images is less than or equal to the specified value. | [optional] 
 **MediaImagesWidthMin** | **string**| This parameter is used for finding stories whose width of images are greater than or equal to the specified value. | [optional] 
 **MediaImagesWidthMax** | **string**| This parameter is used for finding stories whose width of images are less than or equal to the specified value. | [optional] 
 **MediaImagesHeightMin** | **string**| This parameter is used for finding stories whose height of images are greater than or equal to the specified value. | [optional] 
 **MediaImagesHeightMax** | **string**| This parameter is used for finding stories whose height of images are less than or equal to the specified value. | [optional] 
 **MediaImagesContentLengthMin** | **string**| This parameter is used for finding stories whose images content length are greater than or equal to the specified value. | [optional] 
 **MediaImagesContentLengthMax** | **string**| This parameter is used for finding stories whose images content length are less than or equal to the specified value. | [optional] 
 **MediaImagesFormat** | [**[]string**](string.md)| This parameter is used for finding stories whose images format are the specified value. | [optional] 
 **MediaVideosCountMin** | **string**| This parameter is used for finding stories whose number of videos is greater than or equal to the specified value. | [optional] 
 **MediaVideosCountMax** | **string**| This parameter is used for finding stories whose number of videos is less than or equal to the specified value. | [optional] 
 **AuthorId** | [**[]int32**](int32.md)| This parameter is used for finding stories whose author id is the specified value. | [optional] 
 **AuthorName** | **string**| This parameter is used for finding stories whose author full name contains the specified value. | [optional] 
 **SourceId** | [**[]int32**](int32.md)| This parameter is used for finding stories whose source id is the specified value. | [optional] 
 **SourceName** | [**[]string**](string.md)| This parameter is used for finding stories whose source name contains the specified value. | [optional] 
 **SourceDomain** | [**[]string**](string.md)| This parameter is used for finding stories whose source domain is the specified value. | [optional] 
 **SourceLocationsCountry** | [**[]string**](string.md)| This parameter is used for finding stories whose source country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceLocationsState** | [**[]string**](string.md)| This parameter is used for finding stories whose source state/province is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceLocationsCity** | [**[]string**](string.md)| This parameter is used for finding stories whose source city is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceScopesCountry** | [**[]string**](string.md)| This parameter is used for finding stories whose source scopes is the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceScopesState** | [**[]string**](string.md)| This parameter is used for finding stories whose source scopes is the specified state/province value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceScopesCity** | [**[]string**](string.md)| This parameter is used for finding stories whose source scopes is the specified city value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceScopesLevel** | [**[]string**](string.md)| This parameter is used for finding stories whose source scopes is the specified level value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceLinksInCountMin** | **int32**| This parameter is used for finding stories from sources whose Links in count is greater than or equal to the specified value. You can read more about working with Links in count [here](https://newsapi.aylien.com/docs/working-with-links-in-count). | [optional] 
 **SourceLinksInCountMax** | **int32**| This parameter is used for finding stories from sources whose Links in count is less than or equal to the specified value. You can read more about working with Links in count [here](https://newsapi.aylien.com/docs/working-with-links-in-count). | [optional] 
 **SourceRankingsAlexaRankMin** | **int32**| This parameter is used for finding stories from sources whose Alexa rank is greater than or equal to the specified value. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
 **SourceRankingsAlexaRankMax** | **int32**| This parameter is used for finding stories from sources whose Alexa rank is less than or equal to the specified value. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
 **SourceRankingsAlexaCountry** | [**[]string**](string.md)| This parameter is used for finding stories from sources whose Alexa rank is in the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
 **SocialSharesCountFacebookMin** | **string**| This parameter is used for finding stories whose Facebook social shares count is greater than or equal to the specified value. | [optional] 
 **SocialSharesCountFacebookMax** | **string**| This parameter is used for finding stories whose Facebook social shares count is less than or equal to the specified value. | [optional] 
 **SocialSharesCountGooglePlusMin** | **string**| This parameter is used for finding stories whose Google+ social shares count is greater than or equal to the specified value. | [optional] 
 **SocialSharesCountGooglePlusMax** | **string**| This parameter is used for finding stories whose Google+ social shares count is less than or equal to the specified value. | [optional] 
 **SocialSharesCountLinkedinMin** | **string**| This parameter is used for finding stories whose LinkedIn social shares count is greater than or equal to the specified value. | [optional] 
 **SocialSharesCountLinkedinMax** | **string**| This parameter is used for finding stories whose LinkedIn social shares count is less than or equal to the specified value. | [optional] 
 **SocialSharesCountRedditMin** | **string**| This parameter is used for finding stories whose Reddit social shares count is greater than or equal to the specified value. | [optional] 
 **SocialSharesCountRedditMax** | **string**| This parameter is used for finding stories whose Reddit social shares count is less than or equal to the specified value. | [optional] 
 **Field** | **string**| This parameter is used to specify the trend field. | [optional] 

### Return type

[**Trends**](Trends.md)

### Authorization

[app_key](../README.md#app_key), [app_id](../README.md#app_id)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

