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
		tqc.Conn = conn

		log.Println("connect: ", conn.LocalAddr())
		go process(&tqc)
	}
}

func process(tqc *TransQClient) {
	format.FuncStart("process")
	defer func() {
		if err := tqc.Conn.Close(); err != nil {
			slog.Warn("conn.Close", "error", err.Error())
		}
		format.FuncEnd("process")
	}()

	decoder := json.NewDecoder(tqc.Conn)
	for {
		var info Information

		err := decoder.Decode(&info)
		if err != nil {
			if !errors.Is(err, io.EOF) {
				slog.Error("reader.ReadBytes", "error", err.Error())
				break
			} else {
				log.Println("disconnect: ", tqc.Conn.LocalAddr())
				break
			}
		}

		err = executor.Do(&info)
		if err != nil {
			slog.Error("executor.Do", "error", err.Error())
		}
	}
}
