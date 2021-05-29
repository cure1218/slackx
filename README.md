# slackx
Slack package wrapping authentication and app_mention

## Interface
```
func New(botToken, appLevelToken string) (*Slackx, error)

func (*Slackx) SetAppMentionHook(hook func(*Slackx, *slackevents.AppMentionEvent, string) error) *Slackx
func (*Slackx) EventHandlerFunc() http.HandlerFunc
func (*Slackx) PostMessageToChannel(chanID, text string, escape bool) (string, string, error) // respChannel, respTimestamp, err
```
