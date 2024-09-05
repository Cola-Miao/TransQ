package executor

import (
	"encoding/json"
	"errors"
	"fmt"
	. "github.com/Cola-Miao/TransQ/server/config"
	"github.com/Cola-Miao/TransQ/server/format"
	"github.com/Cola-Miao/TransQ/server/utils"
	"io"
	"log"
	"log/slog"
	"net"
)

func init() {
	exec.init()
}

func (e *executor) init() {
	format.FuncStart("executor.init")
	defer format.FuncEnd("executor.init")

	e.handle = make(map[method]handler)
	e.name = make(map[method]string)

	e.register(methodAuth, auth, "auth")
	e.register(methodEcho, echo, "echo")
	e.register(methodTranslate, translate, "translate")
}

func (e *executor) register(method method, handle handler, name string) {
	if _, ok := e.handle[method]; ok {
		log.Panicf("e.handle has method: %d", method)
	}
	e.handle[method] = handle

	if _, ok := e.name[method]; ok {
		log.Panicf("e.name has method: %d", method)
	}
	e.name[method] = name
}

func (e *executor) do(tqc *transQClient) error {
	format.FuncStartWithData("executor.do", tqc)
	defer format.FuncEnd("executor.do")

	if _, ok := e.handle[tqc.Info.Method]; !ok {
		return errNoMethod
	}

	err := e.handle[tqc.Info.Method](tqc)
	if err != nil {
		return fmt.Errorf("method: %s: %w", e.name[tqc.Info.Method], err)
	}

	return nil
}

func Process(conn net.Conn) {
	format.FuncStart("process")
	defer func() {
		if err := conn.Close(); err != nil {
			slog.Warn("conn.Close", "error", err.Error())
		}
		format.FuncEnd("process")
	}()

	var tqc transQClient
	tqc.Conn = conn

	decoder := json.NewDecoder(tqc.Conn)
	for {
		err := decoder.Decode(&tqc.Info)
		if err != nil {
			if errors.Is(err, io.EOF) {
				slog.Info("disconnect", "addr", tqc.Conn.LocalAddr().String())
				break
			} else {
				slog.Error("reader.ReadBytes", "error", err.Error())
				break
			}
		}

		err = tqc.Conn.SetDeadline(utils.GetOutTime(Cfg.ConnTimeout))
		if err != nil {
			slog.Warn("conn.SetDeadline", "error", err.Error())
		}

		err = exec.do(&tqc)
		if err != nil {
			slog.Error("executor.Do", "error", err.Error())
		}
	}
}
