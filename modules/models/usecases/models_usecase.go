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

func (mu *modelsUse) GetData(ctx context.Context, getType string) ([]entities.Data, error) {
	ctx = context.WithValue(ctx, entities.ModelsUse, time.Now().UnixMilli())
	log.Printf("called:\t%v", utils.Trace())
	defer log.Printf("return:\t%v time:%v ms", utils.Trace(), utils.CallTimer(ctx.Value(entities.ModelsUse).(int64)))

	data, err := mu.ModelsRep.GetData(ctx, getType)
	if err != nil {
		return nil, err
	}
	return data, nil
}
