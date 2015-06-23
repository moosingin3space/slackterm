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

type slackRtmStartTeam struct {
	Name string `json:name`
}

type slackRtmStartResponse struct {
	Ok       bool              `json:ok`
	Url      string            `json:url`
	Error    string            `json:error`
	Team     slackRtmStartTeam `json:team`
	Users    []SlackUser       `json:users`
	Channels []Chat            `json:channels`
	Groups   []Chat            `json:groups`
	IMs      []Chat            `json:ims`
}

type slackRtmEvent struct {
	Type             string         `json:type`
	Subtype          *string        `json:subtype`
	Hidden           bool           `json:hidden`
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
	TeamName string
	Roster   []SlackUser
	Channels []Chat
	Groups   []Chat
	DMs      []Chat
}

type SlackUser struct {
	UserId string `json:id`
	Name   string `json:name`
	// TODO look at adding profile information
}

type Chat struct {
	Name      string `json:name`
	IsChannel bool   `json:is_channel`
	IsGroup   bool   `json:is_group`
	IsIM      bool   `json:is_im`
}

type SlackMessage struct {
	User      SlackUser
	Chat      Chat
	Text      string
	Timestamp time.Time
}
