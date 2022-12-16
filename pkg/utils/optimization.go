package utils

import (
	"math"

	"www.github.com/Rayato159/go-gradient-descent/modules/entities"
)

func ObjectiveFunc(params []float64, x float64) float64 {
	return params[0]*x + params[1]
}

func SumSquareError(params []float64, data []entities.Data) float64 {
	var sum float64
	for i := range data {
		sum += math.Pow(data[i].Y-(ObjectiveFunc(params, data[i].X)), 2)
	}
	return sum
}

func Gradient(h float64, params []float64, data []entities.Data) []float64 {
	results := make([]float64, len(params))
	for i := range params {
		f := make([]float64, 0)
		f = append(f, params...)
		b := make([]float64, 0)
		b = append(b, params...)
		f[i] += h
		b[i] -= h
		results[i] = (SumSquareError(f, data) - SumSquareError(b, data)) / (2 * h)
	}
	return results
}

func GradientDescent(lr float64, data *entities.DataGroup) *entities.TrainRes {
	return nil
}
