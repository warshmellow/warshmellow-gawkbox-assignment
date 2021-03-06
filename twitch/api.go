package twitch

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

/*
Get*Response structs are data containers to send back to consumer
*/
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

/*
Twitch API specific response from Get Streams endpoint
*/
type ExtAPIGetStreamResponse struct {
	Stream interface{} `json:"stream"`
}

/*
Consumers will create a struct supporting this interface, which allows for mocking in tests
*/
type Twitchy interface {
	GetChannel(id int) (GetChannelResponse, error)
	GetStream(id int) (GetStreamResponse, error)
	GetUser(id int) (GetUserResponse, error)
}

/*
Twitch API config
*/
type TwitchConfig struct {
	ClientID      string `json:"client_id"`
	AcceptHeader  string `json:"accept_header"`
	GetStreamURI  string `json:"get_stream_uri"`
	GetChannelURI string `json:"get_channel_uri"`
	GetUserURI    string `json:"get_user_uri"`
}

/*
Supports the Twitchy interface and will talk to Twitch v5 API
*/
type TwitchAPI struct {
	ClientID      string
	AcceptHeader  string
	GetStreamURI  string
	GetChannelURI string
	GetUserURI    string
}

/*
Creates a Request that contains the Client ID and Accept Header for Twitch
*/
func (t TwitchAPI) NewRequest(method string, uri string, id int) (*http.Request, error) {
	fullUri := fmt.Sprintf("%s%v?client_id=%s", uri, id, t.ClientID)

	req, err := http.NewRequest(method, fullUri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", t.AcceptHeader)
	return req, nil
}

/*
Get* take user id and call Twitch API and get data, following usual error handling
*/
func (t TwitchAPI) GetChannel(id int) (GetChannelResponse, error) {
	client := &http.Client{}

	result := GetChannelResponse{ID: id}

	req, err := t.NewRequest("GET", t.GetChannelURI, id)
	if err != nil {
		return result, err
	}
	fmt.Printf("Sent %v\n", req)

	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}

	if resp.StatusCode == http.StatusOK {

		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, &result)

		return result, nil
	} else {
		return result, errors.New(strconv.Itoa(resp.StatusCode))
	}
}

func (t TwitchAPI) GetStream(id int) (GetStreamResponse, error) {
	client := &http.Client{}

	result := GetStreamResponse{ID: id}
	extResp := ExtAPIGetStreamResponse{}

	req, err := t.NewRequest("GET", t.GetStreamURI, id)
	if err != nil {
		return result, err
	}
	fmt.Printf("Sent %v\n", req)

	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}

	if resp.StatusCode == http.StatusOK {

		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, &extResp)

		result.StreamingNow = extResp.Stream != nil

		return result, nil
	} else {
		return result, errors.New(strconv.Itoa(resp.StatusCode))
	}
}

func (t TwitchAPI) GetUser(id int) (GetUserResponse, error) {
	client := &http.Client{}

	result := GetUserResponse{ID: id}

	req, err := t.NewRequest("GET", t.GetUserURI, id)
	if err != nil {
		return result, err
	}
	fmt.Printf("Sent %v\n", req)

	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}

	if resp.StatusCode == http.StatusOK {

		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, &result)

		return result, nil
	} else {
		return result, errors.New(strconv.Itoa(resp.StatusCode))
	}
}
