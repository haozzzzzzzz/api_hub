package doc

import (
	"backend/common/business"
	"backend/common/es/indices"
	"github.com/haozzzzzzzz/go-rapid-development/api/code"
	"github.com/haozzzzzzzz/go-rapid-development/web/ginbuilder"
)

var DocEsInitAhDoc ginbuilder.HandleFunc = ginbuilder.HandleFunc{
	HttpMethod: "POST",
	RelativePaths: []string{
		"/api/api_hub/v1/doc/es/init_ah_doc",
	},
	Handle: func(ctx *ginbuilder.Context) (err error) {
		reqCtx := ctx.RequestCtx
		esClient := indices.NewEsApiHub(reqCtx)
		exists, err := esClient.AhDocExists()
		if nil != err {
			ctx.Errorf(code.CodeErrorElasticsearch.Clone(), "get ah doc exists failed. %s", err)
			return
		}

		if exists {
			retCode := code.CodeErrorElasticsearch.Clone()
			retCode.Message = "ah_doc index exists"
			ctx.Errorf(retCode, "ah_doc index exists")
			return
		}

		err = esClient.AhDocCreateIndex()
		if nil != err {
			ctx.Errorf(code.CodeErrorElasticsearchCreate.Clone(), "create index failed. %s", err)
			return
		}

		ctx.Success()
		return
	},
}

var DocEsIndexAllAhDoc ginbuilder.HandleFunc = ginbuilder.HandleFunc{
	HttpMethod: "POST",
	RelativePaths: []string{
		"/api/api_hub/v1/doc/es/index_all_ah_doc",
	},
	Handle: func(ctx *ginbuilder.Context) (err error) {
		reqCtx := ctx.RequestCtx
		bsDoc := business.NewBsDoc(reqCtx)
		err = bsDoc.EsIndexAllAhDoc()
		if nil != err {
			ctx.Errorf(code.CodeErrorElasticsearch.Clone(), "index all ah_doc failed. %s", err)
			return
		}

		ctx.Success()
		return
	},
}
