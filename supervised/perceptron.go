package supervised

import "github.com/bodokaiser/mlearn/common"

type Perceptron struct {
	Count          int32
	Rate, Treshold float32
	Weights        []float64
}

func NewPerceptron() *Perceptron {
	return &Perceptron{
		Count:    0,
		Rate:     .1,
		Treshold: .5,
	}
}

func (p *Perceptron) Fit(x [][]float64, y []float64) error {
	if len(x) != len(y) {
		return common.ErrInputMissmatch
	}
	p.Weights = make([]float64, len(x[0]))

	for i, x := range x {
		e := y[i]

		if utils.DotProduct(x, p.Weights) > p.Treshold {
			e -= 1
		}

		if e != 0 {
			p.Count += 1
		}
	}

	return nil
}
