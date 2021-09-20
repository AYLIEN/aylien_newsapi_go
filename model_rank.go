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
import (
	"time"
)
// Rank struct for Rank
type Rank struct {
	// The country code which the rank is in it
	Country string `json:"country,omitempty"`
	// The fetched date of the rank
	FetchedAt time.Time `json:"fetched_at,omitempty"`
	// The rank
	Rank int32 `json:"rank,omitempty"`
}
