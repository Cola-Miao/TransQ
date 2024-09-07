package models

import "errors"

var (
	ErrNoMethod            = errors.New("no method")
	ErrIDNotExist          = errors.New("id not exist")
	ErrIDExist             = errors.New("id exist")
	ErrNoStructure         = errors.New("no structure")
	ErrNoHandler           = errors.New("no handler")
	ErrNoName              = errors.New("no name")
	ErrBadRequestType      = errors.New("bad request type")
	ErrUnsupportedLanguage = errors.New("unsupported language")
	ErrAssertionType       = errors.New("wrong assertion type")
	ErrAPINotExist         = errors.New("api not exist")
)
