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
// Author struct for Author
type Author struct {
	// A URL which points to the author avatar
	AvatarUrl string `json:"avatar_url,omitempty"`
	// A unique identification for the author
	Id int64 `json:"id,omitempty"`
	// The extracted author full name
	Name string `json:"name,omitempty"`
}
