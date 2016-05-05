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
	term := "term_example"  // string | This parameter is used for finding autocomplete objects whose contains the specified value.
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
 **Term** | **string**| This parameter is used for finding autocomplete objects whose contains the specified value. | 
 **Language** | **string**| This parameter is used for autocompletes whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes. | [optional] [default to en]
 **PerPage** | **int32**| This parameter is used for specifying number of the items in each page. | [optional] [default to 3]

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

This endpoint is used for finding story coverages based on provided parameters. The number of coverages to return, up to a maximum of 100.

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

	id := []int64{112}                                                         // []int64 | This parameter is used for finding stroies by story id.
	title := "title_example"                                                   // string | This parameter is used for finding stories whose title contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
	body := "body_example"                                                     // string | This parameter is used for finding stories whose title contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
	text := "text_example"                                                     // string | This parameter is used for finding stories whose title or body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
	language := []string{"en"}                                                 // []string | This parameter is used for finding stories whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes.
	publishedAtStart := "publishedAtStart_example"                             // string | This parameter is used for finding stories whose published at is greater than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates).
	publishedAtEnd := "publishedAtEnd_example"                                 // string | This parameter is used for finding stories whose published at is less than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates).
	categoriesTaxonomy := "categoriesTaxonomy_example"                         // string | This parameter is used for defining type of the taxonomy for the rest of categories queries.
	categoriesConfident := "categoriesConfident_example"                       // string | This parameter is used for finding stories whose categories are confident.
	categoriesId := []string{"categoriesId_example"}                           // []string | This parameter is used for finding stories whose categories id is the specified value.
	categoriesLevel := []int32{3}                                              // []int32 | This parameter is used for finding stories whose categories level is the specified value.
	entitiesTitleText := []string{"entitiesTitleText_example"}                 // []string | This parameter is used for finding stories whose entities text in title is the specified value.
	entitiesTitleType := []string{"entitiesTitleType_example"}                 // []string| This parameter is used for finding stories whose entities type in title is the specified value.
	entitiesTitleLinksDbpedia := []string{"entitiesTitleLinksDbpedia_example"} // []string | This parameter is used for finding stories whose entities dbpedia url in title is the specified value.
	entitiesBodyText := []string{"entitiesBodyText_example"}                   // []string | This parameter is used for finding stories whose entities text in body is the specified value.
	entitiesBodyType := []string{"entitiesBodyType_example"}                   // []string | This parameter is used for finding stories whose entities type in body is the specified value.
	entitiesBodyLinksDbpedia := []string{"entitiesBodyLinksDbpedia_example"}   // []string | This parameter is used for finding stories whose entities dbpedia url in body is the specified value.
	sentimentTitlePolarity := "sentimentTitlePolarity_example"                 // string | This parameter is used for finding stories whose title sentiment is the specified value.
	sentimentBodyPolarity := "sentimentBodyPolarity_example"                   // string | This parameter is used for finding stories whose body sentiment is the specified value.
	authorId := []int32{25}                                                    // []int32 | This parameter is used for finding stories whose author id is the specified value.
	authorName := "authorName_example"                                         // string | This parameter is used for finding stories whose author full name contains the specified value.
	sourceId := []int32{45}                                                    // []int32 | This parameter is used for finding stories whose source id is the specified value.
	sourceName := []string{"sourceName_example"}                               // []string | This parameter is used for finding stories whose source name contains the specified value.
	sourceDomain := []string{"sourceDomain_example"}                           // []string | This parameter is used for finding stories whose source domain is the specified value.
	sourceLocationsCountry := []string{"sourceLocationsCountry_example"}       // []string | This parameter is used for finding stories whose source country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes.
	sourceLocationsState := []string{"sourceLocationsState_example"}           // []string | This parameter is used for finding stories whose source state/province is the specified value.
	sourceLocationsCity := []string{"sourceLocationsCity_example"}             // []string | This parameter is used for finding stories whose source city is the specified value.
	sourceScopesCountry := []string{"sourceScopesCountry_example"}             // []string | This parameter is used for finding stories whose source scopes country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes.
	sourceScopesState := []string{"sourceScopesState_example"}                 // []string | This parameter is used for finding stories whose source scopes state/province is the specified value.
	sourceScopesCity := []string{"sourceScopesCity_example"}                   // []string | This parameter is used for finding stories whose source scopes city is the specified value.
	sourceScopesLevel := []string{"sourceScopesLevel_example"}                 // []string | This parameter is used for finding stories whose source scopes level is the specified value.
	_return := []string{"_return_example"}                                     // []string | This parameter is used for specifying return fields.
	storyId := 789                                                             // int64 | A story id
	storyUrl := "storyUrl_example"                                             // string | An article or webpage
	storyTitle := "storyTitle_example"                                         //string | Title of the article
	storyBody := "storyBody_example"                                           // string | Body of the article
	storyPublishedAt := time.Now()                                             // time.Time | Publish date of the article. If you use url or title and body for getting coverages, this parameter is required. The format used is a restricted form of the canonical representation of dateTime in the [XML Schema specification (ISO 8601)](https://www.w3.org/TR/xmlschema-2/#dateTime).
	storyLanguage := "auto"                                                    // string | This parameter is used for setting language of the story. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes.
	perPage := 3                                                               // int32 | This parameter is used for specifying number of the items in each page.

	coveragesParams := &newsapi.CoveragesParams{
		Id:                        id,
		Title:                     title,
		Body:                      body,
		Text:                      text,
		Language:                  language,
		PublishedAtStart:          publishedAtStart,
		PublishedAtEnd:            publishedAtEnd,
		CategoriesTaxonomy:        categoriesTaxonomy,
		CategoriesConfident:       categoriesConfident,
		CategoriesId:              categoriesId,
		CategoriesLevel:           categoriesLevel,
		EntitiesTitleText:         entitiesTitleText,
		EntitiesTitleType:         entitiesTitleType,
		EntitiesTitleLinksDbpedia: entitiesTitleLinksDbpedia,
		EntitiesBodyText:          entitiesBodyText,
		EntitiesBodyType:          entitiesBodyType,
		EntitiesBodyLinksDbpedia:  entitiesBodyLinksDbpedia,
		SentimentTitlePolarity:    sentimentTitlePolarity,
		SentimentBodyPolarity:     sentimentBodyPolarity,
		AuthorId:                  authorId,
		AuthorName:                authorName,
		SourceId:                  sourceId,
		SourceName:                sourceName,
		SourceDomain:              sourceDomain,
		SourceLocationsCountry:    sourceLocationsCountry,
		SourceLocationsState:      sourceLocationsState,
		SourceLocationsCity:       sourceLocationsCity,
		SourceScopesCountry:       sourceScopesCountry,
		SourceScopesState:         sourceScopesState,
		SourceScopesCity:          sourceScopesCity,
		SourceScopesLevel:         sourceScopesLevel,
		Return:                    _return,
		StoryId:                   storyId,
		StoryUrl:                  storyUrl,
		StoryTitle:                storyTitle,
		StoryBody:                 storyBody,
		StoryPublishedAt:          storyPublishedAt,
		StoryLanguage:             storyLanguage,
		PerPage:                   perPage}

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
 **PublishedAtStart** | **string**| This parameter is used for finding stories whose published at is greater than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates). | [optional] 
 **PublishedAtEnd** | **string**| This parameter is used for finding stories whose published at is less than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates). | [optional] 
 **CategoriesTaxonomy** | **string**| This parameter is used for defining type of the taxonomy for the rest of categories queries. | [optional] 
 **CategoriesConfident** | **bool**| This parameter is used for finding stories whose categories are confident. | [optional] [default to true]
 **CategoriesId** | [**[]string**](string.md)| This parameter is used for finding stories whose categories id is the specified value. | [optional] 
 **CategoriesLevel** | [**[]int32**](int32.md)| This parameter is used for finding stories whose categories level is the specified value. | [optional] 
 **EntitiesTitleText** | [**[]string**](string.md)| This parameter is used for finding stories whose entities text in title is the specified value. | [optional] 
 **EntitiesTitleType** | [**[]string**](string.md)| This parameter is used for finding stories whose entities type in title is the specified value. | [optional] 
 **EntitiesTitleLinksDbpedia** | [**[]string**](string.md)| This parameter is used for finding stories whose entities dbpedia url in title is the specified value. | [optional] 
 **EntitiesBodyText** | [**[]string**](string.md)| This parameter is used for finding stories whose entities text in body is the specified value. | [optional] 
 **EntitiesBodyType** | [**[]string**](string.md)| This parameter is used for finding stories whose entities type in body is the specified value. | [optional] 
 **EntitiesBodyLinksDbpedia** | [**[]string**](string.md)| This parameter is used for finding stories whose entities dbpedia url in body is the specified value. | [optional] 
 **SentimentTitlePolarity** | **string**| This parameter is used for finding stories whose title sentiment is the specified value. | [optional] 
 **SentimentBodyPolarity** | **string**| This parameter is used for finding stories whose body sentiment is the specified value. | [optional] 
 **AuthorId** | [**[]int32**](int32.md)| This parameter is used for finding stories whose author id is the specified value. | [optional] 
 **AuthorName** | **string**| This parameter is used for finding stories whose author full name contains the specified value. | [optional] 
 **SourceId** | [**[]int32**](int32.md)| This parameter is used for finding stories whose source id is the specified value. | [optional] 
 **SourceName** | [**[]string**](string.md)| This parameter is used for finding stories whose source name contains the specified value. | [optional] 
 **SourceDomain** | [**[]string**](string.md)| This parameter is used for finding stories whose source domain is the specified value. | [optional] 
 **SourceLocationsCountry** | [**[]string**](string.md)| This parameter is used for finding stories whose source country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. | [optional] 
 **SourceLocationsState** | [**[]string**](string.md)| This parameter is used for finding stories whose source state/province is the specified value. | [optional] 
 **SourceLocationsCity** | [**[]string**](string.md)| This parameter is used for finding stories whose source city is the specified value. | [optional] 
 **SourceScopesCountry** | [**[]string**](string.md)| This parameter is used for finding stories whose source scopes country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. | [optional] 
 **SourceScopesState** | [**[]string**](string.md)| This parameter is used for finding stories whose source scopes state/province is the specified value. | [optional] 
 **SourceScopesCity** | [**[]string**](string.md)| This parameter is used for finding stories whose source scopes city is the specified value. | [optional] 
 **SourceScopesLevel** | [**[]string**](string.md)| This parameter is used for finding stories whose source scopes level is the specified value. | [optional] 
 **Return** | [**[]string**](string.md)| This parameter is used for specifying return fields. | [optional] 
 **StoryId** | **int64**| A story id | [optional] 
 **StoryUrl** | **string**| An article or webpage | [optional] 
 **StoryTitle** | **string**| Title of the article | [optional] 
 **StoryBody** | **string**| Body of the article | [optional] 
 **StoryPublishedAt** | **time.Time**| Publish date of the article. If you use url or title and body for getting coverages, this parameter is required. The format used is a restricted form of the canonical representation of dateTime in the [XML Schema specification (ISO 8601)](https://www.w3.org/TR/xmlschema-2/#dateTime). | [optional] 
 **StoryLanguage** | **string**| This parameter is used for setting language of the story. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes. | [optional] [default to auto]
 **PerPage** | **int32**| This parameter is used for specifying number of the items in each page. | [optional] [default to 3]

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

This endpoint is used for getting histograms based on the field parameter is passed.

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

	id := []int64{112}                                                         // []int64 | This parameter is used for finding stroies by story id.
	title := "title_example"                                                   // string | This parameter is used for finding stories whose title contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
	body := "body_example"                                                     // string | This parameter is used for finding stories whose title contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
	text := "text_example"                                                     // string | This parameter is used for finding stories whose title or body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
	language := []string{"en"}                                                 // []string | This parameter is used for finding stories whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes.
	publishedAtStart := "publishedAtStart_example"                             // string | This parameter is used for finding stories whose published at is greater than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates).
	publishedAtEnd := "publishedAtEnd_example"                                 // string | This parameter is used for finding stories whose published at is less than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates).
	categoriesTaxonomy := "categoriesTaxonomy_example"                         // string | This parameter is used for defining type of the taxonomy for the rest of categories queries.
	categoriesConfident := "categoriesConfident_example"                       // string | This parameter is used for finding stories whose categories are confident.
	categoriesId := []string{"categoriesId_example"}                           // []string | This parameter is used for finding stories whose categories id is the specified value.
	categoriesLevel := []int32{3}                                              // []int32 | This parameter is used for finding stories whose categories level is the specified value.
	entitiesTitleText := []string{"entitiesTitleText_example"}                 // []string | This parameter is used for finding stories whose entities text in title is the specified value.
	entitiesTitleType := []string{"entitiesTitleType_example"}                 // []string| This parameter is used for finding stories whose entities type in title is the specified value.
	entitiesTitleLinksDbpedia := []string{"entitiesTitleLinksDbpedia_example"} // []string | This parameter is used for finding stories whose entities dbpedia url in title is the specified value.
	entitiesBodyText := []string{"entitiesBodyText_example"}                   // []string | This parameter is used for finding stories whose entities text in body is the specified value.
	entitiesBodyType := []string{"entitiesBodyType_example"}                   // []string | This parameter is used for finding stories whose entities type in body is the specified value.
	entitiesBodyLinksDbpedia := []string{"entitiesBodyLinksDbpedia_example"}   // []string | This parameter is used for finding stories whose entities dbpedia url in body is the specified value.
	sentimentTitlePolarity := "sentimentTitlePolarity_example"                 // string | This parameter is used for finding stories whose title sentiment is the specified value.
	sentimentBodyPolarity := "sentimentBodyPolarity_example"                   // string | This parameter is used for finding stories whose body sentiment is the specified value.
	authorId := []int32{25}                                                    // []int32 | This parameter is used for finding stories whose author id is the specified value.
	authorName := "authorName_example"                                         // string | This parameter is used for finding stories whose author full name contains the specified value.
	sourceId := []int32{45}                                                    // []int32 | This parameter is used for finding stories whose source id is the specified value.
	sourceName := []string{"sourceName_example"}                               // []string | This parameter is used for finding stories whose source name contains the specified value.
	sourceDomain := []string{"sourceDomain_example"}                           // []string | This parameter is used for finding stories whose source domain is the specified value.
	sourceLocationsCountry := []string{"sourceLocationsCountry_example"}       // []string | This parameter is used for finding stories whose source country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes.
	sourceLocationsState := []string{"sourceLocationsState_example"}           // []string | This parameter is used for finding stories whose source state/province is the specified value.
	sourceLocationsCity := []string{"sourceLocationsCity_example"}             // []string | This parameter is used for finding stories whose source city is the specified value.
	sourceScopesCountry := []string{"sourceScopesCountry_example"}             // []string | This parameter is used for finding stories whose source scopes country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes.
	sourceScopesState := []string{"sourceScopesState_example"}                 // []string | This parameter is used for finding stories whose source scopes state/province is the specified value.
	sourceScopesCity := []string{"sourceScopesCity_example"}                   // []string | This parameter is used for finding stories whose source scopes city is the specified value.
	sourceScopesLevel := []string{"sourceScopesLevel_example"}                 // []string | This parameter is used for finding stories whose source scopes level is the specified value.
	intervalStart := 56                                                        // int32 | This parameter is used for setting the start data point of histogram intervals.
	intervalEnd := 22                                                          // int32 | This parameter is used for setting the end data point of histogram intervals.
	intervalWidth := 20                                                        // int32 | This parameter is used for setting the width of histogram intervals.
	field := "social_shares_count"                                             // string | This parameter is used for specifying the y-axis variable for the histogram.

	histogramsParams := &newsapi.HistogramsParams{
		Id:                        id,
		Title:                     title,
		Body:                      body,
		Text:                      text,
		Language:                  language,
		PublishedAtStart:          publishedAtStart,
		PublishedAtEnd:            publishedAtEnd,
		CategoriesTaxonomy:        categoriesTaxonomy,
		CategoriesConfident:       categoriesConfident,
		CategoriesId:              categoriesId,
		CategoriesLevel:           categoriesLevel,
		EntitiesTitleText:         entitiesTitleText,
		EntitiesTitleType:         entitiesTitleType,
		EntitiesTitleLinksDbpedia: entitiesTitleLinksDbpedia,
		EntitiesBodyText:          entitiesBodyText,
		EntitiesBodyType:          entitiesBodyType,
		EntitiesBodyLinksDbpedia:  entitiesBodyLinksDbpedia,
		SentimentTitlePolarity:    sentimentTitlePolarity,
		SentimentBodyPolarity:     sentimentBodyPolarity,
		AuthorId:                  authorId,
		AuthorName:                authorName,
		SourceId:                  sourceId,
		SourceName:                sourceName,
		SourceDomain:              sourceDomain,
		SourceLocationsCountry:    sourceLocationsCountry,
		SourceLocationsState:      sourceLocationsState,
		SourceLocationsCity:       sourceLocationsCity,
		SourceScopesCountry:       sourceScopesCountry,
		SourceScopesState:         sourceScopesState,
		SourceScopesCity:          sourceScopesCity,
		SourceScopesLevel:         sourceScopesLevel,
		IntervalStart:             intervalStart,
		IntervalEnd:               intervalEnd,
		IntervalWidth:             intervalWidth,
		Field:                     field}

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
 **PublishedAtStart** | **string**| This parameter is used for finding stories whose published at is greater than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates). | [optional] 
 **PublishedAtEnd** | **string**| This parameter is used for finding stories whose published at is less than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates). | [optional] 
 **CategoriesTaxonomy** | **string**| This parameter is used for defining type of the taxonomy for the rest of categories queries. | [optional] 
 **CategoriesConfident** | **bool**| This parameter is used for finding stories whose categories are confident. | [optional] [default to true]
 **CategoriesId** | [**[]string**](string.md)| This parameter is used for finding stories whose categories id is the specified value. | [optional] 
 **CategoriesLevel** | [**[]int32**](int32.md)| This parameter is used for finding stories whose categories level is the specified value. | [optional] 
 **EntitiesTitleText** | [**[]string**](string.md)| This parameter is used for finding stories whose entities text in title is the specified value. | [optional] 
 **EntitiesTitleType** | [**[]string**](string.md)| This parameter is used for finding stories whose entities type in title is the specified value. | [optional] 
 **EntitiesTitleLinksDbpedia** | [**[]string**](string.md)| This parameter is used for finding stories whose entities dbpedia url in title is the specified value. | [optional] 
 **EntitiesBodyText** | [**[]string**](string.md)| This parameter is used for finding stories whose entities text in body is the specified value. | [optional] 
 **EntitiesBodyType** | [**[]string**](string.md)| This parameter is used for finding stories whose entities type in body is the specified value. | [optional] 
 **EntitiesBodyLinksDbpedia** | [**[]string**](string.md)| This parameter is used for finding stories whose entities dbpedia url in body is the specified value. | [optional] 
 **SentimentTitlePolarity** | **string**| This parameter is used for finding stories whose title sentiment is the specified value. | [optional] 
 **SentimentBodyPolarity** | **string**| This parameter is used for finding stories whose body sentiment is the specified value. | [optional] 
 **AuthorId** | [**[]int32**](int32.md)| This parameter is used for finding stories whose author id is the specified value. | [optional] 
 **AuthorName** | **string**| This parameter is used for finding stories whose author full name contains the specified value. | [optional] 
 **SourceId** | [**[]int32**](int32.md)| This parameter is used for finding stories whose source id is the specified value. | [optional] 
 **SourceName** | [**[]string**](string.md)| This parameter is used for finding stories whose source name contains the specified value. | [optional] 
 **SourceDomain** | [**[]string**](string.md)| This parameter is used for finding stories whose source domain is the specified value. | [optional] 
 **SourceLocationsCountry** | [**[]string**](string.md)| This parameter is used for finding stories whose source country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. | [optional] 
 **SourceLocationsState** | [**[]string**](string.md)| This parameter is used for finding stories whose source state/province is the specified value. | [optional] 
 **SourceLocationsCity** | [**[]string**](string.md)| This parameter is used for finding stories whose source city is the specified value. | [optional] 
 **SourceScopesCountry** | [**[]string**](string.md)| This parameter is used for finding stories whose source scopes country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. | [optional] 
 **SourceScopesState** | [**[]string**](string.md)| This parameter is used for finding stories whose source scopes state/province is the specified value. | [optional] 
 **SourceScopesCity** | [**[]string**](string.md)| This parameter is used for finding stories whose source scopes city is the specified value. | [optional] 
 **SourceScopesLevel** | [**[]string**](string.md)| This parameter is used for finding stories whose source scopes level is the specified value. | [optional] 
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

