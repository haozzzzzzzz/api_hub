package api

import (
	account "backend/api/account"
	doc "backend/api/doc"
	info "backend/api/info"
	proxy "backend/api/proxy"
	"github.com/gin-gonic/gin"
)

// 注意：BindRouters函数体内不能自定义添加任何声明，由api compile命令生成api绑定声明
func BindRouters(engine *gin.Engine) (err error) {
	engine.Handle("GET", "/api/api_hub/info/reply/say_hi", info.InfoSayHi.GinHandler)
	engine.Handle("POST", "/api/api_hub/v1/account/account/add", account.AccountAdd.GinHandler)
	engine.Handle("GET", "/api/api_hub/v1/account/account/list", account.AccountList.GinHandler)
	engine.Handle("POST", "/api/api_hub/v1/account/account/update/:account_id", account.AccountUpdate.GinHandler)
	engine.Handle("POST", "/api/api_hub/v1/doc/category/add", doc.CategoryAdd.GinHandler)
	engine.Handle("GET", "/api/api_hub/v1/doc/category/list", doc.CategoryList.GinHandler)
	engine.Handle("POST", "/api/api_hub/v1/doc/category/update/:cat_id", doc.CategoryUpdate.GinHandler)
	engine.Handle("GET", "/api/api_hub/v1/doc/detail/spec/:doc_id", doc.DocDetailSpec.GinHandler)
	engine.Handle("POST", "/api/api_hub/v1/doc/doc/add", doc.DocAdd.GinHandler)
	engine.Handle("POST", "/api/api_hub/v1/doc/doc/change_author/:doc_id", doc.DocChangeAuthor.GinHandler)
	engine.Handle("POST", "/api/api_hub/v1/doc/doc/change_category/:doc_id", doc.DocChangeCategory.GinHandler)
	engine.Handle("POST", "/api/api_hub/v1/doc/doc/change_title/:doc_id", doc.DocChangeTitle.GinHandler)
	engine.Handle("POST", "/api/api_hub/v1/doc/doc/check_post", doc.DocCheckPost.GinHandler)
	engine.Handle("POST", "/api/api_hub/v1/doc/doc/delete/:doc_id", doc.DocDelele.GinHandler)
	engine.Handle("GET", "/api/api_hub/v1/doc/doc/get/:doc_id", doc.DocGet.GinHandler)
	engine.Handle("GET", "/api/api_hub/v1/doc/doc/list", doc.DocList.GinHandler)
	engine.Handle("GET", "/api/api_hub/v1/doc/doc/summary/:doc_id", doc.DocGetSummary.GinHandler)
	engine.Handle("POST", "/api/api_hub/v1/doc/doc/update/:doc_id", doc.DocUpdate.GinHandler)
	engine.Handle("POST", "/api/api_hub/v1/doc/es/index_all_ah_doc", doc.DocEsIndexAllAhDoc.GinHandler)
	engine.Handle("POST", "/api/api_hub/v1/doc/es/init_ah_doc", doc.DocEsInitAhDoc.GinHandler)
	engine.Any("/api/api_hub/v1/reverse_proxy/:proxy_type/*target_uri", proxy.ReverseProxy.GinHandler)
	return
}
