package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Cola-Miao/TransQ/server/format"
	"io"
	"log/slog"
	"net/http"
)

const (
	lingocloudURL = "http://api.interpreter.caiyunai.com/v1/translator"
)

type lingocloud struct {
	header http.Header
}

func (l *lingocloud) initialize(token string) error {
	hd := http.Header{}
	hd.Set("content-type", "application/json")
	hd.Set("x-authorization", "token "+token)
	l.header = hd

	return nil
}

func (l *lingocloud) sendMessage(tq *TransReq) (tp *TransResp) {
	format.FuncStartWithData("sendMessage", tq)
	defer format.FuncEndWithData("sendMessage", tp)

	tp = &TransResp{}

	payload := map[string]any{
		"source":     "Lingocloud is the best translation service.",
		"trans_type": "auto2zh",
		"request_id": "demo",
		"detect":     true,
	}

	js, err := json.Marshal(payload)
	if err != nil {
		tp.err = fmt.Errorf("json.Marshal: %w", err)
		return
	}

	req, err := http.NewRequest("POST", lingocloudURL, bytes.NewReader(js))
	if err != nil {
		tp.err = fmt.Errorf("http.NewRequest: %w", err)
		return
	}
	req.Header = l.header

	resp, err := client.Do(req)
	if err != nil {
		tp.err = fmt.Errorf("client.Do: %w", err)
		return
	}

	defer func() {
		der := resp.Body.Close()
		if der != nil {
			slog.Error("resp.Body.Close", "error", err.Error())
		}
	}()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		tp.err = fmt.Errorf("io.ReadAll: %w", err)
		return
	}

	var kv map[string]any
	err = json.Unmarshal(data, &kv)
	if err != nil {
		tp.err = fmt.Errorf("json.Unmarshal: %w", err)
		return
	}

	message, ok := kv["message"].(string)
	if ok {
		tp.err = errors.New(message)
		return
	}

	target, ok := kv["target"].(string)
	if !ok {
		tp.err = errors.New("target is not a string")
		return
	}

	tp.Message = target
	return
}
