package main

import (
	"github.com/warshmellow/warshmellow-gawkbox-assignment/twitch"
	"net/http"
	"net/http/httptest"
	"testing"
)

var mux *http.ServeMux

type TestTwitchy struct{}

func (t TestTwitchy) GetChannel(id int) twitch.GetChannelResponse {
	return twitch.GetChannelResponse{
		ID:        id,
		Followers: 333,
		Game:      "Dark Souls 3",
		Language:  "en",
		Views:     9001,
	}
}

func (t TestTwitchy) GetStream(id int) twitch.GetStreamResponse {
	return twitch.GetStreamResponse{ID: id, StreamingNow: true}
}

func (t TestTwitchy) GetUser(id int) twitch.GetUserResponse {
	return twitch.GetUserResponse{
		ID:          id,
		Bio:         "Just a gamer playing games and chatting. :)",
		CreatedAt:   "2013-06-03T19:12:02Z",
		DisplayName: "dallas",
	}
}

func init() {
	mux = http.NewServeMux()

	testTwitchy := TestTwitchy{}

	mux.HandleFunc("/channels", handleGetChannel(testTwitchy))
	mux.HandleFunc("/streams", handleGetStream(testTwitchy))
	mux.HandleFunc("/users", handleGetUser(testTwitchy))
}

func TestGetChannelsHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/channels?id=1", nil)

	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

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
	mux.ServeHTTP(rr, req)

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
	mux.ServeHTTP(rr, req)

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
