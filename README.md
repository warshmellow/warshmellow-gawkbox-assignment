## Remote Assignment V3

A REST API to get the following from Twitch v5 API:

Given a streamer’s username, get:
1. user’s channel’s # of views
1. user’s channel’s # of followers
1. user’s channel’s game
1. user’s channel’s language
1. if the user is currently streaming
1. user’s display name
1. user’s bio
1. user’s account creation date

We expose the following API:

    GET /channels?id=
    GET /streams?id=
    GET /users?id=

`channels` will get you channel info, `streams` if the user is streaming, `users` for user info.

## Getting started
### Config
Config for the Twitch API is stored in a JSON file (e.g. `twitch_config.json`) of the following form:

    {
      "client_id": "uo6dggojyb8d6soh92zknwmi5ej1q2",
      "accept_header": "application/vnd.twitchtv.v5+json",
      "get_channel_uri": "https://api.twitch.tv/kraken/channels/",
      "get_stream_uri": "https://api.twitch.tv/kraken/streams/",
      "get_user_uri": "https://api.twitch.tv/kraken/users/"
    }

A sample config is included as `twitch_config.json`.
### Build
Go get the repo with 
    
    go get -u github.com/warshmellow/warshmellow-gawkbox-assignment
    
Build with 

    go build
    
    
### Start
Execute from the command line with the first argument being the path to the config file for Twitch API. This will 
start an HTTP Server listening on 8080.

    ./[executable] twitch_config.json
    
### Try
Try the following request

    GET http://localhost:8080/channels?id=98955702

This will get you channel info for user 98955702, aka "theDemodcracy".

## API

Each endpoint uses an integer `id` query parameter. If you can't parse it, you'll receive a `400 Bad Request`.

#### Get Channel

    GET /channels?id=

Given user id, get:
1. user’s channel’s # of views
1. user’s channel’s # of followers
1. user’s channel’s game
1. user’s channel’s language
    
Sample Successful Request

    GET http://localhost:8080/channels?id=98955702

Sample Response
    
    {"_id":98955702,"followers":16246,"game":"Destiny 2","language":"en","views":197480}
    
Sample Failed Request

    GET http://localhost:8080/channels?id=1135165481
    
Response will be only the error Status Code from Twitch API. (e.g. `404`)

#### Get Stream

    GET /streams?id=

Given user id, get if the user is currently streaming. Note that if the user does not exist, you'll get
`streaming_now = false` and a `200`. This is consistent with Twitch API's `Get Streams` Endpoint.

Sample Successful Request

    GET http://localhost:8080/streams?id=98955702

Sample Response
    
    {"_id":98955702,"streaming_now":true}
    
Sample Failed Request

    GET http://localhost:8080/channels?id=1135165481
    
Sample Response
    
    {"_id":1135165481,"streaming_now":false}
    
#### Get User

    GET /users?id=

Given user id, get:
1. user’s display name
1. user’s bio
1. user’s account creation date
    
Sample Successful Request

    GET http://localhost:8080/users?id=98955702

Sample Response
    
    {
    "_id":98955702,
    "bio":"Variety speedrunner tries to live through memes to get all the PBs.",
    "created_at":"2015-08-11T21:36:31.606157Z",
    "display_name":"theDeModcracy"
    }
    
Sample Failed Request

    GET http://localhost:8080/users?id=1135165481
    
Response will be only the error Status Code from Twitch API. (e.g. `404`)

## Docker
Unfortunately I don't have access to a development machine allowing Docker. 
My development machine uses a version of Windows that doesn't allow the virtualization
needed by Docker.

Further direction with Docker:
1. Refactor the code to consume Twitch API config via environment variables
1. Build the executable
1. Put config as env. variables in Dockerfile
1. Ship docker image with the executable