/*
 * AYLIEN News API
 *
 * The AYLIEN News API is the most powerful way of sourcing, searching and syndicating analyzed and enriched news content. It is accessed by sending HTTP requests to our server, which returns information to your client. 
 *
 * API version: 4.7.4
 * Contact: support@aylien.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package newsapi
// EntityMention struct for EntityMention
type EntityMention struct {
	Index EntityMentionIndex `json:"index,omitempty"`
	Sentiment EntitySentiment `json:"sentiment,omitempty"`
}
