# \CoverageApi

All URIs are relative to *https://api.aylien.com/news*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListCoveragesGet**](CoverageApi.md#ListCoveragesGet) | **Get** /coverages | 
[**ListCoveragesPost**](CoverageApi.md#ListCoveragesPost) | **Post** /coverages | 



## ListCoveragesGet

> Coverages ListCoveragesGet(ctx, optional)


### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListCoveragesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCoveragesGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | [**optional.Interface of []int64**](int64.md)| This parameter is used for finding stories by story id.  | 
 **id2** | [**optional.Interface of []int64**](int64.md)| This parameter is used for excluding stories by story id.  | 
 **title** | **optional.String**| This parameter is used for finding stories whose title contains a specific keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).  | 
 **body** | **optional.String**| This parameter is used for finding stories whose body contains a specific keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).  | 
 **text** | **optional.String**| This parameter is used for finding stories whose title or body contains a specific keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).  | 
 **translationsEnTitle** | **optional.String**| This parameter is used for finding stories whose translation title contains a specific keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).  | 
 **translationsEnBody** | **optional.String**| This parameter is used for finding stories whose translation body contains a specific keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).  | 
 **translationsEnText** | **optional.String**| This parameter is used for finding stories whose translation title or body contains a specific keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).  | 
 **linksPermalink** | [**optional.Interface of []string**](string.md)| This parameter is used to find stories based on their url.  | 
 **linksPermalink2** | [**optional.Interface of []string**](string.md)| This parameter is used to exclude stories based on their url.  | 
 **language** | [**optional.Interface of []string**](string.md)| This parameter is used for finding stories whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes.  | 
 **language2** | [**optional.Interface of []string**](string.md)| This parameter is used for excluding stories whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes.  | 
 **publishedAtStart** | **optional.String**| This parameter is used for finding stories whose published at time is greater than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates).  | 
 **publishedAtEnd** | **optional.String**| This parameter is used for finding stories whose published at time is less than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates).  | 
 **categoriesTaxonomy** | **optional.String**| This parameter is used for defining the type of the taxonomy for the rest of the categories queries. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories).  | 
 **categoriesConfident** | **optional.Bool**| This parameter is used for finding stories whose categories are confident. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories).  | [default to true]
 **categoriesId** | [**optional.Interface of []string**](string.md)| This parameter is used for finding stories by categories id. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories).  | 
 **categoriesId2** | [**optional.Interface of []string**](string.md)| This parameter is used for excluding stories by categories id. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories).  | 
 **categoriesLabel** | [**optional.Interface of []string**](string.md)| This parameter is used for finding stories by categories label. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories).  | 
 **categoriesLabel2** | [**optional.Interface of []string**](string.md)| This parameter is used for excluding stories by categories label. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories).  | 
 **categoriesLevel** | [**optional.Interface of []int32**](int32.md)| This parameter is used for finding stories by categories level. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories).  | 
 **categoriesLevel2** | [**optional.Interface of []int32**](int32.md)| This parameter is used for excluding stories by categories level. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories).  | 
 **entitiesTitleText** | [**optional.Interface of []string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;text&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).  | 
 **entitiesTitleText2** | [**optional.Interface of []string**](string.md)| This parameter is used to exclude stories based on the specified entities &#x60;text&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).  | 
 **entitiesTitleType** | [**optional.Interface of []string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;type&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).  | 
 **entitiesTitleType2** | [**optional.Interface of []string**](string.md)| This parameter is used to exclude stories based on the specified entities &#x60;type&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).  | 
 **entitiesTitleLinksDbpedia** | [**optional.Interface of []string**](string.md)| This parameter is used to find stories based on the specified entities dbpedia URL in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).  | 
 **entitiesTitleLinksDbpedia2** | [**optional.Interface of []string**](string.md)| This parameter is used to exclude stories based on the specified entities dbpedia URL in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).  | 
 **entitiesBodyText** | [**optional.Interface of []string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;text&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).  | 
 **entitiesBodyText2** | [**optional.Interface of []string**](string.md)| This parameter is used to exclude stories based on the specified entities &#x60;text&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).  | 
 **entitiesBodyType** | [**optional.Interface of []string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;type&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).  | 
 **entitiesBodyType2** | [**optional.Interface of []string**](string.md)| This parameter is used to exclude stories based on the specified entities &#x60;type&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).  | 
 **entitiesBodyLinksDbpedia** | [**optional.Interface of []string**](string.md)| This parameter is used to find stories based on the specified entities dbpedia URL in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).  | 
 **entitiesBodyLinksDbpedia2** | [**optional.Interface of []string**](string.md)| This parameter is used to exclude stories based on the specified entities dbpedia URL in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).  | 
 **sentimentTitlePolarity** | **optional.String**| This parameter is used for finding stories whose title sentiment is the specified value.  | 
 **sentimentTitlePolarity2** | **optional.String**| This parameter is used for excluding stories whose title sentiment is the specified value.  | 
 **sentimentBodyPolarity** | **optional.String**| This parameter is used for finding stories whose body sentiment is the specified value.  | 
 **sentimentBodyPolarity2** | **optional.String**| This parameter is used for excluding stories whose body sentiment is the specified value.  | 
 **mediaImagesCountMin** | **optional.Int32**| This parameter is used for finding stories whose number of images is greater than or equal to the specified value.  | 
 **mediaImagesCountMax** | **optional.Int32**| This parameter is used for finding stories whose number of images is less than or equal to the specified value.  | 
 **mediaImagesWidthMin** | **optional.Int32**| This parameter is used for finding stories whose width of images are greater than or equal to the specified value.  | 
 **mediaImagesWidthMax** | **optional.Int32**| This parameter is used for finding stories whose width of images are less than or equal to the specified value.  | 
 **mediaImagesHeightMin** | **optional.Int32**| This parameter is used for finding stories whose height of images are greater than or equal to the specified value.  | 
 **mediaImagesHeightMax** | **optional.Int32**| This parameter is used for finding stories whose height of images are less than or equal to the specified value.  | 
 **mediaImagesContentLengthMin** | **optional.Int32**| This parameter is used for finding stories whose images content length are greater than or equal to the specified value.  | 
 **mediaImagesContentLengthMax** | **optional.Int32**| This parameter is used for finding stories whose images content length are less than or equal to the specified value.  | 
 **mediaImagesFormat** | [**optional.Interface of []string**](string.md)| This parameter is used for finding stories whose images format are the specified value.  | 
 **mediaImagesFormat2** | [**optional.Interface of []string**](string.md)| This parameter is used for excluding stories whose images format are the specified value.  | 
 **mediaVideosCountMin** | **optional.Int32**| This parameter is used for finding stories whose number of videos is greater than or equal to the specified value.  | 
 **mediaVideosCountMax** | **optional.Int32**| This parameter is used for finding stories whose number of videos is less than or equal to the specified value.  | 
 **authorId** | [**optional.Interface of []int32**](int32.md)| This parameter is used for finding stories whose author id is the specified value.  | 
 **authorId2** | [**optional.Interface of []int32**](int32.md)| This parameter is used for excluding stories whose author id is the specified value.  | 
 **authorName** | **optional.String**| This parameter is used for finding stories whose author full name contains the specified value.  | 
 **authorName2** | **optional.String**| This parameter is used for excluding stories whose author full name contains the specified value.  | 
 **sourceId** | [**optional.Interface of []int32**](int32.md)| This parameter is used for finding stories whose source id is the specified value.  | 
 **sourceId2** | [**optional.Interface of []int32**](int32.md)| This parameter is used for excluding stories whose source id is the specified value.  | 
 **sourceName** | [**optional.Interface of []string**](string.md)| This parameter is used for finding stories whose source name contains the specified value.  | 
 **sourceName2** | [**optional.Interface of []string**](string.md)| This parameter is used for excluding stories whose source name contains the specified value.  | 
 **sourceDomain** | [**optional.Interface of []string**](string.md)| This parameter is used for finding stories whose source domain is the specified value.  | 
 **sourceDomain2** | [**optional.Interface of []string**](string.md)| This parameter is used for excluding stories whose source domain is the specified value.  | 
 **sourceLocationsCountry** | [**optional.Interface of []string**](string.md)| This parameter is used for finding stories whose source country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).  | 
 **sourceLocationsCountry2** | [**optional.Interface of []string**](string.md)| This parameter is used for excluding stories whose source country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).  | 
 **sourceLocationsState** | [**optional.Interface of []string**](string.md)| This parameter is used for finding stories whose source state/province is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).  | 
 **sourceLocationsState2** | [**optional.Interface of []string**](string.md)| This parameter is used for excluding stories whose source state/province is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).  | 
 **sourceLocationsCity** | [**optional.Interface of []string**](string.md)| This parameter is used for finding stories whose source city is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).  | 
 **sourceLocationsCity2** | [**optional.Interface of []string**](string.md)| This parameter is used for excluding stories whose source city is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).  | 
 **sourceScopesCountry** | [**optional.Interface of []string**](string.md)| This parameter is used for finding stories whose source scopes is the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).  | 
 **sourceScopesCountry2** | [**optional.Interface of []string**](string.md)| This parameter is used for excluding stories whose source scopes is the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).  | 
 **sourceScopesState** | [**optional.Interface of []string**](string.md)| This parameter is used for finding stories whose source scopes is the specified state/province value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).  | 
 **sourceScopesState2** | [**optional.Interface of []string**](string.md)| This parameter is used for excluding stories whose source scopes is the specified state/province value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).  | 
 **sourceScopesCity** | [**optional.Interface of []string**](string.md)| This parameter is used for finding stories whose source scopes is the specified city value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).  | 
 **sourceScopesCity2** | [**optional.Interface of []string**](string.md)| This parameter is used for excluding stories whose source scopes is the specified city value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).  | 
 **sourceScopesLevel** | [**optional.Interface of []string**](string.md)| This parameter is used for finding stories whose source scopes is the specified level value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).  | 
 **sourceScopesLevel2** | [**optional.Interface of []string**](string.md)| This parameter is used for excluding stories whose source scopes is the specified level value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).  | 
 **sourceLinksInCountMin** | **optional.Int32**| This parameter is used for finding stories from sources whose Links in count is greater than or equal to the specified value. You can read more about working with Links in count [here](https://newsapi.aylien.com/docs/working-with-links-in-count).  | 
 **sourceLinksInCountMax** | **optional.Int32**| This parameter is used for finding stories from sources whose Links in count is less than or equal to the specified value. You can read more about working with Links in count [here](https://newsapi.aylien.com/docs/working-with-links-in-count).  | 
 **sourceRankingsAlexaRankMin** | **optional.Int32**| This parameter is used for finding stories from sources whose Alexa rank is greater than or equal to the specified value. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks).  | 
 **sourceRankingsAlexaRankMax** | **optional.Int32**| This parameter is used for finding stories from sources whose Alexa rank is less than or equal to the specified value. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks).  | 
 **sourceRankingsAlexaCountry** | [**optional.Interface of []string**](string.md)| This parameter is used for finding stories from sources whose Alexa rank is in the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks).  | 
 **socialSharesCountFacebookMin** | **optional.Int32**| This parameter is used for finding stories whose Facebook social shares count is greater than or equal to the specified value.  | 
 **socialSharesCountFacebookMax** | **optional.Int32**| This parameter is used for finding stories whose Facebook social shares count is less than or equal to the specified value.  | 
 **socialSharesCountGooglePlusMin** | **optional.Int32**| This parameter is used for finding stories whose Google+ social shares count is greater than or equal to the specified value.  | 
 **socialSharesCountGooglePlusMax** | **optional.Int32**| This parameter is used for finding stories whose Google+ social shares count is less than or equal to the specified value.  | 
 **socialSharesCountLinkedinMin** | **optional.Int32**| This parameter is used for finding stories whose LinkedIn social shares count is greater than or equal to the specified value.  | 
 **socialSharesCountLinkedinMax** | **optional.Int32**| This parameter is used for finding stories whose LinkedIn social shares count is less than or equal to the specified value.  | 
 **socialSharesCountRedditMin** | **optional.Int32**| This parameter is used for finding stories whose Reddit social shares count is greater than or equal to the specified value.  | 
 **socialSharesCountRedditMax** | **optional.Int32**| This parameter is used for finding stories whose Reddit social shares count is less than or equal to the specified value.  | 
 **clusters** | [**optional.Interface of []string**](string.md)| This parameter is used for finding stories with belonging to one of clusters in a specific set of clusters You can read more about working with clustering [here](https://newsapi.aylien.com/docs/working-with-clustering).  | 
 **return_** | [**optional.Interface of []string**](string.md)| This parameter is used for specifying return fields. | 
 **storyId** | **optional.Int64**| A story id | 
 **storyUrl** | **optional.String**| An article or webpage | 
 **storyTitle** | **optional.String**| Title of the article | 
 **storyBody** | **optional.String**| Body of the article | 
 **storyPublishedAt** | **optional.Time**| Publish date of the article. If you use a url or title and body of an article for getting coverages, this parameter is required. The format used is a restricted form of the canonical representation of dateTime in the [XML Schema specification (ISO 8601)](https://www.w3.org/TR/xmlschema-2/#dateTime).  | 
 **storyLanguage** | **optional.String**| This parameter is used for setting the language of the story. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes.  | [default to auto]
 **perPage** | **optional.Int32**| This parameter is used for specifying number of items in each page.  | [default to 3]

### Return type

[**Coverages**](Coverages.md)

### Authorization

[app_id](../README.md#app_id), [app_key](../README.md#app_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCoveragesPost

> Coverages ListCoveragesPost(ctx, optional)


### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListCoveragesPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCoveragesPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | [**optional.Interface of []int64**](int64.md)| This parameter is used for finding stories by story id.  | 
 **id2** | [**optional.Interface of []int64**](int64.md)| This parameter is used for excluding stories by story id.  | 
 **title** | **optional.String**| This parameter is used for finding stories whose title contains a specific keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).  | 
 **body** | **optional.String**| This parameter is used for finding stories whose body contains a specific keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).  | 
 **text** | **optional.String**| This parameter is used for finding stories whose title or body contains a specific keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).  | 
 **translationsEnTitle** | **optional.String**| This parameter is used for finding stories whose translation title contains a specific keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).  | 
 **translationsEnBody** | **optional.String**| This parameter is used for finding stories whose translation body contains a specific keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).  | 
 **translationsEnText** | **optional.String**| This parameter is used for finding stories whose translation title or body contains a specific keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).  | 
 **linksPermalink** | [**optional.Interface of []string**](string.md)| This parameter is used to find stories based on their url.  | 
 **linksPermalink2** | [**optional.Interface of []string**](string.md)| This parameter is used to exclude stories based on their url.  | 
 **language** | [**optional.Interface of []string**](string.md)| This parameter is used for finding stories whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes.  | 
 **language2** | [**optional.Interface of []string**](string.md)| This parameter is used for excluding stories whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes.  | 
 **publishedAtStart** | **optional.String**| This parameter is used for finding stories whose published at time is greater than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates).  | 
 **publishedAtEnd** | **optional.String**| This parameter is used for finding stories whose published at time is less than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates).  | 
 **categoriesTaxonomy** | **optional.String**| This parameter is used for defining the type of the taxonomy for the rest of the categories queries. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories).  | 
 **categoriesConfident** | **optional.Bool**| This parameter is used for finding stories whose categories are confident. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories).  | [default to true]
 **categoriesId** | [**optional.Interface of []string**](string.md)| This parameter is used for finding stories by categories id. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories).  | 
 **categoriesId2** | [**optional.Interface of []string**](string.md)| This parameter is used for excluding stories by categories id. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories).  | 
 **categoriesLabel** | [**optional.Interface of []string**](string.md)| This parameter is used for finding stories by categories label. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories).  | 
 **categoriesLabel2** | [**optional.Interface of []string**](string.md)| This parameter is used for excluding stories by categories label. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories).  | 
 **categoriesLevel** | [**optional.Interface of []int32**](int32.md)| This parameter is used for finding stories by categories level. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories).  | 
 **categoriesLevel2** | [**optional.Interface of []int32**](int32.md)| This parameter is used for excluding stories by categories level. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories).  | 
 **entitiesTitleText** | [**optional.Interface of []string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;text&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).  | 
 **entitiesTitleText2** | [**optional.Interface of []string**](string.md)| This parameter is used to exclude stories based on the specified entities &#x60;text&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).  | 
 **entitiesTitleType** | [**optional.Interface of []string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;type&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).  | 
 **entitiesTitleType2** | [**optional.Interface of []string**](string.md)| This parameter is used to exclude stories based on the specified entities &#x60;type&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).  | 
 **entitiesTitleLinksDbpedia** | [**optional.Interface of []string**](string.md)| This parameter is used to find stories based on the specified entities dbpedia URL in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).  | 
 **entitiesTitleLinksDbpedia2** | [**optional.Interface of []string**](string.md)| This parameter is used to exclude stories based on the specified entities dbpedia URL in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).  | 
 **entitiesBodyText** | [**optional.Interface of []string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;text&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).  | 
 **entitiesBodyText2** | [**optional.Interface of []string**](string.md)| This parameter is used to exclude stories based on the specified entities &#x60;text&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).  | 
 **entitiesBodyType** | [**optional.Interface of []string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;type&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).  | 
 **entitiesBodyType2** | [**optional.Interface of []string**](string.md)| This parameter is used to exclude stories based on the specified entities &#x60;type&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).  | 
 **entitiesBodyLinksDbpedia** | [**optional.Interface of []string**](string.md)| This parameter is used to find stories based on the specified entities dbpedia URL in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).  | 
 **entitiesBodyLinksDbpedia2** | [**optional.Interface of []string**](string.md)| This parameter is used to exclude stories based on the specified entities dbpedia URL in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).  | 
 **sentimentTitlePolarity** | **optional.String**| This parameter is used for finding stories whose title sentiment is the specified value.  | 
 **sentimentTitlePolarity2** | **optional.String**| This parameter is used for excluding stories whose title sentiment is the specified value.  | 
 **sentimentBodyPolarity** | **optional.String**| This parameter is used for finding stories whose body sentiment is the specified value.  | 
 **sentimentBodyPolarity2** | **optional.String**| This parameter is used for excluding stories whose body sentiment is the specified value.  | 
 **mediaImagesCountMin** | **optional.Int32**| This parameter is used for finding stories whose number of images is greater than or equal to the specified value.  | 
 **mediaImagesCountMax** | **optional.Int32**| This parameter is used for finding stories whose number of images is less than or equal to the specified value.  | 
 **mediaImagesWidthMin** | **optional.Int32**| This parameter is used for finding stories whose width of images are greater than or equal to the specified value.  | 
 **mediaImagesWidthMax** | **optional.Int32**| This parameter is used for finding stories whose width of images are less than or equal to the specified value.  | 
 **mediaImagesHeightMin** | **optional.Int32**| This parameter is used for finding stories whose height of images are greater than or equal to the specified value.  | 
 **mediaImagesHeightMax** | **optional.Int32**| This parameter is used for finding stories whose height of images are less than or equal to the specified value.  | 
 **mediaImagesContentLengthMin** | **optional.Int32**| This parameter is used for finding stories whose images content length are greater than or equal to the specified value.  | 
 **mediaImagesContentLengthMax** | **optional.Int32**| This parameter is used for finding stories whose images content length are less than or equal to the specified value.  | 
 **mediaImagesFormat** | [**optional.Interface of []string**](string.md)| This parameter is used for finding stories whose images format are the specified value.  | 
 **mediaImagesFormat2** | [**optional.Interface of []string**](string.md)| This parameter is used for excluding stories whose images format are the specified value.  | 
 **mediaVideosCountMin** | **optional.Int32**| This parameter is used for finding stories whose number of videos is greater than or equal to the specified value.  | 
 **mediaVideosCountMax** | **optional.Int32**| This parameter is used for finding stories whose number of videos is less than or equal to the specified value.  | 
 **authorId** | [**optional.Interface of []int32**](int32.md)| This parameter is used for finding stories whose author id is the specified value.  | 
 **authorId2** | [**optional.Interface of []int32**](int32.md)| This parameter is used for excluding stories whose author id is the specified value.  | 
 **authorName** | **optional.String**| This parameter is used for finding stories whose author full name contains the specified value.  | 
 **authorName2** | **optional.String**| This parameter is used for excluding stories whose author full name contains the specified value.  | 
 **sourceId** | [**optional.Interface of []int32**](int32.md)| This parameter is used for finding stories whose source id is the specified value.  | 
 **sourceId2** | [**optional.Interface of []int32**](int32.md)| This parameter is used for excluding stories whose source id is the specified value.  | 
 **sourceName** | [**optional.Interface of []string**](string.md)| This parameter is used for finding stories whose source name contains the specified value.  | 
 **sourceName2** | [**optional.Interface of []string**](string.md)| This parameter is used for excluding stories whose source name contains the specified value.  | 
 **sourceDomain** | [**optional.Interface of []string**](string.md)| This parameter is used for finding stories whose source domain is the specified value.  | 
 **sourceDomain2** | [**optional.Interface of []string**](string.md)| This parameter is used for excluding stories whose source domain is the specified value.  | 
 **sourceLocationsCountry** | [**optional.Interface of []string**](string.md)| This parameter is used for finding stories whose source country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).  | 
 **sourceLocationsCountry2** | [**optional.Interface of []string**](string.md)| This parameter is used for excluding stories whose source country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).  | 
 **sourceLocationsState** | [**optional.Interface of []string**](string.md)| This parameter is used for finding stories whose source state/province is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).  | 
 **sourceLocationsState2** | [**optional.Interface of []string**](string.md)| This parameter is used for excluding stories whose source state/province is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).  | 
 **sourceLocationsCity** | [**optional.Interface of []string**](string.md)| This parameter is used for finding stories whose source city is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).  | 
 **sourceLocationsCity2** | [**optional.Interface of []string**](string.md)| This parameter is used for excluding stories whose source city is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).  | 
 **sourceScopesCountry** | [**optional.Interface of []string**](string.md)| This parameter is used for finding stories whose source scopes is the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).  | 
 **sourceScopesCountry2** | [**optional.Interface of []string**](string.md)| This parameter is used for excluding stories whose source scopes is the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).  | 
 **sourceScopesState** | [**optional.Interface of []string**](string.md)| This parameter is used for finding stories whose source scopes is the specified state/province value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).  | 
 **sourceScopesState2** | [**optional.Interface of []string**](string.md)| This parameter is used for excluding stories whose source scopes is the specified state/province value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).  | 
 **sourceScopesCity** | [**optional.Interface of []string**](string.md)| This parameter is used for finding stories whose source scopes is the specified city value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).  | 
 **sourceScopesCity2** | [**optional.Interface of []string**](string.md)| This parameter is used for excluding stories whose source scopes is the specified city value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).  | 
 **sourceScopesLevel** | [**optional.Interface of []string**](string.md)| This parameter is used for finding stories whose source scopes is the specified level value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).  | 
 **sourceScopesLevel2** | [**optional.Interface of []string**](string.md)| This parameter is used for excluding stories whose source scopes is the specified level value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).  | 
 **sourceLinksInCountMin** | **optional.Int32**| This parameter is used for finding stories from sources whose Links in count is greater than or equal to the specified value. You can read more about working with Links in count [here](https://newsapi.aylien.com/docs/working-with-links-in-count).  | 
 **sourceLinksInCountMax** | **optional.Int32**| This parameter is used for finding stories from sources whose Links in count is less than or equal to the specified value. You can read more about working with Links in count [here](https://newsapi.aylien.com/docs/working-with-links-in-count).  | 
 **sourceRankingsAlexaRankMin** | **optional.Int32**| This parameter is used for finding stories from sources whose Alexa rank is greater than or equal to the specified value. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks).  | 
 **sourceRankingsAlexaRankMax** | **optional.Int32**| This parameter is used for finding stories from sources whose Alexa rank is less than or equal to the specified value. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks).  | 
 **sourceRankingsAlexaCountry** | [**optional.Interface of []string**](string.md)| This parameter is used for finding stories from sources whose Alexa rank is in the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks).  | 
 **socialSharesCountFacebookMin** | **optional.Int32**| This parameter is used for finding stories whose Facebook social shares count is greater than or equal to the specified value.  | 
 **socialSharesCountFacebookMax** | **optional.Int32**| This parameter is used for finding stories whose Facebook social shares count is less than or equal to the specified value.  | 
 **socialSharesCountGooglePlusMin** | **optional.Int32**| This parameter is used for finding stories whose Google+ social shares count is greater than or equal to the specified value.  | 
 **socialSharesCountGooglePlusMax** | **optional.Int32**| This parameter is used for finding stories whose Google+ social shares count is less than or equal to the specified value.  | 
 **socialSharesCountLinkedinMin** | **optional.Int32**| This parameter is used for finding stories whose LinkedIn social shares count is greater than or equal to the specified value.  | 
 **socialSharesCountLinkedinMax** | **optional.Int32**| This parameter is used for finding stories whose LinkedIn social shares count is less than or equal to the specified value.  | 
 **socialSharesCountRedditMin** | **optional.Int32**| This parameter is used for finding stories whose Reddit social shares count is greater than or equal to the specified value.  | 
 **socialSharesCountRedditMax** | **optional.Int32**| This parameter is used for finding stories whose Reddit social shares count is less than or equal to the specified value.  | 
 **clusters** | [**optional.Interface of []string**](string.md)| This parameter is used for finding stories with belonging to one of clusters in a specific set of clusters You can read more about working with clustering [here](https://newsapi.aylien.com/docs/working-with-clustering).  | 
 **return_** | [**optional.Interface of []string**](string.md)| This parameter is used for specifying return fields. | 
 **storyId** | **optional.Int64**| A story id | 
 **storyUrl** | **optional.String**| An article or webpage | 
 **storyTitle** | **optional.String**| Title of the article | 
 **storyBody** | **optional.String**| Body of the article | 
 **storyPublishedAt** | **optional.Time**| Publish date of the article. If you use a url or title and body of an article for getting coverages, this parameter is required. The format used is a restricted form of the canonical representation of dateTime in the [XML Schema specification (ISO 8601)](https://www.w3.org/TR/xmlschema-2/#dateTime).  | 
 **storyLanguage** | **optional.String**| This parameter is used for setting the language of the story. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes.  | [default to auto]
 **perPage** | **optional.Int32**| This parameter is used for specifying number of items in each page.  | [default to 3]

### Return type

[**Coverages**](Coverages.md)

### Authorization

[app_id](../README.md#app_id), [app_key](../README.md#app_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

