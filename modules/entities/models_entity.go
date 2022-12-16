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
	GetData(ctx context.Context, req *DataReq) (*DataGroup, error)
}

type DataReq struct {
	GetType    string  `query:"get_type"`
	TrainRatio float64 `query:"train_ratio"`
}

type DataGroup struct {
	TrainData      []Data  `json:"train_data"`
	TestData       []Data  `json:"test_data"`
	TrainDataRatio float64 `json:"train_data_ratio"`
	TestDataRatio  float64 `json:"test_data_ratio"`
}

type Data struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type TrainReq struct {
	TrainDataRatio float64 `query:"train_data_ratio"`
}
