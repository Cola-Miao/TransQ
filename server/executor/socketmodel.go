package executor

type authRequest struct {
	Addr  string `json:"adr"`
	ID    int    `json:"id"`
	Force bool   `json:"foc"`
}

type echoRequest struct {
	Message string `json:"msg"`
}
