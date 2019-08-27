package table

import (
	"context"
	"testing"
)

func TestHubDB_InitTables(t *testing.T) {
	ctx := context.Background()
	dbClient := NewHubDB(ctx)
	err := dbClient.InitTables()
	if nil != err {
		t.Error(err)
		return
	}
}
