package log

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"

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

//---------------------------------------------------------------------------
func init() {
	logs.SetLogFuncCall(false)

	_user, err := user.Current()
	if err != nil {
		logs.Error(err.Error())
	}
	fileName := _user.HomeDir + "/.unichain-go"
	iniConfig, err := config.NewConfig("json", fileName)
	if err != nil {
		return
	}

	myBeegoLogAdapterMultiFile := &MyBeegoLogAdapterMultiFile{}
	myBeegoLogAdapterMultiFile.FileName = iniConfig.String("Log::LogName")
	myBeegoLogAdapterMultiFile.Level, _ = iniConfig.Int("Log::LogSaveLevel")
	logMaxDays, _ := iniConfig.Int("Log::LogMaxDays")
	myBeegoLogAdapterMultiFile.MaxDays = int16(logMaxDays)
	myBeegoLogAdapterMultiFile.MaxLines, _ = iniConfig.Int64("Log::LogMaxLines")
	myBeegoLogAdapterMultiFile.MaxSize, _ = iniConfig.Int64("Log::LogMaxSize")
	myBeegoLogAdapterMultiFile.Rotate, _ = iniConfig.Bool("Log::LogRotate")
	myBeegoLogAdapterMultiFile.Daily, _ = iniConfig.Bool("Log::LogDaily")
	myBeegoLogAdapterMultiFile.Separate = iniConfig.Strings("Log::LogSeparate")

	log_config := NewMyBeegoLogAdapterMultiFile(myBeegoLogAdapterMultiFile)
	log_config_str := _Serialize(log_config)
	//fmt.Println(log_config_str)

	// order 顺序必须按照
	// 1. logs.SetLevel(level)
	// 2. logs.SetLogger(logs.AdapterMultiFile, log_config_str)
	logLevel, _ := iniConfig.Int("Log::LogLevel")
	logs.SetLevel(logLevel)
	logs.SetLogger(logs.AdapterMultiFile, log_config_str)

	enableconsole, _ := iniConfig.Bool("Log::LogEnableConsole")
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
	f("[%s : %d (%s)] %s", slStr[0]+string(filepath.Separator)+filepath.Base(file), line, slStr[len(slStr)-1], _FormatLog(format, v...))
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
