package main

import (
	"errors"
	"fmt"
	"log/slog"
	"net"
	"os"
	"path"
	"runtime"
)

func initWorkDir() (string, error) {
	hd, err := os.UserHomeDir()
	if err != nil {
		slog.Warn("os.UserHomeDir", "error", err.Error())
		hd = "."
	}

	wd := path.Join(hd, folder)
	err = os.MkdirAll(wd, os.FileMode(0755))
	if err != nil {
		return "", fmt.Errorf("os.MkdirAll: %w", err)
	}

	return wd, nil
}

func initSocketListener() (net.Listener, error) {
	err := os.Remove(socketPath)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("os.Remove: %w", err)
	}

	ls, err := net.Listen("unix", socketPath)
	if err != nil {
		return nil, fmt.Errorf("net.Listen: %w", err)
	}

	return ls, nil
}

func initSlog() error {
	err := os.MkdirAll(logDir, os.FileMode(0755))
	if err != nil {
		return fmt.Errorf("os.MkdirAll: %w", err)
	}

	opt := slog.HandlerOptions{
		AddSource:   true,
		Level:       logLever,
		ReplaceAttr: nil,
	}

	fileName := path.Join(logDir, "transQ.log")
	fp, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.FileMode(0755))
	if err != nil {
		return fmt.Errorf("os.OpenFile: %w", err)
	}

	h := slog.NewTextHandler(fp, &opt)
	logger := slog.New(h)
	slog.SetDefault(logger)

	return nil
}

func initEnvWithGOOS() {
	switch runtime.GOOS {
	case "windows":
		logDir = os.Getenv("ProgramData") + "\\Logs"
	case "linux":
		logDir = workDir
	case "darwin":
	default:
		logDir = "./logs/"
	}
}
