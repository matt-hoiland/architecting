package logging

import (
	"strings"

	"github.com/sirupsen/logrus"
)

func LevelFromString(level string) logrus.Level {
	uplevel := strings.ToUpper(level)
	switch uplevel {
	case "PANIC":
		return logrus.PanicLevel
	case "FATAL":
		return logrus.FatalLevel
	case "ERROR":
		return logrus.ErrorLevel
	case "WARN":
		return logrus.WarnLevel
	case "INFO":
		return logrus.InfoLevel
	case "DEBUG":
		return logrus.DebugLevel
	case "TRACE":
		return logrus.TraceLevel
	default:
		logrus.Warn("Invalid level provided: " + level)
		return logrus.InfoLevel
	}
}
