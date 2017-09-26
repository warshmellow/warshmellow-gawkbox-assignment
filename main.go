package main

import (
	"encoding/json"
	"fmt"
	"github.com/warshmellow/warshmellow-gawkbox-assignment/twitch"
	"net/http"
	"strconv"
)

func main() {
	fmt.Println("Booting the server...")

	t := twitch.TwitchAPI{}

	// Configure routes
	http.HandleFunc("/channels", handleGetChannel(t))
	http.HandleFunc("/streams", handleGetStream(t))
	http.HandleFunc("/users", handleGetUser(t))

	// Run your server
	http.ListenAndServe(":8080", nil)
}

func handleGetChannel(t twitch.Twitchy) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Recieved the following request: %s %s \n", r.Method, r.URL)

		v := r.URL.Query()
		id, err := strconv.Atoi(v.Get("id"))
		if err != nil {
			http.Error(w, "Cannot parse id", http.StatusBadRequest)
			return
		}

		resp := t.GetChannel(id)

		respByte, _ := json.Marshal(resp)
		respStr := string(respByte)

		fmt.Fprintf(w, respStr)
	}
}

func handleGetStream(t twitch.Twitchy) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Recieved the following request: %s %s \n", r.Method, r.URL)

		v := r.URL.Query()
		id, err := strconv.Atoi(v.Get("id"))
		if err != nil {
			http.Error(w, "Cannot parse id", http.StatusBadRequest)
			return
		}

		resp := t.GetStream(id)

		respByte, _ := json.Marshal(resp)
		respStr := string(respByte)

		fmt.Fprintf(w, respStr)
	}
}
func handleGetUser(t twitch.Twitchy) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Recieved the following request: %s %s \n", r.Method, r.URL)

		v := r.URL.Query()
		id, err := strconv.Atoi(v.Get("id"))
		if err != nil {
			http.Error(w, "Cannot parse id", http.StatusBadRequest)
			return
		}

		resp := t.GetUser(id)

		respByte, _ := json.Marshal(resp)
		respStr := string(respByte)

		fmt.Fprintf(w, respStr)
	}
}
