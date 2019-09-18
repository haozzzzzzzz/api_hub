package account

import (
	"backend/common/db/model"
	"backend/common/db/table"
	"github.com/haozzzzzzzz/go-rapid-development/api/code"
	"github.com/haozzzzzzzz/go-rapid-development/web/ginbuilder"
	"time"
)

// 账户目录
// @api_doc_tags: 账号
var AccountList ginbuilder.HandleFunc = ginbuilder.HandleFunc{
	HttpMethod: "GET",
	RelativePaths: []string{
		"/api/api_hub/v1/account/account/list",
	},
	Handle: func(ctx *ginbuilder.Context) (err error) {
		// request query data
		type QueryData struct {
			Page  uint32 `json:"page" form:"page" binding:"required"`   // 分页ID
			Limit uint8  `json:"limit" form:"limit" binding:"required"` // 分页大小
		}
		queryData := QueryData{}
		retCode, err := ctx.BindQueryData(&queryData)
		if err != nil {
			ctx.Errorf(retCode, "verify account list query data failed. %s.", err)
			return
		}

		// response data
		type ResponseData struct {
			Count int64              `json:"count"`
			Items []*model.AhAccount `json:"items"`
		}
		respData := &ResponseData{
			Items: make([]*model.AhAccount, 0),
		}

		reqCtx := ctx.RequestCtx
		dbClient := table.NewHubDB(reqCtx)
		respData.Count, err = dbClient.AhAccountCount()
		if nil != err {
			ctx.Errorf(code.CodeErrorDBQueryFailed.Clone(), "get account count failed. %s", err)
			return
		}

		respData.Items, err = dbClient.AhAccountList(queryData.Page, queryData.Limit)
		if nil != err {
			ctx.Errorf(code.CodeErrorDBQueryFailed.Clone(), "get account list failed. %s", err)
			return
		}

		ctx.SuccessReturn(respData)
		return
	},
}

// 添加账户
// @api_doc_tags: 账号
var AccountAdd ginbuilder.HandleFunc = ginbuilder.HandleFunc{
	HttpMethod: "POST",
	RelativePaths: []string{
		"/api/api_hub/v1/account/account/add",
	},
	Handle: func(ctx *ginbuilder.Context) (err error) {
		// request post data
		type PostData struct {
			Name string `json:"name" form:"name"`
		}
		postData := PostData{}
		retCode, err := ctx.BindPostData(&postData)
		if err != nil {
			ctx.Errorf(retCode, "verify account add post data failed. %s.", err)
			return
		}

		// response data
		type ResponseData struct {
			NewAccountId int64 `json:"new_account_id"`
		}
		respData := &ResponseData{}

		reqCtx := ctx.RequestCtx
		dbClient := table.NewHubDB(reqCtx)
		now := time.Now().Unix()
		result, err := dbClient.AhAccountAdd(&model.AhAccount{
			Name:       postData.Name,
			UpdateTime: now,
			CreateTime: now,
		})
		if nil != err {
			ctx.Errorf(code.CodeErrorDBInsertFailed.Clone(), "add ah account failed. %s", err)
			return
		}

		respData.NewAccountId, err = result.LastInsertId()
		if nil != err {
			ctx.Errorf(code.CodeErrorDB.Clone(), "get last insert id failed. %s", err)
			return
		}

		ctx.SuccessReturn(respData)
		return
	},
}

// 更新账户
// @api_doc_tags: 账号
var AccountUpdate ginbuilder.HandleFunc = ginbuilder.HandleFunc{
	HttpMethod: "POST",
	RelativePaths: []string{
		"/api/api_hub/v1/account/account/update/:account_id",
	},
	Handle: func(ctx *ginbuilder.Context) (err error) {
		// request uri data
		type UriData struct {
			AccountId uint32 `json:"account_id" uri:"account_id" binding:"required"`
		}
		uriData := UriData{}
		retCode, err := ctx.BindUriData(&uriData)
		if err != nil {
			ctx.Errorf(retCode, "verify update account uri data failed. %s.", err)
			return
		}

		// request post data
		type PostData struct {
			Name string `json:"name" form:"name" binding:"required"`
		}
		postData := PostData{}
		retCode, err = ctx.BindPostData(&postData)
		if err != nil {
			ctx.Errorf(retCode, "verify update account post data failed. %s.", err)
			return
		}

		reqCtx := ctx.RequestCtx
		dbClient := table.NewHubDB(reqCtx)
		err = dbClient.AhAccountUpdate(uriData.AccountId, postData.Name, time.Now().Unix())
		if nil != err {
			ctx.Errorf(code.CodeErrorDBUpdateFailed.Clone(), "update account failed. %s", err)
			return
		}

		ctx.Success()
		return
	},
}
