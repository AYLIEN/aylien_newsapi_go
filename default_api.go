/*
Copyright 2017 Aylien, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package newsapi

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
)

type DefaultApi struct {
	Configuration Configuration
}

func NewDefaultApi() *DefaultApi {
	configuration := NewConfiguration()
	return &DefaultApi{
		Configuration: *configuration,
	}
}

func NewDefaultApiWithBasePath(basePath string) *DefaultApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &DefaultApi{
		Configuration: *configuration,
	}
}

/**
 * List autocompletes
 * This endpoint is used for getting list of autocompletes by providing a specific term and type.
 *
 * @param params This is an AutocompletesParams struct which accepts following parameters:
 ** @param Type This parameter is used for defining the type of autocompletes.
 ** @param Term This parameter is used for finding autocomplete objects that contain the specified value.
 ** @param Language This parameter is used for autocompletes whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes.
 ** @param PerPage This parameter is used for specifying number of items in each page.
 * @return *Autocompletes
 */
func (a DefaultApi) ListAutocompletes(params *AutocompletesParams) (*Autocompletes, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/autocompletes"

	headerParams := make(map[string]string)
	queryParams := []QueryParams{}
	formParams := []FormParams{}
	// authentication '(app_key)' required
	// set key with prefix in header
	headerParams["X-AYLIEN-NewsAPI-Application-Key"] = a.Configuration.GetAPIKeyWithPrefix("X-AYLIEN-NewsAPI-Application-Key")
	// authentication '(app_id)' required
	// set key with prefix in header
	headerParams["X-AYLIEN-NewsAPI-Application-ID"] = a.Configuration.GetAPIKeyWithPrefix("X-AYLIEN-NewsAPI-Application-ID")
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	if !isZero(reflect.ValueOf(params.Type)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "type",
			Item2: a.Configuration.APIClient.ParameterToString(params.Type)})
	}
	if !isZero(reflect.ValueOf(params.Term)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "term",
			Item2: a.Configuration.APIClient.ParameterToString(params.Term)})
	}
	if !isZero(reflect.ValueOf(params.Language)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "language",
			Item2: a.Configuration.APIClient.ParameterToString(params.Language)})
	}
	if !isZero(reflect.ValueOf(params.PerPage)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "per_page",
			Item2: a.Configuration.APIClient.ParameterToString(params.PerPage)})
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"text/xml",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(Autocompletes)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, headerParams, queryParams, formParams)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse), err
	}

	defer httpResponse.Body.Close()
	resBody, err := ioutil.ReadAll(httpResponse.Body)

	err = json.Unmarshal(resBody, &successPayload)
	return successPayload, NewAPIResponse(httpResponse), err
}

/**
 * List coverages
 * This endpoint is used for finding story coverages based on the parameters provided. The maximum number of related stories returned is 100.
 *
 * @param params This is an CoveragesParams struct which accepts following parameters:
 ** @param Id This parameter is used for finding stories by story id.
 ** @param NotId This parameter is used for excluding stories by story id.
 ** @param Title This parameter is used for finding stories whose title contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
 ** @param Body This parameter is used for finding stories whose body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
 ** @param Text This parameter is used for finding stories whose title or body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
 ** @param Language This parameter is used for finding stories whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes.
 ** @param NotLanguage This parameter is used for excluding stories whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes.
 ** @param PublishedAtStart This parameter is used for finding stories whose published at time is greater than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates).
 ** @param PublishedAtEnd This parameter is used for finding stories whose published at time is less than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates).
 ** @param CategoriesTaxonomy This parameter is used for defining the type of the taxonomy for the rest of the categories queries. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories).
 ** @param CategoriesConfident This parameter is used for finding stories whose categories are confident. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories).
 ** @param CategoriesId This parameter is used for finding stories by categories id. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories).
 ** @param NotCategoriesId This parameter is used for excluding stories by categories id. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories).
 ** @param CategoriesLevel This parameter is used for finding stories by categories level. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories).
 ** @param NotCategoriesLevel This parameter is used for excluding stories by categories level. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories).
 ** @param EntitiesTitleText This parameter is used to find stories based on the specified entities &#x60;text&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param NotEntitiesTitleText This parameter is used to exclude stories based on the specified entities &#x60;text&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param EntitiesTitleType This parameter is used to find stories based on the specified entities &#x60;type&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param NotEntitiesTitleType This parameter is used to exclude stories based on the specified entities &#x60;type&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param EntitiesTitleLinksDbpedia This parameter is used to find stories based on the specified entities dbpedia URL in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param NotEntitiesTitleLinksDbpedia This parameter is used to exclude stories based on the specified entities dbpedia URL in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param EntitiesBodyText This parameter is used to find stories based on the specified entities &#x60;text&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param NotEntitiesBodyText This parameter is used to exclude stories based on the specified entities &#x60;text&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param EntitiesBodyType This parameter is used to find stories based on the specified entities &#x60;type&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param NotEntitiesBodyType This parameter is used to exclude stories based on the specified entities &#x60;type&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param EntitiesBodyLinksDbpedia This parameter is used to find stories based on the specified entities dbpedia URL in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param NotEntitiesBodyLinksDbpedia This parameter is used to exclude stories based on the specified entities dbpedia URL in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param SentimentTitlePolarity This parameter is used for finding stories whose title sentiment is the specified value.
 ** @param NotSentimentTitlePolarity This parameter is used for excluding stories whose title sentiment is the specified value.
 ** @param SentimentBodyPolarity This parameter is used for finding stories whose body sentiment is the specified value.
 ** @param NotSentimentBodyPolarity This parameter is used for excluding stories whose body sentiment is the specified value.
 ** @param MediaImagesCountMin This parameter is used for finding stories whose number of images is greater than or equal to the specified value.
 ** @param MediaImagesCountMax This parameter is used for finding stories whose number of images is less than or equal to the specified value.
 ** @param MediaImagesWidthMin This parameter is used for finding stories whose width of images are greater than or equal to the specified value.
 ** @param MediaImagesWidthMax This parameter is used for finding stories whose width of images are less than or equal to the specified value.
 ** @param MediaImagesHeightMin This parameter is used for finding stories whose height of images are greater than or equal to the specified value.
 ** @param MediaImagesHeightMax This parameter is used for finding stories whose height of images are less than or equal to the specified value.
 ** @param MediaImagesContentLengthMin This parameter is used for finding stories whose images content length are greater than or equal to the specified value.
 ** @param MediaImagesContentLengthMax This parameter is used for finding stories whose images content length are less than or equal to the specified value.
 ** @param MediaImagesFormat This parameter is used for finding stories whose images format are the specified value.
 ** @param NotMediaImagesFormat This parameter is used for excluding stories whose images format are the specified value.
 ** @param MediaVideosCountMin This parameter is used for finding stories whose number of videos is greater than or equal to the specified value.
 ** @param MediaVideosCountMax This parameter is used for finding stories whose number of videos is less than or equal to the specified value.
 ** @param AuthorId This parameter is used for finding stories whose author id is the specified value.
 ** @param NotAuthorId This parameter is used for excluding stories whose author id is the specified value.
 ** @param AuthorName This parameter is used for finding stories whose author full name contains the specified value.
 ** @param NotAuthorName This parameter is used for excluding stories whose author full name contains the specified value.
 ** @param SourceId This parameter is used for finding stories whose source id is the specified value.
 ** @param NotSourceId This parameter is used for excluding stories whose source id is the specified value.
 ** @param SourceName This parameter is used for finding stories whose source name contains the specified value.
 ** @param NotSourceName This parameter is used for excluding stories whose source name contains the specified value.
 ** @param SourceDomain This parameter is used for finding stories whose source domain is the specified value.
 ** @param NotSourceDomain This parameter is used for excluding stories whose source domain is the specified value.
 ** @param SourceLocationsCountry This parameter is used for finding stories whose source country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param NotSourceLocationsCountry This parameter is used for excluding stories whose source country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param SourceLocationsState This parameter is used for finding stories whose source state/province is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param NotSourceLocationsState This parameter is used for excluding stories whose source state/province is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param SourceLocationsCity This parameter is used for finding stories whose source city is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param NotSourceLocationsCity This parameter is used for excluding stories whose source city is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param SourceScopesCountry This parameter is used for finding stories whose source scopes  is the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param NotSourceScopesCountry This parameter is used for excluding stories whose source scopes  is the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param SourceScopesState This parameter is used for finding stories whose source scopes is the specified state/province value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param NotSourceScopesState This parameter is used for excluding stories whose source scopes is the specified state/province value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param SourceScopesCity This parameter is used for finding stories whose source scopes is the specified city value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param NotSourceScopesCity This parameter is used for excluding stories whose source scopes is the specified city value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param SourceScopesLevel This parameter is used for finding stories whose source scopes  is the specified level value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param NotSourceScopesLevel This parameter is used for excluding stories whose source scopes  is the specified level value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param SourceLinksInCountMin This parameter is used for finding stories from sources whose Links in count is greater than or equal to the specified value. You can read more about working with Links in count [here](https://newsapi.aylien.com/docs/working-with-links-in-count).
 ** @param SourceLinksInCountMax This parameter is used for finding stories from sources whose Links in count is less than or equal to the specified value. You can read more about working with Links in count [here](https://newsapi.aylien.com/docs/working-with-links-in-count).
 ** @param SourceRankingsAlexaRankMin This parameter is used for finding stories from sources whose Alexa rank is greater than or equal to the specified value. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks).
 ** @param SourceRankingsAlexaRankMax This parameter is used for finding stories from sources whose Alexa rank is less than or equal to the specified value. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks).
 ** @param SourceRankingsAlexaCountry This parameter is used for finding stories from sources whose Alexa rank is in the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks).
 ** @param SocialSharesCountFacebookMin This parameter is used for finding stories whose Facebook social shares count is greater than or equal to the specified value.
 ** @param SocialSharesCountFacebookMax This parameter is used for finding stories whose Facebook social shares count is less than or equal to the specified value.
 ** @param SocialSharesCountGooglePlusMin This parameter is used for finding stories whose Google+ social shares count is greater than or equal to the specified value.
 ** @param SocialSharesCountGooglePlusMax This parameter is used for finding stories whose Google+ social shares count is less than or equal to the specified value.
 ** @param SocialSharesCountLinkedinMin This parameter is used for finding stories whose LinkedIn social shares count is greater than or equal to the specified value.
 ** @param SocialSharesCountLinkedinMax This parameter is used for finding stories whose LinkedIn social shares count is less than or equal to the specified value.
 ** @param SocialSharesCountRedditMin This parameter is used for finding stories whose Reddit social shares count is greater than or equal to the specified value.
 ** @param SocialSharesCountRedditMax This parameter is used for finding stories whose Reddit social shares count is less than or equal to the specified value.
 ** @param Cluster This parameter enables clustering for the returned stories. You can read more about working with clustering [here](https://newsapi.aylien.com/docs/working-with-clustering).
 ** @param ClusterAlgorithm This parameter is used for specifying the clustering algorithm you wish to use. It supprts STC, Lingo and [k-means](https://en.wikipedia.org/wiki/K-means_clustering) algorithms. You can read more about working with clustering [here](https://newsapi.aylien.com/docs/working-with-clustering).
 ** @param Return This parameter is used for specifying return fields.
 ** @param StoryId A story id
 ** @param StoryUrl An article or webpage
 ** @param StoryTitle Title of the article
 ** @param StoryBody Body of the article
 ** @param StoryPublishedAt Publish date of the article. If you use a url or title and body of an article for getting coverages, this parameter is required. The format used is a restricted form of the canonical representation of dateTime in the [XML Schema specification (ISO 8601)](https://www.w3.org/TR/xmlschema-2/#dateTime).
 ** @param StoryLanguage This parameter is used for setting the language of the story. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes.
 ** @param PerPage This parameter is used for specifying number of items in each page.
 * @return *Coverages
 */
