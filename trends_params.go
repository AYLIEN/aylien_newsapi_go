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

type TrendsParams struct {
	Id                             []int64
	NotId                          []int64
	Title                          string
	Body                           string
	Text                           string
	Language                       []string
	NotLanguage                    []string
	PublishedAtStart               string
	PublishedAtEnd                 string
	CategoriesTaxonomy             string
	CategoriesConfident            string
	CategoriesId                   []string
	NotCategoriesId                []string
	CategoriesLevel                []int32
	NotCategoriesLevel             []int32
	EntitiesTitleText              []string
	NotEntitiesTitleText           []string
	EntitiesTitleType              []string
	NotEntitiesTitleType           []string
	EntitiesTitleLinksDbpedia      []string
	NotEntitiesTitleLinksDbpedia   []string
	EntitiesBodyText               []string
	NotEntitiesBodyText            []string
	EntitiesBodyType               []string
	NotEntitiesBodyType            []string
	EntitiesBodyLinksDbpedia       []string
	NotEntitiesBodyLinksDbpedia    []string
	SentimentTitlePolarity         string
	NotSentimentTitlePolarity      string
	SentimentBodyPolarity          string
	NotSentimentBodyPolarity       string
	MediaImagesCountMin            string
	MediaImagesCountMax            string
	MediaImagesWidthMin            string
	MediaImagesWidthMax            string
	MediaImagesHeightMin           string
	MediaImagesHeightMax           string
	MediaImagesContentLengthMin    string
	MediaImagesContentLengthMax    string
	MediaImagesFormat              []string
	NotMediaImagesFormat           []string
	MediaVideosCountMin            string
	MediaVideosCountMax            string
	AuthorId                       []int32
	NotAuthorId                    []int32
	AuthorName                     string
	NotAuthorName                  string
	SourceId                       []int32
	NotSourceId                    []int32
	SourceName                     []string
	NotSourceName                  []string
	SourceDomain                   []string
	NotSourceDomain                []string
	SourceLocationsCountry         []string
	NotSourceLocationsCountry      []string
	SourceLocationsState           []string
	NotSourceLocationsState        []string
	SourceLocationsCity            []string
	NotSourceLocationsCity         []string
	SourceScopesCountry            []string
	NotSourceScopesCountry         []string
	SourceScopesState              []string
	NotSourceScopesState           []string
	SourceScopesCity               []string
	NotSourceScopesCity            []string
	SourceScopesLevel              []string
	NotSourceScopesLevel           []string
	SourceLinksInCountMin          int32
	SourceLinksInCountMax          int32
	SourceRankingsAlexaRankMin     int32
	SourceRankingsAlexaRankMax     int32
	SourceRankingsAlexaCountry     []string
	SocialSharesCountFacebookMin   string
	SocialSharesCountFacebookMax   string
	SocialSharesCountGooglePlusMin string
	SocialSharesCountGooglePlusMax string
	SocialSharesCountLinkedinMin   string
	SocialSharesCountLinkedinMax   string
	SocialSharesCountRedditMin     string
	SocialSharesCountRedditMax     string
	Field                          string
}
