package entities

import "context"

type ModelsContext string

const (
	ModelsCon ModelsContext = "ModelsController"
	ModelsUse ModelsContext = "ModelsUsecase"
	ModelsRep ModelsContext = "ModelsRepository"
)

type ModelsRepository interface {
	GetData(ctx context.Context, getType string, ratio float64) ([]Data, error)
}
type ModelsUsecase interface {
	GetData(ctx context.Context) (*DataGroup, error)
}

type DataGroup struct {
	TrainData []Data `json:"train_data"`
	TestData  []Data `json:"test_data"`
}

type Data struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type TrainReq struct {
	TrainDataRatio float64 `query:"train_data_ratio"`
}
