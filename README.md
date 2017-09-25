## Remote Assignment V3


Your instructions go here. 

	In order to accomplish the aforementioned mission we need you to build an HTTP Server
as an API with various endpoints where the request contains the streamer’s username and
returns an HTTP response. We would like to access this API to receive the following information:


Older Twitch v5 API
request contains the streamer’s username (NB: * user ID — Identifies a user, channel,
 or channel feed, depending on the endpoint.)
● user’s channel’s # of views (GET Channel)
● user’s channel’s # of followers (GET Channel)
● user’s channel’s game (GET Channel)
● user’s channel’s language (GET Channel)
● if the user is currently streaming (GET Live Stream by User)
● user’s display name (GET user)
● user’s bio (GET user)
● user’s account creation date (GET user)

Consider the following REST API:
GET channels/:id
GET streams/:id
GET users/:id

Question: useful to put channel and user info in same endpoint? benefit is you get info all at once; downside
is that Twitch itself separates into two diff endpoints. I'll have my twitch api conform more closely to theirs
for now.




GET https://api.twitch.tv/kraken/user
Response:
{
    "_id": 44322889,
    "bio": "Just a gamer playing games and chatting. :)",
    "created_at": "2013-06-03T19:12:02Z",
    "display_name": "dallas",
   ...
}

GET https://api.twitch.tv/kraken/streams/<channel ID>
{
   "stream":null
}

OR 

{
   "stream": {
      "_id": 23932774784,
      "game": "BATMAN - The Telltale Series",
      "viewers": 7254,
      "video_height": 720,
      "average_fps": 60,
      "delay": 0,
      "created_at": "2016-12-14T22:49:56Z",
      "is_playlist": false,
      "preview": {
         "small": "https://static-cdn.jtvnw.net/previews-ttv/live_user_dansgaming-80x45.jpg",
         "medium": "https://static-cdn.jtvnw.net/previews-ttv/live_user_dansgaming-320x180.jpg",
         "large": "https://static-cdn.jtvnw.net/previews-ttv/live_user_dansgaming-640x360.jpg",
         "template": "https://static-cdn.jtvnw.net/previews-ttv/live_user_dansgaming-{width}x{height}.jpg"
      },
      "channel": {
         "mature": false,
         "status": "Dan is Batman? - Telltale's Batman",
         "broadcaster_language": "en",
         "display_name": "DansGaming",
         "game": "BATMAN - The Telltale Series",
         "language": "en",
         "_id": 7236692,
         "name": "dansgaming",
         "created_at": "2009-07-15T03:02:41Z",
         "updated_at": "2016-12-15T01:33:58Z",
         "partner": true,
         "logo": "https://static-cdn.jtvnw.net/jtv_user_pictures/dansgaming-profile_image-76e4a4ab9388bc9c-300x300.png",
         "video_banner": "https://static-cdn.jtvnw.net/jtv_user_pictures/dansgaming-channel_offline_image-d3551503c24c08ad-1920x1080.png",
         "profile_banner": "https://static-cdn.jtvnw.net/jtv_user_pictures/dansgaming-profile_banner-4c2b8ece8cd010b4-480.jpeg",
         "profile_banner_background_color": null,
         "url": "https://www.twitch.tv/dansgaming",
         "views": 63906830,
         "followers": 538598
      }
   }
}

GET https://api.twitch.tv/kraken/channels/<channel ID>
{
    "_id": 44322889,
    "broadcaster_language": "en",
    "created_at": "2013-06-03T19:12:02Z",
    "display_name": "dallas",
    "followers": 40,
    "game": "Final Fantasy XV",
    "language": "en",
    "logo": "https://static-cdn.jtvnw.net/jtv_user_pictures/dallas-profile_image-1a2c906ee2c35f12-300x300.png",
    "mature": true,
    "name": "dallas",
    "partner": false,
    "profile_banner": null,
    "profile_banner_background_color": null,
    "status": "The Finalest of Fantasies",
    "updated_at": "2016-12-06T22:02:05Z",
    "url": "https://www.twitch.tv/dallas",
    "video_banner": null,
    "views": 232
}
