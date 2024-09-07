package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Cola-Miao/TransQ/server/format"
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

	req, err := l.generateRequest()
	if err != nil {
		tp.err = fmt.Errorf("generateRequest: %w", err)
		return
	}

	kv, err := getRespBodyMap(req)
	if err != nil {
		tp.err = fmt.Errorf("getRespBodyMap: %w", err)
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

func (l *lingocloud) generateRequest() (*http.Request, error) {
	payload := map[string]any{
		"source":     "Lingocloud is the best translation service.",
		"trans_type": "auto2zh",
		"request_id": "demo",
		"detect":     true,
	}

	js, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal: %w", err)
	}

	req, err := http.NewRequest("POST", lingocloudURL, bytes.NewReader(js))
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest: %w", err)
	}

	req.Header = l.header
	return req, nil
}
