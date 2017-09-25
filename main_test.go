package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetChannelsHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/channels?id=1", nil)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetChannelHandler)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check body is non-empty
	if len(rr.Body.String()) == 0 {
		t.Errorf("body should be nonempty")
	}
}

func TestGetStreamHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/streams?id=1", nil)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetStreamHandler)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check body is non-empty
	if len(rr.Body.String()) == 0 {
		t.Errorf("body should be nonempty")
	}
}

func TestGetUserHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/users?id=1", nil)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetUserHandler)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check body is non-empty
	if len(rr.Body.String()) == 0 {
		t.Errorf("body should be nonempty")
	}
}
