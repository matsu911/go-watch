// +build debug

package gowatch

import "log"

const debugOn = true

func debug(fmt string, args ...interface{}) {
	log.Printf(fmt, args...)
}
