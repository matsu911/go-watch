package gowatch

import (
	"bytes"
	"log"
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type S struct{}

var _ = Suite(&S{})

func (s *S) TestDebug(c *C) {
	// Capture the debug output.
	buf := new(bytes.Buffer)
	log.SetOutput(buf)

	testString := "Test"
	debug(testString)

	result := buf.String()

	if debugOn {
		c.Check(result[len(result)-len(testString)-1:], Equals, testString+"\n")
	} else {
		c.Check(result, Equals, "")
	}
}
