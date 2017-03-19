package logger

import (
	"fmt"
	"log"
	"os"
)

const (
	defaultLogFile        = "/root/log/output.log"
	defaultErrorFile      = "/root/log/error.log"
	outputErrorsToLogFile = false
	outputToStdout        = true
)

var logFile = createLogFile(defaultLogFile)
var errorFile = createLogFile(defaultErrorFile)

func Println(a ...interface{}) {
	log.SetOutput(logFile)
	log.Println(a...)

	if outputToStdout {
		fmt.Println(a...)
	}
}

func Print(a ...interface{}) {
	log.SetOutput(logFile)
	log.Print(a...)

	if outputToStdout {
		fmt.Print(a...)
	}
}

func Error(a ...interface{}) {
	log.SetOutput(errorFile)
	log.Println(a...)

	if outputErrorsToLogFile {
		log.SetOutput(logFile)
		log.Println(a...)
	}
	
	if outputToStdout {
		fmt.Print(a...)
	}
}

func createLogFile(path string) *os.File {
	os.MkdirAll("/root/log", os.ModePerm) // Ensure log folder is created.
	f, _ := os.Create(path)
	return f
}
