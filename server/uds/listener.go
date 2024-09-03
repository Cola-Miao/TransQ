package uds

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/Cola-Miao/TransQ/server/format"
	"io"
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

	reader := bufio.NewReader(conn)
	for {
		data, err := reader.ReadBytes('\n')
		if err != nil {
			if !errors.Is(err, io.EOF) {
				slog.Warn("reader.ReadBytes", "error", err.Error())
			} else {
				break
			}
		}
		fmt.Printf("data: %s", string(data))
	}
}
