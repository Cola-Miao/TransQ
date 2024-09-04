package uds

import (
	"encoding/json"
	"errors"
	"github.com/Cola-Miao/TransQ/server/executor"
	"github.com/Cola-Miao/TransQ/server/format"
	. "github.com/Cola-Miao/TransQ/server/models"
	"io"
	"log"
	"log/slog"
	"net"
	"time"
)

const (
	timeout = 0
)

func Listen(listener net.Listener) {
	format.FuncStart("Listen")
	defer format.FuncEnd("Listen")

	for {
		var tqc TransQClient

		conn, err := listener.Accept()
		if err != nil {
			slog.Warn("listener.Accept", "error", err.Error())
			continue
		}

		err = conn.SetDeadline(time.Now().Add(timeout))
		if err != nil {
			slog.Warn("conn.SetDeadline", "error", err.Error())
		}
		tqc.Conn = conn

		go process(&tqc)
	}
}

func process(tqc *TransQClient) {
	format.FuncStartWithData("process", tqc)
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
				log.Println("disconnect: ", tqc.Conn.LocalAddr())
				break
			} else {
				slog.Error("reader.ReadBytes", "error", err.Error())
				break
			}
		}

		err = tqc.Conn.SetDeadline(time.Now().Add(timeout))
		if err != nil {
			slog.Warn("conn.SetDeadline", "error", err.Error())
		}

		err = executor.Do(&info)
		if err != nil {
			slog.Error("executor.Do", "error", err.Error())
		}
	}
}
