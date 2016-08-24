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

type HistogramsParams struct {
	Id                         []int64
	Title                      string
	Body                       string
	Text                       string
	Language                   []string
	PublishedAtStart           string
	PublishedAtEnd             string
	CategoriesTaxonomy         string
	CategoriesConfident        string
	CategoriesId               []string
	CategoriesLevel            []int32
	EntitiesTitleText          []string
	EntitiesTitleType          []string
	EntitiesTitleLinksDbpedia  []string
	EntitiesBodyText           []string
	EntitiesBodyType           []string
	EntitiesBodyLinksDbpedia   []string
	SentimentTitlePolarity     string
	SentimentBodyPolarity      string
	MediaImagesCountMin        string
	MediaImagesCountMax        string
	MediaVideosCountMin        string
	MediaVideosCountMax        string
	AuthorId                   []int32
	AuthorName                 string
	SourceId                   []int32
	SourceName                 []string
	SourceDomain               []string
	SourceLocationsCountry     []string
	SourceLocationsState       []string
	SourceLocationsCity        []string
	SourceScopesCountry        []string
	SourceScopesState          []string
	SourceScopesCity           []string
	SourceScopesLevel          []string
	SourceLinksInCountMin      int32
	SourceLinksInCountMax      int32
	SourceRankingsAlexaRankMin int32
	SourceRankingsAlexaRankMax int32
	SourceRankingsAlexaCountry []string
	IntervalStart              int32
	IntervalEnd                int32
	IntervalWidth              int32
	Field                      string
}
