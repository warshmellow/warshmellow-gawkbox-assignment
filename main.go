package main

import (
	"encoding/json"
	"fmt"
	"github.com/warshmellow/warshmellow-gawkbox-assignment/twitch"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Booting the server...")

	// Read config.json from directory
	filename := "twitch_config.json"
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Can't read Twitch API config %v\n", filename)
		os.Exit(1)
	}

	twitchConfig := twitch.TwitchConfig{}
	json.Unmarshal(dat, &twitchConfig)
	fmt.Printf("Loaded Twitch Config: %v\n", twitchConfig)

	t := twitch.TwitchAPI{
		ClientID:      twitchConfig.ClientID,
		AcceptHeader:  twitchConfig.AcceptHeader,
		GetChannelURI: twitchConfig.GetChannelURI,
		GetStreamURI:  twitchConfig.GetStreamURI,
		GetUserURI:    twitchConfig.GetUserURI,
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
