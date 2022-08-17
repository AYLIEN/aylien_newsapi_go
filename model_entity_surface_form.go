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
// EntitySurfaceForm struct for EntitySurfaceForm
type EntitySurfaceForm struct {
	// Amount of entity surface form mentions in the article
	Frequency int32 `json:"frequency,omitempty"`
	// Mentions of the entity text
	Mentions []EntityMention `json:"mentions,omitempty"`
	// The entity text
	Text string `json:"text,omitempty"`
}
