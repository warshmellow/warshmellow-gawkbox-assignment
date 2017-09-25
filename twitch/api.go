package twitch

import "fmt"

type GetChannelResponse struct {
	ID        int
	Followers int
	Game      string
	Language  string
	Views     int
}

type GetStreamResponse struct {
	ID           int
	StreamingNow bool
}

type GetUserResponse struct {
	ID          int
	Bio         string
	CreatedAt   string
	DisplayName string
}

func init() {
	fmt.Println("Initializing Twitch API...")
}

func DoSomething() {
	fmt.Println("Something.")
}

func GetChannel(id int) GetChannelResponse {
	r := GetChannelResponse{
		ID:        id,
		Followers: 333,
		Game:      "Dark Souls 3",
		Language:  "en",
		Views:     9001,
	}
	return r
}

func GetStream(id int) GetStreamResponse {
	return GetStreamResponse{ID: id, StreamingNow: true}
}

func GetUser(id int) GetUserResponse {
	return GetUserResponse{
		ID:          id,
		Bio:         "Just a gamer playing games and chatting. :)",
		CreatedAt:   "2013-06-03T19:12:02Z",
		DisplayName: "dallas",
	}
}
