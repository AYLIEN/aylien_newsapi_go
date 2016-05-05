/*
Copyright 2016 Aylien, Inc. All Rights Reserved.

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
	"errors"
	"io/ioutil"
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
 ** @param Term This parameter is used for finding autocomplete objects whose contains the specified value.
 ** @param Language This parameter is used for autocompletes whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes.
 ** @param PerPage This parameter is used for specifying number of the items in each page.
 * @return Autocompletes
 */
func (a DefaultApi) ListAutocompletes(params *AutocompletesParams) (*Autocompletes, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/autocompletes"

	// verify the required parameter 'type_' is set
	if len(params.Type) == 0 {
		return new(Autocompletes), nil, errors.New("Missing required parameter 'Type' when calling DefaultApi->ListAutocompletes")
	}
	// verify the required parameter 'term' is set
	if len(params.Term) == 0 {
		return new(Autocompletes), nil, errors.New("Missing required parameter 'Term' when calling DefaultApi->ListAutocompletes")
	}

	headerParams := make(map[string]string)
	queryParams := []QueryParams{}
	formParams := []FormParams{}
	// authentication (app_key) required

	// set key with prefix in header
	headerParams["X-AYLIEN-NewsAPI-Application-Key"] = a.Configuration.GetAPIKeyWithPrefix("X-AYLIEN-NewsAPI-Application-Key")
	// authentication (app_id) required

	// set key with prefix in header
	headerParams["X-AYLIEN-NewsAPI-Application-ID"] = a.Configuration.GetAPIKeyWithPrefix("X-AYLIEN-NewsAPI-Application-ID")

	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}

	queryParams = append(queryParams, QueryParams{
		Item1: "type",
		Item2: a.Configuration.APIClient.ParameterToString(params.Type)})

	queryParams = append(queryParams, QueryParams{
		Item1: "term",
		Item2: a.Configuration.APIClient.ParameterToString(params.Term)})

	queryParams = append(queryParams, QueryParams{
		Item1: "language",
		Item2: a.Configuration.APIClient.ParameterToString(params.Language)})

	queryParams = append(queryParams, QueryParams{
		Item1: "perPage",
		Item2: a.Configuration.APIClient.ParameterToString(params.PerPage)})

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
 * This endpoint is used for finding story coverages based on provided parameters. The number of coverages to return, up to a maximum of 100.
 *
 * @param params This is an CoveragesParams struct which accepts following parameters:
 ** @param Id This parameter is used for finding stroies by story id.
 ** @param Title This parameter is used for finding stories whose title contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
 ** @param Body This parameter is used for finding stories whose body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
 ** @param Text This parameter is used for finding stories whose title or body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
 ** @param Language This parameter is used for finding stories whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes.
 ** @param PublishedAtStart This parameter is used for finding stories whose published at is greater than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates).
 ** @param PublishedAtEnd This parameter is used for finding stories whose published at is less than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates).
 ** @param CategoriesTaxonomy This parameter is used for defining type of the taxonomy for the rest of categories queries.
 ** @param CategoriesConfident This parameter is used for finding stories whose categories are confident.
 ** @param CategoriesId This parameter is used for finding stories whose categories id is the specified value.
 ** @param CategoriesLevel This parameter is used for finding stories whose categories level is the specified value.
 ** @param EntitiesTitleText This parameter is used for finding stories whose entities text in title is the specified value.
 ** @param EntitiesTitleType This parameter is used for finding stories whose entities type in title is the specified value.
 ** @param EntitiesTitleLinksDbpedia This parameter is used for finding stories whose entities dbpedia url in title is the specified value.
 ** @param EntitiesBodyText This parameter is used for finding stories whose entities text in body is the specified value.
 ** @param EntitiesBodyType This parameter is used for finding stories whose entities type in body is the specified value.
 ** @param EntitiesBodyLinksDbpedia This parameter is used for finding stories whose entities dbpedia url in body is the specified value.
 ** @param SentimentTitlePolarity This parameter is used for finding stories whose title sentiment is the specified value.
 ** @param SentimentBodyPolarity This parameter is used for finding stories whose body sentiment is the specified value.
 ** @param AuthorId This parameter is used for finding stories whose author id is the specified value.
 ** @param AuthorName This parameter is used for finding stories whose author full name contains the specified value.
 ** @param SourceId This parameter is used for finding stories whose source id is the specified value.
 ** @param SourceName This parameter is used for finding stories whose source name contains the specified value.
 ** @param SourceDomain This parameter is used for finding stories whose source domain is the specified value.
 ** @param SourceLocationsCountry This parameter is used for finding stories whose source country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes.
 ** @param SourceLocationsState This parameter is used for finding stories whose source state/province is the specified value.
 ** @param SourceLocationsCity This parameter is used for finding stories whose source city is the specified value.
 ** @param SourceScopesCountry This parameter is used for finding stories whose source scopes country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes.
 ** @param SourceScopesState This parameter is used for finding stories whose source scopes state/province is the specified value.
 ** @param SourceScopesCity This parameter is used for finding stories whose source scopes city is the specified value.
 ** @param SourceScopesLevel This parameter is used for finding stories whose source scopes level is the specified value.
 ** @param Return This parameter is used for specifying return fields.
 ** @param StoryId A story id
 ** @param StoryUrl An article or webpage
 ** @param StoryTitle Title of the article
 ** @param StoryBody Body of the article
 ** @param StoryPublishedAt Publish date of the article. If you use url or title and body for getting coverages, this parameter is required. The format used is a restricted form of the canonical representation of dateTime in the [XML Schema specification (ISO 8601)](https://www.w3.org/TR/xmlschema-2/#dateTime).
 ** @param StoryLanguage This parameter is used for setting language of the story. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes.
 ** @param PerPage This parameter is used for specifying number of the items in each page.
 * @return Coverages
 */
func (a DefaultApi) ListCoverages(params *CoveragesParams) (*Coverages, *APIResponse, error) {

	var httpMethod = "Post"
	// create path and map variables
	path := a.Configuration.BasePath + "/coverages"

	headerParams := make(map[string]string)
	queryParams := []QueryParams{}
	formParams := []FormParams{}
	// authentication (app_key) required

	// set key with prefix in header
	headerParams["X-AYLIEN-NewsAPI-Application-Key"] = a.Configuration.GetAPIKeyWithPrefix("X-AYLIEN-NewsAPI-Application-Key")
	// authentication (app_id) required

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

	if len(params.Title) > 0 {
		formParams = append(formParams, FormParams{
			Item1: "title",
			Item2: a.Configuration.APIClient.ParameterToString(params.Title)})
	}

	if len(params.Body) > 0 {
		formParams = append(formParams, FormParams{
			Item1: "body",
			Item2: a.Configuration.APIClient.ParameterToString(params.Body)})
	}

	if len(params.Text) > 0 {
		formParams = append(formParams, FormParams{
			Item1: "text",
			Item2: a.Configuration.APIClient.ParameterToString(params.Text)})
	}

	for _, f := range params.Language {
		formParams = append(formParams, FormParams{
			Item1: "language[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if len(params.PublishedAtStart) > 0 {
		formParams = append(formParams, FormParams{
			Item1: "published_at.start",
			Item2: a.Configuration.APIClient.ParameterToString(params.PublishedAtStart)})
	}

	if len(params.PublishedAtEnd) > 0 {
		formParams = append(formParams, FormParams{
			Item1: "published_at.end",
			Item2: a.Configuration.APIClient.ParameterToString(params.PublishedAtEnd)})
	}

	if len(params.CategoriesTaxonomy) > 0 {
		formParams = append(formParams, FormParams{
			Item1: "categories.taxonomy",
			Item2: a.Configuration.APIClient.ParameterToString(params.CategoriesTaxonomy)})
	}

	if len(params.CategoriesConfident) > 0 {
		formParams = append(formParams, FormParams{
			Item1: "categories.confident",
			Item2: a.Configuration.APIClient.ParameterToString(params.CategoriesConfident)})
	}

	for _, f := range params.CategoriesId {
		formParams = append(formParams, FormParams{
			Item1: "categories.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.CategoriesLevel {
		formParams = append(formParams, FormParams{
			Item1: "categories.level[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesTitleText {
		formParams = append(formParams, FormParams{
			Item1: "entities.title.text[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesTitleType {
		formParams = append(formParams, FormParams{
			Item1: "entities.title.type[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesTitleLinksDbpedia {
		formParams = append(formParams, FormParams{
			Item1: "entities.title.links.dbpedia[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesBodyText {
		formParams = append(formParams, FormParams{
			Item1: "entities.body.text[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesBodyType {
		formParams = append(formParams, FormParams{
			Item1: "entities.body.type[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesBodyLinksDbpedia {
		formParams = append(formParams, FormParams{
			Item1: "entities.body.links.dbpedia[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if len(params.SentimentTitlePolarity) > 0 {
		formParams = append(formParams, FormParams{
			Item1: "sentiment.title.polarity",
			Item2: a.Configuration.APIClient.ParameterToString(params.SentimentTitlePolarity)})
	}

	if len(params.SentimentBodyPolarity) > 0 {
		formParams = append(formParams, FormParams{
			Item1: "sentiment.body.polarity",
			Item2: a.Configuration.APIClient.ParameterToString(params.SentimentBodyPolarity)})
	}

	for _, f := range params.AuthorId {
		formParams = append(formParams, FormParams{
			Item1: "author.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if len(params.AuthorName) > 0 {
		formParams = append(formParams, FormParams{
			Item1: "author.name",
			Item2: a.Configuration.APIClient.ParameterToString(params.AuthorName)})
	}

	for _, f := range params.SourceId {
		formParams = append(formParams, FormParams{
			Item1: "source.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceName {
		formParams = append(formParams, FormParams{
			Item1: "source.name[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceDomain {
		formParams = append(formParams, FormParams{
			Item1: "source.domain[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceLocationsCountry {
		formParams = append(formParams, FormParams{
			Item1: "source.locations.country[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceLocationsState {
		formParams = append(formParams, FormParams{
			Item1: "source.locations.state[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceLocationsCity {
		formParams = append(formParams, FormParams{
			Item1: "source.locations.city[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceScopesCountry {
		formParams = append(formParams, FormParams{
			Item1: "source.scopes.country[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceScopesState {
		formParams = append(formParams, FormParams{
			Item1: "source.scopes.state[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceScopesCity {
		formParams = append(formParams, FormParams{
			Item1: "source.scopes.city[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceScopesLevel {
		formParams = append(formParams, FormParams{
			Item1: "source.scopes.level[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.Return {
		formParams = append(formParams, FormParams{
			Item1: "return[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if params.StoryId != 0 {
		formParams = append(formParams, FormParams{
			Item1: "story_id",
			Item2: a.Configuration.APIClient.ParameterToString(params.StoryId)})
	}

	if len(params.StoryUrl) > 0 {
		formParams = append(formParams, FormParams{
			Item1: "story_url",
			Item2: a.Configuration.APIClient.ParameterToString(params.StoryUrl)})
	}

	if len(params.StoryTitle) > 0 {
		formParams = append(formParams, FormParams{
			Item1: "story_title",
			Item2: a.Configuration.APIClient.ParameterToString(params.StoryTitle)})
	}

	if len(params.StoryBody) > 0 {
		formParams = append(formParams, FormParams{
			Item1: "story_body",
			Item2: a.Configuration.APIClient.ParameterToString(params.StoryBody)})
	}

	if params.StoryPublishedAt.String() != "0001-01-01 00:00:00 +0000 UTC" {
		formParams = append(formParams, FormParams{
			Item1: "story_published_at",
			Item2: a.Configuration.APIClient.ParameterToString(params.StoryPublishedAt)})
	}

	if len(params.StoryLanguage) > 0 {
		formParams = append(formParams, FormParams{
			Item1: "story_language",
			Item2: a.Configuration.APIClient.ParameterToString(params.StoryLanguage)})
	}

	if params.PerPage != 0 {
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
 * This endpoint is used for getting histograms based on the field parameter is passed.
 *
 * @param params This is an HistogramsParams struct which accepts following parameters:
 ** @param Id This parameter is used for finding stroies by story id.
 ** @param Title This parameter is used for finding stories whose title contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
 ** @param Body This parameter is used for finding stories whose body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
 ** @param Text This parameter is used for finding stories whose title or body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
 ** @param Language This parameter is used for finding stories whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes.
 ** @param PublishedAtStart This parameter is used for finding stories whose published at is greater than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates).
 ** @param PublishedAtEnd This parameter is used for finding stories whose published at is less than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates).
 ** @param CategoriesTaxonomy This parameter is used for defining type of the taxonomy for the rest of categories queries.
 ** @param CategoriesConfident This parameter is used for finding stories whose categories are confident.
 ** @param CategoriesId This parameter is used for finding stories whose categories id is the specified value.
 ** @param CategoriesLevel This parameter is used for finding stories whose categories level is the specified value.
 ** @param EntitiesTitleText This parameter is used for finding stories whose entities text in title is the specified value.
 ** @param EntitiesTitleType This parameter is used for finding stories whose entities type in title is the specified value.
 ** @param EntitiesTitleLinksDbpedia This parameter is used for finding stories whose entities dbpedia url in title is the specified value.
 ** @param EntitiesBodyText This parameter is used for finding stories whose entities text in body is the specified value.
 ** @param EntitiesBodyType This parameter is used for finding stories whose entities type in body is the specified value.
 ** @param EntitiesBodyLinksDbpedia This parameter is used for finding stories whose entities dbpedia url in body is the specified value.
 ** @param SentimentTitlePolarity This parameter is used for finding stories whose title sentiment is the specified value.
 ** @param SentimentBodyPolarity This parameter is used for finding stories whose body sentiment is the specified value.
 ** @param AuthorId This parameter is used for finding stories whose author id is the specified value.
 ** @param AuthorName This parameter is used for finding stories whose author full name contains the specified value.
 ** @param SourceId This parameter is used for finding stories whose source id is the specified value.
 ** @param SourceName This parameter is used for finding stories whose source name contains the specified value.
 ** @param SourceDomain This parameter is used for finding stories whose source domain is the specified value.
 ** @param SourceLocationsCountry This parameter is used for finding stories whose source country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes.
 ** @param SourceLocationsState This parameter is used for finding stories whose source state/province is the specified value.
 ** @param SourceLocationsCity This parameter is used for finding stories whose source city is the specified value.
 ** @param SourceScopesCountry This parameter is used for finding stories whose source scopes country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes.
 ** @param SourceScopesState This parameter is used for finding stories whose source scopes state/province is the specified value.
 ** @param SourceScopesCity This parameter is used for finding stories whose source scopes city is the specified value.
 ** @param SourceScopesLevel This parameter is used for finding stories whose source scopes level is the specified value.
 ** @param IntervalStart This parameter is used for setting the start data point of histogram intervals.
 ** @param IntervalEnd This parameter is used for setting the end data point of histogram intervals.
 ** @param IntervalWidth This parameter is used for setting the width of histogram intervals.
 ** @param Field This parameter is used for specifying the y-axis variable for the histogram.
 * @return Histograms
 */
func (a DefaultApi) ListHistograms(params *HistogramsParams) (*Histograms, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/histograms"

	headerParams := make(map[string]string)
	queryParams := []QueryParams{}
	formParams := []FormParams{}
	// authentication (app_key) required

	// set key with prefix in header
	headerParams["X-AYLIEN-NewsAPI-Application-Key"] = a.Configuration.GetAPIKeyWithPrefix("X-AYLIEN-NewsAPI-Application-Key")
	// authentication (app_id) required

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

	if len(params.Title) > 0 {
		queryParams = append(queryParams, QueryParams{
			Item1: "title",
			Item2: a.Configuration.APIClient.ParameterToString(params.Title)})
	}

	if len(params.Body) > 0 {
		queryParams = append(queryParams, QueryParams{
			Item1: "body",
			Item2: a.Configuration.APIClient.ParameterToString(params.Body)})
	}

	if len(params.Text) > 0 {
		queryParams = append(queryParams, QueryParams{
			Item1: "text",
			Item2: a.Configuration.APIClient.ParameterToString(params.Text)})
	}

	for _, f := range params.Language {
		queryParams = append(queryParams, QueryParams{
			Item1: "language[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if len(params.PublishedAtStart) > 0 {
		queryParams = append(queryParams, QueryParams{
			Item1: "published_at.start",
			Item2: a.Configuration.APIClient.ParameterToString(params.PublishedAtStart)})
	}

	if len(params.PublishedAtEnd) > 0 {
		queryParams = append(queryParams, QueryParams{
			Item1: "published_at.end",
			Item2: a.Configuration.APIClient.ParameterToString(params.PublishedAtEnd)})
	}

	if len(params.CategoriesTaxonomy) > 0 {
		queryParams = append(queryParams, QueryParams{
			Item1: "categories.taxonomy",
			Item2: a.Configuration.APIClient.ParameterToString(params.CategoriesTaxonomy)})
	}

	if len(params.CategoriesConfident) > 0 {
		queryParams = append(queryParams, QueryParams{
			Item1: "categories.confident",
			Item2: a.Configuration.APIClient.ParameterToString(params.CategoriesConfident)})
	}

	for _, f := range params.CategoriesId {
		queryParams = append(queryParams, QueryParams{
			Item1: "categories.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.CategoriesLevel {
		queryParams = append(queryParams, QueryParams{
			Item1: "categories.level[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesTitleText {
		queryParams = append(queryParams, QueryParams{
			Item1: "entities.title.text[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesTitleType {
		queryParams = append(queryParams, QueryParams{
			Item1: "entities.title.type[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesTitleLinksDbpedia {
		queryParams = append(queryParams, QueryParams{
			Item1: "entities.title.links.dbpedia[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesBodyText {
		queryParams = append(queryParams, QueryParams{
			Item1: "entities.body.text[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesBodyType {
		queryParams = append(queryParams, QueryParams{
			Item1: "entities.body.type[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesBodyLinksDbpedia {
		queryParams = append(queryParams, QueryParams{
			Item1: "entities.body.links.dbpedia[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if len(params.SentimentTitlePolarity) > 0 {
		queryParams = append(queryParams, QueryParams{
			Item1: "sentiment.title.polarity",
			Item2: a.Configuration.APIClient.ParameterToString(params.SentimentTitlePolarity)})
	}

	if len(params.SentimentBodyPolarity) > 0 {
		queryParams = append(queryParams, QueryParams{
			Item1: "sentiment.body.polarity",
			Item2: a.Configuration.APIClient.ParameterToString(params.SentimentBodyPolarity)})
	}

	for _, f := range params.AuthorId {
		queryParams = append(queryParams, QueryParams{
			Item1: "author.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if len(params.AuthorName) > 0 {
		queryParams = append(queryParams, QueryParams{
			Item1: "author.name",
			Item2: a.Configuration.APIClient.ParameterToString(params.AuthorName)})
	}

	for _, f := range params.SourceId {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceName {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.name[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceDomain {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.domain[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceLocationsCountry {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.locations.country[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceLocationsState {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.locations.state[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceLocationsCity {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.locations.city[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceScopesCountry {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.scopes.country[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceScopesState {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.scopes.state[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceScopesCity {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.scopes.city[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceScopesLevel {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.scopes.level[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if params.IntervalStart != 0 {
		queryParams = append(queryParams, QueryParams{
			Item1: "interval.start",
			Item2: a.Configuration.APIClient.ParameterToString(params.IntervalStart)})
	}

	if params.IntervalEnd != 0 {
		queryParams = append(queryParams, QueryParams{
			Item1: "interval.end",
			Item2: a.Configuration.APIClient.ParameterToString(params.IntervalEnd)})
	}

	if params.IntervalWidth != 0 {
		queryParams = append(queryParams, QueryParams{
			Item1: "interval.width",
			Item2: a.Configuration.APIClient.ParameterToString(params.IntervalWidth)})
	}

	if len(params.Field) > 0 {
		queryParams = append(queryParams, QueryParams{
			Item1: "interval.field",
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
 * This endpoint is used for finding related stories based on provided parameters. The number of related stories to return, up to a maximum of 100.
 *
 * @param params This is an RelatedStoriesParams struct which accepts following parameters:
 ** @param Id This parameter is used for finding stroies by story id.
 ** @param Title This parameter is used for finding stories whose title contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
 ** @param Body This parameter is used for finding stories whose body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
 ** @param Text This parameter is used for finding stories whose title or body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
 ** @param Language This parameter is used for finding stories whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes.
 ** @param PublishedAtStart This parameter is used for finding stories whose published at is greater than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates).
 ** @param PublishedAtEnd This parameter is used for finding stories whose published at is less than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates).
 ** @param CategoriesTaxonomy This parameter is used for defining type of the taxonomy for the rest of categories queries.
 ** @param CategoriesConfident This parameter is used for finding stories whose categories are confident.
 ** @param CategoriesId This parameter is used for finding stories whose categories id is the specified value.
 ** @param CategoriesLevel This parameter is used for finding stories whose categories level is the specified value.
 ** @param EntitiesTitleText This parameter is used for finding stories whose entities text in title is the specified value.
 ** @param EntitiesTitleType This parameter is used for finding stories whose entities type in title is the specified value.
 ** @param EntitiesTitleLinksDbpedia This parameter is used for finding stories whose entities dbpedia url in title is the specified value.
 ** @param EntitiesBodyText This parameter is used for finding stories whose entities text in body is the specified value.
 ** @param EntitiesBodyType This parameter is used for finding stories whose entities type in body is the specified value.
 ** @param EntitiesBodyLinksDbpedia This parameter is used for finding stories whose entities dbpedia url in body is the specified value.
 ** @param SentimentTitlePolarity This parameter is used for finding stories whose title sentiment is the specified value.
 ** @param SentimentBodyPolarity This parameter is used for finding stories whose body sentiment is the specified value.
 ** @param AuthorId This parameter is used for finding stories whose author id is the specified value.
 ** @param AuthorName This parameter is used for finding stories whose author full name contains the specified value.
 ** @param SourceId This parameter is used for finding stories whose source id is the specified value.
 ** @param SourceName This parameter is used for finding stories whose source name contains the specified value.
 ** @param SourceDomain This parameter is used for finding stories whose source domain is the specified value.
 ** @param SourceLocationsCountry This parameter is used for finding stories whose source country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes.
 ** @param SourceLocationsState This parameter is used for finding stories whose source state/province is the specified value.
 ** @param SourceLocationsCity This parameter is used for finding stories whose source city is the specified value.
 ** @param SourceScopesCountry This parameter is used for finding stories whose source scopes country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes.
 ** @param SourceScopesState This parameter is used for finding stories whose source scopes state/province is the specified value.
 ** @param SourceScopesCity This parameter is used for finding stories whose source scopes city is the specified value.
 ** @param SourceScopesLevel This parameter is used for finding stories whose source scopes level is the specified value.
 ** @param Return This parameter is used for specifying return fields.
 ** @param StoryId A story id
 ** @param StoryUrl An article or webpage
 ** @param StoryTitle Title of the article
 ** @param StoryBody Body of the article
 ** @param BoostBy This parameter is used for boosting result by the specified value.
 ** @param StoryLanguage This parameter is used for setting language of the story. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes.
 ** @param PerPage This parameter is used for specifying number of the items in each page.
 * @return RelatedStories
 */
func (a DefaultApi) ListRelatedStories(params *RelatedStoriesParams) (*RelatedStories, *APIResponse, error) {

	var httpMethod = "Post"
	// create path and map variables
	path := a.Configuration.BasePath + "/related_stories"

	headerParams := make(map[string]string)
	queryParams := []QueryParams{}
	formParams := []FormParams{}
	// authentication (app_key) required

	// set key with prefix in header
	headerParams["X-AYLIEN-NewsAPI-Application-Key"] = a.Configuration.GetAPIKeyWithPrefix("X-AYLIEN-NewsAPI-Application-Key")
	// authentication (app_id) required

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

	if len(params.Title) > 0 {
		formParams = append(formParams, FormParams{
			Item1: "title",
			Item2: a.Configuration.APIClient.ParameterToString(params.Title)})
	}

	if len(params.Body) > 0 {
		formParams = append(formParams, FormParams{
			Item1: "body",
			Item2: a.Configuration.APIClient.ParameterToString(params.Body)})
	}

	if len(params.Text) > 0 {
		formParams = append(formParams, FormParams{
			Item1: "text",
			Item2: a.Configuration.APIClient.ParameterToString(params.Text)})
	}

	for _, f := range params.Language {
		formParams = append(formParams, FormParams{
			Item1: "language[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if len(params.PublishedAtStart) > 0 {
		formParams = append(formParams, FormParams{
			Item1: "published_at.start",
			Item2: a.Configuration.APIClient.ParameterToString(params.PublishedAtStart)})
	}

	if len(params.PublishedAtEnd) > 0 {
		formParams = append(formParams, FormParams{
			Item1: "published_at.end",
			Item2: a.Configuration.APIClient.ParameterToString(params.PublishedAtEnd)})
	}

	if len(params.CategoriesTaxonomy) > 0 {
		formParams = append(formParams, FormParams{
			Item1: "categories.taxonomy",
			Item2: a.Configuration.APIClient.ParameterToString(params.CategoriesTaxonomy)})
	}

	if len(params.CategoriesConfident) > 0 {
		formParams = append(formParams, FormParams{
			Item1: "categories.confident",
			Item2: a.Configuration.APIClient.ParameterToString(params.CategoriesConfident)})
	}

	for _, f := range params.CategoriesId {
		formParams = append(formParams, FormParams{
			Item1: "categories.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.CategoriesLevel {
		formParams = append(formParams, FormParams{
			Item1: "categories.level[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesTitleText {
		formParams = append(formParams, FormParams{
			Item1: "entities.title.text[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesTitleType {
		formParams = append(formParams, FormParams{
			Item1: "entities.title.type[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesTitleLinksDbpedia {
		formParams = append(formParams, FormParams{
			Item1: "entities.title.links.dbpedia[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesBodyText {
		formParams = append(formParams, FormParams{
			Item1: "entities.body.text[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesBodyType {
		formParams = append(formParams, FormParams{
			Item1: "entities.body.type[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesBodyLinksDbpedia {
		formParams = append(formParams, FormParams{
			Item1: "entities.body.links.dbpedia[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if len(params.SentimentTitlePolarity) > 0 {
		formParams = append(formParams, FormParams{
			Item1: "sentiment.title.polarity",
			Item2: a.Configuration.APIClient.ParameterToString(params.SentimentTitlePolarity)})
	}

	if len(params.SentimentBodyPolarity) > 0 {
		formParams = append(formParams, FormParams{
			Item1: "sentiment.body.polarity",
			Item2: a.Configuration.APIClient.ParameterToString(params.SentimentBodyPolarity)})
	}

	for _, f := range params.AuthorId {
		formParams = append(formParams, FormParams{
			Item1: "author.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if len(params.AuthorName) > 0 {
		formParams = append(formParams, FormParams{
			Item1: "author.name",
			Item2: a.Configuration.APIClient.ParameterToString(params.AuthorName)})
	}

	for _, f := range params.SourceId {
		formParams = append(formParams, FormParams{
			Item1: "source.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceName {
		formParams = append(formParams, FormParams{
			Item1: "source.name[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceDomain {
		formParams = append(formParams, FormParams{
			Item1: "source.domain[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceLocationsCountry {
		formParams = append(formParams, FormParams{
			Item1: "source.locations.country[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceLocationsState {
		formParams = append(formParams, FormParams{
			Item1: "source.locations.state[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceLocationsCity {
		formParams = append(formParams, FormParams{
			Item1: "source.locations.city[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceScopesCountry {
		formParams = append(formParams, FormParams{
			Item1: "source.scopes.country[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceScopesState {
		formParams = append(formParams, FormParams{
			Item1: "source.scopes.state[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceScopesCity {
		formParams = append(formParams, FormParams{
			Item1: "source.scopes.city[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceScopesLevel {
		formParams = append(formParams, FormParams{
			Item1: "source.scopes.level[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.Return {
		formParams = append(formParams, FormParams{
			Item1: "return[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if params.StoryId != 0 {
		formParams = append(formParams, FormParams{
			Item1: "story_id",
			Item2: a.Configuration.APIClient.ParameterToString(params.StoryId)})
	}

	if len(params.StoryUrl) > 0 {
		formParams = append(formParams, FormParams{
			Item1: "story_url",
			Item2: a.Configuration.APIClient.ParameterToString(params.StoryUrl)})
	}

	if len(params.StoryTitle) > 0 {
		formParams = append(formParams, FormParams{
			Item1: "story_title",
			Item2: a.Configuration.APIClient.ParameterToString(params.StoryTitle)})
	}

	if len(params.StoryBody) > 0 {
		formParams = append(formParams, FormParams{
			Item1: "story_body",
			Item2: a.Configuration.APIClient.ParameterToString(params.StoryBody)})
	}

	if len(params.BoostBy) > 0 {
		formParams = append(formParams, FormParams{
			Item1: "boost_by",
			Item2: a.Configuration.APIClient.ParameterToString(params.BoostBy)})
	}

	if len(params.StoryLanguage) > 0 {
		formParams = append(formParams, FormParams{
			Item1: "story_language",
			Item2: a.Configuration.APIClient.ParameterToString(params.StoryLanguage)})
	}

	if params.PerPage != 0 {
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
 * This endpoint is used for getting list of stories.
 *
 * @param params This is an StoriesParams struct which accepts following parameters:
 ** @param Id This parameter is used for finding stroies by story id.
 ** @param Title This parameter is used for finding stories whose title contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
 ** @param Body This parameter is used for finding stories whose body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
 ** @param Text This parameter is used for finding stories whose title or body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
 ** @param Language This parameter is used for finding stories whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes.
 ** @param PublishedAtStart This parameter is used for finding stories whose published at is greater than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates).
 ** @param PublishedAtEnd This parameter is used for finding stories whose published at is less than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates).
 ** @param CategoriesTaxonomy This parameter is used for defining type of the taxonomy for the rest of categories queries.
 ** @param CategoriesConfident This parameter is used for finding stories whose categories are confident.
 ** @param CategoriesId This parameter is used for finding stories whose categories id is the specified value.
 ** @param CategoriesLevel This parameter is used for finding stories whose categories level is the specified value.
 ** @param EntitiesTitleText This parameter is used for finding stories whose entities text in title is the specified value.
 ** @param EntitiesTitleType This parameter is used for finding stories whose entities type in title is the specified value.
 ** @param EntitiesTitleLinksDbpedia This parameter is used for finding stories whose entities dbpedia url in title is the specified value.
 ** @param EntitiesBodyText This parameter is used for finding stories whose entities text in body is the specified value.
 ** @param EntitiesBodyType This parameter is used for finding stories whose entities type in body is the specified value.
 ** @param EntitiesBodyLinksDbpedia This parameter is used for finding stories whose entities dbpedia url in body is the specified value.
 ** @param SentimentTitlePolarity This parameter is used for finding stories whose title sentiment is the specified value.
 ** @param SentimentBodyPolarity This parameter is used for finding stories whose body sentiment is the specified value.
 ** @param AuthorId This parameter is used for finding stories whose author id is the specified value.
 ** @param AuthorName This parameter is used for finding stories whose author full name contains the specified value.
 ** @param SourceId This parameter is used for finding stories whose source id is the specified value.
 ** @param SourceName This parameter is used for finding stories whose source name contains the specified value.
 ** @param SourceDomain This parameter is used for finding stories whose source domain is the specified value.
 ** @param SourceLocationsCountry This parameter is used for finding stories whose source country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes.
 ** @param SourceLocationsState This parameter is used for finding stories whose source state/province is the specified value.
 ** @param SourceLocationsCity This parameter is used for finding stories whose source city is the specified value.
 ** @param SourceScopesCountry This parameter is used for finding stories whose source scopes country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes.
 ** @param SourceScopesState This parameter is used for finding stories whose source scopes state/province is the specified value.
 ** @param SourceScopesCity This parameter is used for finding stories whose source scopes city is the specified value.
 ** @param SourceScopesLevel This parameter is used for finding stories whose source scopes level is the specified value.
 ** @param Cluster This parameter enables clustering for the returned stories.
 ** @param ClusterAlgorithm This parameter is used for specifying the clustering algorithm. It supprts STC, Lingo and [k-means](https://en.wikipedia.org/wiki/K-means_clustering) algorithms.
 ** @param Return This parameter is used for specifying return fields.
 ** @param SortBy This parameter is used for changing the order column of the result.
 ** @param SortDirection This parameter is used for changing the order direction of the result.
 ** @param Cursor This parameter is used for finding a specific page.
 ** @param PerPage This parameter is used for specifying number of the items in each page.
 * @return Stories
 */
func (a DefaultApi) ListStories(params *StoriesParams) (*Stories, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/stories"

	headerParams := make(map[string]string)
	queryParams := []QueryParams{}
	formParams := []FormParams{}
	// authentication (app_key) required

	// set key with prefix in header
	headerParams["X-AYLIEN-NewsAPI-Application-Key"] = a.Configuration.GetAPIKeyWithPrefix("X-AYLIEN-NewsAPI-Application-Key")
	// authentication (app_id) required

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

	if len(params.Title) > 0 {
		queryParams = append(queryParams, QueryParams{
			Item1: "title",
			Item2: a.Configuration.APIClient.ParameterToString(params.Title)})
	}

	if len(params.Body) > 0 {
		queryParams = append(queryParams, QueryParams{
			Item1: "body",
			Item2: a.Configuration.APIClient.ParameterToString(params.Body)})
	}

	if len(params.Text) > 0 {
		queryParams = append(queryParams, QueryParams{
			Item1: "text",
			Item2: a.Configuration.APIClient.ParameterToString(params.Text)})
	}

	for _, f := range params.Language {
		queryParams = append(queryParams, QueryParams{
			Item1: "language[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if len(params.PublishedAtStart) > 0 {
		queryParams = append(queryParams, QueryParams{
			Item1: "published_at.start",
			Item2: a.Configuration.APIClient.ParameterToString(params.PublishedAtStart)})
	}

	if len(params.PublishedAtEnd) > 0 {
		queryParams = append(queryParams, QueryParams{
			Item1: "published_at.end",
			Item2: a.Configuration.APIClient.ParameterToString(params.PublishedAtEnd)})
	}

	if len(params.CategoriesTaxonomy) > 0 {
		queryParams = append(queryParams, QueryParams{
			Item1: "categories.taxonomy",
			Item2: a.Configuration.APIClient.ParameterToString(params.CategoriesTaxonomy)})
	}

	if len(params.CategoriesConfident) > 0 {
		queryParams = append(queryParams, QueryParams{
			Item1: "categories.confident",
			Item2: a.Configuration.APIClient.ParameterToString(params.CategoriesConfident)})
	}

	for _, f := range params.CategoriesId {
		queryParams = append(queryParams, QueryParams{
			Item1: "categories.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.CategoriesLevel {
		queryParams = append(queryParams, QueryParams{
			Item1: "categories.level[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesTitleText {
		queryParams = append(queryParams, QueryParams{
			Item1: "entities.title.text[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesTitleType {
		queryParams = append(queryParams, QueryParams{
			Item1: "entities.title.type[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesTitleLinksDbpedia {
		queryParams = append(queryParams, QueryParams{
			Item1: "entities.title.links.dbpedia[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesBodyText {
		queryParams = append(queryParams, QueryParams{
			Item1: "entities.body.text[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesBodyType {
		queryParams = append(queryParams, QueryParams{
			Item1: "entities.body.type[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesBodyLinksDbpedia {
		queryParams = append(queryParams, QueryParams{
			Item1: "entities.body.links.dbpedia[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if len(params.SentimentTitlePolarity) > 0 {
		queryParams = append(queryParams, QueryParams{
			Item1: "sentiment.title.polarity",
			Item2: a.Configuration.APIClient.ParameterToString(params.SentimentTitlePolarity)})
	}

	if len(params.SentimentBodyPolarity) > 0 {
		queryParams = append(queryParams, QueryParams{
			Item1: "sentiment.body.polarity",
			Item2: a.Configuration.APIClient.ParameterToString(params.SentimentBodyPolarity)})
	}

	for _, f := range params.AuthorId {
		queryParams = append(queryParams, QueryParams{
			Item1: "author.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if len(params.AuthorName) > 0 {
		queryParams = append(queryParams, QueryParams{
			Item1: "author.name",
			Item2: a.Configuration.APIClient.ParameterToString(params.AuthorName)})
	}

	for _, f := range params.SourceId {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceName {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.name[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceDomain {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.domain[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceLocationsCountry {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.locations.country[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceLocationsState {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.locations.state[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceLocationsCity {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.locations.city[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceScopesCountry {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.scopes.country[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceScopesState {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.scopes.state[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceScopesCity {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.scopes.city[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceScopesLevel {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.scopes.level[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if len(params.Cluster) > 0 {
		queryParams = append(queryParams, QueryParams{
			Item1: "cluster",
			Item2: a.Configuration.APIClient.ParameterToString(params.Cluster)})
	}

	if len(params.ClusterAlgorithm) > 0 {
		queryParams = append(queryParams, QueryParams{
			Item1: "cluster.algorithm",
			Item2: a.Configuration.APIClient.ParameterToString(params.ClusterAlgorithm)})
	}

	for _, f := range params.Return {
		queryParams = append(queryParams, QueryParams{
			Item1: "return[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if len(params.SortBy) > 0 {
		queryParams = append(queryParams, QueryParams{
			Item1: "sort_by",
			Item2: a.Configuration.APIClient.ParameterToString(params.SortBy)})
	}

	if len(params.SortDirection) > 0 {
		queryParams = append(queryParams, QueryParams{
			Item1: "sort_direction",
			Item2: a.Configuration.APIClient.ParameterToString(params.SortDirection)})
	}

	if len(params.Cursor) > 0 {
		queryParams = append(queryParams, QueryParams{
			Item1: "cursor",
			Item2: a.Configuration.APIClient.ParameterToString(params.Cursor)})
	}

	if params.PerPage != 0 {
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
 ** @param Id This parameter is used for finding stroies by story id.
 ** @param Title This parameter is used for finding stories whose title contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
 ** @param Body This parameter is used for finding stories whose body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
 ** @param Text This parameter is used for finding stories whose title or body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
 ** @param Language This parameter is used for finding stories whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes.
 ** @param CategoriesTaxonomy This parameter is used for defining type of the taxonomy for the rest of categories queries.
 ** @param CategoriesConfident This parameter is used for finding stories whose categories are confident.
 ** @param CategoriesId This parameter is used for finding stories whose categories id is the specified value.
 ** @param CategoriesLevel This parameter is used for finding stories whose categories level is the specified value.
 ** @param EntitiesTitleText This parameter is used for finding stories whose entities text in title is the specified value.
 ** @param EntitiesTitleType This parameter is used for finding stories whose entities type in title is the specified value.
 ** @param EntitiesTitleLinksDbpedia This parameter is used for finding stories whose entities dbpedia url in title is the specified value.
 ** @param EntitiesBodyText This parameter is used for finding stories whose entities text in body is the specified value.
 ** @param EntitiesBodyType This parameter is used for finding stories whose entities type in body is the specified value.
 ** @param EntitiesBodyLinksDbpedia This parameter is used for finding stories whose entities dbpedia url in body is the specified value.
 ** @param SentimentTitlePolarity This parameter is used for finding stories whose title sentiment is the specified value.
 ** @param SentimentBodyPolarity This parameter is used for finding stories whose body sentiment is the specified value.
 ** @param AuthorId This parameter is used for finding stories whose author id is the specified value.
 ** @param AuthorName This parameter is used for finding stories whose author full name contains the specified value.
 ** @param SourceId This parameter is used for finding stories whose source id is the specified value.
 ** @param SourceName This parameter is used for finding stories whose source name contains the specified value.
 ** @param SourceDomain This parameter is used for finding stories whose source domain is the specified value.
 ** @param SourceLocationsCountry This parameter is used for finding stories whose source country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes.
 ** @param SourceLocationsState This parameter is used for finding stories whose source state/province is the specified value.
 ** @param SourceLocationsCity This parameter is used for finding stories whose source city is the specified value.
 ** @param SourceScopesCountry This parameter is used for finding stories whose source scopes country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes.
 ** @param SourceScopesState This parameter is used for finding stories whose source scopes state/province is the specified value.
 ** @param SourceScopesCity This parameter is used for finding stories whose source scopes city is the specified value.
 ** @param SourceScopesLevel This parameter is used for finding stories whose source scopes level is the specified value.
 ** @param PublishedAtStart This parameter is used for finding stories whose published at is less than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates).
 ** @param PublishedAtEnd This parameter is used for finding stories whose published at is greater than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates).
 ** @param Period The size of each date range expressed as an interval to be added to the lower bound. It supports Date Math Syntax. Valid options are &#x60;+&#x60; following an integer number greater than 0 and one of the Date Math keywords. e.g. &#x60;+1DAY&#x60;, &#x60;+2MINUTES&#x60; and &#x60;+1MONTH&#x60;. Here are [Supported keywords](https://newsapi.aylien.com/docs/working-with-dates#date-math).
 * @return TimeSeriesList
 */
func (a DefaultApi) ListTimeSeries(params *TimeSeriesParams) (*TimeSeriesList, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/time_series"

	headerParams := make(map[string]string)
	queryParams := []QueryParams{}
	formParams := []FormParams{}
	// authentication (app_key) required

	// set key with prefix in header
	headerParams["X-AYLIEN-NewsAPI-Application-Key"] = a.Configuration.GetAPIKeyWithPrefix("X-AYLIEN-NewsAPI-Application-Key")
	// authentication (app_id) required

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

	if len(params.Title) > 0 {
		queryParams = append(queryParams, QueryParams{
			Item1: "title",
			Item2: a.Configuration.APIClient.ParameterToString(params.Title)})
	}

	if len(params.Body) > 0 {
		queryParams = append(queryParams, QueryParams{
			Item1: "body",
			Item2: a.Configuration.APIClient.ParameterToString(params.Body)})
	}

	if len(params.Text) > 0 {
		queryParams = append(queryParams, QueryParams{
			Item1: "text",
			Item2: a.Configuration.APIClient.ParameterToString(params.Text)})
	}

	for _, f := range params.Language {
		queryParams = append(queryParams, QueryParams{
			Item1: "language[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if len(params.CategoriesTaxonomy) > 0 {
		queryParams = append(queryParams, QueryParams{
			Item1: "categories.taxonomy",
			Item2: a.Configuration.APIClient.ParameterToString(params.CategoriesTaxonomy)})
	}

	if len(params.CategoriesConfident) > 0 {
		queryParams = append(queryParams, QueryParams{
			Item1: "categories.confident",
			Item2: a.Configuration.APIClient.ParameterToString(params.CategoriesConfident)})
	}

	for _, f := range params.CategoriesId {
		queryParams = append(queryParams, QueryParams{
			Item1: "categories.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.CategoriesLevel {
		queryParams = append(queryParams, QueryParams{
			Item1: "categories.level[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesTitleText {
		queryParams = append(queryParams, QueryParams{
			Item1: "entities.title.text[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesTitleType {
		queryParams = append(queryParams, QueryParams{
			Item1: "entities.title.type[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesTitleLinksDbpedia {
		queryParams = append(queryParams, QueryParams{
			Item1: "entities.title.links.dbpedia[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesBodyText {
		queryParams = append(queryParams, QueryParams{
			Item1: "entities.body.text[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesBodyType {
		queryParams = append(queryParams, QueryParams{
			Item1: "entities.body.type[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesBodyLinksDbpedia {
		queryParams = append(queryParams, QueryParams{
			Item1: "entities.body.links.dbpedia[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if len(params.SentimentTitlePolarity) > 0 {
		queryParams = append(queryParams, QueryParams{
			Item1: "sentiment.title.polarity",
			Item2: a.Configuration.APIClient.ParameterToString(params.SentimentTitlePolarity)})
	}

	if len(params.SentimentBodyPolarity) > 0 {
		queryParams = append(queryParams, QueryParams{
			Item1: "sentiment.body.polarity",
			Item2: a.Configuration.APIClient.ParameterToString(params.SentimentBodyPolarity)})
	}

	for _, f := range params.AuthorId {
		queryParams = append(queryParams, QueryParams{
			Item1: "author.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if len(params.AuthorName) > 0 {
		queryParams = append(queryParams, QueryParams{
			Item1: "author.name",
			Item2: a.Configuration.APIClient.ParameterToString(params.AuthorName)})
	}

	for _, f := range params.SourceId {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceName {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.name[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceDomain {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.domain[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceLocationsCountry {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.locations.country[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceLocationsState {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.locations.state[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceLocationsCity {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.locations.city[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceScopesCountry {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.scopes.country[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceScopesState {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.scopes.state[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceScopesCity {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.scopes.city[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceScopesLevel {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.scopes.level[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if len(params.PublishedAtStart) > 0 {
		queryParams = append(queryParams, QueryParams{
			Item1: "published_at.start",
			Item2: a.Configuration.APIClient.ParameterToString(params.PublishedAtStart)})
	}

	if len(params.PublishedAtEnd) > 0 {
		queryParams = append(queryParams, QueryParams{
			Item1: "published_at.end",
			Item2: a.Configuration.APIClient.ParameterToString(params.PublishedAtEnd)})
	}

	if len(params.Period) > 0 {
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
 * This endpoint is used for finding news trendings based on stories resource.
 *
 * @param params This is an TrendsParams struct which accepts following parameters:
 ** @param Id This parameter is used for finding stroies by story id.
 ** @param Title This parameter is used for finding stories whose title contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
 ** @param Body This parameter is used for finding stories whose body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
 ** @param Text This parameter is used for finding stories whose title or body contains a specfic keyword. It supports [boolean operators](https://newsapi.aylien.com/docs/boolean-operators).
 ** @param Language This parameter is used for finding stories whose language is the specified value. It supports [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language codes.
 ** @param PublishedAtStart This parameter is used for finding stories whose published at is greater than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates).
 ** @param PublishedAtEnd This parameter is used for finding stories whose published at is less than the specified value. [Here](https://newsapi.aylien.com/docs/working-with-dates) you can find more information about how [to work with dates](https://newsapi.aylien.com/docs/working-with-dates).
 ** @param CategoriesTaxonomy This parameter is used for defining type of the taxonomy for the rest of categories queries.
 ** @param CategoriesConfident This parameter is used for finding stories whose categories are confident.
 ** @param CategoriesId This parameter is used for finding stories whose categories id is the specified value.
 ** @param CategoriesLevel This parameter is used for finding stories whose categories level is the specified value.
 ** @param EntitiesTitleText This parameter is used for finding stories whose entities text in title is the specified value.
 ** @param EntitiesTitleType This parameter is used for finding stories whose entities type in title is the specified value.
 ** @param EntitiesTitleLinksDbpedia This parameter is used for finding stories whose entities dbpedia url in title is the specified value.
 ** @param EntitiesBodyText This parameter is used for finding stories whose entities text in body is the specified value.
 ** @param EntitiesBodyType This parameter is used for finding stories whose entities type in body is the specified value.
 ** @param EntitiesBodyLinksDbpedia This parameter is used for finding stories whose entities dbpedia url in body is the specified value.
 ** @param SentimentTitlePolarity This parameter is used for finding stories whose title sentiment is the specified value.
 ** @param SentimentBodyPolarity This parameter is used for finding stories whose body sentiment is the specified value.
 ** @param AuthorId This parameter is used for finding stories whose author id is the specified value.
 ** @param AuthorName This parameter is used for finding stories whose author full name contains the specified value.
 ** @param SourceId This parameter is used for finding stories whose source id is the specified value.
 ** @param SourceName This parameter is used for finding stories whose source name contains the specified value.
 ** @param SourceDomain This parameter is used for finding stories whose source domain is the specified value.
 ** @param SourceLocationsCountry This parameter is used for finding stories whose source country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes.
 ** @param SourceLocationsState This parameter is used for finding stories whose source state/province is the specified value.
 ** @param SourceLocationsCity This parameter is used for finding stories whose source city is the specified value.
 ** @param SourceScopesCountry This parameter is used for finding stories whose source scopes country is the specified value. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes.
 ** @param SourceScopesState This parameter is used for finding stories whose source scopes state/province is the specified value.
 ** @param SourceScopesCity This parameter is used for finding stories whose source scopes city is the specified value.
 ** @param SourceScopesLevel This parameter is used for finding stories whose source scopes level is the specified value.
 ** @param Field This parameter is used to specify the trend field.
 * @return Trends
 */
func (a DefaultApi) ListTrends(params *TrendsParams) (*Trends, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/trends"

	headerParams := make(map[string]string)
	queryParams := []QueryParams{}
	formParams := []FormParams{}
	// authentication (app_key) required

	// set key with prefix in header
	headerParams["X-AYLIEN-NewsAPI-Application-Key"] = a.Configuration.GetAPIKeyWithPrefix("X-AYLIEN-NewsAPI-Application-Key")
	// authentication (app_id) required

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

	if len(params.Title) > 0 {
		queryParams = append(queryParams, QueryParams{
			Item1: "title",
			Item2: a.Configuration.APIClient.ParameterToString(params.Title)})
	}

	if len(params.Body) > 0 {
		queryParams = append(queryParams, QueryParams{
			Item1: "body",
			Item2: a.Configuration.APIClient.ParameterToString(params.Body)})
	}

	if len(params.Text) > 0 {
		queryParams = append(queryParams, QueryParams{
			Item1: "text",
			Item2: a.Configuration.APIClient.ParameterToString(params.Text)})
	}

	for _, f := range params.Language {
		queryParams = append(queryParams, QueryParams{
			Item1: "language[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if len(params.PublishedAtStart) > 0 {
		queryParams = append(queryParams, QueryParams{
			Item1: "published_at.start",
			Item2: a.Configuration.APIClient.ParameterToString(params.PublishedAtStart)})
	}

	if len(params.PublishedAtEnd) > 0 {
		queryParams = append(queryParams, QueryParams{
			Item1: "published_at.end",
			Item2: a.Configuration.APIClient.ParameterToString(params.PublishedAtEnd)})
	}

	if len(params.CategoriesTaxonomy) > 0 {
		queryParams = append(queryParams, QueryParams{
			Item1: "categories.taxonomy",
			Item2: a.Configuration.APIClient.ParameterToString(params.CategoriesTaxonomy)})
	}

	if len(params.CategoriesConfident) > 0 {
		queryParams = append(queryParams, QueryParams{
			Item1: "categories.confident",
			Item2: a.Configuration.APIClient.ParameterToString(params.CategoriesConfident)})
	}

	for _, f := range params.CategoriesId {
		queryParams = append(queryParams, QueryParams{
			Item1: "categories.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.CategoriesLevel {
		queryParams = append(queryParams, QueryParams{
			Item1: "categories.level[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesTitleText {
		queryParams = append(queryParams, QueryParams{
			Item1: "entities.title.text[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesTitleType {
		queryParams = append(queryParams, QueryParams{
			Item1: "entities.title.type[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesTitleLinksDbpedia {
		queryParams = append(queryParams, QueryParams{
			Item1: "entities.title.links.dbpedia[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesBodyText {
		queryParams = append(queryParams, QueryParams{
			Item1: "entities.body.text[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesBodyType {
		queryParams = append(queryParams, QueryParams{
			Item1: "entities.body.type[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.EntitiesBodyLinksDbpedia {
		queryParams = append(queryParams, QueryParams{
			Item1: "entities.body.links.dbpedia[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if len(params.SentimentTitlePolarity) > 0 {
		queryParams = append(queryParams, QueryParams{
			Item1: "sentiment.title.polarity",
			Item2: a.Configuration.APIClient.ParameterToString(params.SentimentTitlePolarity)})
	}

	if len(params.SentimentBodyPolarity) > 0 {
		queryParams = append(queryParams, QueryParams{
			Item1: "sentiment.body.polarity",
			Item2: a.Configuration.APIClient.ParameterToString(params.SentimentBodyPolarity)})
	}

	for _, f := range params.AuthorId {
		queryParams = append(queryParams, QueryParams{
			Item1: "author.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if len(params.AuthorName) > 0 {
		queryParams = append(queryParams, QueryParams{
			Item1: "author.name",
			Item2: a.Configuration.APIClient.ParameterToString(params.AuthorName)})
	}

	for _, f := range params.SourceId {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.id[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceName {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.name[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceDomain {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.domain[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceLocationsCountry {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.locations.country[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceLocationsState {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.locations.state[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceLocationsCity {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.locations.city[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceScopesCountry {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.scopes.country[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceScopesState {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.scopes.state[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceScopesCity {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.scopes.city[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	for _, f := range params.SourceScopesLevel {
		queryParams = append(queryParams, QueryParams{
			Item1: "source.scopes.level[]",
			Item2: a.Configuration.APIClient.ParameterToString(f)})
	}

	if len(params.Field) > 0 {
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
