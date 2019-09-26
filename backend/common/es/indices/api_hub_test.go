package indices

import (
	"backend/common/es/model"
	"fmt"
	"github.com/haozzzzzzzz/go-rapid-development/aws/xray"
	"testing"
	"time"
)

func TestEsApiHub_AhDocIndex(t *testing.T) {
	var err error
	ctx, _, cancel := xray.NewBackgroundContext("test")
	defer func() {
		cancel(err)
	}()

	esClient := NewEsApiHub(ctx)
	err = esClient.AhDocIndex(&model.EsAhDoc{
		DocId:   4,
		Title:   "test",
		SpecUrl: "http://github.com",
		SpecPaths: []string{
			"/what",
			"/罗浩",
		},
		AuthorId:       3,
		AuthorName:     "haozzzzzzz",
		CategoryId:     3,
		CategoryName:   "default",
		CreateTime:     time.Now(),
		CreateTimeUnix: time.Now().Unix(),
	})
	if nil != err {
		t.Error(err)
		return
	}
}

func TestEsApiHub_AhDocSearch(t *testing.T) {
	var err error
	ctx, _, cancel := xray.NewBackgroundContext("test")
	defer func() {
		cancel(err)
	}()

	esClient := NewEsApiHub(ctx)
	total, items, err := esClient.AhDocSearch("内部", 1, 3)
	if nil != err {
		t.Error(err)
		return
	}

	fmt.Println(total)
	fmt.Println(items)
}

func TestEsApiHub_AhDocDeleteIndex(t *testing.T) {
	var err error
	ctx, _, cancel := xray.NewBackgroundContext("test")
	defer func() {
		cancel(err)
	}()

	esClient := NewEsApiHub(ctx)
	err = esClient.AhDocDeleteIndex()
	if nil != err {
		t.Error(err)
		return
	}
}
