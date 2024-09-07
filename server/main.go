package main

import (
	"errors"
	"github.com/Cola-Miao/TransQ/server/cache"
	cfg "github.com/Cola-Miao/TransQ/server/config"
	"github.com/Cola-Miao/TransQ/server/dao"
	"github.com/Cola-Miao/TransQ/server/format"
	"github.com/Cola-Miao/TransQ/server/uds"
	"log"
	"log/slog"
	"net"
	"os"
	"path"
)

const (
	socket     = "transQ.sock"
	folder     = ".transQ/"
	configType = "yaml"
	database   = "transQ.sqlite"
)

var (
	workDir      string
	logDir       string
	socketPath   string
	databasePath string
	listener     net.Listener
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
	databasePath = path.Join(workDir, database)

	initEnvWithGOOS()

	err = cfg.InitViper(workDir, configType)
	if err != nil {
		slog.Warn("cfg.InitViper", "error", err.Error())
	}

	err = initSlog()
	if err != nil {
		log.Panicf("initSlog: %s", err.Error())
	}

	err = dao.InitSqlite(databasePath)
	if err != nil {
		log.Panicf("dao.InitSqlite: %s", err.Error())
	}

	err = cache.InitCache()
	if err != nil {
		log.Panicf("cache.InitCache: %s", err.Error())
	}

	initAPI()

	ls, err := initSocketListener()
	if err != nil {
		log.Panicf("initSocketListener: %s", err.Error())
	}
	listener = ls

	slog.Info("run param", "work dir", workDir, "log dir", logDir, "socket path", socketPath)
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
