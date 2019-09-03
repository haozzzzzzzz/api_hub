package api

import (
	doc "backend/api/doc"
	proxy "backend/api/proxy"
	"github.com/gin-gonic/gin"
)

// 注意：BindRouters函数体内不能自定义添加任何声明，由api compile命令生成api绑定声明
func BindRouters(engine *gin.Engine) (err error) {
	engine.Handle("POST", "/api/api_hub/v1/doc/doc/add", doc.DocAdd.GinHandler)
	engine.Handle("POST", "/api/api_hub/v1/doc/doc/check_add", doc.CheckAndAddDoc.GinHandler)
	engine.Handle("GET", "/api/api_hub/v1/doc/doc/list", doc.DocList.GinHandler)
	engine.Handle("GET", "/api/api_hub/v1/request/proxy/get", proxy.ProxyGet.GinHandler)
	return
}
