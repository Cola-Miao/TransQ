package models

type AuthRequest struct {
	Addr string `json:"adr"`
	ID   int    `json:"id"`
}

type EchoRequest struct {
	Message string `json:"msg"`
}
