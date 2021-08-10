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
// Media struct for Media
type Media struct {
	// The content length of media
	ContentLength int32 `json:"content_length,omitempty"`
	Format MediaFormat `json:"format,omitempty"`
	// The height of media
	Height int32 `json:"height,omitempty"`
	Type MediaType `json:"type,omitempty"`
	// A URL which points to the media file
	Url string `json:"url,omitempty"`
	// The width of media
	Width int32 `json:"width,omitempty"`
}
