package utility

import (
	"github.com/nik/image-fetcher-service/internal/model"
	"path"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func TestLoadConfiguration(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		args    args
		want    *model.Config
		wantErr bool
	}{
		{name: "test_success_test_json",
			args: args{
				file: getFilePath(),
			},
			want: &model.Config{
				ApiKey:           "test",
				Url:              "https://app.zenserp.com/api/v2/search",
				PageSize:         100,
				TotalNumResults:  100,
				SearchImageQuery: "test_search",
			},
		},
		{name: "test_success_test_json",
			args: args{
				file: "file_not_exists.json",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadConfiguration(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadConfiguration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadConfiguration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildUrlWithQueryParameters(t *testing.T) {
	type args struct {
		baseUrl    string
		keyToValue map[string]string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"test_build_url_with_no_parameters",
			args{
				baseUrl: "http://www.google.com",
			},
			"http://www.google.com",
			false,
		},
		// TODO: Add test cases.
		{"test_build_url_with_erroneous_url",
			args{
				baseUrl: "www@!%^Hgoogl",
			},
			"",
			true,
		},
		// TODO: Add test cases.
		{"test_build_url_with_query_parameters",
			args{
				baseUrl: "http://www.google.com",
				keyToValue: map[string]string{
					"query": "1355",
					"r":     "2138",
				},
			},
			"http://www.google.com?query=1355&r=2138",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BuildUrlWithQueryParameters(tt.args.baseUrl, tt.args.keyToValue)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildUrlWithQueryParameters() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("BuildUrlWithQueryParameters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getFilePath() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(filepath.Join(b, "../../"))
	rootDir := filepath.Dir(d)
	return filepath.FromSlash(rootDir + "/config/test.json")
}
