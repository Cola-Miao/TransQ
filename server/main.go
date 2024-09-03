package main

import (
	"errors"
	"github.com/Cola-Miao/TransQ/server/format"
	"github.com/Cola-Miao/TransQ/server/uds"
	"log"
	"log/slog"
	"net"
	"os"
	"path"
)

const (
	socket = "transQ.sock"
	folder = ".transQ/"

	logLever = slog.LevelDebug
)

var (
	workDir    string
	logDir     string
	socketPath string
	listener   net.Listener
)

func init() {
	format.FuncStart("init")
	defer format.FuncEnd("init")

	wd, err := initWorkDir()
	if err != nil {
		log.Panicf("initWorkDir: %s", err.Error())
	}
	workDir = wd
	socketPath = path.Join(workDir, socket)

	initEnvWithGOOS()

	err = initSlog()
	if err != nil {
		log.Panicf("initSlog: %s", err.Error())
	}

	ls, err := initSocketListener()
	if err != nil {
		log.Panicf("initSocketListener: %s", err.Error())
	}
	listener = ls
}

func main() {
	format.FuncStart("main")
	defer func() {
		if err := listener.Close(); err != nil {
			slog.Warn("listener.Close", "error", err.Error())
		}

		if err := os.Remove(socketPath); err != nil && !errors.Is(err, os.ErrNotExist) {
			slog.Warn("os.Remove", "error", err.Error())
		}

		format.FuncEnd("main")
	}()

	uds.Listen(listener)
}
