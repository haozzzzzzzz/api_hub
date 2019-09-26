package business

import (
	"github.com/haozzzzzzzz/go-rapid-development/aws/xray"
	"testing"
)

func TestBsDoc_EsIndexAllAhDoc(t *testing.T) {
	var err error
	ctx, _, cancel := xray.NewBackgroundContext("test")
	defer func() {
		cancel(err)
	}()

	bsDoc := NewBsDoc(ctx)
	err = bsDoc.EsIndexAllAhDoc()
	if nil != err {
		t.Error(err)
		return
	}
}
