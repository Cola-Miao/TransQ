package executor

import (
	"fmt"
	. "github.com/Cola-Miao/TransQ/server/config"
	"github.com/Cola-Miao/TransQ/server/format"
	. "github.com/Cola-Miao/TransQ/server/models"
	"github.com/Cola-Miao/TransQ/server/thirdAPI"
	"github.com/Cola-Miao/TransQ/server/utils"
)

func auth(tqc *transQClient, req *authRequest) error {
	format.FuncStart("auth")
	defer format.FuncEnd("auth")

	if req.Force {
		err := forceAuth(tqc, req)
		if err != nil {
			return fmt.Errorf("forceAuth: %w", err)
		}
		return nil
	}

	err := stdAuth(tqc, req)
	if err != nil {
		return fmt.Errorf("stdAuth: %w", err)
	}
	return nil
}

func forceAuth(tqc *transQClient, req *authRequest) error {
	conn, err := utils.DialSocketWithTimeout(req.Addr, Cfg.Listener.Timeout)
	if err != nil {
		return fmt.Errorf("utils.DialSocketWithTimeout: %w", err)
	}

	// don`t check id exist, cover old conn
	err = exec.setConnForce(req.ID, conn)
	if err != nil {
		return fmt.Errorf("exec.setConnForce: %w", err)
	}

	tqc.Addr = req.Addr
	tqc.ID = req.ID

	return nil
}

func stdAuth(tqc *transQClient, req *authRequest) error {
	// check id exist
	_, err := exec.getConn(req.ID)
	if err == nil {
		return ErrIDExist
	}

	conn, err := utils.DialSocketWithTimeout(req.Addr, Cfg.Listener.Timeout)
	if err != nil {
		return fmt.Errorf("utils.DialSocketWithTimeout: %w", err)
	}

	// set id-conn, double check
	err = exec.setConn(req.ID, conn)
	if err != nil {
		return fmt.Errorf("exec.setConn: %w", err)
	}

	tqc.Addr = req.Addr
	tqc.ID = req.ID

	return nil
}

func translate(req *translateRequest) (*translateResponse, error) {
	tp := thirdAPI.API[1].SendMessage(&thirdAPI.TransReq{
		Source:  req.Source,
		Target:  req.Target,
		Message: req.Message,
	})

	if tp.Error != nil {
		return nil, fmt.Errorf("SendMessage: %w", tp.Error)
	}

	fmt.Println("translate resp: ", tp.Message)
	return nil, nil
}