func (a DefaultApi) ListCoverages(params *CoveragesParams) (*Coverages, *APIResponse, error) {

	var httpMethod = "Post"
	// create path and map variables
	path := a.Configuration.BasePath + "/coverages"

	headerParams := make(map[string]string)
	queryParams := []QueryParams{}
	formParams := []FormParams{}
	// authentication '(app_key)' required
	// set key with prefix in header
	headerParams["X-AYLIEN-NewsAPI-Application-Key"] = a.Configuration.GetAPIKeyWithPrefix("X-AYLIEN-NewsAPI-Application-Key")
	// authentication '(app_id)' required
	// set key with prefix in header
	headerParams["X-AYLIEN-NewsAPI-Application-ID"] = a.Configuration.GetAPIKeyWithPrefix("X-AYLIEN-NewsAPI-Application-ID")
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"text/xml",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	for _, f := range params.Id {
		formParams = append(formParams, FormParams{
			Item1: "id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotId {
		formParams = append(formParams, FormParams{
			Item1: "!id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if !isZero(reflect.ValueOf(params.Title)) {
		formParams = append(formParams, FormParams{
			Item1: "title",
			Item2: a.Configuration.APIClient.ParameterToString(params.Title)})
	}
	if !isZero(reflect.ValueOf(params.Body)) {
		formParams = append(formParams, FormParams{
			Item1: "body",
			Item2: a.Configuration.APIClient.ParameterToString(params.Body)})
	}
	if !isZero(reflect.ValueOf(params.Text)) {
		formParams = append(formParams, FormParams{
			Item1: "text",
			Item2: a.Configuration.APIClient.ParameterToString(params.Text)})
	}
	for _, f := range params.Language {
		formParams = append(formParams, FormParams{
			Item1: "language[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotLanguage {
		formParams = append(formParams, FormParams{
			Item1: "!language[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if !isZero(reflect.ValueOf(params.PublishedAtStart)) {
		formParams = append(formParams, FormParams{
			Item1: "published_at.start",
			Item2: a.Configuration.APIClient.ParameterToString(params.PublishedAtStart)})
	}
	if !isZero(reflect.ValueOf(params.PublishedAtEnd)) {
		formParams = append(formParams, FormParams{
			Item1: "published_at.end",
			Item2: a.Configuration.APIClient.ParameterToString(params.PublishedAtEnd)})
	}
	if !isZero(reflect.ValueOf(params.CategoriesTaxonomy)) {
		formParams = append(formParams, FormParams{
			Item1: "categories.taxonomy",
			Item2: a.Configuration.APIClient.ParameterToString(params.CategoriesTaxonomy)})
	}
	if !isZero(reflect.ValueOf(params.CategoriesConfident)) {
		formParams = append(formParams, FormParams{
			Item1: "categories.confident",
			Item2: a.Configuration.APIClient.ParameterToString(params.CategoriesConfident)})
	}
	for _, f := range params.CategoriesId {
		formParams = append(formParams, FormParams{
			Item1: "categories.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotCategoriesId {
		formParams = append(formParams, FormParams{
			Item1: "!categories.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.CategoriesLevel {
		formParams = append(formParams, FormParams{
			Item1: "categories.level[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotCategoriesLevel {
		formParams = append(formParams, FormParams{
			Item1: "!categories.level[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesTitleText {
		formParams = append(formParams, FormParams{
			Item1: "entities.title.text[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotEntitiesTitleText {
		formParams = append(formParams, FormParams{
			Item1: "!entities.title.text[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesTitleType {
		formParams = append(formParams, FormParams{
			Item1: "entities.title.type[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotEntitiesTitleType {
		formParams = append(formParams, FormParams{
			Item1: "!entities.title.type[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesTitleLinksDbpedia {
		formParams = append(formParams, FormParams{
			Item1: "entities.title.links.dbpedia[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotEntitiesTitleLinksDbpedia {
		formParams = append(formParams, FormParams{
			Item1: "!entities.title.links.dbpedia[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesBodyText {
		formParams = append(formParams, FormParams{
			Item1: "entities.body.text[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotEntitiesBodyText {
		formParams = append(formParams, FormParams{
			Item1: "!entities.body.text[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesBodyType {
		formParams = append(formParams, FormParams{
			Item1: "entities.body.type[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotEntitiesBodyType {
		formParams = append(formParams, FormParams{
			Item1: "!entities.body.type[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesBodyLinksDbpedia {
		formParams = append(formParams, FormParams{
			Item1: "entities.body.links.dbpedia[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotEntitiesBodyLinksDbpedia {
		formParams = append(formParams, FormParams{
			Item1: "!entities.body.links.dbpedia[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if !isZero(reflect.ValueOf(params.SentimentTitlePolarity)) {
		formParams = append(formParams, FormParams{
			Item1: "sentiment.title.polarity",
			Item2: a.Configuration.APIClient.ParameterToString(params.SentimentTitlePolarity)})
	}
	if !isZero(reflect.ValueOf(params.NotSentimentTitlePolarity)) {
		formParams = append(formParams, FormParams{
			Item1: "!sentiment.title.polarity",
			Item2: a.Configuration.APIClient.ParameterToString(params.NotSentimentTitlePolarity)})
	}
	if !isZero(reflect.ValueOf(params.SentimentBodyPolarity)) {
		formParams = append(formParams, FormParams{
			Item1: "sentiment.body.polarity",
			Item2: a.Configuration.APIClient.ParameterToString(params.SentimentBodyPolarity)})
	}
	if !isZero(reflect.ValueOf(params.NotSentimentBodyPolarity)) {
		formParams = append(formParams, FormParams{
			Item1: "!sentiment.body.polarity",
			Item2: a.Configuration.APIClient.ParameterToString(params.NotSentimentBodyPolarity)})
	}
	if !isZero(reflect.ValueOf(params.MediaImagesCountMin)) {
		formParams = append(formParams, FormParams{
			Item1: "media.images.count.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaImagesCountMin)})
	}
	if !isZero(reflect.ValueOf(params.MediaImagesCountMax)) {
		formParams = append(formParams, FormParams{
			Item1: "media.images.count.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaImagesCountMax)})
	}
	if !isZero(reflect.ValueOf(params.MediaImagesWidthMin)) {
		formParams = append(formParams, FormParams{
			Item1: "media.images.width.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaImagesWidthMin)})
	}
	if !isZero(reflect.ValueOf(params.MediaImagesWidthMax)) {
		formParams = append(formParams, FormParams{
			Item1: "media.images.width.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaImagesWidthMax)})
	}
	if !isZero(reflect.ValueOf(params.MediaImagesHeightMin)) {
		formParams = append(formParams, FormParams{
			Item1: "media.images.height.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaImagesHeightMin)})
	}
	if !isZero(reflect.ValueOf(params.MediaImagesHeightMax)) {
		formParams = append(formParams, FormParams{
			Item1: "media.images.height.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaImagesHeightMax)})
	}
	if !isZero(reflect.ValueOf(params.MediaImagesContentLengthMin)) {
		formParams = append(formParams, FormParams{
			Item1: "media.images.content_length.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaImagesContentLengthMin)})
	}
	if !isZero(reflect.ValueOf(params.MediaImagesContentLengthMax)) {
		formParams = append(formParams, FormParams{
			Item1: "media.images.content_length.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaImagesContentLengthMax)})
	}
	for _, f := range params.MediaImagesFormat {
		formParams = append(formParams, FormParams{
			Item1: "media.images.format[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotMediaImagesFormat {
		formParams = append(formParams, FormParams{
			Item1: "!media.images.format[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if !isZero(reflect.ValueOf(params.MediaVideosCountMin)) {
		formParams = append(formParams, FormParams{
			Item1: "media.videos.count.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaVideosCountMin)})
	}
	if !isZero(reflect.ValueOf(params.MediaVideosCountMax)) {
		formParams = append(formParams, FormParams{
			Item1: "media.videos.count.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaVideosCountMax)})
	}
	for _, f := range params.AuthorId {
		formParams = append(formParams, FormParams{
			Item1: "author.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotAuthorId {
		formParams = append(formParams, FormParams{
			Item1: "!author.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if !isZero(reflect.ValueOf(params.AuthorName)) {
		formParams = append(formParams, FormParams{
			Item1: "author.name",
			Item2: a.Configuration.APIClient.ParameterToString(params.AuthorName)})
	}
	if !isZero(reflect.ValueOf(params.NotAuthorName)) {
		formParams = append(formParams, FormParams{
			Item1: "!author.name",
			Item2: a.Configuration.APIClient.ParameterToString(params.NotAuthorName)})
	}
	for _, f := range params.SourceId {
		formParams = append(formParams, FormParams{
			Item1: "source.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceId {
		formParams = append(formParams, FormParams{
			Item1: "!source.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceName {
		formParams = append(formParams, FormParams{
			Item1: "source.name[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceName {
		formParams = append(formParams, FormParams{
			Item1: "!source.name[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceDomain {
		formParams = append(formParams, FormParams{
			Item1: "source.domain[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceDomain {
		formParams = append(formParams, FormParams{
			Item1: "!source.domain[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceLocationsCountry {
		formParams = append(formParams, FormParams{
			Item1: "source.locations.country[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceLocationsCountry {
		formParams = append(formParams, FormParams{
			Item1: "!source.locations.country[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceLocationsState {
		formParams = append(formParams, FormParams{
			Item1: "source.locations.state[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceLocationsState {
		formParams = append(formParams, FormParams{
			Item1: "!source.locations.state[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceLocationsCity {
		formParams = append(formParams, FormParams{
			Item1: "source.locations.city[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceLocationsCity {
		formParams = append(formParams, FormParams{
			Item1: "!source.locations.city[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceScopesCountry {
		formParams = append(formParams, FormParams{
			Item1: "source.scopes.country[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceScopesCountry {
		formParams = append(formParams, FormParams{
			Item1: "!source.scopes.country[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceScopesState {
		formParams = append(formParams, FormParams{
			Item1: "source.scopes.state[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceScopesState {
		formParams = append(formParams, FormParams{
			Item1: "!source.scopes.state[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceScopesCity {
		formParams = append(formParams, FormParams{
			Item1: "source.scopes.city[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceScopesCity {
		formParams = append(formParams, FormParams{
			Item1: "!source.scopes.city[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceScopesLevel {
		formParams = append(formParams, FormParams{
			Item1: "source.scopes.level[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceScopesLevel {
		formParams = append(formParams, FormParams{
			Item1: "!source.scopes.level[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if !isZero(reflect.ValueOf(params.SourceLinksInCountMin)) {
		formParams = append(formParams, FormParams{
			Item1: "source.links_in_count.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.SourceLinksInCountMin)})
	}
	if !isZero(reflect.ValueOf(params.SourceLinksInCountMax)) {
		formParams = append(formParams, FormParams{
			Item1: "source.links_in_count.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.SourceLinksInCountMax)})
	}
	if !isZero(reflect.ValueOf(params.SourceRankingsAlexaRankMin)) {
		formParams = append(formParams, FormParams{
			Item1: "source.rankings.alexa.rank.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.SourceRankingsAlexaRankMin)})
	}
	if !isZero(reflect.ValueOf(params.SourceRankingsAlexaRankMax)) {
		formParams = append(formParams, FormParams{
			Item1: "source.rankings.alexa.rank.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.SourceRankingsAlexaRankMax)})
	}
	for _, f := range params.SourceRankingsAlexaCountry {
		formParams = append(formParams, FormParams{
			Item1: "source.rankings.alexa.country[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if !isZero(reflect.ValueOf(params.SocialSharesCountFacebookMin)) {
		formParams = append(formParams, FormParams{
			Item1: "social_shares_count.facebook.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.SocialSharesCountFacebookMin)})
	}
	if !isZero(reflect.ValueOf(params.SocialSharesCountFacebookMax)) {
		formParams = append(formParams, FormParams{
			Item1: "social_shares_count.facebook.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.SocialSharesCountFacebookMax)})
	}
	if !isZero(reflect.ValueOf(params.SocialSharesCountGooglePlusMin)) {
		formParams = append(formParams, FormParams{
			Item1: "social_shares_count.google_plus.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.SocialSharesCountGooglePlusMin)})
	}
	if !isZero(reflect.ValueOf(params.SocialSharesCountGooglePlusMax)) {
		formParams = append(formParams, FormParams{
			Item1: "social_shares_count.google_plus.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.SocialSharesCountGooglePlusMax)})
	}
	if !isZero(reflect.ValueOf(params.SocialSharesCountLinkedinMin)) {
		formParams = append(formParams, FormParams{
			Item1: "social_shares_count.linkedin.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.SocialSharesCountLinkedinMin)})
	}
	if !isZero(reflect.ValueOf(params.SocialSharesCountLinkedinMax)) {
		formParams = append(formParams, FormParams{
			Item1: "social_shares_count.linkedin.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.SocialSharesCountLinkedinMax)})
	}
	if !isZero(reflect.ValueOf(params.SocialSharesCountRedditMin)) {
		formParams = append(formParams, FormParams{
			Item1: "social_shares_count.reddit.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.SocialSharesCountRedditMin)})
	}
	if !isZero(reflect.ValueOf(params.SocialSharesCountRedditMax)) {
		formParams = append(formParams, FormParams{
			Item1: "social_shares_count.reddit.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.SocialSharesCountRedditMax)})
	}
	if !isZero(reflect.ValueOf(params.Cluster)) {
		formParams = append(formParams, FormParams{
			Item1: "cluster",
			Item2: a.Configuration.APIClient.ParameterToString(params.Cluster)})
	}
	if !isZero(reflect.ValueOf(params.ClusterAlgorithm)) {
		formParams = append(formParams, FormParams{
			Item1: "cluster.algorithm",
			Item2: a.Configuration.APIClient.ParameterToString(params.ClusterAlgorithm)})
	}
	for _, f := range params.Return {
		formParams = append(formParams, FormParams{
			Item1: "return[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if !isZero(reflect.ValueOf(params.StoryId)) {
		formParams = append(formParams, FormParams{
			Item1: "story_id",
			Item2: a.Configuration.APIClient.ParameterToString(params.StoryId)})
	}
	if !isZero(reflect.ValueOf(params.StoryUrl)) {
		formParams = append(formParams, FormParams{
			Item1: "story_url",
			Item2: a.Configuration.APIClient.ParameterToString(params.StoryUrl)})
	}
	if !isZero(reflect.ValueOf(params.StoryTitle)) {
		formParams = append(formParams, FormParams{
			Item1: "story_title",
			Item2: a.Configuration.APIClient.ParameterToString(params.StoryTitle)})
	}
	if !isZero(reflect.ValueOf(params.StoryBody)) {
		formParams = append(formParams, FormParams{
			Item1: "story_body",
			Item2: a.Configuration.APIClient.ParameterToString(params.StoryBody)})
	}
	if params.StoryPublishedAt.String() != "0001-01-01 00:00:00 +0000 UTC" {
		formParams = append(formParams, FormParams{
			Item1: "story_published_at",
			Item2: a.Configuration.APIClient.ParameterToString(params.StoryPublishedAt)})
	}
	if !isZero(reflect.ValueOf(params.StoryLanguage)) {
		formParams = append(formParams, FormParams{
			Item1: "story_language",
			Item2: a.Configuration.APIClient.ParameterToString(params.StoryLanguage)})
	}
	if !isZero(reflect.ValueOf(params.PerPage)) {
		formParams = append(formParams, FormParams{
			Item1: "per_page",
			Item2: a.Configuration.APIClient.ParameterToString(params.PerPage)})
	}
	var successPayload = new(Coverages)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, headerParams, queryParams, formParams)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse), err
	}

	defer httpResponse.Body.Close()
	resBody, err := ioutil.ReadAll(httpResponse.Body)

	err = json.Unmarshal(resBody, &successPayload)
	return successPayload, NewAPIResponse(httpResponse), err
}

/**
 * List histograms
 * This endpoint is used for getting histograms based on the &#x60;field&#x60; parameter passed to the API.
 *
 * @param params This is an HistogramsParams struct which accepts following parameters:
 ** @param Id This parameter is used for finding stories by story id.
 ** @param NotId This parameter is used for excluding stories by story id.
 ** @param Title This parameter is used for finding stories whose title contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
 ** @param Body This parameter is used for finding stories whose body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
 ** @param Text This parameter is used for finding stories whose title or body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
 ** @param Language This parameter is used for finding stories whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes.
 ** @param NotLanguage This parameter is used for excluding stories whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes.
 ** @param PublishedAtStart This parameter is used for finding stories whose published at time is greater than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates).
 ** @param PublishedAtEnd This parameter is used for finding stories whose published at time is less than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates).
 ** @param CategoriesTaxonomy This parameter is used for defining the type of the taxonomy for the rest of the categories queries. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories).
 ** @param CategoriesConfident This parameter is used for finding stories whose categories are confident. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories).
 ** @param CategoriesId This parameter is used for finding stories by categories id. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories).
 ** @param NotCategoriesId This parameter is used for excluding stories by categories id. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories).
 ** @param CategoriesLevel This parameter is used for finding stories by categories level. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories).
 ** @param NotCategoriesLevel This parameter is used for excluding stories by categories level. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories).
 ** @param EntitiesTitleText This parameter is used to find stories based on the specified entities &#x60;text&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param NotEntitiesTitleText This parameter is used to exclude stories based on the specified entities &#x60;text&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param EntitiesTitleType This parameter is used to find stories based on the specified entities &#x60;type&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param NotEntitiesTitleType This parameter is used to exclude stories based on the specified entities &#x60;type&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param EntitiesTitleLinksDbpedia This parameter is used to find stories based on the specified entities dbpedia URL in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param NotEntitiesTitleLinksDbpedia This parameter is used to exclude stories based on the specified entities dbpedia URL in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param EntitiesBodyText This parameter is used to find stories based on the specified entities &#x60;text&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param NotEntitiesBodyText This parameter is used to exclude stories based on the specified entities &#x60;text&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param EntitiesBodyType This parameter is used to find stories based on the specified entities &#x60;type&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param NotEntitiesBodyType This parameter is used to exclude stories based on the specified entities &#x60;type&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param EntitiesBodyLinksDbpedia This parameter is used to find stories based on the specified entities dbpedia URL in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param NotEntitiesBodyLinksDbpedia This parameter is used to exclude stories based on the specified entities dbpedia URL in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param SentimentTitlePolarity This parameter is used for finding stories whose title sentiment is the specified value.
 ** @param NotSentimentTitlePolarity This parameter is used for excluding stories whose title sentiment is the specified value.
 ** @param SentimentBodyPolarity This parameter is used for finding stories whose body sentiment is the specified value.
 ** @param NotSentimentBodyPolarity This parameter is used for excluding stories whose body sentiment is the specified value.
 ** @param MediaImagesCountMin This parameter is used for finding stories whose number of images is greater than or equal to the specified value.
 ** @param MediaImagesCountMax This parameter is used for finding stories whose number of images is less than or equal to the specified value.
 ** @param MediaImagesWidthMin This parameter is used for finding stories whose width of images are greater than or equal to the specified value.
 ** @param MediaImagesWidthMax This parameter is used for finding stories whose width of images are less than or equal to the specified value.
 ** @param MediaImagesHeightMin This parameter is used for finding stories whose height of images are greater than or equal to the specified value.
 ** @param MediaImagesHeightMax This parameter is used for finding stories whose height of images are less than or equal to the specified value.
 ** @param MediaImagesContentLengthMin This parameter is used for finding stories whose images content length are greater than or equal to the specified value.
 ** @param MediaImagesContentLengthMax This parameter is used for finding stories whose images content length are less than or equal to the specified value.
 ** @param MediaImagesFormat This parameter is used for finding stories whose images format are the specified value.
 ** @param NotMediaImagesFormat This parameter is used for excluding stories whose images format are the specified value.
 ** @param MediaVideosCountMin This parameter is used for finding stories whose number of videos is greater than or equal to the specified value.
 ** @param MediaVideosCountMax This parameter is used for finding stories whose number of videos is less than or equal to the specified value.
 ** @param AuthorId This parameter is used for finding stories whose author id is the specified value.
 ** @param NotAuthorId This parameter is used for excluding stories whose author id is the specified value.
 ** @param AuthorName This parameter is used for finding stories whose author full name contains the specified value.
 ** @param NotAuthorName This parameter is used for excluding stories whose author full name contains the specified value.
 ** @param SourceId This parameter is used for finding stories whose source id is the specified value.
 ** @param NotSourceId This parameter is used for excluding stories whose source id is the specified value.
 ** @param SourceName This parameter is used for finding stories whose source name contains the specified value.
 ** @param NotSourceName This parameter is used for excluding stories whose source name contains the specified value.
 ** @param SourceDomain This parameter is used for finding stories whose source domain is the specified value.
 ** @param NotSourceDomain This parameter is used for excluding stories whose source domain is the specified value.
 ** @param SourceLocationsCountry This parameter is used for finding stories whose source country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param NotSourceLocationsCountry This parameter is used for excluding stories whose source country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param SourceLocationsState This parameter is used for finding stories whose source state/province is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param NotSourceLocationsState This parameter is used for excluding stories whose source state/province is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param SourceLocationsCity This parameter is used for finding stories whose source city is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param NotSourceLocationsCity This parameter is used for excluding stories whose source city is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param SourceScopesCountry This parameter is used for finding stories whose source scopes is the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param NotSourceScopesCountry This parameter is used for excluding stories whose source scopes is the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param SourceScopesState This parameter is used for finding stories whose source scopes is the specified state/province value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param NotSourceScopesState This parameter is used for excluding stories whose source scopes is the specified state/province value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param SourceScopesCity This parameter is used for finding stories whose source scopes is the specified city value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param NotSourceScopesCity This parameter is used for excluding stories whose source scopes is the specified city value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param SourceScopesLevel This parameter is used for finding stories whose source scopes is the specified level value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param NotSourceScopesLevel This parameter is used for excluding stories whose source scopes is the specified level value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param SourceLinksInCountMin This parameter is used for finding stories from sources whose Links in count is greater than or equal to the specified value. You can read more about working with Links in count [here](https://newsapi.aylien.com/docs/working-with-links-in-count).
 ** @param SourceLinksInCountMax This parameter is used for finding stories from sources whose Links in count is less than or equal to the specified value. You can read more about working with Links in count [here](https://newsapi.aylien.com/docs/working-with-links-in-count).
 ** @param SourceRankingsAlexaRankMin This parameter is used for finding stories from sources whose Alexa rank is greater than or equal to the specified value. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks).
 ** @param SourceRankingsAlexaRankMax This parameter is used for finding stories from sources whose Alexa rank is less than or equal to the specified value. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks).
 ** @param SourceRankingsAlexaCountry This parameter is used for finding stories from sources whose Alexa rank is in the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks).
 ** @param SocialSharesCountFacebookMin This parameter is used for finding stories whose Facebook social shares count is greater than or equal to the specified value.
 ** @param SocialSharesCountFacebookMax This parameter is used for finding stories whose Facebook social shares count is less than or equal to the specified value.
 ** @param SocialSharesCountGooglePlusMin This parameter is used for finding stories whose Google+ social shares count is greater than or equal to the specified value.
 ** @param SocialSharesCountGooglePlusMax This parameter is used for finding stories whose Google+ social shares count is less than or equal to the specified value.
 ** @param SocialSharesCountLinkedinMin This parameter is used for finding stories whose LinkedIn social shares count is greater than or equal to the specified value.
 ** @param SocialSharesCountLinkedinMax This parameter is used for finding stories whose LinkedIn social shares count is less than or equal to the specified value.
 ** @param SocialSharesCountRedditMin This parameter is used for finding stories whose Reddit social shares count is greater than or equal to the specified value.
 ** @param SocialSharesCountRedditMax This parameter is used for finding stories whose Reddit social shares count is less than or equal to the specified value.
 ** @param IntervalStart This parameter is used for setting the start data point of histogram intervals.
 ** @param IntervalEnd This parameter is used for setting the end data point of histogram intervals.
 ** @param IntervalWidth This parameter is used for setting the width of histogram intervals.
 ** @param Field This parameter is used for specifying the y-axis variable for the histogram.
 * @return *Histograms
 */
func (a DefaultApi) ListHistograms(params *HistogramsParams) (*Histograms, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/histograms"

	headerParams := make(map[string]string)
	queryParams := []QueryParams{}
	formParams := []FormParams{}
	// authentication '(app_key)' required
	// set key with prefix in header
	headerParams["X-AYLIEN-NewsAPI-Application-Key"] = a.Configuration.GetAPIKeyWithPrefix("X-AYLIEN-NewsAPI-Application-Key")
	// authentication '(app_id)' required
	// set key with prefix in header
	headerParams["X-AYLIEN-NewsAPI-Application-ID"] = a.Configuration.GetAPIKeyWithPrefix("X-AYLIEN-NewsAPI-Application-ID")
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	for _, f := range params.Id {
		queryParams = append(queryParams, QueryParams{
			Item1: "id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotId {
		queryParams = append(queryParams, QueryParams{
			Item1: "!id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if !isZero(reflect.ValueOf(params.Title)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "title",
			Item2: a.Configuration.APIClient.ParameterToString(params.Title)})
	}
	if !isZero(reflect.ValueOf(params.Body)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "body",
			Item2: a.Configuration.APIClient.ParameterToString(params.Body)})
	}
	if !isZero(reflect.ValueOf(params.Text)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "text",
			Item2: a.Configuration.APIClient.ParameterToString(params.Text)})
	}
	for _, f := range params.Language {
		queryParams = append(queryParams, QueryParams{
			Item1: "language[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotLanguage {
		queryParams = append(queryParams, QueryParams{
			Item1: "!language[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if !isZero(reflect.ValueOf(params.PublishedAtStart)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "published_at.start",
			Item2: a.Configuration.APIClient.ParameterToString(params.PublishedAtStart)})
	}
	if !isZero(reflect.ValueOf(params.PublishedAtEnd)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "published_at.end",
			Item2: a.Configuration.APIClient.ParameterToString(params.PublishedAtEnd)})
	}
	if !isZero(reflect.ValueOf(params.CategoriesTaxonomy)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "categories.taxonomy",
			Item2: a.Configuration.APIClient.ParameterToString(params.CategoriesTaxonomy)})
	}
	if !isZero(reflect.ValueOf(params.CategoriesConfident)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "categories.confident",
			Item2: a.Configuration.APIClient.ParameterToString(params.CategoriesConfident)})
	}
	for _, f := range params.CategoriesId {
		queryParams = append(queryParams, QueryParams{
			Item1: "categories.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotCategoriesId {
		queryParams = append(queryParams, QueryParams{
			Item1: "!categories.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.CategoriesLevel {
		queryParams = append(queryParams, QueryParams{
			Item1: "categories.level[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotCategoriesLevel {
		queryParams = append(queryParams, QueryParams{
			Item1: "!categories.level[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesTitleText {
		queryParams = append(queryParams, QueryParams{
			Item1: "entities.title.text[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotEntitiesTitleText {
		queryParams = append(queryParams, QueryParams{
			Item1: "!entities.title.text[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesTitleType {
		queryParams = append(queryParams, QueryParams{
			Item1: "entities.title.type[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotEntitiesTitleType {
		queryParams = append(queryParams, QueryParams{
			Item1: "!entities.title.type[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesTitleLinksDbpedia {
		queryParams = append(queryParams, QueryParams{
			Item1: "entities.title.links.dbpedia[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotEntitiesTitleLinksDbpedia {
		queryParams = append(queryParams, QueryParams{
			Item1: "!entities.title.links.dbpedia[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesBodyText {
		queryParams = append(queryParams, QueryParams{
			Item1: "entities.body.text[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotEntitiesBodyText {
		queryParams = append(queryParams, QueryParams{
			Item1: "!entities.body.text[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesBodyType {
		queryParams = append(queryParams, QueryParams{
			Item1: "entities.body.type[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotEntitiesBodyType {
		queryParams = append(queryParams, QueryParams{
			Item1: "!entities.body.type[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesBodyLinksDbpedia {
		queryParams = append(queryParams, QueryParams{
			Item1: "entities.body.links.dbpedia[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotEntitiesBodyLinksDbpedia {
		queryParams = append(queryParams, QueryParams{
			Item1: "!entities.body.links.dbpedia[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if !isZero(reflect.ValueOf(params.SentimentTitlePolarity)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "sentiment.title.polarity",
			Item2: a.Configuration.APIClient.ParameterToString(params.SentimentTitlePolarity)})
	}
	if !isZero(reflect.ValueOf(params.NotSentimentTitlePolarity)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "!sentiment.title.polarity",
			Item2: a.Configuration.APIClient.ParameterToString(params.NotSentimentTitlePolarity)})
	}
	if !isZero(reflect.ValueOf(params.SentimentBodyPolarity)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "sentiment.body.polarity",
			Item2: a.Configuration.APIClient.ParameterToString(params.SentimentBodyPolarity)})
	}
	if !isZero(reflect.ValueOf(params.NotSentimentBodyPolarity)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "!sentiment.body.polarity",
			Item2: a.Configuration.APIClient.ParameterToString(params.NotSentimentBodyPolarity)})
	}
	if !isZero(reflect.ValueOf(params.MediaImagesCountMin)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "media.images.count.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaImagesCountMin)})
	}
	if !isZero(reflect.ValueOf(params.MediaImagesCountMax)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "media.images.count.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaImagesCountMax)})
	}
	if !isZero(reflect.ValueOf(params.MediaImagesWidthMin)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "media.images.width.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaImagesWidthMin)})
	}
	if !isZero(reflect.ValueOf(params.MediaImagesWidthMax)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "media.images.width.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaImagesWidthMax)})
	}
	if !isZero(reflect.ValueOf(params.MediaImagesHeightMin)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "media.images.height.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaImagesHeightMin)})
	}
	if !isZero(reflect.ValueOf(params.MediaImagesHeightMax)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "media.images.height.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaImagesHeightMax)})
	}
	if !isZero(reflect.ValueOf(params.MediaImagesContentLengthMin)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "media.images.content_length.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaImagesContentLengthMin)})
	}
	if !isZero(reflect.ValueOf(params.MediaImagesContentLengthMax)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "media.images.content_length.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaImagesContentLengthMax)})
	}
	for _, f := range params.MediaImagesFormat {
		queryParams = append(queryParams, QueryParams{
			Item1: "media.images.format[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotMediaImagesFormat {
		queryParams = append(queryParams, QueryParams{
			Item1: "!media.images.format[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if !isZero(reflect.ValueOf(params.MediaVideosCountMin)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "media.videos.count.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaVideosCountMin)})
	}
	if !isZero(reflect.ValueOf(params.MediaVideosCountMax)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "media.videos.count.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaVideosCountMax)})
	}
	for _, f := range params.AuthorId {
		queryParams = append(queryParams, QueryParams{
			Item1: "author.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotAuthorId {
		queryParams = append(queryParams, QueryParams{
			Item1: "!author.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if !isZero(reflect.ValueOf(params.AuthorName)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "author.name",
			Item2: a.Configuration.APIClient.ParameterToString(params.AuthorName)})
	}
	if !isZero(reflect.ValueOf(params.NotAuthorName)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "!author.name",
			Item2: a.Configuration.APIClient.ParameterToString(params.NotAuthorName)})
	}
	for _, f := range params.SourceId {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceId {
		queryParams = append(queryParams, QueryParams{
			Item1: "!source.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceName {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.name[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceName {
		queryParams = append(queryParams, QueryParams{
			Item1: "!source.name[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceDomain {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.domain[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceDomain {
		queryParams = append(queryParams, QueryParams{
			Item1: "!source.domain[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceLocationsCountry {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.locations.country[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceLocationsCountry {
		queryParams = append(queryParams, QueryParams{
			Item1: "!source.locations.country[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceLocationsState {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.locations.state[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceLocationsState {
		queryParams = append(queryParams, QueryParams{
			Item1: "!source.locations.state[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceLocationsCity {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.locations.city[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceLocationsCity {
		queryParams = append(queryParams, QueryParams{
			Item1: "!source.locations.city[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceScopesCountry {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.scopes.country[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceScopesCountry {
		queryParams = append(queryParams, QueryParams{
			Item1: "!source.scopes.country[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceScopesState {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.scopes.state[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceScopesState {
		queryParams = append(queryParams, QueryParams{
			Item1: "!source.scopes.state[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceScopesCity {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.scopes.city[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceScopesCity {
		queryParams = append(queryParams, QueryParams{
			Item1: "!source.scopes.city[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceScopesLevel {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.scopes.level[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceScopesLevel {
		queryParams = append(queryParams, QueryParams{
			Item1: "!source.scopes.level[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if !isZero(reflect.ValueOf(params.SourceLinksInCountMin)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.links_in_count.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.SourceLinksInCountMin)})
	}
	if !isZero(reflect.ValueOf(params.SourceLinksInCountMax)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.links_in_count.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.SourceLinksInCountMax)})
	}
	if !isZero(reflect.ValueOf(params.SourceRankingsAlexaRankMin)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.rankings.alexa.rank.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.SourceRankingsAlexaRankMin)})
	}
	if !isZero(reflect.ValueOf(params.SourceRankingsAlexaRankMax)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.rankings.alexa.rank.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.SourceRankingsAlexaRankMax)})
	}
	for _, f := range params.SourceRankingsAlexaCountry {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.rankings.alexa.country[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if !isZero(reflect.ValueOf(params.SocialSharesCountFacebookMin)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "social_shares_count.facebook.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.SocialSharesCountFacebookMin)})
	}
	if !isZero(reflect.ValueOf(params.SocialSharesCountFacebookMax)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "social_shares_count.facebook.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.SocialSharesCountFacebookMax)})
	}
	if !isZero(reflect.ValueOf(params.SocialSharesCountGooglePlusMin)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "social_shares_count.google_plus.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.SocialSharesCountGooglePlusMin)})
	}
	if !isZero(reflect.ValueOf(params.SocialSharesCountGooglePlusMax)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "social_shares_count.google_plus.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.SocialSharesCountGooglePlusMax)})
	}
	if !isZero(reflect.ValueOf(params.SocialSharesCountLinkedinMin)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "social_shares_count.linkedin.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.SocialSharesCountLinkedinMin)})
	}
	if !isZero(reflect.ValueOf(params.SocialSharesCountLinkedinMax)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "social_shares_count.linkedin.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.SocialSharesCountLinkedinMax)})
	}
	if !isZero(reflect.ValueOf(params.SocialSharesCountRedditMin)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "social_shares_count.reddit.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.SocialSharesCountRedditMin)})
	}
	if !isZero(reflect.ValueOf(params.SocialSharesCountRedditMax)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "social_shares_count.reddit.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.SocialSharesCountRedditMax)})
	}
	if !isZero(reflect.ValueOf(params.IntervalStart)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "interval.start",
			Item2: a.Configuration.APIClient.ParameterToString(params.IntervalStart)})
	}
	if !isZero(reflect.ValueOf(params.IntervalEnd)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "interval.end",
			Item2: a.Configuration.APIClient.ParameterToString(params.IntervalEnd)})
	}
	if !isZero(reflect.ValueOf(params.IntervalWidth)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "interval.width",
			Item2: a.Configuration.APIClient.ParameterToString(params.IntervalWidth)})
	}
	if !isZero(reflect.ValueOf(params.Field)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "field",
			Item2: a.Configuration.APIClient.ParameterToString(params.Field)})
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"text/xml",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(Histograms)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, headerParams, queryParams, formParams)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse), err
	}

	defer httpResponse.Body.Close()
	resBody, err := ioutil.ReadAll(httpResponse.Body)

	err = json.Unmarshal(resBody, &successPayload)
	return successPayload, NewAPIResponse(httpResponse), err
}

/**
 * List related stories
 * This endpoint is used for finding related stories based on the parameters provided. The maximum number of related stories returned is 100.
 *
 * @param params This is an RelatedStoriesParams struct which accepts following parameters:
 ** @param Id This parameter is used for finding stories by story id.
 ** @param NotId This parameter is used for excluding stories by story id.
 ** @param Title This parameter is used for finding stories whose title contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
 ** @param Body This parameter is used for finding stories whose body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
 ** @param Text This parameter is used for finding stories whose title or body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
 ** @param Language This parameter is used for finding stories whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes.
 ** @param NotLanguage This parameter is used for excluding stories whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes.
 ** @param PublishedAtStart This parameter is used for finding stories whose published at time is greater than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates).
 ** @param PublishedAtEnd This parameter is used for finding stories whose published at time is less than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates).
 ** @param CategoriesTaxonomy This parameter is used for defining the type of the taxonomy for the rest of the categories queries. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories).
 ** @param CategoriesConfident This parameter is used for finding stories whose categories are confident. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories).
 ** @param CategoriesId This parameter is used for finding stories by categories id. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories).
 ** @param NotCategoriesId This parameter is used for excluding stories by categories id. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories).
 ** @param CategoriesLevel This parameter is used for finding stories by categories level. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories).
 ** @param NotCategoriesLevel This parameter is used for excluding stories by categories level. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories).
 ** @param EntitiesTitleText This parameter is used to find stories based on the specified entities &#x60;text&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param NotEntitiesTitleText This parameter is used to exclude stories based on the specified entities &#x60;text&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param EntitiesTitleType This parameter is used to find stories based on the specified entities &#x60;type&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param NotEntitiesTitleType This parameter is used to exclude stories based on the specified entities &#x60;type&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param EntitiesTitleLinksDbpedia This parameter is used to find stories based on the specified entities dbpedia URL in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param NotEntitiesTitleLinksDbpedia This parameter is used to exclude stories based on the specified entities dbpedia URL in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param EntitiesBodyText This parameter is used to find stories based on the specified entities &#x60;text&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param NotEntitiesBodyText This parameter is used to exclude stories based on the specified entities &#x60;text&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param EntitiesBodyType This parameter is used to find stories based on the specified entities &#x60;type&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param NotEntitiesBodyType This parameter is used to exclude stories based on the specified entities &#x60;type&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param EntitiesBodyLinksDbpedia This parameter is used to find stories based on the specified entities dbpedia URL in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param NotEntitiesBodyLinksDbpedia This parameter is used to exclude stories based on the specified entities dbpedia URL in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param SentimentTitlePolarity This parameter is used for finding stories whose title sentiment is the specified value.
 ** @param NotSentimentTitlePolarity This parameter is used for excluding stories whose title sentiment is the specified value.
 ** @param SentimentBodyPolarity This parameter is used for finding stories whose body sentiment is the specified value.
 ** @param NotSentimentBodyPolarity This parameter is used for excluding stories whose body sentiment is the specified value.
 ** @param MediaImagesCountMin This parameter is used for finding stories whose number of images is greater than or equal to the specified value.
 ** @param MediaImagesCountMax This parameter is used for finding stories whose number of images is less than or equal to the specified value.
 ** @param MediaImagesWidthMin This parameter is used for finding stories whose width of images are greater than or equal to the specified value.
 ** @param MediaImagesWidthMax This parameter is used for finding stories whose width of images are less than or equal to the specified value.
 ** @param MediaImagesHeightMin This parameter is used for finding stories whose height of images are greater than or equal to the specified value.
 ** @param MediaImagesHeightMax This parameter is used for finding stories whose height of images are less than or equal to the specified value.
 ** @param MediaImagesContentLengthMin This parameter is used for finding stories whose images content length are greater than or equal to the specified value.
 ** @param MediaImagesContentLengthMax This parameter is used for finding stories whose images content length are less than or equal to the specified value.
 ** @param MediaImagesFormat This parameter is used for finding stories whose images format are the specified value.
 ** @param NotMediaImagesFormat This parameter is used for excluding stories whose images format are the specified value.
 ** @param MediaVideosCountMin This parameter is used for finding stories whose number of videos is greater than or equal to the specified value.
 ** @param MediaVideosCountMax This parameter is used for finding stories whose number of videos is less than or equal to the specified value.
 ** @param AuthorId This parameter is used for finding stories whose author id is the specified value.
 ** @param NotAuthorId This parameter is used for excluding stories whose author id is the specified value.
 ** @param AuthorName This parameter is used for finding stories whose author full name contains the specified value.
 ** @param NotAuthorName This parameter is used for excluding stories whose author full name contains the specified value.
 ** @param SourceId This parameter is used for finding stories whose source id is the specified value.
 ** @param NotSourceId This parameter is used for excluding stories whose source id is the specified value.
 ** @param SourceName This parameter is used for finding stories whose source name contains the specified value.
 ** @param NotSourceName This parameter is used for excluding stories whose source name contains the specified value.
 ** @param SourceDomain This parameter is used for finding stories whose source domain is the specified value.
 ** @param NotSourceDomain This parameter is used for excluding stories whose source domain is the specified value.
 ** @param SourceLocationsCountry This parameter is used for finding stories whose source country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param NotSourceLocationsCountry This parameter is used for excluding stories whose source country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param SourceLocationsState This parameter is used for finding stories whose source state/province is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param NotSourceLocationsState This parameter is used for excluding stories whose source state/province is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param SourceLocationsCity This parameter is used for finding stories whose source city is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param NotSourceLocationsCity This parameter is used for excluding stories whose source city is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param SourceScopesCountry This parameter is used for finding stories whose source scopes  is the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param NotSourceScopesCountry This parameter is used for excluding stories whose source scopes  is the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param SourceScopesState This parameter is used for finding stories whose source scopes is the specified state/province value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param NotSourceScopesState This parameter is used for excluding stories whose source scopes is the specified state/province value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param SourceScopesCity This parameter is used for finding stories whose source scopes is the specified city value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param NotSourceScopesCity This parameter is used for excluding stories whose source scopes is the specified city value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param SourceScopesLevel This parameter is used for finding stories whose source scopes  is the specified level value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param NotSourceScopesLevel This parameter is used for excluding stories whose source scopes  is the specified level value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param SourceLinksInCountMin This parameter is used for finding stories from sources whose Links in count is greater than or equal to the specified value. You can read more about working with Links in count [here](https://newsapi.aylien.com/docs/working-with-links-in-count).
 ** @param SourceLinksInCountMax This parameter is used for finding stories from sources whose Links in count is less than or equal to the specified value. You can read more about working with Links in count [here](https://newsapi.aylien.com/docs/working-with-links-in-count).
 ** @param SourceRankingsAlexaRankMin This parameter is used for finding stories from sources whose Alexa rank is greater than or equal to the specified value. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks).
 ** @param SourceRankingsAlexaRankMax This parameter is used for finding stories from sources whose Alexa rank is less than or equal to the specified value. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks).
 ** @param SourceRankingsAlexaCountry This parameter is used for finding stories from sources whose Alexa rank is in the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks).
 ** @param SocialSharesCountFacebookMin This parameter is used for finding stories whose Facebook social shares count is greater than or equal to the specified value.
 ** @param SocialSharesCountFacebookMax This parameter is used for finding stories whose Facebook social shares count is less than or equal to the specified value.
 ** @param SocialSharesCountGooglePlusMin This parameter is used for finding stories whose Google+ social shares count is greater than or equal to the specified value.
 ** @param SocialSharesCountGooglePlusMax This parameter is used for finding stories whose Google+ social shares count is less than or equal to the specified value.
 ** @param SocialSharesCountLinkedinMin This parameter is used for finding stories whose LinkedIn social shares count is greater than or equal to the specified value.
 ** @param SocialSharesCountLinkedinMax This parameter is used for finding stories whose LinkedIn social shares count is less than or equal to the specified value.
 ** @param SocialSharesCountRedditMin This parameter is used for finding stories whose Reddit social shares count is greater than or equal to the specified value.
 ** @param SocialSharesCountRedditMax This parameter is used for finding stories whose Reddit social shares count is less than or equal to the specified value.
 ** @param Cluster This parameter enables clustering for the returned stories. You can read more about working with clustering [here](https://newsapi.aylien.com/docs/working-with-clustering).
 ** @param ClusterAlgorithm This parameter is used for specifying the clustering algorithm you wish to use. It supprts STC, Lingo and [k-means](https://en.wikipedia.org/wiki/K-means_clustering) algorithms. You can read more about working with clustering [here](https://newsapi.aylien.com/docs/working-with-clustering).
 ** @param Return This parameter is used for specifying return fields.
 ** @param StoryId A story id
 ** @param StoryUrl An article or webpage
 ** @param StoryTitle Title of the article
 ** @param StoryBody Body of the article
 ** @param BoostBy This parameter is used for boosting the result by the specified value.
 ** @param StoryLanguage This parameter is used for setting the language of the story. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes.
 ** @param PerPage This parameter is used for specifying number of items in each page.
 * @return *RelatedStories
 */
func (a DefaultApi) ListRelatedStories(params *RelatedStoriesParams) (*RelatedStories, *APIResponse, error) {

	var httpMethod = "Post"
	// create path and map variables
	path := a.Configuration.BasePath + "/related_stories"

	headerParams := make(map[string]string)
	queryParams := []QueryParams{}
	formParams := []FormParams{}
	// authentication '(app_key)' required
	// set key with prefix in header
	headerParams["X-AYLIEN-NewsAPI-Application-Key"] = a.Configuration.GetAPIKeyWithPrefix("X-AYLIEN-NewsAPI-Application-Key")
	// authentication '(app_id)' required
	// set key with prefix in header
	headerParams["X-AYLIEN-NewsAPI-Application-ID"] = a.Configuration.GetAPIKeyWithPrefix("X-AYLIEN-NewsAPI-Application-ID")
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"text/xml",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	for _, f := range params.Id {
		formParams = append(formParams, FormParams{
			Item1: "id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotId {
		formParams = append(formParams, FormParams{
			Item1: "!id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if !isZero(reflect.ValueOf(params.Title)) {
		formParams = append(formParams, FormParams{
			Item1: "title",
			Item2: a.Configuration.APIClient.ParameterToString(params.Title)})
	}
	if !isZero(reflect.ValueOf(params.Body)) {
		formParams = append(formParams, FormParams{
			Item1: "body",
			Item2: a.Configuration.APIClient.ParameterToString(params.Body)})
	}
	if !isZero(reflect.ValueOf(params.Text)) {
		formParams = append(formParams, FormParams{
			Item1: "text",
			Item2: a.Configuration.APIClient.ParameterToString(params.Text)})
	}
	for _, f := range params.Language {
		formParams = append(formParams, FormParams{
			Item1: "language[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotLanguage {
		formParams = append(formParams, FormParams{
			Item1: "!language[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if !isZero(reflect.ValueOf(params.PublishedAtStart)) {
		formParams = append(formParams, FormParams{
			Item1: "published_at.start",
			Item2: a.Configuration.APIClient.ParameterToString(params.PublishedAtStart)})
	}
	if !isZero(reflect.ValueOf(params.PublishedAtEnd)) {
		formParams = append(formParams, FormParams{
			Item1: "published_at.end",
			Item2: a.Configuration.APIClient.ParameterToString(params.PublishedAtEnd)})
	}
	if !isZero(reflect.ValueOf(params.CategoriesTaxonomy)) {
		formParams = append(formParams, FormParams{
			Item1: "categories.taxonomy",
			Item2: a.Configuration.APIClient.ParameterToString(params.CategoriesTaxonomy)})
	}
	if !isZero(reflect.ValueOf(params.CategoriesConfident)) {
		formParams = append(formParams, FormParams{
			Item1: "categories.confident",
			Item2: a.Configuration.APIClient.ParameterToString(params.CategoriesConfident)})
	}
	for _, f := range params.CategoriesId {
		formParams = append(formParams, FormParams{
			Item1: "categories.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotCategoriesId {
		formParams = append(formParams, FormParams{
			Item1: "!categories.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.CategoriesLevel {
		formParams = append(formParams, FormParams{
			Item1: "categories.level[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotCategoriesLevel {
		formParams = append(formParams, FormParams{
			Item1: "!categories.level[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesTitleText {
		formParams = append(formParams, FormParams{
			Item1: "entities.title.text[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotEntitiesTitleText {
		formParams = append(formParams, FormParams{
			Item1: "!entities.title.text[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesTitleType {
		formParams = append(formParams, FormParams{
			Item1: "entities.title.type[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotEntitiesTitleType {
		formParams = append(formParams, FormParams{
			Item1: "!entities.title.type[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesTitleLinksDbpedia {
		formParams = append(formParams, FormParams{
			Item1: "entities.title.links.dbpedia[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotEntitiesTitleLinksDbpedia {
		formParams = append(formParams, FormParams{
			Item1: "!entities.title.links.dbpedia[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesBodyText {
		formParams = append(formParams, FormParams{
			Item1: "entities.body.text[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotEntitiesBodyText {
		formParams = append(formParams, FormParams{
			Item1: "!entities.body.text[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesBodyType {
		formParams = append(formParams, FormParams{
			Item1: "entities.body.type[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotEntitiesBodyType {
		formParams = append(formParams, FormParams{
			Item1: "!entities.body.type[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesBodyLinksDbpedia {
		formParams = append(formParams, FormParams{
			Item1: "entities.body.links.dbpedia[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotEntitiesBodyLinksDbpedia {
		formParams = append(formParams, FormParams{
			Item1: "!entities.body.links.dbpedia[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if !isZero(reflect.ValueOf(params.SentimentTitlePolarity)) {
		formParams = append(formParams, FormParams{
			Item1: "sentiment.title.polarity",
			Item2: a.Configuration.APIClient.ParameterToString(params.SentimentTitlePolarity)})
	}
	if !isZero(reflect.ValueOf(params.NotSentimentTitlePolarity)) {
		formParams = append(formParams, FormParams{
			Item1: "!sentiment.title.polarity",
			Item2: a.Configuration.APIClient.ParameterToString(params.NotSentimentTitlePolarity)})
	}
	if !isZero(reflect.ValueOf(params.SentimentBodyPolarity)) {
		formParams = append(formParams, FormParams{
			Item1: "sentiment.body.polarity",
			Item2: a.Configuration.APIClient.ParameterToString(params.SentimentBodyPolarity)})
	}
	if !isZero(reflect.ValueOf(params.NotSentimentBodyPolarity)) {
		formParams = append(formParams, FormParams{
			Item1: "!sentiment.body.polarity",
			Item2: a.Configuration.APIClient.ParameterToString(params.NotSentimentBodyPolarity)})
	}
	if !isZero(reflect.ValueOf(params.MediaImagesCountMin)) {
		formParams = append(formParams, FormParams{
			Item1: "media.images.count.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaImagesCountMin)})
	}
	if !isZero(reflect.ValueOf(params.MediaImagesCountMax)) {
		formParams = append(formParams, FormParams{
			Item1: "media.images.count.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaImagesCountMax)})
	}
	if !isZero(reflect.ValueOf(params.MediaImagesWidthMin)) {
		formParams = append(formParams, FormParams{
			Item1: "media.images.width.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaImagesWidthMin)})
	}
	if !isZero(reflect.ValueOf(params.MediaImagesWidthMax)) {
		formParams = append(formParams, FormParams{
			Item1: "media.images.width.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaImagesWidthMax)})
	}
	if !isZero(reflect.ValueOf(params.MediaImagesHeightMin)) {
		formParams = append(formParams, FormParams{
			Item1: "media.images.height.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaImagesHeightMin)})
	}
	if !isZero(reflect.ValueOf(params.MediaImagesHeightMax)) {
		formParams = append(formParams, FormParams{
			Item1: "media.images.height.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaImagesHeightMax)})
	}
	if !isZero(reflect.ValueOf(params.MediaImagesContentLengthMin)) {
		formParams = append(formParams, FormParams{
			Item1: "media.images.content_length.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaImagesContentLengthMin)})
	}
	if !isZero(reflect.ValueOf(params.MediaImagesContentLengthMax)) {
		formParams = append(formParams, FormParams{
			Item1: "media.images.content_length.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaImagesContentLengthMax)})
	}
	for _, f := range params.MediaImagesFormat {
		formParams = append(formParams, FormParams{
			Item1: "media.images.format[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotMediaImagesFormat {
		formParams = append(formParams, FormParams{
			Item1: "!media.images.format[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if !isZero(reflect.ValueOf(params.MediaVideosCountMin)) {
		formParams = append(formParams, FormParams{
			Item1: "media.videos.count.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaVideosCountMin)})
	}
	if !isZero(reflect.ValueOf(params.MediaVideosCountMax)) {
		formParams = append(formParams, FormParams{
			Item1: "media.videos.count.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaVideosCountMax)})
	}
	for _, f := range params.AuthorId {
		formParams = append(formParams, FormParams{
			Item1: "author.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotAuthorId {
		formParams = append(formParams, FormParams{
			Item1: "!author.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if !isZero(reflect.ValueOf(params.AuthorName)) {
		formParams = append(formParams, FormParams{
			Item1: "author.name",
			Item2: a.Configuration.APIClient.ParameterToString(params.AuthorName)})
	}
	if !isZero(reflect.ValueOf(params.NotAuthorName)) {
		formParams = append(formParams, FormParams{
			Item1: "!author.name",
			Item2: a.Configuration.APIClient.ParameterToString(params.NotAuthorName)})
	}
	for _, f := range params.SourceId {
		formParams = append(formParams, FormParams{
			Item1: "source.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceId {
		formParams = append(formParams, FormParams{
			Item1: "!source.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceName {
		formParams = append(formParams, FormParams{
			Item1: "source.name[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceName {
		formParams = append(formParams, FormParams{
			Item1: "!source.name[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceDomain {
		formParams = append(formParams, FormParams{
			Item1: "source.domain[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceDomain {
		formParams = append(formParams, FormParams{
			Item1: "!source.domain[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceLocationsCountry {
		formParams = append(formParams, FormParams{
			Item1: "source.locations.country[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceLocationsCountry {
		formParams = append(formParams, FormParams{
			Item1: "!source.locations.country[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceLocationsState {
		formParams = append(formParams, FormParams{
			Item1: "source.locations.state[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceLocationsState {
		formParams = append(formParams, FormParams{
			Item1: "!source.locations.state[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceLocationsCity {
		formParams = append(formParams, FormParams{
			Item1: "source.locations.city[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceLocationsCity {
		formParams = append(formParams, FormParams{
			Item1: "!source.locations.city[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceScopesCountry {
		formParams = append(formParams, FormParams{
			Item1: "source.scopes.country[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceScopesCountry {
		formParams = append(formParams, FormParams{
			Item1: "!source.scopes.country[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceScopesState {
		formParams = append(formParams, FormParams{
			Item1: "source.scopes.state[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceScopesState {
		formParams = append(formParams, FormParams{
			Item1: "!source.scopes.state[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceScopesCity {
		formParams = append(formParams, FormParams{
			Item1: "source.scopes.city[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceScopesCity {
		formParams = append(formParams, FormParams{
			Item1: "!source.scopes.city[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceScopesLevel {
		formParams = append(formParams, FormParams{
			Item1: "source.scopes.level[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceScopesLevel {
		formParams = append(formParams, FormParams{
			Item1: "!source.scopes.level[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if !isZero(reflect.ValueOf(params.SourceLinksInCountMin)) {
		formParams = append(formParams, FormParams{
			Item1: "source.links_in_count.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.SourceLinksInCountMin)})
	}
	if !isZero(reflect.ValueOf(params.SourceLinksInCountMax)) {
		formParams = append(formParams, FormParams{
			Item1: "source.links_in_count.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.SourceLinksInCountMax)})
	}
	if !isZero(reflect.ValueOf(params.SourceRankingsAlexaRankMin)) {
		formParams = append(formParams, FormParams{
			Item1: "source.rankings.alexa.rank.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.SourceRankingsAlexaRankMin)})
	}
	if !isZero(reflect.ValueOf(params.SourceRankingsAlexaRankMax)) {
		formParams = append(formParams, FormParams{
			Item1: "source.rankings.alexa.rank.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.SourceRankingsAlexaRankMax)})
	}
	for _, f := range params.SourceRankingsAlexaCountry {
		formParams = append(formParams, FormParams{
			Item1: "source.rankings.alexa.country[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if !isZero(reflect.ValueOf(params.SocialSharesCountFacebookMin)) {
		formParams = append(formParams, FormParams{
			Item1: "social_shares_count.facebook.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.SocialSharesCountFacebookMin)})
	}
	if !isZero(reflect.ValueOf(params.SocialSharesCountFacebookMax)) {
		formParams = append(formParams, FormParams{
			Item1: "social_shares_count.facebook.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.SocialSharesCountFacebookMax)})
	}
	if !isZero(reflect.ValueOf(params.SocialSharesCountGooglePlusMin)) {
		formParams = append(formParams, FormParams{
			Item1: "social_shares_count.google_plus.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.SocialSharesCountGooglePlusMin)})
	}
	if !isZero(reflect.ValueOf(params.SocialSharesCountGooglePlusMax)) {
		formParams = append(formParams, FormParams{
			Item1: "social_shares_count.google_plus.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.SocialSharesCountGooglePlusMax)})
	}
	if !isZero(reflect.ValueOf(params.SocialSharesCountLinkedinMin)) {
		formParams = append(formParams, FormParams{
			Item1: "social_shares_count.linkedin.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.SocialSharesCountLinkedinMin)})
	}
	if !isZero(reflect.ValueOf(params.SocialSharesCountLinkedinMax)) {
		formParams = append(formParams, FormParams{
			Item1: "social_shares_count.linkedin.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.SocialSharesCountLinkedinMax)})
	}
	if !isZero(reflect.ValueOf(params.SocialSharesCountRedditMin)) {
		formParams = append(formParams, FormParams{
			Item1: "social_shares_count.reddit.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.SocialSharesCountRedditMin)})
	}
	if !isZero(reflect.ValueOf(params.SocialSharesCountRedditMax)) {
		formParams = append(formParams, FormParams{
			Item1: "social_shares_count.reddit.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.SocialSharesCountRedditMax)})
	}
	if !isZero(reflect.ValueOf(params.Cluster)) {
		formParams = append(formParams, FormParams{
			Item1: "cluster",
			Item2: a.Configuration.APIClient.ParameterToString(params.Cluster)})
	}
	if !isZero(reflect.ValueOf(params.ClusterAlgorithm)) {
		formParams = append(formParams, FormParams{
			Item1: "cluster.algorithm",
			Item2: a.Configuration.APIClient.ParameterToString(params.ClusterAlgorithm)})
	}
	for _, f := range params.Return {
		formParams = append(formParams, FormParams{
			Item1: "return[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if !isZero(reflect.ValueOf(params.StoryId)) {
		formParams = append(formParams, FormParams{
			Item1: "story_id",
			Item2: a.Configuration.APIClient.ParameterToString(params.StoryId)})
	}
	if !isZero(reflect.ValueOf(params.StoryUrl)) {
		formParams = append(formParams, FormParams{
			Item1: "story_url",
			Item2: a.Configuration.APIClient.ParameterToString(params.StoryUrl)})
	}
	if !isZero(reflect.ValueOf(params.StoryTitle)) {
		formParams = append(formParams, FormParams{
			Item1: "story_title",
			Item2: a.Configuration.APIClient.ParameterToString(params.StoryTitle)})
	}
	if !isZero(reflect.ValueOf(params.StoryBody)) {
		formParams = append(formParams, FormParams{
			Item1: "story_body",
			Item2: a.Configuration.APIClient.ParameterToString(params.StoryBody)})
	}
	if !isZero(reflect.ValueOf(params.BoostBy)) {
		formParams = append(formParams, FormParams{
			Item1: "boost_by",
			Item2: a.Configuration.APIClient.ParameterToString(params.BoostBy)})
	}
	if !isZero(reflect.ValueOf(params.StoryLanguage)) {
		formParams = append(formParams, FormParams{
			Item1: "story_language",
			Item2: a.Configuration.APIClient.ParameterToString(params.StoryLanguage)})
	}
	if !isZero(reflect.ValueOf(params.PerPage)) {
		formParams = append(formParams, FormParams{
			Item1: "per_page",
			Item2: a.Configuration.APIClient.ParameterToString(params.PerPage)})
	}
	var successPayload = new(RelatedStories)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, headerParams, queryParams, formParams)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse), err
	}

	defer httpResponse.Body.Close()
	resBody, err := ioutil.ReadAll(httpResponse.Body)

	err = json.Unmarshal(resBody, &successPayload)
	return successPayload, NewAPIResponse(httpResponse), err
}

/**
 * List Stories
 * This endpoint is used for getting a list of stories.
 *
 * @param params This is an StoriesParams struct which accepts following parameters:
 ** @param Id This parameter is used for finding stories by story id.
 ** @param NotId This parameter is used for excluding stories by story id.
 ** @param Title This parameter is used for finding stories whose title contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
 ** @param Body This parameter is used for finding stories whose body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
 ** @param Text This parameter is used for finding stories whose title or body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
 ** @param Language This parameter is used for finding stories whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes.
 ** @param NotLanguage This parameter is used for excluding stories whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes.
 ** @param PublishedAtStart This parameter is used for finding stories whose published at time is greater than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates).
 ** @param PublishedAtEnd This parameter is used for finding stories whose published at time is less than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates).
 ** @param CategoriesTaxonomy This parameter is used for defining the type of the taxonomy for the rest of the categories queries. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories).
 ** @param CategoriesConfident This parameter is used for finding stories whose categories are confident. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories).
 ** @param CategoriesId This parameter is used for finding stories by categories id. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories).
 ** @param NotCategoriesId This parameter is used for excluding stories by categories id. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories).
 ** @param CategoriesLevel This parameter is used for finding stories by categories level. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories).
 ** @param NotCategoriesLevel This parameter is used for excluding stories by categories level. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories).
 ** @param EntitiesTitleText This parameter is used to find stories based on the specified entities &#x60;text&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param NotEntitiesTitleText This parameter is used to exclude stories based on the specified entities &#x60;text&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param EntitiesTitleType This parameter is used to find stories based on the specified entities &#x60;type&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param NotEntitiesTitleType This parameter is used to exclude stories based on the specified entities &#x60;type&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param EntitiesTitleLinksDbpedia This parameter is used to find stories based on the specified entities dbpedia URL in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param NotEntitiesTitleLinksDbpedia This parameter is used to exclude stories based on the specified entities dbpedia URL in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param EntitiesBodyText This parameter is used to find stories based on the specified entities &#x60;text&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param NotEntitiesBodyText This parameter is used to exclude stories based on the specified entities &#x60;text&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param EntitiesBodyType This parameter is used to find stories based on the specified entities &#x60;type&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param NotEntitiesBodyType This parameter is used to exclude stories based on the specified entities &#x60;type&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param EntitiesBodyLinksDbpedia This parameter is used to find stories based on the specified entities dbpedia URL in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param NotEntitiesBodyLinksDbpedia This parameter is used to exclude stories based on the specified entities dbpedia URL in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param SentimentTitlePolarity This parameter is used for finding stories whose title sentiment is the specified value.
 ** @param NotSentimentTitlePolarity This parameter is used for excluding stories whose title sentiment is the specified value.
 ** @param SentimentBodyPolarity This parameter is used for finding stories whose body sentiment is the specified value.
 ** @param NotSentimentBodyPolarity This parameter is used for excluding stories whose body sentiment is the specified value.
 ** @param MediaImagesCountMin This parameter is used for finding stories whose number of images is greater than or equal to the specified value.
 ** @param MediaImagesCountMax This parameter is used for finding stories whose number of images is less than or equal to the specified value.
 ** @param MediaImagesWidthMin This parameter is used for finding stories whose width of images are greater than or equal to the specified value.
 ** @param MediaImagesWidthMax This parameter is used for finding stories whose width of images are less than or equal to the specified value.
 ** @param MediaImagesHeightMin This parameter is used for finding stories whose height of images are greater than or equal to the specified value.
 ** @param MediaImagesHeightMax This parameter is used for finding stories whose height of images are less than or equal to the specified value.
 ** @param MediaImagesContentLengthMin This parameter is used for finding stories whose images content length are greater than or equal to the specified value.
 ** @param MediaImagesContentLengthMax This parameter is used for finding stories whose images content length are less than or equal to the specified value.
 ** @param MediaImagesFormat This parameter is used for finding stories whose images format are the specified value.
 ** @param NotMediaImagesFormat This parameter is used for excluding stories whose images format are the specified value.
 ** @param MediaVideosCountMin This parameter is used for finding stories whose number of videos is greater than or equal to the specified value.
 ** @param MediaVideosCountMax This parameter is used for finding stories whose number of videos is less than or equal to the specified value.
 ** @param AuthorId This parameter is used for finding stories whose author id is the specified value.
 ** @param NotAuthorId This parameter is used for excluding stories whose author id is the specified value.
 ** @param AuthorName This parameter is used for finding stories whose author full name contains the specified value.
 ** @param NotAuthorName This parameter is used for excluding stories whose author full name contains the specified value.
 ** @param SourceId This parameter is used for finding stories whose source id is the specified value.
 ** @param NotSourceId This parameter is used for excluding stories whose source id is the specified value.
 ** @param SourceName This parameter is used for finding stories whose source name contains the specified value.
 ** @param NotSourceName This parameter is used for excluding stories whose source name contains the specified value.
 ** @param SourceDomain This parameter is used for finding stories whose source domain is the specified value.
 ** @param NotSourceDomain This parameter is used for excluding stories whose source domain is the specified value.
 ** @param SourceLocationsCountry This parameter is used for finding stories whose source country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param NotSourceLocationsCountry This parameter is used for excluding stories whose source country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param SourceLocationsState This parameter is used for finding stories whose source state/province is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param NotSourceLocationsState This parameter is used for excluding stories whose source state/province is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param SourceLocationsCity This parameter is used for finding stories whose source city is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param NotSourceLocationsCity This parameter is used for excluding stories whose source city is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param SourceScopesCountry This parameter is used for finding stories whose source scopes is the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param NotSourceScopesCountry This parameter is used for excluding stories whose source scopes is the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param SourceScopesState This parameter is used for finding stories whose source scopes is the specified state/province value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param NotSourceScopesState This parameter is used for excluding stories whose source scopes is the specified state/province value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param SourceScopesCity This parameter is used for finding stories whose source scopes is the specified city value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param NotSourceScopesCity This parameter is used for excluding stories whose source scopes is the specified city value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param SourceScopesLevel This parameter is used for finding stories whose source scopes is the specified level value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param NotSourceScopesLevel This parameter is used for excluding stories whose source scopes is the specified level value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param SourceLinksInCountMin This parameter is used for finding stories from sources whose Links in count is greater than or equal to the specified value. You can read more about working with Links in count [here](https://newsapi.aylien.com/docs/working-with-links-in-count).
 ** @param SourceLinksInCountMax This parameter is used for finding stories from sources whose Links in count is less than or equal to the specified value. You can read more about working with Links in count [here](https://newsapi.aylien.com/docs/working-with-links-in-count).
 ** @param SourceRankingsAlexaRankMin This parameter is used for finding stories from sources whose Alexa rank is greater than or equal to the specified value. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks).
 ** @param SourceRankingsAlexaRankMax This parameter is used for finding stories from sources whose Alexa rank is less than or equal to the specified value. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks).
 ** @param SourceRankingsAlexaCountry This parameter is used for finding stories from sources whose Alexa rank is in the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks).
 ** @param SocialSharesCountFacebookMin This parameter is used for finding stories whose Facebook social shares count is greater than or equal to the specified value.
 ** @param SocialSharesCountFacebookMax This parameter is used for finding stories whose Facebook social shares count is less than or equal to the specified value.
 ** @param SocialSharesCountGooglePlusMin This parameter is used for finding stories whose Google+ social shares count is greater than or equal to the specified value.
 ** @param SocialSharesCountGooglePlusMax This parameter is used for finding stories whose Google+ social shares count is less than or equal to the specified value.
 ** @param SocialSharesCountLinkedinMin This parameter is used for finding stories whose LinkedIn social shares count is greater than or equal to the specified value.
 ** @param SocialSharesCountLinkedinMax This parameter is used for finding stories whose LinkedIn social shares count is less than or equal to the specified value.
 ** @param SocialSharesCountRedditMin This parameter is used for finding stories whose Reddit social shares count is greater than or equal to the specified value.
 ** @param SocialSharesCountRedditMax This parameter is used for finding stories whose Reddit social shares count is less than or equal to the specified value.
 ** @param Cluster This parameter enables clustering for the returned stories. You can read more about working with clustering [here](https://newsapi.aylien.com/docs/working-with-clustering).
 ** @param ClusterAlgorithm This parameter is used for specifying the clustering algorithm you wish to use. It supprts STC, Lingo and [k-means](https://en.wikipedia.org/wiki/K-means_clustering) algorithms. You can read more about working with clustering [here](https://newsapi.aylien.com/docs/working-with-clustering).
 ** @param Return This parameter is used for specifying return fields.
 ** @param SortBy This parameter is used for changing the order column of the results.
 ** @param SortDirection This parameter is used for changing the order direction of the result.
 ** @param Cursor This parameter is used for finding a specific page. You can read more about pagination of results [here](https://newsapi.aylien.com/docs/pagination-of-results).
 ** @param PerPage This parameter is used for specifying number of items in each page You can read more about pagination of results [here](https://newsapi.aylien.com/docs/pagination-of-results)
 * @return *Stories
 */
func (a DefaultApi) ListStories(params *StoriesParams) (*Stories, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/stories"

	headerParams := make(map[string]string)
	queryParams := []QueryParams{}
	formParams := []FormParams{}
	// authentication '(app_key)' required
	// set key with prefix in header
	headerParams["X-AYLIEN-NewsAPI-Application-Key"] = a.Configuration.GetAPIKeyWithPrefix("X-AYLIEN-NewsAPI-Application-Key")
	// authentication '(app_id)' required
	// set key with prefix in header
	headerParams["X-AYLIEN-NewsAPI-Application-ID"] = a.Configuration.GetAPIKeyWithPrefix("X-AYLIEN-NewsAPI-Application-ID")
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	for _, f := range params.Id {
		queryParams = append(queryParams, QueryParams{
			Item1: "id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotId {
		queryParams = append(queryParams, QueryParams{
			Item1: "!id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if !isZero(reflect.ValueOf(params.Title)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "title",
			Item2: a.Configuration.APIClient.ParameterToString(params.Title)})
	}
	if !isZero(reflect.ValueOf(params.Body)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "body",
			Item2: a.Configuration.APIClient.ParameterToString(params.Body)})
	}
	if !isZero(reflect.ValueOf(params.Text)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "text",
			Item2: a.Configuration.APIClient.ParameterToString(params.Text)})
	}
	for _, f := range params.Language {
		queryParams = append(queryParams, QueryParams{
			Item1: "language[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotLanguage {
		queryParams = append(queryParams, QueryParams{
			Item1: "!language[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if !isZero(reflect.ValueOf(params.PublishedAtStart)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "published_at.start",
			Item2: a.Configuration.APIClient.ParameterToString(params.PublishedAtStart)})
	}
	if !isZero(reflect.ValueOf(params.PublishedAtEnd)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "published_at.end",
			Item2: a.Configuration.APIClient.ParameterToString(params.PublishedAtEnd)})
	}
	if !isZero(reflect.ValueOf(params.CategoriesTaxonomy)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "categories.taxonomy",
			Item2: a.Configuration.APIClient.ParameterToString(params.CategoriesTaxonomy)})
	}
	if !isZero(reflect.ValueOf(params.CategoriesConfident)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "categories.confident",
			Item2: a.Configuration.APIClient.ParameterToString(params.CategoriesConfident)})
	}
	for _, f := range params.CategoriesId {
		queryParams = append(queryParams, QueryParams{
			Item1: "categories.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotCategoriesId {
		queryParams = append(queryParams, QueryParams{
			Item1: "!categories.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.CategoriesLevel {
		queryParams = append(queryParams, QueryParams{
			Item1: "categories.level[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotCategoriesLevel {
		queryParams = append(queryParams, QueryParams{
			Item1: "!categories.level[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesTitleText {
		queryParams = append(queryParams, QueryParams{
			Item1: "entities.title.text[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotEntitiesTitleText {
		queryParams = append(queryParams, QueryParams{
			Item1: "!entities.title.text[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesTitleType {
		queryParams = append(queryParams, QueryParams{
			Item1: "entities.title.type[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotEntitiesTitleType {
		queryParams = append(queryParams, QueryParams{
			Item1: "!entities.title.type[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesTitleLinksDbpedia {
		queryParams = append(queryParams, QueryParams{
			Item1: "entities.title.links.dbpedia[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotEntitiesTitleLinksDbpedia {
		queryParams = append(queryParams, QueryParams{
			Item1: "!entities.title.links.dbpedia[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesBodyText {
		queryParams = append(queryParams, QueryParams{
			Item1: "entities.body.text[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotEntitiesBodyText {
		queryParams = append(queryParams, QueryParams{
			Item1: "!entities.body.text[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesBodyType {
		queryParams = append(queryParams, QueryParams{
			Item1: "entities.body.type[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotEntitiesBodyType {
		queryParams = append(queryParams, QueryParams{
			Item1: "!entities.body.type[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesBodyLinksDbpedia {
		queryParams = append(queryParams, QueryParams{
			Item1: "entities.body.links.dbpedia[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotEntitiesBodyLinksDbpedia {
		queryParams = append(queryParams, QueryParams{
			Item1: "!entities.body.links.dbpedia[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if !isZero(reflect.ValueOf(params.SentimentTitlePolarity)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "sentiment.title.polarity",
			Item2: a.Configuration.APIClient.ParameterToString(params.SentimentTitlePolarity)})
	}
	if !isZero(reflect.ValueOf(params.NotSentimentTitlePolarity)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "!sentiment.title.polarity",
			Item2: a.Configuration.APIClient.ParameterToString(params.NotSentimentTitlePolarity)})
	}
	if !isZero(reflect.ValueOf(params.SentimentBodyPolarity)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "sentiment.body.polarity",
			Item2: a.Configuration.APIClient.ParameterToString(params.SentimentBodyPolarity)})
	}
	if !isZero(reflect.ValueOf(params.NotSentimentBodyPolarity)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "!sentiment.body.polarity",
			Item2: a.Configuration.APIClient.ParameterToString(params.NotSentimentBodyPolarity)})
	}
	if !isZero(reflect.ValueOf(params.MediaImagesCountMin)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "media.images.count.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaImagesCountMin)})
	}
	if !isZero(reflect.ValueOf(params.MediaImagesCountMax)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "media.images.count.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaImagesCountMax)})
	}
	if !isZero(reflect.ValueOf(params.MediaImagesWidthMin)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "media.images.width.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaImagesWidthMin)})
	}
	if !isZero(reflect.ValueOf(params.MediaImagesWidthMax)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "media.images.width.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaImagesWidthMax)})
	}
	if !isZero(reflect.ValueOf(params.MediaImagesHeightMin)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "media.images.height.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaImagesHeightMin)})
	}
	if !isZero(reflect.ValueOf(params.MediaImagesHeightMax)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "media.images.height.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaImagesHeightMax)})
	}
	if !isZero(reflect.ValueOf(params.MediaImagesContentLengthMin)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "media.images.content_length.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaImagesContentLengthMin)})
	}
	if !isZero(reflect.ValueOf(params.MediaImagesContentLengthMax)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "media.images.content_length.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaImagesContentLengthMax)})
	}
	for _, f := range params.MediaImagesFormat {
		queryParams = append(queryParams, QueryParams{
			Item1: "media.images.format[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotMediaImagesFormat {
		queryParams = append(queryParams, QueryParams{
			Item1: "!media.images.format[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if !isZero(reflect.ValueOf(params.MediaVideosCountMin)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "media.videos.count.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaVideosCountMin)})
	}
	if !isZero(reflect.ValueOf(params.MediaVideosCountMax)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "media.videos.count.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaVideosCountMax)})
	}
	for _, f := range params.AuthorId {
		queryParams = append(queryParams, QueryParams{
			Item1: "author.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotAuthorId {
		queryParams = append(queryParams, QueryParams{
			Item1: "!author.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if !isZero(reflect.ValueOf(params.AuthorName)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "author.name",
			Item2: a.Configuration.APIClient.ParameterToString(params.AuthorName)})
	}
	if !isZero(reflect.ValueOf(params.NotAuthorName)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "!author.name",
			Item2: a.Configuration.APIClient.ParameterToString(params.NotAuthorName)})
	}
	for _, f := range params.SourceId {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceId {
		queryParams = append(queryParams, QueryParams{
			Item1: "!source.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceName {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.name[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceName {
		queryParams = append(queryParams, QueryParams{
			Item1: "!source.name[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceDomain {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.domain[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceDomain {
		queryParams = append(queryParams, QueryParams{
			Item1: "!source.domain[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceLocationsCountry {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.locations.country[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceLocationsCountry {
		queryParams = append(queryParams, QueryParams{
			Item1: "!source.locations.country[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceLocationsState {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.locations.state[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceLocationsState {
		queryParams = append(queryParams, QueryParams{
			Item1: "!source.locations.state[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceLocationsCity {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.locations.city[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceLocationsCity {
		queryParams = append(queryParams, QueryParams{
			Item1: "!source.locations.city[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceScopesCountry {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.scopes.country[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceScopesCountry {
		queryParams = append(queryParams, QueryParams{
			Item1: "!source.scopes.country[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceScopesState {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.scopes.state[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceScopesState {
		queryParams = append(queryParams, QueryParams{
			Item1: "!source.scopes.state[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceScopesCity {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.scopes.city[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceScopesCity {
		queryParams = append(queryParams, QueryParams{
			Item1: "!source.scopes.city[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceScopesLevel {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.scopes.level[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceScopesLevel {
		queryParams = append(queryParams, QueryParams{
			Item1: "!source.scopes.level[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if !isZero(reflect.ValueOf(params.SourceLinksInCountMin)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.links_in_count.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.SourceLinksInCountMin)})
	}
	if !isZero(reflect.ValueOf(params.SourceLinksInCountMax)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.links_in_count.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.SourceLinksInCountMax)})
	}
	if !isZero(reflect.ValueOf(params.SourceRankingsAlexaRankMin)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.rankings.alexa.rank.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.SourceRankingsAlexaRankMin)})
	}
	if !isZero(reflect.ValueOf(params.SourceRankingsAlexaRankMax)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.rankings.alexa.rank.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.SourceRankingsAlexaRankMax)})
	}
	for _, f := range params.SourceRankingsAlexaCountry {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.rankings.alexa.country[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if !isZero(reflect.ValueOf(params.SocialSharesCountFacebookMin)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "social_shares_count.facebook.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.SocialSharesCountFacebookMin)})
	}
	if !isZero(reflect.ValueOf(params.SocialSharesCountFacebookMax)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "social_shares_count.facebook.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.SocialSharesCountFacebookMax)})
	}
	if !isZero(reflect.ValueOf(params.SocialSharesCountGooglePlusMin)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "social_shares_count.google_plus.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.SocialSharesCountGooglePlusMin)})
	}
	if !isZero(reflect.ValueOf(params.SocialSharesCountGooglePlusMax)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "social_shares_count.google_plus.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.SocialSharesCountGooglePlusMax)})
	}
	if !isZero(reflect.ValueOf(params.SocialSharesCountLinkedinMin)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "social_shares_count.linkedin.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.SocialSharesCountLinkedinMin)})
	}
	if !isZero(reflect.ValueOf(params.SocialSharesCountLinkedinMax)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "social_shares_count.linkedin.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.SocialSharesCountLinkedinMax)})
	}
	if !isZero(reflect.ValueOf(params.SocialSharesCountRedditMin)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "social_shares_count.reddit.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.SocialSharesCountRedditMin)})
	}
	if !isZero(reflect.ValueOf(params.SocialSharesCountRedditMax)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "social_shares_count.reddit.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.SocialSharesCountRedditMax)})
	}
	if !isZero(reflect.ValueOf(params.Cluster)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "cluster",
			Item2: a.Configuration.APIClient.ParameterToString(params.Cluster)})
	}
	if !isZero(reflect.ValueOf(params.ClusterAlgorithm)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "cluster.algorithm",
			Item2: a.Configuration.APIClient.ParameterToString(params.ClusterAlgorithm)})
	}
	for _, f := range params.Return {
		queryParams = append(queryParams, QueryParams{
			Item1: "return[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if !isZero(reflect.ValueOf(params.SortBy)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "sort_by",
			Item2: a.Configuration.APIClient.ParameterToString(params.SortBy)})
	}
	if !isZero(reflect.ValueOf(params.SortDirection)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "sort_direction",
			Item2: a.Configuration.APIClient.ParameterToString(params.SortDirection)})
	}
	if !isZero(reflect.ValueOf(params.Cursor)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "cursor",
			Item2: a.Configuration.APIClient.ParameterToString(params.Cursor)})
	}
	if !isZero(reflect.ValueOf(params.PerPage)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "per_page",
			Item2: a.Configuration.APIClient.ParameterToString(params.PerPage)})
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"text/xml",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(Stories)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, headerParams, queryParams, formParams)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse), err
	}

	defer httpResponse.Body.Close()
	resBody, err := ioutil.ReadAll(httpResponse.Body)

	err = json.Unmarshal(resBody, &successPayload)
	return successPayload, NewAPIResponse(httpResponse), err
}

/**
 * List time series
 * This endpoint is used for getting time series by stories.
 *
 * @param params This is an TimeSeriesParams struct which accepts following parameters:
 ** @param Id This parameter is used for finding stories by story id.
 ** @param NotId This parameter is used for excluding stories by story id.
 ** @param Title This parameter is used for finding stories whose title contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
 ** @param Body This parameter is used for finding stories whose body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
 ** @param Text This parameter is used for finding stories whose title or body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
 ** @param Language This parameter is used for finding stories whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes.
 ** @param NotLanguage This parameter is used for excluding stories whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes.
 ** @param CategoriesTaxonomy This parameter is used for defining the type of the taxonomy for the rest of the categories queries. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories).
 ** @param CategoriesConfident This parameter is used for finding stories whose categories are confident. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories).
 ** @param CategoriesId This parameter is used for finding stories by categories id. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories).
 ** @param NotCategoriesId This parameter is used for excluding stories by categories id. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories).
 ** @param CategoriesLevel This parameter is used for finding stories by categories level. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories).
 ** @param NotCategoriesLevel This parameter is used for excluding stories by categories level. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories).
 ** @param EntitiesTitleText This parameter is used to find stories based on the specified entities &#x60;text&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param NotEntitiesTitleText This parameter is used to exclude stories based on the specified entities &#x60;text&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param EntitiesTitleType This parameter is used to find stories based on the specified entities &#x60;type&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param NotEntitiesTitleType This parameter is used to exclude stories based on the specified entities &#x60;type&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param EntitiesTitleLinksDbpedia This parameter is used to find stories based on the specified entities dbpedia URL in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param NotEntitiesTitleLinksDbpedia This parameter is used to exclude stories based on the specified entities dbpedia URL in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param EntitiesBodyText This parameter is used to find stories based on the specified entities &#x60;text&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param NotEntitiesBodyText This parameter is used to exclude stories based on the specified entities &#x60;text&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param EntitiesBodyType This parameter is used to find stories based on the specified entities &#x60;type&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param NotEntitiesBodyType This parameter is used to exclude stories based on the specified entities &#x60;type&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param EntitiesBodyLinksDbpedia This parameter is used to find stories based on the specified entities dbpedia URL in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param NotEntitiesBodyLinksDbpedia This parameter is used to exclude stories based on the specified entities dbpedia URL in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param SentimentTitlePolarity This parameter is used for finding stories whose title sentiment is the specified value.
 ** @param NotSentimentTitlePolarity This parameter is used for excluding stories whose title sentiment is the specified value.
 ** @param SentimentBodyPolarity This parameter is used for finding stories whose body sentiment is the specified value.
 ** @param NotSentimentBodyPolarity This parameter is used for excluding stories whose body sentiment is the specified value.
 ** @param MediaImagesCountMin This parameter is used for finding stories whose number of images is greater than or equal to the specified value.
 ** @param MediaImagesCountMax This parameter is used for finding stories whose number of images is less than or equal to the specified value.
 ** @param MediaImagesWidthMin This parameter is used for finding stories whose width of images are greater than or equal to the specified value.
 ** @param MediaImagesWidthMax This parameter is used for finding stories whose width of images are less than or equal to the specified value.
 ** @param MediaImagesHeightMin This parameter is used for finding stories whose height of images are greater than or equal to the specified value.
 ** @param MediaImagesHeightMax This parameter is used for finding stories whose height of images are less than or equal to the specified value.
 ** @param MediaImagesContentLengthMin This parameter is used for finding stories whose images content length are greater than or equal to the specified value.
 ** @param MediaImagesContentLengthMax This parameter is used for finding stories whose images content length are less than or equal to the specified value.
 ** @param MediaImagesFormat This parameter is used for finding stories whose images format are the specified value.
 ** @param NotMediaImagesFormat This parameter is used for excluding stories whose images format are the specified value.
 ** @param MediaVideosCountMin This parameter is used for finding stories whose number of videos is greater than or equal to the specified value.
 ** @param MediaVideosCountMax This parameter is used for finding stories whose number of videos is less than or equal to the specified value.
 ** @param AuthorId This parameter is used for finding stories whose author id is the specified value.
 ** @param NotAuthorId This parameter is used for excluding stories whose author id is the specified value.
 ** @param AuthorName This parameter is used for finding stories whose author full name contains the specified value.
 ** @param NotAuthorName This parameter is used for excluding stories whose author full name contains the specified value.
 ** @param SourceId This parameter is used for finding stories whose source id is the specified value.
 ** @param NotSourceId This parameter is used for excluding stories whose source id is the specified value.
 ** @param SourceName This parameter is used for finding stories whose source name contains the specified value.
 ** @param NotSourceName This parameter is used for excluding stories whose source name contains the specified value.
 ** @param SourceDomain This parameter is used for finding stories whose source domain is the specified value.
 ** @param NotSourceDomain This parameter is used for excluding stories whose source domain is the specified value.
 ** @param SourceLocationsCountry This parameter is used for finding stories whose source country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param NotSourceLocationsCountry This parameter is used for excluding stories whose source country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param SourceLocationsState This parameter is used for finding stories whose source state/province is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param NotSourceLocationsState This parameter is used for excluding stories whose source state/province is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param SourceLocationsCity This parameter is used for finding stories whose source city is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param NotSourceLocationsCity This parameter is used for excluding stories whose source city is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param SourceScopesCountry This parameter is used for finding stories whose source scopes is the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param NotSourceScopesCountry This parameter is used for excluding stories whose source scopes is the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param SourceScopesState This parameter is used for finding stories whose source scopes is the specified state/province value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param NotSourceScopesState This parameter is used for excluding stories whose source scopes is the specified state/province value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param SourceScopesCity This parameter is used for finding stories whose source scopes is the specified city value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param NotSourceScopesCity This parameter is used for excluding stories whose source scopes is the specified city value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param SourceScopesLevel This parameter is used for finding stories whose source scopes is the specified level value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param NotSourceScopesLevel This parameter is used for excluding stories whose source scopes is the specified level value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param SourceLinksInCountMin This parameter is used for finding stories from sources whose Links in count is greater than or equal to the specified value. You can read more about working with Links in count [here](https://newsapi.aylien.com/docs/working-with-links-in-count).
 ** @param SourceLinksInCountMax This parameter is used for finding stories from sources whose Links in count is less than or equal to the specified value. You can read more about working with Links in count [here](https://newsapi.aylien.com/docs/working-with-links-in-count).
 ** @param SourceRankingsAlexaRankMin This parameter is used for finding stories from sources whose Alexa rank is greater than or equal to the specified value. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks).
 ** @param SourceRankingsAlexaRankMax This parameter is used for finding stories from sources whose Alexa rank is less than or equal to the specified value. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks).
 ** @param SourceRankingsAlexaCountry This parameter is used for finding stories from sources whose Alexa rank is in the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks).
 ** @param SocialSharesCountFacebookMin This parameter is used for finding stories whose Facebook social shares count is greater than or equal to the specified value.
 ** @param SocialSharesCountFacebookMax This parameter is used for finding stories whose Facebook social shares count is less than or equal to the specified value.
 ** @param SocialSharesCountGooglePlusMin This parameter is used for finding stories whose Google+ social shares count is greater than or equal to the specified value.
 ** @param SocialSharesCountGooglePlusMax This parameter is used for finding stories whose Google+ social shares count is less than or equal to the specified value.
 ** @param SocialSharesCountLinkedinMin This parameter is used for finding stories whose LinkedIn social shares count is greater than or equal to the specified value.
 ** @param SocialSharesCountLinkedinMax This parameter is used for finding stories whose LinkedIn social shares count is less than or equal to the specified value.
 ** @param SocialSharesCountRedditMin This parameter is used for finding stories whose Reddit social shares count is greater than or equal to the specified value.
 ** @param SocialSharesCountRedditMax This parameter is used for finding stories whose Reddit social shares count is less than or equal to the specified value.
 ** @param PublishedAtStart This parameter is used for finding stories whose published at time is less than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates).
 ** @param PublishedAtEnd This parameter is used for finding stories whose published at time is greater than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates).
 ** @param Period The size of each date range is expressed as an interval to be added to the lower bound. It supports Date Math Syntax. Valid options are &#x60;+&#x60; following an integer number greater than 0 and one of the Date Math keywords. e.g. &#x60;+1DAY&#x60;, &#x60;+2MINUTES&#x60; and &#x60;+1MONTH&#x60;. Here are [Supported keywords](https://newsapi.aylien.com/docs/working-with-dates#date-math).
 * @return *TimeSeriesList
 */
func (a DefaultApi) ListTimeSeries(params *TimeSeriesParams) (*TimeSeriesList, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/time_series"

	headerParams := make(map[string]string)
	queryParams := []QueryParams{}
	formParams := []FormParams{}
	// authentication '(app_key)' required
	// set key with prefix in header
	headerParams["X-AYLIEN-NewsAPI-Application-Key"] = a.Configuration.GetAPIKeyWithPrefix("X-AYLIEN-NewsAPI-Application-Key")
	// authentication '(app_id)' required
	// set key with prefix in header
	headerParams["X-AYLIEN-NewsAPI-Application-ID"] = a.Configuration.GetAPIKeyWithPrefix("X-AYLIEN-NewsAPI-Application-ID")
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	for _, f := range params.Id {
		queryParams = append(queryParams, QueryParams{
			Item1: "id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotId {
		queryParams = append(queryParams, QueryParams{
			Item1: "!id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if !isZero(reflect.ValueOf(params.Title)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "title",
			Item2: a.Configuration.APIClient.ParameterToString(params.Title)})
	}
	if !isZero(reflect.ValueOf(params.Body)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "body",
			Item2: a.Configuration.APIClient.ParameterToString(params.Body)})
	}
	if !isZero(reflect.ValueOf(params.Text)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "text",
			Item2: a.Configuration.APIClient.ParameterToString(params.Text)})
	}
	for _, f := range params.Language {
		queryParams = append(queryParams, QueryParams{
			Item1: "language[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotLanguage {
		queryParams = append(queryParams, QueryParams{
			Item1: "!language[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if !isZero(reflect.ValueOf(params.CategoriesTaxonomy)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "categories.taxonomy",
			Item2: a.Configuration.APIClient.ParameterToString(params.CategoriesTaxonomy)})
	}
	if !isZero(reflect.ValueOf(params.CategoriesConfident)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "categories.confident",
			Item2: a.Configuration.APIClient.ParameterToString(params.CategoriesConfident)})
	}
	for _, f := range params.CategoriesId {
		queryParams = append(queryParams, QueryParams{
			Item1: "categories.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotCategoriesId {
		queryParams = append(queryParams, QueryParams{
			Item1: "!categories.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.CategoriesLevel {
		queryParams = append(queryParams, QueryParams{
			Item1: "categories.level[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotCategoriesLevel {
		queryParams = append(queryParams, QueryParams{
			Item1: "!categories.level[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesTitleText {
		queryParams = append(queryParams, QueryParams{
			Item1: "entities.title.text[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotEntitiesTitleText {
		queryParams = append(queryParams, QueryParams{
			Item1: "!entities.title.text[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesTitleType {
		queryParams = append(queryParams, QueryParams{
			Item1: "entities.title.type[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotEntitiesTitleType {
		queryParams = append(queryParams, QueryParams{
			Item1: "!entities.title.type[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesTitleLinksDbpedia {
		queryParams = append(queryParams, QueryParams{
			Item1: "entities.title.links.dbpedia[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotEntitiesTitleLinksDbpedia {
		queryParams = append(queryParams, QueryParams{
			Item1: "!entities.title.links.dbpedia[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesBodyText {
		queryParams = append(queryParams, QueryParams{
			Item1: "entities.body.text[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotEntitiesBodyText {
		queryParams = append(queryParams, QueryParams{
			Item1: "!entities.body.text[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesBodyType {
		queryParams = append(queryParams, QueryParams{
			Item1: "entities.body.type[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotEntitiesBodyType {
		queryParams = append(queryParams, QueryParams{
			Item1: "!entities.body.type[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesBodyLinksDbpedia {
		queryParams = append(queryParams, QueryParams{
			Item1: "entities.body.links.dbpedia[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotEntitiesBodyLinksDbpedia {
		queryParams = append(queryParams, QueryParams{
			Item1: "!entities.body.links.dbpedia[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if !isZero(reflect.ValueOf(params.SentimentTitlePolarity)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "sentiment.title.polarity",
			Item2: a.Configuration.APIClient.ParameterToString(params.SentimentTitlePolarity)})
	}
	if !isZero(reflect.ValueOf(params.NotSentimentTitlePolarity)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "!sentiment.title.polarity",
			Item2: a.Configuration.APIClient.ParameterToString(params.NotSentimentTitlePolarity)})
	}
	if !isZero(reflect.ValueOf(params.SentimentBodyPolarity)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "sentiment.body.polarity",
			Item2: a.Configuration.APIClient.ParameterToString(params.SentimentBodyPolarity)})
	}
	if !isZero(reflect.ValueOf(params.NotSentimentBodyPolarity)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "!sentiment.body.polarity",
			Item2: a.Configuration.APIClient.ParameterToString(params.NotSentimentBodyPolarity)})
	}
	if !isZero(reflect.ValueOf(params.MediaImagesCountMin)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "media.images.count.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaImagesCountMin)})
	}
	if !isZero(reflect.ValueOf(params.MediaImagesCountMax)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "media.images.count.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaImagesCountMax)})
	}
	if !isZero(reflect.ValueOf(params.MediaImagesWidthMin)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "media.images.width.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaImagesWidthMin)})
	}
	if !isZero(reflect.ValueOf(params.MediaImagesWidthMax)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "media.images.width.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaImagesWidthMax)})
	}
	if !isZero(reflect.ValueOf(params.MediaImagesHeightMin)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "media.images.height.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaImagesHeightMin)})
	}
	if !isZero(reflect.ValueOf(params.MediaImagesHeightMax)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "media.images.height.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaImagesHeightMax)})
	}
	if !isZero(reflect.ValueOf(params.MediaImagesContentLengthMin)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "media.images.content_length.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaImagesContentLengthMin)})
	}
	if !isZero(reflect.ValueOf(params.MediaImagesContentLengthMax)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "media.images.content_length.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaImagesContentLengthMax)})
	}
	for _, f := range params.MediaImagesFormat {
		queryParams = append(queryParams, QueryParams{
			Item1: "media.images.format[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotMediaImagesFormat {
		queryParams = append(queryParams, QueryParams{
			Item1: "!media.images.format[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if !isZero(reflect.ValueOf(params.MediaVideosCountMin)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "media.videos.count.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaVideosCountMin)})
	}
	if !isZero(reflect.ValueOf(params.MediaVideosCountMax)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "media.videos.count.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaVideosCountMax)})
	}
	for _, f := range params.AuthorId {
		queryParams = append(queryParams, QueryParams{
			Item1: "author.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotAuthorId {
		queryParams = append(queryParams, QueryParams{
			Item1: "!author.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if !isZero(reflect.ValueOf(params.AuthorName)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "author.name",
			Item2: a.Configuration.APIClient.ParameterToString(params.AuthorName)})
	}
	if !isZero(reflect.ValueOf(params.NotAuthorName)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "!author.name",
			Item2: a.Configuration.APIClient.ParameterToString(params.NotAuthorName)})
	}
	for _, f := range params.SourceId {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceId {
		queryParams = append(queryParams, QueryParams{
			Item1: "!source.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceName {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.name[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceName {
		queryParams = append(queryParams, QueryParams{
			Item1: "!source.name[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceDomain {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.domain[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceDomain {
		queryParams = append(queryParams, QueryParams{
			Item1: "!source.domain[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceLocationsCountry {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.locations.country[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceLocationsCountry {
		queryParams = append(queryParams, QueryParams{
			Item1: "!source.locations.country[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceLocationsState {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.locations.state[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceLocationsState {
		queryParams = append(queryParams, QueryParams{
			Item1: "!source.locations.state[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceLocationsCity {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.locations.city[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceLocationsCity {
		queryParams = append(queryParams, QueryParams{
			Item1: "!source.locations.city[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceScopesCountry {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.scopes.country[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceScopesCountry {
		queryParams = append(queryParams, QueryParams{
			Item1: "!source.scopes.country[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceScopesState {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.scopes.state[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceScopesState {
		queryParams = append(queryParams, QueryParams{
			Item1: "!source.scopes.state[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceScopesCity {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.scopes.city[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceScopesCity {
		queryParams = append(queryParams, QueryParams{
			Item1: "!source.scopes.city[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceScopesLevel {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.scopes.level[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceScopesLevel {
		queryParams = append(queryParams, QueryParams{
			Item1: "!source.scopes.level[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if !isZero(reflect.ValueOf(params.SourceLinksInCountMin)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.links_in_count.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.SourceLinksInCountMin)})
	}
	if !isZero(reflect.ValueOf(params.SourceLinksInCountMax)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.links_in_count.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.SourceLinksInCountMax)})
	}
	if !isZero(reflect.ValueOf(params.SourceRankingsAlexaRankMin)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.rankings.alexa.rank.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.SourceRankingsAlexaRankMin)})
	}
	if !isZero(reflect.ValueOf(params.SourceRankingsAlexaRankMax)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.rankings.alexa.rank.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.SourceRankingsAlexaRankMax)})
	}
	for _, f := range params.SourceRankingsAlexaCountry {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.rankings.alexa.country[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if !isZero(reflect.ValueOf(params.SocialSharesCountFacebookMin)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "social_shares_count.facebook.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.SocialSharesCountFacebookMin)})
	}
	if !isZero(reflect.ValueOf(params.SocialSharesCountFacebookMax)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "social_shares_count.facebook.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.SocialSharesCountFacebookMax)})
	}
	if !isZero(reflect.ValueOf(params.SocialSharesCountGooglePlusMin)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "social_shares_count.google_plus.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.SocialSharesCountGooglePlusMin)})
	}
	if !isZero(reflect.ValueOf(params.SocialSharesCountGooglePlusMax)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "social_shares_count.google_plus.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.SocialSharesCountGooglePlusMax)})
	}
	if !isZero(reflect.ValueOf(params.SocialSharesCountLinkedinMin)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "social_shares_count.linkedin.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.SocialSharesCountLinkedinMin)})
	}
	if !isZero(reflect.ValueOf(params.SocialSharesCountLinkedinMax)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "social_shares_count.linkedin.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.SocialSharesCountLinkedinMax)})
	}
	if !isZero(reflect.ValueOf(params.SocialSharesCountRedditMin)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "social_shares_count.reddit.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.SocialSharesCountRedditMin)})
	}
	if !isZero(reflect.ValueOf(params.SocialSharesCountRedditMax)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "social_shares_count.reddit.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.SocialSharesCountRedditMax)})
	}
	if !isZero(reflect.ValueOf(params.PublishedAtStart)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "published_at.start",
			Item2: a.Configuration.APIClient.ParameterToString(params.PublishedAtStart)})
	}
	if !isZero(reflect.ValueOf(params.PublishedAtEnd)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "published_at.end",
			Item2: a.Configuration.APIClient.ParameterToString(params.PublishedAtEnd)})
	}
	if !isZero(reflect.ValueOf(params.Period)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "period",
			Item2: a.Configuration.APIClient.ParameterToString(params.Period)})
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"text/xml",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(TimeSeriesList)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, headerParams, queryParams, formParams)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse), err
	}

	defer httpResponse.Body.Close()
	resBody, err := ioutil.ReadAll(httpResponse.Body)

	err = json.Unmarshal(resBody, &successPayload)
	return successPayload, NewAPIResponse(httpResponse), err
}

/**
 * List trends
 * This endpoint is used for finding trends based on stories.
 *
 * @param params This is an TrendsParams struct which accepts following parameters:
 ** @param Field This parameter is used to specify the trend field.
 ** @param Id This parameter is used for finding stories by story id.
 ** @param NotId This parameter is used for excluding stories by story id.
 ** @param Title This parameter is used for finding stories whose title contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
 ** @param Body This parameter is used for finding stories whose body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
 ** @param Text This parameter is used for finding stories whose title or body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
 ** @param Language This parameter is used for finding stories whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes.
 ** @param NotLanguage This parameter is used for excluding stories whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes.
 ** @param PublishedAtStart This parameter is used for finding stories whose published at time is greater than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates).
 ** @param PublishedAtEnd This parameter is used for finding stories whose published at time is less than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates).
 ** @param CategoriesTaxonomy This parameter is used for defining the type of the taxonomy for the rest of the categories queries. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories).
 ** @param CategoriesConfident This parameter is used for finding stories whose categories are confident. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories).
 ** @param CategoriesId This parameter is used for finding stories by categories id. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories).
 ** @param NotCategoriesId This parameter is used for excluding stories by categories id. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories).
 ** @param CategoriesLevel This parameter is used for finding stories by categories level. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories).
 ** @param NotCategoriesLevel This parameter is used for excluding stories by categories level. You can read more about working with categories [here](https://newsapi.aylien.com/docs/working-with-categories).
 ** @param EntitiesTitleText This parameter is used to find stories based on the specified entities &#x60;text&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param NotEntitiesTitleText This parameter is used to exclude stories based on the specified entities &#x60;text&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param EntitiesTitleType This parameter is used to find stories based on the specified entities &#x60;type&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param NotEntitiesTitleType This parameter is used to exclude stories based on the specified entities &#x60;type&#x60; in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param EntitiesTitleLinksDbpedia This parameter is used to find stories based on the specified entities dbpedia URL in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param NotEntitiesTitleLinksDbpedia This parameter is used to exclude stories based on the specified entities dbpedia URL in story titles. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param EntitiesBodyText This parameter is used to find stories based on the specified entities &#x60;text&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param NotEntitiesBodyText This parameter is used to exclude stories based on the specified entities &#x60;text&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param EntitiesBodyType This parameter is used to find stories based on the specified entities &#x60;type&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param NotEntitiesBodyType This parameter is used to exclude stories based on the specified entities &#x60;type&#x60; in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param EntitiesBodyLinksDbpedia This parameter is used to find stories based on the specified entities dbpedia URL in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param NotEntitiesBodyLinksDbpedia This parameter is used to exclude stories based on the specified entities dbpedia URL in the body of stories. You can read more about working with entities [here](https://newsapi.aylien.com/docs/working-with-entities).
 ** @param SentimentTitlePolarity This parameter is used for finding stories whose title sentiment is the specified value.
 ** @param NotSentimentTitlePolarity This parameter is used for excluding stories whose title sentiment is the specified value.
 ** @param SentimentBodyPolarity This parameter is used for finding stories whose body sentiment is the specified value.
 ** @param NotSentimentBodyPolarity This parameter is used for excluding stories whose body sentiment is the specified value.
 ** @param MediaImagesCountMin This parameter is used for finding stories whose number of images is greater than or equal to the specified value.
 ** @param MediaImagesCountMax This parameter is used for finding stories whose number of images is less than or equal to the specified value.
 ** @param MediaImagesWidthMin This parameter is used for finding stories whose width of images are greater than or equal to the specified value.
 ** @param MediaImagesWidthMax This parameter is used for finding stories whose width of images are less than or equal to the specified value.
 ** @param MediaImagesHeightMin This parameter is used for finding stories whose height of images are greater than or equal to the specified value.
 ** @param MediaImagesHeightMax This parameter is used for finding stories whose height of images are less than or equal to the specified value.
 ** @param MediaImagesContentLengthMin This parameter is used for finding stories whose images content length are greater than or equal to the specified value.
 ** @param MediaImagesContentLengthMax This parameter is used for finding stories whose images content length are less than or equal to the specified value.
 ** @param MediaImagesFormat This parameter is used for finding stories whose images format are the specified value.
 ** @param NotMediaImagesFormat This parameter is used for excluding stories whose images format are the specified value.
 ** @param MediaVideosCountMin This parameter is used for finding stories whose number of videos is greater than or equal to the specified value.
 ** @param MediaVideosCountMax This parameter is used for finding stories whose number of videos is less than or equal to the specified value.
 ** @param AuthorId This parameter is used for finding stories whose author id is the specified value.
 ** @param NotAuthorId This parameter is used for excluding stories whose author id is the specified value.
 ** @param AuthorName This parameter is used for finding stories whose author full name contains the specified value.
 ** @param NotAuthorName This parameter is used for excluding stories whose author full name contains the specified value.
 ** @param SourceId This parameter is used for finding stories whose source id is the specified value.
 ** @param NotSourceId This parameter is used for excluding stories whose source id is the specified value.
 ** @param SourceName This parameter is used for finding stories whose source name contains the specified value.
 ** @param NotSourceName This parameter is used for excluding stories whose source name contains the specified value.
 ** @param SourceDomain This parameter is used for finding stories whose source domain is the specified value.
 ** @param NotSourceDomain This parameter is used for excluding stories whose source domain is the specified value.
 ** @param SourceLocationsCountry This parameter is used for finding stories whose source country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param NotSourceLocationsCountry This parameter is used for excluding stories whose source country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param SourceLocationsState This parameter is used for finding stories whose source state/province is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param NotSourceLocationsState This parameter is used for excluding stories whose source state/province is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param SourceLocationsCity This parameter is used for finding stories whose source city is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param NotSourceLocationsCity This parameter is used for excluding stories whose source city is the specified value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param SourceScopesCountry This parameter is used for finding stories whose source scopes is the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param NotSourceScopesCountry This parameter is used for excluding stories whose source scopes is the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param SourceScopesState This parameter is used for finding stories whose source scopes is the specified state/province value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param NotSourceScopesState This parameter is used for excluding stories whose source scopes is the specified state/province value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param SourceScopesCity This parameter is used for finding stories whose source scopes is the specified city value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param NotSourceScopesCity This parameter is used for excluding stories whose source scopes is the specified city value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param SourceScopesLevel This parameter is used for finding stories whose source scopes is the specified level value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param NotSourceScopesLevel This parameter is used for excluding stories whose source scopes is the specified level value. [Here](https://newsapi.aylien.com/docs/working-with-locations) you can find more information about how [to work with locations](https://newsapi.aylien.com/docs/working-with-locations).
 ** @param SourceLinksInCountMin This parameter is used for finding stories from sources whose Links in count is greater than or equal to the specified value. You can read more about working with Links in count [here](https://newsapi.aylien.com/docs/working-with-links-in-count).
 ** @param SourceLinksInCountMax This parameter is used for finding stories from sources whose Links in count is less than or equal to the specified value. You can read more about working with Links in count [here](https://newsapi.aylien.com/docs/working-with-links-in-count).
 ** @param SourceRankingsAlexaRankMin This parameter is used for finding stories from sources whose Alexa rank is greater than or equal to the specified value. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks).
 ** @param SourceRankingsAlexaRankMax This parameter is used for finding stories from sources whose Alexa rank is less than or equal to the specified value. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks).
 ** @param SourceRankingsAlexaCountry This parameter is used for finding stories from sources whose Alexa rank is in the specified country value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. You can read more about working with Alexa ranks [here](https://newsapi.aylien.com/docs/working-with-alexa-ranks).
 ** @param SocialSharesCountFacebookMin This parameter is used for finding stories whose Facebook social shares count is greater than or equal to the specified value.
 ** @param SocialSharesCountFacebookMax This parameter is used for finding stories whose Facebook social shares count is less than or equal to the specified value.
 ** @param SocialSharesCountGooglePlusMin This parameter is used for finding stories whose Google+ social shares count is greater than or equal to the specified value.
 ** @param SocialSharesCountGooglePlusMax This parameter is used for finding stories whose Google+ social shares count is less than or equal to the specified value.
 ** @param SocialSharesCountLinkedinMin This parameter is used for finding stories whose LinkedIn social shares count is greater than or equal to the specified value.
 ** @param SocialSharesCountLinkedinMax This parameter is used for finding stories whose LinkedIn social shares count is less than or equal to the specified value.
 ** @param SocialSharesCountRedditMin This parameter is used for finding stories whose Reddit social shares count is greater than or equal to the specified value.
 ** @param SocialSharesCountRedditMax This parameter is used for finding stories whose Reddit social shares count is less than or equal to the specified value.
 * @return *Trends
 */
func (a DefaultApi) ListTrends(params *TrendsParams) (*Trends, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/trends"

	headerParams := make(map[string]string)
	queryParams := []QueryParams{}
	formParams := []FormParams{}
	// authentication '(app_key)' required
	// set key with prefix in header
	headerParams["X-AYLIEN-NewsAPI-Application-Key"] = a.Configuration.GetAPIKeyWithPrefix("X-AYLIEN-NewsAPI-Application-Key")
	// authentication '(app_id)' required
	// set key with prefix in header
	headerParams["X-AYLIEN-NewsAPI-Application-ID"] = a.Configuration.GetAPIKeyWithPrefix("X-AYLIEN-NewsAPI-Application-ID")
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	for _, f := range params.Id {
		queryParams = append(queryParams, QueryParams{
			Item1: "id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotId {
		queryParams = append(queryParams, QueryParams{
			Item1: "!id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if !isZero(reflect.ValueOf(params.Title)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "title",
			Item2: a.Configuration.APIClient.ParameterToString(params.Title)})
	}
	if !isZero(reflect.ValueOf(params.Body)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "body",
			Item2: a.Configuration.APIClient.ParameterToString(params.Body)})
	}
	if !isZero(reflect.ValueOf(params.Text)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "text",
			Item2: a.Configuration.APIClient.ParameterToString(params.Text)})
	}
	for _, f := range params.Language {
		queryParams = append(queryParams, QueryParams{
			Item1: "language[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotLanguage {
		queryParams = append(queryParams, QueryParams{
			Item1: "!language[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if !isZero(reflect.ValueOf(params.PublishedAtStart)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "published_at.start",
			Item2: a.Configuration.APIClient.ParameterToString(params.PublishedAtStart)})
	}
	if !isZero(reflect.ValueOf(params.PublishedAtEnd)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "published_at.end",
			Item2: a.Configuration.APIClient.ParameterToString(params.PublishedAtEnd)})
	}
	if !isZero(reflect.ValueOf(params.CategoriesTaxonomy)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "categories.taxonomy",
			Item2: a.Configuration.APIClient.ParameterToString(params.CategoriesTaxonomy)})
	}
	if !isZero(reflect.ValueOf(params.CategoriesConfident)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "categories.confident",
			Item2: a.Configuration.APIClient.ParameterToString(params.CategoriesConfident)})
	}
	for _, f := range params.CategoriesId {
		queryParams = append(queryParams, QueryParams{
			Item1: "categories.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotCategoriesId {
		queryParams = append(queryParams, QueryParams{
			Item1: "!categories.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.CategoriesLevel {
		queryParams = append(queryParams, QueryParams{
			Item1: "categories.level[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotCategoriesLevel {
		queryParams = append(queryParams, QueryParams{
			Item1: "!categories.level[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesTitleText {
		queryParams = append(queryParams, QueryParams{
			Item1: "entities.title.text[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotEntitiesTitleText {
		queryParams = append(queryParams, QueryParams{
			Item1: "!entities.title.text[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesTitleType {
		queryParams = append(queryParams, QueryParams{
			Item1: "entities.title.type[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotEntitiesTitleType {
		queryParams = append(queryParams, QueryParams{
			Item1: "!entities.title.type[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesTitleLinksDbpedia {
		queryParams = append(queryParams, QueryParams{
			Item1: "entities.title.links.dbpedia[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotEntitiesTitleLinksDbpedia {
		queryParams = append(queryParams, QueryParams{
			Item1: "!entities.title.links.dbpedia[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesBodyText {
		queryParams = append(queryParams, QueryParams{
			Item1: "entities.body.text[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotEntitiesBodyText {
		queryParams = append(queryParams, QueryParams{
			Item1: "!entities.body.text[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesBodyType {
		queryParams = append(queryParams, QueryParams{
			Item1: "entities.body.type[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotEntitiesBodyType {
		queryParams = append(queryParams, QueryParams{
			Item1: "!entities.body.type[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesBodyLinksDbpedia {
		queryParams = append(queryParams, QueryParams{
			Item1: "entities.body.links.dbpedia[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotEntitiesBodyLinksDbpedia {
		queryParams = append(queryParams, QueryParams{
			Item1: "!entities.body.links.dbpedia[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if !isZero(reflect.ValueOf(params.SentimentTitlePolarity)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "sentiment.title.polarity",
			Item2: a.Configuration.APIClient.ParameterToString(params.SentimentTitlePolarity)})
	}
	if !isZero(reflect.ValueOf(params.NotSentimentTitlePolarity)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "!sentiment.title.polarity",
			Item2: a.Configuration.APIClient.ParameterToString(params.NotSentimentTitlePolarity)})
	}
	if !isZero(reflect.ValueOf(params.SentimentBodyPolarity)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "sentiment.body.polarity",
			Item2: a.Configuration.APIClient.ParameterToString(params.SentimentBodyPolarity)})
	}
	if !isZero(reflect.ValueOf(params.NotSentimentBodyPolarity)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "!sentiment.body.polarity",
			Item2: a.Configuration.APIClient.ParameterToString(params.NotSentimentBodyPolarity)})
	}
	if !isZero(reflect.ValueOf(params.MediaImagesCountMin)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "media.images.count.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaImagesCountMin)})
	}
	if !isZero(reflect.ValueOf(params.MediaImagesCountMax)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "media.images.count.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaImagesCountMax)})
	}
	if !isZero(reflect.ValueOf(params.MediaImagesWidthMin)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "media.images.width.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaImagesWidthMin)})
	}
	if !isZero(reflect.ValueOf(params.MediaImagesWidthMax)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "media.images.width.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaImagesWidthMax)})
	}
	if !isZero(reflect.ValueOf(params.MediaImagesHeightMin)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "media.images.height.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaImagesHeightMin)})
	}
	if !isZero(reflect.ValueOf(params.MediaImagesHeightMax)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "media.images.height.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaImagesHeightMax)})
	}
	if !isZero(reflect.ValueOf(params.MediaImagesContentLengthMin)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "media.images.content_length.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaImagesContentLengthMin)})
	}
	if !isZero(reflect.ValueOf(params.MediaImagesContentLengthMax)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "media.images.content_length.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaImagesContentLengthMax)})
	}
	for _, f := range params.MediaImagesFormat {
		queryParams = append(queryParams, QueryParams{
			Item1: "media.images.format[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotMediaImagesFormat {
		queryParams = append(queryParams, QueryParams{
			Item1: "!media.images.format[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if !isZero(reflect.ValueOf(params.MediaVideosCountMin)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "media.videos.count.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaVideosCountMin)})
	}
	if !isZero(reflect.ValueOf(params.MediaVideosCountMax)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "media.videos.count.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.MediaVideosCountMax)})
	}
	for _, f := range params.AuthorId {
		queryParams = append(queryParams, QueryParams{
			Item1: "author.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotAuthorId {
		queryParams = append(queryParams, QueryParams{
			Item1: "!author.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if !isZero(reflect.ValueOf(params.AuthorName)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "author.name",
			Item2: a.Configuration.APIClient.ParameterToString(params.AuthorName)})
	}
	if !isZero(reflect.ValueOf(params.NotAuthorName)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "!author.name",
			Item2: a.Configuration.APIClient.ParameterToString(params.NotAuthorName)})
	}
	for _, f := range params.SourceId {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceId {
		queryParams = append(queryParams, QueryParams{
			Item1: "!source.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceName {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.name[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceName {
		queryParams = append(queryParams, QueryParams{
			Item1: "!source.name[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceDomain {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.domain[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceDomain {
		queryParams = append(queryParams, QueryParams{
			Item1: "!source.domain[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceLocationsCountry {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.locations.country[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceLocationsCountry {
		queryParams = append(queryParams, QueryParams{
			Item1: "!source.locations.country[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceLocationsState {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.locations.state[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceLocationsState {
		queryParams = append(queryParams, QueryParams{
			Item1: "!source.locations.state[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceLocationsCity {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.locations.city[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceLocationsCity {
		queryParams = append(queryParams, QueryParams{
			Item1: "!source.locations.city[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceScopesCountry {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.scopes.country[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceScopesCountry {
		queryParams = append(queryParams, QueryParams{
			Item1: "!source.scopes.country[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceScopesState {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.scopes.state[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceScopesState {
		queryParams = append(queryParams, QueryParams{
			Item1: "!source.scopes.state[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceScopesCity {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.scopes.city[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceScopesCity {
		queryParams = append(queryParams, QueryParams{
			Item1: "!source.scopes.city[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceScopesLevel {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.scopes.level[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.NotSourceScopesLevel {
		queryParams = append(queryParams, QueryParams{
			Item1: "!source.scopes.level[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if !isZero(reflect.ValueOf(params.SourceLinksInCountMin)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.links_in_count.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.SourceLinksInCountMin)})
	}
	if !isZero(reflect.ValueOf(params.SourceLinksInCountMax)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.links_in_count.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.SourceLinksInCountMax)})
	}
	if !isZero(reflect.ValueOf(params.SourceRankingsAlexaRankMin)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.rankings.alexa.rank.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.SourceRankingsAlexaRankMin)})
	}
	if !isZero(reflect.ValueOf(params.SourceRankingsAlexaRankMax)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.rankings.alexa.rank.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.SourceRankingsAlexaRankMax)})
	}
	for _, f := range params.SourceRankingsAlexaCountry {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.rankings.alexa.country[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if !isZero(reflect.ValueOf(params.SocialSharesCountFacebookMin)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "social_shares_count.facebook.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.SocialSharesCountFacebookMin)})
	}
	if !isZero(reflect.ValueOf(params.SocialSharesCountFacebookMax)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "social_shares_count.facebook.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.SocialSharesCountFacebookMax)})
	}
	if !isZero(reflect.ValueOf(params.SocialSharesCountGooglePlusMin)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "social_shares_count.google_plus.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.SocialSharesCountGooglePlusMin)})
	}
	if !isZero(reflect.ValueOf(params.SocialSharesCountGooglePlusMax)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "social_shares_count.google_plus.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.SocialSharesCountGooglePlusMax)})
	}
	if !isZero(reflect.ValueOf(params.SocialSharesCountLinkedinMin)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "social_shares_count.linkedin.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.SocialSharesCountLinkedinMin)})
	}
	if !isZero(reflect.ValueOf(params.SocialSharesCountLinkedinMax)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "social_shares_count.linkedin.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.SocialSharesCountLinkedinMax)})
	}
	if !isZero(reflect.ValueOf(params.SocialSharesCountRedditMin)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "social_shares_count.reddit.min",
			Item2: a.Configuration.APIClient.ParameterToString(params.SocialSharesCountRedditMin)})
	}
	if !isZero(reflect.ValueOf(params.SocialSharesCountRedditMax)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "social_shares_count.reddit.max",
			Item2: a.Configuration.APIClient.ParameterToString(params.SocialSharesCountRedditMax)})
	}
	if !isZero(reflect.ValueOf(params.Field)) {
		queryParams = append(queryParams, QueryParams{
			Item1: "field",
			Item2: a.Configuration.APIClient.ParameterToString(params.Field)})
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"text/xml",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(Trends)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, headerParams, queryParams, formParams)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse), err
	}

	defer httpResponse.Body.Close()
	resBody, err := ioutil.ReadAll(httpResponse.Body)

	err = json.Unmarshal(resBody, &successPayload)
	return successPayload, NewAPIResponse(httpResponse), err
}
