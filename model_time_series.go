/*
 * AYLIEN News API
 *
 * The AYLIEN News API is the most powerful way of sourcing, searching and syndicating analyzed and enriched news content. It is accessed by sending HTTP requests to our server, which returns information to your client. 
 *
 * API version: 4.7.4
 * Contact: support@aylien.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package newsapi
import (
	"time"
)
// TimeSeries struct for TimeSeries
type TimeSeries struct {
	// The count of time series bin
	Count int32 `json:"count,omitempty"`
	// The published date of the time series bin
	PublishedAt time.Time `json:"published_at,omitempty"`
	Sentiment AggregatedSentiment `json:"sentiment,omitempty"`
}
