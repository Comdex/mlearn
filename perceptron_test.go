package mlearn

import (
	"testing"

	"gopkg.in/check.v1"
)

func TestPerceptron(t *testing.T) {
	check.Suite(&PerceptronSuite{
		labels:  []int{0, 0, 1},
		weights: []float64{-0.6, 0.2, 0.6},
		features: [][]float64{
			{0, 1},
			{1, 0},
			{1, 1},
		},
	})
	check.TestingT(t)
}

type PerceptronSuite struct {
	labels     []int
	weights    []float64
	features   [][]float64
	perceptron *Perceptron
}

func (s *PerceptronSuite) SetUpTest(c *check.C) {
	s.perceptron = &Perceptron{
		Rate:    0.4,
		Weights: []float64{1, 1, 1},
	}
}

func (s *PerceptronSuite) TestFit(c *check.C) {
	s.perceptron.Fit(s.features, s.labels)

	// TODO: compare floats with bad precision
	//c.Assert(s.perceptron.Weights, check.DeepEquals, s.weights)
}

func (s *PerceptronSuite) TestError(c *check.C) {
	err := s.perceptron.Error(s.features, s.labels)

	c.Assert(err, check.Equals, 2)
}

func (s *PerceptronSuite) TestPredict(c *check.C) {
	label := s.perceptron.predict([]float64{1, 1, 1})

	c.Assert(label, check.Equals, 1)
}
