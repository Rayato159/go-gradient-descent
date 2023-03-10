package usecases

import (
	"context"
	"log"
	"math/rand"
	"time"

	"www.github.com/Rayato159/go-gradient-descent/modules/entities"
	"www.github.com/Rayato159/go-gradient-descent/pkg/utils"
)

type modelsUse struct {
	ModelsRep entities.ModelsRepository
}

func NewModelsUsecase(modelsRep entities.ModelsRepository) *modelsUse {
	return &modelsUse{
		ModelsRep: modelsRep,
	}
}

func (mu *modelsUse) GetData(ctx context.Context) (*entities.DataGroup, error) {
	ctx = context.WithValue(ctx, entities.ModelsUse, time.Now().UnixMilli())
	log.Printf("called:\t%v", utils.Trace())
	defer log.Printf("return:\t%v time:%v ms", utils.Trace(), utils.CallTimer(ctx.Value(entities.ModelsUse).(int64)))

	train, err := mu.ModelsRep.GetData(ctx, "train_data", 1)
	if err != nil {
		return nil, err
	}
	test, err := mu.ModelsRep.GetData(ctx, "test_data", 0)
	if err != nil {
		return nil, err
	}
	res := &entities.DataGroup{
		TrainData: train,
		TestData:  test,
	}
	return res, nil
}

func (mu *modelsUse) TrainModel(ctx context.Context, req *entities.TrainReq) (*entities.TrainRes, error) {
	ctx = context.WithValue(ctx, entities.ModelsUse, time.Now().UnixMilli())
	log.Printf("called:\t%v", utils.Trace())
	defer log.Printf("return:\t%v time:%v ms", utils.Trace(), utils.CallTimer(ctx.Value(entities.ModelsUse).(int64)))

	// Get data prepare for training
	train, err := mu.ModelsRep.GetData(ctx, "train_data", req.TrainDataRatio)
	if err != nil {
		return nil, err
	}
	test, err := mu.ModelsRep.GetData(ctx, "test_data", 0)
	if err != nil {
		return nil, err
	}
	data := &entities.DataGroup{
		TrainData: train,
		TestData:  test,
	}

	// Hyper params set
	rand.Seed(time.Now().UnixNano())
	weights := []float64{
		rand.NormFloat64(),
		rand.NormFloat64(),
	}
	result := utils.GradientDescent(req.LearningRate, req.StepSize, weights, data)
	if err := mu.ModelsRep.InsertTrainResult(ctx, result); err != nil {
		return nil, err
	}

	// Calculate an error from test data
	result.ErrorTest = utils.AverageTestError(result.Weights, test)

	return result, nil
}

func (mu *modelsUse) ClearData(ctx context.Context) error {
	ctx = context.WithValue(ctx, entities.ModelsUse, time.Now().UnixMilli())
	log.Printf("called:\t%v", utils.Trace())
	defer log.Printf("return:\t%v time:%v ms", utils.Trace(), utils.CallTimer(ctx.Value(entities.ModelsUse).(int64)))

	if err := mu.ModelsRep.ClearData(ctx); err != nil {
		return err
	}
	return nil
}

func (mu *modelsUse) ClearRecord(ctx context.Context) error {
	ctx = context.WithValue(ctx, entities.ModelsUse, time.Now().UnixMilli())
	log.Printf("called:\t%v", utils.Trace())
	defer log.Printf("return:\t%v time:%v ms", utils.Trace(), utils.CallTimer(ctx.Value(entities.ModelsUse).(int64)))

	if err := mu.ModelsRep.ClearRecord(ctx); err != nil {
		return err
	}
	return nil
}

func (mu *modelsUse) Predict(ctx context.Context, req *entities.Predict) error {
	ctx = context.WithValue(ctx, entities.ModelsUse, time.Now().UnixMilli())
	log.Printf("called:\t%v", utils.Trace())
	defer log.Printf("return:\t%v time:%v ms", utils.Trace(), utils.CallTimer(ctx.Value(entities.ModelsUse).(int64)))

	weights, err := mu.ModelsRep.GetWeights(ctx)
	if err != nil {
		return err
	}
	req.Result = weights[0]*req.Feature + weights[1]

	return nil
}

func (mu *modelsUse) InsertData(ctx context.Context, req *entities.DataGroup) error {
	ctx = context.WithValue(ctx, entities.ModelsUse, time.Now().UnixMilli())
	log.Printf("called:\t%v", utils.Trace())
	defer log.Printf("return:\t%v time:%v ms", utils.Trace(), utils.CallTimer(ctx.Value(entities.ModelsUse).(int64)))

	if err := mu.ModelsRep.InsertData(ctx, req.TrainData, req.TrainData); err != nil {
		return err
	}
	return nil
}
