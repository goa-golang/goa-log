package goa_log

import (
	"log"
	"testing"
)

func TestColor(t *testing.T) {

	log.Printf("\x1b[%dm[E] hello world 31: 红 \x1b[0m\n", 31)
	log.Printf("\x1b[%dm[I]hello world 30: 黑 \x1b[0m\n", 30)
	log.Printf("\x1b[%dm[D]hello world 31: 红 \x1b[0m\n", 31)
	log.Printf("\x1b[%dm[W]hello world 32: 绿 \x1b[0m\n", 32)
	log.Printf("\x1b[%dmhello world 33: 黄 \x1b[0m\n", 33)
	log.Printf("\x1b[%dmhello world 34: 蓝 \x1b[0m\n", 34)
	log.Printf("\x1b[%dmhello world 35: 紫 \x1b[0m\n", 35)
	log.Printf("\x1b[%dmhello world 36: 深绿 \x1b[0m\n", 36)
	log.Printf("\x1b[%dmhello world 37: 白色 \x1b[0m\n", 37)

	log.Printf("\x1b[%d;%dmhello world \x1b[0m 47: 白色 30: 黑 \n", 47, 30)
	log.Printf("\x1b[%d;%dmhello world \x1b[0m 46: 深绿 31: 红 \n", 46, 31)
	log.Printf("\x1b[%d;%dmhello world \x1b[0m 45: 紫   32: 绿 \n", 45, 32)
	log.Printf("\x1b[%d;%dmhello world \x1b[0m 44: 蓝   33: 黄 \n", 44, 33)
	log.Printf("\x1b[%d;%dmhello world \x1b[0m 43: 黄   34: 蓝 \n", 43, 34)
	log.Printf("\x1b[%d;%dmhello world \x1b[0m 42: 绿   35: 紫 \n", 42, 35)
	log.Printf("\x1b[%d;%dmhello world \x1b[0m 41: 红   36: 深绿 \n", 41, 36)
	log.Printf("\x1b[%d;%dmhello world \x1b[0m 40: 黑   37: 白色 \n", 40, 37)

}

func TestLog(t *testing.T) {
	Debug("测试输出语句")
	Info("Info")
	Warn("Warn")
	Error("Error")
}
