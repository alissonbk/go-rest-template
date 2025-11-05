package config

import (
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)


func InitLog() {
	lvl, err := logrus.ParseLevel("debug")
	if err != nil {
		panic("failed to parse logrus log level" + err.Error())
	}

	log.SetLevel(lvl)
	log.SetReportCaller(true)
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	// log.SetFormatter(&nested.Formatter{
	// 		HideKeys:        true,
	// 		FieldsOrder:     []string{"component", "category"},
	// 		TimestampFormat: "2006-01-02 15:04:05",
	// 		ShowFullLevel:   true,
	// 		CallerFirst:     true,
	// 	})

}

// package config

// import (
// 	"sync"

// 	"go.uber.org/zap"
// 	"go.uber.org/zap/zapcore"
// )

// var (
// 	logger *zap.Logger
// 	once   sync.Once
// )

// func InitLogger() {
// 	once.Do(func() {
// 		config := zap.Config{
// 			Level:       zap.NewAtomicLevelAt(zap.FatalLevel),
// 			Development: false,
// 			Encoding:    "json",
// 			EncoderConfig: zapcore.EncoderConfig{
// 				TimeKey:        "timestamp",
// 				LevelKey:       "level",
// 				NameKey:        "logger",
// 				CallerKey:      "caller",
// 				MessageKey:     "message",
// 				StacktraceKey:  "stacktrace",
// 				LineEnding:     zapcore.DefaultLineEnding,
// 				EncodeLevel:    zapcore.LowercaseLevelEncoder,
// 				EncodeTime:     zapcore.ISO8601TimeEncoder,
// 				EncodeDuration: zapcore.SecondsDurationEncoder,
// 				EncodeCaller:   zapcore.ShortCallerEncoder,
// 			},
// 			OutputPaths: []string{
// 				"stdout",
// 				//"/var/log/myapp/app.log"
// 			},
// 			ErrorOutputPaths: []string{"stderr"},
// 		}
// 		logger, err := config.Build()
// 		if err != nil {
// 			panic("failed to initialize zap logger, cause: " + err.Error())
// 		}
// 		// flushes buffer, if any
// 		logger.Sync()
// 	})
// }

// func Logger() *zap.Logger {
// 	if logger == nil {
// 		panic("the logger have not been initialized yet!")
// 	}

// 	return logger
// }
