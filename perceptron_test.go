package mlearn

import (
	"testing"

	"gopkg.in/check.v1"
)

func TestPerceptron(t *testing.T) {
	check.Suite(&PerceptronSuite{})
	check.TestingT(t)
}

type PerceptronSuite struct {
	perceptron *Perceptron
}

func (s *PerceptronSuite) SetUpTest(c *check.C) {
	s.perceptron = &Perceptron{
		Rate:    0.4,
		Weights: []float64{1, 1, 1},
	}
}

func (s *PerceptronSuite) TestFit(c *check.C) {
	s.testUpdate(c)
}

func (s *PerceptronSuite) TestScore(c *check.C) {

}

func (s *PerceptronSuite) TestPredict(c *check.C) {
	s.testPredict(c)
}

func (s *PerceptronSuite) testUpdate(c *check.C) {
	s.perceptron.update([]float64{1, 1, 0}, 0)

	c.Assert(s.perceptron.Weights, check.DeepEquals, []float64{.6, .6, 1})
}

func (s *PerceptronSuite) testPredict(c *check.C) {
	c.Assert(s.perceptron.predict([]float64{1, 1, 1}), check.Equals, 1)
}
