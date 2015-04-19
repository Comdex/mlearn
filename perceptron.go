package mlearn

import "math"
import "errors"

// ErrInvalidInputLen is used when there are more or less features than labels.
var ErrInvalidInputLen = errors.New("invalid length of input vectors")

// Perceptron is a binary classifier algorithm for linearly separable data.
type Perceptron struct {
	Rate    float64
	Weights []float64
}

// Fit tries to find a pattern in datasets x assigned to y.
func (p *Perceptron) Fit(x [][]float64, y []int) error {
	for i := 0; i < len(x); i++ {
		x[i] = append([]float64{1}, x[i]...)
	}

	if len(x) != len(y) {
		return ErrInvalidInputLen
	}

	for p.Error(x, y) > 0 {
		for i := 0; i < len(x); i++ {
			if p.error(x[i], y[i]) > 0 {
				p.update(x[i], y[i])

				break
			}
		}
	}

	return nil
}

// Error returns the aggregated error on label classification.
func (p *Perceptron) Error(x [][]float64, y []int) int {
	sum := 0

	for i := 0; i < len(x); i++ {
		sum += p.error(x[i], y[i])
	}

	return sum
}

// Predict returns the predicted label for the given dataset x.
func (p *Perceptron) Predict(x [][]float64) []int {
	res := make([]int, len(x))

	for i := range x {
		res[i] = p.predict(x[i])
	}

	return res
}

func (p *Perceptron) predict(x []float64) int {
	sum := 0.0

	for i := 0; i < len(x); i++ {
		sum += p.Weights[i] * x[i]
	}

	if sum > 0 {
		return 1
	}

	return 0
}

func (p *Perceptron) update(x []float64, y int) {
	z := p.Rate * float64(y-p.predict(x))

	for i := 0; i < len(x); i++ {
		p.Weights[i] += z * x[i]
	}
}

func (p *Perceptron) error(x []float64, y int) int {
	return int(math.Abs(float64(y - p.predict(x))))
}
