package executor

import (
	"encoding/json"
	"fmt"
	"github.com/Cola-Miao/TransQ/server/format"
	"log/slog"
)

func mtdAuth(tqc *transQClient) error {
	format.FuncStart("mtdAuth")
	defer format.FuncEnd("mtdAuth")

	var req authRequest
	err := json.Unmarshal([]byte(tqc.Info.Data), &req)
	if err != nil {
		return fmt.Errorf("json.Unmarshal: %w", err)
	}

	resp := authResponse{
		common{
			Sequence: req.Sequence,
			Code:     success,
		}}
	defer func() {
		der := exec.writeConn(req.ID, &resp)
		if der != nil {
			slog.Error("exec.writeConn", "error", err.Error(), "id", req.ID)
		}
	}()

	err = auth(tqc, &req)
	if err != nil {
		resp.Code = failed
		return fmt.Errorf("auth: %w", err)
	}

	return nil
}

func mtdEcho(tqc *transQClient) error {
	format.FuncStart("mtdEcho")
	defer format.FuncEnd("mtdEcho")

	var req echoRequest

	err := json.Unmarshal([]byte(tqc.Info.Data), &req)
	if err != nil {
		return fmt.Errorf("json.Unmarshal: %w", err)
	}

	fmt.Println("echo resp: ", req.Message)

	return nil
}

func mtdTranslate(tqc *transQClient) error {
	return nil
}
