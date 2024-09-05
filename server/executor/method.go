package executor

import (
	"encoding/json"
	"fmt"
	"github.com/Cola-Miao/TransQ/server/format"
	. "github.com/Cola-Miao/TransQ/server/models"
)

func mtdAuth(tqc *transQClient) error {
	format.FuncStart("auth")
	defer format.FuncEnd("auth")

	var req AuthRequest
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
	format.FuncStart("echo")
	defer format.FuncEnd("echo")

	var req EchoRequest

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
