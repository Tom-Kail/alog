// Package alog is a tiny wrap of log module of beego
package alog

import "github.com/astaxie/beego/logs"

// Log global var
var Log = logs.NewLogger(10000)

func init() {
	Log.SetLogger("console", "")
	Log.EnableFuncCallDepth(true)
	Log.SetLogFuncCallDepth(3)
	Log.SetLevel(logs.LevelDebug)

}

//	LevelEmergency
//	LevelAlert
//	LevelCritical
//	LevelError
//	LevelWarning
//	LevelNotice
//	LevelInformational
//	LevelDebug

// Debug ...
func Debug(v ...interface{}) {
	Log.Debug("%v", v)
}

// Info ...
func Info(v ...interface{}) {
	Log.Info("%v", v)
}

// Notice ...
func Notice(v ...interface{}) {
	Log.Notice("%v", v)
}

// Warning ...
func Warning(v ...interface{}) {
	Log.Warning("%v", v)
}

// Error ...
func Error(v ...interface{}) {
	Log.Error("%v", v)
}

// Critical ...
func Critical(v ...interface{}) {
	Log.Critical("%v", v)
}

// Alert ...
func Alert(v ...interface{}) {
	Log.Alert("%v", v)
}

// Emergency ...
func Emergency(v ...interface{}) {
	Log.Emergency("%v", v)
}
