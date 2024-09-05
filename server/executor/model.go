package executor

import (
	"errors"
	"net"
)

type handler func(tqc *transQClient) error

const (
	methodAuth = iota
	methodEcho
	methodTranslate
)

var exec executor

var (
	errNoMethod = errors.New("no method")
)

type executor struct {
	handle map[method]handler
	name   map[method]string
	conn   map[int]net.Conn
}

type transQClient struct {
	Conn net.Conn
	Info *information
	Addr string
	ID   int
}

type method uint8

type information struct {
	Method method `json:"mtd"`
	Data   string `json:"dat"`
}
