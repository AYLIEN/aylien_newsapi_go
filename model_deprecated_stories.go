/*
 * AYLIEN News API
 *
 * The AYLIEN News API is the most powerful way of sourcing, searching and syndicating analyzed and enriched news content. It is accessed by sending HTTP requests to our server, which returns information to your client. 
 *
 * API version: 3.1
 * Contact: support@aylien.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package newsapi
import (
	"time"
)
// DeprecatedStories Stories containing deprecated entities
type DeprecatedStories struct {
	// The next page cursor
	NextPageCursor string `json:"next_page_cursor,omitempty"`
	// The end of a period in which searched stories were published
	PublishedAtEnd time.Time `json:"published_at.end,omitempty"`
	// The start of a period in which searched stories were published
	PublishedAtStart time.Time `json:"published_at.start,omitempty"`
	// An array of stories
	Stories []DeprecatedStory `json:"stories,omitempty"`
	// Notifies about possible issues that occurred when searching for stories
	Warnings []Warning `json:"warnings,omitempty"`
}
