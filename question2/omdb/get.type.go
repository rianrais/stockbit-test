package omdb

// OmdbResponse is used to get result from OMDB API
type OmdbResponse struct {
	Search      []SearchResult `json:"Search"`
	TotalResult string         `json:"totalResults"`
	Response    string         `json:"Response"`
}

// SearchResult is used to map search result
type SearchResult struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

// GetResponse is used as the response of the API
type GetResponse struct {
	CurrentPage int            `json:"currentPage"`
	TotalPage   int            `json:"totalPage"`
	Data        []SearchResult `json:"data"`
}
