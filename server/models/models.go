package models

import "net"

type Method uint8

type Information struct {
	Method Method `json:"mtd"`
	Data   string `json:"dat"`
}

type TransQClient struct {
	Conn net.Conn
	Addr string
	ID   string
}
