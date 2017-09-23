package log

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
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
	fmt.Println("starting log test")
	config.Conf.ServerConfig.LogLevel = "DEBUG"
	config.Conf.ServerConfig.LogPath = "."
	logFile = filepath.Join(config.Conf.ServerConfig.LogPath, "gomotics.log")
	/* 	if err := os.Remove(logFile); err != nil {
		fmt.Println("could not delete log file")
	} */

	Init()
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
		//fmt.Println("log reader: ", &LogLine.Msg)
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

func TestDebugf(t *testing.T) {
	moreInfo := retrieveCallInfo()
	level := "debugf"
	str := "test " + level
	Debugf(str)
	res := searchLogMsg(str, moreInfo.funcName)
	if res.Msg != str {
		t.Error("test debugf: logged msg not found")
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

func TestFatal(t *testing.T) {
	fmt.Println("Starting test fatal")
	moreInfo := retrieveCallInfo()
	level := "fatal"
	str := "test " + level

	if os.Getenv("FATAL") == "1" {
		Fatal(str)
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestFatal")
	fmt.Println("args ", os.Args[0])
	fmt.Println("command: ", cmd)
	cmd.Env = append(os.Environ(), "FATAL=1")
	err := cmd.Run()
	fmt.Println("Error: ", err)
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		fmt.Println("cmd in error .. ok")
		res := searchLogMsg(str, moreInfo.funcName)
		if res.Msg != str {
			t.Errorf("test %v: logged msg not found", level)
		} else {
			//fmt.Println(res.Msg)
			return
		}
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}

func TestFatalf(t *testing.T) {
	moreInfo := retrieveCallInfo()
	level := "fatalf"
	str := "test " + level

	if os.Getenv("FATALF") == "1" {
		Fatalf(str)
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestFatalf")
	cmd.Env = append(os.Environ(), "FATALF=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		res := searchLogMsg(str, moreInfo.funcName)
		if res.Msg != str {
			t.Errorf("test %v: logged msg not found", level)
		} else {
			fmt.Println(res.Msg)
			return
		}
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}
