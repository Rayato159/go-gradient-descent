package usecases

import (
	"context"
	"fmt"
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
	test, err := mu.ModelsRep.GetData(ctx, "test_data", req.TrainDataRatio)
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
	fmt.Println(utils.Gradient(req.StepSize, weights, data.TrainData))

	return nil, nil
}
