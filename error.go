package slackx

import (
	"errors"
)

const (
	DefMsgUnknownCmd      string = "抱歉，我不了解您的意思"
	DefMsgSvcUnavailable  string = "功能目前無法服務"
	DefMsgSvcErr          string = "功能目前發生異常"
	DefMsgSvcOnDeveloping string = "功能開發中"
)

var (
	ErrUnknownCmd      error = errors.New(DefMsgUnknownCmd)
	ErrSvcUnavailable  error = errors.New(DefMsgSvcUnavailable)
	ErrSvcErr          error = errors.New(DefMsgSvcErr)
	ErrSvcOnDeveloping error = errors.New(DefMsgSvcOnDeveloping)
)
