package goalog

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"runtime/debug"
)

// 输出测试信息
func Debug(v ...interface{}) {
	log.Printf("\x1b[%dmDebug %s\x1b[0m\n", 34, toJsonString(v))
}

// 输出提示信息
func Info(v ...interface{}) {
	log.Printf("\x1b[%dmInfo %s\x1b[0m\n", 30, toJsonString(v))
}

// 输出警告信息
func Warn(v ...interface{}) {
	log.Printf("\x1b[%dmWarn %s\x1b[0m\n", 33, toJsonString(v))
}

// 输出错误信息
func Fatal(v ...interface{}) {

	str := toJsonString(v)
	debug.PrintStack()
	log.Fatal(fmt.Sprintf("\x1b[%dmError %s\x1b[0m\n", 31, str))
}

func toJsonString(vs ...interface{}) string {
	s := ""
	for i, v := range vs {
		kind := reflect.TypeOf(reflect.ValueOf(v)).Kind().String()
		if kind == "struct" || kind == "ptr" {
			d, _ := json.Marshal(vs[i])
			s += " " + string(d)
		} else {
			s += fmt.Sprint(vs[i])
		}
	}

	return fmt.Sprintf(s)
}
