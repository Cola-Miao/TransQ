package thirdAPI

const (
	Lingocloud = iota + 1
)

var API = map[int]api{
	Lingocloud: &lingocloud{},
}

type api interface {
	Initialize(token string) error
	SendMessage(tq *TransReq) (tp *TransResp)
}
