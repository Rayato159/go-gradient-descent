package controllers

import (
	"context"
	"log"
	"strings"
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
	r.Get("/data", controller.GetData)
}

func (mc *modelsCon) GetData(c *fiber.Ctx) error {
	ctx := context.WithValue(c.Context(), entities.ModelsCon, time.Now().UnixMilli())
	log.Printf("called:\t%v", utils.Trace())
	defer log.Printf("return:\t%v time:%v ms", utils.Trace(), utils.CallTimer(ctx.Value(entities.ModelsCon).(int64)))

	getTypeMap := map[string]string{
		"train": "train",
		"test":  "test",
	}
	getTypeQuery := strings.ToLower(c.Query("get_type"))
	if getTypeMap[getTypeQuery] == "" || getTypeMap[getTypeQuery] != getTypeQuery {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ErrResponse{
			Status:  fiber.ErrBadRequest.Message,
			Message: "error, get_type is invalid",
		})
	}

	res, err := mc.ModelsUse.GetData(ctx, getTypeQuery)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entities.ErrResponse{
			Status:  fiber.ErrInternalServerError.Message,
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(res)
}
