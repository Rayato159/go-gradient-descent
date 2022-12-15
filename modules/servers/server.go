package servers

import (
	"fmt"
	"log"

	_monitorsHttp "www.github.com/Rayato159/go-gradient-descent/modules/monitors/controllers/http"
	_monitorsUsecase "www.github.com/Rayato159/go-gradient-descent/modules/monitors/usecases"

	_modelsHttp "www.github.com/Rayato159/go-gradient-descent/modules/models/controllers/http"
	_modelsRepository "www.github.com/Rayato159/go-gradient-descent/modules/models/repositories"
	_modelsUsecase "www.github.com/Rayato159/go-gradient-descent/modules/models/usecases"

	"www.github.com/Rayato159/go-gradient-descent/configs"
	"www.github.com/Rayato159/go-gradient-descent/modules/entities"
	"www.github.com/Rayato159/go-gradient-descent/pkg/middlewares"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type server struct {
	Cfg *configs.Config
	App *fiber.App
	Db  *mongo.Database
}

func NewServer(cfg *configs.Config, db *mongo.Database) *server {
	fiberConfigs := configs.NewFiberConfig(cfg.App)
	return &server{
		Cfg: cfg,
		App: fiber.New(fiberConfigs),
		Db:  db,
	}
}

func (s *server) Start() {
	// Map all routes
	if err := s.mapHandlers(); err != nil {
		log.Fatalln(err.Error())
	}

	// Server config
	host := s.Cfg.App.Host
	port := s.Cfg.App.Port
	fiberConnURL := fmt.Sprintf("%s:%s", s.Cfg.App.Host, s.Cfg.App.Port)
	log.Printf("server has been started on %s:%s ⚡", host, port)

	// Start server
	if err := s.App.Listen(fiberConnURL); err != nil {
		log.Fatalln(err.Error())
	}
}

func (s *server) mapHandlers() error {
	// Cors config
	middlewares.NewCorsFiberHandler(s.App)

	// Group a version
	v1 := s.App.Group("/v1")

	//* Monitors group.
	monitorsUsecase := _monitorsUsecase.NewMonitorsUsecase()
	_monitorsHttp.NewMonitorsController(v1, s.Cfg, monitorsUsecase)

	//* Users group
	modelsGroup := v1.Group("/models")
	modelsRepository := _modelsRepository.NewModelsRepository(s.Db)
	modelsUsecase := _modelsUsecase.NewModelsUsecase(modelsRepository)
	_modelsHttp.NewModelsController(modelsGroup, modelsUsecase)

	// End point not found error response
	s.App.Use(func(c *fiber.Ctx) error {
		log.Println("error, endpoint is not found")
		return c.Status(fiber.StatusNotFound).JSON(entities.ErrResponse{
			Status:  fiber.ErrNotFound.Message,
			Message: "error, endpoint is not found",
		})
	})
	return nil
}
