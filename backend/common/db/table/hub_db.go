package table

import (
	db2 "backend/common/db"
	"backend/common/db/model"
	"context"
	"database/sql"

	"github.com/sirupsen/logrus"

	"github.com/haozzzzzzzz/go-rapid-development/db"
	"github.com/jmoiron/sqlx"
)

type HubDB struct {
	db.Client
}

func NewHubDB(ctx context.Context) *HubDB {
	return &HubDB{
		Client: db.Client{
			Ctx:    ctx,
			DB:     db2.HubSqlxDB,
			Config: db2.HubDBConfig,
		},
	}
}

// ah_doc
func (m *HubDB) AhDocAddTx(
	tx *sqlx.Tx,
	doc *model.AhDoc,
) (newId int64, err error) {
	if tx == nil {
		tx, err = m.DB.BeginTxx(m.Ctx, nil)
		if nil != err {
			logrus.Errorf("begin add ah doc tx failed. error: %s.", err)
			return
		}
		defer func() {
			if err == nil {
				err = tx.Commit()
				if nil != err {
					logrus.Errorf("commit add ah doc tx failed. error: %s.", err)
					return
				} else {
					errRoll := tx.Rollback()
					if errRoll != nil {
						logrus.Errorf("rollback add ah doc tx failed. error: %s.", err)
					}
				}
			}
		}()
	}

	strSql := `
		INSERT INTO ah_doc
		(
			doc_id, title, spec_url, spec_content, category_id, author_id, post_status, update_time, create_time
		)
		VALUES
		(
			:doc_id, :title, :spec_url, :spec_content, :category_id, :author_id, :post_status, :update_time, :create_time
		)
	`
	result, err := tx.NamedExec(strSql, doc)
	if nil != err {
		logrus.Errorf("add ah_doc record failed. error: %s.", err)
		return
	}

	newId, err = result.LastInsertId()
	if nil != err {
		logrus.Errorf("get ah_doc last insert id failed. error: %s.", err)
		return
	}

	return
}

func (m *HubDB) AhDocUpdate(
	docId uint32,
	title string,
	specUrl string,
	specContent string,
	categoryId string,
	postStatus uint8,
	updateTime int64,
) (result sql.Result, err error) {
	strSql := `
		UPDATE ah_doc
		SET
			title=?,
			spec_url=?,
			spec_content=?,
			category_id=?,
			post_status=?,
			update_time=?
		WHERE
			doc_id=?
		LIMIT 1
	`
	result, err = m.Exec(strSql, title, specUrl, specContent, categoryId, postStatus, updateTime)
	if nil != err {
		logrus.Errorf("update ah_doc failed. error: %s.", err)
		return
	}
	return
}

func (m *HubDB) AhDocGet(docId uint32) (doc *model.AhDoc, err error) {
	doc = &model.AhDoc{}
	strSql := `
		SELECT
			doc_id,
			title,
			spec_url,
			spec_content,
			category_id,
			author_id,
			post_status,
			update_time,
			create_time
		FROM
			ah_doc
		WHERE
			doc_id=?
	`
	err = m.Get(doc, strSql, docId)
	if nil != err && err != sql.ErrNoRows {
		logrus.Errorf("get ah doc failed. error: %s.", err)
		return
	}
	return
}

func (m *HubDB) AhDocGetByTitle(
	title string,
) (
	doc *model.AhDoc,
	err error,
) {
	doc = &model.AhDoc{}
	strSql := `
		SELECT
			doc_id,
			title,
			spec_url,
			category_id,
			author_id,
			post_status,
			update_time,
			create_time
		FROM
			ah_doc
		WHERE
			title=?
		LIMIT 1
	`
	err = m.Get(doc, strSql, title)
	if nil != err && err != sql.ErrNoRows {
		logrus.Errorf("get ah doc failed. error: %s.", err)
		return
	}
	return
}

func (m *HubDB) AhDocList(
	pageId uint32,
	limit uint8,
) (docs []*model.AhDoc, err error) {
	docs = make([]*model.AhDoc, 0)
	if pageId <= 0 || limit == 0 {
		return
	}

	strSql := `
		SELECT
			doc_id,
			title,
			spec_url,
			category_id,
			author_id,
			post_status,
			update_time,
			create_time
		FROM
			ah_doc
		ORDER BY doc_id DESC
		LIMIT ? OFFSET ?
`
	err = m.Select(&docs, strSql, limit, (pageId-1)*uint32(limit))
	if nil != err {
		logrus.Errorf("db get docs failed. error: %s.", err)
		return
	}

	return
}

func (m *HubDB) AhDocCount() (count int64, err error) {
	strSql := `
		SELECT COUNT(*) FROM ah_doc
	`
	err = m.Get(&count, strSql)
	if nil != err {
		logrus.Errorf("get ah_doc count failed. error: %s.", err)
		return
	}
	return
}

// ah_account
func (m *HubDB) AhAccountGet(
	accId uint32,
) (account *model.AhAccount, err error) {
	account = &model.AhAccount{}
	strSql := `
		SELECT
			acc_id,
			name,
			update_time,
			create_time
		FROM
			ah_account
		WHERE
			acc_id=?
	`
	err = m.Get(account, strSql, accId)
	if nil != err && err != sql.ErrNoRows {
		logrus.Errorf("get account failed. error: %s.", err)
		return
	}
	return
}

