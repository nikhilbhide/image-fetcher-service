package downloader

import "github.com/nik/image-fetcher-service/internal/model"

// Captures the use cases of an image downloader
type Downloader interface {
	GetSearchResponse() (*model.QueryResponse, error)
	GetLinks() ([]string, error)
}
