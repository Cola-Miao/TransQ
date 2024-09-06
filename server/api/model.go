package api

type TransReq struct {
	Source  int
	Target  int
	Message string
}

type TransResp struct {
	Message string
	err     error
}
