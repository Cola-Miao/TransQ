package executor

import (
	"fmt"
	"github.com/Cola-Miao/TransQ/server/format"
	"github.com/Cola-Miao/TransQ/server/models"
	"net"
)

func auth(tqc *transQClient, req *models.AuthRequest) error {
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

func forceAuth(tqc *transQClient, req *models.AuthRequest) error {
	conn, err := net.Dial("unix", req.Addr)
	if err != nil {
		return fmt.Errorf("net.Dial: %w", err)
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

func stdAuth(tqc *transQClient, req *models.AuthRequest) error {
	// check id exist
	_, err := exec.getConn(req.ID)
	if err == nil {
		return errIDExist
	}

	conn, err := net.Dial("unix", req.Addr)
	if err != nil {
		return fmt.Errorf("net.Dial: %w", err)
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
