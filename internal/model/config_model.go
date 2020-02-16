package model

type Config struct {
	Url string `json:"url"`
	ApiKey string `json:"api_key"`
	PageSize int `json:"page_size"`
	TotalNumResults float64 `json:"total_num_results"`
	SearchImageQuery string `json:"search_query"`
}