package entities

import (
	"context"

	"www.github.com/Rayato159/go-gradient-descent/configs"
)

type MonitorsContext string

const (
	MonitorsCon MonitorsContext = "MonitorsController"
	MonitorsUse MonitorsContext = "MonitorsUsecase"
	MonitorsRep MonitorsContext = "MonitorsRepository"
)

type MonitorsUsecase interface {
	HealthCheck(ctx context.Context, cfg *configs.Config) Monitor
}

type Monitor struct {
	Health  string
	Version string
}
