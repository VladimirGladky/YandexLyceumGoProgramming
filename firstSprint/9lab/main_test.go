package main

import (
	"io"
	"net/http"
	"testing"

	"net/http/httptest"
)

func TestRequestHandlerSuccessCase(t *testing.T) {
	expected := "Hello John"
	req := httptest.NewRequest(http.MethodGet, "/greet?name=John", nil)
	w := httptest.NewRecorder()
	RequestHandler(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if string(data) != expected {
		t.Errorf("Expected Hello John but got %v", string(data))
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("wrong status code")
	}
}

func TestRequestHandlerBadRequestCase(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/greet", nil)
	w := httptest.NewRecorder()
	RequestHandler(w, req)
	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("wrong status code")
	}
}
