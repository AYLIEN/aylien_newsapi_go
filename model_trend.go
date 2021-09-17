/*
 * AYLIEN News API
 *
 * The AYLIEN News API is the most powerful way of sourcing, searching and syndicating analyzed and enriched news content. It is accessed by sending HTTP requests to our server, which returns information to your client. 
 *
 * API version: 5.1.0
 * Contact: support@aylien.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package newsapi
// Trend struct for Trend
type Trend struct {
	// The count of the trend
	Count int32 `json:"count,omitempty"`
	Sentiment AggregatedSentiment `json:"sentiment,omitempty"`
	// The value of the trend
	Value string `json:"value,omitempty"`
}
