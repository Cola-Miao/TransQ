package uds

import "net"

type transQClient struct {
	Conn net.Conn
	Addr string
	ID   string
}
