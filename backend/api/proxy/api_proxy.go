package proxy

import (
	"github.com/gin-gonic/gin"
	"github.com/haozzzzzzzz/go-rapid-development/api/code"
	"github.com/haozzzzzzzz/go-rapid-development/http"
	"github.com/haozzzzzzzz/go-rapid-development/web/ginbuilder"
)

var ProxyGet ginbuilder.HandleFunc = ginbuilder.HandleFunc{
	HttpMethod: "GET",
	RelativePaths: []string{
		"/api/api_hub/v1/request/proxy/get",
	},
	Handle: func(ctx *ginbuilder.Context) (err error) {
		// request query data
		type QueryData struct {
			Source string `json:"source" form:"source" binding:"required"`
		}
		queryData := QueryData{}
		retCode, err := ctx.BindQueryData(&queryData)
		if err != nil {
			ctx.Errorf(retCode, "verify  query data failed. %s.", err)
			return
		}

		reqCtx := ctx.RequestCtx

		req, err := http.NewRequest(queryData.Source, reqCtx, http.LongTimeoutRequestClient)
		if nil != err {
			retCode = code.CodeErrorServer.Clone()
			retCode.Message = "proxy new request failed"
			ctx.Errorf(retCode, "proxy new request failed. %s", err)
			return
		}

		text, err := req.GetText()
		if nil != err {
			retCode = code.CodeErrorServer.Clone()
			retCode.Message = "proxy get text failed"
			ctx.Errorf(retCode, "proxy get text failed", err)
			return
		}

		ctx.GinContext.Header("Content-Type", gin.MIMEJSON)
		ctx.String(text)
		return
	},
}
