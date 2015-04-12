package utils

import (
	"testing"

	"gopkg.in/check.v1"
)

func TestVector(t *testing.T) {
	check.Suite(&VectorSuite{})
	check.TestingT(t)
}

type VectorSuite struct {
	a, b []float64
}

func (s *VectorSuite) SetUpTest(c *check.C) {
	s.a = []float64{2, 4, 6}
	s.b = []float64{1, 3, 5}
}

func (s *VectorSuite) TestDotProduct(c *check.C) {
	c.Check(DotProduct(s.a, s.b), check.Equals, float64(2*1+4*3+6*5))
}
