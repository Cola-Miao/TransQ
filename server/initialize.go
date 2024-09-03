package main

import (
	"errors"
	"fmt"
	"log/slog"
	"net"
	"os"
	"path"
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
