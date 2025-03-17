package utils

import (
	"fmt"
	"log"
	"runtime"
)

const (
	Debug = iota
	Warn
	Info
	Error
)

const (
	green   = "\033[97;42m"
	white   = "\033[90;47m"
	yellow  = "\033[90;43m"
	red     = "\033[97;41m"
	blue    = "\033[97;44m"
	magenta = "\033[97;45m"
	cyan    = "\033[97;46m"
	reset   = "\033[0m"
)

type (
	logLevel    int
	logInstance struct {
		level logLevel
		color bool
	}
)

var (
	l = &logInstance{
		level: Info,
		color: true,
	}
)

func SetLevel(setLevel logLevel) {
	l.level = setLevel
}

func Infof(format string, args ...interface{}) {
	if l.level <= Info {
		output(green, "INFO", format, args...)
	}
}

func Debugf(format string, args ...interface{}) {
	if l.level == Debug {
		output(magenta, "DEBUG", format, args...)
	}
}

func Warnf(format string, args ...interface{}) {
	if l.level <= Warn {
		output(yellow, "WARN", format, args...)
	}
}

func Errorf(format string, args ...interface{}) {
	output(red, "ERROR", format, args...)
}

func output(color, level, format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(2)
	if l.color {
		log.Printf(fmt.Sprintf("[%s%s%s]\t%s:%d\t|", color, level, reset, file, line)+format, args...)
	} else {
		log.Printf(fmt.Sprintf("[%s]\t%s:%d\t|", level, file, line)+format, args...)
	}
}
