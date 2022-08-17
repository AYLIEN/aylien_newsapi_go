/*
 * AYLIEN News API
 *
 * The AYLIEN News API is the most powerful way of sourcing, searching and syndicating analyzed and enriched news content. It is accessed by sending HTTP requests to our server, which returns information to your client. 
 *
 * API version: 4.7.2
 * Contact: support@aylien.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package newsapi
// Sentiments struct for Sentiments
type Sentiments struct {
	Body Sentiment `json:"body,omitempty"`
	Title Sentiment `json:"title,omitempty"`
}
