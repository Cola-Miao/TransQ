package format

import "log"

func FuncStart(s string) {
	log.Printf("%s start\n", s)
}

func FuncEnd(s string) {
	log.Printf("%s end\n", s)
}
