package gowatch

import (
	"time"

	"github.com/beefsack/go-rate"
	. "gopkg.in/check.v1"
)

func (s *S) TestIntSeq(c *C) {
	i := intSeq()
	c.Check(i(), Equals, 2)
	c.Check(i(), Equals, 3)
}

func (s *S) TestRateLimit(c *C) {
	c.Check(testRateLimitHelper(1, 0), Equals, 1)
	c.Check(testRateLimitHelper(2, 0), Equals, 1)
	c.Check(testRateLimitHelper(3, 0), Equals, 1)
	c.Check(testRateLimitHelper(1, 1), Equals, 2)
	c.Check(testRateLimitHelper(2, 1), Equals, 2)
	c.Check(testRateLimitHelper(2, 2), Equals, 2)
}

func testRateLimitHelper(n, m int) int {
	ch := make(chan bool, 1)
	rl := rate.New(1, time.Millisecond/4)

	for i := 0; i < n; i++ {
		rateLimit(ch, rl)
	}

	j := 0
	if <-ch {
		j++
	}
	time.Sleep(time.Millisecond / 4)

	for i := 0; i < m; i++ {
		rateLimit(ch, rl)
	}

	close(ch)

	for _ = range ch {
		j++
	}
	return j
}
