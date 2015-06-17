package slackapi

import (
	"golang.org/x/net/websocket"
)

type slackRtmStartResponse struct {
	Ok    bool   `json:ok`
	Url   string `json:url`
	Error string `json:error`
}

type SlackRTM struct {
	ws websocket.Conn
}
