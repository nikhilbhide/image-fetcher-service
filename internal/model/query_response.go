package model

type QueryResponse struct {
	Query struct {
		Apikey string `json:"apikey"`
		Q      string `json:"q"`
		Tbm    string `json:"tbm"`
		Device string `json:"device"`
		URL    string `json:"url"`
	} `json:"query"`
	RelatedSearches []interface{} `json:"related_searches"`
	ImageResults    []struct {
		Position  int    `json:"position"`
		Thumbnail string `json:"thumbnail"`
		SourceURL string `json:"sourceUrl"`
		Title     string `json:"title"`
		Link      string `json:"link"`
		Source    string `json:"source"`
	} `json:"image_results"`
}
