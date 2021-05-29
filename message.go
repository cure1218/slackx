package slackx

import (
	"github.com/slack-go/slack"
)

func (s *Slackx) PostMessageToChannel(chanID, text string, escape bool) (string, string, error) {
	return s.PostMessage(chanID, slack.MsgOptionText(text, escape))
}

//type Msg struct {
//	// Basic Message
//	ClientMsgID     string       `json:"client_msg_id,omitempty"`
//	Type            string       `json:"type,omitempty"`
//	Channel         string       `json:"channel,omitempty"`
//	User            string       `json:"user,omitempty"`
//	Text            string       `json:"text,omitempty"`
//	Timestamp       string       `json:"ts,omitempty"`
//	ThreadTimestamp string       `json:"thread_ts,omitempty"`
//	IsStarred       bool         `json:"is_starred,omitempty"`
//	PinnedTo        []string     `json:"pinned_to,omitempty"`
//	Attachments     []Attachment `json:"attachments,omitempty"`
//	Edited          *Edited      `json:"edited,omitempty"`
//	LastRead        string       `json:"last_read,omitempty"`
//	Subscribed      bool         `json:"subscribed,omitempty"`
//	UnreadCount     int          `json:"unread_count,omitempty"`
//
//	// Message Subtypes
//	SubType string `json:"subtype,omitempty"`
//
//	// Hidden Subtypes
//	Hidden           bool   `json:"hidden,omitempty"`     // message_changed, message_deleted, unpinned_item
//	DeletedTimestamp string `json:"deleted_ts,omitempty"` // message_deleted
//	EventTimestamp   string `json:"event_ts,omitempty"`
//
//	// bot_message (https://api.slack.com/events/message/bot_message)
//	BotID      string      `json:"bot_id,omitempty"`
//	Username   string      `json:"username,omitempty"`
//	Icons      *Icon       `json:"icons,omitempty"`
//	BotProfile *BotProfile `json:"bot_profile,omitempty"`
//
//	// channel_join, group_join
//	Inviter string `json:"inviter,omitempty"`
//
//	// channel_topic, group_topic
//	Topic string `json:"topic,omitempty"`
//
//	// channel_purpose, group_purpose
//	Purpose string `json:"purpose,omitempty"`
//
//	// channel_name, group_name
//	Name    string `json:"name,omitempty"`
//	OldName string `json:"old_name,omitempty"`
//
//	// channel_archive, group_archive
//	Members []string `json:"members,omitempty"`
//
//	// channels.replies, groups.replies, im.replies, mpim.replies
//	ReplyCount   int     `json:"reply_count,omitempty"`
//	Replies      []Reply `json:"replies,omitempty"`
//	ParentUserId string  `json:"parent_user_id,omitempty"`
//
//	// file_share, file_comment, file_mention
//	Files []File `json:"files,omitempty"`
//
//	// file_share
//	Upload bool `json:"upload,omitempty"`
//
//	// file_comment
//	Comment *Comment `json:"comment,omitempty"`
//
//	// pinned_item
//	ItemType string `json:"item_type,omitempty"`
//
//	// https://api.slack.com/rtm
//	ReplyTo int    `json:"reply_to,omitempty"`
//	Team    string `json:"team,omitempty"`
//
//	// reactions
//	Reactions []ItemReaction `json:"reactions,omitempty"`
//
//	// slash commands and interactive messages
//	ResponseType    string `json:"response_type,omitempty"`
//	ReplaceOriginal bool   `json:"replace_original"`
//	DeleteOriginal  bool   `json:"delete_original"`
//
//	// Block type Message
//	Blocks Blocks `json:"blocks,omitempty"`
//}
