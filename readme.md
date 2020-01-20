# 常规使用
```go
package main

import (
	logEx "github.com/lylib/go-log"
)

func main() {
	logEx.Debug("文本", 123, true)
	logEx.Debugf("文本[%s], 数字[%d], 布尔[%t]", "测试", 123, true)
	logEx.Error("文本", 123, true)
	logEx.Errorf("文本[%s], 数字[%d], 布尔[%t]", "测试", 123, true)
	// ...
}

```

# 定制打印规则
```go
package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"

	logEx "github.com/lylib/go-log"
)

func main() {
	/*
		Custom: 颜色/标签
		Tips:
			[数组下标对应日志等级]
			lvEmergency = 0
			lvAlert = 1
			lvCritical = 2
			lvError = 3
			lvWarning = 4
			lvNotice = 5
			lvInformational = 6
			lvDebug = 7
	*/
	logEx.Logger.SetColor([]int{37, 36, 35, 31, 33, 32, 34, 30})
	logEx.Logger.SetTag([]rune{'M', 'A', 'C', 'E', 'W', 'N', 'I', 'D'})

	/*
		Custom: 打印方法
		Tips:
			logEx.Logger.GetTag(lv)
			logEx.Logger.GetColor(lv)
	*/
	logEx.Logger.SetMsgPrintFunc(func(lv int, msg string) {
		// position
		_, file, line, ok := runtime.Caller(2)
		if !ok {
			file = "???"
			line = 0
		}
		position := file + ":" + strconv.Itoa(line)

		// info
		logInfo := fmt.Sprintf("%s [%c] [%s]", time.Now().Format("2006/01/02 15:04:05.000"), logEx.Logger.GetTag(lv), position)

		// msg
		logMsg := fmt.Sprintf("%c[%d;%dm\n%s\n%c[0m", 0x1B, 1, logEx.Logger.GetColor(lv), msg, 0x1B)

		// print
		os.Stdout.WriteString(logInfo + logMsg)
	})
	logEx.Logger.SetMsgGenFunc(func(format string, v ...interface{}) string {
		return fmt.Sprintf(format, v...)
	})
}
```