package executor

type common struct {
	Sequence string `json:"seq"`
	Code     int    `json:"cod"`
}

type authRequest struct {
	common
	Addr  string `json:"adr"`
	ID    int    `json:"id"`
	Force bool   `json:"foc"`
}

type authResponse struct {
	common
}

type echoRequest struct {
	common
	Message string `json:"msg"`
}

type echoResponse struct {
	common
	Message string `json:"msg"`
}

type translateRequest struct {
	common
	Source  int    `json:"source"`
	Target  int    `json:"target"`
	Message string `json:"msg"`
}

type translateResponse struct {
	common
	Message string `json:"msg"`
}
