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
	mongoClient := monitors.NewMongoClient()

	healthService.Monitor(types.Monitor{
		Name: "Redis",
		Critical: true,
		Handler: redisClient,
	})

	healthService.Monitor(types.Monitor{
		Name: "MongoDB",
		Critical: false,
		Handler: mongoClient,
	})

	ticker := time.NewTicker(5 * time.Second)

	for {
		select {
		case <-ticker.C:
			healthService.Check()
		}
	}
}
