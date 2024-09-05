package uds

import (
	"encoding/json"
	"errors"
	"github.com/Cola-Miao/TransQ/server/executor"
	"github.com/Cola-Miao/TransQ/server/format"
	. "github.com/Cola-Miao/TransQ/server/models"
	"io"
	"log/slog"
	"net"
	"time"
)

const (
	timeout = time.Minute * 0
)

func Listen(listener net.Listener) {
	format.FuncStart("Listen")
	defer format.FuncEnd("Listen")

	for {
		var tqc transQClient

		conn, err := listener.Accept()
		if err != nil {
			slog.Warn("listener.Accept", "error", err.Error())
			continue
		}
		slog.Info("connect", "addr", conn.LocalAddr().String())

		err = conn.SetDeadline(getOutTime())
		if err != nil {
			slog.Warn("conn.SetDeadline", "error", err.Error())
		}
		tqc.Conn = conn

		go tqc.process()
	}
}

func (tqc *transQClient) process() {
	format.FuncStart("process")
	defer func() {
		if err := tqc.Conn.Close(); err != nil {
			slog.Warn("conn.Close", "error", err.Error())
		}
		format.FuncEndWithData("process", tqc)
	}()

	decoder := json.NewDecoder(tqc.Conn)
	for {
		var info Information

		err := decoder.Decode(&info)
		if err != nil {
			if errors.Is(err, io.EOF) {
				slog.Info("disconnect", "addr", tqc.Conn.LocalAddr().String())
				break
			} else {
				slog.Error("reader.ReadBytes", "error", err.Error())
				break
			}
		}

		err = tqc.Conn.SetDeadline(getOutTime())
		if err != nil {
			slog.Warn("conn.SetDeadline", "error", err.Error())
		}

		err = executor.Do(&info)
		if err != nil {
			slog.Error("executor.Do", "error", err.Error())
		}
	}
}

func getOutTime() time.Time {
	if timeout == 0 {
		return time.Time{}
	}
	return time.Now().Add(timeout)
}
