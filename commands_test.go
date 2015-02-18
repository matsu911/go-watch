package gowatch

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type S struct{}

var _ = Suite(&S{})

func (s *S) TestTermWidth(c *C) {
	c.Check(termWidth(), Equals, -1)

}

func (s *S) TestRunCommand(c *C) {
	c.Check(runCommand("echo Hello"), Equals, "Hello\n")

}
