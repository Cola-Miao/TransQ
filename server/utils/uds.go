package utils

import (
	"fmt"
	"log/slog"
	"net"
	"time"
)

func DialSocketWithTimeout(addr string, timeout time.Duration) (net.Conn, error) {
	conn, err := net.Dial("unix", addr)
	if err != nil {
		return nil, fmt.Errorf("net.Dial: %w", err)
	}

	err = conn.SetDeadline(GetOutTime(timeout))
	if err != nil {
		der := conn.Close()
		if der != nil {
			slog.Error("conn.Close", "error", err.Error())
		}
		return nil, fmt.Errorf("SetDeadline: %w", err)
	}

	return conn, nil
}

func AcceptSocketWithTimeout(ls net.Listener, timeout time.Duration) (net.Conn, error) {
	conn, err := ls.Accept()
	if err != nil {
		return nil, fmt.Errorf("ls.Accept: %w", err)
	}

	err = conn.SetDeadline(GetOutTime(timeout))
	if err != nil {
		return nil, fmt.Errorf("SetDeadline: %w", err)
	}

	return conn, nil
}
