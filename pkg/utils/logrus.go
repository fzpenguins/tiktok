package utils

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var Logrus *logrus.Logger

func InitLogrus() {
	if Logrus != nil {
		src, _ := setOutput()
		Logrus.Out = src
		return
	}

	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	src, err := setOutput()
	if err != nil {
		panic(err)
	}
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logger.Out = src
	Logrus = logger
}

func setOutput() (*os.File, error) {

	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/log/"
	}
	_, err := os.Stat(logFilePath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(logFilePath, 0777)
		if err != nil {
			return nil, err
		}
	}
	logFileName := time.Now().Format("2006-01-02") + ".log"
	fileName := logFilePath + logFileName
	_, err = os.Stat(fileName)
	if err != nil {
		_, err = os.Create(fileName)
		if err != nil {
			return nil, err
		}
	}
	src, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if err != nil {
		return nil, errors.Wrap(err, "utils.logger failed")
	}
	return src, nil
}
