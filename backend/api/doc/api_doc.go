package doc

import (
	"backend/api/session"
	"backend/common/business"
	"backend/common/db/model"
	"backend/common/db/table"
	"database/sql"
	"github.com/gin-gonic/gin"
	"strings"
	"time"

	"github.com/haozzzzzzzz/go-rapid-development/api/code"

	"github.com/haozzzzzzzz/go-rapid-development/web/ginbuilder"
)

// 文档列表
// @api_doc_tags: 文档
var DocList ginbuilder.HandleFunc = ginbuilder.HandleFunc{
	HttpMethod: "GET",
	RelativePaths: []string{
		"/api/api_hub/v1/doc/doc/list",
	},
	Handle: func(ctx *ginbuilder.Context) (err error) {
		// request query data
		type QueryData struct {
			Page   uint32 `json:"page" form:"page" binding:"required"`   // 分页ID
			Limit  uint8  `json:"limit" form:"limit" binding:"required"` // 分页大小
			Search string `json:"search" form:"search"`
		}
		queryData := QueryData{}
		retCode, err := ctx.BindQueryData(&queryData)
		if err != nil {
			ctx.Errorf(retCode, "verify doc list query data failed. %s.", err)
			return
		}

		queryData.Search = strings.TrimSpace(queryData.Search)

		// response data
		type ResponseData struct {
			Count int64                   `json:"count"` // 记录数目
			Items []*business.DocListItem `json:"items"` // 文档记录
		}
		respData := &ResponseData{
			Items: make([]*business.DocListItem, 0),
		}

		reqCtx := ctx.RequestCtx

		dbClient := table.NewHubDB(reqCtx)
		respData.Count, err = dbClient.AhDocCount(queryData.Search)
		if nil != err {
			ctx.Errorf(code.CodeErrorDBQueryFailed.Clone(), "query doc count failed. %s", err)
			return
		}

		bsDoc := business.NewBsDoc(reqCtx)
		respData.Items, err = bsDoc.DocList(queryData.Page, queryData.Limit, queryData.Search)
		if nil != err {
			ctx.Errorf(code.CodeErrorServer.Clone(), "get doc list failed. %s", err)
			return
		}

		ctx.SuccessReturn(respData)
		return
	},
}

// 添加文档
// @api_doc_tags: 文档
var DocAdd ginbuilder.HandleFunc = ginbuilder.HandleFunc{
	HttpMethod: "POST",
	RelativePaths: []string{
		"/api/api_hub/v1/doc/doc/add",
	},
	Handle: func(ctx *ginbuilder.Context) (err error) {
		ses := session.GetSession(ctx)

		// request post data
		type PostData struct {
			Title       string `json:"title" form:"title" binding:"required"`             // 文档标题
			CategoryId  uint32 `json:"category_id" form:"category_id" binding:"required"` // 分类ID
			SpecUrl     string `json:"spec_url" form:"spec_url"`                          // swagger.json url
			SpecContent string `json:"spec_content" form:"spec_content"`                  // swagger content
		}
		postData := PostData{}
		retCode, err := ctx.BindPostData(&postData)
		if err != nil {
			ctx.Errorf(retCode, "verify add doc post data failed. %s.", err)
			return
		}

		if postData.SpecUrl == "" && postData.SpecContent == "" {
			retCode.Message = "require spec"
			ctx.Errorf(retCode, "spec is empty")
			return
		}

		// response data
		type ResponseData struct {
			NewDocId int64 `json:"new_doc_id"` // 新建的文档ID
		}
		respData := &ResponseData{}

		reqCtx := ctx.RequestCtx

		bsDoc := business.NewBsDoc(reqCtx)
		now := time.Now().Unix()
		respData.NewDocId, err = bsDoc.DocAdd(&model.AhDoc{
			Title:       postData.Title,
			SpecUrl:     postData.SpecUrl,
			SpecContent: postData.SpecContent,
			CategoryId:  postData.CategoryId,
			AuthorId:    ses.Auth.AccountId,
			PostStatus:  model.PostStatusPublished,
			UpdateTime:  now,
			CreateTime:  now,
		})
		if nil != err {
			ctx.Errorf(code.CodeErrorServer.Clone(), "bs doc add doc failed. %s", err)
			return
		}

		ctx.SuccessReturn(respData)
		return
	},
}