This endpoint is used for finding related stories based on provided parameters. The number of related stories to return, up to a maximum of 100.

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

	id := []int64{112}                                                         // []int64 | This parameter is used for finding stroies by story id.
	title := "title_example"                                                   // string | This parameter is used for finding stories whose title contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
	body := "body_example"                                                     // string | This parameter is used for finding stories whose title contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
	text := "text_example"                                                     // string | This parameter is used for finding stories whose title or body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
	language := []string{"en"}                                                 // []string | This parameter is used for finding stories whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes.
	publishedAtStart := "publishedAtStart_example"                             // string | This parameter is used for finding stories whose published at is greater than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates).
	publishedAtEnd := "publishedAtEnd_example"                                 // string | This parameter is used for finding stories whose published at is less than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates).
	categoriesTaxonomy := "categoriesTaxonomy_example"                         // string | This parameter is used for defining type of the taxonomy for the rest of categories queries.
	categoriesConfident := "categoriesConfident_example"                       // string | This parameter is used for finding stories whose categories are confident.
	categoriesId := []string{"categoriesId_example"}                           // []string | This parameter is used for finding stories whose categories id is the specified value.
	categoriesLevel := []int32{3}                                              // []int32 | This parameter is used for finding stories whose categories level is the specified value.
	entitiesTitleText := []string{"entitiesTitleText_example"}                 // []string | This parameter is used for finding stories whose entities text in title is the specified value.
	entitiesTitleType := []string{"entitiesTitleType_example"}                 // []string| This parameter is used for finding stories whose entities type in title is the specified value.
	entitiesTitleLinksDbpedia := []string{"entitiesTitleLinksDbpedia_example"} // []string | This parameter is used for finding stories whose entities dbpedia url in title is the specified value.
	entitiesBodyText := []string{"entitiesBodyText_example"}                   // []string | This parameter is used for finding stories whose entities text in body is the specified value.
	entitiesBodyType := []string{"entitiesBodyType_example"}                   // []string | This parameter is used for finding stories whose entities type in body is the specified value.
	entitiesBodyLinksDbpedia := []string{"entitiesBodyLinksDbpedia_example"}   // []string | This parameter is used for finding stories whose entities dbpedia url in body is the specified value.
	sentimentTitlePolarity := "sentimentTitlePolarity_example"                 // string | This parameter is used for finding stories whose title sentiment is the specified value.
	sentimentBodyPolarity := "sentimentBodyPolarity_example"                   // string | This parameter is used for finding stories whose body sentiment is the specified value.
	authorId := []int32{25}                                                    // []int32 | This parameter is used for finding stories whose author id is the specified value.
	authorName := "authorName_example"                                         // string | This parameter is used for finding stories whose author full name contains the specified value.
	sourceId := []int32{45}                                                    // []int32 | This parameter is used for finding stories whose source id is the specified value.
	sourceName := []string{"sourceName_example"}                               // []string | This parameter is used for finding stories whose source name contains the specified value.
	sourceDomain := []string{"sourceDomain_example"}                           // []string | This parameter is used for finding stories whose source domain is the specified value.
	sourceLocationsCountry := []string{"sourceLocationsCountry_example"}       // []string | This parameter is used for finding stories whose source country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes.
	sourceLocationsState := []string{"sourceLocationsState_example"}           // []string | This parameter is used for finding stories whose source state/province is the specified value.
	sourceLocationsCity := []string{"sourceLocationsCity_example"}             // []string | This parameter is used for finding stories whose source city is the specified value.
	sourceScopesCountry := []string{"sourceScopesCountry_example"}             // []string | This parameter is used for finding stories whose source scopes country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes.
	sourceScopesState := []string{"sourceScopesState_example"}                 // []string | This parameter is used for finding stories whose source scopes state/province is the specified value.
	sourceScopesCity := []string{"sourceScopesCity_example"}                   // []string | This parameter is used for finding stories whose source scopes city is the specified value.
	sourceScopesLevel := []string{"sourceScopesLevel_example"}                 // []string | This parameter is used for finding stories whose source scopes level is the specified value.
	_return := []string{"_return_example"}                                     // []string | This parameter is used for specifying return fields.
	storyId := 789                                                             // int64 | A story id
	storyUrl := "storyUrl_example"                                             // string | An article or webpage
	storyTitle := "storyTitle_example"                                         //string | Title of the article
	storyBody := "storyBody_example"                                           // string | Body of the article
	boostBy := "boostBy_example"                                               // string | This parameter is used for boosting result by the specified value.
	storyLanguage := "auto"                                                    // string | This parameter is used for setting language of the story. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes.
	perPage := 3                                                               // int32 | This parameter is used for specifying number of the items in each page.

	relatedStoriesParams := &newsapi.RelatedStoriesParams{
		Id:                        id,
		Title:                     title,
		Body:                      body,
		Text:                      text,
		Language:                  language,
		PublishedAtStart:          publishedAtStart,
		PublishedAtEnd:            publishedAtEnd,
		CategoriesTaxonomy:        categoriesTaxonomy,
		CategoriesConfident:       categoriesConfident,
		CategoriesId:              categoriesId,
		CategoriesLevel:           categoriesLevel,
		EntitiesTitleText:         entitiesTitleText,
		EntitiesTitleType:         entitiesTitleType,
		EntitiesTitleLinksDbpedia: entitiesTitleLinksDbpedia,
		EntitiesBodyText:          entitiesBodyText,
		EntitiesBodyType:          entitiesBodyType,
		EntitiesBodyLinksDbpedia:  entitiesBodyLinksDbpedia,
		SentimentTitlePolarity:    sentimentTitlePolarity,
		SentimentBodyPolarity:     sentimentBodyPolarity,
		AuthorId:                  authorId,
		AuthorName:                authorName,
		SourceId:                  sourceId,
		SourceName:                sourceName,
		SourceDomain:              sourceDomain,
		SourceLocationsCountry:    sourceLocationsCountry,
		SourceLocationsState:      sourceLocationsState,
		SourceLocationsCity:       sourceLocationsCity,
		SourceScopesCountry:       sourceScopesCountry,
		SourceScopesState:         sourceScopesState,
		SourceScopesCity:          sourceScopesCity,
		SourceScopesLevel:         sourceScopesLevel,
		Return:                    _return,
		StoryId:                   storyId,
		StoryUrl:                  storyUrl,
		StoryTitle:                storyTitle,
		StoryBody:                 storyBody,
		BoostBy:                   boostBy,
		StoryLanguage:             storyLanguage,
		PerPage:                   perPage}

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
 **PublishedAtStart** | **string**| This parameter is used for finding stories whose published at is greater than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates). | [optional] 
 **PublishedAtEnd** | **string**| This parameter is used for finding stories whose published at is less than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates). | [optional] 
 **CategoriesTaxonomy** | **string**| This parameter is used for defining type of the taxonomy for the rest of categories queries. | [optional] 
 **CategoriesConfident** | **bool**| This parameter is used for finding stories whose categories are confident. | [optional] [default to true]
 **CategoriesId** | [**[]string**](string.md)| This parameter is used for finding stories whose categories id is the specified value. | [optional] 
 **CategoriesLevel** | [**[]int32**](int32.md)| This parameter is used for finding stories whose categories level is the specified value. | [optional] 
 **EntitiesTitleText** | [**[]string**](string.md)| This parameter is used for finding stories whose entities text in title is the specified value. | [optional] 
 **EntitiesTitleType** | [**[]string**](string.md)| This parameter is used for finding stories whose entities type in title is the specified value. | [optional] 
 **EntitiesTitleLinksDbpedia** | [**[]string**](string.md)| This parameter is used for finding stories whose entities dbpedia url in title is the specified value. | [optional] 
 **EntitiesBodyText** | [**[]string**](string.md)| This parameter is used for finding stories whose entities text in body is the specified value. | [optional] 
 **EntitiesBodyType** | [**[]string**](string.md)| This parameter is used for finding stories whose entities type in body is the specified value. | [optional] 
 **EntitiesBodyLinksDbpedia** | [**[]string**](string.md)| This parameter is used for finding stories whose entities dbpedia url in body is the specified value. | [optional] 
 **SentimentTitlePolarity** | **string**| This parameter is used for finding stories whose title sentiment is the specified value. | [optional] 
 **SentimentBodyPolarity** | **string**| This parameter is used for finding stories whose body sentiment is the specified value. | [optional] 
 **AuthorId** | [**[]int32**](int32.md)| This parameter is used for finding stories whose author id is the specified value. | [optional] 
 **AuthorName** | **string**| This parameter is used for finding stories whose author full name contains the specified value. | [optional] 
 **SourceId** | [**[]int32**](int32.md)| This parameter is used for finding stories whose source id is the specified value. | [optional] 
 **SourceName** | [**[]string**](string.md)| This parameter is used for finding stories whose source name contains the specified value. | [optional] 
 **SourceDomain** | [**[]string**](string.md)| This parameter is used for finding stories whose source domain is the specified value. | [optional] 
 **SourceLocationsCountry** | [**[]string**](string.md)| This parameter is used for finding stories whose source country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. | [optional] 
 **SourceLocationsState** | [**[]string**](string.md)| This parameter is used for finding stories whose source state/province is the specified value. | [optional] 
 **SourceLocationsCity** | [**[]string**](string.md)| This parameter is used for finding stories whose source city is the specified value. | [optional] 
 **SourceScopesCountry** | [**[]string**](string.md)| This parameter is used for finding stories whose source scopes country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. | [optional] 
 **SourceScopesState** | [**[]string**](string.md)| This parameter is used for finding stories whose source scopes state/province is the specified value. | [optional] 
 **SourceScopesCity** | [**[]string**](string.md)| This parameter is used for finding stories whose source scopes city is the specified value. | [optional] 
 **sourceScopesLevel** | [**[]string**](string.md)| This parameter is used for finding stories whose source scopes level is the specified value. | [optional] 
 **Return** | [**[]string**](string.md)| This parameter is used for specifying return fields. | [optional] 
 **StoryId** | **int64**| A story id | [optional] 
 **StoryUrl** | **string**| An article or webpage | [optional] 
 **StoryTitle** | **string**| Title of the article | [optional] 
 **StoryBody** | **string**| Body of the article | [optional] 
 **BoostBy** | **string**| This parameter is used for boosting result by the specified value. | [optional] 
 **StoryLanguage** | **string**| This parameter is used for setting language of the story. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes. | [optional] [default to auto]
 **PerPage** | **int32**| This parameter is used for specifying number of the items in each page. | [optional] [default to 3]

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

