package logger

import (
	"fmt"
	"log"
	"os"
)

var logDir = os.Getenv("LOG_DIR")
var defaultLogFile = logDir + "/analysis.log"
var defaultErrorFile = logDir + "/analysis.err"

const (
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
		fmt.Println(a...)
	}
}

func Fatal(a ...interface{}) {
	Error(a...)
	os.Exit(1)
}

func createLogFile(path string) *os.File {
	os.MkdirAll(logDir, os.ModePerm) // Ensure log folder is created.
	f, _ := os.Create(path)
	return f
}