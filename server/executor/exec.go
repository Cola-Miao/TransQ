package executor

import (
	"fmt"
	"github.com/Cola-Miao/TransQ/server/format"
	. "github.com/Cola-Miao/TransQ/server/models"
	"log"
)

func init() {
	exec.init()
}

func (e *executor) init() {
	format.FuncStart("executor.init")
	defer format.FuncEnd("executor.init")

	e.handle = make(map[Method]handler)
	e.name = make(map[Method]string)

	e.register(methodEcho, echo, "echo")
	e.register(methodTranslate, translate, "translate")
}

func (e *executor) register(method Method, handle handler, name string) {
	if _, ok := e.handle[method]; ok {
		log.Panicf("e.handle has method: %d", method)
	}
	e.handle[method] = handle

	if _, ok := e.name[method]; ok {
		log.Panicf("e.name has method: %d", method)
	}
	e.name[method] = name
}

func (e *executor) do(info *Information) error {
	format.FuncStart("executor.do")
	defer format.FuncEnd("executor.do")

	if _, ok := e.handle[info.Method]; !ok {
		return errNoMethod
	}

	err := e.handle[info.Method](info.Data)
	if err != nil {
		return fmt.Errorf("method: %s: %w", e.name[info.Method], err)
	}

	return nil
}

func Do(info *Information) error {
	return exec.do(info)
}
