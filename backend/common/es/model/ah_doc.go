package model

import (
	"backend/common/db/model"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"time"
)

type EsAhDoc struct {
	DocId          uint32    `json:"doc_id"`
	Title          string    `json:"title"`
	SpecUrl        string    `json:"spec_url"`
	SpecPaths      []string  `json:"spec_paths"`
	AuthorId       uint32    `json:"author_id"`
	AuthorName     string    `json:"author_name"`
	CategoryId     uint32    `json:"category_id"`
	CategoryName   string    `json:"category_name"`
	CreateTime     time.Time `json:"create_time"`
	PostStatus     uint8     `json:"post_status"`
	CreateTimeUnix int64     `json:"create_time_unix"`
}

func NewEsAhDoc(
	doc *model.AhDocWithName,
) (esDoc *EsAhDoc, err error) {
	esDoc = &EsAhDoc{
		DocId:          doc.DocId,
		Title:          doc.Title,
		SpecUrl:        doc.SpecUrl,
		SpecPaths:      []string{},
		AuthorId:       doc.AuthorId,
		AuthorName:     doc.AuthorName,
		CategoryId:     doc.CategoryId,
		CategoryName:   doc.CategoryName,
		PostStatus:     doc.PostStatus,
		CreateTime:     time.Unix(doc.CreateTime, 0),
		CreateTimeUnix: doc.CreateTime,
	}

	if doc.SpecContent != "" {
		swaggerSpec := &model.SwaggerSpec{}
		err = json.Unmarshal([]byte(doc.SpecContent), swaggerSpec)
		if nil != err {
			logrus.Errorf("unmarshal spec content failed. error: %s.", err)
			return
		}

		for path, _ := range swaggerSpec.Paths {
			esDoc.SpecPaths = append(esDoc.SpecPaths, path)
		}

	}

	return
}
