package goa_log

import "log"

// 输出测试信息
func Debug(v string) {
	log.Printf("\x1b[%dmDebug %s\x1b[0m\n", 34, v)
}

// 输出提示信息
func Info(v string) {
	log.Printf("\x1b[%dmInfo %s\x1b[0m\n", 30, v)
}

// 输出警告信息
func Warn(v string) {
	log.Printf("\x1b[%dmWarn %s\x1b[0m\n", 33, v)
}

// 输出错误信息
func Error(v string) {
	log.Printf("\x1b[%dmError %s\x1b[0m\n", 31, v)
}
