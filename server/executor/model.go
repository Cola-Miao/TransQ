package executor

import (
	"net"
	"sync"
)

type handler func(tqc *transQClient, req any) (any, error)

const (
	methodAuth = iota
	methodEcho
	methodTranslate
)

var exec executor

type executor struct {
	mu        sync.RWMutex
	handle    map[method]handler
	name      map[method]string
	structure map[method]any
	conn      map[int]net.Conn
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
