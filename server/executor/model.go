package executor

import (
	"errors"
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

var (
	errNoMethod       = errors.New("no method")
	errIDNotExist     = errors.New("id not exist")
	errIDExist        = errors.New("id exist")
	errNoStructure    = errors.New("no structure")
	errNoHandler      = errors.New("no handler")
	errNoName         = errors.New("no name")
	errBadRequestType = errors.New("bad request type")
)

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
