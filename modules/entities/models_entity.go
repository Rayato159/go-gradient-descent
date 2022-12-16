package entities

import "context"

type ModelsContext string

const (
	ModelsCon ModelsContext = "ModelsController"
	ModelsUse ModelsContext = "ModelsUsecase"
	ModelsRep ModelsContext = "ModelsRepository"
)

type ModelsRepository interface {
	GetData(ctx context.Context, getType string) ([]Data, error)
}
type ModelsUsecase interface {
	GetData(ctx context.Context, getType string) ([]Data, error)
}

type Data struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}
