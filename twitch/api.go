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

type ExtAPIGetChannelResponse struct {
	ID        int    `json:"_id"`
	Followers int    `json:"followers"`
	Game      string `json:"game"`
	Language  string `json:"language"`
	Views     int    `json:"views"`
}

type ExtAPIGetStreamResponse struct {
	Stream interface{} `json:"stream"`
}

type ExtAPIGetUserResponse struct {
	ID          int    `json:"_id"`
	Bio         string `json:"bio"`
	CreatedAt   string `json:"created_at"`
	DisplayName string `json:"display_name"`
}

type Twitchy interface {
	GetChannel(id int) (GetChannelResponse, error)
	GetStream(id int) (GetStreamResponse, error)
	GetUser(id int) (GetUserResponse, error)
}

type TwitchAPI struct {
	ClientID      string
	GetStreamURI  string
	GetChannelURI string
	GetUserIdURI  string
}

func (t TwitchAPI) GetChannel(id int) (GetChannelResponse, error) {
	client := &http.Client{}

	result := GetChannelResponse{ID: id}

	uri := fmt.Sprintf("%s%v?client_id=%s", t.GetChannelURI, id, t.ClientID)
	fmt.Printf("Sent request to %s", uri)

	req, _ := http.NewRequest("GET", uri, nil)
	req.Header.Add("Accept", `application/vnd.twitchtv.v5+json"`)

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

	uri := fmt.Sprintf("%s%v?client_id=%s", t.GetStreamURI, id, t.ClientID)
	fmt.Printf("Sent request to %s", uri)

	req, _ := http.NewRequest("GET", uri, nil)
	req.Header.Add("Accept", `application/vnd.twitchtv.v5+json"`)

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

	uri := fmt.Sprintf("%s%v?client_id=%s", t.GetUserIdURI, id, t.ClientID)
	fmt.Printf("Sent request to %s", uri)

	req, _ := http.NewRequest("GET", uri, nil)
	req.Header.Add("Accept", `application/vnd.twitchtv.v5+json"`)

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
