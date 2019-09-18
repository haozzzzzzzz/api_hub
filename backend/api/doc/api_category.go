package doc

import (
	"backend/common/db/model"
	"backend/common/db/table"
	"github.com/haozzzzzzzz/go-rapid-development/api/code"
	"github.com/haozzzzzzzz/go-rapid-development/web/ginbuilder"
	"time"
)

// 目录列表
// @api_doc_tags: 目录
var CategoryList ginbuilder.HandleFunc = ginbuilder.HandleFunc{
	HttpMethod: "GET",
	RelativePaths: []string{
		"/api/api_hub/v1/doc/category/list",
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
			ctx.Errorf(retCode, "verify category list query data failed. %s.", err)
			return
		}

		// response data
		type ResponseData struct {
			Count int64               `json:"count"`
			Items []*model.AhCategory `json:"items"`
		}
		respData := &ResponseData{
			Items: make([]*model.AhCategory, 0),
		}

		reqCtx := ctx.RequestCtx
		dbClient := table.NewHubDB(reqCtx)
		respData.Count, err = dbClient.AhCategoryCount()
		if nil != err {
			ctx.Errorf(code.CodeErrorDBQueryFailed.Clone(), "get category count failed. %s", err)
			return
		}

		respData.Items, err = dbClient.AhCategoryList(queryData.Page, queryData.Limit)
		if nil != err {
			ctx.Errorf(code.CodeErrorDBQueryFailed.Clone(), "get category list failed. %s", err)
			return
		}

		ctx.SuccessReturn(respData)
		return
	},
}

// 添加目录
// @api_doc_tags: 目录
var CategoryAdd ginbuilder.HandleFunc = ginbuilder.HandleFunc{
	HttpMethod: "POST",
	RelativePaths: []string{
		"/api/api_hub/v1/doc/category/add",
	},
	Handle: func(ctx *ginbuilder.Context) (err error) {
		// request post data
		type PostData struct {
			Name string `json:"name" form:"name"`
		}
		postData := PostData{}
		retCode, err := ctx.BindPostData(&postData)
		if err != nil {
			ctx.Errorf(retCode, "verify category add post data failed. %s.", err)
			return
		}

		// response data
		type ResponseData struct {
			NewCatId int64 `json:"new_cat_id"`
		}
		respData := &ResponseData{}

		reqCtx := ctx.RequestCtx
		dbClient := table.NewHubDB(reqCtx)
		now := time.Now().Unix()
		result, err := dbClient.AhCategoryAdd(&model.AhCategory{
			Name:       postData.Name,
			UpdateTime: now,
			CreateTime: now,
		})
		if nil != err {
			ctx.Errorf(code.CodeErrorDBInsertFailed.Clone(), "add category failed. %s", err)
			return
		}

		respData.NewCatId, err = result.LastInsertId()
		if nil != err {
			ctx.Errorf(code.CodeErrorDB.Clone(), "get last insert id failed. %s", err)
			return
		}

		ctx.SuccessReturn(respData)
		return
	},
}

/*
更新目录
@api_doc_tags: 目录
*/
var CategoryUpdate ginbuilder.HandleFunc = ginbuilder.HandleFunc{
	HttpMethod: "POST",
	RelativePaths: []string{
		"/api/api_hub/v1/doc/category/update/:cat_id",
	},
	Handle: func(ctx *ginbuilder.Context) (err error) {
		// request uri data
		type UriData struct {
			CatId uint32 `json:"cat_id" uri:"cat_id" binding:"required"`
		}
		uriData := UriData{}
		retCode, err := ctx.BindUriData(&uriData)
		if err != nil {
			ctx.Errorf(retCode, "verify category update uri data failed. %s.", err)
			return
		}

		// request post data
		type PostData struct {
			Name string `json:"name" form:"name" binding:"required"`
		}
		postData := PostData{}
		retCode, err = ctx.BindPostData(&postData)
		if err != nil {
			ctx.Errorf(retCode, "verify category update post data failed. %s.", err)
			return
		}

		reqCtx := ctx.RequestCtx
		dbClient := table.NewHubDB(reqCtx)
		err = dbClient.AhCategoryUpdate(uriData.CatId, postData.Name, time.Now().Unix())
		if nil != err {
			ctx.Errorf(code.CodeErrorDBUpdateFailed.Clone(), "update ah category failed. %s", err)
			return
		}

		ctx.Success()
		return
	},
}