This endpoint is used for getting list of stories.

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

	id := []int64{112}                                                         // []int64 | This parameter is used for finding stroies by story id.
	title := "title_example"                                                   // string | This parameter is used for finding stories whose title contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
	body := "body_example"                                                     // string | This parameter is used for finding stories whose title contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
	text := "text_example"                                                     // string | This parameter is used for finding stories whose title or body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
	language := []string{"en"}                                                 // []string | This parameter is used for finding stories whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes.
	publishedAtStart := "publishedAtStart_example"                             // string | This parameter is used for finding stories whose published at is greater than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates).
	publishedAtEnd := "publishedAtEnd_example"                                 // string | This parameter is used for finding stories whose published at is less than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates).
	categoriesTaxonomy := "categoriesTaxonomy_example"                         // string | This parameter is used for defining type of the taxonomy for the rest of categories queries.
	categoriesConfident := "categoriesConfident_example"                       // string | This parameter is used for finding stories whose categories are confident.
	categoriesId := []string{"categoriesId_example"}                           // []string | This parameter is used for finding stories whose categories id is the specified value.
	categoriesLevel := []int32{3}                                              // []int32 | This parameter is used for finding stories whose categories level is the specified value.
	entitiesTitleText := []string{"entitiesTitleText_example"}                 // []string | This parameter is used for finding stories whose entities text in title is the specified value.
	entitiesTitleType := []string{"entitiesTitleType_example"}                 // []string| This parameter is used for finding stories whose entities type in title is the specified value.
	entitiesTitleLinksDbpedia := []string{"entitiesTitleLinksDbpedia_example"} // []string | This parameter is used for finding stories whose entities dbpedia url in title is the specified value.
	entitiesBodyText := []string{"entitiesBodyText_example"}                   // []string | This parameter is used for finding stories whose entities text in body is the specified value.
	entitiesBodyType := []string{"entitiesBodyType_example"}                   // []string | This parameter is used for finding stories whose entities type in body is the specified value.
	entitiesBodyLinksDbpedia := []string{"entitiesBodyLinksDbpedia_example"}   // []string | This parameter is used for finding stories whose entities dbpedia url in body is the specified value.
	sentimentTitlePolarity := "sentimentTitlePolarity_example"                 // string | This parameter is used for finding stories whose title sentiment is the specified value.
	sentimentBodyPolarity := "sentimentBodyPolarity_example"                   // string | This parameter is used for finding stories whose body sentiment is the specified value.
	authorId := []int32{25}                                                    // []int32 | This parameter is used for finding stories whose author id is the specified value.
	authorName := "authorName_example"                                         // string | This parameter is used for finding stories whose author full name contains the specified value.
	sourceId := []int32{45}                                                    // []int32 | This parameter is used for finding stories whose source id is the specified value.
	sourceName := []string{"sourceName_example"}                               // []string | This parameter is used for finding stories whose source name contains the specified value.
	sourceDomain := []string{"sourceDomain_example"}                           // []string | This parameter is used for finding stories whose source domain is the specified value.
	sourceLocationsCountry := []string{"sourceLocationsCountry_example"}       // []string | This parameter is used for finding stories whose source country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes.
	sourceLocationsState := []string{"sourceLocationsState_example"}           // []string | This parameter is used for finding stories whose source state/province is the specified value.
	sourceLocationsCity := []string{"sourceLocationsCity_example"}             // []string | This parameter is used for finding stories whose source city is the specified value.
	sourceScopesCountry := []string{"sourceScopesCountry_example"}             // []string | This parameter is used for finding stories whose source scopes country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes.
	sourceScopesState := []string{"sourceScopesState_example"}                 // []string | This parameter is used for finding stories whose source scopes state/province is the specified value.
	sourceScopesCity := []string{"sourceScopesCity_example"}                   // []string | This parameter is used for finding stories whose source scopes city is the specified value.
	sourceScopesLevel := []string{"sourceScopesLevel_example"}                 // []string | This parameter is used for finding stories whose source scopes level is the specified value.
	cluster := "false"                                                         // string | This parameter enables clustering for the returned stories.
	clusterAlgorithm := "lingo"                                                // string | This parameter is used for specifying the clustering algorithm. It supprts STC, Lingo and [k-means](https://en.wikipedia.org/wiki/K-means_clustering) algorithms.
	_return := []string{"_return_example"}                                     // []string | This parameter is used for specifying return fields.
	sortBy := "published_at"                                                   // string | This parameter is used for changing the order column of the result.
	sortDirection := "desc"                                                    // string | This parameter is used for changing the order direction of the result.
	cursor := "*"                                                              // string | This parameter is used for finding a specific page.
	perPage := 10                                                              // int32 | This parameter is used for specifying number of the items in each page.

	storiesParams := &newsapi.StoriesParams{
		Id:                        id,
		Title:                     title,
		Body:                      body,
		Text:                      text,
		Language:                  language,
		PublishedAtStart:          publishedAtStart,
		PublishedAtEnd:            publishedAtEnd,
		CategoriesTaxonomy:        categoriesTaxonomy,
		CategoriesConfident:       categoriesConfident,
		CategoriesId:              categoriesId,
		CategoriesLevel:           categoriesLevel,
		EntitiesTitleText:         entitiesTitleText,
		EntitiesTitleType:         entitiesTitleType,
		EntitiesTitleLinksDbpedia: entitiesTitleLinksDbpedia,
		EntitiesBodyText:          entitiesBodyText,
		EntitiesBodyType:          entitiesBodyType,
		EntitiesBodyLinksDbpedia:  entitiesBodyLinksDbpedia,
		SentimentTitlePolarity:    sentimentTitlePolarity,
		SentimentBodyPolarity:     sentimentBodyPolarity,
		AuthorId:                  authorId,
		AuthorName:                authorName,
		SourceId:                  sourceId,
		SourceName:                sourceName,
		SourceDomain:              sourceDomain,
		SourceLocationsCountry:    sourceLocationsCountry,
		SourceLocationsState:      sourceLocationsState,
		SourceLocationsCity:       sourceLocationsCity,
		SourceScopesCountry:       sourceScopesCountry,
		SourceScopesState:         sourceScopesState,
		SourceScopesCity:          sourceScopesCity,
		SourceScopesLevel:         sourceScopesLevel,
		Cluster:                   cluster,
		ClusterAlgorithm:          clusterAlgorithm,
		Return:                    _return,
		SortBy:                    sortBy,
		SortDirection:             sortDirection,
		Cursor:                    cursor,
		PerPage:                   perPage}

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
 **PublishedAtStart** | **string**| This parameter is used for finding stories whose published at is greater than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates). | [optional] 
 **PublishedAtEnd** | **string**| This parameter is used for finding stories whose published at is less than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates). | [optional] 
 **CategoriesTaxonomy** | **string**| This parameter is used for defining type of the taxonomy for the rest of categories queries. | [optional] 
 **CategoriesConfident** | **bool**| This parameter is used for finding stories whose categories are confident. | [optional] [default to true]
 **CategoriesId** | [**[]string**](string.md)| This parameter is used for finding stories whose categories id is the specified value. | [optional] 
 **CategoriesLevel** | [**[]int32**](int32.md)| This parameter is used for finding stories whose categories level is the specified value. | [optional] 
 **EntitiesTitleText** | [**[]string**](string.md)| This parameter is used for finding stories whose entities text in title is the specified value. | [optional] 
 **EntitiesTitleType** | [**[]string**](string.md)| This parameter is used for finding stories whose entities type in title is the specified value. | [optional] 
 **EntitiesTitleLinksDbpedia** | [**[]string**](string.md)| This parameter is used for finding stories whose entities dbpedia url in title is the specified value. | [optional] 
 **EntitiesBodyText** | [**[]string**](string.md)| This parameter is used for finding stories whose entities text in body is the specified value. | [optional] 
 **EntitiesBodyType** | [**[]string**](string.md)| This parameter is used for finding stories whose entities type in body is the specified value. | [optional] 
 **EntitiesBodyLinksDbpedia** | [**[]string**](string.md)| This parameter is used for finding stories whose entities dbpedia url in body is the specified value. | [optional] 
 **SentimentTitlePolarity** | **string**| This parameter is used for finding stories whose title sentiment is the specified value. | [optional] 
 **SentimentBodyPolarity** | **string**| This parameter is used for finding stories whose body sentiment is the specified value. | [optional] 
 **AuthorId** | [**[]int32**](int32.md)| This parameter is used for finding stories whose author id is the specified value. | [optional] 
 **AuthorName** | **string**| This parameter is used for finding stories whose author full name contains the specified value. | [optional] 
 **SourceId** | [**[]int32**](int32.md)| This parameter is used for finding stories whose source id is the specified value. | [optional] 
 **SourceName** | [**[]string**](string.md)| This parameter is used for finding stories whose source name contains the specified value. | [optional] 
 **SourceDomain** | [**[]string**](string.md)| This parameter is used for finding stories whose source domain is the specified value. | [optional] 
 **SourceLocationsCountry** | [**[]string**](string.md)| This parameter is used for finding stories whose source country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. | [optional] 
 **SourceLocationsState** | [**[]string**](string.md)| This parameter is used for finding stories whose source state/province is the specified value. | [optional] 
 **SourceLocationsCity** | [**[]string**](string.md)| This parameter is used for finding stories whose source city is the specified value. | [optional] 
 **SourceScopesCountry** | [**[]string**](string.md)| This parameter is used for finding stories whose source scopes country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. | [optional] 
 **SourceScopesState** | [**[]string**](string.md)| This parameter is used for finding stories whose source scopes state/province is the specified value. | [optional] 
 **SourceScopesCity** | [**[]string**](string.md)| This parameter is used for finding stories whose source scopes city is the specified value. | [optional] 
 **SourceScopesLevel** | [**[]string**](string.md)| This parameter is used for finding stories whose source scopes level is the specified value. | [optional] 
 **Cluster** | **bool**| This parameter enables clustering for the returned stories. | [optional] [default to false]
 **ClusterAlgorithm** | **string**| This parameter is used for specifying the clustering algorithm. It supprts STC, Lingo and [k-means](https://en.wikipedia.org/wiki/K-means_clustering) algorithms. | [optional] [default to lingo]
 **Return** | [**[]string**](string.md)| This parameter is used for specifying return fields. | [optional] 
 **SortBy** | **string**| This parameter is used for changing the order column of the result. | [optional] [default to published_at]
 **SortDirection** | **string**| This parameter is used for changing the order direction of the result. | [optional] [default to desc]
 **Cursor** | **string**| This parameter is used for finding a specific page. | [optional] [default to *]
 **PerPage** | **int32**| This parameter is used for specifying number of the items in each page. | [optional] [default to 10]

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

	id := []int64{112}                                                         // []int64 | This parameter is used for finding stroies by story id.
	title := "title_example"                                                   // string | This parameter is used for finding stories whose title contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
	body := "body_example"                                                     // string | This parameter is used for finding stories whose title contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
	text := "text_example"                                                     // string | This parameter is used for finding stories whose title or body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
	language := []string{"en"}                                                 // []string | This parameter is used for finding stories whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes.
	categoriesTaxonomy := "categoriesTaxonomy_example"                         // string | This parameter is used for defining type of the taxonomy for the rest of categories queries.
	categoriesConfident := "categoriesConfident_example"                       // string | This parameter is used for finding stories whose categories are confident.
	categoriesId := []string{"categoriesId_example"}                           // []string | This parameter is used for finding stories whose categories id is the specified value.
	categoriesLevel := []int32{3}                                              // []int32 | This parameter is used for finding stories whose categories level is the specified value.
	entitiesTitleText := []string{"entitiesTitleText_example"}                 // []string | This parameter is used for finding stories whose entities text in title is the specified value.
	entitiesTitleType := []string{"entitiesTitleType_example"}                 // []string| This parameter is used for finding stories whose entities type in title is the specified value.
	entitiesTitleLinksDbpedia := []string{"entitiesTitleLinksDbpedia_example"} // []string | This parameter is used for finding stories whose entities dbpedia url in title is the specified value.
	entitiesBodyText := []string{"entitiesBodyText_example"}                   // []string | This parameter is used for finding stories whose entities text in body is the specified value.
	entitiesBodyType := []string{"entitiesBodyType_example"}                   // []string | This parameter is used for finding stories whose entities type in body is the specified value.
	entitiesBodyLinksDbpedia := []string{"entitiesBodyLinksDbpedia_example"}   // []string | This parameter is used for finding stories whose entities dbpedia url in body is the specified value.
	sentimentTitlePolarity := "sentimentTitlePolarity_example"                 // string | This parameter is used for finding stories whose title sentiment is the specified value.
	sentimentBodyPolarity := "sentimentBodyPolarity_example"                   // string | This parameter is used for finding stories whose body sentiment is the specified value.
	authorId := []int32{25}                                                    // []int32 | This parameter is used for finding stories whose author id is the specified value.
	authorName := "authorName_example"                                         // string | This parameter is used for finding stories whose author full name contains the specified value.
	sourceId := []int32{45}                                                    // []int32 | This parameter is used for finding stories whose source id is the specified value.
	sourceName := []string{"sourceName_example"}                               // []string | This parameter is used for finding stories whose source name contains the specified value.
	sourceDomain := []string{"sourceDomain_example"}                           // []string | This parameter is used for finding stories whose source domain is the specified value.
	sourceLocationsCountry := []string{"sourceLocationsCountry_example"}       // []string | This parameter is used for finding stories whose source country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes.
	sourceLocationsState := []string{"sourceLocationsState_example"}           // []string | This parameter is used for finding stories whose source state/province is the specified value.
	sourceLocationsCity := []string{"sourceLocationsCity_example"}             // []string | This parameter is used for finding stories whose source city is the specified value.
	sourceScopesCountry := []string{"sourceScopesCountry_example"}             // []string | This parameter is used for finding stories whose source scopes country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes.
	sourceScopesState := []string{"sourceScopesState_example"}                 // []string | This parameter is used for finding stories whose source scopes state/province is the specified value.
	sourceScopesCity := []string{"sourceScopesCity_example"}                   // []string | This parameter is used for finding stories whose source scopes city is the specified value.
	sourceScopesLevel := []string{"sourceScopesLevel_example"}                 // []string | This parameter is used for finding stories whose source scopes level is the specified value.
	publishedAtStart := "NOW-7DAYS/DAY"                                        // string | This parameter is used for finding stories whose published at is greater than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates).
	publishedAtEnd := "NOW/DAY"                                                // string | This parameter is used for finding stories whose published at is less than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates).
	period := "+1DAY"                                                          // string | The size of each date range expressed as an interval to be added to the lower bound. It supports Date Math Syntax. Valid options are `+` following an integer number greater than 0 and one of the Date Math keywords. e.g. `+1DAY`, `+2MINUTES` and `+1MONTH`. Here are [Supported keywords](https://newsapi.aylien.com/docs/working-with-dates#date-math).

	timeSeriesParams := &newsapi.TimeSeriesParams{
		Id:                        id,
		Title:                     title,
		Body:                      body,
		Text:                      text,
		Language:                  language,
		CategoriesTaxonomy:        categoriesTaxonomy,
		CategoriesConfident:       categoriesConfident,
		CategoriesId:              categoriesId,
		CategoriesLevel:           categoriesLevel,
		EntitiesTitleText:         entitiesTitleText,
		EntitiesTitleType:         entitiesTitleType,
		EntitiesTitleLinksDbpedia: entitiesTitleLinksDbpedia,
		EntitiesBodyText:          entitiesBodyText,
		EntitiesBodyType:          entitiesBodyType,
		EntitiesBodyLinksDbpedia:  entitiesBodyLinksDbpedia,
		SentimentTitlePolarity:    sentimentTitlePolarity,
		SentimentBodyPolarity:     sentimentBodyPolarity,
		AuthorId:                  authorId,
		AuthorName:                authorName,
		SourceId:                  sourceId,
		SourceName:                sourceName,
		SourceDomain:              sourceDomain,
		SourceLocationsCountry:    sourceLocationsCountry,
		SourceLocationsState:      sourceLocationsState,
		SourceLocationsCity:       sourceLocationsCity,
		SourceScopesCountry:       sourceScopesCountry,
		SourceScopesState:         sourceScopesState,
		SourceScopesCity:          sourceScopesCity,
		SourceScopesLevel:         sourceScopesLevel,
		PublishedAtStart:          publishedAtStart,
		PublishedAtEnd:            publishedAtEnd,
		Period:                    period}

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
 **CategoriesTaxonomy** | **string**| This parameter is used for defining type of the taxonomy for the rest of categories queries. | [optional] 
 **CategoriesConfident** | **bool**| This parameter is used for finding stories whose categories are confident. | [optional] [default to true]
 **CategoriesId** | [**[]string**](string.md)| This parameter is used for finding stories whose categories id is the specified value. | [optional] 
 **CategoriesLevel** | [**[]int32**](int32.md)| This parameter is used for finding stories whose categories level is the specified value. | [optional] 
 **EntitiesTitleText** | [**[]string**](string.md)| This parameter is used for finding stories whose entities text in title is the specified value. | [optional] 
 **EntitiesTitleType** | [**[]string**](string.md)| This parameter is used for finding stories whose entities type in title is the specified value. | [optional] 
 **EntitiesTitleLinksDbpedia** | [**[]string**](string.md)| This parameter is used for finding stories whose entities dbpedia url in title is the specified value. | [optional] 
 **EntitiesBodyText** | [**[]string**](string.md)| This parameter is used for finding stories whose entities text in body is the specified value. | [optional] 
 **EntitiesBodyType** | [**[]string**](string.md)| This parameter is used for finding stories whose entities type in body is the specified value. | [optional] 
 **EntitiesBodyLinksDbpedia** | [**[]string**](string.md)| This parameter is used for finding stories whose entities dbpedia url in body is the specified value. | [optional] 
 **SentimentTitlePolarity** | **string**| This parameter is used for finding stories whose title sentiment is the specified value. | [optional] 
 **SentimentBodyPolarity** | **string**| This parameter is used for finding stories whose body sentiment is the specified value. | [optional] 
 **AuthorId** | [**[]int32**](int32.md)| This parameter is used for finding stories whose author id is the specified value. | [optional] 
 **AuthorName** | **string**| This parameter is used for finding stories whose author full name contains the specified value. | [optional] 
 **SourceId** | [**[]int32**](int32.md)| This parameter is used for finding stories whose source id is the specified value. | [optional] 
 **SourceName** | [**[]string**](string.md)| This parameter is used for finding stories whose source name contains the specified value. | [optional] 
 **SourceDomain** | [**[]string**](string.md)| This parameter is used for finding stories whose source domain is the specified value. | [optional] 
 **SourceLocationsCountry** | [**[]string**](string.md)| This parameter is used for finding stories whose source country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. | [optional] 
 **SourceLocationsState** | [**[]string**](string.md)| This parameter is used for finding stories whose source state/province is the specified value. | [optional] 
 **SourceLocationsCity** | [**[]string**](string.md)| This parameter is used for finding stories whose source city is the specified value. | [optional] 
 **SourceScopesCountry** | [**[]string**](string.md)| This parameter is used for finding stories whose source scopes country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. | [optional] 
 **SourceScopesState** | [**[]string**](string.md)| This parameter is used for finding stories whose source scopes state/province is the specified value. | [optional] 
 **SourceScopesCity** | [**[]string**](string.md)| This parameter is used for finding stories whose source scopes city is the specified value. | [optional] 
 **SourceScopesLevel** | [**[]string**](string.md)| This parameter is used for finding stories whose source scopes level is the specified value. | [optional] 
 **PublishedAtStart** | **string**| This parameter is used for finding stories whose published at is less than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates). | [optional] [default to NOW-7DAYS/DAY]
 **PublishedAtEnd** | **string**| This parameter is used for finding stories whose published at is greater than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates). | [optional] [default to NOW/DAY]
 **Period** | **string**| The size of each date range expressed as an interval to be added to the lower bound. It supports Date Math Syntax. Valid options are &#x60;+&#x60; following an integer number greater than 0 and one of the Date Math keywords. e.g. &#x60;+1DAY&#x60;, &#x60;+2MINUTES&#x60; and &#x60;+1MONTH&#x60;. Here are [Supported keywords](https://newsapi.aylien.com/docs/working-with-dates#date-math). | [optional] [default to +1DAY]

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

