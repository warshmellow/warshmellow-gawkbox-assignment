package twitch

import "fmt"

type GetChannelResponse struct {
	ID        int    `json:"_id"`
	Followers int    `json:"followers"`
	Game      string `json:"game"`
	Language  string `json:"language"`
	Views     int    `json:"views"`
}

type GetStreamResponse struct {
	ID           int  `json:"_id"`
	StreamingNow bool `json:"streaming_now"`
}

type GetUserResponse struct {
	ID          int    `json:"_id"`
	Bio         string `json:"bio"`
	CreatedAt   string `json:"created_at"`
	DisplayName string `json:"display_name"`
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
