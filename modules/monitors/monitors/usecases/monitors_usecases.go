package usecases

import (
	"context"
	"log"
	"time"

	"www.github.com/Rayato159/go-gradient-descent/configs"
	"www.github.com/Rayato159/go-gradient-descent/modules/entities"
	"www.github.com/Rayato159/go-gradient-descent/pkg/utils"
)

type monitorsUse struct {
	MonitorsRepo any
}

func NewMonitorsUsecase() entities.MonitorsUsecase {
	return &monitorsUse{
		MonitorsRepo: nil,
	}
}

func (mu *monitorsUse) HealthCheck(ctx context.Context, cfg *configs.Config) entities.Monitor {
	ctx = context.WithValue(ctx, entities.MonitorsUse, time.Now().UnixMilli())
	log.Printf("called:\t%v", utils.Trace())
	defer log.Printf("return:\t%v time:%v ms", utils.Trace(), utils.CallTimer(ctx.Value(entities.MonitorsUse).(int64)))

	return entities.Monitor{
		Health:  "health is 100% 👌" + time.Now().Format("2006-01-02 15:04:05"),
		Version: cfg.App.Version,
	}
}
