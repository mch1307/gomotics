package log

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/mch1307/gomotics/config"
)

var (
	globalConf config.GlobalConfig
)

type callInfo struct {
	packageName string
	fileName    string
	funcName    string
	line        int
}

// Init initialize logger
// Don't use init() otherwise get called before the conf file is parsed
func Init() {

	log.SetFormatter(&log.JSONFormatter{})

	var err error
	logLevel, _ := log.ParseLevel(config.Conf.ServerConfig.LogLevel)

	log.SetLevel(logLevel)
	logFilePath := filepath.Join(config.Conf.ServerConfig.LogPath, "gomotics.log")
	file, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		panic(err)
	}
	if config.Conf.ServerConfig.LogPath == "stdout" {
		log.SetOutput(os.Stdout)
	} else {
		log.SetOutput(file)
	}
	log.Info("logger initialized")
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {
	moreInfo := retrieveCallInfo()
	log.WithFields(log.Fields{
		"filename": moreInfo.fileName,
		"package":  moreInfo.packageName,
		"function": moreInfo.funcName,
		"line":     moreInfo.line,
	}).Debugln(args...)
}

// Debugf logs a message at level Debug on the standard logger.
func Debugf(format string, args ...interface{}) {
	moreInfo := retrieveCallInfo()
	log.WithFields(log.Fields{
		"filename": moreInfo.fileName,
		"package":  moreInfo.packageName,
		"function": moreInfo.funcName,
		"line":     moreInfo.line,
	}).Debugf(format, args...)
}

// Info logs a message at level Info on the standard logger.
func Info(args ...interface{}) {
	moreInfo := retrieveCallInfo()
	log.WithFields(log.Fields{
		"filename": moreInfo.fileName,
		"package":  moreInfo.packageName,
		"function": moreInfo.funcName,
		"line":     moreInfo.line,
	}).Infoln(args...)
}

// Infof logs a message at level Info on the standard logger.
func Infof(format string, args ...interface{}) {
	moreInfo := retrieveCallInfo()
	log.WithFields(log.Fields{
		"filename": moreInfo.fileName,
		"package":  moreInfo.packageName,
		"function": moreInfo.funcName,
		"line":     moreInfo.line,
	}).Infof(format, args...)
}

// Warn logs a message at level Warn on the standard logger.
func Warn(args ...interface{}) {
	moreInfo := retrieveCallInfo()
	log.WithFields(log.Fields{
		"filename": moreInfo.fileName,
		"package":  moreInfo.packageName,
		"function": moreInfo.funcName,
		"line":     moreInfo.line,
	}).Warnln(args...)
}

// Warnf logs a message at level Warn on the standard logger.
func Warnf(format string, args ...interface{}) {
	moreInfo := retrieveCallInfo()
	log.WithFields(log.Fields{
		"filename": moreInfo.fileName,
		"package":  moreInfo.packageName,
		"function": moreInfo.funcName,
		"line":     moreInfo.line,
	}).Warnf(format, args...)
}

// Error logs a message at level Error on the standard logger.
func Error(args ...interface{}) {
	moreInfo := retrieveCallInfo()
	log.WithFields(log.Fields{
		"filename": moreInfo.fileName,
		"package":  moreInfo.packageName,
		"function": moreInfo.funcName,
		"line":     moreInfo.line,
	}).Errorln(args...)
}

// Errorf logs a message at level Error on the standard logger.
func Errorf(format string, args ...interface{}) {
	moreInfo := retrieveCallInfo()
	log.WithFields(log.Fields{
		"filename": moreInfo.fileName,
		"package":  moreInfo.packageName,
		"function": moreInfo.funcName,
		"line":     moreInfo.line,
	}).Errorf(format, args...)
}

// Fatal logs a message at level Fatal on the standard logger.
func Fatal(args ...interface{}) {
	moreInfo := retrieveCallInfo()
	log.WithFields(log.Fields{
		"filename": moreInfo.fileName,
		"package":  moreInfo.packageName,
		"function": moreInfo.funcName,
		"line":     moreInfo.line,
	}).Fatalln(args...)
}

// Fatalf logs a message at level Fatal on the standard logger.
func Fatalf(format string, args ...interface{}) {
	moreInfo := retrieveCallInfo()
	log.WithFields(log.Fields{
		"filename": moreInfo.fileName,
		"package":  moreInfo.packageName,
		"function": moreInfo.funcName,
		"line":     moreInfo.line,
	}).Fatalf(format, args...)
}

func retrieveCallInfo() *callInfo {
	pc, file, line, _ := runtime.Caller(2)
	_, fileName := path.Split(file)
	parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	pl := len(parts)
	packageName := ""
	funcName := parts[pl-1]

	if parts[pl-2][0] == '(' {
		funcName = parts[pl-2] + "." + funcName
		packageName = strings.Join(parts[0:pl-2], ".")
	} else {
		packageName = strings.Join(parts[0:pl-1], ".")
	}

	return &callInfo{
		packageName: packageName,
		fileName:    fileName,
		funcName:    funcName,
		line:        line,
	}
}
