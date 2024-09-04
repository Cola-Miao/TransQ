package format

import (
	"log/slog"
)

func FuncStart(funcName string) {
	slog.Debug("start", "func", funcName)
}

func FuncEnd(funcName string) {
	slog.Debug("end", "func", funcName)
}

func FuncStartWithData(funcName string, data any) {
	slog.Debug("start", "func", funcName, "data", data)
}

func FuncEndWithData(funcName string, data any) {
	slog.Debug("end", "func", funcName, "data", data)
}
