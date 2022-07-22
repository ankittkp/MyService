package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			"HealthCheck",
			"Welcome to My Service, Author: Ankit Kumar!",
			false,
		},
		{
			"failed HealthCheck",
			"failed!",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/", nil)
			if err != nil {
				t.Errorf("HealthCheck() error = %v", err)
			}
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(HealthCheck)
			handler.ServeHTTP(rr, req)
			if status := rr.Code; status != http.StatusOK {
				t.Errorf("HealthCheck() status = %v, want %v", status, http.StatusOK)
			}

			if (rr.Body.String() != tt.want) != tt.wantErr {
				t.Errorf("HealthCheck() body = %v, want %v", rr.Body.String(), tt.want)
			}
		})
	}
}
