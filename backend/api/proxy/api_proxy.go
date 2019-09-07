package proxy

import (
	"backend/api/session"
	"backend/common/config"
	"fmt"
	"github.com/haozzzzzzzz/go-rapid-development/api/code"
	"github.com/haozzzzzzzz/go-rapid-development/web/ginbuilder"
	"net/http/httputil"
	"net/url"
)

var ReverseProxy ginbuilder.HandleFunc = ginbuilder.HandleFunc{
	HttpMethod: "ANY",
	RelativePaths: []string{
		"/api/api_hub/v1/reverse_proxy/:proxy_type/*target_uri",
	},
	BeforeHandle: func(ctx *ginbuilder.Context) (stop bool, err error) {
		ctx.GinContext.Set(session.KeySkipSessionValidate, true)
		return
	},
	Handle: func(ctx *ginbuilder.Context) (err error) {
		// request uri data
		type UriData struct {
			ProxyType string `json:"proxy_type" uri:"proxy_type" binding:"required"`
			TargetUri string `json:"target_uri" uri:"target_uri"`
		}
		uriData := UriData{}
		retCode, err := ctx.BindUriData(&uriData)
		if err != nil {
			ctx.Errorf(retCode, "verify  uri data failed. %s.", err)
			return
		}

		proxyUrlPrefix, ok := config.ReverseProxyConfig.ProxyTypeMap[uriData.ProxyType]
		if !ok {
			retCode = code.CodeErrorRequestFailed.Clone()
			retCode.Message = fmt.Sprintf("not supported reverse proxy type: %s", uriData.ProxyType)
			ctx.Errorf(retCode, retCode.Message)
			return
		}

		prefixUrl, err := url.Parse(proxyUrlPrefix)
		if nil != err {
			retCode = code.CodeErrorRequestFailed.Clone()
			retCode.Message = fmt.Sprintf("parse prefix url failed. prefix: %s, error: %s", prefixUrl, err)
			ctx.Errorf(retCode, retCode.Message)
			return
		}

		ginCtx := ctx.GinContext
		ginCtx.Request.URL.Path = uriData.TargetUri
		ginCtx.Request.RequestURI = ginCtx.Request.URL.RequestURI()
		reverseProxy := httputil.NewSingleHostReverseProxy(prefixUrl)
		ctx.Logger.Infof("make reverse proxy request, url: %s", fmt.Sprintf("%s%s", prefixUrl, ginCtx.Request.RequestURI))
		reverseProxy.ServeHTTP(ginCtx.Writer, ginCtx.Request)
		return
	},
}
