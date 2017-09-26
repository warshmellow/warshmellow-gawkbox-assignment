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

	t := twitch.TwitchAPI{
		ClientID:      "l3p9c840mspj37hw3845gnfu0pg2ar",
		AcceptHeader:  `application/vnd.twitchtv.v5+json`,
		GetChannelURI: "https://api.twitch.tv/kraken/channels/",
		GetStreamURI:  "https://api.twitch.tv/kraken/streams/",
		GetUserURI:    "https://api.twitch.tv/kraken/users/",
	}

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

		resp, err := t.GetChannel(id)
		if err != nil {
			errStatusCode, _ := strconv.Atoi(err.Error())
			http.Error(w, err.Error(), errStatusCode)
			return
		}

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

		resp, err := t.GetStream(id)
		if err != nil {
			errStatusCode, _ := strconv.Atoi(err.Error())
			http.Error(w, err.Error(), errStatusCode)
			return
		}

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

		resp, err := t.GetUser(id)
		if err != nil {
			errStatusCode, _ := strconv.Atoi(err.Error())
			http.Error(w, err.Error(), errStatusCode)
			return
		}

		respByte, _ := json.Marshal(resp)
		respStr := string(respByte)

		fmt.Fprintf(w, respStr)
	}
}
