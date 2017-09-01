package log

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/mch1307/gomotics/config"
)

// TODO check how to avoid code dup
// how to test fatal and panic

type LogEntries []LogEntry

var LogLine LogEntry
var logFile string

// LogEntry struct to parse log file
type LogEntry struct {
	Filename string    `json:"filename"`
	Function string    `json:"function"`
	Level    string    `json:"level"`
	Line     int       `json:"line"`
	Msg      string    `json:"msg"`
	Package  string    `json:"package"`
	Time     time.Time `json:"time"`
}

func init() {
	logFile = filepath.Join(config.Conf.ServerConfig.LogPath, "gomotics.log")
	if err := os.Remove(logFile); err != nil {
		fmt.Println("could not delete log file")
	}
	config.Conf.ServerConfig.LogLevel = "DEBUG"
	config.Conf.ServerConfig.LogPath = "."
	//Init()
	Debug("Init log testing")
	//readLogs()
}

func readLogs() (list LogEntries) {
	res := LogEntries{}
	jsonFile, err := os.Open(logFile)
	if err != nil {
		fmt.Println("unable to open log file ", filepath.Join(config.Conf.ServerConfig.LogPath, "gomotics.log"), err)
	}
	defer jsonFile.Close()
	reader := bufio.NewScanner(jsonFile)
	for reader.Scan() {
		if err := json.Unmarshal(reader.Bytes(), &LogLine); err != nil {
			fmt.Println(err)
		}
		fmt.Println("log reader: ", &LogLine.Msg)
		res = append(res, LogLine)
	}
	return res
}

func searchLogMsg(msg, function string) LogEntry {
	list := readLogs()
	var val LogEntry
	for _, val = range list {
		if (val.Filename == msg) && val.Function == function {
			return val
		}
	}
	return val
}

func TestDebug(t *testing.T) {
	moreInfo := retrieveCallInfo()
	str := "test Debug"
	Debug(str)
	res := searchLogMsg(str, moreInfo.funcName)
	if res.Msg != str {
		t.Error("test debug: logged msg not found")
	}
}

func TestInfo(t *testing.T) {
	moreInfo := retrieveCallInfo()
	level := "info"
	str := "test " + level
	Info(str)
	res := searchLogMsg(str, moreInfo.funcName)
	if res.Msg != str {
		t.Errorf("test %v: logged msg not found", level)
	}
}

func TestInfof(t *testing.T) {
	moreInfo := retrieveCallInfo()
	level := "info"
	str := "test %v"
	Infof(str, level)
	//time.Sleep(300 * time.Millisecond)
	res := searchLogMsg(str, moreInfo.funcName)
	if res.Msg != "test info" {
		t.Errorf("test %v: logged msg not found", level)
	}
}

func TestWarn(t *testing.T) {
	moreInfo := retrieveCallInfo()
	level := "warn"
	str := "test " + level
	Warn(str)
	res := searchLogMsg(str, moreInfo.funcName)
	if res.Msg != str {
		t.Errorf("test %v: logged msg not found", level)
	}
}

func TestWarnf(t *testing.T) {
	moreInfo := retrieveCallInfo()
	level := "warn"
	str := "test %v"
	Warnf(str, level)
	//time.Sleep(300 * time.Millisecond)
	res := searchLogMsg(str, moreInfo.funcName)
	if res.Msg != "test warn" {
		t.Errorf("test %v: logged msg not found", level)
	}
}

func TestError(t *testing.T) {
	moreInfo := retrieveCallInfo()
	level := "error"
	str := "test " + level
	Error(str)
	res := searchLogMsg(str, moreInfo.funcName)
	if res.Msg != str {
		t.Errorf("test %v: logged msg not found", level)
	}
}

func TestErrorf(t *testing.T) {
	moreInfo := retrieveCallInfo()
	level := "error"
	str := "test %v"
	Errorf(str, level)
	//time.Sleep(300 * time.Millisecond)
	res := searchLogMsg(str, moreInfo.funcName)
	if res.Msg != "test error" {
		t.Errorf("test %v: logged msg not found", level)
	}
}
