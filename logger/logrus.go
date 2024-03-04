package logger

import (
	"fmt"
	"io"
	internalLog "log"
	"os"

	"github.com/manicar2093/winter/stages"
	"github.com/sirupsen/logrus"
)

var (
	createBlackBox = func(stage string) (*os.File, error) {
		file, err := os.OpenFile(fmt.Sprintf("black_box_%s.log", stage), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Info("Failed to log to file, using default stderr")
			return nil, err
		}
		return file, err
	}
	initializers = map[string]func(l *logrus.Logger) error{
		stages.Test: func(l *logrus.Logger) error {
			file, err := createBlackBox(stages.Test)
			if err != nil {
				return err
			}
			loggerOutput := io.MultiWriter(file, os.Stdout)
			l.SetOutput(loggerOutput)
			internalLog.SetOutput(loggerOutput)
			l.SetFormatter(&logrus.JSONFormatter{
				PrettyPrint: true,
			})
			l.SetReportCaller(true)
			l.SetLevel(logrus.DebugLevel)
			return nil
		},
		stages.Dev: func(l *logrus.Logger) error {
			file, err := createBlackBox(stages.Dev)
			if err != nil {
				return err
			}
			loggerOutput := io.MultiWriter(file, os.Stdout)
			l.SetOutput(loggerOutput)
			internalLog.SetOutput(loggerOutput)
			l.SetFormatter(&logrus.JSONFormatter{
				PrettyPrint: true,
			})
			l.SetReportCaller(true)
			l.SetLevel(logrus.DebugLevel)
			return nil
		},
		stages.Prod: func(l *logrus.Logger) error {
			l.SetOutput(os.Stdout)
			internalLog.SetOutput(os.Stdout)
			l.SetFormatter(&logrus.JSONFormatter{})
			l.SetReportCaller(true)
			return nil
		},
	}

	log *Logger
)

type (
	Logger struct {
		*logrus.Logger
	}
)

func GetLogger() *Logger {
	return log
}

func init() {
	logger := logrus.New()
	currentStage, err := stages.GetCurrentStage()
	if err != nil {
		panic(err)
	}
	if err := initializers[currentStage](logger); err != nil {
		panic(err)
	}
	log = &Logger{Logger: logger}
}
