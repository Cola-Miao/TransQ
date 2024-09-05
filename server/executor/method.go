package executor

import (
	"encoding/json"
	"fmt"
	"github.com/Cola-Miao/TransQ/server/format"
)

func mtdAuth(tqc *transQClient) error {
	format.FuncStart("mtdAuth")
	defer format.FuncEnd("mtdAuth")

	var req authRequest
	err := json.Unmarshal([]byte(tqc.Info.Data), &req)
	if err != nil {
		return fmt.Errorf("json.Unmarshal: %w", err)
	}

	err = auth(tqc, &req)
	if err != nil {
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
