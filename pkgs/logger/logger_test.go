package logger

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

// Already declared in logger package
//logDir := os.Getenv("LOG_DIR")
//defaultLogFile := os.Open(logDir + "/analysis.log")
//defaultErrFile := os.Open(logDir + "/analysis.err")

func TestLogFileWrite(t *testing.T) {
	fmt.Println("* Testing writing to log file")
	str := "Testing Log File Write"
	Println(str)
	text, err := ioutil.ReadFile(defaultLogFile)
	if err != nil {
		t.Error("Could not open file")
		return
	}
	if !strings.HasSuffix(string(text), str+"\n") {
		t.Error("String outputted does not match file contents")
	} else {
		fmt.Println("PASSED!")
	}
}

func TestErrorLogFileWrite(t *testing.T) {
	fmt.Println("* Testing writing to err file")
	str := "Testing Error Log File Write"
	Error(str)
	text, err := ioutil.ReadFile(defaultErrorFile)
	if err != nil {
		t.Error("Could not open file")
		return
	}
	if !strings.HasSuffix(string(text), str+"\n") {
		t.Error("String outputted does not match file contents")
	} else {
		fmt.Println("PASSED!")
	}
}
