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

func (mu *modelsUse) GetTrainData(ctx context.Context, ratio float64) ([]entities.Data, error) {
	ctx = context.WithValue(ctx, entities.ModelsUse, time.Now().UnixMilli())
	log.Printf("called:\t%v", utils.Trace())
	defer log.Printf("return:\t%v time:%v ms", utils.Trace(), utils.CallTimer(ctx.Value(entities.ModelsUse).(int64)))

	data, err := mu.ModelsRep.GetTrainData(ctx, ratio)
	if err != nil {
		return nil, err
	}
	return data, nil
}
