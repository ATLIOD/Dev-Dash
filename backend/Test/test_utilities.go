package Test

import "testing"

type Checker struct {
	t      *testing.T
	failed bool
}

func NewChecker(t *testing.T) *Checker {
	t.Helper()
	return &Checker{t: t}
}

func (c *Checker) Check(ok bool) {
	c.t.Helper()
	if !ok {
		c.failed = true
	}
}

func (c *Checker) Failed() bool {
	c.t.Helper()
	return c.failed
}