/*
更新文档
@api_doc_tags: 文档
更新文档
*/
var DocUpdate ginbuilder.HandleFunc = ginbuilder.HandleFunc{
	HttpMethod: "POST",
	RelativePaths: []string{
		"/api/api_hub/v1/doc/doc/update/:doc_id",
	},
	Handle: func(ctx *ginbuilder.Context) (err error) {
		ses := session.GetSession(ctx)

		// request uri data
		type UriData struct {
			DocId uint32 `json:"doc_id" uri:"doc_id" binding:"required"`
		}
		uriData := UriData{}
		retCode, err := ctx.BindUriData(&uriData)
		if err != nil {
			ctx.Errorf(retCode, "verify  uri data failed. %s.", err)
			return
		}

		// request post data
		type PostData struct {
			Title       string `json:"title" form:"title" binding:"required"`             // 文档标题
			CategoryId  uint32 `json:"category_id" form:"category_id" binding:"required"` // 分类ID
			SpecUrl     string `json:"spec_url" form:"spec_url"`                          // swagger.json url
			SpecContent string `json:"spec_content" form:"spec_content"`                  // swagger content
		}
		postData := PostData{}
		retCode, err = ctx.BindPostData(&postData)
		if err != nil {
			ctx.Errorf(retCode, "verify  post data failed. %s.", err)
			return
		}

		if postData.SpecUrl == "" && postData.SpecContent == "" {
			retCode.Message = "require spec"
			ctx.Errorf(retCode, "spec is empty")
			return
		}

		reqCtx := ctx.RequestCtx
		bsDoc := business.NewBsDoc(reqCtx)
		err = bsDoc.DocUpdate(
			uriData.DocId,
			postData.Title,
			postData.CategoryId,
			ses.Auth.AccountId,
			postData.SpecUrl,
			postData.SpecContent,
			time.Now().Unix(),
		)
		if nil != err {
			ctx.Errorf(code.CodeErrorDBUpdateFailed.Clone(), "update doc failed. %s", err)
			return
		}

		ctx.Success()
		return
	},
}

/*
更改目录
@api_doc_tags: 文档
*/
var DocChangeCategory ginbuilder.HandleFunc = ginbuilder.HandleFunc{
	HttpMethod: "POST",
	RelativePaths: []string{
		"/api/api_hub/v1/doc/doc/change_category/:doc_id",
	},
	Handle: func(ctx *ginbuilder.Context) (err error) {
		// request uri data
		type UriData struct {
			DocId uint32 `json:"doc_id" uri:"doc_id" binding:"required"`
		}
		uriData := UriData{}
		retCode, err := ctx.BindUriData(&uriData)
		if err != nil {
			ctx.Errorf(retCode, "verify change category uri data failed. %s.", err)
			return
		}

		// request post data
		type PostData struct {
			CategoryId uint32 `json:"category_id" form:"category_id" binding:"required"`
		}
		postData := PostData{}
		retCode, err = ctx.BindPostData(&postData)
		if err != nil {
			ctx.Errorf(retCode, "verify  post data failed. %s.", err)
			return
		}

		reqCtx := ctx.RequestCtx

		dbClient := table.NewHubDB(reqCtx)
		doc, err := dbClient.AhDocGet(uriData.DocId)
		if nil != err {
			ctx.Errorf(code.CodeErrorDBQueryFailed.Clone(), "get doc failed. %s", err)
			return
		}

		if doc.CategoryId == postData.CategoryId {
			ctx.Success()
			return
		}

		bsDoc := business.NewBsDoc(reqCtx)
		err = bsDoc.DocUpdate(doc.DocId, doc.Title, postData.CategoryId, doc.AuthorId, doc.SpecUrl, doc.SpecContent, time.Now().Unix())
		if nil != err {
			ctx.Errorf(code.CodeErrorDBUpdateFailed.Clone(), "update doc category failed. %s", err)
			return
		}

		ctx.Success()
		return
	},
}

