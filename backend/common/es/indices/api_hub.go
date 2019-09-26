package indices

import (
	"backend/common/config"
	"backend/common/es/model"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esutil"
	"github.com/gosexy/to"
	"github.com/haozzzzzzzz/go-rapid-development/es"
	"github.com/haozzzzzzzz/go-rapid-development/utils/uerrors"
	"github.com/haozzzzzzzz/go-rapid-development/utils/uio"
	"github.com/sirupsen/logrus"
	"net/http"
)

type EsApiHub struct {
	Ctx    context.Context
	Client *elasticsearch.Client
}

func NewEsApiHub(ctx context.Context) *EsApiHub {
	return &EsApiHub{
		Ctx:    ctx,
		Client: EsClient,
	}
}

func (m *EsApiHub) AhDocIndexName() string {
	return config.ElasticSearchConfig.IndexAhDoc
}

func (m *EsApiHub) AhDocExists() (exists bool, err error) {
	indexName := m.AhDocIndexName()
	existsFunc := m.Client.Indices.Exists
	resp, err := existsFunc([]string{indexName},
		existsFunc.WithContext(m.Ctx),
	)
	if nil != err {
		logrus.Errorf("get ah doc exists failed. error: %s.", err)
		return
	}

	switch resp.StatusCode {
	case http.StatusOK:
		exists = true
	case http.StatusNotFound:
	default:
		err = uerrors.Newf("get exits return unsupported code: %d", resp.StatusCode)
	}

	return
}

func (m *EsApiHub) AhDocCreateIndex() (err error) {
	indexName := m.AhDocIndexName()
	mapping := `{
  "mappings": {
	"dynamic_templates": [
	  {
		"ik_max_word_type": {
		  "match": "*",
		  "mapping": {
			  "analyzer": "ik_max_word",
			  "search_analyzer": "ik_smart"
		  }
		}
	  }
	]
  }
}`

	mappingReader := bytes.NewBuffer([]byte(mapping))
	createFunc := m.Client.Indices.Create
	resp, err := createFunc(
		indexName,
		createFunc.WithContext(m.Ctx),
		createFunc.WithBody(mappingReader),
	)
	if nil != err {
		logrus.Errorf("create index %s failed. error: %s.", indexName, err)
		return
	}

	if resp.IsError() {
		bBody, errRead := uio.ReadAllAndClose(resp.Body)
		err = errRead
		if nil != err {
			logrus.Errorf("read response body failed. error: %s.", err)
			return
		}

		err = es.RespError(bBody)
		if nil != err {
			logrus.Errorf("response error. error: %s.", err)
			return
		}
	}
	return
}

func (m *EsApiHub) AhDocIndex(esAhDoc *model.EsAhDoc) (err error) {
	indexName := m.AhDocIndexName()
	indexFunc := m.Client.Index
	resp, err := indexFunc(
		indexName,
		esutil.NewJSONReader(esAhDoc),
		indexFunc.WithDocumentID(to.String(esAhDoc.DocId)),
		indexFunc.WithContext(m.Ctx),
	)
	if nil != err {
		logrus.Errorf("index ah doc failed. error: %s.", err)
		return
	}

	bBody, err := uio.ReadAllAndClose(resp.Body)
	if nil != err {
		logrus.Errorf("read response body failed. error: %s.", err)
		return
	}

	if resp.IsError() {
		err = es.RespError(bBody)
		if nil != err {
			logrus.Errorf("response error. error: %s.", err)
			return
		}
		return
	}

	return
}

func (m *EsApiHub) AhDocDeleteIndex() (err error) {
	indexName := m.AhDocIndexName()
	deleteFunc := m.Client.Indices.Delete
	resp, err := deleteFunc([]string{indexName}, deleteFunc.WithContext(m.Ctx))
	if nil != err {
		logrus.Errorf("es delete ah_doc index failed. error: %s.", err)
		return
	}

	bRespBody, err := uio.ReadAllAndClose(resp.Body)
	if nil != err {
		logrus.Errorf("read response body failed. error: %s.", err)
		return
	}

	if resp.IsError() {
		err = es.RespError(bRespBody)
		if nil != err {
			logrus.Errorf("response error. error: %s.", err)
			return
		}
		return
	}

	return
}

func (m *EsApiHub) AhDocDelete(docId uint32) (err error) {
	indexName := m.AhDocIndexName()
	deleteFunc := m.Client.Delete
	resp, err := deleteFunc(
		indexName,
		to.String(docId),
		deleteFunc.WithContext(m.Ctx),
	)
	if nil != err {
		logrus.Errorf("delete ah doc failed. error: %s.", err)
		return
	}

	bBody, err := uio.ReadAllAndClose(resp.Body)
	if nil != err {
		logrus.Errorf("read response body failed. error: %s.", err)
		return
	}

	if resp.IsError() {
		err = es.RespError(bBody)
		if nil != err {
			logrus.Errorf("response error. error: %s.", err)
			return
		}
		return
	}

	return
}

type SearchAhDocResponseData struct {
	Took    int64 `json:"took"`
	Timeout bool  `json:"timeout"`
	Hits    struct {
		Total struct {
			Value    int64  `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		Hits []struct {
			InIndex  string         `json:"_index"`
			InType   string         `json:"_type"`
			InId     string         `json:"_id"`
			InSource *model.EsAhDoc `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

func (m *EsApiHub) AhDocSearch(
	word string,
	page uint32,
	size uint8,
) (total int64, items []*model.EsAhDoc, err error) {
	items = make([]*model.EsAhDoc, 0)

	queryBody := `
{
  "from": %d,
  "size": %d,
  "sort": {
    "_score": {
      "order": "desc"
    },
    "doc_id": {
      "order": "desc"
    }
  },
  "query": {
    "bool": {
      "should": [
        {
          "multi_match": {
          "query": "%s",
          "fields": ["title","spec_url","spec_paths","author_name",   "category_name"]
          }
        }
      ]
    }
  }
}`
	queryBody = fmt.Sprintf(queryBody, (page-1)*uint32(size), size, word)
	bodyReader := bytes.NewBuffer([]byte(queryBody))

	indexName := m.AhDocIndexName()
	searchFunc := m.Client.Search
	resp, err := searchFunc(
		searchFunc.WithContext(m.Ctx),
		searchFunc.WithIndex(indexName),
		searchFunc.WithBody(bodyReader),
	)
	if nil != err {
		logrus.Errorf("search ah_doc failed. error: %s.", err)
		return
	}

	bBody, err := uio.ReadAllAndClose(resp.Body)
	if nil != err {
		logrus.Errorf("read response body failed. error: %s.", err)
		return
	}

	if resp.IsError() {
		err = es.RespError(bBody)
		if nil != err {
			logrus.Errorf("response error. error: %s.", err)
			return
		}
		return
	}

	respData := &SearchAhDocResponseData{}
	err = json.Unmarshal(bBody, respData)
	if nil != err {
		logrus.Errorf("unmarshal resp ah_doc failed. error: %s.", err)
		return
	}

	total = respData.Hits.Total.Value

	for _, hit := range respData.Hits.Hits {
		items = append(items, hit.InSource)
	}

	return
}
