/*
 * AYLIEN News API
 *
 * The AYLIEN News API is the most powerful way of sourcing, searching and syndicating analyzed and enriched news content. It is accessed by sending HTTP requests to our server, which returns information to your client. 
 *
 * API version: 5.2.3
 * Contact: support@aylien.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package newsapi
// Clusters struct for Clusters
type Clusters struct {
	// The total number of clusters
	ClusterCount int64 `json:"cluster_count,omitempty"`
	// An array of clusters
	Clusters []Cluster `json:"clusters,omitempty"`
	// The next page cursor
	NextPageCursor string `json:"next_page_cursor,omitempty"`
}
