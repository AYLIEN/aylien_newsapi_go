/*
 * AYLIEN News API
 *
 * The AYLIEN News API is the most powerful way of sourcing, searching and syndicating analyzed and enriched news content. It is accessed by sending HTTP requests to our server, which returns information to your client. 
 *
 * API version: 5.1.0
 * Contact: support@aylien.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package newsapi
// Autocomplete struct for Autocomplete
type Autocomplete struct {
	// ID of the autocomplete
	Id string `json:"id,omitempty"`
	// Text of the autocomplete
	Text string `json:"text,omitempty"`
}
