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

type Coverages struct {

	// The input story title
	StoryTitle string `json:"story_title,omitempty"`

	// The input story body
	StoryBody string `json:"story_body,omitempty"`

	// The input story published date
	StoryPublishedAt time.Time `json:"story_published_at,omitempty"`

	// The input story language
	StoryLanguage string `json:"story_language,omitempty"`

	// An array of coverages for the input story
	Coverages []Story `json:"coverages,omitempty"`

	// An array of clusters
	Clusters []StoryCluster `json:"clusters,omitempty"`
}