/*
更改发布者
@api_doc_tags: 文档
*/
var DocChangeAuthor ginbuilder.HandleFunc = ginbuilder.HandleFunc{
	HttpMethod: "POST",
	RelativePaths: []string{
		"/api/api_hub/v1/doc/doc/change_author/:doc_id",
	},
	Handle: func(ctx *ginbuilder.Context) (err error) {
		// request uri data
		type UriData struct {
			DocId uint32 `json:"doc_id" uri:"doc_id" binding:"required"`
		}
		uriData := UriData{}
		retCode, err := ctx.BindUriData(&uriData)
		if err != nil {
			ctx.Errorf(retCode, "verify change author uri data failed. %s.", err)
			return
		}

		// request post data
		type PostData struct {
			AuthorId uint32 `json:"author_id" form:"author_id" binding:"required"`
		}
		postData := PostData{}
		retCode, err = ctx.BindPostData(&postData)
		if err != nil {
			ctx.Errorf(retCode, "verify change author post data failed. %s.", err)
			return
		}

		reqCtx := ctx.RequestCtx
		dbClient := table.NewHubDB(reqCtx)
		doc, err := dbClient.AhDocGet(uriData.DocId)
		if nil != err {
			ctx.Errorf(code.CodeErrorDBQueryFailed.Clone(), "get doc failed. %s", err)
			return
		}

		if doc.AuthorId == postData.AuthorId {
			ctx.Success()
			return
		}

		bsDoc := business.NewBsDoc(reqCtx)
		err = bsDoc.DocUpdate(doc.DocId, doc.Title, doc.CategoryId, postData.AuthorId, doc.SpecUrl, doc.SpecContent, time.Now().Unix())
		if nil != err {
			ctx.Errorf(code.CodeErrorDBUpdateFailed.Clone(), "update doc category failed. %s", err)
			return
		}

		ctx.Success()
		return
	},
}

// 删除文档
// @api_doc_tags: 文档
var DocDelele ginbuilder.HandleFunc = ginbuilder.HandleFunc{
	HttpMethod: "POST",
	RelativePaths: []string{
		"/api/api_hub/v1/doc/doc/delete/:doc_id",
	},
	Handle: func(ctx *ginbuilder.Context) (err error) {
		ses := session.GetSession(ctx)

		// request uri data
		type UriData struct {
			DocId uint32 `json:"doc_id" uri:"doc_id" binding:"required"` // 文档ID
		}
		uriData := UriData{}
		retCode, err := ctx.BindUriData(&uriData)
		if err != nil {
			ctx.Errorf(retCode, "verify delete doc uri data failed. %s.", err)
			return
		}

		reqCtx := ctx.RequestCtx
		bsDoc := business.NewBsDoc(reqCtx)
		err = bsDoc.DocDel(uriData.DocId)
		if nil != err {
			ctx.Errorf(code.CodeErrorDBDeleteFailed.Clone(), "delete doc failed. %s", err)
			return
		}

		ctx.Logger.Warnf("delete doc. doc_id: %d, %#v", uriData.DocId, ses.Auth)

		ctx.Success()
		return
	},
}

