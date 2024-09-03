package format

import "log"

func FuncStart(funcName string) {
	log.Printf("%s start\n", funcName)
}

func FuncEnd(funcName string) {
	log.Printf("%s end\n", funcName)
}

func FuncStartWithData(funcName string, data any) {
	log.Printf("%s start with data: %+v\n", funcName, data)
}

func FuncEndWithData(funcName string, data any) {
	log.Printf("%s end with data: %+v\n", funcName, data)
}
