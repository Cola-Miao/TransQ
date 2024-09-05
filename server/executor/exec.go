package executor

import (
	"fmt"
	"github.com/Cola-Miao/TransQ/server/format"
	"log"
	"net"
)

func init() {
	exec.init()
}

func (e *executor) init() {
	format.FuncStart("executor.init")
	defer format.FuncEnd("executor.init")

	e.handle = make(map[method]handler)
	e.name = make(map[method]string)
	e.conn = make(map[int]net.Conn)

	e.register(methodAuth, mtdAuth, "auth")
	e.register(methodEcho, mtdEcho, "echo")
	e.register(methodTranslate, mtdTranslate, "translate")
}

func (e *executor) register(method method, handle handler, name string) {
	if _, ok := e.handle[method]; ok {
		log.Panicf("e.handle has method: %d", method)
	}
	e.handle[method] = handle

	if _, ok := e.name[method]; ok {
		log.Panicf("e.name has method: %d", method)
	}
	e.name[method] = name
}

func (e *executor) do(tqc *transQClient) error {
	format.FuncStartWithData("executor.do", tqc)
	defer format.FuncEnd("executor.do")

	if _, ok := e.handle[tqc.Info.Method]; !ok {
		return errNoMethod
	}

	err := e.handle[tqc.Info.Method](tqc)
	if err != nil {
		return fmt.Errorf("method: %s: %w", e.name[tqc.Info.Method], err)
	}

	return nil
}

func (e *executor) getConn(id int) (net.Conn, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()

	conn, ok := e.conn[id]
	if !ok {
		return nil, errIDNotExist
	}
	return conn, nil
}

func (e *executor) setConn(id int, conn net.Conn) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	_, ok := e.conn[id]
	if ok {
		return errIDExist
	}

	e.conn[id] = conn
	return nil
}

func (e *executor) setConnForce(id int, conn net.Conn) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.conn[id] = conn
	return nil
}
