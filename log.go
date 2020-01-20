package logEx

import (
	"strings"
)

type basicLogger struct {
	mode     ModeType
	lvColor  []int
	lvTag    []rune
	msgPrint func(lv int, msg string)
	msgGen   func(format string, v ...interface{}) string
}

// setting
func (this *basicLogger) SetMode(mode ModeType) {
	this.mode = mode
}

func (this *basicLogger) SetColor(color []int) {
	this.lvColor = color
}

func (this *basicLogger) SetTag(tag []rune) {
	this.lvTag = tag
}

func (this *basicLogger) SetMsgPrintFunc(f func(lv int, msg string)) {
	this.msgPrint = f
}

func (this *basicLogger) SetMsgGenFunc(f func(format string, v ...interface{}) string) {
	this.msgGen = f
}

func (this *basicLogger) GetColor(lv int) int {
	return this.lvColor[lv]
}

func (this *basicLogger) GetTag(lv int) rune {
	return this.lvTag[lv]
}

// operation
func (this *basicLogger) Emergency(v ...interface{}) {
	fmtStr := strings.Repeat("%v ", len(v))
	this.msgPrint(lvEmergency, this.msgGen(fmtStr, v...))
}

func (this *basicLogger) Emergencyf(fmtStr string, v ...interface{}) {
	this.msgPrint(lvEmergency, this.msgGen(fmtStr, v...))
}

func (this *basicLogger) Alert(v ...interface{}) {
	fmtStr := strings.Repeat("%v ", len(v))
	this.msgPrint(lvAlert, this.msgGen(fmtStr, v...))
}

func (this *basicLogger) Alertf(fmtStr string, v ...interface{}) {
	this.msgPrint(lvAlert, this.msgGen(fmtStr, v...))
}

func (this *basicLogger) Critical(v ...interface{}) {
	fmtStr := strings.Repeat("%v ", len(v))
	this.msgPrint(lvCritical, this.msgGen(fmtStr, v...))
}

func (this *basicLogger) Criticalf(fmtStr string, v ...interface{}) {
	this.msgPrint(lvCritical, this.msgGen(fmtStr, v...))
}

func (this *basicLogger) Error(v ...interface{}) {
	fmtStr := strings.Repeat("%v ", len(v))
	this.msgPrint(lvError, this.msgGen(fmtStr, v...))
}

func (this *basicLogger) Errorf(fmtStr string, v ...interface{}) {
	this.msgPrint(lvError, this.msgGen(fmtStr, v...))
}

func (this *basicLogger) Warn(v ...interface{}) {
	fmtStr := strings.Repeat("%v ", len(v))
	this.msgPrint(lvWarning, this.msgGen(fmtStr, v...))
}

func (this *basicLogger) Warnf(fmtStr string, v ...interface{}) {
	this.msgPrint(lvWarning, this.msgGen(fmtStr, v...))
}

func (this *basicLogger) Notice(v ...interface{}) {
	fmtStr := strings.Repeat("%v ", len(v))
	this.msgPrint(lvNotice, this.msgGen(fmtStr, v...))
}

func (this *basicLogger) Noticef(fmtStr string, v ...interface{}) {
	this.msgPrint(lvNotice, this.msgGen(fmtStr, v...))
}

func (this *basicLogger) Info(v ...interface{}) {
	fmtStr := strings.Repeat("%v ", len(v))
	this.msgPrint(lvInformational, this.msgGen(fmtStr, v...))
}

func (this *basicLogger) Infof(fmtStr string, v ...interface{}) {
	this.msgPrint(lvInformational, this.msgGen(fmtStr, v...))
}

func (this *basicLogger) Debug(v ...interface{}) {
	if this.mode == Mode.Prod {
		return
	}
	fmtStr := strings.Repeat("%v ", len(v))
	this.msgPrint(lvDebug, this.msgGen(fmtStr, v...))
}

func (this *basicLogger) Debugf(fmtStr string, v ...interface{}) {
	if this.mode == Mode.Prod {
		return
	}
	this.msgPrint(lvDebug, this.msgGen(fmtStr, v...))
}
