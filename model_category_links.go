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
// CategoryLinks struct for CategoryLinks
type CategoryLinks struct {
	// A URL pointing to the parent category
	Parent string `json:"parent,omitempty"`
	// A URL pointing to the category
	Self string `json:"self,omitempty"`
}
