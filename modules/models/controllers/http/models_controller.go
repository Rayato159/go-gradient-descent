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
	r.Get("/data", controller.GetData)
	r.Post("/train", controller.TrainModel)
}

func (mc *modelsCon) GetData(c *fiber.Ctx) error {
	ctx := context.WithValue(c.Context(), entities.ModelsCon, time.Now().UnixMilli())
	log.Printf("called:\t%v", utils.Trace())
	defer log.Printf("return:\t%v time:%v ms", utils.Trace(), utils.CallTimer(ctx.Value(entities.ModelsCon).(int64)))

	res, err := mc.ModelsUse.GetData(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entities.ErrResponse{
			Status:  fiber.ErrInternalServerError.Message,
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

func (mc *modelsCon) TrainModel(c *fiber.Ctx) error {
	ctx := context.WithValue(c.Context(), entities.ModelsCon, time.Now().UnixMilli())
	log.Printf("called:\t%v", utils.Trace())
	defer log.Printf("return:\t%v time:%v ms", utils.Trace(), utils.CallTimer(ctx.Value(entities.ModelsCon).(int64)))

	req := new(entities.TrainReq)
	if err := c.QueryParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ErrResponse{
			Status:  fiber.ErrBadRequest.Message,
			Message: err.Error(),
		})
	}

	if req.TrainDataRatio <= 0 || req.TrainDataRatio > 1 {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ErrResponse{
			Status:  fiber.ErrBadRequest.Message,
			Message: "error, train_data_ratio must be (0, 1]",
		})
	}
	if req.LearningRate <= 0 || req.LearningRate > 1 {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ErrResponse{
			Status:  fiber.ErrBadRequest.Message,
			Message: "error, learning_rate must be (0, 1]",
		})
	}
	if req.StepSize <= 0 || req.StepSize > 1 {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ErrResponse{
			Status:  fiber.ErrBadRequest.Message,
			Message: "error, step_size must be (0, 1]",
		})
	}

	res, err := mc.ModelsUse.TrainModel(ctx, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entities.ErrResponse{
			Status:  fiber.ErrInternalServerError.Message,
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(res)
}
