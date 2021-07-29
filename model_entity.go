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
// Entity struct for Entity
type Entity struct {
	Body EntityInText `json:"body,omitempty"`
	// The unique ID of the entity
	Id string `json:"id,omitempty"`
	Links EntityLinks `json:"links,omitempty"`
	// Amount of entity surface form mentions in the article
	OverallFrequency int32 `json:"overall_frequency,omitempty"`
	// Describes how relevant an entity is to the article
	OverallProminence float32 `json:"overall_prominence,omitempty"`
	OverallSentiment EntitySentiment `json:"overall_sentiment,omitempty"`
	// The stock tickers of the entity (might be empty)
	StockTickers []string `json:"stock_tickers,omitempty"`
	Title EntityInText `json:"title,omitempty"`
	// An array of the entity types
	Types []string `json:"types,omitempty"`
}
