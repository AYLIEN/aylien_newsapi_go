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
// CategoryTaxonomy The taxonomy of the category
type CategoryTaxonomy string

// List of CategoryTaxonomy
const (
	IAB_QAG CategoryTaxonomy = "iab-qag"
	IPTC_SUBJECTCODE CategoryTaxonomy = "iptc-subjectcode"
	AYLIEN CategoryTaxonomy = "aylien"
)
