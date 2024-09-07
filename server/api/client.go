package api

import (
	"encoding/json"
	"fmt"
	"github.com/Cola-Miao/TransQ/server/format"
	"io"
	"log/slog"
	"net/http"
)

var client http.Client

func init() {
	c := http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       0,
	}

	client = c
}

func getRespBodyMap(req *http.Request) (map[string]any, error) {
	format.FuncStart("getRespBodyMap")
	defer format.FuncEnd("getRespBodyMap")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("client.Do: %w", err)
	}

	defer func() {
		der := resp.Body.Close()
		if der != nil {
			slog.Error("resp.Body.Close", "error", err.Error())
		}
	}()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll: %w", err)
	}

	var kv map[string]any
	err = json.Unmarshal(data, &kv)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return kv, nil
}
