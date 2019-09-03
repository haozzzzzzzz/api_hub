package info

import (
	"github.com/haozzzzzzzz/go-rapid-development/web/ginbuilder"
)

var InfoSayHi ginbuilder.HandleFunc = ginbuilder.HandleFunc{
	HttpMethod: "GET",
	RelativePaths: []string{
		"/api/api_hub/info/reply/say_hi",
	},
	Handle: func(ctx *ginbuilder.Context) (err error) {
		ctx.String("hi")
		return
	},
}
