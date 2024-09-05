package models

type AuthRequest struct {
	Addr  string `json:"adr"`
	ID    int    `json:"id"`
	Force bool   `json:"foc"`
}

type EchoRequest struct {
	Message string `json:"msg"`
}
