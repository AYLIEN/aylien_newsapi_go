/*
 * AYLIEN News API
 *
 * The AYLIEN News API is the most powerful way of sourcing, searching and syndicating analyzed and enriched news content. It is accessed by sending HTTP requests to our server, which returns information to your client. 
 *
 * API version: 3.0
 * Contact: support@aylien.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package newsapi
// Scope struct for Scope
type Scope struct {
	// The scope by city
	City string `json:"city,omitempty"`
	// The source scope by country code. It supports [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. 
	Country string `json:"country,omitempty"`
	Level ScopeLevel `json:"level,omitempty"`
	// The scope by state
	State string `json:"state,omitempty"`
}
