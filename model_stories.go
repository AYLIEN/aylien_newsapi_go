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
// Stories Stories containing new V3 entities - available for new_v3_entities feature flag
type Stories struct {
	// The next page cursor
	NextPageCursor string `json:"next_page_cursor,omitempty"`
	// An array of stories
	Stories []Story `json:"stories,omitempty"`
	// The end of a period in which searched stories were published
	PublishedAtEnd time.Time `json:"published_at.end,omitempty"`
	// The start of a period in which searched stories were published
	PublishedAtStart time.Time `json:"published_at.start,omitempty"`
	// Notifies about possible issues that occurred when searching for stories
	Warnings []Warning `json:"warnings,omitempty"`
}
