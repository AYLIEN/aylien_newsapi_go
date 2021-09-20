/*
 * AYLIEN News API
 *
 * The AYLIEN News API is the most powerful way of sourcing, searching and syndicating analyzed and enriched news content. It is accessed by sending HTTP requests to our server, which returns information to your client. 
 *
 * API version: 4.7
 * Contact: support@aylien.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package newsapi
// EntitySentiment struct for EntitySentiment
type EntitySentiment struct {
	// Polarity confidence of the sentiment
	Confidence float64 `json:"confidence,omitempty"`
	Polarity SentimentPolarity `json:"polarity,omitempty"`
}
