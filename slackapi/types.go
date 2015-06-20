package slackapi

import (
	"golang.org/x/net/websocket"
	"time"
)

type RTMChatType int

const (
	Channel RTMChatType = iota
	Group
	DM
)

type slackRtmStartResponse struct {
	Ok    bool   `json:ok`
	Url   string `json:url`
	Error string `json:error`
}

type slackRtmEvent struct {
	Type             string         `json:type`
	Subtype          *string        `json:subtype`
	Hidden           *bool          `json:hidden`
	DeletedTimestamp *string        `json:deleted_ts`
	EventTimestamp   *string        `json:event_ts`
	Timestamp        *string        `json:ts`
	User             *string        `json:user`
	Text             *string        `json:text`
	Error            *slackRtmError `json:error`
}

type slackRtmError struct {
	Code    int    `json:code`
	Message string `json:msg`
}

type SlackRTM struct {
	ws       websocket.Conn
	Roster   []SlackUser
	Channels []Chat
	Groups   []Chat
	DMs      []Chat
}

type SlackUser struct {
	UserId string
	Name   string
}

type Chat struct {
	Type RTMChatType
	Name string
}

type SlackMessage struct {
	User      SlackUser
	Chat      Chat
	Text      string
	Timestamp time.Time
}
