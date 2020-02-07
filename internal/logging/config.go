package logging

import (
    "os"
    "github.com/sirupsen/logrus"
)

func LoadLoggingLevel() logrus.Level {
    lvl := os.Getenv("LEVEL")

    switch lvl {
    case "panic":
        return logrus.PanicLevel
    case "fatal":
        return logrus.FatalLevel
    case "error":
        return logrus.ErrorLevel
    case "warn":
        return logrus.WarnLevel
    case "info":
        return logrus.InfoLevel
    case "debug":
        return logrus.DebugLevel
    case "trace":
        return logrus.TraceLevel
    }

    return logrus.InfoLevel
}
