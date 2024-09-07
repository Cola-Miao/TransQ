package thirdAPI

import (
	. "github.com/Cola-Miao/TransQ/server/models"
	"log/slog"
)

const (
	Lingocloud = iota + 1
)

var apis = map[int]api{
	Lingocloud: &lingocloud{},
}

type api interface {
	initialize() error
	SendMessage(tq *TransReq) (tp *TransResp)
}

func GetAPIsByID(ids ...int) ([]api, error) {
	res := make([]api, len(ids))
	for idx, id := range ids {
		eng, ok := apis[id]
		if !ok {
			return nil, ErrAPINotExist
		}
		res[idx] = eng
	}

	return res, nil
}

func InitAPIs() {
	var err error
	for i, a := range apis {
		err = a.initialize()
		if err != nil {
			slog.Error("initialize", "index", i, "error", err.Error())
		}
	}
}
