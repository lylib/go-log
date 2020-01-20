package logEx

const (
	lvEmergency = iota
	lvAlert
	lvCritical
	lvError
	lvWarning
	lvNotice
	lvInformational
	lvDebug
)

type ModeType uint8

var Mode = struct {
	Dev  ModeType
	Prod ModeType
}{
	Dev:  1,
	Prod: 2,
}

type logger interface {
	SetMode(mode ModeType)
	SetColor(color []int)
	SetTag(tag []rune)
	SetMsgPrintFunc(func(lv int, msg string))
	SetMsgGenFunc(func(format string, v ...interface{}) string)
	GetColor(lv int) int
	GetTag(lv int) rune

	Emergency(v ...interface{})
	Emergencyf(format string, v ...interface{})

	Alert(v ...interface{})
	Alertf(format string, v ...interface{})

	Critical(v ...interface{})
	Criticalf(format string, v ...interface{})

	Error(v ...interface{})
	Errorf(format string, v ...interface{})

	Warn(v ...interface{})
	Warnf(format string, v ...interface{})

	Notice(v ...interface{})
	Noticef(format string, v ...interface{})

	Info(v ...interface{})
	Infof(format string, v ...interface{})

	Debug(v ...interface{})
	Debugf(format string, v ...interface{})
}
