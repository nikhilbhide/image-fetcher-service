package utility

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"errors"
	"github.com/nik/image-fetcher-service/internal/model"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

// Loads the configuration from the properties  file
func LoadConfiguration(file string) (*model.Config, error) {
	var config model.Config
	configFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	return &config, nil
}

// BuildUrlWithQueryParameters builds url by combining the base url and query parameters
func BuildUrlWithQueryParameters(baseUrl string, keyToValue map[string]string) (string, error) {
	url, err := url.Parse(baseUrl)
	if err == nil {
		if keyToValue == nil || len(keyToValue) == 0 {
			// if there are no query parameters passed
			return baseUrl, nil
		}

		//build url by adding query param
		queryParam := url.Query()
		for key, value := range keyToValue {
			queryParam.Set(key, value)
		}

		url.RawQuery = queryParam.Encode()
		return url.String(), nil

	} else {
		return "", errors.New("Wrong url format")
	}
}

//Download the images as byte array
func DownloadImage(urlPath string) ([]byte, error) {
	resp, err := http.Get(urlPath)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Error in downloading images")
	}

	return responseData, nil
}

//Convert object to byte array
func GetBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
