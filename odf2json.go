package main

import (
	"fmt"
	xlx "github.com/360EntSecGroup-Skylar/excelize"
	log "github.com/sirupsen/logrus"
	"os"
)

func initMain() {
	InitFlags()
	InitLogger()
}

func main() {

	initMain()

	log.Trace("starting main")
	log.Trace("opening file ", GetInfile())

	f, err := xlx.OpenFile(GetInfile())
	if nil != err || nil == f {
		log.Fatal("could not open ", GetInfile())
		os.Exit(-1)
	}

	sMap := f.GetSheetMap()
	sName := sMap[1]

	rowIterator, err := f.Rows(sName)
	if rowIterator == nil || err != nil {
		var p string
		if nil != err {
			p = err.Error()
		} else {
			p = "err was null?"
		}
		s := fmt.Sprintf("rowIterator might be nil (%p), error (%p): %s\n", rowIterator, err, p)
		log.Fatal(s)
	}

	if rowIterator.Next() {

		var headers []string

		headers = rowIterator.Columns()
		log.Trace("Opening outfile ", GetOutfile())
		StartJson(GetOutfile())

		for rowIterator.Next() {
			OutputJson(headers, rowIterator.Columns())

		}

		EndJson()

	} else {
		log.Info("empty spreadsheet / no data for ", GetInfile(), "?")
	}

}
