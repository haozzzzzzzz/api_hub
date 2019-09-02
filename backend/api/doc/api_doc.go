package doc

import (
	"backend/api/session"
	"backend/common/business"
	"backend/common/db/model"
	"backend/common/db/table"
	"database/sql"
	"time"

	"github.com/haozzzzzzzz/go-rapid-development/api/code"

	"github.com/haozzzzzzzz/go-rapid-development/web/ginbuilder"
)

var DocList ginbuilder.HandleFunc = ginbuilder.HandleFunc{
	HttpMethod: "GET",
	RelativePaths: []string{
		"/api/api_hub/v1/doc/doc/list",
	},
	Handle: func(ctx *ginbuilder.Context) (err error) {
		// request query data
		type QueryData struct {
			Page  uint32 `json:"page" form:"page" binding:"required"`
			Limit uint8  `json:"limit" form:"limit" binding:"required"`
		}
		queryData := QueryData{}
		retCode, err := ctx.BindQueryData(&queryData)
		if err != nil {
			ctx.Errorf(retCode, "verify doc list query data failed. %s.", err)
			return
		}

		// response data
		type ResponseData struct {
			Count int64                   `json:"count"`
			Items []*business.DocListItem `json:"items"`
		}
		respData := &ResponseData{
			Items: make([]*business.DocListItem, 0),
		}

		reqCtx := ctx.RequestCtx

		dbClient := table.NewHubDB(reqCtx)
		respData.Count, err = dbClient.AhDocCount()
		if nil != err {
			ctx.Errorf(code.CodeErrorDBQueryFailed.Clone(), "query doc count failed. %s", err)
			return
		}

		bsDoc := business.NewBsDoc(reqCtx)
		respData.Items, err = bsDoc.DocList(queryData.Page, queryData.Limit)
		if nil != err {
			ctx.Errorf(code.CodeErrorServer.Clone(), "get doc list failed. %s", err)
			return
		}

		ctx.SuccessReturn(respData)
		return
	},
}

var DocAdd ginbuilder.HandleFunc = ginbuilder.HandleFunc{
	HttpMethod: "POST",
	RelativePaths: []string{
		"/api/api_hub/v1/doc/doc/add",
	},
	Handle: func(ctx *ginbuilder.Context) (err error) {
		ses := session.GetSession(ctx)

		// request post data
		type PostData struct {
			Title      string `json:"title" form:"title" binding:"required"`
			CategoryId uint32 `json:"category_id" form:"category_id" binding:"required"`
			SpecUrl    string `json:"spec_url" form:"spec_url" binding:"required"`
		}
		postData := PostData{}
		retCode, err := ctx.BindPostData(&postData)
		if err != nil {
			ctx.Errorf(retCode, "verify add doc post data failed. %s.", err)
			return
		}

		// response data
		type ResponseData struct {
			NewDocId int64 `json:"new_doc_id"`
		}
		respData := &ResponseData{}

		reqCtx := ctx.RequestCtx

		bsDoc := business.NewBsDoc(reqCtx)
		now := time.Now().Unix()
		respData.NewDocId, err = bsDoc.DocAdd(&model.AhDoc{
			Title:      postData.Title,
			SpecUrl:    postData.SpecUrl,
			CategoryId: postData.CategoryId,
			AuthorId:   ses.Auth.AccountId,
			PostStatus: model.PostStatusPublished,
			UpdateTime: now,
			CreateTime: now,
		})
		if nil != err {
			ctx.Errorf(code.CodeErrorServer.Clone(), "bs doc add doc failed. %s", err)
			return
		}

		ctx.SuccessReturn(respData)
		return
	},
}

var CheckAndAddDoc ginbuilder.HandleFunc = ginbuilder.HandleFunc{
	HttpMethod: "POST",
	RelativePaths: []string{
		"/api/api_hub/v1/doc/doc/check_add",
	},
	Handle: func(ctx *ginbuilder.Context) (err error) {
		ses := session.GetSession(ctx)

		// request post data
		type PostData struct {
			Title   string `json:"title" form:"title" binding:"required"`
			SpecUrl string `json:"spec_url" form:"spec_url" binding:"required"`
		}
		postData := PostData{}
		retCode, err := ctx.BindPostData(&postData)
		if err != nil {
			ctx.Errorf(retCode, "verify check and add doc post data failed. %s.", err)
			return
		}

		// response data
		type ResponseData struct {
			DocId int64 `json:"doc_id"`
		}
		respData := &ResponseData{}

		reqCtx := ctx.RequestCtx

		dbClient := table.NewHubDB(reqCtx)
		doc, err := dbClient.AhDocGetByTitleSpecUrl(postData.Title, postData.SpecUrl)
		if nil != err && err != sql.ErrNoRows {
			ctx.Errorf(code.CodeErrorDBQueryFailed.Clone(), "get ah_doc by title and spec_url failed. %s", err)
			return
		}

		if err == nil && doc != nil {
			respData.DocId = int64(doc.DocId)
			ctx.SuccessReturn(respData)
			return
		}
		err = nil

		bsDoc := business.NewBsDoc(reqCtx)
		now := time.Now().Unix()
		respData.DocId, err = bsDoc.DocAdd(&model.AhDoc{
			Title:      postData.Title,
			SpecUrl:    postData.SpecUrl,
			CategoryId: model.DefaultCategoryId,
			AuthorId:   ses.Auth.AccountId,
			PostStatus: model.PostStatusPublished,
			UpdateTime: now,
			CreateTime: now,
		})
		if nil != err {
			ctx.Errorf(code.CodeErrorServer.Clone(), "bs doc add doc failed. %s", err)
			return
		}

		ctx.SuccessReturn(respData)
		return
	},
}
