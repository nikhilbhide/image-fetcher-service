package downloader

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

func TestGoogleImageDownloader_populateMap(t *testing.T) {
	type args struct {
		keyToValue map[string]string
		key        string
		value      string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			populateMap(tt.args.keyToValue, tt.args.key, tt.args.value)
		})
	}
}

func Test_getKeyToValueFromQueryParameters(t *testing.T) {
	type args struct {
		queryParam *QueryParameters
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr bool
	}{
		{
			name:    "test_success_no_parameters",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "test_success_empty_parameters",
			args:    args{},
			want:    nil,
			wantErr: true,
		},
		{
			name: "test_success_parameters_with_no_values",
			args: args{
				queryParam: &QueryParameters{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "test_success_parameters_with_values",
			args: args{
				queryParam: &QueryParameters{
					apiKey:     "akldjsad",
					searchTerm: "q",
					device:     "pc",
				},
			},
			want: map[string]string{
				"apikey": "akldjsad",
				"q":      "q",
				"device": "pc",
				"start":  "0",
			},
			wantErr: false,
		},
		{
			name: "test_success_parameters_with_start_num",
			args: args{
				queryParam: &QueryParameters{
					apiKey:     "akldjsad",
					searchTerm: "q",
					device:     "pc",
					startNum:   100,
				},
			},
			want: map[string]string{
				"apikey": "akldjsad",
				"q":      "q",
				"device": "pc",
				"start":  "100",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getKeyToValueFromQueryParameters(tt.args.queryParam)
			if (err != nil) != tt.wantErr {
				t.Errorf("getKeyToValueFromQueryParameters() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getKeyToValueFromQueryParameters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGoogleImageDownloader_GetSearchResponse(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.Write(loadMockResponse())
		}),
	)
	// Close the server when test finishes
	defer server.Close()
	d := NewDownloader(server.URL, "test", "xyz")
	queryResponse, _ := d.GetSearchResponse()
	//get actual values
	apiKeyGot := queryResponse.Query.Apikey
	deviceGot := queryResponse.Query.Device
	//get desired values
	apiKeyWant := queryResponse.Query.Apikey
	deviceWant := queryResponse.Query.Device

	if !reflect.DeepEqual(apiKeyGot, apiKeyWant) {
		t.Errorf("GetSearchResponse() = %v, want %v", apiKeyGot, apiKeyWant)
	}

	if !reflect.DeepEqual(deviceGot, deviceWant) {
		t.Errorf("GetSearchResponse() = %v, want %v", deviceGot, deviceWant)
	}
}

func TestGoogleImageDownloader_GetLinks(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.Write(loadMockResponse())
		}),
	)
	// Close the server when test finishes
	defer server.Close()
	d := NewDownloader(server.URL, "test", "xyz")
	queryResponse, _ := d.GetSearchResponse()
	linksWant := []string{"https://www.outbrain.com/techblog/2017/05/effective-testing-with-loan-pattern-in-scala/",
		"https://the-test-fun-for-friends.en.softonic.com/android",
		"https://www.spectrum.com/internet/speed-test.html"}
	//get actual values
	linksGot := []string{}
	for _, imageResult := range queryResponse.ImageResults {
		linksGot = append(linksGot, imageResult.Link)
	}

	if !reflect.DeepEqual(linksGot, linksWant) {
		t.Errorf("GetLinks() = %v, want %v", linksGot, linksWant)
	}
}

func TestGoogleImageDownloader_GetImages(t *testing.T) {
	server1 := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.Write(loadGetImageMockResponse()[0])
		}),
	)
	// Close the server when test finishes
	defer server1.Close()

	server2 := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.Write(loadGetImageMockResponse()[1])
		}),
	)
	// Close the server when test finishes
	defer server2.Close()

	server3 := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.Write(loadGetImageMockResponse()[2])
		}),
	)
	// Close the server when test finishes
	defer server3.Close()

	linksWant := loadGetImageMockResponse()
	var linksGot [][]byte

	d := NewDownloader("test", "test", "xyz")
	urls := []string{server1.URL, server2.URL, server3.URL}
	urlToImages, _ := d.GetImages(urls)

	if urlToImages == nil {
		t.Errorf("GetImages() length of the returned images = %v, want %v", 0, len(linksWant))
	} else if len(urlToImages) != len(linksWant) {
		t.Errorf("GetImages() length of the returned images = %v, want %v", len(urlToImages), len(linksWant))
	}

	linksGot = append(linksGot, urlToImages[server1.URL])
	linksGot = append(linksGot, urlToImages[server2.URL])
	linksGot = append(linksGot, urlToImages[server3.URL])

	if !reflect.DeepEqual(linksGot, linksWant) {
		t.Errorf("GetImages() = %v, want %v", linksGot, linksWant)
	}
}

func TestNewDownloader(t *testing.T) {
	type args struct {
		url    string
		apiKey string
		query  string
	}
	tests := []struct {
		name string
		args args
		want *GoogleImageDownloader
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDownloader(tt.args.url, tt.args.apiKey, tt.args.query); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDownloader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func loadMockResponse() []byte {
	jsonFile, err := os.Open("/app/image-fetcher-service/test/test_response.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// if we os.Open returns an error then raise panic
	if err != nil {
		panic(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue
}

//mock image response
func loadGetImageMockResponse() [][]byte {
	image1 := []byte{'r', 'o', 'a', 'd'}
	image2 := []byte{'t', 'e', 's', 't'}
	image3 := []byte{'n', 'i', 'l'}
	images := [][]byte{image1, image2, image3}

	return images
}