This endpoint is used for finding news trendings based on stories resource.

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

	id := []int64{112}                                                         // []int64 | This parameter is used for finding stroies by story id.
	title := "title_example"                                                   // string | This parameter is used for finding stories whose title contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
	body := "body_example"                                                     // string | This parameter is used for finding stories whose title contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
	text := "text_example"                                                     // string | This parameter is used for finding stories whose title or body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
	language := []string{"en"}                                                 // []string | This parameter is used for finding stories whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes.
	publishedAtStart := "publishedAtStart_example"                             // string | This parameter is used for finding stories whose published at is greater than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates).
	publishedAtEnd := "publishedAtEnd_example"                                 // string | This parameter is used for finding stories whose published at is less than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates).
	categoriesTaxonomy := "categoriesTaxonomy_example"                         // string | This parameter is used for defining type of the taxonomy for the rest of categories queries.
	categoriesConfident := "categoriesConfident_example"                       // string | This parameter is used for finding stories whose categories are confident.
	categoriesId := []string{"categoriesId_example"}                           // []string | This parameter is used for finding stories whose categories id is the specified value.
	categoriesLevel := []int32{3}                                              // []int32 | This parameter is used for finding stories whose categories level is the specified value.
	entitiesTitleText := []string{"entitiesTitleText_example"}                 // []string | This parameter is used for finding stories whose entities text in title is the specified value.
	entitiesTitleType := []string{"entitiesTitleType_example"}                 // []string| This parameter is used for finding stories whose entities type in title is the specified value.
	entitiesTitleLinksDbpedia := []string{"entitiesTitleLinksDbpedia_example"} // []string | This parameter is used for finding stories whose entities dbpedia url in title is the specified value.
	entitiesBodyText := []string{"entitiesBodyText_example"}                   // []string | This parameter is used for finding stories whose entities text in body is the specified value.
	entitiesBodyType := []string{"entitiesBodyType_example"}                   // []string | This parameter is used for finding stories whose entities type in body is the specified value.
	entitiesBodyLinksDbpedia := []string{"entitiesBodyLinksDbpedia_example"}   // []string | This parameter is used for finding stories whose entities dbpedia url in body is the specified value.
	sentimentTitlePolarity := "sentimentTitlePolarity_example"                 // string | This parameter is used for finding stories whose title sentiment is the specified value.
	sentimentBodyPolarity := "sentimentBodyPolarity_example"                   // string | This parameter is used for finding stories whose body sentiment is the specified value.
	authorId := []int32{25}                                                    // []int32 | This parameter is used for finding stories whose author id is the specified value.
	authorName := "authorName_example"                                         // string | This parameter is used for finding stories whose author full name contains the specified value.
	sourceId := []int32{45}                                                    // []int32 | This parameter is used for finding stories whose source id is the specified value.
	sourceName := []string{"sourceName_example"}                               // []string | This parameter is used for finding stories whose source name contains the specified value.
	sourceDomain := []string{"sourceDomain_example"}                           // []string | This parameter is used for finding stories whose source domain is the specified value.
	sourceLocationsCountry := []string{"sourceLocationsCountry_example"}       // []string | This parameter is used for finding stories whose source country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes.
	sourceLocationsState := []string{"sourceLocationsState_example"}           // []string | This parameter is used for finding stories whose source state/province is the specified value.
	sourceLocationsCity := []string{"sourceLocationsCity_example"}             // []string | This parameter is used for finding stories whose source city is the specified value.
	sourceScopesCountry := []string{"sourceScopesCountry_example"}             // []string | This parameter is used for finding stories whose source scopes country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes.
	sourceScopesState := []string{"sourceScopesState_example"}                 // []string | This parameter is used for finding stories whose source scopes state/province is the specified value.
	sourceScopesCity := []string{"sourceScopesCity_example"}                   // []string | This parameter is used for finding stories whose source scopes city is the specified value.
	sourceScopesLevel := []string{"sourceScopesLevel_example"}                 // []string | This parameter is used for finding stories whose source scopes level is the specified value.
	field := "field_example"                                                   // string | This parameter is used to specify the trend field.

	trendsParams := &newsapi.TrendsParams{
		Id:                        id,
		Title:                     title,
		Body:                      body,
		Text:                      text,
		Language:                  language,
		PublishedAtStart:          publishedAtStart,
		PublishedAtEnd:            publishedAtEnd,
		CategoriesTaxonomy:        categoriesTaxonomy,
		CategoriesConfident:       categoriesConfident,
		CategoriesId:              categoriesId,
		CategoriesLevel:           categoriesLevel,
		EntitiesTitleText:         entitiesTitleText,
		EntitiesTitleType:         entitiesTitleType,
		EntitiesTitleLinksDbpedia: entitiesTitleLinksDbpedia,
		EntitiesBodyText:          entitiesBodyText,
		EntitiesBodyType:          entitiesBodyType,
		EntitiesBodyLinksDbpedia:  entitiesBodyLinksDbpedia,
		SentimentTitlePolarity:    sentimentTitlePolarity,
		SentimentBodyPolarity:     sentimentBodyPolarity,
		AuthorId:                  authorId,
		AuthorName:                authorName,
		SourceId:                  sourceId,
		SourceName:                sourceName,
		SourceDomain:              sourceDomain,
		SourceLocationsCountry:    sourceLocationsCountry,
		SourceLocationsState:      sourceLocationsState,
		SourceLocationsCity:       sourceLocationsCity,
		SourceScopesCountry:       sourceScopesCountry,
		SourceScopesState:         sourceScopesState,
		SourceScopesCity:          sourceScopesCity,
		SourceScopesLevel:         sourceScopesLevel,
		Field:                     field}

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
 **PublishedAtStart** | **string**| This parameter is used for finding stories whose published at is greater than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates). | [optional] 
 **PublishedAtEnd** | **string**| This parameter is used for finding stories whose published at is less than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates). | [optional] 
 **CategoriesTaxonomy** | **string**| This parameter is used for defining type of the taxonomy for the rest of categories queries. | [optional] 
 **CategoriesConfident** | **bool**| This parameter is used for finding stories whose categories are confident. | [optional] [default to true]
 **CategoriesId** | [**[]string**](string.md)| This parameter is used for finding stories whose categories id is the specified value. | [optional] 
 **CategoriesLevel** | [**[]int32**](int32.md)| This parameter is used for finding stories whose categories level is the specified value. | [optional] 
 **EntitiesTitleText** | [**[]string**](string.md)| This parameter is used for finding stories whose entities text in title is the specified value. | [optional] 
 **EntitiesTitleType** | [**[]string**](string.md)| This parameter is used for finding stories whose entities type in title is the specified value. | [optional] 
 **EntitiesTitleLinksDbpedia** | [**[]string**](string.md)| This parameter is used for finding stories whose entities dbpedia url in title is the specified value. | [optional] 
 **EntitiesBodyText** | [**[]string**](string.md)| This parameter is used for finding stories whose entities text in body is the specified value. | [optional] 
 **EntitiesBodyType** | [**[]string**](string.md)| This parameter is used for finding stories whose entities type in body is the specified value. | [optional] 
 **EntitiesBodyLinksDbpedia** | [**[]string**](string.md)| This parameter is used for finding stories whose entities dbpedia url in body is the specified value. | [optional] 
 **SentimentTitlePolarity** | **string**| This parameter is used for finding stories whose title sentiment is the specified value. | [optional] 
 **SentimentBodyPolarity** | **string**| This parameter is used for finding stories whose body sentiment is the specified value. | [optional] 
 **AuthorId** | [**[]int32**](int32.md)| This parameter is used for finding stories whose author id is the specified value. | [optional] 
 **AuthorName** | **string**| This parameter is used for finding stories whose author full name contains the specified value. | [optional] 
 **SourceId** | [**[]int32**](int32.md)| This parameter is used for finding stories whose source id is the specified value. | [optional] 
 **SourceName** | [**[]string**](string.md)| This parameter is used for finding stories whose source name contains the specified value. | [optional] 
 **SourceDomain** | [**[]string**](string.md)| This parameter is used for finding stories whose source domain is the specified value. | [optional] 
 **SourceLocationsCountry** | [**[]string**](string.md)| This parameter is used for finding stories whose source country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. | [optional] 
 **SourceLocationsState** | [**[]string**](string.md)| This parameter is used for finding stories whose source state/province is the specified value. | [optional] 
 **SourceLocationsCity** | [**[]string**](string.md)| This parameter is used for finding stories whose source city is the specified value. | [optional] 
 **SourceScopesCountry** | [**[]string**](string.md)| This parameter is used for finding stories whose source scopes country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. | [optional] 
 **SourceScopesState** | [**[]string**](string.md)| This parameter is used for finding stories whose source scopes state/province is the specified value. | [optional] 
 **SourceScopesCity** | [**[]string**](string.md)| This parameter is used for finding stories whose source scopes city is the specified value. | [optional] 
 **SourceScopesLevel** | [**[]string**](string.md)| This parameter is used for finding stories whose source scopes level is the specified value. | [optional] 
 **Field** | **string**| This parameter is used to specify the trend field. | [optional] 

### Return type

[**Trends**](Trends.md)

### Authorization

[app_key](../README.md#app_key), [app_id](../README.md#app_id)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

