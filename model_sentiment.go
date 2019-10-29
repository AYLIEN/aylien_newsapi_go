/*
 * AYLIEN News API
 *
 * The AYLIEN News API is the most powerful way of sourcing, searching and syndicating analyzed and enriched news content. It is accessed by sending HTTP requests to our server, which returns information to your client. 
 *
 * API version: 2.0
 * Contact: support@aylien.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package newsapi
// Sentiment struct for Sentiment
type Sentiment struct {
	// Polarity of the sentiment
	Polarity string `json:"polarity,omitempty"`
	// Polarity score of the sentiment
	Score float64 `json:"score,omitempty"`
}