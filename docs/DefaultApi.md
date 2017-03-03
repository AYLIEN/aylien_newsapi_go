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
> Autocompletes ListAutocompletes(params)

List autocompletes

This endpoint is used for getting list of autocompletes by providing a specific term and type.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Type** | **string**| This parameter is used for defining the type of autocompletes. |  [enum: source_names, source_domains, entity_types, dbpedia_resources]
 **Term** | **string**| This parameter is used for finding autocomplete objects that contain the specified value. | 
 **Language** | **string**| This parameter is used for autocompletes whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes. | [optional] [default to en] [enum: en, de, fr, it, es, pt]
 **PerPage** | **int32**| This parameter is used for specifying number of items in each page. | [optional] [default to 3]

### Return type

[**Autocompletes**](Autocompletes.md)

### Authorization

[app_key](../README.md#app_key), [app_id](../README.md#app_id)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCoverages**
> Coverages ListCoverages(params)

List coverages

This endpoint is used for finding story coverages based on the parameters provided. The maximum number of related stories returned is 100.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | **[]int64**| This parameter is used for finding stories by story id. | [optional] 
 **NotId** | **[]int64**| This parameter is used for excluding stories by story id. | [optional] 
 **Title** | **string**| This parameter is used for finding stories whose title contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Body** | **string**| This parameter is used for finding stories whose body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Text** | **string**| This parameter is used for finding stories whose title or body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Language** | **[]string**| This parameter is used for finding stories whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes. | [optional]  [enum: en, de, fr, it, es, pt]
 **NotLanguage** | **[]string**| This parameter is used for excluding stories whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes. | [optional]  [enum: en, de, fr, it, es, pt]
 **PublishedAtStart** | **string**| This parameter is used for finding stories whose published at time is greater than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates). | [optional] 
 **PublishedAtEnd** | **string**| This parameter is used for finding stories whose published at time is less than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates). | [optional] 
 **CategoriesTaxonomy** | **string**| This parameter is used for defining the type of the taxonomy for the rest of the categories queries. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional]  [enum: iab-qag, iptc-subjectcode]
 **CategoriesConfident** | **bool**| This parameter is used for finding stories whose categories are confident. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] [default to true]
 **CategoriesId** | **[]string**| This parameter is used for finding stories by categories id. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] 
 **NotCategoriesId** | **[]string**| This parameter is used for excluding stories by categories id. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] 
 **CategoriesLevel** | **[]int32**| This parameter is used for finding stories by categories level. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] 
 **NotCategoriesLevel** | **[]int32**| This parameter is used for excluding stories by categories level. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] 
 **EntitiesTitleText** | **[]string**| This parameter is used to find stories based on the specified entities &#x60;text&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **NotEntitiesTitleText** | **[]string**| This parameter is used to exclude stories based on the specified entities &#x60;text&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesTitleType** | **[]string**| This parameter is used to find stories based on the specified entities &#x60;type&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **NotEntitiesTitleType** | **[]string**| This parameter is used to exclude stories based on the specified entities &#x60;type&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesTitleLinksDbpedia** | **[]string**| This parameter is used to find stories based on the specified entities dbpedia URL in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **NotEntitiesTitleLinksDbpedia** | **[]string**| This parameter is used to exclude stories based on the specified entities dbpedia URL in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesBodyText** | **[]string**| This parameter is used to find stories based on the specified entities &#x60;text&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **NotEntitiesBodyText** | **[]string**| This parameter is used to exclude stories based on the specified entities &#x60;text&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesBodyType** | **[]string**| This parameter is used to find stories based on the specified entities &#x60;type&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **NotEntitiesBodyType** | **[]string**| This parameter is used to exclude stories based on the specified entities &#x60;type&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesBodyLinksDbpedia** | **[]string**| This parameter is used to find stories based on the specified entities dbpedia URL in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **NotEntitiesBodyLinksDbpedia** | **[]string**| This parameter is used to exclude stories based on the specified entities dbpedia URL in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **SentimentTitlePolarity** | **string**| This parameter is used for finding stories whose title sentiment is the specified value. | [optional]  [enum: positive, neutral, negative]
 **NotSentimentTitlePolarity** | **string**| This parameter is used for excluding stories whose title sentiment is the specified value. | [optional]  [enum: positive, neutral, negative]
 **SentimentBodyPolarity** | **string**| This parameter is used for finding stories whose body sentiment is the specified value. | [optional]  [enum: positive, neutral, negative]
 **NotSentimentBodyPolarity** | **string**| This parameter is used for excluding stories whose body sentiment is the specified value. | [optional]  [enum: positive, neutral, negative]
 **MediaImagesCountMin** | **int32**| This parameter is used for finding stories whose number of images is greater than or equal to the specified value. | [optional] 
 **MediaImagesCountMax** | **int32**| This parameter is used for finding stories whose number of images is less than or equal to the specified value. | [optional] 
 **MediaImagesWidthMin** | **int32**| This parameter is used for finding stories whose width of images are greater than or equal to the specified value. | [optional] 
 **MediaImagesWidthMax** | **int32**| This parameter is used for finding stories whose width of images are less than or equal to the specified value. | [optional] 
 **MediaImagesHeightMin** | **int32**| This parameter is used for finding stories whose height of images are greater than or equal to the specified value. | [optional] 
 **MediaImagesHeightMax** | **int32**| This parameter is used for finding stories whose height of images are less than or equal to the specified value. | [optional] 
 **MediaImagesContentLengthMin** | **int32**| This parameter is used for finding stories whose images content length are greater than or equal to the specified value. | [optional] 
 **MediaImagesContentLengthMax** | **int32**| This parameter is used for finding stories whose images content length are less than or equal to the specified value. | [optional] 
 **MediaImagesFormat** | **[]string**| This parameter is used for finding stories whose images format are the specified value. | [optional]  [enum: BMP, GIF, JPEG, PNG, TIFF, PSD, ICO, CUR, WEBP, SVG]
 **NotMediaImagesFormat** | **[]string**| This parameter is used for excluding stories whose images format are the specified value. | [optional]  [enum: BMP, GIF, JPEG, PNG, TIFF, PSD, ICO, CUR, WEBP, SVG]
 **MediaVideosCountMin** | **int32**| This parameter is used for finding stories whose number of videos is greater than or equal to the specified value. | [optional] 
 **MediaVideosCountMax** | **int32**| This parameter is used for finding stories whose number of videos is less than or equal to the specified value. | [optional] 
 **AuthorId** | **[]int32**| This parameter is used for finding stories whose author id is the specified value. | [optional] 
 **NotAuthorId** | **[]int32**| This parameter is used for excluding stories whose author id is the specified value. | [optional] 
 **AuthorName** | **string**| This parameter is used for finding stories whose author full name contains the specified value. | [optional] 
 **NotAuthorName** | **string**| This parameter is used for excluding stories whose author full name contains the specified value. | [optional] 
 **SourceId** | **[]int32**| This parameter is used for finding stories whose source id is the specified value. | [optional] 
 **NotSourceId** | **[]int32**| This parameter is used for excluding stories whose source id is the specified value. | [optional] 
 **SourceName** | **[]string**| This parameter is used for finding stories whose source name contains the specified value. | [optional] 
 **NotSourceName** | **[]string**| This parameter is used for excluding stories whose source name contains the specified value. | [optional] 
 **SourceDomain** | **[]string**| This parameter is used for finding stories whose source domain is the specified value. | [optional] 
 **NotSourceDomain** | **[]string**| This parameter is used for excluding stories whose source domain is the specified value. | [optional] 
 **SourceLocationsCountry** | **[]string**| This parameter is used for finding stories whose source country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **NotSourceLocationsCountry** | **[]string**| This parameter is used for excluding stories whose source country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceLocationsState** | **[]string**| This parameter is used for finding stories whose source state/province is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **NotSourceLocationsState** | **[]string**| This parameter is used for excluding stories whose source state/province is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceLocationsCity** | **[]string**| This parameter is used for finding stories whose source city is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **NotSourceLocationsCity** | **[]string**| This parameter is used for excluding stories whose source city is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceScopesCountry** | **[]string**| This parameter is used for finding stories whose source scopes  is the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **NotSourceScopesCountry** | **[]string**| This parameter is used for excluding stories whose source scopes  is the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceScopesState** | **[]string**| This parameter is used for finding stories whose source scopes is the specified state/province value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **NotSourceScopesState** | **[]string**| This parameter is used for excluding stories whose source scopes is the specified state/province value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceScopesCity** | **[]string**| This parameter is used for finding stories whose source scopes is the specified city value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **NotSourceScopesCity** | **[]string**| This parameter is used for excluding stories whose source scopes is the specified city value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceScopesLevel** | **[]string**| This parameter is used for finding stories whose source scopes  is the specified level value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional]  [enum: international, national, local]
 **NotSourceScopesLevel** | **[]string**| This parameter is used for excluding stories whose source scopes  is the specified level value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional]  [enum: international, national, local]
 **SourceLinksInCountMin** | **int32**| This parameter is used for finding stories from sources whose Links in count is greater than or equal to the specified value. You can read more about working with Links in count [here](https://newsapi.aylien.com/docs/working-with-links-in-count). | [optional] 
 **SourceLinksInCountMax** | **int32**| This parameter is used for finding stories from sources whose Links in count is less than or equal to the specified value. You can read more about working with Links in count [here](https://newsapi.aylien.com/docs/working-with-links-in-count). | [optional] 
 **SourceRankingsAlexaRankMin** | **int32**| This parameter is used for finding stories from sources whose Alexa rank is greater than or equal to the specified value. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
 **SourceRankingsAlexaRankMax** | **int32**| This parameter is used for finding stories from sources whose Alexa rank is less than or equal to the specified value. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
 **SourceRankingsAlexaCountry** | **[]string**| This parameter is used for finding stories from sources whose Alexa rank is in the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
 **SocialSharesCountFacebookMin** | **int32**| This parameter is used for finding stories whose Facebook social shares count is greater than or equal to the specified value. | [optional] 
 **SocialSharesCountFacebookMax** | **int32**| This parameter is used for finding stories whose Facebook social shares count is less than or equal to the specified value. | [optional] 
 **SocialSharesCountGooglePlusMin** | **int32**| This parameter is used for finding stories whose Google+ social shares count is greater than or equal to the specified value. | [optional] 
 **SocialSharesCountGooglePlusMax** | **int32**| This parameter is used for finding stories whose Google+ social shares count is less than or equal to the specified value. | [optional] 
 **SocialSharesCountLinkedinMin** | **int32**| This parameter is used for finding stories whose LinkedIn social shares count is greater than or equal to the specified value. | [optional] 
 **SocialSharesCountLinkedinMax** | **int32**| This parameter is used for finding stories whose LinkedIn social shares count is less than or equal to the specified value. | [optional] 
 **SocialSharesCountRedditMin** | **int32**| This parameter is used for finding stories whose Reddit social shares count is greater than or equal to the specified value. | [optional] 
 **SocialSharesCountRedditMax** | **int32**| This parameter is used for finding stories whose Reddit social shares count is less than or equal to the specified value. | [optional] 
 **Cluster** | **bool**| This parameter enables clustering for the returned stories. You can read more about working with clustering [here](https://newsapi.aylien.com/docs/working-with-clustering). | [optional] [default to false]
 **ClusterAlgorithm** | **string**| This parameter is used for specifying the clustering algorithm you wish to use. It supprts STC, Lingo and [k-means](https://en.wikipedia.org/wiki/K-means_clustering) algorithms. You can read more about working with clustering [here](https://newsapi.aylien.com/docs/working-with-clustering). | [optional] [default to lingo] [enum: stc, lingo, kmeans]
 **Return** | **[]string**| This parameter is used for specifying return fields. | [optional]  [enum: id, title, body, summary, source, author, entities, keywords, hashtags, characters_count, words_count, sentences_count, paragraphs_count, categories, social_shares_count, media, sentiment, language, published_at, links]
 **StoryId** | **int64**| A story id | [optional] 
 **StoryUrl** | **string**| An article or webpage | [optional] 
 **StoryTitle** | **string**| Title of the article | [optional] 
 **StoryBody** | **string**| Body of the article | [optional] 
 **StoryPublishedAt** | **time.Time**| Publish date of the article. If you use a url or title and body of an article for getting coverages, this parameter is required. The format used is a restricted form of the canonical representation of dateTime in the [XML Schema specification (ISO 8601)](https://www.w3.org/TR/xmlschema-2/#dateTime). | [optional] 
 **StoryLanguage** | **string**| This parameter is used for setting the language of the story. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes. | [optional] [default to auto] [enum: auto, en, de, fr, it, es, pt]
 **PerPage** | **int32**| This parameter is used for specifying number of items in each page. | [optional] [default to 3]

### Return type

[**Coverages**](Coverages.md)

### Authorization

[app_key](../README.md#app_key), [app_id](../README.md#app_id)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListHistograms**
> Histograms ListHistograms(params)

List histograms

This endpoint is used for getting histograms based on the `field` parameter passed to the API.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | **[]int64**| This parameter is used for finding stories by story id. | [optional] 
 **NotId** | **[]int64**| This parameter is used for excluding stories by story id. | [optional] 
 **Title** | **string**| This parameter is used for finding stories whose title contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Body** | **string**| This parameter is used for finding stories whose body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Text** | **string**| This parameter is used for finding stories whose title or body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Language** | **[]string**| This parameter is used for finding stories whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes. | [optional]  [enum: en, de, fr, it, es, pt]
 **NotLanguage** | **[]string**| This parameter is used for excluding stories whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes. | [optional]  [enum: en, de, fr, it, es, pt]
 **PublishedAtStart** | **string**| This parameter is used for finding stories whose published at time is greater than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates). | [optional] 
 **PublishedAtEnd** | **string**| This parameter is used for finding stories whose published at time is less than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates). | [optional] 
 **CategoriesTaxonomy** | **string**| This parameter is used for defining the type of the taxonomy for the rest of the categories queries. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional]  [enum: iab-qag, iptc-subjectcode]
 **CategoriesConfident** | **bool**| This parameter is used for finding stories whose categories are confident. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] [default to true]
 **CategoriesId** | **[]string**| This parameter is used for finding stories by categories id. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] 
 **NotCategoriesId** | **[]string**| This parameter is used for excluding stories by categories id. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] 
 **CategoriesLevel** | **[]int32**| This parameter is used for finding stories by categories level. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] 
 **NotCategoriesLevel** | **[]int32**| This parameter is used for excluding stories by categories level. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] 
 **EntitiesTitleText** | **[]string**| This parameter is used to find stories based on the specified entities &#x60;text&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **NotEntitiesTitleText** | **[]string**| This parameter is used to exclude stories based on the specified entities &#x60;text&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesTitleType** | **[]string**| This parameter is used to find stories based on the specified entities &#x60;type&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **NotEntitiesTitleType** | **[]string**| This parameter is used to exclude stories based on the specified entities &#x60;type&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesTitleLinksDbpedia** | **[]string**| This parameter is used to find stories based on the specified entities dbpedia URL in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **NotEntitiesTitleLinksDbpedia** | **[]string**| This parameter is used to exclude stories based on the specified entities dbpedia URL in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesBodyText** | **[]string**| This parameter is used to find stories based on the specified entities &#x60;text&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **NotEntitiesBodyText** | **[]string**| This parameter is used to exclude stories based on the specified entities &#x60;text&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesBodyType** | **[]string**| This parameter is used to find stories based on the specified entities &#x60;type&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **NotEntitiesBodyType** | **[]string**| This parameter is used to exclude stories based on the specified entities &#x60;type&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesBodyLinksDbpedia** | **[]string**| This parameter is used to find stories based on the specified entities dbpedia URL in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **NotEntitiesBodyLinksDbpedia** | **[]string**| This parameter is used to exclude stories based on the specified entities dbpedia URL in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **SentimentTitlePolarity** | **string**| This parameter is used for finding stories whose title sentiment is the specified value. | [optional]  [enum: positive, neutral, negative]
 **NotSentimentTitlePolarity** | **string**| This parameter is used for excluding stories whose title sentiment is the specified value. | [optional]  [enum: positive, neutral, negative]
 **SentimentBodyPolarity** | **string**| This parameter is used for finding stories whose body sentiment is the specified value. | [optional]  [enum: positive, neutral, negative]
 **NotSentimentBodyPolarity** | **string**| This parameter is used for excluding stories whose body sentiment is the specified value. | [optional]  [enum: positive, neutral, negative]
 **MediaImagesCountMin** | **int32**| This parameter is used for finding stories whose number of images is greater than or equal to the specified value. | [optional] 
 **MediaImagesCountMax** | **int32**| This parameter is used for finding stories whose number of images is less than or equal to the specified value. | [optional] 
 **MediaImagesWidthMin** | **int32**| This parameter is used for finding stories whose width of images are greater than or equal to the specified value. | [optional] 
 **MediaImagesWidthMax** | **int32**| This parameter is used for finding stories whose width of images are less than or equal to the specified value. | [optional] 
 **MediaImagesHeightMin** | **int32**| This parameter is used for finding stories whose height of images are greater than or equal to the specified value. | [optional] 
 **MediaImagesHeightMax** | **int32**| This parameter is used for finding stories whose height of images are less than or equal to the specified value. | [optional] 
 **MediaImagesContentLengthMin** | **int32**| This parameter is used for finding stories whose images content length are greater than or equal to the specified value. | [optional] 
 **MediaImagesContentLengthMax** | **int32**| This parameter is used for finding stories whose images content length are less than or equal to the specified value. | [optional] 
 **MediaImagesFormat** | **[]string**| This parameter is used for finding stories whose images format are the specified value. | [optional]  [enum: BMP, GIF, JPEG, PNG, TIFF, PSD, ICO, CUR, WEBP, SVG]
 **NotMediaImagesFormat** | **[]string**| This parameter is used for excluding stories whose images format are the specified value. | [optional]  [enum: BMP, GIF, JPEG, PNG, TIFF, PSD, ICO, CUR, WEBP, SVG]
 **MediaVideosCountMin** | **int32**| This parameter is used for finding stories whose number of videos is greater than or equal to the specified value. | [optional] 
 **MediaVideosCountMax** | **int32**| This parameter is used for finding stories whose number of videos is less than or equal to the specified value. | [optional] 
 **AuthorId** | **[]int32**| This parameter is used for finding stories whose author id is the specified value. | [optional] 
 **NotAuthorId** | **[]int32**| This parameter is used for excluding stories whose author id is the specified value. | [optional] 
 **AuthorName** | **string**| This parameter is used for finding stories whose author full name contains the specified value. | [optional] 
 **NotAuthorName** | **string**| This parameter is used for excluding stories whose author full name contains the specified value. | [optional] 
 **SourceId** | **[]int32**| This parameter is used for finding stories whose source id is the specified value. | [optional] 
 **NotSourceId** | **[]int32**| This parameter is used for excluding stories whose source id is the specified value. | [optional] 
 **SourceName** | **[]string**| This parameter is used for finding stories whose source name contains the specified value. | [optional] 
 **NotSourceName** | **[]string**| This parameter is used for excluding stories whose source name contains the specified value. | [optional] 
 **SourceDomain** | **[]string**| This parameter is used for finding stories whose source domain is the specified value. | [optional] 
 **NotSourceDomain** | **[]string**| This parameter is used for excluding stories whose source domain is the specified value. | [optional] 
 **SourceLocationsCountry** | **[]string**| This parameter is used for finding stories whose source country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **NotSourceLocationsCountry** | **[]string**| This parameter is used for excluding stories whose source country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceLocationsState** | **[]string**| This parameter is used for finding stories whose source state/province is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **NotSourceLocationsState** | **[]string**| This parameter is used for excluding stories whose source state/province is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceLocationsCity** | **[]string**| This parameter is used for finding stories whose source city is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **NotSourceLocationsCity** | **[]string**| This parameter is used for excluding stories whose source city is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceScopesCountry** | **[]string**| This parameter is used for finding stories whose source scopes is the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **NotSourceScopesCountry** | **[]string**| This parameter is used for excluding stories whose source scopes is the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceScopesState** | **[]string**| This parameter is used for finding stories whose source scopes is the specified state/province value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **NotSourceScopesState** | **[]string**| This parameter is used for excluding stories whose source scopes is the specified state/province value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceScopesCity** | **[]string**| This parameter is used for finding stories whose source scopes is the specified city value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **NotSourceScopesCity** | **[]string**| This parameter is used for excluding stories whose source scopes is the specified city value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceScopesLevel** | **[]string**| This parameter is used for finding stories whose source scopes is the specified level value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional]  [enum: international, national, local]
 **NotSourceScopesLevel** | **[]string**| This parameter is used for excluding stories whose source scopes is the specified level value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional]  [enum: international, national, local]
 **SourceLinksInCountMin** | **int32**| This parameter is used for finding stories from sources whose Links in count is greater than or equal to the specified value. You can read more about working with Links in count [here](https://newsapi.aylien.com/docs/working-with-links-in-count). | [optional] 
 **SourceLinksInCountMax** | **int32**| This parameter is used for finding stories from sources whose Links in count is less than or equal to the specified value. You can read more about working with Links in count [here](https://newsapi.aylien.com/docs/working-with-links-in-count). | [optional] 
 **SourceRankingsAlexaRankMin** | **int32**| This parameter is used for finding stories from sources whose Alexa rank is greater than or equal to the specified value. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
 **SourceRankingsAlexaRankMax** | **int32**| This parameter is used for finding stories from sources whose Alexa rank is less than or equal to the specified value. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
 **SourceRankingsAlexaCountry** | **[]string**| This parameter is used for finding stories from sources whose Alexa rank is in the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
 **SocialSharesCountFacebookMin** | **int32**| This parameter is used for finding stories whose Facebook social shares count is greater than or equal to the specified value. | [optional] 
 **SocialSharesCountFacebookMax** | **int32**| This parameter is used for finding stories whose Facebook social shares count is less than or equal to the specified value. | [optional] 
 **SocialSharesCountGooglePlusMin** | **int32**| This parameter is used for finding stories whose Google+ social shares count is greater than or equal to the specified value. | [optional] 
 **SocialSharesCountGooglePlusMax** | **int32**| This parameter is used for finding stories whose Google+ social shares count is less than or equal to the specified value. | [optional] 
 **SocialSharesCountLinkedinMin** | **int32**| This parameter is used for finding stories whose LinkedIn social shares count is greater than or equal to the specified value. | [optional] 
 **SocialSharesCountLinkedinMax** | **int32**| This parameter is used for finding stories whose LinkedIn social shares count is less than or equal to the specified value. | [optional] 
 **SocialSharesCountRedditMin** | **int32**| This parameter is used for finding stories whose Reddit social shares count is greater than or equal to the specified value. | [optional] 
 **SocialSharesCountRedditMax** | **int32**| This parameter is used for finding stories whose Reddit social shares count is less than or equal to the specified value. | [optional] 
 **IntervalStart** | **int32**| This parameter is used for setting the start data point of histogram intervals. | [optional] 
 **IntervalEnd** | **int32**| This parameter is used for setting the end data point of histogram intervals. | [optional] 
 **IntervalWidth** | **int32**| This parameter is used for setting the width of histogram intervals. | [optional] 
 **Field** | **string**| This parameter is used for specifying the y-axis variable for the histogram. | [optional] [default to social_shares_count] [enum: social_shares_count, social_shares_count.facebook, social_shares_count.linkedin, social_shares_count.reddit, social_shares_count.google_plus, characters_count, words_count, sentences_count, paragraphs_count, media.images.count, media.videos.count, source.links_in_count, source.rankings.alexa.rank, source.rankings.alexa.rank.AF, source.rankings.alexa.rank.AX, source.rankings.alexa.rank.AL, source.rankings.alexa.rank.DZ, source.rankings.alexa.rank.AS, source.rankings.alexa.rank.AD, source.rankings.alexa.rank.AO, source.rankings.alexa.rank.AI, source.rankings.alexa.rank.AQ, source.rankings.alexa.rank.AG, source.rankings.alexa.rank.AR, source.rankings.alexa.rank.AM, source.rankings.alexa.rank.AW, source.rankings.alexa.rank.AU, source.rankings.alexa.rank.AT, source.rankings.alexa.rank.AZ, source.rankings.alexa.rank.BS, source.rankings.alexa.rank.BH, source.rankings.alexa.rank.BD, source.rankings.alexa.rank.BB, source.rankings.alexa.rank.BY, source.rankings.alexa.rank.BE, source.rankings.alexa.rank.BZ, source.rankings.alexa.rank.BJ, source.rankings.alexa.rank.BM, source.rankings.alexa.rank.BT, source.rankings.alexa.rank.BO, source.rankings.alexa.rank.BQ, source.rankings.alexa.rank.BA, source.rankings.alexa.rank.BW, source.rankings.alexa.rank.BV, source.rankings.alexa.rank.BR, source.rankings.alexa.rank.IO, source.rankings.alexa.rank.BN, source.rankings.alexa.rank.BG, source.rankings.alexa.rank.BF, source.rankings.alexa.rank.BI, source.rankings.alexa.rank.KH, source.rankings.alexa.rank.CM, source.rankings.alexa.rank.CA, source.rankings.alexa.rank.CV, source.rankings.alexa.rank.KY, source.rankings.alexa.rank.CF, source.rankings.alexa.rank.TD, source.rankings.alexa.rank.CL, source.rankings.alexa.rank.CN, source.rankings.alexa.rank.CX, source.rankings.alexa.rank.CC, source.rankings.alexa.rank.CO, source.rankings.alexa.rank.KM, source.rankings.alexa.rank.CG, source.rankings.alexa.rank.CD, source.rankings.alexa.rank.CK, source.rankings.alexa.rank.CR, source.rankings.alexa.rank.CI, source.rankings.alexa.rank.HR, source.rankings.alexa.rank.CU, source.rankings.alexa.rank.CW, source.rankings.alexa.rank.CY, source.rankings.alexa.rank.CZ, source.rankings.alexa.rank.DK, source.rankings.alexa.rank.DJ, source.rankings.alexa.rank.DM, source.rankings.alexa.rank.DO, source.rankings.alexa.rank.EC, source.rankings.alexa.rank.EG, source.rankings.alexa.rank.SV, source.rankings.alexa.rank.GQ, source.rankings.alexa.rank.ER, source.rankings.alexa.rank.EE, source.rankings.alexa.rank.ET, source.rankings.alexa.rank.FK, source.rankings.alexa.rank.FO, source.rankings.alexa.rank.FJ, source.rankings.alexa.rank.FI, source.rankings.alexa.rank.FR, source.rankings.alexa.rank.GF, source.rankings.alexa.rank.PF, source.rankings.alexa.rank.TF, source.rankings.alexa.rank.GA, source.rankings.alexa.rank.GM, source.rankings.alexa.rank.GE, source.rankings.alexa.rank.DE, source.rankings.alexa.rank.GH, source.rankings.alexa.rank.GI, source.rankings.alexa.rank.GR, source.rankings.alexa.rank.GL, source.rankings.alexa.rank.GD, source.rankings.alexa.rank.GP, source.rankings.alexa.rank.GU, source.rankings.alexa.rank.GT, source.rankings.alexa.rank.GG, source.rankings.alexa.rank.GN, source.rankings.alexa.rank.GW, source.rankings.alexa.rank.GY, source.rankings.alexa.rank.HT, source.rankings.alexa.rank.HM, source.rankings.alexa.rank.VA, source.rankings.alexa.rank.HN, source.rankings.alexa.rank.HK, source.rankings.alexa.rank.HU, source.rankings.alexa.rank.IS, source.rankings.alexa.rank.IN, source.rankings.alexa.rank.ID, source.rankings.alexa.rank.IR, source.rankings.alexa.rank.IQ, source.rankings.alexa.rank.IE, source.rankings.alexa.rank.IM, source.rankings.alexa.rank.IL, source.rankings.alexa.rank.IT, source.rankings.alexa.rank.JM, source.rankings.alexa.rank.JP, source.rankings.alexa.rank.JE, source.rankings.alexa.rank.JO, source.rankings.alexa.rank.KZ, source.rankings.alexa.rank.KE, source.rankings.alexa.rank.KI, source.rankings.alexa.rank.KP, source.rankings.alexa.rank.KR, source.rankings.alexa.rank.KW, source.rankings.alexa.rank.KG, source.rankings.alexa.rank.LA, source.rankings.alexa.rank.LV, source.rankings.alexa.rank.LB, source.rankings.alexa.rank.LS, source.rankings.alexa.rank.LR, source.rankings.alexa.rank.LY, source.rankings.alexa.rank.LI, source.rankings.alexa.rank.LT, source.rankings.alexa.rank.LU, source.rankings.alexa.rank.MO, source.rankings.alexa.rank.MK, source.rankings.alexa.rank.MG, source.rankings.alexa.rank.MW, source.rankings.alexa.rank.MY, source.rankings.alexa.rank.MV, source.rankings.alexa.rank.ML, source.rankings.alexa.rank.MT, source.rankings.alexa.rank.MH, source.rankings.alexa.rank.MQ, source.rankings.alexa.rank.MR, source.rankings.alexa.rank.MU, source.rankings.alexa.rank.YT, source.rankings.alexa.rank.MX, source.rankings.alexa.rank.FM, source.rankings.alexa.rank.MD, source.rankings.alexa.rank.MC, source.rankings.alexa.rank.MN, source.rankings.alexa.rank.ME, source.rankings.alexa.rank.MS, source.rankings.alexa.rank.MA, source.rankings.alexa.rank.MZ, source.rankings.alexa.rank.MM, source.rankings.alexa.rank.NA, source.rankings.alexa.rank.NR, source.rankings.alexa.rank.NP, source.rankings.alexa.rank.NL, source.rankings.alexa.rank.NC, source.rankings.alexa.rank.NZ, source.rankings.alexa.rank.NI, source.rankings.alexa.rank.NE, source.rankings.alexa.rank.NG, source.rankings.alexa.rank.NU, source.rankings.alexa.rank.NF, source.rankings.alexa.rank.MP, source.rankings.alexa.rank.NO, source.rankings.alexa.rank.OM, source.rankings.alexa.rank.PK, source.rankings.alexa.rank.PW, source.rankings.alexa.rank.PS, source.rankings.alexa.rank.PA, source.rankings.alexa.rank.PG, source.rankings.alexa.rank.PY, source.rankings.alexa.rank.PE, source.rankings.alexa.rank.PH, source.rankings.alexa.rank.PN, source.rankings.alexa.rank.PL, source.rankings.alexa.rank.PT, source.rankings.alexa.rank.PR, source.rankings.alexa.rank.QA, source.rankings.alexa.rank.RE, source.rankings.alexa.rank.RO, source.rankings.alexa.rank.RU, source.rankings.alexa.rank.RW, source.rankings.alexa.rank.BL, source.rankings.alexa.rank.SH, source.rankings.alexa.rank.KN, source.rankings.alexa.rank.LC, source.rankings.alexa.rank.MF, source.rankings.alexa.rank.PM, source.rankings.alexa.rank.VC, source.rankings.alexa.rank.WS, source.rankings.alexa.rank.SM, source.rankings.alexa.rank.ST, source.rankings.alexa.rank.SA, source.rankings.alexa.rank.SN, source.rankings.alexa.rank.RS, source.rankings.alexa.rank.SC, source.rankings.alexa.rank.SL, source.rankings.alexa.rank.SG, source.rankings.alexa.rank.SX, source.rankings.alexa.rank.SK, source.rankings.alexa.rank.SI, source.rankings.alexa.rank.SB, source.rankings.alexa.rank.SO, source.rankings.alexa.rank.ZA, source.rankings.alexa.rank.GS, source.rankings.alexa.rank.SS, source.rankings.alexa.rank.ES, source.rankings.alexa.rank.LK, source.rankings.alexa.rank.SD, source.rankings.alexa.rank.SR, source.rankings.alexa.rank.SJ, source.rankings.alexa.rank.SZ, source.rankings.alexa.rank.SE, source.rankings.alexa.rank.CH, source.rankings.alexa.rank.SY, source.rankings.alexa.rank.TW, source.rankings.alexa.rank.TJ, source.rankings.alexa.rank.TZ, source.rankings.alexa.rank.TH, source.rankings.alexa.rank.TL, source.rankings.alexa.rank.TG, source.rankings.alexa.rank.TK, source.rankings.alexa.rank.TO, source.rankings.alexa.rank.TT, source.rankings.alexa.rank.TN, source.rankings.alexa.rank.TR, source.rankings.alexa.rank.TM, source.rankings.alexa.rank.TC, source.rankings.alexa.rank.TV, source.rankings.alexa.rank.UG, source.rankings.alexa.rank.UA, source.rankings.alexa.rank.AE, source.rankings.alexa.rank.GB, source.rankings.alexa.rank.US, source.rankings.alexa.rank.UM, source.rankings.alexa.rank.UY, source.rankings.alexa.rank.UZ, source.rankings.alexa.rank.VU, source.rankings.alexa.rank.VE, source.rankings.alexa.rank.VN, source.rankings.alexa.rank.VG, source.rankings.alexa.rank.VI, source.rankings.alexa.rank.WF, source.rankings.alexa.rank.EH, source.rankings.alexa.rank.YE, source.rankings.alexa.rank.ZM, source.rankings.alexa.rank.ZW]

### Return type

[**Histograms**](Histograms.md)

### Authorization

[app_key](../README.md#app_key), [app_id](../README.md#app_id)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRelatedStories**
> RelatedStories ListRelatedStories(params)

List related stories

This endpoint is used for finding related stories based on the parameters provided. The maximum number of related stories returned is 100.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | **[]int64**| This parameter is used for finding stories by story id. | [optional] 
 **NotId** | **[]int64**| This parameter is used for excluding stories by story id. | [optional] 
 **Title** | **string**| This parameter is used for finding stories whose title contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Body** | **string**| This parameter is used for finding stories whose body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Text** | **string**| This parameter is used for finding stories whose title or body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Language** | **[]string**| This parameter is used for finding stories whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes. | [optional]  [enum: en, de, fr, it, es, pt]
 **NotLanguage** | **[]string**| This parameter is used for excluding stories whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes. | [optional]  [enum: en, de, fr, it, es, pt]
 **PublishedAtStart** | **string**| This parameter is used for finding stories whose published at time is greater than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates). | [optional] 
 **PublishedAtEnd** | **string**| This parameter is used for finding stories whose published at time is less than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates). | [optional] 
 **CategoriesTaxonomy** | **string**| This parameter is used for defining the type of the taxonomy for the rest of the categories queries. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional]  [enum: iab-qag, iptc-subjectcode]
 **CategoriesConfident** | **bool**| This parameter is used for finding stories whose categories are confident. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] [default to true]
 **CategoriesId** | **[]string**| This parameter is used for finding stories by categories id. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] 
 **NotCategoriesId** | **[]string**| This parameter is used for excluding stories by categories id. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] 
 **CategoriesLevel** | **[]int32**| This parameter is used for finding stories by categories level. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] 
 **NotCategoriesLevel** | **[]int32**| This parameter is used for excluding stories by categories level. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] 
 **EntitiesTitleText** | **[]string**| This parameter is used to find stories based on the specified entities &#x60;text&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **NotEntitiesTitleText** | **[]string**| This parameter is used to exclude stories based on the specified entities &#x60;text&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesTitleType** | **[]string**| This parameter is used to find stories based on the specified entities &#x60;type&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **NotEntitiesTitleType** | **[]string**| This parameter is used to exclude stories based on the specified entities &#x60;type&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesTitleLinksDbpedia** | **[]string**| This parameter is used to find stories based on the specified entities dbpedia URL in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **NotEntitiesTitleLinksDbpedia** | **[]string**| This parameter is used to exclude stories based on the specified entities dbpedia URL in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesBodyText** | **[]string**| This parameter is used to find stories based on the specified entities &#x60;text&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **NotEntitiesBodyText** | **[]string**| This parameter is used to exclude stories based on the specified entities &#x60;text&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesBodyType** | **[]string**| This parameter is used to find stories based on the specified entities &#x60;type&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **NotEntitiesBodyType** | **[]string**| This parameter is used to exclude stories based on the specified entities &#x60;type&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesBodyLinksDbpedia** | **[]string**| This parameter is used to find stories based on the specified entities dbpedia URL in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **NotEntitiesBodyLinksDbpedia** | **[]string**| This parameter is used to exclude stories based on the specified entities dbpedia URL in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **SentimentTitlePolarity** | **string**| This parameter is used for finding stories whose title sentiment is the specified value. | [optional]  [enum: positive, neutral, negative]
 **NotSentimentTitlePolarity** | **string**| This parameter is used for excluding stories whose title sentiment is the specified value. | [optional]  [enum: positive, neutral, negative]
 **SentimentBodyPolarity** | **string**| This parameter is used for finding stories whose body sentiment is the specified value. | [optional]  [enum: positive, neutral, negative]
 **NotSentimentBodyPolarity** | **string**| This parameter is used for excluding stories whose body sentiment is the specified value. | [optional]  [enum: positive, neutral, negative]
 **MediaImagesCountMin** | **int32**| This parameter is used for finding stories whose number of images is greater than or equal to the specified value. | [optional] 
 **MediaImagesCountMax** | **int32**| This parameter is used for finding stories whose number of images is less than or equal to the specified value. | [optional] 
 **MediaImagesWidthMin** | **int32**| This parameter is used for finding stories whose width of images are greater than or equal to the specified value. | [optional] 
 **MediaImagesWidthMax** | **int32**| This parameter is used for finding stories whose width of images are less than or equal to the specified value. | [optional] 
 **MediaImagesHeightMin** | **int32**| This parameter is used for finding stories whose height of images are greater than or equal to the specified value. | [optional] 
 **MediaImagesHeightMax** | **int32**| This parameter is used for finding stories whose height of images are less than or equal to the specified value. | [optional] 
 **MediaImagesContentLengthMin** | **int32**| This parameter is used for finding stories whose images content length are greater than or equal to the specified value. | [optional] 
 **MediaImagesContentLengthMax** | **int32**| This parameter is used for finding stories whose images content length are less than or equal to the specified value. | [optional] 
 **MediaImagesFormat** | **[]string**| This parameter is used for finding stories whose images format are the specified value. | [optional]  [enum: BMP, GIF, JPEG, PNG, TIFF, PSD, ICO, CUR, WEBP, SVG]
 **NotMediaImagesFormat** | **[]string**| This parameter is used for excluding stories whose images format are the specified value. | [optional]  [enum: BMP, GIF, JPEG, PNG, TIFF, PSD, ICO, CUR, WEBP, SVG]
 **MediaVideosCountMin** | **int32**| This parameter is used for finding stories whose number of videos is greater than or equal to the specified value. | [optional] 
 **MediaVideosCountMax** | **int32**| This parameter is used for finding stories whose number of videos is less than or equal to the specified value. | [optional] 
 **AuthorId** | **[]int32**| This parameter is used for finding stories whose author id is the specified value. | [optional] 
 **NotAuthorId** | **[]int32**| This parameter is used for excluding stories whose author id is the specified value. | [optional] 
 **AuthorName** | **string**| This parameter is used for finding stories whose author full name contains the specified value. | [optional] 
 **NotAuthorName** | **string**| This parameter is used for excluding stories whose author full name contains the specified value. | [optional] 
 **SourceId** | **[]int32**| This parameter is used for finding stories whose source id is the specified value. | [optional] 
 **NotSourceId** | **[]int32**| This parameter is used for excluding stories whose source id is the specified value. | [optional] 
 **SourceName** | **[]string**| This parameter is used for finding stories whose source name contains the specified value. | [optional] 
 **NotSourceName** | **[]string**| This parameter is used for excluding stories whose source name contains the specified value. | [optional] 
 **SourceDomain** | **[]string**| This parameter is used for finding stories whose source domain is the specified value. | [optional] 
 **NotSourceDomain** | **[]string**| This parameter is used for excluding stories whose source domain is the specified value. | [optional] 
 **SourceLocationsCountry** | **[]string**| This parameter is used for finding stories whose source country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **NotSourceLocationsCountry** | **[]string**| This parameter is used for excluding stories whose source country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceLocationsState** | **[]string**| This parameter is used for finding stories whose source state/province is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **NotSourceLocationsState** | **[]string**| This parameter is used for excluding stories whose source state/province is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceLocationsCity** | **[]string**| This parameter is used for finding stories whose source city is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **NotSourceLocationsCity** | **[]string**| This parameter is used for excluding stories whose source city is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceScopesCountry** | **[]string**| This parameter is used for finding stories whose source scopes  is the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **NotSourceScopesCountry** | **[]string**| This parameter is used for excluding stories whose source scopes  is the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceScopesState** | **[]string**| This parameter is used for finding stories whose source scopes is the specified state/province value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **NotSourceScopesState** | **[]string**| This parameter is used for excluding stories whose source scopes is the specified state/province value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceScopesCity** | **[]string**| This parameter is used for finding stories whose source scopes is the specified city value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **NotSourceScopesCity** | **[]string**| This parameter is used for excluding stories whose source scopes is the specified city value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceScopesLevel** | **[]string**| This parameter is used for finding stories whose source scopes  is the specified level value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional]  [enum: international, national, local]
 **NotSourceScopesLevel** | **[]string**| This parameter is used for excluding stories whose source scopes  is the specified level value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional]  [enum: international, national, local]
 **SourceLinksInCountMin** | **int32**| This parameter is used for finding stories from sources whose Links in count is greater than or equal to the specified value. You can read more about working with Links in count [here](https://newsapi.aylien.com/docs/working-with-links-in-count). | [optional] 
 **SourceLinksInCountMax** | **int32**| This parameter is used for finding stories from sources whose Links in count is less than or equal to the specified value. You can read more about working with Links in count [here](https://newsapi.aylien.com/docs/working-with-links-in-count). | [optional] 
 **SourceRankingsAlexaRankMin** | **int32**| This parameter is used for finding stories from sources whose Alexa rank is greater than or equal to the specified value. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
 **SourceRankingsAlexaRankMax** | **int32**| This parameter is used for finding stories from sources whose Alexa rank is less than or equal to the specified value. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
 **SourceRankingsAlexaCountry** | **[]string**| This parameter is used for finding stories from sources whose Alexa rank is in the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
 **SocialSharesCountFacebookMin** | **int32**| This parameter is used for finding stories whose Facebook social shares count is greater than or equal to the specified value. | [optional] 
 **SocialSharesCountFacebookMax** | **int32**| This parameter is used for finding stories whose Facebook social shares count is less than or equal to the specified value. | [optional] 
 **SocialSharesCountGooglePlusMin** | **int32**| This parameter is used for finding stories whose Google+ social shares count is greater than or equal to the specified value. | [optional] 
 **SocialSharesCountGooglePlusMax** | **int32**| This parameter is used for finding stories whose Google+ social shares count is less than or equal to the specified value. | [optional] 
 **SocialSharesCountLinkedinMin** | **int32**| This parameter is used for finding stories whose LinkedIn social shares count is greater than or equal to the specified value. | [optional] 
 **SocialSharesCountLinkedinMax** | **int32**| This parameter is used for finding stories whose LinkedIn social shares count is less than or equal to the specified value. | [optional] 
 **SocialSharesCountRedditMin** | **int32**| This parameter is used for finding stories whose Reddit social shares count is greater than or equal to the specified value. | [optional] 
 **SocialSharesCountRedditMax** | **int32**| This parameter is used for finding stories whose Reddit social shares count is less than or equal to the specified value. | [optional] 
 **Cluster** | **bool**| This parameter enables clustering for the returned stories. You can read more about working with clustering [here](https://newsapi.aylien.com/docs/working-with-clustering). | [optional] [default to false]
 **ClusterAlgorithm** | **string**| This parameter is used for specifying the clustering algorithm you wish to use. It supprts STC, Lingo and [k-means](https://en.wikipedia.org/wiki/K-means_clustering) algorithms. You can read more about working with clustering [here](https://newsapi.aylien.com/docs/working-with-clustering). | [optional] [default to lingo] [enum: stc, lingo, kmeans]
 **Return** | **[]string**| This parameter is used for specifying return fields. | [optional]  [enum: id, title, body, summary, source, author, entities, keywords, hashtags, characters_count, words_count, sentences_count, paragraphs_count, categories, social_shares_count, media, sentiment, language, published_at, links]
 **StoryId** | **int64**| A story id | [optional] 
 **StoryUrl** | **string**| An article or webpage | [optional] 
 **StoryTitle** | **string**| Title of the article | [optional] 
 **StoryBody** | **string**| Body of the article | [optional] 
 **BoostBy** | **string**| This parameter is used for boosting the result by the specified value. | [optional]  [enum: recency, popularity]
 **StoryLanguage** | **string**| This parameter is used for setting the language of the story. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes. | [optional] [default to auto] [enum: auto, en, de, fr, it, es, pt]
 **PerPage** | **int32**| This parameter is used for specifying number of items in each page. | [optional] [default to 3]

### Return type

[**RelatedStories**](RelatedStories.md)

### Authorization

[app_key](../README.md#app_key), [app_id](../README.md#app_id)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListStories**
> Stories ListStories(params)

List Stories

This endpoint is used for getting a list of stories.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | **[]int64**| This parameter is used for finding stories by story id. | [optional] 
 **NotId** | **[]int64**| This parameter is used for excluding stories by story id. | [optional] 
 **Title** | **string**| This parameter is used for finding stories whose title contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Body** | **string**| This parameter is used for finding stories whose body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Text** | **string**| This parameter is used for finding stories whose title or body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Language** | **[]string**| This parameter is used for finding stories whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes. | [optional]  [enum: en, de, fr, it, es, pt]
 **NotLanguage** | **[]string**| This parameter is used for excluding stories whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes. | [optional]  [enum: en, de, fr, it, es, pt]
 **PublishedAtStart** | **string**| This parameter is used for finding stories whose published at time is greater than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates). | [optional] 
 **PublishedAtEnd** | **string**| This parameter is used for finding stories whose published at time is less than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates). | [optional] 
 **CategoriesTaxonomy** | **string**| This parameter is used for defining the type of the taxonomy for the rest of the categories queries. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional]  [enum: iab-qag, iptc-subjectcode]
 **CategoriesConfident** | **bool**| This parameter is used for finding stories whose categories are confident. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] [default to true]
 **CategoriesId** | **[]string**| This parameter is used for finding stories by categories id. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] 
 **NotCategoriesId** | **[]string**| This parameter is used for excluding stories by categories id. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] 
 **CategoriesLevel** | **[]int32**| This parameter is used for finding stories by categories level. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] 
 **NotCategoriesLevel** | **[]int32**| This parameter is used for excluding stories by categories level. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] 
 **EntitiesTitleText** | **[]string**| This parameter is used to find stories based on the specified entities &#x60;text&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **NotEntitiesTitleText** | **[]string**| This parameter is used to exclude stories based on the specified entities &#x60;text&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesTitleType** | **[]string**| This parameter is used to find stories based on the specified entities &#x60;type&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **NotEntitiesTitleType** | **[]string**| This parameter is used to exclude stories based on the specified entities &#x60;type&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesTitleLinksDbpedia** | **[]string**| This parameter is used to find stories based on the specified entities dbpedia URL in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **NotEntitiesTitleLinksDbpedia** | **[]string**| This parameter is used to exclude stories based on the specified entities dbpedia URL in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesBodyText** | **[]string**| This parameter is used to find stories based on the specified entities &#x60;text&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **NotEntitiesBodyText** | **[]string**| This parameter is used to exclude stories based on the specified entities &#x60;text&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesBodyType** | **[]string**| This parameter is used to find stories based on the specified entities &#x60;type&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **NotEntitiesBodyType** | **[]string**| This parameter is used to exclude stories based on the specified entities &#x60;type&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesBodyLinksDbpedia** | **[]string**| This parameter is used to find stories based on the specified entities dbpedia URL in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **NotEntitiesBodyLinksDbpedia** | **[]string**| This parameter is used to exclude stories based on the specified entities dbpedia URL in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **SentimentTitlePolarity** | **string**| This parameter is used for finding stories whose title sentiment is the specified value. | [optional]  [enum: positive, neutral, negative]
 **NotSentimentTitlePolarity** | **string**| This parameter is used for excluding stories whose title sentiment is the specified value. | [optional]  [enum: positive, neutral, negative]
 **SentimentBodyPolarity** | **string**| This parameter is used for finding stories whose body sentiment is the specified value. | [optional]  [enum: positive, neutral, negative]
 **NotSentimentBodyPolarity** | **string**| This parameter is used for excluding stories whose body sentiment is the specified value. | [optional]  [enum: positive, neutral, negative]
 **MediaImagesCountMin** | **int32**| This parameter is used for finding stories whose number of images is greater than or equal to the specified value. | [optional] 
 **MediaImagesCountMax** | **int32**| This parameter is used for finding stories whose number of images is less than or equal to the specified value. | [optional] 
 **MediaImagesWidthMin** | **int32**| This parameter is used for finding stories whose width of images are greater than or equal to the specified value. | [optional] 
 **MediaImagesWidthMax** | **int32**| This parameter is used for finding stories whose width of images are less than or equal to the specified value. | [optional] 
 **MediaImagesHeightMin** | **int32**| This parameter is used for finding stories whose height of images are greater than or equal to the specified value. | [optional] 
 **MediaImagesHeightMax** | **int32**| This parameter is used for finding stories whose height of images are less than or equal to the specified value. | [optional] 
 **MediaImagesContentLengthMin** | **int32**| This parameter is used for finding stories whose images content length are greater than or equal to the specified value. | [optional] 
 **MediaImagesContentLengthMax** | **int32**| This parameter is used for finding stories whose images content length are less than or equal to the specified value. | [optional] 
 **MediaImagesFormat** | **[]string**| This parameter is used for finding stories whose images format are the specified value. | [optional]  [enum: BMP, GIF, JPEG, PNG, TIFF, PSD, ICO, CUR, WEBP, SVG]
 **NotMediaImagesFormat** | **[]string**| This parameter is used for excluding stories whose images format are the specified value. | [optional]  [enum: BMP, GIF, JPEG, PNG, TIFF, PSD, ICO, CUR, WEBP, SVG]
 **MediaVideosCountMin** | **int32**| This parameter is used for finding stories whose number of videos is greater than or equal to the specified value. | [optional] 
 **MediaVideosCountMax** | **int32**| This parameter is used for finding stories whose number of videos is less than or equal to the specified value. | [optional] 
 **AuthorId** | **[]int32**| This parameter is used for finding stories whose author id is the specified value. | [optional] 
 **NotAuthorId** | **[]int32**| This parameter is used for excluding stories whose author id is the specified value. | [optional] 
 **AuthorName** | **string**| This parameter is used for finding stories whose author full name contains the specified value. | [optional] 
 **NotAuthorName** | **string**| This parameter is used for excluding stories whose author full name contains the specified value. | [optional] 
 **SourceId** | **[]int32**| This parameter is used for finding stories whose source id is the specified value. | [optional] 
 **NotSourceId** | **[]int32**| This parameter is used for excluding stories whose source id is the specified value. | [optional] 
 **SourceName** | **[]string**| This parameter is used for finding stories whose source name contains the specified value. | [optional] 
 **NotSourceName** | **[]string**| This parameter is used for excluding stories whose source name contains the specified value. | [optional] 
 **SourceDomain** | **[]string**| This parameter is used for finding stories whose source domain is the specified value. | [optional] 
 **NotSourceDomain** | **[]string**| This parameter is used for excluding stories whose source domain is the specified value. | [optional] 
 **SourceLocationsCountry** | **[]string**| This parameter is used for finding stories whose source country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **NotSourceLocationsCountry** | **[]string**| This parameter is used for excluding stories whose source country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceLocationsState** | **[]string**| This parameter is used for finding stories whose source state/province is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **NotSourceLocationsState** | **[]string**| This parameter is used for excluding stories whose source state/province is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceLocationsCity** | **[]string**| This parameter is used for finding stories whose source city is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **NotSourceLocationsCity** | **[]string**| This parameter is used for excluding stories whose source city is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceScopesCountry** | **[]string**| This parameter is used for finding stories whose source scopes is the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **NotSourceScopesCountry** | **[]string**| This parameter is used for excluding stories whose source scopes is the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceScopesState** | **[]string**| This parameter is used for finding stories whose source scopes is the specified state/province value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **NotSourceScopesState** | **[]string**| This parameter is used for excluding stories whose source scopes is the specified state/province value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceScopesCity** | **[]string**| This parameter is used for finding stories whose source scopes is the specified city value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **NotSourceScopesCity** | **[]string**| This parameter is used for excluding stories whose source scopes is the specified city value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceScopesLevel** | **[]string**| This parameter is used for finding stories whose source scopes is the specified level value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional]  [enum: international, national, local]
 **NotSourceScopesLevel** | **[]string**| This parameter is used for excluding stories whose source scopes is the specified level value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional]  [enum: international, national, local]
 **SourceLinksInCountMin** | **int32**| This parameter is used for finding stories from sources whose Links in count is greater than or equal to the specified value. You can read more about working with Links in count [here](https://newsapi.aylien.com/docs/working-with-links-in-count). | [optional] 
 **SourceLinksInCountMax** | **int32**| This parameter is used for finding stories from sources whose Links in count is less than or equal to the specified value. You can read more about working with Links in count [here](https://newsapi.aylien.com/docs/working-with-links-in-count). | [optional] 
 **SourceRankingsAlexaRankMin** | **int32**| This parameter is used for finding stories from sources whose Alexa rank is greater than or equal to the specified value. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
 **SourceRankingsAlexaRankMax** | **int32**| This parameter is used for finding stories from sources whose Alexa rank is less than or equal to the specified value. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
 **SourceRankingsAlexaCountry** | **[]string**| This parameter is used for finding stories from sources whose Alexa rank is in the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
 **SocialSharesCountFacebookMin** | **int32**| This parameter is used for finding stories whose Facebook social shares count is greater than or equal to the specified value. | [optional] 
 **SocialSharesCountFacebookMax** | **int32**| This parameter is used for finding stories whose Facebook social shares count is less than or equal to the specified value. | [optional] 
 **SocialSharesCountGooglePlusMin** | **int32**| This parameter is used for finding stories whose Google+ social shares count is greater than or equal to the specified value. | [optional] 
 **SocialSharesCountGooglePlusMax** | **int32**| This parameter is used for finding stories whose Google+ social shares count is less than or equal to the specified value. | [optional] 
 **SocialSharesCountLinkedinMin** | **int32**| This parameter is used for finding stories whose LinkedIn social shares count is greater than or equal to the specified value. | [optional] 
 **SocialSharesCountLinkedinMax** | **int32**| This parameter is used for finding stories whose LinkedIn social shares count is less than or equal to the specified value. | [optional] 
 **SocialSharesCountRedditMin** | **int32**| This parameter is used for finding stories whose Reddit social shares count is greater than or equal to the specified value. | [optional] 
 **SocialSharesCountRedditMax** | **int32**| This parameter is used for finding stories whose Reddit social shares count is less than or equal to the specified value. | [optional] 
 **Cluster** | **bool**| This parameter enables clustering for the returned stories. You can read more about working with clustering [here](https://newsapi.aylien.com/docs/working-with-clustering). | [optional] [default to false]
 **ClusterAlgorithm** | **string**| This parameter is used for specifying the clustering algorithm you wish to use. It supprts STC, Lingo and [k-means](https://en.wikipedia.org/wiki/K-means_clustering) algorithms. You can read more about working with clustering [here](https://newsapi.aylien.com/docs/working-with-clustering). | [optional] [default to lingo] [enum: stc, lingo, kmeans]
 **Return** | **[]string**| This parameter is used for specifying return fields. | [optional]  [enum: id, title, body, summary, source, author, entities, keywords, hashtags, characters_count, words_count, sentences_count, paragraphs_count, categories, social_shares_count, media, sentiment, language, published_at, links]
 **SortBy** | **string**| This parameter is used for changing the order column of the results. | [optional] [default to published_at] [enum: relevance, recency, hotness, published_at, social_shares_count, social_shares_count.facebook, social_shares_count.linkedin, social_shares_count.google_plus, social_shares_count.reddit, media.images.count, media.videos.count, source.links_in_count, source.rankings.alexa.rank, source.rankings.alexa.rank.AF, source.rankings.alexa.rank.AX, source.rankings.alexa.rank.AL, source.rankings.alexa.rank.DZ, source.rankings.alexa.rank.AS, source.rankings.alexa.rank.AD, source.rankings.alexa.rank.AO, source.rankings.alexa.rank.AI, source.rankings.alexa.rank.AQ, source.rankings.alexa.rank.AG, source.rankings.alexa.rank.AR, source.rankings.alexa.rank.AM, source.rankings.alexa.rank.AW, source.rankings.alexa.rank.AU, source.rankings.alexa.rank.AT, source.rankings.alexa.rank.AZ, source.rankings.alexa.rank.BS, source.rankings.alexa.rank.BH, source.rankings.alexa.rank.BD, source.rankings.alexa.rank.BB, source.rankings.alexa.rank.BY, source.rankings.alexa.rank.BE, source.rankings.alexa.rank.BZ, source.rankings.alexa.rank.BJ, source.rankings.alexa.rank.BM, source.rankings.alexa.rank.BT, source.rankings.alexa.rank.BO, source.rankings.alexa.rank.BQ, source.rankings.alexa.rank.BA, source.rankings.alexa.rank.BW, source.rankings.alexa.rank.BV, source.rankings.alexa.rank.BR, source.rankings.alexa.rank.IO, source.rankings.alexa.rank.BN, source.rankings.alexa.rank.BG, source.rankings.alexa.rank.BF, source.rankings.alexa.rank.BI, source.rankings.alexa.rank.KH, source.rankings.alexa.rank.CM, source.rankings.alexa.rank.CA, source.rankings.alexa.rank.CV, source.rankings.alexa.rank.KY, source.rankings.alexa.rank.CF, source.rankings.alexa.rank.TD, source.rankings.alexa.rank.CL, source.rankings.alexa.rank.CN, source.rankings.alexa.rank.CX, source.rankings.alexa.rank.CC, source.rankings.alexa.rank.CO, source.rankings.alexa.rank.KM, source.rankings.alexa.rank.CG, source.rankings.alexa.rank.CD, source.rankings.alexa.rank.CK, source.rankings.alexa.rank.CR, source.rankings.alexa.rank.CI, source.rankings.alexa.rank.HR, source.rankings.alexa.rank.CU, source.rankings.alexa.rank.CW, source.rankings.alexa.rank.CY, source.rankings.alexa.rank.CZ, source.rankings.alexa.rank.DK, source.rankings.alexa.rank.DJ, source.rankings.alexa.rank.DM, source.rankings.alexa.rank.DO, source.rankings.alexa.rank.EC, source.rankings.alexa.rank.EG, source.rankings.alexa.rank.SV, source.rankings.alexa.rank.GQ, source.rankings.alexa.rank.ER, source.rankings.alexa.rank.EE, source.rankings.alexa.rank.ET, source.rankings.alexa.rank.FK, source.rankings.alexa.rank.FO, source.rankings.alexa.rank.FJ, source.rankings.alexa.rank.FI, source.rankings.alexa.rank.FR, source.rankings.alexa.rank.GF, source.rankings.alexa.rank.PF, source.rankings.alexa.rank.TF, source.rankings.alexa.rank.GA, source.rankings.alexa.rank.GM, source.rankings.alexa.rank.GE, source.rankings.alexa.rank.DE, source.rankings.alexa.rank.GH, source.rankings.alexa.rank.GI, source.rankings.alexa.rank.GR, source.rankings.alexa.rank.GL, source.rankings.alexa.rank.GD, source.rankings.alexa.rank.GP, source.rankings.alexa.rank.GU, source.rankings.alexa.rank.GT, source.rankings.alexa.rank.GG, source.rankings.alexa.rank.GN, source.rankings.alexa.rank.GW, source.rankings.alexa.rank.GY, source.rankings.alexa.rank.HT, source.rankings.alexa.rank.HM, source.rankings.alexa.rank.VA, source.rankings.alexa.rank.HN, source.rankings.alexa.rank.HK, source.rankings.alexa.rank.HU, source.rankings.alexa.rank.IS, source.rankings.alexa.rank.IN, source.rankings.alexa.rank.ID, source.rankings.alexa.rank.IR, source.rankings.alexa.rank.IQ, source.rankings.alexa.rank.IE, source.rankings.alexa.rank.IM, source.rankings.alexa.rank.IL, source.rankings.alexa.rank.IT, source.rankings.alexa.rank.JM, source.rankings.alexa.rank.JP, source.rankings.alexa.rank.JE, source.rankings.alexa.rank.JO, source.rankings.alexa.rank.KZ, source.rankings.alexa.rank.KE, source.rankings.alexa.rank.KI, source.rankings.alexa.rank.KP, source.rankings.alexa.rank.KR, source.rankings.alexa.rank.KW, source.rankings.alexa.rank.KG, source.rankings.alexa.rank.LA, source.rankings.alexa.rank.LV, source.rankings.alexa.rank.LB, source.rankings.alexa.rank.LS, source.rankings.alexa.rank.LR, source.rankings.alexa.rank.LY, source.rankings.alexa.rank.LI, source.rankings.alexa.rank.LT, source.rankings.alexa.rank.LU, source.rankings.alexa.rank.MO, source.rankings.alexa.rank.MK, source.rankings.alexa.rank.MG, source.rankings.alexa.rank.MW, source.rankings.alexa.rank.MY, source.rankings.alexa.rank.MV, source.rankings.alexa.rank.ML, source.rankings.alexa.rank.MT, source.rankings.alexa.rank.MH, source.rankings.alexa.rank.MQ, source.rankings.alexa.rank.MR, source.rankings.alexa.rank.MU, source.rankings.alexa.rank.YT, source.rankings.alexa.rank.MX, source.rankings.alexa.rank.FM, source.rankings.alexa.rank.MD, source.rankings.alexa.rank.MC, source.rankings.alexa.rank.MN, source.rankings.alexa.rank.ME, source.rankings.alexa.rank.MS, source.rankings.alexa.rank.MA, source.rankings.alexa.rank.MZ, source.rankings.alexa.rank.MM, source.rankings.alexa.rank.NA, source.rankings.alexa.rank.NR, source.rankings.alexa.rank.NP, source.rankings.alexa.rank.NL, source.rankings.alexa.rank.NC, source.rankings.alexa.rank.NZ, source.rankings.alexa.rank.NI, source.rankings.alexa.rank.NE, source.rankings.alexa.rank.NG, source.rankings.alexa.rank.NU, source.rankings.alexa.rank.NF, source.rankings.alexa.rank.MP, source.rankings.alexa.rank.NO, source.rankings.alexa.rank.OM, source.rankings.alexa.rank.PK, source.rankings.alexa.rank.PW, source.rankings.alexa.rank.PS, source.rankings.alexa.rank.PA, source.rankings.alexa.rank.PG, source.rankings.alexa.rank.PY, source.rankings.alexa.rank.PE, source.rankings.alexa.rank.PH, source.rankings.alexa.rank.PN, source.rankings.alexa.rank.PL, source.rankings.alexa.rank.PT, source.rankings.alexa.rank.PR, source.rankings.alexa.rank.QA, source.rankings.alexa.rank.RE, source.rankings.alexa.rank.RO, source.rankings.alexa.rank.RU, source.rankings.alexa.rank.RW, source.rankings.alexa.rank.BL, source.rankings.alexa.rank.SH, source.rankings.alexa.rank.KN, source.rankings.alexa.rank.LC, source.rankings.alexa.rank.MF, source.rankings.alexa.rank.PM, source.rankings.alexa.rank.VC, source.rankings.alexa.rank.WS, source.rankings.alexa.rank.SM, source.rankings.alexa.rank.ST, source.rankings.alexa.rank.SA, source.rankings.alexa.rank.SN, source.rankings.alexa.rank.RS, source.rankings.alexa.rank.SC, source.rankings.alexa.rank.SL, source.rankings.alexa.rank.SG, source.rankings.alexa.rank.SX, source.rankings.alexa.rank.SK, source.rankings.alexa.rank.SI, source.rankings.alexa.rank.SB, source.rankings.alexa.rank.SO, source.rankings.alexa.rank.ZA, source.rankings.alexa.rank.GS, source.rankings.alexa.rank.SS, source.rankings.alexa.rank.ES, source.rankings.alexa.rank.LK, source.rankings.alexa.rank.SD, source.rankings.alexa.rank.SR, source.rankings.alexa.rank.SJ, source.rankings.alexa.rank.SZ, source.rankings.alexa.rank.SE, source.rankings.alexa.rank.CH, source.rankings.alexa.rank.SY, source.rankings.alexa.rank.TW, source.rankings.alexa.rank.TJ, source.rankings.alexa.rank.TZ, source.rankings.alexa.rank.TH, source.rankings.alexa.rank.TL, source.rankings.alexa.rank.TG, source.rankings.alexa.rank.TK, source.rankings.alexa.rank.TO, source.rankings.alexa.rank.TT, source.rankings.alexa.rank.TN, source.rankings.alexa.rank.TR, source.rankings.alexa.rank.TM, source.rankings.alexa.rank.TC, source.rankings.alexa.rank.TV, source.rankings.alexa.rank.UG, source.rankings.alexa.rank.UA, source.rankings.alexa.rank.AE, source.rankings.alexa.rank.GB, source.rankings.alexa.rank.US, source.rankings.alexa.rank.UM, source.rankings.alexa.rank.UY, source.rankings.alexa.rank.UZ, source.rankings.alexa.rank.VU, source.rankings.alexa.rank.VE, source.rankings.alexa.rank.VN, source.rankings.alexa.rank.VG, source.rankings.alexa.rank.VI, source.rankings.alexa.rank.WF, source.rankings.alexa.rank.EH, source.rankings.alexa.rank.YE, source.rankings.alexa.rank.ZM, source.rankings.alexa.rank.ZW]
 **SortDirection** | **string**| This parameter is used for changing the order direction of the result. | [optional] [default to desc] [enum: asc, desc]
 **Cursor** | **string**| This parameter is used for finding a specific page. You can read more about pagination of results [here](https://newsapi.aylien.com/docs/pagination-of-results). | [optional] [default to *]
 **PerPage** | **int32**| This parameter is used for specifying number of items in each page You can read more about pagination of results [here](https://newsapi.aylien.com/docs/pagination-of-results) | [optional] [default to 10]

### Return type

[**Stories**](Stories.md)

### Authorization

[app_key](../README.md#app_key), [app_id](../README.md#app_id)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTimeSeries**
> TimeSeriesList ListTimeSeries(params)

List time series

This endpoint is used for getting time series by stories.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | **[]int64**| This parameter is used for finding stories by story id. | [optional] 
 **NotId** | **[]int64**| This parameter is used for excluding stories by story id. | [optional] 
 **Title** | **string**| This parameter is used for finding stories whose title contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Body** | **string**| This parameter is used for finding stories whose body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Text** | **string**| This parameter is used for finding stories whose title or body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Language** | **[]string**| This parameter is used for finding stories whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes. | [optional]  [enum: en, de, fr, it, es, pt]
 **NotLanguage** | **[]string**| This parameter is used for excluding stories whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes. | [optional]  [enum: en, de, fr, it, es, pt]
 **CategoriesTaxonomy** | **string**| This parameter is used for defining the type of the taxonomy for the rest of the categories queries. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional]  [enum: iab-qag, iptc-subjectcode]
 **CategoriesConfident** | **bool**| This parameter is used for finding stories whose categories are confident. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] [default to true]
 **CategoriesId** | **[]string**| This parameter is used for finding stories by categories id. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] 
 **NotCategoriesId** | **[]string**| This parameter is used for excluding stories by categories id. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] 
 **CategoriesLevel** | **[]int32**| This parameter is used for finding stories by categories level. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] 
 **NotCategoriesLevel** | **[]int32**| This parameter is used for excluding stories by categories level. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] 
 **EntitiesTitleText** | **[]string**| This parameter is used to find stories based on the specified entities &#x60;text&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **NotEntitiesTitleText** | **[]string**| This parameter is used to exclude stories based on the specified entities &#x60;text&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesTitleType** | **[]string**| This parameter is used to find stories based on the specified entities &#x60;type&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **NotEntitiesTitleType** | **[]string**| This parameter is used to exclude stories based on the specified entities &#x60;type&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesTitleLinksDbpedia** | **[]string**| This parameter is used to find stories based on the specified entities dbpedia URL in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **NotEntitiesTitleLinksDbpedia** | **[]string**| This parameter is used to exclude stories based on the specified entities dbpedia URL in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesBodyText** | **[]string**| This parameter is used to find stories based on the specified entities &#x60;text&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **NotEntitiesBodyText** | **[]string**| This parameter is used to exclude stories based on the specified entities &#x60;text&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesBodyType** | **[]string**| This parameter is used to find stories based on the specified entities &#x60;type&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **NotEntitiesBodyType** | **[]string**| This parameter is used to exclude stories based on the specified entities &#x60;type&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesBodyLinksDbpedia** | **[]string**| This parameter is used to find stories based on the specified entities dbpedia URL in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **NotEntitiesBodyLinksDbpedia** | **[]string**| This parameter is used to exclude stories based on the specified entities dbpedia URL in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **SentimentTitlePolarity** | **string**| This parameter is used for finding stories whose title sentiment is the specified value. | [optional]  [enum: positive, neutral, negative]
 **NotSentimentTitlePolarity** | **string**| This parameter is used for excluding stories whose title sentiment is the specified value. | [optional]  [enum: positive, neutral, negative]
 **SentimentBodyPolarity** | **string**| This parameter is used for finding stories whose body sentiment is the specified value. | [optional]  [enum: positive, neutral, negative]
 **NotSentimentBodyPolarity** | **string**| This parameter is used for excluding stories whose body sentiment is the specified value. | [optional]  [enum: positive, neutral, negative]
 **MediaImagesCountMin** | **int32**| This parameter is used for finding stories whose number of images is greater than or equal to the specified value. | [optional] 
 **MediaImagesCountMax** | **int32**| This parameter is used for finding stories whose number of images is less than or equal to the specified value. | [optional] 
 **MediaImagesWidthMin** | **int32**| This parameter is used for finding stories whose width of images are greater than or equal to the specified value. | [optional] 
 **MediaImagesWidthMax** | **int32**| This parameter is used for finding stories whose width of images are less than or equal to the specified value. | [optional] 
 **MediaImagesHeightMin** | **int32**| This parameter is used for finding stories whose height of images are greater than or equal to the specified value. | [optional] 
 **MediaImagesHeightMax** | **int32**| This parameter is used for finding stories whose height of images are less than or equal to the specified value. | [optional] 
 **MediaImagesContentLengthMin** | **int32**| This parameter is used for finding stories whose images content length are greater than or equal to the specified value. | [optional] 
 **MediaImagesContentLengthMax** | **int32**| This parameter is used for finding stories whose images content length are less than or equal to the specified value. | [optional] 
 **MediaImagesFormat** | **[]string**| This parameter is used for finding stories whose images format are the specified value. | [optional]  [enum: BMP, GIF, JPEG, PNG, TIFF, PSD, ICO, CUR, WEBP, SVG]
 **NotMediaImagesFormat** | **[]string**| This parameter is used for excluding stories whose images format are the specified value. | [optional]  [enum: BMP, GIF, JPEG, PNG, TIFF, PSD, ICO, CUR, WEBP, SVG]
 **MediaVideosCountMin** | **int32**| This parameter is used for finding stories whose number of videos is greater than or equal to the specified value. | [optional] 
 **MediaVideosCountMax** | **int32**| This parameter is used for finding stories whose number of videos is less than or equal to the specified value. | [optional] 
 **AuthorId** | **[]int32**| This parameter is used for finding stories whose author id is the specified value. | [optional] 
 **NotAuthorId** | **[]int32**| This parameter is used for excluding stories whose author id is the specified value. | [optional] 
 **AuthorName** | **string**| This parameter is used for finding stories whose author full name contains the specified value. | [optional] 
 **NotAuthorName** | **string**| This parameter is used for excluding stories whose author full name contains the specified value. | [optional] 
 **SourceId** | **[]int32**| This parameter is used for finding stories whose source id is the specified value. | [optional] 
 **NotSourceId** | **[]int32**| This parameter is used for excluding stories whose source id is the specified value. | [optional] 
 **SourceName** | **[]string**| This parameter is used for finding stories whose source name contains the specified value. | [optional] 
 **NotSourceName** | **[]string**| This parameter is used for excluding stories whose source name contains the specified value. | [optional] 
 **SourceDomain** | **[]string**| This parameter is used for finding stories whose source domain is the specified value. | [optional] 
 **NotSourceDomain** | **[]string**| This parameter is used for excluding stories whose source domain is the specified value. | [optional] 
 **SourceLocationsCountry** | **[]string**| This parameter is used for finding stories whose source country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **NotSourceLocationsCountry** | **[]string**| This parameter is used for excluding stories whose source country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceLocationsState** | **[]string**| This parameter is used for finding stories whose source state/province is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **NotSourceLocationsState** | **[]string**| This parameter is used for excluding stories whose source state/province is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceLocationsCity** | **[]string**| This parameter is used for finding stories whose source city is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **NotSourceLocationsCity** | **[]string**| This parameter is used for excluding stories whose source city is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceScopesCountry** | **[]string**| This parameter is used for finding stories whose source scopes is the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **NotSourceScopesCountry** | **[]string**| This parameter is used for excluding stories whose source scopes is the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceScopesState** | **[]string**| This parameter is used for finding stories whose source scopes is the specified state/province value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **NotSourceScopesState** | **[]string**| This parameter is used for excluding stories whose source scopes is the specified state/province value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceScopesCity** | **[]string**| This parameter is used for finding stories whose source scopes is the specified city value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **NotSourceScopesCity** | **[]string**| This parameter is used for excluding stories whose source scopes is the specified city value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceScopesLevel** | **[]string**| This parameter is used for finding stories whose source scopes is the specified level value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional]  [enum: international, national, local]
 **NotSourceScopesLevel** | **[]string**| This parameter is used for excluding stories whose source scopes is the specified level value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional]  [enum: international, national, local]
 **SourceLinksInCountMin** | **int32**| This parameter is used for finding stories from sources whose Links in count is greater than or equal to the specified value. You can read more about working with Links in count [here](https://newsapi.aylien.com/docs/working-with-links-in-count). | [optional] 
 **SourceLinksInCountMax** | **int32**| This parameter is used for finding stories from sources whose Links in count is less than or equal to the specified value. You can read more about working with Links in count [here](https://newsapi.aylien.com/docs/working-with-links-in-count). | [optional] 
 **SourceRankingsAlexaRankMin** | **int32**| This parameter is used for finding stories from sources whose Alexa rank is greater than or equal to the specified value. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
 **SourceRankingsAlexaRankMax** | **int32**| This parameter is used for finding stories from sources whose Alexa rank is less than or equal to the specified value. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
 **SourceRankingsAlexaCountry** | **[]string**| This parameter is used for finding stories from sources whose Alexa rank is in the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
 **SocialSharesCountFacebookMin** | **int32**| This parameter is used for finding stories whose Facebook social shares count is greater than or equal to the specified value. | [optional] 
 **SocialSharesCountFacebookMax** | **int32**| This parameter is used for finding stories whose Facebook social shares count is less than or equal to the specified value. | [optional] 
 **SocialSharesCountGooglePlusMin** | **int32**| This parameter is used for finding stories whose Google+ social shares count is greater than or equal to the specified value. | [optional] 
 **SocialSharesCountGooglePlusMax** | **int32**| This parameter is used for finding stories whose Google+ social shares count is less than or equal to the specified value. | [optional] 
 **SocialSharesCountLinkedinMin** | **int32**| This parameter is used for finding stories whose LinkedIn social shares count is greater than or equal to the specified value. | [optional] 
 **SocialSharesCountLinkedinMax** | **int32**| This parameter is used for finding stories whose LinkedIn social shares count is less than or equal to the specified value. | [optional] 
 **SocialSharesCountRedditMin** | **int32**| This parameter is used for finding stories whose Reddit social shares count is greater than or equal to the specified value. | [optional] 
 **SocialSharesCountRedditMax** | **int32**| This parameter is used for finding stories whose Reddit social shares count is less than or equal to the specified value. | [optional] 
 **PublishedAtStart** | **string**| This parameter is used for finding stories whose published at time is less than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates). | [optional] [default to NOW-7DAYS/DAY]
 **PublishedAtEnd** | **string**| This parameter is used for finding stories whose published at time is greater than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates). | [optional] [default to NOW/DAY]
 **Period** | **string**| The size of each date range is expressed as an interval to be added to the lower bound. It supports Date Math Syntax. Valid options are &#x60;+&#x60; following an integer number greater than 0 and one of the Date Math keywords. e.g. &#x60;+1DAY&#x60;, &#x60;+2MINUTES&#x60; and &#x60;+1MONTH&#x60;. Here are [Supported keywords](https://newsapi.aylien.com/docs/working-with-dates#date-math). | [optional] [default to +1DAY]

