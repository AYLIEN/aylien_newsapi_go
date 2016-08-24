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

type Source struct {

	// The source id which is a unique value
	Id int32 `json:"id,omitempty"`

	// The source name
	Name string `json:"name,omitempty"`

	// The title of the home page URL
	Title string `json:"title,omitempty"`

	// A general explanation about the source
	Description string `json:"description,omitempty"`

	// The number of websites that link to the source
	LinksInCount int32 `json:"links_in_count,omitempty"`

	// The home page URL of the source
	HomePageUrl string `json:"home_page_url,omitempty"`

	// The domain name of the source which is extracted from the source URL
	Domain string `json:"domain,omitempty"`

	// A URL which points to the source logo
	LogoUrl string `json:"logo_url,omitempty"`

	// The source locations which are tend to be the physical locations of the source, e.g. BBC headquarter is located in London.
	Locations []Location `json:"locations,omitempty"`

	// The source scopes which is tend to be scope locations of the source, e.g. BBC scopes is international.
	Scopes []Scope `json:"scopes,omitempty"`

	// The web rankings of the source
	Rankings Rankings `json:"rankings,omitempty"`
}
