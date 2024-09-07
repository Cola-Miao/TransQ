package thirdAPI

const (
	auto = iota
	zh
	en
	ja
)

var lcLanguageCode = map[int]string{
	auto: "auto",
	zh:   "zh",
	en:   "en",
	ja:   "ja",
}

type TransReq struct {
	Source  int
	Target  int
	Message string
}

type TransResp struct {
	Message string
	Error   error
}
