package usecases

import (
	"context"
	"log"
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

func (mu *modelsUse) GetData(ctx context.Context, req *entities.DataReq) (*entities.DataGroup, error) {
	ctx = context.WithValue(ctx, entities.ModelsUse, time.Now().UnixMilli())
	log.Printf("called:\t%v", utils.Trace())
	defer log.Printf("return:\t%v time:%v ms", utils.Trace(), utils.CallTimer(ctx.Value(entities.ModelsUse).(int64)))

	switch req.GetType {
	case "train":
		train, err := mu.ModelsRep.GetData(ctx, "train_data", req.TrainRatio)
		if err != nil {
			return nil, err
		}
		res := &entities.DataGroup{
			TrainData:      train,
			TestData:       make([]entities.Data, 0),
			TrainDataRatio: req.TrainRatio,
			TestDataRatio:  1 - req.TrainRatio,
		}
		return res, nil
	case "test":
		test, err := mu.ModelsRep.GetData(ctx, "test_data", req.TrainRatio)
		if err != nil {
			return nil, err
		}
		res := &entities.DataGroup{
			TrainData:      make([]entities.Data, 0),
			TestData:       test,
			TrainDataRatio: req.TrainRatio,
			TestDataRatio:  1 - req.TrainRatio,
		}
		return res, nil
	default:
		train, err := mu.ModelsRep.GetData(ctx, "train_data", 1)
		if err != nil {
			return nil, err
		}
		test, err := mu.ModelsRep.GetData(ctx, "test_data", 0)
		if err != nil {
			return nil, err
		}
		res := &entities.DataGroup{
			TrainData:      train,
			TestData:       test,
			TrainDataRatio: 1,
			TestDataRatio:  1,
		}
		return res, nil
	}
}
