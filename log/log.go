package log

import (
	"bytes"
	"encoding/json"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

import (
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
)

//---------------------------------------------------------------------------
const (
	LevelError = iota
	LevelWarn
	LevelInfo
	LevelDebug
)

var mapLevelKeys = map[int]string{
	LevelError: "ERROR",
	LevelWarn:  "WARN",
	LevelInfo:  "INFO",
	LevelDebug: "DEBUG",
}

var (
	confPath = "log/log.conf"
)

//---------------------------------------------------------------------------
func init() {
	logs.SetLogFuncCall(false)

	iniConfig, err := config.NewConfig("ini", confPath)
	if err != nil {
		panic(err)
	}

	myBeegoLogAdapterMultiFile := &MyBeegoLogAdapterMultiFile{}
	myBeegoLogAdapterMultiFile.FileName = iniConfig.String("log::LogName")
	myBeegoLogAdapterMultiFile.Level, _ = iniConfig.Int("log::LogSaveLevel")
	logMaxDays, _ := iniConfig.Int("log::LogMaxDays")
	myBeegoLogAdapterMultiFile.MaxDays = int16(logMaxDays)
	myBeegoLogAdapterMultiFile.MaxLines, _ = iniConfig.Int64("log::LogMaxLines")
	myBeegoLogAdapterMultiFile.MaxSize, _ = iniConfig.Int64("log::LogMaxSize")
	myBeegoLogAdapterMultiFile.Rotate, _ = iniConfig.Bool("log::LogRotate")
	myBeegoLogAdapterMultiFile.Daily, _ = iniConfig.Bool("log::LogDaily")
	myBeegoLogAdapterMultiFile.Separate = iniConfig.Strings("log::LogSeparate")

	log_config := NewMyBeegoLogAdapterMultiFile(myBeegoLogAdapterMultiFile)
	log_config_str := _Serialize(log_config)
	//fmt.Println(log_config_str)

	// order 顺序必须按照
	// 1. logs.SetLevel(level)
	// 2. logs.SetLogger(logs.AdapterMultiFile, log_config_str)
	logLevel, _ := iniConfig.Int("log::LogLevel")
	logs.SetLevel(logLevel)
	logs.SetLogger(logs.AdapterMultiFile, log_config_str)

	enableconsole, _ := iniConfig.Bool("log::LogEnableConsole")
	if enableconsole {
		logs.SetLogger(logs.AdapterConsole)
	}
}

//---------------------------------------------------------------------------
func _Serialize(obj interface{}, escapeHTML ...bool) string {
	setEscapeHTML := false
	if len(escapeHTML) >= 1 {
		setEscapeHTML = escapeHTML[0]
	}
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	// disabled the HTMLEscape for &, <, and > to \u0026, \u003c, and \u003e in json string
	enc.SetEscapeHTML(setEscapeHTML)
	err := enc.Encode(obj)
	if err != nil {
		logs.Error(err.Error())
		return ""
	}
	return strings.TrimSpace(buf.String())
}

//---------------------------------------------------------------------------
func _FormatLog(f interface{}, v ...interface{}) string {
	var msg string
	switch f.(type) {
	case string:
		msg = f.(string)
		if len(v) == 0 {
			return msg
		}
		if strings.Contains(msg, "%") && !strings.Contains(msg, "%%") {
			//format string
		} else {
			//do not contain format char
			msg += strings.Repeat(" %v", len(v))
		}
	default:
		msg = fmt.Sprint(f)
		if len(v) == 0 {
			return msg
		}
		msg += strings.Repeat(" %v", len(v))
	}
	return fmt.Sprintf(msg, v...)
}

//---------------------------------------------------------------------------
func _WriteLog(key int, format interface{}, v ...interface{}) {
	pc, file, line, _ := runtime.Caller(2)
	func_ := runtime.FuncForPC(pc)
	var f func(f interface{}, v ...interface{})
	switch key {
	case LevelError:
		f = logs.Error
	case LevelWarn:
		f = logs.Warn
	case LevelInfo:
		f = logs.Info
	case LevelDebug:
		f = logs.Debug
	}
	slStr := strings.Split(func_.Name(), ".")
	f("[%s] [%s : %d (%s)] %s", mapLevelKeys[key],
		slStr[0]+string(filepath.Separator)+filepath.Base(file), line, slStr[len(slStr)-1], _FormatLog(format, v...))
}

//---------------------------------------------------------------------------
func Error(f interface{}, v ...interface{}) {
	_WriteLog(LevelError, f, v...)
}

//---------------------------------------------------------------------------
func Warn(f interface{}, v ...interface{}) {
	_WriteLog(LevelWarn, f, v...)
}

//---------------------------------------------------------------------------
func Info(f interface{}, v ...interface{}) {
	_WriteLog(LevelInfo, f, v...)
}

//---------------------------------------------------------------------------
func Debug(f interface{}, v ...interface{}) {
	_WriteLog(LevelDebug, f, v...)
}

//---------------------------------------------------------------------------
