package main

import (
	"fmt"
	"github.com/warshmellow/warshmellow-gawkbox-assignment/twitch"
	"net/http"
)

func main() {
	fmt.Println("Booting the server...")

	// Configure a sample route
	http.HandleFunc("/channels", GetChannelHandler)

	// Run your server
	http.ListenAndServe(":8080", nil)
}

// myHandlerFunc - A sample handler function for the route /sample_route for your HTTP server
func GetChannelHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recieved the following request:", r.Body)

	twitch.DoSomething()

	// YOUR ROUTES LOGIC GOES HERE
	//
	// Feel free to structure your routing however you see fit, this is just an example to get you started.

	fmt.Fprintf(w, "hello")
}
