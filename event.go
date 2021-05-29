package slackx

import (
	"encoding/json"
	"github.com/slack-go/slack/slackevents"
	"io/ioutil"
	"net/http"
	"strings"
)

//================================================================
//
//================================================================
const (
	DefTextHello          string = "hello"
	DefSlackMsgUnknownCmd string = "抱歉，我不了解您的意思。"
)

func (s *Slackx) EventHandlerFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if body, err := ioutil.ReadAll(r.Body); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		} else if event, err := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionNoVerifyToken()); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		} else {
			switch event.Type {
			case slackevents.URLVerification:
				if r, ok := event.Data.(*slackevents.EventsAPIURLVerificationEvent); !ok {
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				} else {
					withJSON(w, http.StatusOK, map[string]string{"challenge": r.Challenge})
				}
			case slackevents.CallbackEvent:
				s.innerEvent(w, event.InnerEvent)
			default:
				http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
			}
		}
	}
}

func (s *Slackx) innerEvent(w http.ResponseWriter, event slackevents.EventsAPIInnerEvent) {
	switch event.Type {
	case slackevents.AppMention:
		if ev, ok := event.Data.(*slackevents.AppMentionEvent); !ok {
			http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		} else {
			withJSON(w, http.StatusOK, nil)
			s.parser(ev)
		}
	// TODO: More cases to implement.
	default:
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
	}
}

//================================================================
//
//================================================================
func (s *Slackx) parser(ev *slackevents.AppMentionEvent) {
	me := "<@" + s.UserID + ">"
	if strings.HasPrefix(ev.Text, me) {
		cmd := strings.TrimSpace(strings.Replace(ev.Text, me, "", 1))
		switch {
		case strings.EqualFold(DefTextHello, cmd):
			s.PostMessageToChannel(ev.Channel, "Hello! <@"+ev.User+">", false)
		default:
			if s.AppMentionHook != nil {
				if err := s.AppMentionHook(s, ev, cmd); err != nil {
					s.PostMessageToChannel(ev.Channel, err.Error(), false)
				}
			} else {
				s.PostMessageToChannel(ev.Channel, DefSlackMsgUnknownCmd, false)
			}
		}
	}
}

//================================================================
//
//================================================================
func withJSON(w http.ResponseWriter, code int, data interface{}) {
	if js, err := json.Marshal(data); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		w.Write(js)
	}
}
