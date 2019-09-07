package account

import (
	"backend/common/db/model"
	"github.com/haozzzzzzzz/go-rapid-development/web/ginbuilder"
)

// 账户目录
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

		// TODO

		ctx.SuccessReturn(respData)
		return
	},
}

// 添加账户
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

		// TODO

		ctx.SuccessReturn(respData)
		return
	},
}

// 更新账户
var AccountUpdate ginbuilder.HandleFunc = ginbuilder.HandleFunc{
	HttpMethod: "POST",
	RelativePaths: []string{
		"/api/api_hub/v1/account/account/update/:account_id",
	},
	Handle: func(ctx *ginbuilder.Context) (err error) {
		// request uri data
		type UriData struct {
			AccountId string `json:"account_id" uri:"account_id" binding:"required"`
		}
		uriData := UriData{}
		retCode, err := ctx.BindUriData(&uriData)
		if err != nil {
			ctx.Errorf(retCode, "verify  uri data failed. %s.", err)
			return
		}

		// request post data
		type PostData struct {
			Name string `json:"name" form:"name" binding:"required"`
		}
		postData := PostData{}
		retCode, err = ctx.BindPostData(&postData)
		if err != nil {
			ctx.Errorf(retCode, "verify  post data failed. %s.", err)
			return
		}

		// TODO

		ctx.Success()
		return
	},
}
