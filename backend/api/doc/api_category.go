package doc

import (
	"backend/common/db/model"
	"github.com/haozzzzzzzz/go-rapid-development/web/ginbuilder"
)

// 目录列表
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

		// TODO

		ctx.SuccessReturn(respData)
		return
	},
}

// 添加目录
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

		// TODO

		ctx.SuccessReturn(respData)
		return
	},
}
