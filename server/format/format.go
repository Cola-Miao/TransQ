package format

import (
	"log/slog"
)

func FuncStart(funcName string) {
	slog.Debug(funcName, "status", "start")
}

func FuncEnd(funcName string) {
	slog.Debug(funcName, "status", "end")
}

func FuncStartWithData(funcName string, data any) {
	slog.Debug(funcName, "status", "start", "data", data)
}

func FuncEndWithData(funcName string, data any) {
	slog.Debug(funcName, "status", "end", "data", data)
}
