// news_aggregator.go
package utils

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

// ArxivResponse represents the top-level feed from arXiv.
type ArxivResponse struct {
	XMLName xml.Name     `xml:"feed"`
	Entries []ArxivEntry `xml:"entry"`
}

// ArxivEntry represents a single research paper in the arXiv feed.
type ArxivEntry struct {
	Title     string        `xml:"title"`
	Summary   string        `xml:"summary"`
	ID        string        `xml:"id"`
	Published string        `xml:"published"`
	Updated   string        `xml:"updated"`
	Authors   []ArxivAuthor `xml:"author"`
	Links     []ArxivLink   `xml:"link"`
}

// ArxivAuthor represents an author element.
type ArxivAuthor struct {
	Name string `xml:"name"`
}

// ArxivLink represents a link element (for example, to the PDF).
type ArxivLink struct {
	Href string `xml:"href,attr"`
	Rel  string `xml:"rel,attr"`
}

// Mapping from your category keys to arXiv categories.
// Adjust these mappings as needed based on your domain.
var arxivCategoryMapping = map[string]string{
	"ml":           "cs.LG", // machine learning (or stat.ML, if you prefer)
	"quantum":      "quant-ph",
	"neuroscience": "q-bio.NC", // example: neural computation in quantitative biology
	"genetics":     "q-bio.GN",
	"renewables":   "physics.optics", // or another related category
	"astrophysics": "astro-ph",
	"robotics":     "cs.RO", // if available or use a related category
	"biotech":      "q-bio.BM",
	"blockchain":   "cs.CR", // may need a custom approach; arXiv doesn't have a direct blockchain category
	"materials":    "cond-mat.mtrl-sci",
	"medicine":     "q-bio.MN",
	"social":       "cs.CY", // computational social science isn’t directly supported, so you might choose an adjacent category
	"engineering":  "eess",
	"cs":           "cs",
	"data_science": "stat.ML", // or "cs.LG", depending on your focus
	"economics":    "econ.EM", // note: economics papers on arXiv are limited; may not have a large presence
}

func BuildArxivQuery(userInterests []string) string {
	var queries []string
	for _, interest := range userInterests {
		if arxivCat, ok := arxivCategoryMapping[interest]; ok {
			// Prepend "cat:" to specify a category search.
			queries = append(queries, fmt.Sprintf("cat:%s", arxivCat))
		}
	}
	// Combine all category queries with the OR operator.
	// The arXiv API expects the Boolean operators to be URL-encoded,
	// but you can build the string and then URL-encode it later if needed.
	return strings.Join(queries, "+OR+")
}

// FetchArxivPapers retrieves research papers from arXiv based on the provided query string.
// The query can follow arXiv’s search syntax (e.g., "all:machine+learning" or "cat:cs.LG").
func FetchArxivPapers(query string) (ArxivResponse, error) {
	log.Println("Fetching research papers from arXiv")

	// Construct the arXiv API URL.
	// For example, to search in all fields: "all:machine+learning"
	// encodedQuery := url.QueryEscape(query)
	queryURL := fmt.Sprintf("http://export.arxiv.org/api/query?search_query=%s&start=0&max_results=10", query)

	log.Println("Query URL:", queryURL)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(queryURL)
	if err != nil {
		return ArxivResponse{}, err
	}
	defer resp.Body.Close()

	// Check the HTTP response status code.
	if resp.StatusCode != http.StatusOK {
		return ArxivResponse{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ArxivResponse{}, err
	}

	var arxivResp ArxivResponse
	if err := xml.Unmarshal(body, &arxivResp); err != nil {
		return ArxivResponse{}, err
	}

	return arxivResp, nil
}
