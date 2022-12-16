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
	TrainModel(ctx context.Context, req *TrainReq) (*TrainRes, error)
}

type DataGroup struct {
	TrainData []Data `json:"train_data"`
	TestData  []Data `json:"test_data"`
}

type Data struct {
	X float64 `bson:"x" json:"x"`
	Y float64 `bson:"y" json:"y"`
}

type TrainReq struct {
	TrainDataRatio float64 `query:"train_data_ratio"`
	LearningRate   float64 `query:"learning_rate"`
	StepSize       float64 `query:"step_size"`
}

type TrainRes struct {
	Slope      float64   `json:"slope"`
	YIntercept float64   `json:"y_intercept"`
	Error      float64   `bson:"error" json:"error"`
	Weights    []float64 `bson:"weights" json:"weights"`
}
