package config

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

func InitLog() {
	logType := os.Getenv("LOG_TYPE")
	logLevel := os.Getenv("LOG_LEVEL")
	log.SetLevel(getLoggerLevel(logLevel))
	log.SetReportCaller(true)
	if strings.ToUpper(logType) == "JSON" {
		log.SetFormatter(&log.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
	}
	if strings.ToUpper(logType) == "NESTED" {
		log.SetFormatter(&nested.Formatter{
			HideKeys:        true,
			FieldsOrder:     []string{"component", "category"},
			TimestampFormat: "2006-01-02 15:04:05",
			ShowFullLevel:   true,
			CallerFirst:     true,
		})
	}
}

func getLoggerLevel(value string) log.Level {
	switch value {
	case "TRACE":
		return log.TraceLevel
	case "DEBUG":
		return log.DebugLevel
	default:
		return log.InfoLevel
	}
}
