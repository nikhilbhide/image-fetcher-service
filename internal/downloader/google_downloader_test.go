package downloader

import (
	"path"
	"path/filepath"
	"reflect"
	"runtime"
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
			name:"test_success_no_parameters",
			want: nil,
			wantErr:true,
		},
		{
			name:"test_success_empty_parameters",
			args:args{

			},
			want: nil,
			wantErr:true,

		},
		{
			name:"test_success_parameters_with_no_values",
			args:args{
				queryParam:&QueryParameters{

				},
			},
			want: nil,
			wantErr:true,
		},
		{
			name:"test_success_parameters_with_values",
			args:args{
				queryParam:&QueryParameters{
					apiKey:"akldjsad",
					searchTerm:"q",
					device:"pc",
				},
			},
			want: map[string]string{
				"apikey":"akldjsad",
				"q":"q",
				"device":"pc",
				"start":"0",
			},
			wantErr:false,
		},
		{
			name:"test_success_parameters_with_start_num",
			args:args{
				queryParam:&QueryParameters{
					apiKey:"akldjsad",
					searchTerm:"q",
					device:"pc",
					startNum:100,
				},
			},
			want: map[string]string{
				"apikey":"akldjsad",
				"q":"q",
				"device":"pc",
				"start":"100",
			},
			wantErr:false,
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

func TestGoogleImageDownloader_GetLinks(t *testing.T) {
	type fields struct {
		url             string
		queryParameters *QueryParameters
	}
	tests := []struct {
		name    string
		fields  fields
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			downloader := &GoogleImageDownloader{
				url:             tt.fields.url,
				queryParameters: tt.fields.queryParameters,
			}
			got, err := downloader.GetLinks()
			if (err != nil) != tt.wantErr {
				t.Errorf("GoogleImageDownloader.GetLinks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GoogleImageDownloader.GetLinks() = %v, want %v", got, tt.want)
			}
		})
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

func getFilePath() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(filepath.Join(b, "../../"))
	rootDir:= filepath.Dir(d)
	return filepath.FromSlash(rootDir+"/test/test_response.json")
}