package log

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

func InitLogger() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logrus.SetOutput(os.Stdout)

	logrus.SetLevel(logrus.InfoLevel)
}

func Debug(format string, args ...interface{}) {
	logrus.Debug(fmt.Sprintf(format, args...))
}

func Info(format string, args ...interface{}) {
	logrus.Info(fmt.Sprintf(format, args...))
}

func Warning(format string, args ...interface{}) {
	logrus.Warning(fmt.Sprintf(format, args...))
}

func Error(format string, args ...interface{}) {
	logrus.Error(fmt.Sprintf(format, args...))
}
