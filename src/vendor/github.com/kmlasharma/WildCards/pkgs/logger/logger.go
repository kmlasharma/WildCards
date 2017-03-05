package logger

import (
	"fmt"
	"log"
	"os"
)

const (
	defaultLogFile = "/root/log/output.log"
	outputToStdout = true
)

var f, _ = os.OpenFile(defaultLogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

func Println(a ...interface{}) {
	log.SetOutput(f)
	log.Println(a...)

	if outputToStdout {
		fmt.Println(a...)
	}
}

func Print(a ...interface{}) {
	log.SetOutput(f)
	log.Print(a...)

	if outputToStdout {
		fmt.Print(a...)
	}
}
