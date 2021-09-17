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
// HistogramInterval struct for HistogramInterval
type HistogramInterval struct {
	// Histogram bin
	Bin int32 `json:"bin,omitempty"`
	// Histogram bin size
	Count int32 `json:"count,omitempty"`
}
