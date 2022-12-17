package utils

import (
	"fmt"
	"math"

	"www.github.com/Rayato159/go-gradient-descent/modules/entities"
)

func ToFixed(data []float64, n int) {
	for i := range data {
		data[i] = math.Round(data[i]*math.Pow10(n)) / math.Pow10(n)
	}
}

func ObjectiveFunc(params []float64, x float64) float64 {
	return params[0]*x + params[1]
}

func SumSquareError(params []float64, data []entities.Data) float64 {
	var result float64
	for i := range data {
		result += math.Pow(data[i].Y-(ObjectiveFunc(params, data[i].X)), 2)
	}
	return result
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

func GradientDescent(h float64, params []float64, data *entities.DataGroup) *entities.TrainRes {
	var errValue float64 = 100
	max := 1000
	iter := 0
	p := params
	n := make([]float64, len(params))

	for errValue > 0.001 && max > 0 {
		grad := Gradient(h, p, data.TrainData)
		for i := range params {
			n[i] = p[i] - h*grad[i]
		}
		copy(p, n)

		n = make([]float64, len(params))
		errValue = SumSquareError(p, data.TrainData)
		max--
		iter++
		ToFixed(p, 6)
		fmt.Printf("iter:\t%dparams:%v\terr: %v\n", iter, p, errValue)
	}

	result := &entities.TrainRes{
		Slope:      p[0],
		YIntercept: p[1],
		Error:      errValue,
		Weights:    p,
	}
	return result
}
