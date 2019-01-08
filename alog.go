// Package alog is a tiny wrap of log module of beego
package alog

import (
	"strings"

	"fmt"
	"path"
	"runtime"
	"time"
)

const depth = 1

// Debug ...
func Debug(v ...interface{}) {
	_, file, line, _ := runtime.Caller(depth)
	fmt.Printf(generateFmtStr("D", file, line, len(v)), v...)
}

// Info ...
func Info(v ...interface{}) {
	_, file, line, _ := runtime.Caller(depth)
	fmt.Printf(generateFmtStr("I", file, line, len(v)), v...)
}

// Notice ...
func Notice(v ...interface{}) {
	_, file, line, _ := runtime.Caller(depth)
	fmt.Printf(generateFmtStr("N", file, line, len(v)), v...)
}

// Warning ...
func Warning(v ...interface{}) {
	_, file, line, _ := runtime.Caller(depth)
	fmt.Printf(generateFmtStr("W", file, line, len(v)), v...)
}

// Error ...
func Error(v ...interface{}) {
	_, file, line, _ := runtime.Caller(depth)
	fmt.Printf(generateFmtStr("E", file, line, len(v)), v...)
}

// Critical ...
func Critical(v ...interface{}) {
	_, file, line, _ := runtime.Caller(depth)
	fmt.Printf(generateFmtStr("C", file, line, len(v)), v...)
}

// Alert ...
func Alert(v ...interface{}) {
	_, file, line, _ := runtime.Caller(depth)
	fmt.Printf(generateFmtStr("A", file, line, len(v)), v...)
}

// Emergency ...
func Emergency(v ...interface{}) {
	_, file, line, _ := runtime.Caller(depth)
	fmt.Printf(generateFmtStr("M", file, line, len(v)), v...)
}


func generateFmtStr(levelFlag, file string, line, n int) string {
	_, filename := path.Split(file)
	dir := path.Dir(file)
	lastDir := path.Base(dir)

	formatStr := strings.Repeat("%v ", n)
	formatStr = time.Now().Format("2006/01/02 15:04:05.999") + fmt.Sprintf(" [%s] [%s/%s:%d] ", levelFlag, lastDir, filename, line) + formatStr + "\n"
	return formatStr
}
