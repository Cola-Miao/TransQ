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
		conn, err := listener.Accept()
		if err != nil {
			slog.Warn("listener.Accept", "error", err.Error())
			continue
		}
		slog.Info("connect", "addr", conn.LocalAddr().String())

		err = conn.SetDeadline(utils.GetOutTime(Cfg.ConnTimeout))
		if err != nil {
			slog.Warn("conn.SetDeadline", "error", err.Error())
		}

		go executor.Process(conn)
	}
}
