/*
 * AYLIEN News API
 *
 * The AYLIEN News API is the most powerful way of sourcing, searching and syndicating analyzed and enriched news content. It is accessed by sending HTTP requests to our server, which returns information to your client. 
 *
 * API version: 4.7.3
 * Contact: support@aylien.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package newsapi
// AggregatedSentiment The aggregation of sentiments
type AggregatedSentiment struct {
	// Negative sentiments count
	Negative int32 `json:"negative,omitempty"`
	// Neutral sentiments count
	Neutral int32 `json:"neutral,omitempty"`
	// Positive sentiments count
	Positive int32 `json:"positive,omitempty"`
}