### Return type

[**TimeSeriesList**](TimeSeriesList.md)

### Authorization

[app_key](../README.md#app_key), [app_id](../README.md#app_id)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTrends**
> Trends ListTrends(params)

List trends

This endpoint is used for finding trends based on stories.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Field** | **string**| This parameter is used to specify the trend field. |  [enum: author.name, source.name, source.domain, keywords, entities.title.text, entities.title.type, entities.title.links.dbpedia, entities.body.text, entities.body.type, entities.body.links.dbpedia, hashtags, categories.id, sentiment.title.polarity, sentiment.body.polarity]
 **Id** | **[]int64**| This parameter is used for finding stories by story id. | [optional] 
 **NotId** | **[]int64**| This parameter is used for excluding stories by story id. | [optional] 
 **Title** | **string**| This parameter is used for finding stories whose title contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Body** | **string**| This parameter is used for finding stories whose body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Text** | **string**| This parameter is used for finding stories whose title or body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Language** | **[]string**| This parameter is used for finding stories whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes. | [optional]  [enum: en, de, fr, it, es, pt]
 **NotLanguage** | **[]string**| This parameter is used for excluding stories whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes. | [optional]  [enum: en, de, fr, it, es, pt]
 **PublishedAtStart** | **string**| This parameter is used for finding stories whose published at time is greater than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates). | [optional] 
 **PublishedAtEnd** | **string**| This parameter is used for finding stories whose published at time is less than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates). | [optional] 
 **CategoriesTaxonomy** | **string**| This parameter is used for defining the type of the taxonomy for the rest of the categories queries. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional]  [enum: iab-qag, iptc-subjectcode]
 **CategoriesConfident** | **bool**| This parameter is used for finding stories whose categories are confident. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] [default to true]
 **CategoriesId** | **[]string**| This parameter is used for finding stories by categories id. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] 
 **NotCategoriesId** | **[]string**| This parameter is used for excluding stories by categories id. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] 
 **CategoriesLevel** | **[]int32**| This parameter is used for finding stories by categories level. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] 
 **NotCategoriesLevel** | **[]int32**| This parameter is used for excluding stories by categories level. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories). | [optional] 
 **EntitiesTitleText** | **[]string**| This parameter is used to find stories based on the specified entities &#x60;text&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **NotEntitiesTitleText** | **[]string**| This parameter is used to exclude stories based on the specified entities &#x60;text&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesTitleType** | **[]string**| This parameter is used to find stories based on the specified entities &#x60;type&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **NotEntitiesTitleType** | **[]string**| This parameter is used to exclude stories based on the specified entities &#x60;type&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesTitleLinksDbpedia** | **[]string**| This parameter is used to find stories based on the specified entities dbpedia URL in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **NotEntitiesTitleLinksDbpedia** | **[]string**| This parameter is used to exclude stories based on the specified entities dbpedia URL in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesBodyText** | **[]string**| This parameter is used to find stories based on the specified entities &#x60;text&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **NotEntitiesBodyText** | **[]string**| This parameter is used to exclude stories based on the specified entities &#x60;text&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesBodyType** | **[]string**| This parameter is used to find stories based on the specified entities &#x60;type&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **NotEntitiesBodyType** | **[]string**| This parameter is used to exclude stories based on the specified entities &#x60;type&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **EntitiesBodyLinksDbpedia** | **[]string**| This parameter is used to find stories based on the specified entities dbpedia URL in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **NotEntitiesBodyLinksDbpedia** | **[]string**| This parameter is used to exclude stories based on the specified entities dbpedia URL in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities). | [optional] 
 **SentimentTitlePolarity** | **string**| This parameter is used for finding stories whose title sentiment is the specified value. | [optional]  [enum: positive, neutral, negative]
 **NotSentimentTitlePolarity** | **string**| This parameter is used for excluding stories whose title sentiment is the specified value. | [optional]  [enum: positive, neutral, negative]
 **SentimentBodyPolarity** | **string**| This parameter is used for finding stories whose body sentiment is the specified value. | [optional]  [enum: positive, neutral, negative]
 **NotSentimentBodyPolarity** | **string**| This parameter is used for excluding stories whose body sentiment is the specified value. | [optional]  [enum: positive, neutral, negative]
 **MediaImagesCountMin** | **int32**| This parameter is used for finding stories whose number of images is greater than or equal to the specified value. | [optional] 
 **MediaImagesCountMax** | **int32**| This parameter is used for finding stories whose number of images is less than or equal to the specified value. | [optional] 
 **MediaImagesWidthMin** | **int32**| This parameter is used for finding stories whose width of images are greater than or equal to the specified value. | [optional] 
 **MediaImagesWidthMax** | **int32**| This parameter is used for finding stories whose width of images are less than or equal to the specified value. | [optional] 
 **MediaImagesHeightMin** | **int32**| This parameter is used for finding stories whose height of images are greater than or equal to the specified value. | [optional] 
 **MediaImagesHeightMax** | **int32**| This parameter is used for finding stories whose height of images are less than or equal to the specified value. | [optional] 
 **MediaImagesContentLengthMin** | **int32**| This parameter is used for finding stories whose images content length are greater than or equal to the specified value. | [optional] 
 **MediaImagesContentLengthMax** | **int32**| This parameter is used for finding stories whose images content length are less than or equal to the specified value. | [optional] 
 **MediaImagesFormat** | **[]string**| This parameter is used for finding stories whose images format are the specified value. | [optional]  [enum: BMP, GIF, JPEG, PNG, TIFF, PSD, ICO, CUR, WEBP, SVG]
 **NotMediaImagesFormat** | **[]string**| This parameter is used for excluding stories whose images format are the specified value. | [optional]  [enum: BMP, GIF, JPEG, PNG, TIFF, PSD, ICO, CUR, WEBP, SVG]
 **MediaVideosCountMin** | **int32**| This parameter is used for finding stories whose number of videos is greater than or equal to the specified value. | [optional] 
 **MediaVideosCountMax** | **int32**| This parameter is used for finding stories whose number of videos is less than or equal to the specified value. | [optional] 
 **AuthorId** | **[]int32**| This parameter is used for finding stories whose author id is the specified value. | [optional] 
 **NotAuthorId** | **[]int32**| This parameter is used for excluding stories whose author id is the specified value. | [optional] 
 **AuthorName** | **string**| This parameter is used for finding stories whose author full name contains the specified value. | [optional] 
 **NotAuthorName** | **string**| This parameter is used for excluding stories whose author full name contains the specified value. | [optional] 
 **SourceId** | **[]int32**| This parameter is used for finding stories whose source id is the specified value. | [optional] 
 **NotSourceId** | **[]int32**| This parameter is used for excluding stories whose source id is the specified value. | [optional] 
 **SourceName** | **[]string**| This parameter is used for finding stories whose source name contains the specified value. | [optional] 
 **NotSourceName** | **[]string**| This parameter is used for excluding stories whose source name contains the specified value. | [optional] 
 **SourceDomain** | **[]string**| This parameter is used for finding stories whose source domain is the specified value. | [optional] 
 **NotSourceDomain** | **[]string**| This parameter is used for excluding stories whose source domain is the specified value. | [optional] 
 **SourceLocationsCountry** | **[]string**| This parameter is used for finding stories whose source country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **NotSourceLocationsCountry** | **[]string**| This parameter is used for excluding stories whose source country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceLocationsState** | **[]string**| This parameter is used for finding stories whose source state/province is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **NotSourceLocationsState** | **[]string**| This parameter is used for excluding stories whose source state/province is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceLocationsCity** | **[]string**| This parameter is used for finding stories whose source city is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **NotSourceLocationsCity** | **[]string**| This parameter is used for excluding stories whose source city is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceScopesCountry** | **[]string**| This parameter is used for finding stories whose source scopes is the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **NotSourceScopesCountry** | **[]string**| This parameter is used for excluding stories whose source scopes is the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceScopesState** | **[]string**| This parameter is used for finding stories whose source scopes is the specified state/province value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **NotSourceScopesState** | **[]string**| This parameter is used for excluding stories whose source scopes is the specified state/province value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceScopesCity** | **[]string**| This parameter is used for finding stories whose source scopes is the specified city value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **NotSourceScopesCity** | **[]string**| This parameter is used for excluding stories whose source scopes is the specified city value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional] 
 **SourceScopesLevel** | **[]string**| This parameter is used for finding stories whose source scopes is the specified level value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional]  [enum: international, national, local]
 **NotSourceScopesLevel** | **[]string**| This parameter is used for excluding stories whose source scopes is the specified level value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations). | [optional]  [enum: international, national, local]
 **SourceLinksInCountMin** | **int32**| This parameter is used for finding stories from sources whose Links in count is greater than or equal to the specified value. You can read more about working with Links in count [here](https://newsapi.aylien.com/docs/working-with-links-in-count). | [optional] 
 **SourceLinksInCountMax** | **int32**| This parameter is used for finding stories from sources whose Links in count is less than or equal to the specified value. You can read more about working with Links in count [here](https://newsapi.aylien.com/docs/working-with-links-in-count). | [optional] 
 **SourceRankingsAlexaRankMin** | **int32**| This parameter is used for finding stories from sources whose Alexa rank is greater than or equal to the specified value. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
 **SourceRankingsAlexaRankMax** | **int32**| This parameter is used for finding stories from sources whose Alexa rank is less than or equal to the specified value. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
 **SourceRankingsAlexaCountry** | **[]string**| This parameter is used for finding stories from sources whose Alexa rank is in the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
 **SocialSharesCountFacebookMin** | **int32**| This parameter is used for finding stories whose Facebook social shares count is greater than or equal to the specified value. | [optional] 
 **SocialSharesCountFacebookMax** | **int32**| This parameter is used for finding stories whose Facebook social shares count is less than or equal to the specified value. | [optional] 
 **SocialSharesCountGooglePlusMin** | **int32**| This parameter is used for finding stories whose Google+ social shares count is greater than or equal to the specified value. | [optional] 
 **SocialSharesCountGooglePlusMax** | **int32**| This parameter is used for finding stories whose Google+ social shares count is less than or equal to the specified value. | [optional] 
 **SocialSharesCountLinkedinMin** | **int32**| This parameter is used for finding stories whose LinkedIn social shares count is greater than or equal to the specified value. | [optional] 
 **SocialSharesCountLinkedinMax** | **int32**| This parameter is used for finding stories whose LinkedIn social shares count is less than or equal to the specified value. | [optional] 
 **SocialSharesCountRedditMin** | **int32**| This parameter is used for finding stories whose Reddit social shares count is greater than or equal to the specified value. | [optional] 
 **SocialSharesCountRedditMax** | **int32**| This parameter is used for finding stories whose Reddit social shares count is less than or equal to the specified value. | [optional] 

### Return type

[**Trends**](Trends.md)

### Authorization

[app_key](../README.md#app_key), [app_id](../README.md#app_id)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

