// Package alog is a tiny wrap of log module of beego
package alog

import (
	"strings"

	"github.com/astaxie/beego/logs"
)

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
	Log.Debug(generateFmtStr(len(v)), v...)
}

// Info ...
func Info(v ...interface{}) {
	Log.Info(generateFmtStr(len(v)), v...)
}

// Notice ...
func Notice(v ...interface{}) {
	Log.Notice(generateFmtStr(len(v)), v...)
}

// Warning ...
func Warning(v ...interface{}) {
	Log.Warning(generateFmtStr(len(v)), v...)
}

// Error ...
func Error(v ...interface{}) {
	Log.Error(generateFmtStr(len(v)), v...)
}

// Critical ...
func Critical(v ...interface{}) {
	Log.Critical(generateFmtStr(len(v)), v...)
}

// Alert ...
func Alert(v ...interface{}) {
	Log.Alert(generateFmtStr(len(v)), v...)
}

// Emergency ...
func Emergency(v ...interface{}) {
	Log.Emergency(generateFmtStr(len(v)), v...)
}

func generateFmtStr(n int) string {
	return strings.Repeat("%v ", n)
}
