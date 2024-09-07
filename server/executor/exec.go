package executor

import (
	"encoding/json"
	"fmt"
	"github.com/Cola-Miao/TransQ/server/format"
	. "github.com/Cola-Miao/TransQ/server/models"
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
	e.structure = make(map[method]any)
	e.conn = make(map[int]net.Conn)

	e.register(methodAuth, mtdAuth, "auth", &authRequest{})
	e.register(methodEcho, mtdEcho, "echo", &echoRequest{})
	e.register(methodTranslate, mtdTranslate, "translate", &translateRequest{})
}

func (e *executor) register(method method, handle handler, name string, structure any) {
	if _, ok := e.handle[method]; ok {
		log.Panicf("e.handle has method: %d", method)
	}
	e.handle[method] = handle

	if _, ok := e.name[method]; ok {
		log.Panicf("e.name has method: %d", method)
	}
	e.name[method] = name

	if _, ok := e.structure[method]; ok {
		log.Panicf("e.structure has method: %d", method)
	}
	e.structure[method] = structure
}

func (e *executor) do(tqc *transQClient) error {
	format.FuncStartWithData("executor.do", tqc)
	defer format.FuncEnd("executor.do")

	mtd := tqc.Info.Method

	name, str, hdl, err := e.mtdCheck(mtd)
	if err != nil {
		return fmt.Errorf("mtdCheck: %w", err)
	}

	err = json.Unmarshal([]byte(tqc.Info.Data), &str)
	if err != nil {
		return fmt.Errorf("json.Unmarshal: %w", err)
	}

	resp, err := hdl(tqc, str)
	if err != nil {
		return fmt.Errorf("method: %s: %w", name, err)
	}

	err = e.writeConn(tqc.ID, resp)
	if err != nil {
		return fmt.Errorf("e.writeConn: %w", err)
	}

	return nil
}

func (e *executor) getConn(id int) (net.Conn, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()

	conn, ok := e.conn[id]
	if !ok {
		return nil, ErrIDNotExist
	}
	return conn, nil
}

func (e *executor) setConn(id int, conn net.Conn) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	_, ok := e.conn[id]
	if ok {
		return ErrIDExist
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

func (e *executor) writeConn(id int, resp any) error {
	conn, err := e.getConn(id)
	if err != nil {
		return fmt.Errorf("getConn: %w", err)
	}

	data, err := json.Marshal(resp)
	if err != nil {
		return fmt.Errorf("json.Marshal: %w", err)
	}

	_, err = conn.Write(data)
	if err != nil {
		return fmt.Errorf("conn.Write: %w", err)
	}

	return nil
}

func (e *executor) mtdCheck(mtd method) (name string, structure any, handle handler, err error) {
	if _, ok := e.handle[mtd]; !ok {
		err = ErrNoMethod
		return
	}

	name, ok := e.name[mtd]
	if !ok {
		err = ErrNoName
		return
	}

	structure, ok = e.structure[mtd]
	if !ok {
		err = ErrNoStructure
		return
	}

	handle, ok = e.handle[mtd]
	if !ok {
		err = ErrNoHandler
		return
	}

	return
}
