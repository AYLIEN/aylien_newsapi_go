/*
 * AYLIEN News API
 *
 * The AYLIEN News API is the most powerful way of sourcing, searching and syndicating analyzed and enriched news content. It is accessed by sending HTTP requests to our server, which returns information to your client. 
 *
 * API version: 5.1.1
 * Contact: support@aylien.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package newsapi
import (
	"time"
)
// ShareCount struct for ShareCount
type ShareCount struct {
	// The number of shares
	Count int32 `json:"count,omitempty"`
	// The fetched date of the shares
	FetchedAt time.Time `json:"fetched_at,omitempty"`
}
