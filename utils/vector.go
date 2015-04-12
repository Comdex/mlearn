package utils

import "github.com/bodokaiser/mlearn/common"

func DotProduct(a []float64, b []float64) float64 {
	if len(a) != len(b) {
		panic(common.ErrInputMissmatch)
	}

	s := 0.0

	for i := 0; i < len(a); i++ {
		s += a[i] * b[i]
	}

	return s
}
