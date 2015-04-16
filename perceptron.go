package mlearn

import "math"

// Perceptron is a binary classifier algorithm for linearly separable data.
type Perceptron struct {
	Rate    float64
	Weights []float64
}

// Fit tries to find a pattern in datasets x assigned to y.
func (p *Perceptron) Fit(x [][]float64, y []int) error {
	return nil
}

// Score returns the aggregated error on label classification.
func (p *Perceptron) Score(x [][]float64, y []int) float64 {
	sum := 0.0

	for i := 0; i < len(x); i++ {
		sum += p.score(x[i], y[i])
	}

	return sum
}

// Predict returns the predicted label for the given dataset x.
func (p *Perceptron) Predict(x [][]float64) ([]int, error) {
	//x = append([]float64{1}, x...)

	return []int{}, nil
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

func (p *Perceptron) score(x []float64, y int) float64 {
	return math.Abs(float64(y - p.predict(x)))
}
