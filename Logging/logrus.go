package logging

// 日志系统

import (
	"TGPersonInfo/config"
	"bytes"
	"fmt"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type LogFormatter struct{}

var _log *logrus.Logger

func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var conf = config.GetConfig()
	var levelColor int
	// 等级颜色选择
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.PanicLevel, logrus.FatalLevel:
		levelColor = red
	default:
		levelColor = blue
	}

	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	// 自定义日期格式
	timeStamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		// 自定义文件路径
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)

		// 自定义输出格式
		fmt.Fprintf(b, "[%s][%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", conf.Log.Projectname, timeStamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		fmt.Fprintf(b, "[%s][%s] \x1b[%dm[%s]\x1b[0m %s\n", conf.Log.Projectname, timeStamp, levelColor, entry.Level, entry.Message)
	}
	return b.Bytes(), nil
}

func InitLogger() {
	var conf = config.GetConfig()
	mylog := logrus.New()                    // 新建一个实例
	mylog.SetOutput(os.Stdout)               // 设置输出类型
	mylog.SetReportCaller(conf.Log.Showline) // 开启返回函数名和行号
	mylog.SetFormatter(&LogFormatter{})      // 设置自定义的formatter

	level, err := logrus.ParseLevel(conf.Log.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	mylog.SetLevel(level) // 设置最低的level
	_log = mylog
}

func GetLog() *logrus.Logger {
	return _log
}
