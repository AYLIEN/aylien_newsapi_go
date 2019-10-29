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
// Category struct for Category
type Category struct {
	// It defines whether the extracted category is confident or not
	Confident bool `json:"confident,omitempty"`
	// The ID of the category
	Id string `json:"id,omitempty"`
	// The level of the category
	Level int32 `json:"level,omitempty"`
	Links CategoryLinks `json:"links,omitempty"`
	// The score of the category
	Score float64 `json:"score,omitempty"`
	// The taxonomy of the category
	Taxonomy string `json:"taxonomy,omitempty"`
}
