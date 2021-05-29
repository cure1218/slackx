package slackx

import (
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

//================================================================
//
//================================================================
type Slackx struct {
	BotToken      string
	AppLevelToken string
	BotID         string
	UserID        string
	*slack.Client
	AppMentionHook func(*Slackx, *slackevents.AppMentionEvent, string) error
}

func New(botToken, appLevelToken string) (*Slackx, error) {
	client := slack.New(botToken, slack.OptionAppLevelToken(appLevelToken))
	if resp, err := client.AuthTest(); err != nil {
		return nil, err
	} else {
		return &Slackx{
			BotToken:       botToken,
			AppLevelToken:  appLevelToken,
			BotID:          resp.BotID,
			UserID:         resp.UserID,
			Client:         client,
			AppMentionHook: nil,
		}, nil
	}
}

func (s *Slackx) SetAppMentionHook(hook func(*Slackx, *slackevents.AppMentionEvent, string) error) *Slackx {
	s.AppMentionHook = hook
	return s
}
