package uds

import (
	. "github.com/Cola-Miao/TransQ/server/config"
	"github.com/Cola-Miao/TransQ/server/executor"
	"github.com/Cola-Miao/TransQ/server/format"
	"github.com/Cola-Miao/TransQ/server/utils"
	"log/slog"
	"net"
)

func Listen(listener net.Listener) {
	format.FuncStart("Listen")
	defer format.FuncEnd("Listen")

	for {
		conn, err := utils.AcceptSocketWithTimeout(listener, Cfg.Listener.Timeout)
		if err != nil {
			slog.Error("utils.AcceptSocketWithTimeout", "error", err.Error())
			continue
		}
		slog.Info("connect", "addr", conn.LocalAddr().String())

		go executor.Process(conn)
	}
}
