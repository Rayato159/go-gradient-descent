package configs

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Config struct {
	Database MongoDb
	App      App
}

type App struct {
	Host                 string
	Port                 string
	Version              string
	ServerRequestTimeout string
	Stage                string
}

// Database
type MongoDb struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

func NewFiberConfig(cfg App) fiber.Config {
	readTimeoutSecondCount, _ := strconv.Atoi(cfg.ServerRequestTimeout)

	// Time limit -> 30 seconds or setting in .env
	// Body limit -> 10 MiB
	return fiber.Config{
		ReadTimeout: time.Second * time.Duration(readTimeoutSecondCount),
		BodyLimit:   10 * 1024 * 1024,
	}
}
