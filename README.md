# AYLIEN News API

AYLIEN News API is the most powerful way of sourcing, searching and syndicating analyzed and enriched news content. If you haven't already done so, you will need to [sign up](https://newsapi.aylien.com/signup).


## Installation
To install, simply use `go get`:

```bash
$ go get github.com/AYLIEN/aylien_newsapi_go
```

## Getting Started

Please follow the [installation](#installation) procedure and then run the following code:

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

	title := "trump"
	until := "NOW"
	since := "NOW-7DAYS"
	sortBy := "social_shares_count.facebook"
	language := []string{"en"}
	entities := []string{
		"http://dbpedia.org/resource/Donald_Trump",
		"http://dbpedia.org/resource/Hillary_Rodham_Clinton",
	}

	storiesParams := &newsapi.StoriesParams{
		Title:                    title,
		Language:                 language,
		PublishedAtStart:         since,
		PublishedAtEnd:           until,
		EntitiesBodyLinksDbpedia: entities,
		SortBy: sortBy}

	storiesResponse, res, err := api.ListStories(storiesParams)
	if err != nil {
		panic(err)
	}
	_ = res

	fmt.Println(storiesResponse)
}
```

## Documentation for API Endpoints

All URIs are relative to *https://api.newsapi.aylien.com/api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**ListAutocompletes**](docs/DefaultApi.md#listautocompletes) | **Get** /autocompletes | List autocompletes
*DefaultApi* | [**ListCoverages**](docs/DefaultApi.md#listcoverages) | **Post** /coverages | List coverages
*DefaultApi* | [**ListHistograms**](docs/DefaultApi.md#listhistograms) | **Get** /histograms | List histograms
*DefaultApi* | [**ListRelatedStories**](docs/DefaultApi.md#listrelatedstories) | **Post** /related_stories | List related stories
*DefaultApi* | [**ListStories**](docs/DefaultApi.md#liststories) | **Get** /stories | List Stories
*DefaultApi* | [**ListTimeSeries**](docs/DefaultApi.md#listtimeseries) | **Get** /time_series | List time series
*DefaultApi* | [**ListTrends**](docs/DefaultApi.md#listtrends) | **Get** /trends | List trends


## Documentation For Models

 - [Author](docs/Author.md)
 - [Autocomplete](docs/Autocomplete.md)
 - [Autocompletes](docs/Autocompletes.md)
 - [Category](docs/Category.md)
 - [CategoryLinks](docs/CategoryLinks.md)
 - [Coverages](docs/Coverages.md)
 - [Entities](docs/Entities.md)
 - [Entity](docs/Entity.md)
 - [EntityLinks](docs/EntityLinks.md)
 - [ErrorLinks](docs/ErrorLinks.md)
 - [ErrorModel](docs/ErrorModel.md)
 - [Errors](docs/Errors.md)
 - [HistogramInterval](docs/HistogramInterval.md)
 - [Histograms](docs/Histograms.md)
 - [Location](docs/Location.md)
 - [Media](docs/Media.md)
 - [RelatedStories](docs/RelatedStories.md)
 - [Scope](docs/Scope.md)
 - [Sentiment](docs/Sentiment.md)
 - [Sentiments](docs/Sentiments.md)
 - [ShareCount](docs/ShareCount.md)
 - [ShareCounts](docs/ShareCounts.md)
 - [Source](docs/Source.md)
 - [Stories](docs/Stories.md)
 - [Story](docs/Story.md)
 - [StoryCluster](docs/StoryCluster.md)
 - [StoryLinks](docs/StoryLinks.md)
 - [Summary](docs/Summary.md)
 - [TimeSeries](docs/TimeSeries.md)
 - [TimeSeriesList](docs/TimeSeriesList.md)
 - [Trend](docs/Trend.md)
 - [Trends](docs/Trends.md)


## Documentation For Authorization


## app_id

- **Type**: API key 
- **API key parameter name**: X-AYLIEN-NewsAPI-Application-ID
- **Location**: HTTP header

## app_key

- **Type**: API key 
- **API key parameter name**: X-AYLIEN-NewsAPI-Application-Key
- **Location**: HTTP header
