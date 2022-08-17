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
// StoryTranslation struct for StoryTranslation
type StoryTranslation struct {
	// Translation of body
	Body string `json:"body,omitempty"`
	// Translation of title
	Title string `json:"title,omitempty"`
}
