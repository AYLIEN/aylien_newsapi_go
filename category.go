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

type Category struct {

	// The ID of the category
	Id string `json:"id,omitempty"`

	// The taxonomy of the category
	Taxonomy string `json:"taxonomy,omitempty"`

	// The level of the category
	Level int32 `json:"level,omitempty"`

	// The score of the category
	Score float64 `json:"score,omitempty"`

	// It defines whether the extracted category is confident or not
	Confident bool `json:"confident,omitempty"`

	// Related links for the category
	Links CategoryLinks `json:"links,omitempty"`
}
