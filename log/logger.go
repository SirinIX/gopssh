package log

import (
	"fmt"
	"os"
	"time"
)

const (
	format = "2006-01-02 15:04:05"
)

func Info(format string, args ...interface{}) {
	fmt.Fprintf(os.Stdout, timeLog()+format+"\n", args...)
}

func Error(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, timeLog()+format+"\n", args...)
}

func timeLog() string {
	return fmt.Sprintf("[%v] ", time.Now().Format(format))
}
