package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_myHandler(t *testing.T) {
	type args struct {
		url          string
		expectedName string
		status       int
	}
	tests := []struct {
		name string
		args args
	}{
		{"t1", args{"/json/test", "test", http.StatusOK}},
		{"t2", args{"/x/y", "", http.StatusNotImplemented}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := tt.args.url
			req, err := http.NewRequest(http.MethodGet, url, nil)
			if err != nil {
				t.Errorf("could't create request: %v", tt.args.url)
			}
			recorder := httptest.NewRecorder()
			myHandler(recorder, req)
			if recorder.Result().StatusCode != tt.args.status {
				t.Errorf("Wrong status code %v, expected %v", recorder.Result().StatusCode, tt.args.status)
			}
			if !strings.Contains(recorder.Body.String(), tt.args.expectedName) {
				t.Errorf("The response does't contain '%v': '%v'", tt.args.expectedName, recorder.Body.String())
			}
		})
	}
}
