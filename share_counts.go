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

type ShareCounts struct {

	// Facebook shares count
	Facebook []ShareCount `json:"facebook,omitempty"`

	// Google Plus shares count
	GooglePlus []ShareCount `json:"google_plus,omitempty"`

	// LinkedIn shares count
	Linkedin []ShareCount `json:"linkedin,omitempty"`

	// Reddit shares count
	Reddit []ShareCount `json:"reddit,omitempty"`
}
