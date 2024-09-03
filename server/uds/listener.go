package uds

import (
	"encoding/json"
	"errors"
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
		conn, err := listener.Accept()
		if err != nil {
			slog.Warn("listener.Accept", "error", err.Error())
			continue
		}
		log.Println("connect: ", conn.LocalAddr())
		go process(conn)
	}
}

func process(conn net.Conn) {
	format.FuncStart("process")
	defer func() {
		if err := conn.Close(); err != nil {
			slog.Warn("conn.Close()", "error", err.Error())
		}
		format.FuncEnd("process")
	}()

	decoder := json.NewDecoder(conn)
	for {
		var info Information
		err := decoder.Decode(&info)
		if err != nil {
			if !errors.Is(err, io.EOF) {
				slog.Warn("reader.ReadBytes", "error", err.Error())
			} else {
				log.Println("disconnect: ", conn.LocalAddr())
				break
			}
		}
	}
}
