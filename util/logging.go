package util

import (
	"flag"
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
	"time"
)

var _log *logrus.Entry
var logFileRelPath string
var rotationPeriod int
var loggingOnce sync.Once
var terminalOnce sync.Once

func init() {
	flag.StringVar(&logFileRelPath, "logFileRelPath",
		"log", "path to save log files relative to application root")
	flag.IntVar(&rotationPeriod, "logRotationPeriod",
		2520, "time interval to rotate log files")

	os.MkdirAll(logFileRelPath, os.ModePerm)
}

func GetLogger() *logrus.Entry {
	loggingOnce.Do(func() {
		log := logrus.New()
		infoPath := "log/info"
		writer, _ := rotatelogs.New(
			infoPath+".%Y%m%d_%H%M.log",
			rotatelogs.WithRotationTime(time.Duration(24)*time.Hour),
		)

		log.AddHook(lfshook.NewHook(
			lfshook.WriterMap{
				logrus.DebugLevel: writer,
				logrus.InfoLevel:  writer,
				logrus.WarnLevel:  writer,
				logrus.ErrorLevel: writer,
			},
			nil,
		))

		_log = log.WithField("terminalId", "00000000")

	})
	return _log
}