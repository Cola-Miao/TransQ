package executor

import (
	"encoding/json"
	"errors"
	. "github.com/Cola-Miao/TransQ/server/config"
	"github.com/Cola-Miao/TransQ/server/format"
	"github.com/Cola-Miao/TransQ/server/utils"
	"io"
	"log/slog"
	"net"
)

func Process(conn net.Conn) {
	format.FuncStart("Process")
	defer func() {
		if err := conn.Close(); err != nil {
			slog.Warn("conn.Close", "error", err.Error())
		}
		format.FuncEnd("Process")
	}()

	var tqc transQClient
	tqc.Conn = conn

	processLoop(&tqc)
}

func processLoop(tqc *transQClient) {
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

		err = tqc.Conn.SetDeadline(utils.GetOutTime(Cfg.Listener.Timeout))
		if err != nil {
			slog.Warn("conn.SetDeadline", "error", err.Error())
		}

		err = exec.do(tqc)
		if err != nil {
			slog.Error("executor.Do", "error", err.Error())
		}
	}
}
