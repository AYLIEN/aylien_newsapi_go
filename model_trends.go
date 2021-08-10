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
import (
	"time"
)
// Trends struct for Trends
type Trends struct {
	// The field of trends
	Field string `json:"field,omitempty"`
	// The end of a period in which searched stories were published
	PublishedAtEnd time.Time `json:"published_at.end,omitempty"`
	// The start of a period in which searched stories were published
	PublishedAtStart time.Time `json:"published_at.start,omitempty"`
	// An array of trends
	Trends []Trend `json:"trends,omitempty"`
}
