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
// StoryCluster struct for StoryCluster
type StoryCluster struct {
	// A unique identification for the cluster
	Id int64 `json:"id,omitempty"`
	// Suggested labels for the cluster
	Phrases []string `json:"phrases,omitempty"`
	// The cluster score
	Score float64 `json:"score,omitempty"`
	// Size of the cluster
	Size int32 `json:"size,omitempty"`
	// Story ids which are in the cluster
	Stories []int64 `json:"stories,omitempty"`
}
