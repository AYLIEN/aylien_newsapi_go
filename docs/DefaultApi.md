# DefaultApi

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
> Autocompletes ListAutocompletes(params *AutocompletesParams)

List autocompletes

This endpoint is used for getting list of autocompletes by providing a specific term and type.

### Example

``` go
package main

// Import the library
import (
	"fmt"
	newsapi "github.com/AYLIEN/aylien_newsapi_go"
)

func main() {
	api := newsapi.NewDefaultApi()

	// Configure API key authorization: app_id
	api.Configuration.APIKeyPrefix["X-AYLIEN-NewsAPI-Application-ID"] = "YOUR APP ID"

	// Configure API key authorization: app_key
	api.Configuration.APIKeyPrefix["X-AYLIEN-NewsAPI-Application-Key"] = "YOUR APP KEY"

	type_ := "type_example" // string | This parameter is used for defining the type of autocompletes.
	term := "term_example"  // string | This parameter is used for finding autocomplete objects that contain the specified value.
	language := "en"        // string | This parameter is used for autocompletes whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes.
	perPage := 3            // int32 | This parameter is used for specifying number of the items in each page.

	autocompletesParams := &newsapi.AutocompletesParams{
		Type:     type_,
		Term:     term,
		Language: language,
		PerPage:  perPage}

	autocompletesResponse, res, err := api.ListAutocompletes(autocompletesParams)
	if err != nil {
		panic(err)
	}
	_ = res

	fmt.Println(autocompletesResponse)
}
```

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
 - **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCoverages**
> Coverages ListCoverages(params *CoveragesParams)

List coverages

This endpoint is used for finding story coverages based on the parameters provided. The maximum number of related stories returned is 100.

### Example

``` go
package main

// Import the library
import (
	"fmt"
	newsapi "github.com/AYLIEN/aylien_newsapi_go"
	"time"
)

func main() {
	api := newsapi.NewDefaultApi()

	// Configure API key authorization: app_id
	api.Configuration.APIKeyPrefix["X-AYLIEN-NewsAPI-Application-ID"] = "YOUR APP ID"

	// Configure API key authorization: app_key
	api.Configuration.APIKeyPrefix["X-AYLIEN-NewsAPI-Application-Key"] = "YOUR APP KEY"

	id := []int64{112}
	title := "title_example"
	body := "body_example"
	text := "text_example"
	language := []string{"en"}
	publishedAtStart := "publishedAtStart_example"
	publishedAtEnd := "publishedAtEnd_example"
	categoriesTaxonomy := "categoriesTaxonomy_example"
	categoriesConfident := "categoriesConfident_example"
	categoriesId := []string{"categoriesId_example"}
	categoriesLevel := []int32{3}
	entitiesTitleText := []string{"entitiesTitleText_example"}
	entitiesTitleType := []string{"entitiesTitleType_example"}
	entitiesTitleLinksDbpedia := []string{"entitiesTitleLinksDbpedia_example"}
	entitiesBodyText := []string{"entitiesBodyText_example"}
	entitiesBodyType := []string{"entitiesBodyType_example"}
	entitiesBodyLinksDbpedia := []string{"entitiesBodyLinksDbpedia_example"}
	sentimentTitlePolarity := "sentimentTitlePolarity_example"
	sentimentBodyPolarity := "sentimentBodyPolarity_example"
	mediaImagesCountMin := 56
	mediaImagesCountMax := 56
	mediaVideosCountMin := 56
	mediaVideosCountMax := 56
	authorId := []int32{25}
	authorName := "authorName_example"
	sourceId := []int32{45}
	sourceName := []string{"sourceName_example"}
	sourceDomain := []string{"sourceDomain_example"}
	sourceLocationsCountry := []string{"sourceLocationsCountry_example"}
	sourceLocationsState := []string{"sourceLocationsState_example"}
	sourceLocationsCity := []string{"sourceLocationsCity_example"}
	sourceScopesCountry := []string{"sourceScopesCountry_example"}
	sourceScopesState := []string{"sourceScopesState_example"}
	sourceScopesCity := []string{"sourceScopesCity_example"}
	sourceScopesLevel := []string{"sourceScopesLevel_example"}
	sourceLinksInCountMin := 56
	sourceLinksInCountMax := 56
	sourceRankingsAlexaRankMin := 56
	sourceRankingsAlexaRankMax := 56
	sourceRankingsAlexaCountry := []string{"sourceRankingsAlexaCountry_example"}
	cluster := true
	clusterAlgorithm := "lingo"
	_return := []string{"_return_example"}
	storyId := 789
	storyUrl := "storyUrl_example"
	storyTitle := "storyTitle_example"
	storyBody := "storyBody_example"
	storyPublishedAt := time.Now()
	storyLanguage := "auto"
	perPage := 3

	coveragesParams := &newsapi.CoveragesParams{
		Id:                         id,
		Title:                      title,
		Body:                       body,
		Text:                       text,
		Language:                   language,
		PublishedAtStart:           publishedAtStart,
		PublishedAtEnd:             publishedAtEnd,
		CategoriesTaxonomy:         categoriesTaxonomy,
		CategoriesConfident:        categoriesConfident,
		CategoriesId:               categoriesId,
		CategoriesLevel:            categoriesLevel,
		EntitiesTitleText:          entitiesTitleText,
		EntitiesTitleType:          entitiesTitleType,
		EntitiesTitleLinksDbpedia:  entitiesTitleLinksDbpedia,
		EntitiesBodyText:           entitiesBodyText,
		EntitiesBodyType:           entitiesBodyType,
		EntitiesBodyLinksDbpedia:   entitiesBodyLinksDbpedia,
		SentimentTitlePolarity:     sentimentTitlePolarity,
		SentimentBodyPolarity:      sentimentBodyPolarity,
		MediaImagesCountMin:        mediaImagesCountMin,
		MediaImagesCountMax:        mediaImagesCountMax,
		MediaVideosCountMin:        mediaVideosCountMin,
		MediaVideosCountMin:        mediaVideosCountMin,
		AuthorId:                   authorId,
		AuthorName:                 authorName,
		SourceId:                   sourceId,
		SourceName:                 sourceName,
		SourceDomain:               sourceDomain,
		SourceLocationsCountry:     sourceLocationsCountry,
		SourceLocationsState:       sourceLocationsState,
		SourceLocationsCity:        sourceLocationsCity,
		SourceScopesCountry:        sourceScopesCountry,
		SourceScopesState:          sourceScopesState,
		SourceScopesCity:           sourceScopesCity,
		SourceScopesLevel:          sourceScopesLevel,
		SourceLinksInCountMin:      sourceLinksInCountMin,
		SourceLinksInCountMax:      sourceLinksInCountMax,
		SourceRankingsAlexaRankMin: sourceRankingsAlexaRankMin,
		SourceRankingsAlexaRankMax: sourceRankingsAlexaRankMax,
		SourceRankingsAlexaCountry: sourceRankingsAlexaCountry,
		Cluster:                    cluster,
		ClusterAlgorithm:           clusterAlgorithm,
		Return:                     _return,
		StoryId:                    storyId,
		StoryUrl:                   storyUrl,
		StoryTitle:                 storyTitle,
		StoryBody:                  storyBody,
		StoryPublishedAt:           storyPublishedAt,
		StoryLanguage:              storyLanguage,
		PerPage:                    perPage}

	coveragesResponse, res, err := api.ListCoverages(coveragesParams)
	if err != nil {
		panic(err)
	}
	_ = res

	fmt.Println(coveragesResponse)
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | [**[]int64**](int64.md)| This parameter is used for finding stroies by story id. | [optional] 
 **Title** | **string**| This parameter is used for finding stories whose title contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Body** | **string**| This parameter is used for finding stories whose body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Text** | **string**| This parameter is used for finding stories whose title or body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Language** | [**[]string**](string.md)| This parameter is used for finding stories whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes. | [optional] 
 **PublishedAtStart** | **string**| This parameter is used for finding stories whose published at time is greater than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates). | [optional] 
 **PublishedAtEnd** | **string**| This parameter is used for finding stories whose published at time is less than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates). | [optional] 
 **CategoriesTaxonomy** | **string**| This parameter is used for defining the type of the taxonomy for the rest of the categories queries. | [optional] 
 **CategoriesConfident** | **bool**| This parameter is used for finding stories whose categories are confident. | [optional] [default to true]
 **CategoriesId** | [**[]string**](string.md)| This parameter is used for finding stories by categories id. | [optional] 
 **CategoriesLevel** | [**[]int32**](int32.md)| This parameter is used for finding stories by categories level. | [optional] 
 **EntitiesTitleText** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;text&#x60; in story titles. | [optional] 
 **EntitiesTitleType** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;type&#x60; in story titles. | [optional] 
 **EntitiesTitleLinksDbpedia** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities dbpedia URL in story titles. | [optional] 
 **EntitiesBodyText** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;text&#x60; in the body of stories. | [optional] 
 **EntitiesBodyType** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;type&#x60; in the body of stories. | [optional] 
 **EntitiesBodyLinksDbpedia** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities dbpedia URL in the body of stories. | [optional] 
 **SentimentTitlePolarity** | **string**| This parameter is used for finding stories whose title sentiment is the specified value. | [optional] 
 **SentimentBodyPolarity** | **string**| This parameter is used for finding stories whose body sentiment is the specified value. | [optional] 
 **MediaImagesCountMin** | **int32**| This parameter is used for finding stories whose number of images is greater than or equal to the specified value. | [optional] 
 **MediaImagesCountMax** | **int32**| This parameter is used for finding stories whose number of images is less than or equal to the specified value. | [optional] 
 **MediaVideosCountMin** | **int32**| This parameter is used for finding stories whose number of videos is greater than or equal to the specified value. | [optional] 
 **MediaVideosCountMax** | **int32**| This parameter is used for finding stories whose number of videos is less than or equal to the specified value. | [optional] 
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
 **SourceLinksInCountMin** | **int32**| This parameter is used for finding stories from sources whose Links in count is greater than or equal to the specified value. You can read more about working with Links in count here [https://newsapi.aylien.com/docs/working-with-links-in-count](https://newsapi.aylien.com/docs/working-with-links-in-count). | [optional] 
 **SourceLinksInCountMax** | **int32**| This parameter is used for finding stories from sources whose Links in count is less than or equal to the specified value. You can read more about working with Links in count here [https://newsapi.aylien.com/docs/working-with-links-in-count](https://newsapi.aylien.com/docs/working-with-links-in-count). | [optional] 
 **SourceRankingsAlexaRankMin** | **int32**| This parameter is used for finding stories from sources whose Alexa rank is greater than or equal to the specified value. You can read more about working with Alexa ranks here [https://newsapi.aylien.com/docs/working-with-alexa-ranks](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
 **SourceRankingsAlexaRankMax** | **int32**| This parameter is used for finding stories from sources whose Alexa rank is less than or equal to the specified value. You can read more about working with Alexa ranks here [https://newsapi.aylien.com/docs/working-with-alexa-ranks](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
 **SourceRankingsAlexaCountry** | [**[]string**](string.md)| This parameter is used for finding stories from sources whose Alexa rank is in the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. You can read more about working with Alexa ranks here [https://newsapi.aylien.com/docs/working-with-alexa-ranks](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
 **Cluster** | **bool**| This parameter enables clustering for the returned stories. | [optional] [default to false]
 **ClusterAlgorithm** | **string**| This parameter is used for specifying the clustering algorithm you wish to use. It supprts STC, Lingo and [k-means](https://en.wikipedia.org/wiki/K-means_clustering) algorithms. | [optional] [default to lingo]
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
 - **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListHistograms**
> Histograms ListHistograms(params *HistogramsParams)

List histograms

This endpoint is used for getting histograms based on the `field` parameter passed to the API.

### Example

``` go
package main

