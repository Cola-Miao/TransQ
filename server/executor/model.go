package executor

import (
	"errors"
	"net"
)

type handler func(tqc *transQClient) error

const (
	methodEcho = iota + 1
	methodTranslate
)

var exec executor

var (
	errNoMethod = errors.New("no method")
)

type executor struct {
	handle map[method]handler
	name   map[method]string
}

type transQClient struct {
	Conn net.Conn
	Info *information
	Addr string
	ID   string
}

type method uint8

type information struct {
	Method method `json:"mtd"`
	Data   string `json:"dat"`
}
