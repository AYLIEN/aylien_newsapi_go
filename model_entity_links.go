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
// EntityLinks struct for EntityLinks
type EntityLinks struct {
	// A dbpedia resource URL (deprecated)
	Dbpedia string `json:"dbpedia,omitempty"`
	// A wikidata resource URL
	Wikidata string `json:"wikidata,omitempty"`
	// A wikipedia resource URL
	Wikipedia string `json:"wikipedia,omitempty"`
}
