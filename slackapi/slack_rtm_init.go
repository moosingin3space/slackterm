package slackapi

import (
	"encoding/json"
	"errors"
	"golang.org/x/net/websocket"
	"io/ioutil"
	"net/http"
)

const rtmUrl = "https://slack.com/api/rtm.start"

func DialSlack(token string) (*SlackRTM, error) {
	// Firstly, get the URL
	var err error
	if resp, err := http.Get(rtmUrl + "?token=" + token); err != nil {
		return nil, err
	}

	if data, err := ioutil.ReadAll(resp); err != nil {
		return nil, err
	}

	var rtmResponse slackRtmStartResponse
	if err = json.Unmarshal(data, rtmResponse); err != nil {
		return nil, err
	}

	if !rtmResponse.Ok {
		return nil, errors.New(rtmResponse.Error)
	}

	ws, err := websocket.Dial(rtmResponse.Url, "", "http://localhost/")
	if err != nil {
		return nil, err
	}

	return &SlackRTM{
		ws:       ws,
		TeamName: rtmResponse.Team.Name,
		Roster:   rtmResponse.Users,
		Channels: rtmResponse.Channels,
		Groups:   rtmResponse.Groups,
		DMs:      rtmResponse.IMs,
	}, nil
}
