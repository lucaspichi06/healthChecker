package main

import (
	"github.com/lucaspichi06/healthChecker/pkg/health"
	"github.com/lucaspichi06/healthChecker/pkg/health/monitors"
	"github.com/lucaspichi06/healthChecker/pkg/health/types"
	"time"
)

func main() {
	healthService := health.GetInstance()
	redisClient := monitors.NewRedisClient()

	healthService.Monitor(types.Monitor{
		Name: "Redis",
		Critical: true,
		Handler: redisClient,
	})

	ticker := time.NewTicker(5 * time.Second)

	for {
		select {
		case <-ticker.C:
			healthService.Check()
		}
	}
}
