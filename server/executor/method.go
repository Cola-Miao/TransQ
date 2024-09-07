package executor

import (
	"fmt"
	"github.com/Cola-Miao/TransQ/server/format"
	. "github.com/Cola-Miao/TransQ/server/models"
)

func mtdAuth(tqc *transQClient, str any) (any, error) {
	format.FuncStart("mtdAuth")
	defer format.FuncEnd("mtdAuth")

	req, ok := str.(*authRequest)
	if !ok {
		return nil, ErrBadRequestType
	}

	resp := &authResponse{
		common{
			Sequence: req.Sequence,
			Code:     success,
		}}

	err := auth(tqc, req)
	if err != nil {
		resp.Code = failed
		return resp, err
	}

	return resp, nil
}

func mtdEcho(tqc *transQClient, str any) (any, error) {
	format.FuncStart("mtdEcho")
	defer format.FuncEnd("mtdEcho")

	req, ok := str.(*echoRequest)
	if !ok {
		return nil, ErrBadRequestType
	}

	fmt.Println("echo resp: ", req.Message)

	return nil, nil
}

func mtdTranslate(tqc *transQClient, str any) (any, error) {
	format.FuncStart("mtdTranslate")
	defer format.FuncEnd("mtdTranslate")

	req, ok := str.(*translateRequest)
	if !ok {
		return nil, ErrBadRequestType
	}

	resp, err := translate(req)
	if err != nil {
		resp.Code = failed
		return resp, err
	}

	return nil, nil
}
