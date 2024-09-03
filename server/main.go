package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net"
	"os"
	"path"
)

const (
	socket = "transQ.sock"
	folder = ".transQ/"
)

var (
	workDir    string
	socketPath string
	listener   net.Listener
)

func init() {
	wd, err := initWorkDir()
	if err != nil {
		log.Panicf("initWorkDir: %s", err.Error())
	}
	workDir = wd
	socketPath = path.Join(workDir, socket)

	ls, err := initSocketListener()
	if err != nil {
		log.Panicf("initSocketListener: %s", err.Error())
	}
	listener = ls
}

func main() {
	defer func() {
		if err := listener.Close(); err != nil {
			slog.Warn("listener.Close", "error", err.Error())
		}

		if err := os.Remove(socketPath); err != nil && !errors.Is(err, os.ErrNotExist) {
			slog.Warn("os.Remove", "error", err.Error())
		}
	}()

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
	defer func() {
		if err := conn.Close(); err != nil {
			slog.Warn("conn.Close()", "error", err.Error())
		}
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
		fmt.Println(string(data))
	}
}
