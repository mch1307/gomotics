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
// Don't use ini() otherwise get called before the conf file is parsed
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
	log.SetOutput(file)
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

// Warn logs a message at level Warn on the standard logger.
func Warn(args ...interface{}) {
	moreInfo := retrieveCallInfo()
	log.WithFields(log.Fields{
		"filename": moreInfo.fileName,
		"package":  moreInfo.packageName,
		"function": moreInfo.funcName,
		"line":     moreInfo.line,
	}).Warningln(args...)
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

// Panic logs a message at level Panic on the standard logger.
func Panic(args ...interface{}) {
	moreInfo := retrieveCallInfo()
	log.WithFields(log.Fields{
		"filename": moreInfo.fileName,
		"package":  moreInfo.packageName,
		"function": moreInfo.funcName,
		"line":     moreInfo.line,
	}).Panicln(args...)
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