// 检查并且创建文档
// @api_doc_tags: 文档
// 用在swagger.sh上
var DocCheckPost ginbuilder.HandleFunc = ginbuilder.HandleFunc{
	HttpMethod: "POST",
	RelativePaths: []string{
		"/api/api_hub/v1/doc/doc/check_post",
	},
	Handle: func(ctx *ginbuilder.Context) (err error) {
		ses := session.GetSession(ctx)

		// request post data
		type PostData struct {
			Title       string `json:"title" form:"title" binding:"required"`             // 标题
			CategoryId  uint32 `json:"category_id" form:"category_id" binding:"required"` // 分类ID
			SpecUrl     string `json:"spec_url" form:"spec_url"`                          // swagger.json url
			SpecContent string `json:"spec_content" form:"spec_content"`
		}
		postData := PostData{}
		retCode, err := ctx.BindPostData(&postData)
		if err != nil {
			ctx.Errorf(retCode, "verify check and add doc post data failed. %s.", err)
			return
		}

		if postData.SpecUrl == "" && postData.SpecContent == "" {
			retCode.Message = "require spec"
			ctx.Errorf(retCode, "spec is empty")
			return
		}

		// response data
		type ResponseData struct {
			DocId int64 `json:"doc_id"` // 文档ID
		}
		respData := &ResponseData{}

		reqCtx := ctx.RequestCtx

		dbClient := table.NewHubDB(reqCtx)
		doc, err := dbClient.AhDocGetByTitle(postData.Title)
		if nil != err && err != sql.ErrNoRows {
			ctx.Errorf(code.CodeErrorDBQueryFailed.Clone(), "get ah_doc by title and spec_url failed. %s", err)
			return
		}

		bsDoc := business.NewBsDoc(reqCtx)
		now := time.Now().Unix()

		if err == nil && doc != nil { // update
			respData.DocId = int64(doc.DocId)

			// update
			err = bsDoc.DocUpdate(
				doc.DocId,
				doc.Title,
				postData.CategoryId,
				ses.Auth.AccountId,
				postData.SpecUrl,
				postData.SpecContent,
				now,
			)
			if nil != err {
				ctx.Errorf(code.CodeErrorDBUpdateFailed.Clone(), "update doc failed. %s", err)
				return
			}

			ctx.SuccessReturn(respData)
			return
		}
		err = nil

		respData.DocId, err = bsDoc.DocAdd(&model.AhDoc{
			Title:       postData.Title,
			SpecUrl:     postData.SpecUrl,
			SpecContent: postData.SpecContent,
			CategoryId:  postData.CategoryId,
			AuthorId:    ses.Auth.AccountId,
			PostStatus:  model.PostStatusPublished,
			UpdateTime:  now,
			CreateTime:  now,
		})
		if nil != err {
			ctx.Errorf(code.CodeErrorServer.Clone(), "bs doc add doc failed. %s", err)
			return
		}

		ctx.SuccessReturn(respData)
		return
	},
}

// 查看文档的swagger.json
// @api_doc_tags: 文档
var DocDetailSpec ginbuilder.HandleFunc = ginbuilder.HandleFunc{
	HttpMethod: "GET",
	RelativePaths: []string{
		"/api/api_hub/v1/doc/detail/spec/:doc_id",
	},
	Handle: func(ctx *ginbuilder.Context) (err error) {
		// request path data
		type PathData struct {
			DocId uint32 `json:"doc_id" uri:"doc_id" binding:"required"` // 文档ID
		}
		pathData := PathData{}
		retCode, err := ctx.BindUriData(&pathData)
		if err != nil {
			ctx.Errorf(retCode, "verify doc spec path data failed. %s.", err)
			return
		}

		reqCtx := ctx.RequestCtx
		dbClient := table.NewHubDB(reqCtx)
		doc, err := dbClient.AhDocGet(pathData.DocId)
		if nil != err {
			ctx.Errorf(code.CodeErrorDBQueryFailed.Clone(), "get doc failed. %s", err)
			return
		}

		if doc.SpecContent != "" {
			ctx.GinContext.Header("Content-Type", gin.MIMEJSON)
			ctx.String(doc.SpecContent)

		} else if doc.SpecUrl != "" {
			ctx.TemporaryRedirect(doc.SpecUrl)

		}

		return
	},
}
