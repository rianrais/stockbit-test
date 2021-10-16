package model

// A Response struct to map the Entire Response
type Response struct {
	Search      []SearchResult `json:"Search"`
	TotalResult string         `json:"totalResults"`
	Response    string         `json:"Response"`
}

// Struct used to map search result
type SearchResult struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	imdbID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}
