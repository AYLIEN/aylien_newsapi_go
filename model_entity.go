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
// Entity struct for Entity
type Entity struct {
	// The indices of the entity text
	Indices [][]int32 `json:"indices,omitempty"`
	Links EntityLinks `json:"links,omitempty"`
	// The entity score
	Score float64 `json:"score,omitempty"`
	// The entity text
	Text string `json:"text,omitempty"`
	// An array of the dbpedia types
	Types []string `json:"types,omitempty"`
}
