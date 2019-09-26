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
	SpecTags       []string  `json:"spec_tags"`
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
		SpecTags:       []string{},
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

		pathMap := make(map[string]bool)
		tagMap := make(map[string]bool)
		for path, actionMap := range swaggerSpec.Paths {
			pathMap[path] = true
			for method, action := range actionMap {
				_ = method
				for _, tag := range action.Tags {
					tagMap[tag] = true
				}
			}
		}

		for path, _ := range pathMap {
			esDoc.SpecPaths = append(esDoc.SpecPaths, path)
		}

		for tag, _ := range tagMap {
			esDoc.SpecTags = append(esDoc.SpecTags, tag)
		}
	}

	return
}
