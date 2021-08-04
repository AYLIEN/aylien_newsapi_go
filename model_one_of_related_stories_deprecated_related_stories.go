/*
 * AYLIEN News API
 *
 * The AYLIEN News API is the most powerful way of sourcing, searching and syndicating analyzed and enriched news content. It is accessed by sending HTTP requests to our server, which returns information to your client. 
 *
 * API version: 4.6
 * Contact: support@aylien.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package newsapi
import (
	"time"
)
// OneOfRelatedStoriesDeprecatedRelatedStories Wrapper of RelatedStories or DeprecatedRelatedStories
type OneOfRelatedStoriesDeprecatedRelatedStories struct {
	// The end of a period in which searched stories were published
	PublishedAtEnd time.Time `json:"published_at.end,omitempty"`
	// The start of a period in which searched stories were published
	PublishedAtStart time.Time `json:"published_at.start,omitempty"`
	// An array of related stories for the input story
	RelatedStories []DeprecatedStory `json:"related_stories,omitempty"`
	// The input story body
	StoryBody string `json:"story_body,omitempty"`
	// The input story language
	StoryLanguage string `json:"story_language,omitempty"`
	// The input story title
	StoryTitle string `json:"story_title,omitempty"`
}
