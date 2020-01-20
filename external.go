package logEx

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"
)

var Logger logger

func init() {
	Logger = new(basicLogger)
	Logger.SetMode(Mode.Dev)
	Logger.SetColor([]int{37, 36, 35, 31, 33, 32, 34, 30})
	Logger.SetTag([]rune{'M', 'A', 'C', 'E', 'W', 'N', 'I', 'D'})
	Logger.SetMsgPrintFunc(func(lv int, msg string) {
		// position
		_, file, line, ok := runtime.Caller(2)
		if !ok {
			file = "???"
			line = 0
		}
		position := file + ":" + strconv.Itoa(line)

		// info
		logInfo := fmt.Sprintf("%s [%c] [%s]", time.Now().Format("2006/01/02 15:04:05.000"), Logger.GetTag(lv), position)

		// msg
		logMsg := fmt.Sprintf("%c[%d;%dm\n%s\n%c[0m", 0x1B, 1, Logger.GetColor(lv), msg, 0x1B)

		// print
		os.Stdout.WriteString(logInfo + logMsg)
	})
	Logger.SetMsgGenFunc(func(format string, v ...interface{}) string {
		return fmt.Sprintf(format, v...)
	})
}

func Emergency(v ...interface{}) {
	Logger.Emergency(v...)
}

func Emergencyf(format string, v ...interface{}) {
	Logger.Emergencyf(format, v...)
}

func Alert(v ...interface{}) {
	Logger.Alert(v...)
}

func Alertf(format string, v ...interface{}) {
	Logger.Alertf(format, v...)
}

func Critical(v ...interface{}) {
	Logger.Critical(v...)
}

func Criticalf(format string, v ...interface{}) {
	Logger.Criticalf(format, v...)
}

func Error(v ...interface{}) {
	Logger.Error(v...)
}

func Errorf(format string, v ...interface{}) {
	Logger.Errorf(format, v...)
}

func Warn(v ...interface{}) {
	Logger.Warn(v...)
}

func Warnf(format string, v ...interface{}) {
	Logger.Warnf(format, v...)
}

func Notice(v ...interface{}) {
	Logger.Notice(v...)
}

func Noticef(format string, v ...interface{}) {
	Logger.Noticef(format, v...)
}

func Info(v ...interface{}) {
	Logger.Info(v...)
}

func Infof(format string, v ...interface{}) {
	Logger.Infof(format, v...)
}

func Debug(v ...interface{}) {
	Logger.Debug(v...)
}

func Debugf(format string, v ...interface{}) {
	Logger.Debugf(format, v...)
}
