package utils

import (
	"fmt"
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
		conn.Close()
		return nil, fmt.Errorf("SetDeadline: %w", err)
	}

	return conn, nil
}
