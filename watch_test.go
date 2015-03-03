package gowatch

import (
	"testing"
	"time"

	"github.com/beefsack/go-rate"
	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type S struct{}

var _ = Suite(&S{})

func (s *S) TestIntSeq(c *C) {
	i := intSeq()
	c.Check(i(), Equals, 2)
	c.Check(i(), Equals, 3)
}

func testRateLimitHelper(n, m int) int {
	ch := make(chan bool, 1)
	rl := rate.New(1, time.Second/4)

	for i := 0; i < n; i++ {
		rateLimit(ch, rl)
	}
	close(ch)

	j := 0
	for _ = range ch {
		j++
	}
	return j
}

func (s *S) TestRateLimit(c *C) {
	c.Check(testRateLimitHelper(1, 0), Equals, 1)
	c.Check(testRateLimitHelper(2, 0), Equals, 1)
	c.Check(testRateLimitHelper(2, 1), Equals, 2)
}
