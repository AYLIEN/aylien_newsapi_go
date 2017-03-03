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

type Stories struct {

	// An array of stories
	Stories []Story `json:"stories,omitempty"`

	// An array of clusters
	Clusters []StoryCluster `json:"clusters,omitempty"`

	// The next page cursor
	NextPageCursor string `json:"next_page_cursor,omitempty"`
}
