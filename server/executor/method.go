package executor

import (
	"encoding/json"
	"fmt"
	"github.com/Cola-Miao/TransQ/server/format"
)

func echo(data string) error {
	format.FuncStart("echo")
	defer format.FuncEnd("echo")

	var req echoRequest

	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return fmt.Errorf("echo: %w", err)
	}

	fmt.Println("echo resp: ", req.Message)

	return nil
}

func translate(data string) error {
	return nil
}
