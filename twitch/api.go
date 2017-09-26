package twitch

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

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

type ExtAPIGetStreamResponse struct {
	Stream interface{} `json:"stream"`
}

type Twitchy interface {
	GetChannel(id int) (GetChannelResponse, error)
	GetStream(id int) (GetStreamResponse, error)
	GetUser(id int) (GetUserResponse, error)
}

type TwitchConfig struct {
	ClientID      string `json:"client_id"`
	AcceptHeader  string `json:"accept_header"`
	GetStreamURI  string `json:"get_stream_uri"`
	GetChannelURI string `json:"get_channel_uri"`
	GetUserURI    string `json:"get_user_uri"`
}

type TwitchAPI struct {
	ClientID      string
	AcceptHeader  string
	GetStreamURI  string
	GetChannelURI string
	GetUserURI    string
}

func (t TwitchAPI) NewRequest(method string, uri string, id int) (*http.Request, error) {
	fullUri := fmt.Sprintf("%s%v?client_id=%s", uri, id, t.ClientID)

	req, err := http.NewRequest(method, fullUri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", t.AcceptHeader)
	return req, nil
}

func (t TwitchAPI) GetChannel(id int) (GetChannelResponse, error) {
	client := &http.Client{}

	result := GetChannelResponse{ID: id}

	req, _ := t.NewRequest("GET", t.GetChannelURI, id)
	fmt.Printf("Sent %v\n", req)

	resp, _ := client.Do(req)

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

	req, _ := t.NewRequest("GET", t.GetStreamURI, id)
	fmt.Printf("Sent %v\n", req)

	resp, _ := client.Do(req)

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

	req, _ := t.NewRequest("GET", t.GetUserURI, id)
	fmt.Printf("Sent %v\n", req)

	resp, _ := client.Do(req)

	if resp.StatusCode == http.StatusOK {

		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, &result)

		return result, nil
	} else {
		return result, errors.New(strconv.Itoa(resp.StatusCode))
	}
}
