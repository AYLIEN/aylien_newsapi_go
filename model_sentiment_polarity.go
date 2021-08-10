/*
 * AYLIEN News API
 *
 * The AYLIEN News API is the most powerful way of sourcing, searching and syndicating analyzed and enriched news content. It is accessed by sending HTTP requests to our server, which returns information to your client. 
 *
 * API version: 5.0.1
 * Contact: support@aylien.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package newsapi
// SentimentPolarity Polarity of the sentiment
type SentimentPolarity string

// List of SentimentPolarity
const (
	POSITIVE SentimentPolarity = "positive"
	NEUTRAL SentimentPolarity = "neutral"
	NEGATIVE SentimentPolarity = "negative"
)
