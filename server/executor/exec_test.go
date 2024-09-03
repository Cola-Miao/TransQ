package executor

import (
	. "github.com/Cola-Miao/TransQ/server/models"
	"testing"
)

var (
	echoInfo = &Information{
		Method: 1,
		Data:   "{\"msg\":\"ping\"}",
	}
)

func TestDo(t *testing.T) {
	err := Do(echoInfo)
	if err != nil {
		t.Error(err)
	}
}
