package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
	"path"
)

var fInfile string
var fOutfile string
var fVerbose bool
var typeName string
var fDebug bool

func InitFlags() {

	flag.StringVar(&typeName, "type", "record", "Name of record type conversion")
	flag.StringVar(&fInfile, "infile", "", "ODS file to process")
	flag.StringVar(&fOutfile, "outfile", "", "Output JSON file")
	flag.BoolVar(&fVerbose, "verbose", true, "Lotsa information")
	flag.BoolVar(&fDebug, "debug", false, "log data to STDERR as well as log file")
	flag.Parse()

	if ("" == fInfile) || ("" == fOutfile) {
		fmt.Print("\n")
		flag.Usage()
		log.Exit(0)
	}

	fInfile = path.Clean(fInfile)
	fOutfile = path.Clean(fOutfile)
}

func GetDebug() bool {
	return fDebug
}

func GetTypeName() string {
	return typeName
}

func GetInfile() string {
	return fInfile
}

func GetOutfile() string {
	return fOutfile
}

func GetVerbose() bool {
	return fVerbose
}
