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
	}
	tests := []struct {
		name string
		args args
	}{
		{"simple URL", args{"/json/testing", "testing"}},
		{"simple URL", args{"/json/test/2", "test"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request, err := http.NewRequest(http.MethodGet, tt.args.url, nil)
			if err != nil {
				t.Errorf("could not create request for url: %v", tt.args.url)
			}
			recorder := httptest.NewRecorder()
			myHandler(recorder, request)
			if !strings.Contains(recorder.Body.String(), tt.args.expectedName) {
				t.Errorf("The response does't contain %v", tt.args.expectedName)
			}
		})
	}
}

