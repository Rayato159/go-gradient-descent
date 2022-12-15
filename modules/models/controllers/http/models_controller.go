package controllers

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"www.github.com/Rayato159/go-gradient-descent/modules/entities"
	"www.github.com/Rayato159/go-gradient-descent/pkg/utils"
)

type modelsCon struct {
	ModelsUse entities.ModelsUsecase
}

func NewModelsController(r fiber.Router, modelsUse entities.ModelsUsecase) {
	controller := &modelsCon{
		ModelsUse: modelsUse,
	}
	r.Get("/", controller.GetTrainData)
}

func (mc *modelsCon) GetTrainData(c *fiber.Ctx) error {
	ctx := context.WithValue(c.Context(), entities.ModelsCon, time.Now().UnixMilli())
	log.Printf("called:\t%v", utils.Trace())
	defer log.Printf("return:\t%v time:%v ms", utils.Trace(), utils.CallTimer(ctx.Value(entities.ModelsCon).(int64)))

	res, err := mc.ModelsUse.GetTrainData(ctx, 0.8)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entities.ErrResponse{
			Status:  fiber.ErrInternalServerError.Message,
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(res)
}