// Import the library
import (
	"fmt"
	newsapi "github.com/AYLIEN/aylien_newsapi_go"
)

func main() {
	api := newsapi.NewDefaultApi()

	// Configure API key authorization: app_id
	api.Configuration.APIKeyPrefix["X-AYLIEN-NewsAPI-Application-ID"] = "YOUR APP ID"

	// Configure API key authorization: app_key
	api.Configuration.APIKeyPrefix["X-AYLIEN-NewsAPI-Application-Key"] = "YOUR APP KEY"

	id := []int64{112}
	title := "title_example"
	body := "body_example"
	text := "text_example"
	language := []string{"en"}
	publishedAtStart := "publishedAtStart_example"
	publishedAtEnd := "publishedAtEnd_example"
	categoriesTaxonomy := "categoriesTaxonomy_example"
	categoriesConfident := "categoriesConfident_example"
	categoriesId := []string{"categoriesId_example"}
	categoriesLevel := []int32{3}
	entitiesTitleText := []string{"entitiesTitleText_example"}
	entitiesTitleType := []string{"entitiesTitleType_example"}
	entitiesTitleLinksDbpedia := []string{"entitiesTitleLinksDbpedia_example"}
	entitiesBodyText := []string{"entitiesBodyText_example"}
	entitiesBodyType := []string{"entitiesBodyType_example"}
	entitiesBodyLinksDbpedia := []string{"entitiesBodyLinksDbpedia_example"}
	sentimentTitlePolarity := "sentimentTitlePolarity_example"
	sentimentBodyPolarity := "sentimentBodyPolarity_example"
	mediaImagesCountMin := 56
	mediaImagesCountMax := 56
	mediaVideosCountMin := 56
	mediaVideosCountMax := 56
	authorId := []int32{25}
	authorName := "authorName_example"
	sourceId := []int32{45}
	sourceName := []string{"sourceName_example"}
	sourceDomain := []string{"sourceDomain_example"}
	sourceLocationsCountry := []string{"sourceLocationsCountry_example"}
	sourceLocationsState := []string{"sourceLocationsState_example"}
	sourceLocationsCity := []string{"sourceLocationsCity_example"}
	sourceScopesCountry := []string{"sourceScopesCountry_example"}
	sourceScopesState := []string{"sourceScopesState_example"}
	sourceScopesCity := []string{"sourceScopesCity_example"}
	sourceScopesLevel := []string{"sourceScopesLevel_example"}
	sourceLinksInCountMin := 56
  sourceLinksInCountMax := 56
  sourceRankingsAlexaRankMin := 56
  sourceRankingsAlexaRankMax := 56
  sourceRankingsAlexaCountry := []string{"sourceRankingsAlexaCountry_example"}
	intervalStart := 56
	intervalEnd := 56
	intervalWidth := 56
	field := "social_shares_count"

	histogramsParams := &newsapi.HistogramsParams{
		Id:                         id,
		Title:                      title,
		Body:                       body,
		Text:                       text,
		Language:                   language,
		PublishedAtStart:           publishedAtStart,
		PublishedAtEnd:             publishedAtEnd,
		CategoriesTaxonomy:         categoriesTaxonomy,
		CategoriesConfident:        categoriesConfident,
		CategoriesId:               categoriesId,
		CategoriesLevel:            categoriesLevel,
		EntitiesTitleText:          entitiesTitleText,
		EntitiesTitleType:          entitiesTitleType,
		EntitiesTitleLinksDbpedia:  entitiesTitleLinksDbpedia,
		EntitiesBodyText:           entitiesBodyText,
		EntitiesBodyType:           entitiesBodyType,
		EntitiesBodyLinksDbpedia:   entitiesBodyLinksDbpedia,
		SentimentTitlePolarity:     sentimentTitlePolarity,
		SentimentBodyPolarity:      sentimentBodyPolarity,
		MediaImagesCountMin:        mediaImagesCountMin,
		MediaImagesCountMax:        mediaImagesCountMax,
		MediaVideosCountMin:        mediaVideosCountMin,
		MediaVideosCountMin:        mediaVideosCountMin,
		AuthorId:                   authorId,
		AuthorName:                 authorName,
		SourceId:                   sourceId,
		SourceName:                 sourceName,
		SourceDomain:               sourceDomain,
		SourceLocationsCountry:     sourceLocationsCountry,
		SourceLocationsState:       sourceLocationsState,
		SourceLocationsCity:        sourceLocationsCity,
		SourceScopesCountry:        sourceScopesCountry,
		SourceScopesState:          sourceScopesState,
		SourceScopesCity:           sourceScopesCity,
		SourceScopesLevel:          sourceScopesLevel,
		SourceLinksInCountMin:      sourceLinksInCountMin,
		SourceLinksInCountMax:      sourceLinksInCountMax,
		SourceRankingsAlexaRankMin: sourceRankingsAlexaRankMin,
		SourceRankingsAlexaRankMax: sourceRankingsAlexaRankMax,
		SourceRankingsAlexaCountry: sourceRankingsAlexaCountry,
		IntervalStart:              intervalStart,
		IntervalEnd:                intervalEnd,
		IntervalWidth:              intervalWidth,
		Field:                      field}

	histogramsResponse, res, err := api.ListHistograms(histogramsParams)
	if err != nil {
		panic(err)
	}
	_ = res

	fmt.Println(histogramsResponse)
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | [**[]int64**](int64.md)| This parameter is used for finding stroies by story id. | [optional] 
 **Title** | **string**| This parameter is used for finding stories whose title contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Body** | **string**| This parameter is used for finding stories whose body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Text** | **string**| This parameter is used for finding stories whose title or body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Language** | [**[]string**](string.md)| This parameter is used for finding stories whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes. | [optional] 
 **PublishedAtStart** | **string**| This parameter is used for finding stories whose published at time is greater than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates). | [optional] 
 **PublishedAtEnd** | **string**| This parameter is used for finding stories whose published at time is less than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates). | [optional] 
 **CategoriesTaxonomy** | **string**| This parameter is used for defining the type of the taxonomy for the rest of the categories queries. | [optional] 
 **CategoriesConfident** | **bool**| This parameter is used for finding stories whose categories are confident. | [optional] [default to true]
 **CategoriesId** | [**[]string**](string.md)| This parameter is used for finding stories by categories id. | [optional] 
 **CategoriesLevel** | [**[]int32**](int32.md)| This parameter is used for finding stories by categories level. | [optional] 
 **EntitiesTitleText** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;text&#x60; in story titles. | [optional] 
 **EntitiesTitleType** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;type&#x60; in story titles. | [optional] 
 **EntitiesTitleLinksDbpedia** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities dbpedia URL in story titles. | [optional] 
 **EntitiesBodyText** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;text&#x60; in the body of stories. | [optional] 
 **EntitiesBodyType** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;type&#x60; in the body of stories. | [optional] 
 **EntitiesBodyLinksDbpedia** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities dbpedia URL in the body of stories. | [optional] 
 **SentimentTitlePolarity** | **string**| This parameter is used for finding stories whose title sentiment is the specified value. | [optional] 
 **SentimentBodyPolarity** | **string**| This parameter is used for finding stories whose body sentiment is the specified value. | [optional] 
 **MediaImagesCountMin** | **int32**| This parameter is used for finding stories whose number of images is greater than or equal to the specified value. | [optional] 
 **MediaImagesCountMax** | **int32**| This parameter is used for finding stories whose number of images is less than or equal to the specified value. | [optional] 
 **MediaVideosCountMin** | **int32**| This parameter is used for finding stories whose number of videos is greater than or equal to the specified value. | [optional] 
 **MediaVideosCountMax** | **int32**| This parameter is used for finding stories whose number of videos is less than or equal to the specified value. | [optional] 
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
 **SourceLinksInCountMin** | **int32**| This parameter is used for finding stories from sources whose Links in count is greater than or equal to the specified value. You can read more about working with Links in count here [https://newsapi.aylien.com/docs/working-with-links-in-count](https://newsapi.aylien.com/docs/working-with-links-in-count). | [optional] 
 **SourceLinksInCountMax** | **int32**| This parameter is used for finding stories from sources whose Links in count is less than or equal to the specified value. You can read more about working with Links in count here [https://newsapi.aylien.com/docs/working-with-links-in-count](https://newsapi.aylien.com/docs/working-with-links-in-count). | [optional] 
 **SourceRankingsAlexaRankMin** | **int32**| This parameter is used for finding stories from sources whose Alexa rank is greater than or equal to the specified value. You can read more about working with Alexa ranks here [https://newsapi.aylien.com/docs/working-with-alexa-ranks](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
 **SourceRankingsAlexaRankMax** | **int32**| This parameter is used for finding stories from sources whose Alexa rank is less than or equal to the specified value. You can read more about working with Alexa ranks here [https://newsapi.aylien.com/docs/working-with-alexa-ranks](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
 **SourceRankingsAlexaCountry** | [**[]string**](string.md)| This parameter is used for finding stories from sources whose Alexa rank is in the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. You can read more about working with Alexa ranks here [https://newsapi.aylien.com/docs/working-with-alexa-ranks](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
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
 - **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRelatedStories**
> RelatedStories ListRelatedStories(params *RelatedStoriesParams)

List related stories

This endpoint is used for finding related stories based on the parameters provided. The maximum number of related stories returned is 100.

### Example

``` go
package main

// Import the library
import (
	"fmt"
	newsapi "github.com/AYLIEN/aylien_newsapi_go"
)

func main() {
	api := newsapi.NewDefaultApi()

	// Configure API key authorization: app_id
	api.Configuration.APIKeyPrefix["X-AYLIEN-NewsAPI-Application-ID"] = "YOUR APP ID"

	// Configure API key authorization: app_key
	api.Configuration.APIKeyPrefix["X-AYLIEN-NewsAPI-Application-Key"] = "YOUR APP KEY"

	id := []int64{112}
	title := "title_example"
	body := "body_example"
	text := "text_example"
	language := []string{"en"}
	publishedAtStart := "publishedAtStart_example"
	publishedAtEnd := "publishedAtEnd_example"
	categoriesTaxonomy := "categoriesTaxonomy_example"
	categoriesConfident := "categoriesConfident_example"
	categoriesId := []string{"categoriesId_example"}
	categoriesLevel := []int32{3}
	entitiesTitleText := []string{"entitiesTitleText_example"}
	entitiesTitleType := []string{"entitiesTitleType_example"}
	entitiesTitleLinksDbpedia := []string{"entitiesTitleLinksDbpedia_example"}
	entitiesBodyText := []string{"entitiesBodyText_example"}
	entitiesBodyType := []string{"entitiesBodyType_example"}
	entitiesBodyLinksDbpedia := []string{"entitiesBodyLinksDbpedia_example"}
	sentimentTitlePolarity := "sentimentTitlePolarity_example"
	sentimentBodyPolarity := "sentimentBodyPolarity_example"
	mediaImagesCountMin := 56
	mediaImagesCountMax := 56
	mediaVideosCountMin := 56
	mediaVideosCountMax := 56
	authorId := []int32{25}
	authorName := "authorName_example"
	sourceId := []int32{45}
	sourceName := []string{"sourceName_example"}
	sourceDomain := []string{"sourceDomain_example"}
	sourceLocationsCountry := []string{"sourceLocationsCountry_example"}
	sourceLocationsState := []string{"sourceLocationsState_example"}
	sourceLocationsCity := []string{"sourceLocationsCity_example"}
	sourceScopesCountry := []string{"sourceScopesCountry_example"}
	sourceScopesState := []string{"sourceScopesState_example"}
	sourceScopesCity := []string{"sourceScopesCity_example"}
	sourceScopesLevel := []string{"sourceScopesLevel_example"}
	sourceLinksInCountMin := 56
  sourceLinksInCountMax := 56
  sourceRankingsAlexaRankMin := 56
  sourceRankingsAlexaRankMax := 56
  sourceRankingsAlexaCountry := []string{"sourceRankingsAlexaCountry_example"}
	cluster := true
	clusterAlgorithm := "lingo"
	_return := []string{"_return_example"}
	storyId := 789
	storyUrl := "storyUrl_example"
	storyTitle := "storyTitle_example"
	storyBody := "storyBody_example"
	boostBy := "recency"
	storyLanguage := "auto"
	perPage := 3

	relatedStoriesParams := &newsapi.RelatedStoriesParams{
		Id:                         id,
		Title:                      title,
		Body:                       body,
		Text:                       text,
		Language:                   language,
		PublishedAtStart:           publishedAtStart,
		PublishedAtEnd:             publishedAtEnd,
		CategoriesTaxonomy:         categoriesTaxonomy,
		CategoriesConfident:        categoriesConfident,
		CategoriesId:               categoriesId,
		CategoriesLevel:            categoriesLevel,
		EntitiesTitleText:          entitiesTitleText,
		EntitiesTitleType:          entitiesTitleType,
		EntitiesTitleLinksDbpedia:  entitiesTitleLinksDbpedia,
		EntitiesBodyText:           entitiesBodyText,
		EntitiesBodyType:           entitiesBodyType,
		EntitiesBodyLinksDbpedia:   entitiesBodyLinksDbpedia,
		SentimentTitlePolarity:     sentimentTitlePolarity,
		SentimentBodyPolarity:      sentimentBodyPolarity,
		MediaImagesCountMin:        mediaImagesCountMin,
		MediaImagesCountMax:        mediaImagesCountMax,
		MediaVideosCountMin:        mediaVideosCountMin,
		MediaVideosCountMin:        mediaVideosCountMin,
		AuthorId:                   authorId,
		AuthorName:                 authorName,
		SourceId:                   sourceId,
		SourceName:                 sourceName,
		SourceDomain:               sourceDomain,
		SourceLocationsCountry:     sourceLocationsCountry,
		SourceLocationsState:       sourceLocationsState,
		SourceLocationsCity:        sourceLocationsCity,
		SourceScopesCountry:        sourceScopesCountry,
		SourceScopesState:          sourceScopesState,
		SourceScopesCity:           sourceScopesCity,
		SourceScopesLevel:          sourceScopesLevel,
		SourceLinksInCountMin:      sourceLinksInCountMin,
		SourceLinksInCountMax:      sourceLinksInCountMax,
		SourceRankingsAlexaRankMin: sourceRankingsAlexaRankMin,
		SourceRankingsAlexaRankMax: sourceRankingsAlexaRankMax,
		SourceRankingsAlexaCountry: sourceRankingsAlexaCountry,
		Cluster:                    cluster,
		ClusterAlgorithm:           clusterAlgorithm,
		Return:                     _return,
		StoryId:                    storyId,
		StoryUrl:                   storyUrl,
		StoryTitle:                 storyTitle,
		StoryBody:                  storyBody,
		BoostBy:                    boostBy,
		StoryLanguage:              storyLanguage,
		PerPage:                    perPage}

	relatedStoriesResponse, res, err := api.ListRelatedStories(relatedStoriesParams)
	if err != nil {
		panic(err)
	}
	_ = res

	fmt.Println(relatedStoriesResponse)
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | [**[]int64**](int64.md)| This parameter is used for finding stroies by story id. | [optional] 
 **Title** | **string**| This parameter is used for finding stories whose title contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Body** | **string**| This parameter is used for finding stories whose body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Text** | **string**| This parameter is used for finding stories whose title or body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Language** | [**[]string**](string.md)| This parameter is used for finding stories whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes. | [optional] 
 **PublishedAtStart** | **string**| This parameter is used for finding stories whose published at time is greater than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates). | [optional] 
 **PublishedAtEnd** | **string**| This parameter is used for finding stories whose published at time is less than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates). | [optional] 
 **CategoriesTaxonomy** | **string**| This parameter is used for defining the type of the taxonomy for the rest of the categories queries. | [optional] 
 **CategoriesConfident** | **bool**| This parameter is used for finding stories whose categories are confident. | [optional] [default to true]
 **CategoriesId** | [**[]string**](string.md)| This parameter is used for finding stories by categories id. | [optional] 
 **CategoriesLevel** | [**[]int32**](int32.md)| This parameter is used for finding stories by categories level. | [optional] 
 **EntitiesTitleText** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;text&#x60; in story titles. | [optional] 
 **EntitiesTitleType** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;type&#x60; in story titles. | [optional] 
 **EntitiesTitleLinksDbpedia** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities dbpedia URL in story titles. | [optional] 
 **EntitiesBodyText** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;text&#x60; in the body of stories. | [optional] 
 **EntitiesBodyType** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;type&#x60; in the body of stories. | [optional] 
 **EntitiesBodyLinksDbpedia** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities dbpedia URL in the body of stories. | [optional] 
 **SentimentTitlePolarity** | **string**| This parameter is used for finding stories whose title sentiment is the specified value. | [optional] 
 **SentimentBodyPolarity** | **string**| This parameter is used for finding stories whose body sentiment is the specified value. | [optional] 
 **MediaImagesCountMin** | **int32**| This parameter is used for finding stories whose number of images is greater than or equal to the specified value. | [optional] 
 **MediaImagesCountMax** | **int32**| This parameter is used for finding stories whose number of images is less than or equal to the specified value. | [optional] 
 **MediaVideosCountMin** | **int32**| This parameter is used for finding stories whose number of videos is greater than or equal to the specified value. | [optional] 
 **MediaVideosCountMax** | **int32**| This parameter is used for finding stories whose number of videos is less than or equal to the specified value. | [optional] 
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
 **SourceLinksInCountMin** | **int32**| This parameter is used for finding stories from sources whose Links in count is greater than or equal to the specified value. You can read more about working with Links in count here [https://newsapi.aylien.com/docs/working-with-links-in-count](https://newsapi.aylien.com/docs/working-with-links-in-count). | [optional] 
 **SourceLinksInCountMax** | **int32**| This parameter is used for finding stories from sources whose Links in count is less than or equal to the specified value. You can read more about working with Links in count here [https://newsapi.aylien.com/docs/working-with-links-in-count](https://newsapi.aylien.com/docs/working-with-links-in-count). | [optional] 
 **SourceRankingsAlexaRankMin** | **int32**| This parameter is used for finding stories from sources whose Alexa rank is greater than or equal to the specified value. You can read more about working with Alexa ranks here [https://newsapi.aylien.com/docs/working-with-alexa-ranks](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
 **SourceRankingsAlexaRankMax** | **int32**| This parameter is used for finding stories from sources whose Alexa rank is less than or equal to the specified value. You can read more about working with Alexa ranks here [https://newsapi.aylien.com/docs/working-with-alexa-ranks](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
 **SourceRankingsAlexaCountry** | [**[]string**](string.md)| This parameter is used for finding stories from sources whose Alexa rank is in the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. You can read more about working with Alexa ranks here [https://newsapi.aylien.com/docs/working-with-alexa-ranks](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
 **Cluster** | **bool**| This parameter enables clustering for the returned stories. | [optional] [default to false]
 **ClusterAlgorithm** | **string**| This parameter is used for specifying the clustering algorithm you wish to use. It supprts STC, Lingo and [k-means](https://en.wikipedia.org/wiki/K-means_clustering) algorithms. | [optional] [default to lingo]
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
 - **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListStories**
> Stories ListStories(params *StoriesParams)

List Stories

This endpoint is used for getting a list of stories.

### Example

``` go
package main

// Import the library
import (
	"fmt"
	newsapi "github.com/AYLIEN/aylien_newsapi_go"
)

func main() {
	api := newsapi.NewDefaultApi()

	// Configure API key authorization: app_id
	api.Configuration.APIKeyPrefix["X-AYLIEN-NewsAPI-Application-ID"] = "YOUR APP ID"

	// Configure API key authorization: app_key
	api.Configuration.APIKeyPrefix["X-AYLIEN-NewsAPI-Application-Key"] = "YOUR APP KEY"

	id := []int64{112}
	title := "title_example"
	body := "body_example"
	text := "text_example"
	language := []string{"en"}
	publishedAtStart := "publishedAtStart_example"
	publishedAtEnd := "publishedAtEnd_example"
	categoriesTaxonomy := "categoriesTaxonomy_example"
	categoriesConfident := "categoriesConfident_example"
	categoriesId := []string{"categoriesId_example"}
	categoriesLevel := []int32{3}
	entitiesTitleText := []string{"entitiesTitleText_example"}
	entitiesTitleType := []string{"entitiesTitleType_example"}
	entitiesTitleLinksDbpedia := []string{"entitiesTitleLinksDbpedia_example"}
	entitiesBodyText := []string{"entitiesBodyText_example"}
	entitiesBodyType := []string{"entitiesBodyType_example"}
	entitiesBodyLinksDbpedia := []string{"entitiesBodyLinksDbpedia_example"}
	sentimentTitlePolarity := "sentimentTitlePolarity_example"
	sentimentBodyPolarity := "sentimentBodyPolarity_example"
	mediaImagesCountMin := 56
	mediaImagesCountMax := 56
	mediaVideosCountMin := 56
	mediaVideosCountMax := 56
	authorId := []int32{25}
	authorName := "authorName_example"
	sourceId := []int32{45}
	sourceName := []string{"sourceName_example"}
	sourceDomain := []string{"sourceDomain_example"}
	sourceLocationsCountry := []string{"sourceLocationsCountry_example"}
	sourceLocationsState := []string{"sourceLocationsState_example"}
	sourceLocationsCity := []string{"sourceLocationsCity_example"}
	sourceScopesCountry := []string{"sourceScopesCountry_example"}
	sourceScopesState := []string{"sourceScopesState_example"}
	sourceScopesCity := []string{"sourceScopesCity_example"}
	sourceScopesLevel := []string{"sourceScopesLevel_example"}
	sourceLinksInCountMin := 56
  sourceLinksInCountMax := 56
  sourceRankingsAlexaRankMin := 56
  sourceRankingsAlexaRankMax := 56
  sourceRankingsAlexaCountry := []string{"sourceRankingsAlexaCountry_example"}
	cluster := true
	clusterAlgorithm := "lingo"
	_return := []string{"_return_example"}
	sortBy := "hotness"
	sortDirection := "desc"
	cursor := "*"
	perPage := 10

	storiesParams := &newsapi.StoriesParams{
		Id:                         id,
		Title:                      title,
		Body:                       body,
		Text:                       text,
		Language:                   language,
		PublishedAtStart:           publishedAtStart,
		PublishedAtEnd:             publishedAtEnd,
		CategoriesTaxonomy:         categoriesTaxonomy,
		CategoriesConfident:        categoriesConfident,
		CategoriesId:               categoriesId,
		CategoriesLevel:            categoriesLevel,
		EntitiesTitleText:          entitiesTitleText,
		EntitiesTitleType:          entitiesTitleType,
		EntitiesTitleLinksDbpedia:  entitiesTitleLinksDbpedia,
		EntitiesBodyText:           entitiesBodyText,
		EntitiesBodyType:           entitiesBodyType,
		EntitiesBodyLinksDbpedia:   entitiesBodyLinksDbpedia,
		SentimentTitlePolarity:     sentimentTitlePolarity,
		SentimentBodyPolarity:      sentimentBodyPolarity,
		MediaImagesCountMin:        mediaImagesCountMin,
		MediaImagesCountMax:        mediaImagesCountMax,
		MediaVideosCountMin:        mediaVideosCountMin,
		MediaVideosCountMin:        mediaVideosCountMin,
		AuthorId:                   authorId,
		AuthorName:                 authorName,
		SourceId:                   sourceId,
		SourceName:                 sourceName,
		SourceDomain:               sourceDomain,
		SourceLocationsCountry:     sourceLocationsCountry,
		SourceLocationsState:       sourceLocationsState,
		SourceLocationsCity:        sourceLocationsCity,
		SourceScopesCountry:        sourceScopesCountry,
		SourceScopesState:          sourceScopesState,
		SourceScopesCity:           sourceScopesCity,
		SourceScopesLevel:          sourceScopesLevel,
		SourceLinksInCountMin:      sourceLinksInCountMin,
		SourceLinksInCountMax:      sourceLinksInCountMax,
		SourceRankingsAlexaRankMin: sourceRankingsAlexaRankMin,
		SourceRankingsAlexaRankMax: sourceRankingsAlexaRankMax,
		SourceRankingsAlexaCountry: sourceRankingsAlexaCountry,
		Cluster:                    cluster,
		ClusterAlgorithm:           clusterAlgorithm,
		Return:                     _return,
		SortBy:                     sortBy,
		SortDirection:              sortDirection,
		Cursor:                     cursor,
		PerPage:                    perPage}

	storiesResponse, res, err := api.ListStories(storiesParams)
	if err != nil {
		panic(err)
	}
	_ = res

	fmt.Println(storiesResponse)
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | [**[]int64**](int64.md)| This parameter is used for finding stroies by story id. | [optional] 
 **Title** | **string**| This parameter is used for finding stories whose title contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Body** | **string**| This parameter is used for finding stories whose body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Text** | **string**| This parameter is used for finding stories whose title or body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Language** | [**[]string**](string.md)| This parameter is used for finding stories whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes. | [optional] 
 **PublishedAtStart** | **string**| This parameter is used for finding stories whose published at time is greater than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates). | [optional] 
 **PublishedAtEnd** | **string**| This parameter is used for finding stories whose published at time is less than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates). | [optional] 
 **CategoriesTaxonomy** | **string**| This parameter is used for defining the type of the taxonomy for the rest of the categories queries. | [optional] 
 **CategoriesConfident** | **bool**| This parameter is used for finding stories whose categories are confident. | [optional] [default to true]
 **CategoriesId** | [**[]string**](string.md)| This parameter is used for finding stories by categories id. | [optional] 
 **CategoriesLevel** | [**[]int32**](int32.md)| This parameter is used for finding stories by categories level. | [optional] 
 **EntitiesTitleText** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;text&#x60; in story titles. | [optional] 
 **EntitiesTitleType** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;type&#x60; in story titles. | [optional] 
 **EntitiesTitleLinksDbpedia** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities dbpedia URL in story titles. | [optional] 
 **EntitiesBodyText** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;text&#x60; in the body of stories. | [optional] 
 **EntitiesBodyType** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;type&#x60; in the body of stories. | [optional] 
 **EntitiesBodyLinksDbpedia** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities dbpedia URL in the body of stories. | [optional] 
 **SentimentTitlePolarity** | **string**| This parameter is used for finding stories whose title sentiment is the specified value. | [optional] 
 **SentimentBodyPolarity** | **string**| This parameter is used for finding stories whose body sentiment is the specified value. | [optional] 
 **MediaImagesCountMin** | **int32**| This parameter is used for finding stories whose number of images is greater than or equal to the specified value. | [optional] 
 **MediaImagesCountMax** | **int32**| This parameter is used for finding stories whose number of images is less than or equal to the specified value. | [optional] 
 **MediaVideosCountMin** | **int32**| This parameter is used for finding stories whose number of videos is greater than or equal to the specified value. | [optional] 
 **MediaVideosCountMax** | **int32**| This parameter is used for finding stories whose number of videos is less than or equal to the specified value. | [optional] 
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
 **SourceLinksInCountMin** | **int32**| This parameter is used for finding stories from sources whose Links in count is greater than or equal to the specified value. You can read more about working with Links in count here [https://newsapi.aylien.com/docs/working-with-links-in-count](https://newsapi.aylien.com/docs/working-with-links-in-count). | [optional] 
 **SourceLinksInCountMax** | **int32**| This parameter is used for finding stories from sources whose Links in count is less than or equal to the specified value. You can read more about working with Links in count here [https://newsapi.aylien.com/docs/working-with-links-in-count](https://newsapi.aylien.com/docs/working-with-links-in-count). | [optional] 
 **SourceRankingsAlexaRankMin** | **int32**| This parameter is used for finding stories from sources whose Alexa rank is greater than or equal to the specified value. You can read more about working with Alexa ranks here [https://newsapi.aylien.com/docs/working-with-alexa-ranks](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
 **SourceRankingsAlexaRankMax** | **int32**| This parameter is used for finding stories from sources whose Alexa rank is less than or equal to the specified value. You can read more about working with Alexa ranks here [https://newsapi.aylien.com/docs/working-with-alexa-ranks](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
 **SourceRankingsAlexaCountry** | [**[]string**](string.md)| This parameter is used for finding stories from sources whose Alexa rank is in the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. You can read more about working with Alexa ranks here [https://newsapi.aylien.com/docs/working-with-alexa-ranks](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
 **Cluster** | **bool**| This parameter enables clustering for the returned stories. | [optional] [default to false]
 **ClusterAlgorithm** | **string**| This parameter is used for specifying the clustering algorithm you wish to use. It supprts STC, Lingo and [k-means](https://en.wikipedia.org/wiki/K-means_clustering) algorithms. | [optional] [default to lingo]
 **Return** | [**[]string**](string.md)| This parameter is used for specifying return fields. | [optional] 
 **SortBy** | **string**| This parameter is used for changing the order column of the results. | [optional] [default to published_at]
 **SortDirection** | **string**| This parameter is used for changing the order direction of the result. | [optional] [default to desc]
 **Cursor** | **string**| This parameter is used for finding a specific page. | [optional] [default to *]
 **PerPage** | **int32**| This parameter is used for specifying number of items in each page. | [optional] [default to 10]

### Return type

[**Stories**](Stories.md)

### Authorization

[app_key](../README.md#app_key), [app_id](../README.md#app_id)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTimeSeries**
> TimeSeriesList ListTimeSeries(params *TimeSeriesParams)

List time series

This endpoint is used for getting time series by stories.

### Example

``` go
package main

// Import the library
import (
	"fmt"
	newsapi "github.com/AYLIEN/aylien_newsapi_go"
)

func main() {
	api := newsapi.NewDefaultApi()

	// Configure API key authorization: app_id
	api.Configuration.APIKeyPrefix["X-AYLIEN-NewsAPI-Application-ID"] = "YOUR APP ID"

	// Configure API key authorization: app_key
	api.Configuration.APIKeyPrefix["X-AYLIEN-NewsAPI-Application-Key"] = "YOUR APP KEY"

	id := []int64{112}
	title := "title_example"
	body := "body_example"
	text := "text_example"
	language := []string{"en"}
	categoriesTaxonomy := "categoriesTaxonomy_example"
	categoriesConfident := "categoriesConfident_example"
	categoriesId := []string{"categoriesId_example"}
	categoriesLevel := []int32{3}
	entitiesTitleText := []string{"entitiesTitleText_example"}
	entitiesTitleType := []string{"entitiesTitleType_example"}
	entitiesTitleLinksDbpedia := []string{"entitiesTitleLinksDbpedia_example"}
	entitiesBodyText := []string{"entitiesBodyText_example"}
	entitiesBodyType := []string{"entitiesBodyType_example"}
	entitiesBodyLinksDbpedia := []string{"entitiesBodyLinksDbpedia_example"}
	sentimentTitlePolarity := "sentimentTitlePolarity_example"
	sentimentBodyPolarity := "sentimentBodyPolarity_example"
	mediaImagesCountMin := 56
	mediaImagesCountMax := 56
	mediaVideosCountMin := 56
	mediaVideosCountMax := 56
	authorId := []int32{25}
	authorName := "authorName_example"
	sourceId := []int32{45}
	sourceName := []string{"sourceName_example"}
	sourceDomain := []string{"sourceDomain_example"}
	sourceLocationsCountry := []string{"sourceLocationsCountry_example"}
	sourceLocationsState := []string{"sourceLocationsState_example"}
	sourceLocationsCity := []string{"sourceLocationsCity_example"}
	sourceScopesCountry := []string{"sourceScopesCountry_example"}
	sourceScopesState := []string{"sourceScopesState_example"}
	sourceScopesCity := []string{"sourceScopesCity_example"}
	sourceScopesLevel := []string{"sourceScopesLevel_example"}
	sourceLinksInCountMin := 56
  sourceLinksInCountMax := 56
  sourceRankingsAlexaRankMin := 56
  sourceRankingsAlexaRankMax := 56
  sourceRankingsAlexaCountry := []string{"sourceRankingsAlexaCountry_example"}
	publishedAtStart := "publishedAtStart_example"
	publishedAtEnd := "publishedAtEnd_example"
	period := "+1DAY"

	timeSeriesParams := &newsapi.TimeSeriesParams{
		Id:                         id,
		Title:                      title,
		Body:                       body,
		Text:                       text,
		Language:                   language,
		CategoriesTaxonomy:         categoriesTaxonomy,
		CategoriesConfident:        categoriesConfident,
		CategoriesId:               categoriesId,
		CategoriesLevel:            categoriesLevel,
		EntitiesTitleText:          entitiesTitleText,
		EntitiesTitleType:          entitiesTitleType,
		EntitiesTitleLinksDbpedia:  entitiesTitleLinksDbpedia,
		EntitiesBodyText:           entitiesBodyText,
		EntitiesBodyType:           entitiesBodyType,
		EntitiesBodyLinksDbpedia:   entitiesBodyLinksDbpedia,
		SentimentTitlePolarity:     sentimentTitlePolarity,
		SentimentBodyPolarity:      sentimentBodyPolarity,
		MediaImagesCountMin:        mediaImagesCountMin,
		MediaImagesCountMax:        mediaImagesCountMax,
		MediaVideosCountMin:        mediaVideosCountMin,
		MediaVideosCountMin:        mediaVideosCountMin,
		AuthorId:                   authorId,
		AuthorName:                 authorName,
		SourceId:                   sourceId,
		SourceName:                 sourceName,
		SourceDomain:               sourceDomain,
		SourceLocationsCountry:     sourceLocationsCountry,
		SourceLocationsState:       sourceLocationsState,
		SourceLocationsCity:        sourceLocationsCity,
		SourceScopesCountry:        sourceScopesCountry,
		SourceScopesState:          sourceScopesState,
		SourceScopesCity:           sourceScopesCity,
		SourceScopesLevel:          sourceScopesLevel,
		SourceLinksInCountMin:      sourceLinksInCountMin,
		SourceLinksInCountMax:      sourceLinksInCountMax,
		SourceRankingsAlexaRankMin: sourceRankingsAlexaRankMin,
		SourceRankingsAlexaRankMax: sourceRankingsAlexaRankMax,
		SourceRankingsAlexaCountry: sourceRankingsAlexaCountry,
		PublishedAtStart:           publishedAtStart,
		PublishedAtEnd:             publishedAtEnd,
		Period:                     period}

	timeSeriesResponse, res, err := api.ListTimeSeries(timeSeriesParams)
	if err != nil {
		panic(err)
	}
	_ = res

	fmt.Println(timeSeriesResponse)
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | [**[]int64**](int64.md)| This parameter is used for finding stroies by story id. | [optional] 
 **Title** | **string**| This parameter is used for finding stories whose title contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Body** | **string**| This parameter is used for finding stories whose body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Text** | **string**| This parameter is used for finding stories whose title or body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Language** | [**[]string**](string.md)| This parameter is used for finding stories whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes. | [optional] 
 **CategoriesTaxonomy** | **string**| This parameter is used for defining the type of the taxonomy for the rest of the categories queries. | [optional] 
 **CategoriesConfident** | **bool**| This parameter is used for finding stories whose categories are confident. | [optional] [default to true]
 **CategoriesId** | [**[]string**](string.md)| This parameter is used for finding stories by categories id. | [optional] 
 **CategoriesLevel** | [**[]int32**](int32.md)| This parameter is used for finding stories by categories level. | [optional] 
 **EntitiesTitleText** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;text&#x60; in story titles. | [optional] 
 **EntitiesTitleType** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;type&#x60; in story titles. | [optional] 
 **EntitiesTitleLinksDbpedia** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities dbpedia URL in story titles. | [optional] 
 **EntitiesBodyText** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;text&#x60; in the body of stories. | [optional] 
 **EntitiesBodyType** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;type&#x60; in the body of stories. | [optional] 
 **EntitiesBodyLinksDbpedia** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities dbpedia URL in the body of stories. | [optional] 
 **SentimentTitlePolarity** | **string**| This parameter is used for finding stories whose title sentiment is the specified value. | [optional] 
 **SentimentBodyPolarity** | **string**| This parameter is used for finding stories whose body sentiment is the specified value. | [optional] 
 **MediaImagesCountMin** | **int32**| This parameter is used for finding stories whose number of images is greater than or equal to the specified value. | [optional] 
 **MediaImagesCountMax** | **int32**| This parameter is used for finding stories whose number of images is less than or equal to the specified value. | [optional] 
 **MediaVideosCountMin** | **int32**| This parameter is used for finding stories whose number of videos is greater than or equal to the specified value. | [optional] 
 **MediaVideosCountMax** | **int32**| This parameter is used for finding stories whose number of videos is less than or equal to the specified value. | [optional] 
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
 **SourceLinksInCountMin** | **int32**| This parameter is used for finding stories from sources whose Links in count is greater than or equal to the specified value. You can read more about working with Links in count here [https://newsapi.aylien.com/docs/working-with-links-in-count](https://newsapi.aylien.com/docs/working-with-links-in-count). | [optional] 
 **SourceLinksInCountMax** | **int32**| This parameter is used for finding stories from sources whose Links in count is less than or equal to the specified value. You can read more about working with Links in count here [https://newsapi.aylien.com/docs/working-with-links-in-count](https://newsapi.aylien.com/docs/working-with-links-in-count). | [optional] 
 **SourceRankingsAlexaRankMin** | **int32**| This parameter is used for finding stories from sources whose Alexa rank is greater than or equal to the specified value. You can read more about working with Alexa ranks here [https://newsapi.aylien.com/docs/working-with-alexa-ranks](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
 **SourceRankingsAlexaRankMax** | **int32**| This parameter is used for finding stories from sources whose Alexa rank is less than or equal to the specified value. You can read more about working with Alexa ranks here [https://newsapi.aylien.com/docs/working-with-alexa-ranks](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
 **SourceRankingsAlexaCountry** | [**[]string**](string.md)| This parameter is used for finding stories from sources whose Alexa rank is in the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. You can read more about working with Alexa ranks here [https://newsapi.aylien.com/docs/working-with-alexa-ranks](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
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
> Trends ListTrends(params *TrendsParams)

List trends

This endpoint is used for finding trends based on stories.

### Example

``` go
package main

// Import the library
import (
	"fmt"
	newsapi "github.com/AYLIEN/aylien_newsapi_go"
)

func main() {
	api := newsapi.NewDefaultApi()

	// Configure API key authorization: app_id
	api.Configuration.APIKeyPrefix["X-AYLIEN-NewsAPI-Application-ID"] = "YOUR APP ID"

	// Configure API key authorization: app_key
	api.Configuration.APIKeyPrefix["X-AYLIEN-NewsAPI-Application-Key"] = "YOUR APP KEY"

	id := []int64{112}
	title := "title_example"
	body := "body_example"
	text := "text_example"
	language := []string{"en"}
	publishedAtStart := "publishedAtStart_example"
	publishedAtEnd := "publishedAtEnd_example"
	categoriesTaxonomy := "categoriesTaxonomy_example"
	categoriesConfident := "categoriesConfident_example"
	categoriesId := []string{"categoriesId_example"}
	categoriesLevel := []int32{3}
	entitiesTitleText := []string{"entitiesTitleText_example"}
	entitiesTitleType := []string{"entitiesTitleType_example"}
	entitiesTitleLinksDbpedia := []string{"entitiesTitleLinksDbpedia_example"}
	entitiesBodyText := []string{"entitiesBodyText_example"}
	entitiesBodyType := []string{"entitiesBodyType_example"}
	entitiesBodyLinksDbpedia := []string{"entitiesBodyLinksDbpedia_example"}
	sentimentTitlePolarity := "sentimentTitlePolarity_example"
	sentimentBodyPolarity := "sentimentBodyPolarity_example"
	mediaImagesCountMin := 56
	mediaImagesCountMax := 56
	mediaVideosCountMin := 56
	mediaVideosCountMax := 56
	authorId := []int32{25}
	authorName := "authorName_example"
	sourceId := []int32{45}
	sourceName := []string{"sourceName_example"}
	sourceDomain := []string{"sourceDomain_example"}
	sourceLocationsCountry := []string{"sourceLocationsCountry_example"}
	sourceLocationsState := []string{"sourceLocationsState_example"}
	sourceLocationsCity := []string{"sourceLocationsCity_example"}
	sourceScopesCountry := []string{"sourceScopesCountry_example"}
	sourceScopesState := []string{"sourceScopesState_example"}
	sourceScopesCity := []string{"sourceScopesCity_example"}
	sourceScopesLevel := []string{"sourceScopesLevel_example"}
	sourceLinksInCountMin := 56
  sourceLinksInCountMax := 56
  sourceRankingsAlexaRankMin := 56
  sourceRankingsAlexaRankMax := 56
  sourceRankingsAlexaCountry := []string{"sourceRankingsAlexaCountry_example"}
	field := "keywords"

	trendsParams := &newsapi.TrendsParams{
		Id:                         id,
		Title:                      title,
		Body:                       body,
		Text:                       text,
		Language:                   language,
		PublishedAtStart:           publishedAtStart,
		PublishedAtEnd:             publishedAtEnd,
		CategoriesTaxonomy:         categoriesTaxonomy,
		CategoriesConfident:        categoriesConfident,
		CategoriesId:               categoriesId,
		CategoriesLevel:            categoriesLevel,
		EntitiesTitleText:          entitiesTitleText,
		EntitiesTitleType:          entitiesTitleType,
		EntitiesTitleLinksDbpedia:  entitiesTitleLinksDbpedia,
		EntitiesBodyText:           entitiesBodyText,
		EntitiesBodyType:           entitiesBodyType,
		EntitiesBodyLinksDbpedia:   entitiesBodyLinksDbpedia,
		SentimentTitlePolarity:     sentimentTitlePolarity,
		SentimentBodyPolarity:      sentimentBodyPolarity,
		MediaImagesCountMin:        mediaImagesCountMin,
		MediaImagesCountMax:        mediaImagesCountMax,
		MediaVideosCountMin:        mediaVideosCountMin,
		MediaVideosCountMin:        mediaVideosCountMin,
		AuthorId:                   authorId,
		AuthorName:                 authorName,
		SourceId:                   sourceId,
		SourceName:                 sourceName,
		SourceDomain:               sourceDomain,
		SourceLocationsCountry:     sourceLocationsCountry,
		SourceLocationsState:       sourceLocationsState,
		SourceLocationsCity:        sourceLocationsCity,
		SourceScopesCountry:        sourceScopesCountry,
		SourceScopesState:          sourceScopesState,
		SourceScopesCity:           sourceScopesCity,
		SourceScopesLevel:          sourceScopesLevel,
		SourceLinksInCountMin:      sourceLinksInCountMin,
		SourceLinksInCountMax:      sourceLinksInCountMax,
		SourceRankingsAlexaRankMin: sourceRankingsAlexaRankMin,
		SourceRankingsAlexaRankMax: sourceRankingsAlexaRankMax,
		SourceRankingsAlexaCountry: sourceRankingsAlexaCountry,
		Field:                      field}

	trendsResponse, res, err := api.ListTrends(trendsParams)
	if err != nil {
		panic(err)
	}
	_ = res

	fmt.Println(trendsResponse)
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | [**[]int64**](int64.md)| This parameter is used for finding stroies by story id. | [optional] 
 **Title** | **string**| This parameter is used for finding stories whose title contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Body** | **string**| This parameter is used for finding stories whose body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Text** | **string**| This parameter is used for finding stories whose title or body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators). | [optional] 
 **Language** | [**[]string**](string.md)| This parameter is used for finding stories whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes. | [optional] 
 **PublishedAtStart** | **string**| This parameter is used for finding stories whose published at time is greater than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates). | [optional] 
 **PublishedAtEnd** | **string**| This parameter is used for finding stories whose published at time is less than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates). | [optional] 
 **CategoriesTaxonomy** | **string**| This parameter is used for defining the type of the taxonomy for the rest of the categories queries. | [optional] 
 **CategoriesConfident** | **bool**| This parameter is used for finding stories whose categories are confident. | [optional] [default to true]
 **CategoriesId** | [**[]string**](string.md)| This parameter is used for finding stories by categories id. | [optional] 
 **CategoriesLevel** | [**[]int32**](int32.md)| This parameter is used for finding stories by categories level. | [optional] 
 **EntitiesTitleText** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;text&#x60; in story titles. | [optional] 
 **EntitiesTitleType** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;type&#x60; in story titles. | [optional] 
 **EntitiesTitleLinksDbpedia** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities dbpedia URL in story titles. | [optional] 
 **EntitiesBodyText** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;text&#x60; in the body of stories. | [optional] 
 **EntitiesBodyType** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities &#x60;type&#x60; in the body of stories. | [optional] 
 **EntitiesBodyLinksDbpedia** | [**[]string**](string.md)| This parameter is used to find stories based on the specified entities dbpedia URL in the body of stories. | [optional] 
 **SentimentTitlePolarity** | **string**| This parameter is used for finding stories whose title sentiment is the specified value. | [optional] 
 **SentimentBodyPolarity** | **string**| This parameter is used for finding stories whose body sentiment is the specified value. | [optional] 
 **MediaImagesCountMin** | **int32**| This parameter is used for finding stories whose number of images is greater than or equal to the specified value. | [optional] 
 **MediaImagesCountMax** | **int32**| This parameter is used for finding stories whose number of images is less than or equal to the specified value. | [optional] 
 **MediaVideosCountMin** | **int32**| This parameter is used for finding stories whose number of videos is greater than or equal to the specified value. | [optional] 
 **MediaVideosCountMax** | **int32**| This parameter is used for finding stories whose number of videos is less than or equal to the specified value. | [optional] 
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
 **SourceLinksInCountMin** | **int32**| This parameter is used for finding stories from sources whose Links in count is greater than or equal to the specified value. You can read more about working with Links in count here [https://newsapi.aylien.com/docs/working-with-links-in-count](https://newsapi.aylien.com/docs/working-with-links-in-count). | [optional] 
 **SourceLinksInCountMax** | **int32**| This parameter is used for finding stories from sources whose Links in count is less than or equal to the specified value. You can read more about working with Links in count here [https://newsapi.aylien.com/docs/working-with-links-in-count](https://newsapi.aylien.com/docs/working-with-links-in-count). | [optional] 
 **SourceRankingsAlexaRankMin** | **int32**| This parameter is used for finding stories from sources whose Alexa rank is greater than or equal to the specified value. You can read more about working with Alexa ranks here [https://newsapi.aylien.com/docs/working-with-alexa-ranks](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
 **SourceRankingsAlexaRankMax** | **int32**| This parameter is used for finding stories from sources whose Alexa rank is less than or equal to the specified value. You can read more about working with Alexa ranks here [https://newsapi.aylien.com/docs/working-with-alexa-ranks](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
 **SourceRankingsAlexaCountry** | [**[]string**](string.md)| This parameter is used for finding stories from sources whose Alexa rank is in the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. You can read more about working with Alexa ranks here [https://newsapi.aylien.com/docs/working-with-alexa-ranks](https://newsapi.aylien.com/docs/working-with-alexa-ranks). | [optional] 
 **Field** | **string**| This parameter is used to specify the trend field. | [optional] 

### Return type

[**Trends**](Trends.md)

### Authorization

[app_key](../README.md#app_key), [app_id](../README.md#app_id)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

