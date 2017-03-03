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

import (
	"time"
)

type TimeSeriesList struct {

	// An array of time series
	TimeSeries []TimeSeries `json:"time_series,omitempty"`

	// The size of each date range expressed as an interval to be added to the lower bound.
	Period string `json:"period,omitempty"`

	// The start published date of the time series
	PublishedAtStart time.Time `json:"published_at.start,omitempty"`

	// The end published date of the time series
	PublishedAtEnd time.Time `json:"published_at.end,omitempty"`
}
