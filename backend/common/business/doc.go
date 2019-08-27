package business

import (
	"backend/common/db/model"
	"backend/common/db/table"
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

type BsDoc struct {
	Ctx context.Context
}

func NewBsDoc(ctx context.Context) *BsDoc {
	return &BsDoc{Ctx: ctx}
}

type DocListItem struct {
	DocId        uint32 `json:"doc_id"`
	Title        string `json:"title"`
	CategoryName string `json:"category_name"`
	AuthorName   string `json:"author_name"`
	SpecUrl      string `json:"spec_url"`
	PostStatus   uint8  `json:"post_status"`
	CreateTime   int64  `json:"create_time"`
}

func (m *BsDoc) DocList(
	page uint32,
	limit uint8,
) (items []*DocListItem, err error) {
	items = make([]*DocListItem, 0)

	dbClient := table.NewHubDB(m.Ctx)
	docs, err := dbClient.AhDocList(page, limit)
	if nil != err {
		logrus.Errorf("get doc list items failed. error: %s.", err)
		return
	}

	for _, doc := range docs {
		item := &DocListItem{
			DocId:      doc.DocId,
			Title:      doc.Title,
			SpecUrl:    doc.SpecUrl,
			PostStatus: doc.PostStatus,
			CreateTime: doc.CreateTime,
		}

		// category
		cat, errGet := dbClient.AhCategoryGet(doc.CategoryId)
		err = errGet
		if nil != err || cat == nil {
			logrus.Errorf("get ah_category failed. cat_id: %d, error: %s.", doc.CategoryId, err)
			return
		}

		item.CategoryName = cat.Name

		// author
		author, errGet := dbClient.AhAccountGet(doc.AuthorId)
		err = errGet
		if nil != err || author == nil {
			logrus.Errorf("get ah_account failed. acc_id: %d, error: %s.", doc.AuthorId, err)
			return
		}

		item.AuthorName = author.Name
		items = append(items, item)
	}

	return
}

func (m *BsDoc) DocAdd(doc *model.AhDoc) (newDocId int64, err error) {
	dbClient := table.NewHubDB(m.Ctx)
	_, err = dbClient.AhCategoryGet(doc.CategoryId)
	if nil != err {
		logrus.Errorf("get category failed. %s", err)
		return
	}

	tx, err := dbClient.DB.BeginTxx(m.Ctx, nil)
	if nil != err {
		logrus.Errorf("add doc failed. error: %s.", err)
		return
	}
	defer func() {
		if nil != err {
			errRoll := tx.Rollback()
			if nil != errRoll {
				logrus.Errorf("rollback add doc tx failed. error: %s.", err)
				return
			}
			return
		} else {
			err = tx.Commit()
			if nil != err {
				logrus.Errorf("commit add doc tx failed. error: %s.", err)
				return
			}
		}
	}()

	newDocId, err = dbClient.AhDocAddTx(tx, doc)
	if nil != err {
		logrus.Errorf("add doc failed. error: %s.", err)
		return
	}

	// 更新数目
	err = dbClient.AhCategoryDocNumIncrTx(tx, doc.CategoryId, 1, time.Now().Unix())
	if nil != err {
		logrus.Errorf("incr ah_category doc_num failed. error: %s.", err)
		return
	}

	return
}
