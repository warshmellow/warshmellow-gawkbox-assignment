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

	// Configure a sample route
	http.HandleFunc("/channels", GetChannelHandler)

	// Run your server
	http.ListenAndServe(":8080", nil)
}

func GetChannelHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Recieved the following request: %s %s \n", r.Method, r.URL)

	v := r.URL.Query()
	id, err := strconv.Atoi(v.Get("id"))
	if err != nil {
		http.Error(w, "Cannot parse id", http.StatusBadRequest)
		return
	}

	resp := twitch.GetChannel(id)

	respByte, _ := json.Marshal(resp)
	respStr := string(respByte)

	fmt.Fprintf(w, respStr)
}

func GetStreamHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Recieved the following request: %s %s \n", r.Method, r.URL)

	v := r.URL.Query()
	id, err := strconv.Atoi(v.Get("id"))
	if err != nil {
		http.Error(w, "Cannot parse id", http.StatusBadRequest)
		return
	}

	resp := twitch.GetStream(id)

	respByte, _ := json.Marshal(resp)
	respStr := string(respByte)

	fmt.Fprintf(w, respStr)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Recieved the following request: %s %s \n", r.Method, r.URL)

	v := r.URL.Query()
	id, err := strconv.Atoi(v.Get("id"))
	if err != nil {
		http.Error(w, "Cannot parse id", http.StatusBadRequest)
		return
	}

	resp := twitch.GetUser(id)

	respByte, _ := json.Marshal(resp)
	respStr := string(respByte)

	fmt.Fprintf(w, respStr)
}
