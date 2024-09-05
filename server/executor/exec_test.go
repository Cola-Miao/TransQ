package executor

import (
	. "github.com/Cola-Miao/TransQ/server/models"
	"testing"
)

var (
	echoInfo = &information{
		Method: 1,
		Data:   "{\"msg\":\"ping\"}",
	}
)

func TestDo(t *testing.T) {
	err := exec.do(echoInfo)
	if err != nil {
		t.Error(err)
	}
}
