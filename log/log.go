package log

import (
	"fmt"
	"io"
	"path"
	"runtime"
	"strings"
	"time"
)

var Logger interface {
	Info(i ...interface{})
	Infof(fmt string, i ...interface{})
	Debug(i ...interface{})
	Debugf(fmt string, i ...interface{})
}

type defaultLog struct {
	writer io.Writer
}

func (this *defaultLog) Info(i ...interface{}) {
	this.log("I", strings.Repeat("%v", len(i)), i...)
}

func (this *defaultLog) Infof(fmt string, i ...interface{}) {
	this.log("I", fmt, i...)
}

func (this *defaultLog) Debug(i ...interface{}) {
	this.log("D", strings.Repeat("%v", len(i)), i...)
}

func (this *defaultLog) Debugf(fmt string, i ...interface{}) {
	this.log("D", fmt, i...)
}

func (this *defaultLog) log(typ string, fmt_ string, i ...interface{}) {
	// [时间][类型][文件名:行号] 日志
	logFmt := "[%s][%s][%s:%d] %s"
	when := time.Now()

	_, file, line, ok := runtime.Caller(3)
	if !ok {
		file = "???"
		line = 0
	}
	_, filename := path.Split(file)
	m := fmt.Sprintf(fmt_, i...)
	msg := fmt.Sprintf(logFmt, when.Format("2006-01-02 15:04:05"), typ, filename, line, m)
	this.writer.Write([]byte(msg))
}
