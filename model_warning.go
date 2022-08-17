/*
 * AYLIEN News API
 *
 * The AYLIEN News API is the most powerful way of sourcing, searching and syndicating analyzed and enriched news content. It is accessed by sending HTTP requests to our server, which returns information to your client. 
 *
 * API version: 5.2.3
 * Contact: support@aylien.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package newsapi
// Warning struct for Warning
type Warning struct {
	// The detailed description of the warning.
	Detail string `json:"detail,omitempty"`
	// The identfier of the warning, represents its origin.
	Id string `json:"id,omitempty"`
	Links ErrorLinks `json:"links,omitempty"`
}
