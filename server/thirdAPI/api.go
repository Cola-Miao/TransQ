package thirdAPI

type api interface {
	Initialize(token string) error
	SendMessage(tq *TransReq) (tp *TransResp)
}

var Lingocloud lingocloud
