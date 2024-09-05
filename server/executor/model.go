package executor

import (
	"errors"
	. "github.com/Cola-Miao/TransQ/server/models"
)

type handler func(data string) error

const (
	methodEcho = iota + 1
	methodTranslate
)

var exec executor

var (
	errNoMethod = errors.New("no method")
)

type executor struct {
	handle map[Method]handler
	name   map[Method]string
}
