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
	"time"
)

type Story struct {

	// ID of the story which is unique identification
	Id int64 `json:"id,omitempty"`

	// Title of the story
	Title string `json:"title,omitempty"`

	// Body of the story
	Body string `json:"body,omitempty"`

	// The suggested story summary
	Summary Summary `json:"summary,omitempty"`

	// The story source
	Source Source `json:"source,omitempty"`

	// The story author
	Author Author `json:"author,omitempty"`

	// Extracted entities from the story title or body
	Entities Entities `json:"entities,omitempty"`

	// Extracted keywords mentioned in the story title or body
	Keywords []string `json:"keywords,omitempty"`

	// An array of suggested Story hashtags
	Hashtags []string `json:"hashtags,omitempty"`

	// Characters count of the story body
	CharactersCount int32 `json:"characters_count,omitempty"`

	// Words count of the story body
	WordsCount int32 `json:"words_count,omitempty"`

	// Sentences count of the story body
	SentencesCount int32 `json:"sentences_count,omitempty"`

	// Paragraphs count of the story body
	ParagraphsCount int32 `json:"paragraphs_count,omitempty"`

	// Suggested categories for the story
	Categories []Category `json:"categories,omitempty"`

	// Social shares count for the story
	SocialSharesCount ShareCounts `json:"social_shares_count,omitempty"`

	// An array of extracted media such as images and videos
	Media []Media `json:"media,omitempty"`

	// Suggested sentiments for the story title or body
	Sentiment Sentiments `json:"sentiment,omitempty"`

	// Language of the story
	Language string `json:"language,omitempty"`

	// Published date of the story
	PublishedAt time.Time `json:"published_at,omitempty"`

	// Links which is related to the story
	Links StoryLinks `json:"links,omitempty"`
}
