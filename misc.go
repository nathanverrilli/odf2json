package main

import (
	log "github.com/sirupsen/logrus"
)

func StripErrorString(f func() (string, error)) string {
	s, err := f()
	if nil != err {
		log.Infof("error %s getting a string (got %s)", err.Error(), s)
	}
	return s
}

func TrueOrFalse(q bool) string {
	if q {
		return "true"
	}
	return "false"
}
