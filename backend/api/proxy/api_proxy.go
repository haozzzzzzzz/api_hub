package proxy

import (
	"backend/api/session"
	"backend/common/config"
	"fmt"
	"github.com/haozzzzzzzz/go-rapid-development/api/code"
	"github.com/haozzzzzzzz/go-rapid-development/web/ginbuilder"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

// 代理接口
// @api_doc_tags: 代理
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
			ctx.Errorf(retCode, "verify reverse proxy uri data failed. %s.", err)
			return
		}

		proxyUrlHost, ok := config.ReverseProxyConfig.ProxyTypeHost[uriData.ProxyType]
		if !ok {
			retCode = code.CodeErrorRequestFailed.Clone()
			retCode.Message = fmt.Sprintf("not supported reverse proxy type: %s", uriData.ProxyType)
			ctx.Errorf(retCode, retCode.Message)
			return
		}

		prefixUrl, err := url.Parse("http://" + proxyUrlHost)
		if nil != err {
			retCode = code.CodeErrorRequestFailed.Clone()
			retCode.Message = fmt.Sprintf("parse prefix url failed. prefix: %s, error: %s", prefixUrl, err)
			ctx.Errorf(retCode, retCode.Message)
			return
		}

		ginCtx := ctx.GinContext
		ginCtx.Request.URL.Path = uriData.TargetUri
		ginCtx.Request.Host = proxyUrlHost
		ginCtx.Request.RequestURI = ginCtx.Request.URL.RequestURI()
		reverseProxy := httputil.NewSingleHostReverseProxy(prefixUrl)
		reverseProxy.ModifyResponse = func(response *http.Response) (proxyRespError error) {
			// deal with as been blocked by CORS policy: The 'Access-Control-Allow-Origin' header contains multiple values '*, *', but only one is allowed
			upHeaders := response.Header
			for key, _ := range upHeaders {
				if strings.HasPrefix(key, "Access-Control-Allow") {
					upHeaders[key] = []string{}
				}
			}
			return
		}

		ctx.Logger.Infof("make reverse proxy request, url: %s", fmt.Sprintf("%s%s", prefixUrl, ginCtx.Request.RequestURI))
		// clear response header
		reverseProxy.ServeHTTP(ginCtx.Writer, ginCtx.Request)
		return
	},
}
