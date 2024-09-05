package utils

import (
	"fmt"
	"net"
)

func DialSocketWithTimeout(addr string, timeout int) (net.Conn, error) {
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
