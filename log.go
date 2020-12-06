package goa_log

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"runtime/debug"
)

// 输出测试信息
func Debug(v interface{}) {
	log.Printf("\x1b[%dmDebug %s\x1b[0m\n", 34, toJsonStringWithNamed(v))
}

// 输出提示信息
func Info(v interface{}) {
	log.Printf("\x1b[%dmInfo %s\x1b[0m\n", 30, toJsonStringWithNamed(v))
}

// 输出警告信息
func Warn(v interface{}) {
	log.Printf("\x1b[%dmWarn %s\x1b[0m\n", 33, toJsonStringWithNamed(v))
}

// 输出错误信息
func Fatal(v interface{}) {
	str := toJsonStringWithNamed(v)
	log.Fatal(fmt.Sprintf("\x1b[%dmError %s\x1b[0m\n", 31, fmt.Sprint(str, " ", toJsonString(string(debug.Stack()), "Stack"))))
}

func toJsonString(v interface{}, name string) string {
	kind := reflect.TypeOf(v).Kind().String()
	if kind == "struct" || kind == "ptr" {
		d, _ := json.Marshal(v)
		return string(d)
	}
	return fmt.Sprintf(`{"%s":"%s"}`, name, fmt.Sprint(v))
}

func toJsonStringWithNamed(v interface{}) string {
	return toJsonString(v, "Data")
}
