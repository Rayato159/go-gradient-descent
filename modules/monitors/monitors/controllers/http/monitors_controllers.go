package http

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"www.github.com/Rayato159/go-gradient-descent/configs"
	"www.github.com/Rayato159/go-gradient-descent/modules/entities"
	"www.github.com/Rayato159/go-gradient-descent/pkg/utils"
)

type monitorsCon struct {
	MonitorsUse entities.MonitorsUsecase
	Cfg         *configs.Config
}

func NewMonitorsController(r fiber.Router, cfg *configs.Config, monitorsUse entities.MonitorsUsecase) {
	controller := &monitorsCon{
		MonitorsUse: monitorsUse,
		Cfg:         cfg,
	}
	r.Get("/", controller.HealthCheck)
}

func (mc *monitorsCon) HealthCheck(c *fiber.Ctx) error {
	ctx := context.WithValue(c.Context(), entities.MonitorsCon, time.Now().UnixMilli())
	log.Printf("called:\t%v", utils.Trace())
	defer log.Printf("return:\t%v time:%v ms", utils.Trace(), utils.CallTimer(ctx.Value(entities.MonitorsCon).(int64)))

	res := mc.MonitorsUse.HealthCheck(ctx, mc.Cfg)
	return c.Status(fiber.StatusOK).JSON(res)
}
