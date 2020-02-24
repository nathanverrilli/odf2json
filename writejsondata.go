package main

import (
	"bufio"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

var f *os.File = nil
var bw *bufio.Writer = nil
var err error = nil
var replacer *strings.Replacer = nil

func StartJson(filename string) {
	log.Trace("in StartJson opening file ", filename)

	/* some characters have special meaning in JSON, so escape them in data */
	if nil == replacer {
		replacer = strings.NewReplacer(
			"\u0022", "\\u0022", // replace quote " with \"
			"\u0027", "\\u0027", // replace singlequote with \'
			"\u0008", "\\u0008", // replace backspace with \<backspace>
			"\u000C", "\\u000C", // replace formfeed with \<formfeed>
			"\u000D", "\\u000D", // replace carriage return with \<carriage return>
			"\u000A", "\\u000A", // replace newline with \<newline>
			"\u0009", "\\u0009", // replace tab with \<tab>
		)
	}

	if nil != f || nil != bw {
		log.Error("This is a single-threaded stateful writer. Attempt to open a new file ",
			filename,
			" before closing ",
			f.Name(), " is an error. Attempting to close old file and aborting ... ")
		EndJson()
		log.Fatal(" attempted to cleanly close old file, dying now")
	}

	f, err = os.Create(filename)
	if nil != err {
		log.Fatal("Could not open file ", filename,
			"because ", err.Error())
	}

	bw = bufio.NewWriter(f)
}

func EndJson() {

	err1 := bw.Flush()
	err2 := f.Close()
	if nil != err1 || nil != err2 {
		if nil != err1 {
			log.Error("bufio writer failure ", err1.Error())
		}
		if nil != err2 {
			log.Error("file failure ", err2.Error())
		}
		log.Fatal("io error")
	}
	bw = nil
	f = nil
}

func endJsonEntry() {
	jsonWriteString(" } ")
}

func startJsonEntry() {
	jsonWriteString(" { ")
}

func jsonString(s string) {
	jsonWriteString("\"")
	jsonWriteString(replacer.Replace(s))
	jsonWriteString("\"")
}

func OutputJson(headers []string, data []string) {
	startJsonEntry()

	if len(headers) != len(data) {
		for len(headers) > len(data) {
			data = append(data, "dummy data")
		}

		for ix := 0; len(headers) < len(data); ix++ {
			headers = append(headers, fmt.Sprintf("dummyHeader%04d", ix))
		}
	}
	for ix := range headers {

		jsonString(headers[ix])
		jsonWriteString(" : ")
		jsonString(data[ix])

		/* comma between each data pair, but not the end */
		if ix < (len(headers) - 1) {
			jsonWriteString(",")
		}
	}
	endJsonEntry()

	// newline for cbimport compatibility
	jsonWriteString("\n")
}

func jsonWriteString(s string) {
	_, err = bw.WriteString(s)
	if nil != err {
		log.Fatal("Failed to write JSON data to buffered IO-- please check disk space and file")
	}
}
