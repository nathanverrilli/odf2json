package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
)

const filename string = "odf2json.log"

var logReady bool = false

func isLogReady() bool {
	return logReady
}

func InitLogger() {

	var sb strings.Builder
	sb.WriteString(StripErrorString(os.Executable))
	fmt.Println("\n\t", sb.String(), "\n")

	logLevel := log.DebugLevel
	if GetVerbose() {
		logLevel = log.TraceLevel
		log.Trace("setting logLevel to TraceLevel")
	}
	log.SetLevel(logLevel)

	// Create the log file if doesn't exist. And append to it if it already exists.
	logFile, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	Formatter := new(log.TextFormatter)
	// You can change the Timestamp format. But you have to use the same date and time.
	// "2006-02-02 15:04:06" Works. If you change any digit, it won't work
	// ie "Mon Jan 2 15:04:05 MST 2006" is the reference time. You can't change it
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true
	log.SetFormatter(Formatter)

	var writers []io.Writer

	logToStderr := GetDebug()

	if nil == err {
		writers = append(writers, logFile)
	} else {
		// could not open log file?
		logToStderr = true
		fmt.Println("\n\tError opening logfile ", logFile,
			"\n\tnerror: ", err.Error(),
			"\n\tenabling debug data to STDERR")
	}

	if logToStderr {
		writers = append(writers, os.Stderr)
	}

	log.SetOutput(io.MultiWriter(writers...))

	if nil == err {
		log.Info("Writing to ", filename)
	} else {
		log.Debug("Could not open file ", filename)
	}

	log.Infof("running program %s", sb.String())

	logReady = true
}
