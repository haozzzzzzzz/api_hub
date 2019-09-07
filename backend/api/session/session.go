package session

import (
	"github.com/haozzzzzzzz/go-rapid-development/web/ginbuilder"
)

const KeySkipSessionValidate = "skip_session_validate"

type Session struct {
	RequestData struct {
	} `json:"request_data"`

	Auth struct {
		AccountId uint32 `json:"account_id"`
	} `json:"auth"`
}

func (m *Session) BeforeHandle(ctx *ginbuilder.Context, method string, url string) {
	// ...
}

func (m *Session) AfterHandle(err error) {
	// ...
}

func (m *Session) Panic(errPanic interface{}) {
	// ...
}

func Builder(ctx *ginbuilder.Context) (err error) {
	ses := &Session{}
	ses.Auth.AccountId = 1 // TODO 目前写死一个用户

	ginCtx := ctx.GinContext
	skipValidate := ginCtx.GetBool(KeySkipSessionValidate)

	if !skipValidate { // validate session

	}

	ctx.Session = ses
	return
}

func GetSession(ctx *ginbuilder.Context) *Session {
	return ctx.Session.(*Session)
}

func init() {
	ginbuilder.BindSessionBuilder(Builder)
}
