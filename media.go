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

type Media struct {

	// The type of media
	Type_ string `json:"type,omitempty"`

	// A URL which points to the media file
	Url string `json:"url,omitempty"`

	// The format of media
	Format string `json:"format,omitempty"`

	// The content length of media
	ContentLength int32 `json:"content_length,omitempty"`

	// The width of media
	Width int32 `json:"width,omitempty"`

	// The height of media
	Height int32 `json:"height,omitempty"`
}
