package twitch

import "fmt"

type GetChannelResponse struct {
	ID        int
	Followers int
	Game      string
	Language  string
	Views     int
}

func init() {
	fmt.Println("Initializing Twitch API...")
}

func DoSomething() {
	fmt.Println("Something.")
}

func GetChannel(id int) GetChannelResponse {
	r := GetChannelResponse{
		ID:        1,
		Followers: 333,
		Game:      "Dark Souls 3",
		Language:  "en",
		Views:     9001,
	}
	return r
}
