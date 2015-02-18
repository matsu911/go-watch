// +build debug

package gowatch

import "log"

func debug(fmt string, args ...interface{}) {
	log.Printf(fmt, args...)
}
