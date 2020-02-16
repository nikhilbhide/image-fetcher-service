package downloader

import (
	"encoding/json"
	"fmt"
	"github.com/nik/image-fetcher-service/internal/model"
	"github.com/nik/image-fetcher-service/internal/utility"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

type GoogleImageDownloader struct {
	url             string
	queryParameters *QueryParameters
}

type QueryParameters struct {
	searchTerm string
	apiKey     string
	tbm        string
	pageSize   int
	device     string
	startNum   int
}

func populateMap(keyToValue map[string]string, key string, value string) {
	if (key != "" && value != "") {
		keyToValue[key] = value
	}
}

// Creates the map of key to value to be used in building url
func getKeyToValueFromQueryParameters(queryParam *QueryParameters) (map[string]string, error) {
	if (queryParam == nil || (queryParam.searchTerm == "" || queryParam.apiKey == "")) {
		return nil, errors.New("Mandatory query parameters are missing")
	} else {
		keyToValue := make(map[string]string)
		populateMap(keyToValue, "apikey", queryParam.apiKey)
		populateMap(keyToValue, "tbm", queryParam.tbm)
		populateMap(keyToValue, "device", queryParam.device)
		populateMap(keyToValue, "q", queryParam.searchTerm)
		populateMap(keyToValue, "start", strconv.Itoa(queryParam.startNum))

		return keyToValue, nil
	}
}

// It queries the api to retreive the results for results for image search based on search term.
// Before it queries the api, it validates the query parameters and builds the url.
func (downloader *GoogleImageDownloader) GetSearchResponse() (*model.QueryResponse, error) {
	var queryResponse model.QueryResponse
	keyToValue, err := getKeyToValueFromQueryParameters(downloader.queryParameters)
	//return in case of error
	if (err != nil) {
		return nil, err
	}

	url, err := utility.BuildUrlWithQueryParameters(downloader.url, keyToValue)
	if (err != nil) {
		return nil, err
	}

	//fetch the results
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		err := json.NewDecoder(response.Body).Decode(&queryResponse)
		if err != nil {
			panic(err)
		}
	}

	return &queryResponse, nil
}

// It queries the api to retreive the results for results for image search based on search term.
// Before it queries the api, it validates the query parameters and builds the url.
func (downloader *GoogleImageDownloader) GetLinks() ([]string, error) {
	queryResponse, err := downloader.GetSearchResponse()
	if (err != nil) {
		//return in case of error
		return nil, err
	}

	links := make([]string, downloader.queryParameters.pageSize)
	//iterate over the result to extract the links of the images
	for _, element := range queryResponse.ImageResults {
		links = append(links, element.SourceURL)
	}

	return links, nil
}

// Builds object for downloader based on the url and query parameters
// If successful, methods on the returned File can be used for I/O.
// If there is an error, it will be of type *PathError.
func NewDownloader(url string, apiKey string, query string) (*GoogleImageDownloader) {
	queryParam := &QueryParameters{
		apiKey:     apiKey,
		searchTerm: query,
		tbm:"isch",
	}

	downloader := GoogleImageDownloader{
		queryParameters: queryParam,
		url:             url,
	}

	return &downloader
}
