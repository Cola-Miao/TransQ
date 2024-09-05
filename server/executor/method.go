package executor

import (
	"fmt"
	"github.com/Cola-Miao/TransQ/server/format"
	"log/slog"
)

func mtdAuth(tqc *transQClient, str any) error {
	format.FuncStart("mtdAuth")
	defer format.FuncEnd("mtdAuth")

	req, ok := str.(*authRequest)
	if !ok {
		return errBadRequestType
	}

	resp := authResponse{
		common{
			Sequence: req.Sequence,
			Code:     success,
		}}
	defer func() {
		der := exec.writeConn(req.ID, &resp)
		if der != nil {
			slog.Error("exec.writeConn", "error", der.Error(), "id", req.ID)
		}
	}()

	err := auth(tqc, req)
	if err != nil {
		resp.Code = failed
		return fmt.Errorf("auth: %w", err)
	}

	return nil
}

func mtdEcho(tqc *transQClient, str any) error {
	format.FuncStart("mtdEcho")
	defer format.FuncEnd("mtdEcho")

	req, ok := str.(*echoRequest)
	if !ok {
		return errBadRequestType
	}

	fmt.Println("echo resp: ", req.Message)

	return nil
}

func mtdTranslate(tqc *transQClient, str any) error {
	return nil
}
