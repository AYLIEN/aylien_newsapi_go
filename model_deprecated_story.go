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
import (
	"time"
)
// DeprecatedStory struct for DeprecatedStory
type DeprecatedStory struct {
	Author Author `json:"author,omitempty"`
	// Body of the story
	Body string `json:"body,omitempty"`
	// Suggested categories for the story
	Categories []Category `json:"categories,omitempty"`
	// Character count of the story body
	CharactersCount int32 `json:"characters_count,omitempty"`
	// An array of clusters the story is associated with
	Clusters []int64 `json:"clusters,omitempty"`
	Entities DeprecatedEntities `json:"entities,omitempty"`
	// An array of suggested Story hashtags
	Hashtags []string `json:"hashtags,omitempty"`
	// ID of the story which is a unique identification
	Id int64 `json:"id,omitempty"`
	// Extracted keywords mentioned in the story title or body
	Keywords []string `json:"keywords,omitempty"`
	// Language of the story
	Language string `json:"language,omitempty"`
	// License type of the story
	LicenseType int32 `json:"license_type,omitempty"`
	Links StoryLinks `json:"links,omitempty"`
	// An array of extracted media such as images and videos
	Media []Media `json:"media,omitempty"`
	// Paragraph count of the story body
	ParagraphsCount int32 `json:"paragraphs_count,omitempty"`
	// Published date of the story
	PublishedAt time.Time `json:"published_at,omitempty"`
	// Sentence count of the story body
	SentencesCount int32 `json:"sentences_count,omitempty"`
	Sentiment Sentiments `json:"sentiment,omitempty"`
	SocialSharesCount ShareCounts `json:"social_shares_count,omitempty"`
	Source Source `json:"source,omitempty"`
	Summary Summary `json:"summary,omitempty"`
	// Title of the story
	Title string `json:"title,omitempty"`
	Translations StoryTranslations `json:"translations,omitempty"`
	// Word count of the story body
	WordsCount int32 `json:"words_count,omitempty"`
}
