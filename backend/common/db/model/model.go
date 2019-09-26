package model

// api 文档
const (
	PostStatusNotPublished uint8 = iota
	PostStatusPublished
	PostStatusDeleted
)

const DefaultAccountId uint32 = 1
const DefaultCategoryId uint32 = 1

type SwaggerSpec struct {
	Paths map[string]map[string]struct {
		Tags []string `json:"tags"`
	} `json:"paths"`
}

type AhDoc struct {
	DocId       uint32 `json:"doc_id" db:"doc_id"`
	Title       string `json:"title" db:"title"`
	SpecUrl     string `json:"spec_url" db:"spec_url"`
	SpecContent string `json:"spec_content" db:"spec_content"`
	CategoryId  uint32 `json:"category_id" db:"category_id"`
	AuthorId    uint32 `json:"author_id" db:"author_id"`
	PostStatus  uint8  `json:"post_status" db:"post_status"`
	UpdateTime  int64  `json:"update_time" db:"update_time"`
	CreateTime  int64  `json:"create_time" db:"create_time"`
}

type AhDocWithName struct {
	AhDoc
	AuthorName   string
	CategoryName string
}

// 账户
type AhAccount struct {
	AccId      uint32 `json:"acc_id" db:"acc_id"`
	Name       string `json:"name" db:"name"`
	UpdateTime int64  `json:"update_time" db:"update_time"`
	CreateTime int64  `json:"create_time" db:"create_time"`
}

// 目录
type AhCategory struct {
	CatId      uint32 `json:"cat_id" db:"cat_id"`
	Name       string `json:"name" db:"name"`
	DocNum     uint32 `json:"doc_num" db:"doc_num"`
	UpdateTime int64  `json:"update_time" db:"update_time"`
	CreateTime int64  `json:"create_time" db:"create_time"`
}

//// 标签
//type AhTag struct {
//	TagId      uint32 `json:"tag_id" db:"tag_id"`
//	Name       string `json:"name" db:"name"`
//	DocNum     uint32 `json:"doc_num" db:"doc_num"`
//	UpdateTime int64  `json:"update_time" db:"update_time"`
//	CreateTime int64  `json:"create_time" db:"create_time"`
//}
//
//type AhDocTag struct {
//	DocId      uint32 `json:"doc_id" db:"doc_id"`
//	TagId      uint32 `json:"tag_id" db:"tag_id"`
//	TagOrder   uint32 `json:"tag_order" db:"tag_order"`
//	CreateTime uint32 `json:"create_time" db:"create_time"`
//	TagName    string `json:"tag_name"`
//}
