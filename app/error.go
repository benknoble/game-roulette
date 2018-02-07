package app

import "log"

func Halt(msg string, a ...interface{}) {
	log.Fatalf(msg+"\n", a...)
}
