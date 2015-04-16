package mlearn

type Classifier interface {
	Fit() error
	Predict() error
}