package entities

import "context"

type ModelsContext string

const (
	ModelsCon ModelsContext = "ModelsController"
	ModelsUse ModelsContext = "ModelsUsecase"
	ModelsRep ModelsContext = "ModelsRepository"
)

type ModelsRepository interface {
	GetTrainData(ctx context.Context, ratio float64) ([]Data, error)
}
type ModelsUsecase interface {
	GetTrainData(ctx context.Context, ratio float64) ([]Data, error)
}

type Data struct {
	X float64
	Y float64
}
