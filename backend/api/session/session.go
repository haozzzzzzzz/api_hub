package session

import (
	"backend/common/db/model"
	"github.com/gosexy/to"
	"github.com/haozzzzzzzz/go-rapid-development/web/ginbuilder"
)

const KeySkipSessionValidate = "skip_session_validate"

type Session struct {
	RequestData struct {
		HeaderAccountId uint32 `json:"header_account_id"`
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

	ginCtx := ctx.GinContext
	skipValidate := ginCtx.GetBool(KeySkipSessionValidate)

	reqHeader := ginCtx.Request.Header
	ses.RequestData.HeaderAccountId = uint32(to.Uint64(reqHeader.Get("X-Ah-Account-Id")))
	ses.Auth.AccountId = ses.RequestData.HeaderAccountId // 临时
	if ses.Auth.AccountId == 0 {
		ses.Auth.AccountId = model.DefaultAccountId
	}

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
