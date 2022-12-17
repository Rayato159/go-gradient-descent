package entities

import (
	"context"
	"time"
)

type ModelsContext string

const (
	ModelsCon ModelsContext = "ModelsController"
	ModelsUse ModelsContext = "ModelsUsecase"
	ModelsRep ModelsContext = "ModelsRepository"
)

type ModelsRepository interface {
	GetData(ctx context.Context, getType string, ratio float64) ([]Data, error)
	InsertTrainResult(ctx context.Context, req *TrainRes) error
	ClearData(ctx context.Context) error
	ClearRecord(ctx context.Context) error
	GetWeights(ctx context.Context) ([]float64, error)
}
type ModelsUsecase interface {
	GetData(ctx context.Context) (*DataGroup, error)
	TrainModel(ctx context.Context, req *TrainReq) (*TrainRes, error)
	ClearData(ctx context.Context) error
	ClearRecord(ctx context.Context) error
	Predict(ctx context.Context, req *Predict) error
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
	ErrorTrain float64   `bson:"error_train" json:"error_train"`
	ErrorTest  float64   `bson:"error_test" json:"error_test"`
	Weights    []float64 `bson:"weights" json:"weights"`
	Timestamp  time.Time `bson:"timestamp" json:"timestamp"`
}

type Predict struct {
	Feature float64 `json:"feature"`
	Result  float64 `json:"result"`
}
