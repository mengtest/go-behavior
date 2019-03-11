package log

import "os"

var LogHandler = defaultLog{writer: os.Stdout}

func Info(i ...interface{}) {
	LogHandler.Info(i...)
}

func Infof(fmt string, i ...interface{}) {
	LogHandler.Infof(fmt, i...)
}

func Debug(i ...interface{}) {
	LogHandler.Debug(i...)
}

func Debugf(fmt string, i ...interface{}) {
	LogHandler.Debugf(fmt, i...)
}
