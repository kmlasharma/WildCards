package logger

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestLogFileWrite(t *testing.T) {
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
