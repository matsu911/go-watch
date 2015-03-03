package gowatch

import (
	"github.com/fatih/color"
	. "gopkg.in/check.v1"
)

func (s *S) TestTermWidth(c *C) {
	c.Check(termWidth(), Equals, 80)
}

func (s *S) TestRunCommand(c *C) {
	c.Check(runCommand("echo Hello"), Equals, "Hello\n")
}

func (s *S) TestHeader(c *C) {
	c.Check(header("Test", color.FgRed), Equals, "\x1b[31m====\x1b[0m\x1b[31m Test \x1b[0m\x1b[31m======================================================================\x1b[0m")
}

func (s *S) TestPrinter(c *C) {
	c.Check(printer("Test", ""), Equals, "")
}
