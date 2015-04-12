package supervised

import (
	"testing"

	"gopkg.in/check.v1"
)

func TestPerceptron(t *testing.T) {
	check.Suite(&PerceptronSuite{})
	check.TestingT(t)
}

type PerceptronSuite struct {

}

func (s *PerceptronSuite) SetUpTest(c *check.C) {

}

func (s *PerceptronSuite) TestFit(c *check.C) {

}