func (m *HubDB) AhAccountBatch(accIds []uint32) (mAccount map[uint32]*model.AhAccount, err error) {
	mAccount = make(map[uint32]*model.AhAccount)
	if len(accIds) == 0 {
		return
	}
	accounts := make([]*model.AhAccount, 0)
	strSql := `
		SELECT
			acc_id,
			name,
			update_time,
			create_time
		FROM
			ah_account
		WHERE
			acc_id IN (?)
	`
	query, args, err := sqlx.In(strSql, accIds)
	if nil != err {
		logrus.Errorf("use sqlx in failed. error: %s.", err)
		return
	}

	err = m.Select(&accounts, query, args...)
	if nil != err {
		logrus.Errorf("batch get ah_account failed. error: %s.", err)
		return
	}

	for _, acc := range accounts {
		mAccount[acc.AccId] = acc
	}

	return
}

// 目录
func (m *HubDB) AhCategoryGet(
	catId uint32,
) (cat *model.AhCategory, err error) {
	cat = &model.AhCategory{}
	strSql := `
		SELECT
			cat_id,
			name,
			doc_num,
			update_time,
			create_time
		FROM
			ah_category
		WHERE
			cat_id=?
	`
	err = m.Get(cat, strSql, catId)
	if nil != err {
		logrus.Errorf("get ah_category failed. error: %s.", err)
		return
	}
	return
}

func (m *HubDB) AhCategoryBatch(catIds []uint32) (mCategory map[uint32]*model.AhCategory, err error) {
	mCategory = make(map[uint32]*model.AhCategory)
	if len(catIds) == 0 {
		return
	}

	categories := make([]*model.AhCategory, 0)
	strSql := `
		SELECT
			cat_id,
			name,
			doc_num,
			update_time,
			create_time
		FROM
			ah_category
		WHERE
			cat_id IN (?)
	`
	query, args, err := sqlx.In(strSql, catIds)
	if nil != err {
		logrus.Errorf("use sqlx in failed. error: %s.", err)
		return
	}

	err = m.Select(&categories, query, args...)
	if nil != err {
		logrus.Errorf("get categories failed. error: %s.", err)
		return
	}

	for _, cat := range categories {
		mCategory[cat.CatId] = cat
	}

	return
}

func (m *HubDB) AhCategoryDocNumIncrTx(
	tx *sqlx.Tx,
	catId uint32,
	incrNum uint32,
	updateTime int64,
) (err error) {
	strSql := `
		UPDATE ah_category
		SET
			doc_num = doc_num + ?,
			update_time = ?
		WHERE
			cat_id = ?
		LIMIT 1
	`
	_, err = tx.Exec(strSql, incrNum, updateTime, catId)
	if nil != err {
		logrus.Errorf("incr ah_category doc num failed. error: %s.", err)
		return
	}
	return
}

//// 标签
//func (m *HubDB) AhTagByNames(
//	names []string,
//) (tags []*model.AhTag, err error) {
//	tags = make([]*model.AhTag, 0)
//	strSql := `
//		SELECT
//			tag_id,
//			name,
//			doc_num,
//			update_time,
//			create_time
//		FROM
//			ah_tag
//		WHERE
//			name IN (?)
//	`
//	query, args, err := sqlx.In(strSql, names)
//	if nil != err {
//		logrus.Errorf("use sqlx in failed. error: %s.", err)
//		return
//	}
//
//	err = m.Select(tags, query, args...)
//	if nil != err {
//		logrus.Errorf("get ah_tag by names failed. error: %s.", err)
//		return
//	}
//
//	return
//}
//
//func (m *HubDB) AhDocTagByDocId(
//	docId uint32,
//) (docTags []*model.AhDocTag, err error) {
//	docTags = make([]*model.AhDocTag, 0)
//	strSql := `
//		SELECT
//			doc_id,
//			tag_id,
//			create_time
//		FROM
//			ah_doc_tag
//		WHERE
//			doc_id=?
//	`
//	err = m.Select(docTags, strSql, docId)
//	if nil != err {
//		logrus.Errorf("get ah_doc_tag by doc_id failed. doc_id: %d, error: %s.", docId, err)
//		return
//	}
//
//	if len(docTags) == 0 {
//		return
//	}
//
//	tagIds := make([]uint32, 0)
//	for _, docTag := range docTags {
//		tagIds = append(tagIds, docTag.TagId)
//	}
//
//	tags := make([]*model.AhTag, 0)
//	strSql = `
//		SELECT
//			tag_id,
//			name,
//			doc_num,
//			update_time,
//			create_time
//		FROM
//			ah_tag
//		WHERE
//			tag_id IN (?)
//	`
//	query, args, err := sqlx.In(strSql, tagIds)
//	if nil != err {
//		logrus.Errorf("use in failed. str_sql: %s, tag_ids: %#v, error: %s.", tagIds, err)
//		return
//	}
//
//	err = m.Select(tags, query, args...)
//	if nil != err {
//		logrus.Errorf("get ah_doc_tag by doc_id failed. error: %s.", err)
//		return
//	}
//
//	mTags := make(map[uint32]*model.AhTag)
//	for _, tag := range tags {
//		mTags[tag.TagId] = tag
//	}
//
//	for _, docTag := range docTags {
//		tag, ok := mTags[docTag.TagId]
//		if !ok {
//			continue
//		}
//
//		docTag.TagName = tag.Name
//	}
//
//	return
//}
