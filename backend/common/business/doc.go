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
	DocId        uint32 `json:"doc_id"`        // 文档ID
	Title        string `json:"title"`         // 文档标题
	CategoryName string `json:"category_name"` // 分类名称
	AuthorName   string `json:"author_name"`   // 作者名称
	SpecUrl      string `json:"spec_url"`      // swagger.json url
	PostStatus   uint8  `json:"post_status"`   // 状态
	CreateTime   int64  `json:"create_time"`   // 创建时间
}

func (m *BsDoc) DocList(
	page uint32,
	limit uint8,
	search string,
) (items []*DocListItem, err error) {
	items = make([]*DocListItem, 0)

	dbClient := table.NewHubDB(m.Ctx)
	docs, err := dbClient.AhDocList(page, limit, search)
	if nil != err {
		logrus.Errorf("get doc list items failed. error: %s.", err)
		return
	}

	mCatIds := make(map[uint32]bool, 0)
	mAccIds := make(map[uint32]bool, 0)
	for _, doc := range docs {
		mCatIds[doc.CategoryId] = true
		mAccIds[doc.AuthorId] = true
	}

	catIds := make([]uint32, 0)
	accIds := make([]uint32, 0)
	for catId, _ := range mCatIds {
		catIds = append(catIds, catId)
	}

	for accId, _ := range mAccIds {
		accIds = append(accIds, accId)
	}

	mCategory, err := dbClient.AhCategoryBatch(catIds)
	if nil != err {
		logrus.Errorf("batch get ah category failed. error: %s.", err)
		return
	}

	mAuthor, err := dbClient.AhAccountBatch(accIds)
	if nil != err {
		logrus.Errorf("batch get ah account failed. error: %s.", err)
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
		cat, ok := mCategory[doc.CategoryId]
		if ok {
			item.CategoryName = cat.Name
		}

		// author
		author, ok := mAuthor[doc.AuthorId]
		if ok {
			item.AuthorName = author.Name
		}

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

func (m *BsDoc) DocUpdate(
	docId uint32,
	title string,
	categoryId uint32,
	accountId uint32,
	specUrl string,
	specContent string,
	updateTime int64,
) (err error) {
	dbClient := table.NewHubDB(m.Ctx)
	_, err = dbClient.AhCategoryGet(categoryId)
	if nil != err {
		logrus.Errorf("get category failed. error: %s.", err)
		return
	}

	oldDoc, err := dbClient.AhDocGet(docId)
	if nil != err {
		logrus.Errorf("get doc failed. error: %s.", err)
		return
	}

	tx, err := dbClient.DB.BeginTxx(m.Ctx, nil)
	if nil != err {
		logrus.Errorf("doc.go begin transaction failed. error: %s.", err)
		return
	}

	defer func() {
		if nil != err {
			errRoll := tx.Rollback()
			if nil != errRoll {
				logrus.Errorf("doc.go transaction rollback failed. error: %s.", err)
				return
			}
			return

		} else {
			err = tx.Commit()
			if nil != err {
				logrus.Errorf("doc.go transaction commit failed. error: %s.", err)
				return
			}
		}

	}()

	_, err = dbClient.AhDocUpdateTx(tx, docId, title, specUrl, specContent, categoryId, accountId, oldDoc.PostStatus, updateTime)
	if nil != err {
		logrus.Errorf("ah doc update failed. error: %s.", err)
		return
	}

	if categoryId != oldDoc.CategoryId {
		// 减少旧目录数目
		err = dbClient.AhCategoryDocNumIncrTx(tx, oldDoc.CategoryId, -1, updateTime)
		if nil != err {
			logrus.Errorf("descr ah_category doc_num failed. error: %s.", err)
			return
		}

		// 更新新目录数目
		err = dbClient.AhCategoryDocNumIncrTx(tx, categoryId, 1, updateTime)
		if nil != err {
			logrus.Errorf("incr ah_category doc_num failed. error: %s.", err)
			return
		}

	}

	return
}

func (m *BsDoc) DocDel(docId uint32) (err error) {
	dbClient := table.NewHubDB(m.Ctx)

	oldDoc, err := dbClient.AhDocGet(docId)
	if nil != err {
		logrus.Errorf("get doc failed. %s", err)
		return
	}

	tx, err := dbClient.DB.BeginTxx(m.Ctx, nil)
	if nil != err {
		logrus.Errorf("doc del begin transaction failed. error: %s.", err)
		return
	}

	defer func() {
		if nil != err {
			errRoll := tx.Rollback()
			if nil != errRoll {
				logrus.Errorf("doc del transaction rollback failed. error: %s.", err)
				return
			}
			return

		} else {
			err = tx.Commit()
			if nil != err {
				logrus.Errorf("doc del transaction commit failed. error: %s.", err)
				return
			}
		}

	}()

	err = dbClient.AhDocDelTx(tx, docId)
	if nil != err {
		logrus.Errorf("delete doc failed. error: %s.", err)
		return
	}

	// 更新数目
	err = dbClient.AhCategoryDocNumIncrTx(tx, oldDoc.CategoryId, -1, time.Now().Unix())
	if nil != err {
		logrus.Errorf("desc ah_category doc_num failed. error: %s.", err)
		return
	}

	return
}
